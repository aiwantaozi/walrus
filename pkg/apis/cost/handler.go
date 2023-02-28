package cost

import (
	"fmt"
	"net/http"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqljson"
	"github.com/gin-gonic/gin"

	"github.com/seal-io/seal/pkg/apis/cost/view"
	"github.com/seal-io/seal/pkg/apis/runtime"
	"github.com/seal-io/seal/pkg/costs/distributor"
	"github.com/seal-io/seal/pkg/dao/model"
	"github.com/seal-io/seal/pkg/dao/model/allocationcost"
	"github.com/seal-io/seal/pkg/dao/model/clustercost"
	"github.com/seal-io/seal/pkg/dao/types"
)

func Handle(mc model.ClientSet) Handler {
	return Handler{
		modelClient: mc,
		distributor: distributor.New(mc),
	}
}

type Handler struct {
	modelClient model.ClientSet
	distributor *distributor.Distributor
}

func (h Handler) Kind() string {
	return "Cost"
}

func (h Handler) Validating() any {
	return h.modelClient
}

// Basic APIs

// Batch APIs

// Extensional APIs

func (h Handler) CollectionRouteAllocationCost(ctx *gin.Context, req view.AllocationCostRequest) (*runtime.ResponseCollection, error) {
	items, count, err := h.distributor.Distribute(ctx, req.StartTime, req.EndTime, req.QueryCondition)
	if err != nil {
		return nil, runtime.Errorf(http.StatusInternalServerError, "error query allocation cost: %w", err)
	}

	resp := runtime.GetResponseCollection(ctx, items, count)
	return &resp, nil
}

func (h Handler) CollectionRouteSummaryCost(ctx *gin.Context, req view.SummaryCostRequest) (*view.SummaryCostResponse, error) {
	// total
	var ps = []*sql.Predicate{
		sql.GTE(allocationcost.FieldStartTime, req.StartTime),
		sql.LTE(allocationcost.FieldEndTime, req.EndTime),
	}

	totalCost, err := h.modelClient.ClusterCosts().Query().
		Modify(func(s *sql.Selector) {
			s.Where(
				sql.And(ps...),
			).SelectExpr(
				sql.ExprFunc(func(b *sql.Builder) {
					b.WriteString(fmt.Sprintf(`COALESCE(SUM(%s),0)`, clustercost.FieldTotalCost))
				}),
			)
		}).
		Float64(ctx)
	if err != nil {
		return nil, fmt.Errorf("error summary total cost: %w", err)
	}

	// cluster
	clusters, err := h.modelClient.ClusterCosts().Query().
		Modify(func(s *sql.Selector) {
			s.Where(
				sql.And(ps...),
			).Select(
				clustercost.FieldConnectorID,
			).Distinct()
		}).Strings(ctx)
	if err != nil {
		return nil, fmt.Errorf("error summary each cluster's cost: %w", err)
	}

	// project
	var projects []struct {
		Value string `json:"value"`
	}
	ps = append(ps, sqljson.ValueIsNotNull(allocationcost.FieldLabels))
	err = h.modelClient.AllocationCosts().Query().
		Modify(func(s *sql.Selector) {
			s.Where(
				sql.And(ps...),
			).SelectExpr(
				sql.Expr(fmt.Sprintf(`DISTINCT(labels ->> '%s') AS value`, types.LabelSealProject)),
			)
		}).Scan(ctx, &projects)
	if err != nil {
		return nil, fmt.Errorf("error summary each project's cost: %w", err)
	}

	var projectCount int
	for _, v := range projects {
		if v.Value != "" {
			projectCount += 1
		}
	}

	// average
	averageDailyCost := averageDaily(req.StartTime, req.EndTime, totalCost)
	return &view.SummaryCostResponse{
		TotalCost:        totalCost,
		AverageDailyCost: averageDailyCost,
		ClusterCount:     len(clusters),
		ProjectCount:     projectCount,
	}, nil
}

func (h Handler) CollectionRouteSummaryClusterCost(ctx *gin.Context, req view.SummaryClusterCostRequest) (*view.SummaryClusterCostResponse, error) {
	var ps = []*sql.Predicate{
		sql.GTE(allocationcost.FieldStartTime, req.StartTime),
		sql.LTE(allocationcost.FieldEndTime, req.EndTime),
	}

	var s []view.SummaryClusterCostResponse
	err := h.modelClient.ClusterCosts().Query().
		Modify(func(s *sql.Selector) {
			s.Where(
				sql.And(ps...),
			).SelectExpr(
				sql.Expr(model.As(model.Sum(clustercost.FieldTotalCost), "totalCost")(s)),
				sql.Expr(model.As(model.Sum(clustercost.FieldManagementCost), "managementCost")(s)),
				sql.Expr(model.As(model.Sum(clustercost.FieldIdleCost), "idleCost")(s)),
				sql.Expr(model.As(model.Sum(clustercost.FieldAllocationCost), "allocationCost")(s)),
			)
		}).
		Scan(ctx, &s)
	if err != nil {
		return nil, fmt.Errorf("error summary cluster cost: %w", err)
	}

	if len(s) == 0 {
		return nil, nil
	}

	summary := s[0]
	summary.AverageDailyCost = averageDaily(req.StartTime, req.EndTime, summary.TotalCost)
	return &summary, nil
}

func (h Handler) CollectionRouteSummaryProjectCost(ctx *gin.Context, req view.SummaryProjectCostRequest) (*view.SummaryCostCommonResponse, error) {
	var ps = []*sql.Predicate{
		sql.GTE(allocationcost.FieldStartTime, req.StartTime),
		sql.LTE(allocationcost.FieldEndTime, req.EndTime),
		sqljson.ValueEQ(allocationcost.FieldLabels, req.Project, sqljson.Path(types.LabelSealProject)),
	}

	var s []view.SummaryCostCommonResponse
	err := h.modelClient.AllocationCosts().Query().
		Modify(func(s *sql.Selector) {
			s.Where(
				sql.And(ps...),
			).SelectExpr(
				sql.Expr(model.As(model.Sum(clustercost.FieldTotalCost), "totalCost")(s)),
			)
		}).
		Scan(ctx, &s)
	if err != nil {
		return nil, fmt.Errorf("error summary project cost: %w", err)
	}

	if len(s) == 0 {
		return nil, nil
	}

	summary := s[0]
	summary.AverageDailyCost = averageDaily(req.StartTime, req.EndTime, s[0].TotalCost)
	return &summary, nil
}

func (h Handler) CollectionRouteSummaryQueriedCost(ctx *gin.Context, req view.SummaryQueriedCostRequest) (*view.SummaryQueriedCostResponse, error) {
	cond := types.QueryCondition{
		Filters:     req.Filters,
		GroupBy:     types.GroupByFieldConnectorID,
		SharedCosts: req.SharedCosts,
	}

	items, count, err := h.distributor.Distribute(ctx, req.StartTime, req.EndTime, cond)
	if err != nil {
		return nil, runtime.Errorf(http.StatusInternalServerError, "error query allocation cost: %w", err)
	}

	summary := &view.SummaryQueriedCostResponse{}
	summary.ConnectorCount = count
	for _, v := range items {
		summary.TotalCost += v.TotalCost
		summary.TotalCost += v.TotalCost
		summary.SharedCost += v.SharedCost
		summary.CpuCost += v.CpuCost
		summary.GpuCost += v.GpuCost
		summary.RamCost += v.RamCost
		summary.PvCost += v.PvCost
	}

	return summary, nil
}

func averageDaily(startTime, endTime time.Time, total float64) float64 {
	if total == 0 {
		return 0
	}

	// average
	day := endTime.Sub(startTime).Hours() / 24.0
	if day == 0 {
		return 0
	}
	return total / day
}
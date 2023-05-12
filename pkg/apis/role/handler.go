package role

import (
	"github.com/gin-gonic/gin"

	"github.com/seal-io/seal/pkg/apis/role/view"
	"github.com/seal-io/seal/pkg/dao/model"
	"github.com/seal-io/seal/pkg/dao/model/predicate"
	"github.com/seal-io/seal/pkg/dao/model/role"
)

func Handle(mc model.ClientSet) Handler {
	return Handler{
		modelClient: mc,
	}
}

type Handler struct {
	modelClient model.ClientSet
}

func (h Handler) Kind() string {
	return "Role"
}

// Basic APIs.

// Batch APIs.

var (
	queryFields = []string{
		role.FieldName,
	}
	getFields = role.WithoutFields(
		role.FieldCreateTime,
		role.FieldUpdateTime)
	sortFields = []string{
		role.FieldName,
		role.FieldCreateTime,
	}
)

func (h Handler) CollectionGet(
	ctx *gin.Context,
	req view.CollectionGetRequest,
) (view.CollectionGetResponse, int, error) {
	// Do not export session level roles.
	input := []predicate.Role{
		role.Session(false),
	}
	if req.Domain != "" {
		input = append(input, role.Domain(req.Domain))
	}

	query := h.modelClient.Roles().Query().
		Where(input...)
	if queries, ok := req.Querying(queryFields); ok {
		query.Where(queries)
	}

	// Get count.
	cnt, err := query.Clone().Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	// Get entities.
	if limit, offset, ok := req.Paging(); ok {
		query.Limit(limit).Offset(offset)
	}

	if fields, ok := req.Extracting(getFields, getFields...); ok {
		query.Select(fields...)
	}

	if orders, ok := req.Sorting(sortFields, model.Desc(role.FieldCreateTime)); ok {
		query.Order(orders...)
	}

	entities, err := query.
		// Allow returning without sorting keys.
		Unique(false).
		All(ctx)
	if err != nil {
		return nil, 0, err
	}

	return model.ExposeRoles(entities), cnt, nil
}

// Extensional APIs.

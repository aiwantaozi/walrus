package project

import (
	"entgo.io/ent/dialect/sql"
	"github.com/gin-gonic/gin"

	"github.com/seal-io/seal/pkg/apis/project/view"
	"github.com/seal-io/seal/pkg/dao"
	"github.com/seal-io/seal/pkg/dao/model"
	"github.com/seal-io/seal/pkg/dao/model/project"
	"github.com/seal-io/seal/pkg/dao/model/secret"
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
	return "Project"
}

func (h Handler) Validating() any {
	return h.modelClient
}

// Basic APIs.

func (h Handler) Create(ctx *gin.Context, req view.CreateRequest) (view.CreateResponse, error) {
	entity := req.Model()

	creates, err := dao.ProjectCreates(h.modelClient, entity)
	if err != nil {
		return nil, err
	}

	entity, err = creates[0].Save(ctx)
	if err != nil {
		return nil, err
	}

	return model.ExposeProject(entity), nil
}

func (h Handler) Delete(ctx *gin.Context, req view.DeleteRequest) error {
	return h.modelClient.Projects().DeleteOne(req.Model()).Exec(ctx)
}

func (h Handler) Update(ctx *gin.Context, req view.UpdateRequest) error {
	entity := req.Model()

	updates, err := dao.ProjectUpdates(h.modelClient, entity)
	if err != nil {
		return err
	}

	return updates[0].Exec(ctx)
}

func (h Handler) Get(ctx *gin.Context, req view.GetRequest) (view.GetResponse, error) {
	entity, err := h.modelClient.Projects().Get(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return model.ExposeProject(entity), nil
}

// Batch APIs.

func (h Handler) CollectionDelete(ctx *gin.Context, req view.CollectionDeleteRequest) error {
	return h.modelClient.WithTx(ctx, func(tx *model.Tx) (err error) {
		for i := range req {
			err = tx.Projects().DeleteOne(req[i].Model()).
				Exec(ctx)
			if err != nil {
				return err
			}
		}

		return
	})
}

var (
	queryFields = []string{
		project.FieldName,
	}
	getFields = project.WithoutFields(
		project.FieldUpdateTime)
	sortFields = []string{
		project.FieldName,
		project.FieldCreateTime,
	}
)

func (h Handler) CollectionGet(
	ctx *gin.Context,
	req view.CollectionGetRequest,
) (view.CollectionGetResponse, int, error) {
	query := h.modelClient.Projects().Query()
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

	if orders, ok := req.Sorting(sortFields, model.Desc(project.FieldCreateTime)); ok {
		query.Order(orders...)
	}

	entities, err := query.
		// Allow returning without sorting keys.
		Unique(false).
		All(ctx)
	if err != nil {
		return nil, 0, err
	}

	return model.ExposeProjects(entities), cnt, nil
}

// Extensional APIs.

var secretQueryFields = []string{
	secret.FieldName,
}

func (h Handler) GetSecrets(
	ctx *gin.Context,
	req view.GetSecretsRequest,
) (view.GetSecretsResponse, error) {
	query := h.modelClient.Secrets().Query()
	if queries, ok := req.Querying(secretQueryFields); ok {
		query.Where(queries)
	}

	var entities model.Secrets

	err := query.
		Where(secret.Or(
			secret.ProjectIDIsNil(),
			secret.ProjectID(req.ID),
		)).
		Modify(func(s *sql.Selector) {
			s.Select(secret.FieldName).
				Distinct()
		}).
		Scan(ctx, &entities)
	if err != nil {
		return nil, err
	}

	return model.ExposeSecrets(entities), nil
}

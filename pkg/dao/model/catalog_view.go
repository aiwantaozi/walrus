// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus". DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/seal-io/walrus/pkg/dao/model/catalog"
	"github.com/seal-io/walrus/pkg/dao/model/predicate"
	"github.com/seal-io/walrus/pkg/dao/schema/intercept"
	"github.com/seal-io/walrus/pkg/dao/types"
	"github.com/seal-io/walrus/pkg/dao/types/object"
	"github.com/seal-io/walrus/pkg/dao/types/status"
)

// CatalogCreateInput holds the creation input of the Catalog entity,
// please tags with `path:",inline" json:",inline"` if embedding.
type CatalogCreateInput struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Project indicates to create Catalog entity CAN under the Project route.
	Project *ProjectQueryInput `path:",inline" query:"-" json:"-"`

	// Source of the catalog.
	Source string `path:"-" query:"-" json:"source"`
	// Type of the catalog.
	Type string `path:"-" query:"-" json:"type"`
	// Name holds the value of the "name" field.
	Name string `path:"-" query:"-" json:"name"`
	// Description holds the value of the "description" field.
	Description string `path:"-" query:"-" json:"description,omitempty"`
	// Labels holds the value of the "labels" field.
	Labels map[string]string `path:"-" query:"-" json:"labels,omitempty"`
}

// Model returns the Catalog entity for creating,
// after validating.
func (cci *CatalogCreateInput) Model() *Catalog {
	if cci == nil {
		return nil
	}

	_c := &Catalog{
		Source:      cci.Source,
		Type:        cci.Type,
		Name:        cci.Name,
		Description: cci.Description,
		Labels:      cci.Labels,
	}

	if cci.Project != nil {
		_c.ProjectID = cci.Project.ID
	}

	return _c
}

// Validate checks the CatalogCreateInput entity.
func (cci *CatalogCreateInput) Validate() error {
	if cci == nil {
		return errors.New("nil receiver")
	}

	return cci.ValidateWith(cci.inputConfig.Context, cci.inputConfig.Client, nil)
}

// ValidateWith checks the CatalogCreateInput entity with the given context and client set.
func (cci *CatalogCreateInput) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if cci == nil {
		return errors.New("nil receiver")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	// Validate when creating under the Project route.
	if cci.Project != nil {
		if err := cci.Project.ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				cci.Project = nil
			}
		}
	}

	return nil
}

// CatalogCreateInputs holds the creation input item of the Catalog entities.
type CatalogCreateInputsItem struct {
	// Source of the catalog.
	Source string `path:"-" query:"-" json:"source"`
	// Type of the catalog.
	Type string `path:"-" query:"-" json:"type"`
	// Name holds the value of the "name" field.
	Name string `path:"-" query:"-" json:"name"`
	// Description holds the value of the "description" field.
	Description string `path:"-" query:"-" json:"description,omitempty"`
	// Labels holds the value of the "labels" field.
	Labels map[string]string `path:"-" query:"-" json:"labels,omitempty"`
}

// ValidateWith checks the CatalogCreateInputsItem entity with the given context and client set.
func (cci *CatalogCreateInputsItem) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if cci == nil {
		return errors.New("nil receiver")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	return nil
}

// CatalogCreateInputs holds the creation input of the Catalog entities,
// please tags with `path:",inline" json:",inline"` if embedding.
type CatalogCreateInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Project indicates to create Catalog entity CAN under the Project route.
	Project *ProjectQueryInput `path:",inline" query:"-" json:"-"`

	// Items holds the entities to create, which MUST not be empty.
	Items []*CatalogCreateInputsItem `path:"-" query:"-" json:"items"`
}

// Model returns the Catalog entities for creating,
// after validating.
func (cci *CatalogCreateInputs) Model() []*Catalog {
	if cci == nil || len(cci.Items) == 0 {
		return nil
	}

	_cs := make([]*Catalog, len(cci.Items))

	for i := range cci.Items {
		_c := &Catalog{
			Source:      cci.Items[i].Source,
			Type:        cci.Items[i].Type,
			Name:        cci.Items[i].Name,
			Description: cci.Items[i].Description,
			Labels:      cci.Items[i].Labels,
		}

		if cci.Project != nil {
			_c.ProjectID = cci.Project.ID
		}

		_cs[i] = _c
	}

	return _cs
}

// Validate checks the CatalogCreateInputs entity .
func (cci *CatalogCreateInputs) Validate() error {
	if cci == nil {
		return errors.New("nil receiver")
	}

	return cci.ValidateWith(cci.inputConfig.Context, cci.inputConfig.Client, nil)
}

// ValidateWith checks the CatalogCreateInputs entity with the given context and client set.
func (cci *CatalogCreateInputs) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if cci == nil {
		return errors.New("nil receiver")
	}

	if len(cci.Items) == 0 {
		return errors.New("empty items")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	// Validate when creating under the Project route.
	if cci.Project != nil {
		if err := cci.Project.ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				cci.Project = nil
			}
		}
	}

	for i := range cci.Items {
		if cci.Items[i] == nil {
			continue
		}

		if err := cci.Items[i].ValidateWith(ctx, cs, cache); err != nil {
			return err
		}
	}

	return nil
}

// CatalogDeleteInput holds the deletion input of the Catalog entity,
// please tags with `path:",inline"` if embedding.
type CatalogDeleteInput struct {
	CatalogQueryInput `path:",inline"`
}

// CatalogDeleteInputs holds the deletion input item of the Catalog entities.
type CatalogDeleteInputsItem struct {
	// ID of the Catalog entity, tries to retrieve the entity with the following unique index parts if no ID provided.
	ID object.ID `path:"-" query:"-" json:"id,omitempty"`
	// Name of the Catalog entity, a part of the unique index.
	Name string `path:"-" query:"-" json:"name,omitempty"`
}

// CatalogDeleteInputs holds the deletion input of the Catalog entities,
// please tags with `path:",inline" json:",inline"` if embedding.
type CatalogDeleteInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Project indicates to delete Catalog entity CAN under the Project route.
	Project *ProjectQueryInput `path:",inline" query:"-" json:"-"`

	// Items holds the entities to create, which MUST not be empty.
	Items []*CatalogDeleteInputsItem `path:"-" query:"-" json:"items"`
}

// Model returns the Catalog entities for deleting,
// after validating.
func (cdi *CatalogDeleteInputs) Model() []*Catalog {
	if cdi == nil || len(cdi.Items) == 0 {
		return nil
	}

	_cs := make([]*Catalog, len(cdi.Items))
	for i := range cdi.Items {
		_cs[i] = &Catalog{
			ID: cdi.Items[i].ID,
		}
	}
	return _cs
}

// IDs returns the ID list of the Catalog entities for deleting,
// after validating.
func (cdi *CatalogDeleteInputs) IDs() []object.ID {
	if cdi == nil || len(cdi.Items) == 0 {
		return nil
	}

	ids := make([]object.ID, len(cdi.Items))
	for i := range cdi.Items {
		ids[i] = cdi.Items[i].ID
	}
	return ids
}

// Validate checks the CatalogDeleteInputs entity.
func (cdi *CatalogDeleteInputs) Validate() error {
	if cdi == nil {
		return errors.New("nil receiver")
	}

	return cdi.ValidateWith(cdi.inputConfig.Context, cdi.inputConfig.Client, nil)
}

// ValidateWith checks the CatalogDeleteInputs entity with the given context and client set.
func (cdi *CatalogDeleteInputs) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if cdi == nil {
		return errors.New("nil receiver")
	}

	if len(cdi.Items) == 0 {
		return errors.New("empty items")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	q := cs.Catalogs().Query()

	// Validate when deleting under the Project route.
	if cdi.Project != nil {
		if err := cdi.Project.ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				cdi.Project = nil
				q.Where(
					catalog.ProjectIDIsNil())
			}
		} else {
			ctx = valueContext(ctx, intercept.WithProjectInterceptor)
			q.Where(
				catalog.ProjectID(cdi.Project.ID))
		}
	} else {
		q.Where(
			catalog.ProjectIDIsNil())
	}

	ids := make([]object.ID, 0, len(cdi.Items))
	ors := make([]predicate.Catalog, 0, len(cdi.Items))
	indexers := make(map[any][]int)

	for i := range cdi.Items {
		if cdi.Items[i] == nil {
			return errors.New("nil item")
		}

		if cdi.Items[i].ID != "" {
			ids = append(ids, cdi.Items[i].ID)
			ors = append(ors, catalog.ID(cdi.Items[i].ID))
			indexers[cdi.Items[i].ID] = append(indexers[cdi.Items[i].ID], i)
		} else if cdi.Items[i].Name != "" {
			ors = append(ors, catalog.And(
				catalog.Name(cdi.Items[i].Name)))
			indexerKey := fmt.Sprint("/", cdi.Items[i].Name)
			indexers[indexerKey] = append(indexers[indexerKey], i)
		} else {
			return errors.New("found item hasn't identify")
		}
	}

	p := catalog.IDIn(ids...)
	if len(ids) != cap(ids) {
		p = catalog.Or(ors...)
	}

	es, err := q.
		Where(p).
		Select(
			catalog.FieldID,
			catalog.FieldName,
		).
		All(ctx)
	if err != nil {
		return err
	}

	if len(es) != cap(ids) {
		return errors.New("found unrecognized item")
	}

	for i := range es {
		indexer := indexers[es[i].ID]
		if indexer == nil {
			indexerKey := fmt.Sprint("/", es[i].Name)
			indexer = indexers[indexerKey]
		}
		for _, j := range indexer {
			cdi.Items[j].ID = es[i].ID
			cdi.Items[j].Name = es[i].Name
		}
	}

	return nil
}

// CatalogQueryInput holds the query input of the Catalog entity,
// please tags with `path:",inline"` if embedding.
type CatalogQueryInput struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Project indicates to query Catalog entity CAN under the Project route.
	Project *ProjectQueryInput `path:",inline" query:"-" json:"project,omitempty"`

	// Refer holds the route path reference of the Catalog entity.
	Refer *object.Refer `path:"catalog,default=" query:"-" json:"-"`
	// ID of the Catalog entity, tries to retrieve the entity with the following unique index parts if no ID provided.
	ID object.ID `path:"-" query:"-" json:"id,omitempty"`
	// Name of the Catalog entity, a part of the unique index.
	Name string `path:"-" query:"-" json:"name,omitempty"`
}

// Model returns the Catalog entity for querying,
// after validating.
func (cqi *CatalogQueryInput) Model() *Catalog {
	if cqi == nil {
		return nil
	}

	return &Catalog{
		ID:   cqi.ID,
		Name: cqi.Name,
	}
}

// Validate checks the CatalogQueryInput entity.
func (cqi *CatalogQueryInput) Validate() error {
	if cqi == nil {
		return errors.New("nil receiver")
	}

	return cqi.ValidateWith(cqi.inputConfig.Context, cqi.inputConfig.Client, nil)
}

// ValidateWith checks the CatalogQueryInput entity with the given context and client set.
func (cqi *CatalogQueryInput) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if cqi == nil {
		return errors.New("nil receiver")
	}

	if cqi.Refer != nil && *cqi.Refer == "" {
		return fmt.Errorf("model: %s : %w", catalog.Label, ErrBlankResourceRefer)
	}

	if cache == nil {
		cache = map[string]any{}
	}

	q := cs.Catalogs().Query()

	// Validate when querying under the Project route.
	if cqi.Project != nil {
		if err := cqi.Project.ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				cqi.Project = nil
				q.Where(
					catalog.ProjectIDIsNil())
			}
		} else {
			ctx = valueContext(ctx, intercept.WithProjectInterceptor)
			q.Where(
				catalog.ProjectID(cqi.Project.ID))
		}
	} else {
		q.Where(
			catalog.ProjectIDIsNil())
	}

	if cqi.Refer != nil {
		if cqi.Refer.IsID() {
			q.Where(
				catalog.ID(cqi.Refer.ID()))
		} else if refers := cqi.Refer.Split(1); len(refers) == 1 {
			q.Where(
				catalog.Name(refers[0].String()))
		} else {
			return errors.New("invalid identify refer of catalog")
		}
	} else if cqi.ID != "" {
		q.Where(
			catalog.ID(cqi.ID))
	} else if cqi.Name != "" {
		q.Where(
			catalog.Name(cqi.Name))
	} else {
		return errors.New("invalid identify of catalog")
	}

	q.Select(
		catalog.FieldID,
		catalog.FieldName,
	)

	var e *Catalog
	{
		// Get cache from previous validation.
		queryStmt, queryArgs := q.sqlQuery(setContextOp(ctx, q.ctx, "cache")).Query()
		ck := fmt.Sprintf("stmt=%v, args=%v", queryStmt, queryArgs)
		if cv, existed := cache[ck]; !existed {
			var err error
			e, err = q.Only(ctx)
			if err != nil {
				return err
			}

			// Set cache for other validation.
			cache[ck] = e
		} else {
			e = cv.(*Catalog)
		}
	}

	cqi.ID = e.ID
	cqi.Name = e.Name
	return nil
}

// CatalogQueryInputs holds the query input of the Catalog entities,
// please tags with `path:",inline" query:",inline"` if embedding.
type CatalogQueryInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Project indicates to query Catalog entity CAN under the Project route.
	Project *ProjectQueryInput `path:",inline" query:"-" json:"-"`
}

// Validate checks the CatalogQueryInputs entity.
func (cqi *CatalogQueryInputs) Validate() error {
	if cqi == nil {
		return errors.New("nil receiver")
	}

	return cqi.ValidateWith(cqi.inputConfig.Context, cqi.inputConfig.Client, nil)
}

// ValidateWith checks the CatalogQueryInputs entity with the given context and client set.
func (cqi *CatalogQueryInputs) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if cqi == nil {
		return errors.New("nil receiver")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	// Validate when querying under the Project route.
	if cqi.Project != nil {
		if err := cqi.Project.ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				cqi.Project = nil
			}
		}
	}

	return nil
}

// CatalogUpdateInput holds the modification input of the Catalog entity,
// please tags with `path:",inline" json:",inline"` if embedding.
type CatalogUpdateInput struct {
	CatalogQueryInput `path:",inline" query:"-" json:"-"`

	// Description holds the value of the "description" field.
	Description string `path:"-" query:"-" json:"description,omitempty"`
	// Labels holds the value of the "labels" field.
	Labels map[string]string `path:"-" query:"-" json:"labels,omitempty"`
}

// Model returns the Catalog entity for modifying,
// after validating.
func (cui *CatalogUpdateInput) Model() *Catalog {
	if cui == nil {
		return nil
	}

	_c := &Catalog{
		ID:          cui.ID,
		Name:        cui.Name,
		Description: cui.Description,
		Labels:      cui.Labels,
	}

	return _c
}

// Validate checks the CatalogUpdateInput entity.
func (cui *CatalogUpdateInput) Validate() error {
	if cui == nil {
		return errors.New("nil receiver")
	}

	return cui.ValidateWith(cui.inputConfig.Context, cui.inputConfig.Client, nil)
}

// ValidateWith checks the CatalogUpdateInput entity with the given context and client set.
func (cui *CatalogUpdateInput) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if cache == nil {
		cache = map[string]any{}
	}

	if err := cui.CatalogQueryInput.ValidateWith(ctx, cs, cache); err != nil {
		return err
	}

	return nil
}

// CatalogUpdateInputs holds the modification input item of the Catalog entities.
type CatalogUpdateInputsItem struct {
	// ID of the Catalog entity, tries to retrieve the entity with the following unique index parts if no ID provided.
	ID object.ID `path:"-" query:"-" json:"id,omitempty"`
	// Name of the Catalog entity, a part of the unique index.
	Name string `path:"-" query:"-" json:"name,omitempty"`

	// Description holds the value of the "description" field.
	Description string `path:"-" query:"-" json:"description,omitempty"`
	// Labels holds the value of the "labels" field.
	Labels map[string]string `path:"-" query:"-" json:"labels,omitempty"`
}

// ValidateWith checks the CatalogUpdateInputsItem entity with the given context and client set.
func (cui *CatalogUpdateInputsItem) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if cui == nil {
		return errors.New("nil receiver")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	return nil
}

// CatalogUpdateInputs holds the modification input of the Catalog entities,
// please tags with `path:",inline" json:",inline"` if embedding.
type CatalogUpdateInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Project indicates to update Catalog entity CAN under the Project route.
	Project *ProjectQueryInput `path:",inline" query:"-" json:"-"`

	// Items holds the entities to create, which MUST not be empty.
	Items []*CatalogUpdateInputsItem `path:"-" query:"-" json:"items"`
}

// Model returns the Catalog entities for modifying,
// after validating.
func (cui *CatalogUpdateInputs) Model() []*Catalog {
	if cui == nil || len(cui.Items) == 0 {
		return nil
	}

	_cs := make([]*Catalog, len(cui.Items))

	for i := range cui.Items {
		_c := &Catalog{
			ID:          cui.Items[i].ID,
			Name:        cui.Items[i].Name,
			Description: cui.Items[i].Description,
			Labels:      cui.Items[i].Labels,
		}

		_cs[i] = _c
	}

	return _cs
}

// IDs returns the ID list of the Catalog entities for modifying,
// after validating.
func (cui *CatalogUpdateInputs) IDs() []object.ID {
	if cui == nil || len(cui.Items) == 0 {
		return nil
	}

	ids := make([]object.ID, len(cui.Items))
	for i := range cui.Items {
		ids[i] = cui.Items[i].ID
	}
	return ids
}

// Validate checks the CatalogUpdateInputs entity.
func (cui *CatalogUpdateInputs) Validate() error {
	if cui == nil {
		return errors.New("nil receiver")
	}

	return cui.ValidateWith(cui.inputConfig.Context, cui.inputConfig.Client, nil)
}

// ValidateWith checks the CatalogUpdateInputs entity with the given context and client set.
func (cui *CatalogUpdateInputs) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if cui == nil {
		return errors.New("nil receiver")
	}

	if len(cui.Items) == 0 {
		return errors.New("empty items")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	q := cs.Catalogs().Query()

	// Validate when updating under the Project route.
	if cui.Project != nil {
		if err := cui.Project.ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				cui.Project = nil
				q.Where(
					catalog.ProjectIDIsNil())
			}
		} else {
			ctx = valueContext(ctx, intercept.WithProjectInterceptor)
			q.Where(
				catalog.ProjectID(cui.Project.ID))
		}
	} else {
		q.Where(
			catalog.ProjectIDIsNil())
	}

	ids := make([]object.ID, 0, len(cui.Items))
	ors := make([]predicate.Catalog, 0, len(cui.Items))
	indexers := make(map[any][]int)

	for i := range cui.Items {
		if cui.Items[i] == nil {
			return errors.New("nil item")
		}

		if cui.Items[i].ID != "" {
			ids = append(ids, cui.Items[i].ID)
			ors = append(ors, catalog.ID(cui.Items[i].ID))
			indexers[cui.Items[i].ID] = append(indexers[cui.Items[i].ID], i)
		} else if cui.Items[i].Name != "" {
			ors = append(ors, catalog.And(
				catalog.Name(cui.Items[i].Name)))
			indexerKey := fmt.Sprint("/", cui.Items[i].Name)
			indexers[indexerKey] = append(indexers[indexerKey], i)
		} else {
			return errors.New("found item hasn't identify")
		}
	}

	p := catalog.IDIn(ids...)
	if len(ids) != cap(ids) {
		p = catalog.Or(ors...)
	}

	es, err := q.
		Where(p).
		Select(
			catalog.FieldID,
			catalog.FieldName,
		).
		All(ctx)
	if err != nil {
		return err
	}

	if len(es) != cap(ids) {
		return errors.New("found unrecognized item")
	}

	for i := range es {
		indexer := indexers[es[i].ID]
		if indexer == nil {
			indexerKey := fmt.Sprint("/", es[i].Name)
			indexer = indexers[indexerKey]
		}
		for _, j := range indexer {
			cui.Items[j].ID = es[i].ID
			cui.Items[j].Name = es[i].Name
		}
	}

	for i := range cui.Items {
		if err := cui.Items[i].ValidateWith(ctx, cs, cache); err != nil {
			return err
		}
	}

	return nil
}

// CatalogOutput holds the output of the Catalog entity.
type CatalogOutput struct {
	ID          object.ID          `json:"id,omitempty"`
	Name        string             `json:"name,omitempty"`
	Description string             `json:"description,omitempty"`
	Labels      map[string]string  `json:"labels,omitempty"`
	CreateTime  *time.Time         `json:"createTime,omitempty"`
	UpdateTime  *time.Time         `json:"updateTime,omitempty"`
	Status      status.Status      `json:"status,omitempty"`
	Type        string             `json:"type,omitempty"`
	Source      string             `json:"source,omitempty"`
	Sync        *types.CatalogSync `json:"sync,omitempty"`

	Project *ProjectOutput `json:"project,omitempty"`
}

// View returns the output of Catalog entity.
func (_c *Catalog) View() *CatalogOutput {
	return ExposeCatalog(_c)
}

// View returns the output of Catalog entities.
func (_cs Catalogs) View() []*CatalogOutput {
	return ExposeCatalogs(_cs)
}

// ExposeCatalog converts the Catalog to CatalogOutput.
func ExposeCatalog(_c *Catalog) *CatalogOutput {
	if _c == nil {
		return nil
	}

	co := &CatalogOutput{
		ID:          _c.ID,
		Name:        _c.Name,
		Description: _c.Description,
		Labels:      _c.Labels,
		CreateTime:  _c.CreateTime,
		UpdateTime:  _c.UpdateTime,
		Status:      _c.Status,
		Type:        _c.Type,
		Source:      _c.Source,
		Sync:        _c.Sync,
	}

	if _c.Edges.Project != nil {
		co.Project = ExposeProject(_c.Edges.Project)
	} else if _c.ProjectID != "" {
		co.Project = &ProjectOutput{
			ID: _c.ProjectID,
		}
	}
	return co
}

// ExposeCatalogs converts the Catalog slice to CatalogOutput pointer slice.
func ExposeCatalogs(_cs []*Catalog) []*CatalogOutput {
	if len(_cs) == 0 {
		return nil
	}

	cos := make([]*CatalogOutput, len(_cs))
	for i := range _cs {
		cos[i] = ExposeCatalog(_cs[i])
	}
	return cos
}

// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "seal". DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/seal-io/seal/pkg/dao/model/predicate"
	"github.com/seal-io/seal/pkg/dao/model/project"
	"github.com/seal-io/seal/pkg/dao/schema/intercept"
	"github.com/seal-io/seal/pkg/dao/types/object"
)

// ProjectCreateInput holds the creation input of the Project entity,
// please tags with `path:",inline" json:",inline"` if embedding.
type ProjectCreateInput struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Name holds the value of the "name" field.
	Name string `path:"-" query:"-" json:"name"`
	// Description holds the value of the "description" field.
	Description string `path:"-" query:"-" json:"description,omitempty"`
	// Labels holds the value of the "labels" field.
	Labels map[string]string `path:"-" query:"-" json:"labels,omitempty"`
}

// Model returns the Project entity for creating,
// after validating.
func (pci *ProjectCreateInput) Model() *Project {
	if pci == nil {
		return nil
	}

	_p := &Project{
		Name:        pci.Name,
		Description: pci.Description,
		Labels:      pci.Labels,
	}

	return _p
}

// Validate checks the ProjectCreateInput entity.
func (pci *ProjectCreateInput) Validate() error {
	if pci == nil {
		return errors.New("nil receiver")
	}

	return pci.ValidateWith(pci.inputConfig.Context, pci.inputConfig.Client)
}

// ValidateWith checks the ProjectCreateInput entity with the given context and client set.
func (pci *ProjectCreateInput) ValidateWith(ctx context.Context, cs ClientSet) error {
	if pci == nil {
		return errors.New("nil receiver")
	}

	return nil
}

// ProjectCreateInputs holds the creation input item of the Project entities.
type ProjectCreateInputsItem struct {
	// Name holds the value of the "name" field.
	Name string `path:"-" query:"-" json:"name"`
	// Description holds the value of the "description" field.
	Description string `path:"-" query:"-" json:"description,omitempty"`
	// Labels holds the value of the "labels" field.
	Labels map[string]string `path:"-" query:"-" json:"labels,omitempty"`
}

// ValidateWith checks the ProjectCreateInputsItem entity with the given context and client set.
func (pci *ProjectCreateInputsItem) ValidateWith(ctx context.Context, cs ClientSet) error {
	if pci == nil {
		return errors.New("nil receiver")
	}

	return nil
}

// ProjectCreateInputs holds the creation input of the Project entities,
// please tags with `path:",inline" json:",inline"` if embedding.
type ProjectCreateInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Items holds the entities to create, which MUST not be empty.
	Items []*ProjectCreateInputsItem `path:"-" query:"-" json:"items"`
}

// Model returns the Project entities for creating,
// after validating.
func (pci *ProjectCreateInputs) Model() []*Project {
	if pci == nil || len(pci.Items) == 0 {
		return nil
	}

	_ps := make([]*Project, len(pci.Items))

	for i := range pci.Items {
		_p := &Project{
			Name:        pci.Items[i].Name,
			Description: pci.Items[i].Description,
			Labels:      pci.Items[i].Labels,
		}

		_ps[i] = _p
	}

	return _ps
}

// Validate checks the ProjectCreateInputs entity .
func (pci *ProjectCreateInputs) Validate() error {
	if pci == nil {
		return errors.New("nil receiver")
	}

	return pci.ValidateWith(pci.inputConfig.Context, pci.inputConfig.Client)
}

// ValidateWith checks the ProjectCreateInputs entity with the given context and client set.
func (pci *ProjectCreateInputs) ValidateWith(ctx context.Context, cs ClientSet) error {
	if pci == nil {
		return errors.New("nil receiver")
	}

	if len(pci.Items) == 0 {
		return errors.New("empty items")
	}

	for i := range pci.Items {
		if pci.Items[i] == nil {
			continue
		}

		if err := pci.Items[i].ValidateWith(ctx, cs); err != nil {
			return err
		}
	}

	return nil
}

// ProjectDeleteInput holds the deletion input of the Project entity,
// please tags with `path:",inline"` if embedding.
type ProjectDeleteInput = ProjectQueryInput

// ProjectDeleteInputs holds the deletion input item of the Project entities.
type ProjectDeleteInputsItem struct {
	// ID of the Project entity, tries to retrieve the entity with the following unique index parts if no ID provided.
	ID object.ID `path:"-" query:"-" json:"id,omitempty"`
	// Name of the Project entity, a part of the unique index.
	Name string `path:"-" query:"-" json:"name,omitempty"`
}

// ProjectDeleteInputs holds the deletion input of the Project entities,
// please tags with `path:",inline" json:",inline"` if embedding.
type ProjectDeleteInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Items holds the entities to create, which MUST not be empty.
	Items []*ProjectDeleteInputsItem `path:"-" query:"-" json:"items"`
}

// Model returns the Project entities for deleting,
// after validating.
func (pdi *ProjectDeleteInputs) Model() []*Project {
	if pdi == nil || len(pdi.Items) == 0 {
		return nil
	}

	_ps := make([]*Project, len(pdi.Items))
	for i := range pdi.Items {
		_ps[i] = &Project{
			ID: pdi.Items[i].ID,
		}
	}
	return _ps
}

// IDs returns the ID list of the Project entities for deleting,
// after validating.
func (pdi *ProjectDeleteInputs) IDs() []object.ID {
	if pdi == nil || len(pdi.Items) == 0 {
		return nil
	}

	ids := make([]object.ID, len(pdi.Items))
	for i := range pdi.Items {
		ids[i] = pdi.Items[i].ID
	}
	return ids
}

// Validate checks the ProjectDeleteInputs entity.
func (pdi *ProjectDeleteInputs) Validate() error {
	if pdi == nil {
		return errors.New("nil receiver")
	}

	return pdi.ValidateWith(pdi.inputConfig.Context, pdi.inputConfig.Client)
}

// ValidateWith checks the ProjectDeleteInputs entity with the given context and client set.
func (pdi *ProjectDeleteInputs) ValidateWith(ctx context.Context, cs ClientSet) error {
	if pdi == nil {
		return errors.New("nil receiver")
	}

	if len(pdi.Items) == 0 {
		return errors.New("empty items")
	}

	q := cs.Projects().Query()

	ids := make([]object.ID, 0, len(pdi.Items))
	ors := make([]predicate.Project, 0, len(pdi.Items))

	for i := range pdi.Items {
		if pdi.Items[i] == nil {
			return errors.New("nil item")
		}

		if pdi.Items[i].ID != "" {
			ids = append(ids, pdi.Items[i].ID)
			ors = append(ors, project.ID(pdi.Items[i].ID))
		} else if pdi.Items[i].Name != "" {
			ors = append(ors, project.And(
				project.Name(pdi.Items[i].Name)))
		} else {
			return errors.New("found item hasn't identify")
		}
	}

	ctx = valueContext(ctx, intercept.WithProjectInterceptor)

	p := project.IDIn(ids...)
	if len(ids) != cap(ids) {
		p = project.Or(ors...)
	}

	es, err := q.
		Where(p).
		Select(
			project.FieldID,
			project.FieldName,
		).
		All(ctx)
	if err != nil {
		return err
	}

	if len(es) != cap(ids) {
		return errors.New("found unrecognized item")
	}

	for i := range es {
		pdi.Items[i].ID = es[i].ID
		pdi.Items[i].Name = es[i].Name
	}

	return nil
}

// ProjectQueryInput holds the query input of the Project entity,
// please tags with `path:",inline"` if embedding.
type ProjectQueryInput struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Refer holds the route path reference of the Project entity.
	Refer *object.Refer `path:"project,default=" query:"-" json:"-"`
	// ID of the Project entity, tries to retrieve the entity with the following unique index parts if no ID provided.
	ID object.ID `path:"-" query:"-" json:"id,omitempty"`
	// Name of the Project entity, a part of the unique index.
	Name string `path:"-" query:"-" json:"name,omitempty"`
}

// Model returns the Project entity for querying,
// after validating.
func (pqi *ProjectQueryInput) Model() *Project {
	if pqi == nil {
		return nil
	}

	return &Project{
		ID:   pqi.ID,
		Name: pqi.Name,
	}
}

// Validate checks the ProjectQueryInput entity.
func (pqi *ProjectQueryInput) Validate() error {
	if pqi == nil {
		return errors.New("nil receiver")
	}

	return pqi.ValidateWith(pqi.inputConfig.Context, pqi.inputConfig.Client)
}

// ValidateWith checks the ProjectQueryInput entity with the given context and client set.
func (pqi *ProjectQueryInput) ValidateWith(ctx context.Context, cs ClientSet) error {
	if pqi == nil {
		return errors.New("nil receiver")
	}

	if pqi.Refer != nil && *pqi.Refer == "" {
		return fmt.Errorf("model: %s : %w", project.Label, ErrBlankResourceRefer)
	}

	q := cs.Projects().Query()

	if pqi.Refer != nil {
		if pqi.Refer.IsID() {
			q.Where(
				project.ID(pqi.Refer.ID()))
		} else if refers := pqi.Refer.Split(1); len(refers) == 1 {
			q.Where(
				project.Name(refers[0].String()))
		} else {
			return errors.New("invalid identify refer of project")
		}
	} else if pqi.ID != "" {
		q.Where(
			project.ID(pqi.ID))
	} else if pqi.Name != "" {
		q.Where(
			project.Name(pqi.Name))
	} else {
		return errors.New("invalid identify of project")
	}

	ctx = valueContext(ctx, intercept.WithProjectInterceptor)

	e, err := q.
		Select(
			project.FieldID,
			project.FieldName,
		).
		Only(ctx)
	if err == nil {
		pqi.ID = e.ID
		pqi.Name = e.Name
	}
	return err
}

// ProjectQueryInputs holds the query input of the Project entities,
// please tags with `path:",inline" query:",inline"` if embedding.
type ProjectQueryInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`
}

// Validate checks the ProjectQueryInputs entity.
func (pqi *ProjectQueryInputs) Validate() error {
	if pqi == nil {
		return errors.New("nil receiver")
	}

	return pqi.ValidateWith(pqi.inputConfig.Context, pqi.inputConfig.Client)
}

// ValidateWith checks the ProjectQueryInputs entity with the given context and client set.
func (pqi *ProjectQueryInputs) ValidateWith(ctx context.Context, cs ClientSet) error {
	if pqi == nil {
		return errors.New("nil receiver")
	}

	return nil
}

// ProjectUpdateInput holds the modification input of the Project entity,
// please tags with `path:",inline" json:",inline"` if embedding.
type ProjectUpdateInput struct {
	ProjectQueryInput `path:",inline" query:"-" json:"-"`

	// Description holds the value of the "description" field.
	Description string `path:"-" query:"-" json:"description,omitempty"`
	// Labels holds the value of the "labels" field.
	Labels map[string]string `path:"-" query:"-" json:"labels,omitempty"`
}

// Model returns the Project entity for modifying,
// after validating.
func (pui *ProjectUpdateInput) Model() *Project {
	if pui == nil {
		return nil
	}

	_p := &Project{
		ID:          pui.ID,
		Description: pui.Description,
		Labels:      pui.Labels,
	}

	return _p
}

// Validate checks the ProjectUpdateInput entity.
func (pui *ProjectUpdateInput) Validate() error {
	if pui == nil {
		return errors.New("nil receiver")
	}

	return pui.ValidateWith(pui.inputConfig.Context, pui.inputConfig.Client)
}

// ValidateWith checks the ProjectUpdateInput entity with the given context and client set.
func (pui *ProjectUpdateInput) ValidateWith(ctx context.Context, cs ClientSet) error {
	if err := pui.ProjectQueryInput.ValidateWith(ctx, cs); err != nil {
		return err
	}

	return nil
}

// ProjectUpdateInputs holds the modification input item of the Project entities.
type ProjectUpdateInputsItem struct {
	// ID of the Project entity, tries to retrieve the entity with the following unique index parts if no ID provided.
	ID object.ID `path:"-" query:"-" json:"id,omitempty"`
	// Name of the Project entity, a part of the unique index.
	Name string `path:"-" query:"-" json:"name,omitempty"`

	// Description holds the value of the "description" field.
	Description string `path:"-" query:"-" json:"description,omitempty"`
	// Labels holds the value of the "labels" field.
	Labels map[string]string `path:"-" query:"-" json:"labels,omitempty"`
}

// ValidateWith checks the ProjectUpdateInputsItem entity with the given context and client set.
func (pui *ProjectUpdateInputsItem) ValidateWith(ctx context.Context, cs ClientSet) error {
	if pui == nil {
		return errors.New("nil receiver")
	}

	return nil
}

// ProjectUpdateInputs holds the modification input of the Project entities,
// please tags with `path:",inline" json:",inline"` if embedding.
type ProjectUpdateInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Items holds the entities to create, which MUST not be empty.
	Items []*ProjectUpdateInputsItem `path:"-" query:"-" json:"items"`
}

// Model returns the Project entities for modifying,
// after validating.
func (pui *ProjectUpdateInputs) Model() []*Project {
	if pui == nil || len(pui.Items) == 0 {
		return nil
	}

	_ps := make([]*Project, len(pui.Items))

	for i := range pui.Items {
		_p := &Project{
			ID:          pui.Items[i].ID,
			Description: pui.Items[i].Description,
			Labels:      pui.Items[i].Labels,
		}

		_ps[i] = _p
	}

	return _ps
}

// IDs returns the ID list of the Project entities for modifying,
// after validating.
func (pui *ProjectUpdateInputs) IDs() []object.ID {
	if pui == nil || len(pui.Items) == 0 {
		return nil
	}

	ids := make([]object.ID, len(pui.Items))
	for i := range pui.Items {
		ids[i] = pui.Items[i].ID
	}
	return ids
}

// Validate checks the ProjectUpdateInputs entity.
func (pui *ProjectUpdateInputs) Validate() error {
	if pui == nil {
		return errors.New("nil receiver")
	}

	return pui.ValidateWith(pui.inputConfig.Context, pui.inputConfig.Client)
}

// ValidateWith checks the ProjectUpdateInputs entity with the given context and client set.
func (pui *ProjectUpdateInputs) ValidateWith(ctx context.Context, cs ClientSet) error {
	if pui == nil {
		return errors.New("nil receiver")
	}

	if len(pui.Items) == 0 {
		return errors.New("empty items")
	}

	q := cs.Projects().Query()

	ids := make([]object.ID, 0, len(pui.Items))
	ors := make([]predicate.Project, 0, len(pui.Items))

	for i := range pui.Items {
		if pui.Items[i] == nil {
			return errors.New("nil item")
		}

		if pui.Items[i].ID != "" {
			ids = append(ids, pui.Items[i].ID)
			ors = append(ors, project.ID(pui.Items[i].ID))
		} else if pui.Items[i].Name != "" {
			ors = append(ors, project.And(
				project.Name(pui.Items[i].Name)))
		} else {
			return errors.New("found item hasn't identify")
		}
	}

	ctx = valueContext(ctx, intercept.WithProjectInterceptor)

	p := project.IDIn(ids...)
	if len(ids) != cap(ids) {
		p = project.Or(ors...)
	}

	es, err := q.
		Where(p).
		Select(
			project.FieldID,
			project.FieldName,
		).
		All(ctx)
	if err != nil {
		return err
	}

	if len(es) != cap(ids) {
		return errors.New("found unrecognized item")
	}

	for i := range es {
		pui.Items[i].ID = es[i].ID
		pui.Items[i].Name = es[i].Name
	}

	for i := range pui.Items {
		if pui.Items[i] == nil {
			continue
		}

		if err := pui.Items[i].ValidateWith(ctx, cs); err != nil {
			return err
		}
	}

	return nil
}

// ProjectOutput holds the output of the Project entity.
type ProjectOutput struct {
	ID          object.ID         `json:"id,omitempty"`
	Name        string            `json:"name,omitempty"`
	Description string            `json:"description,omitempty"`
	Labels      map[string]string `json:"labels,omitempty"`
	CreateTime  *time.Time        `json:"createTime,omitempty"`
	UpdateTime  *time.Time        `json:"updateTime,omitempty"`
}

// View returns the output of Project entity.
func (_p *Project) View() *ProjectOutput {
	return ExposeProject(_p)
}

// View returns the output of Project entities.
func (_ps Projects) View() []*ProjectOutput {
	return ExposeProjects(_ps)
}

// ExposeProject converts the Project to ProjectOutput.
func ExposeProject(_p *Project) *ProjectOutput {
	if _p == nil {
		return nil
	}

	po := &ProjectOutput{
		ID:          _p.ID,
		Name:        _p.Name,
		Description: _p.Description,
		Labels:      _p.Labels,
		CreateTime:  _p.CreateTime,
		UpdateTime:  _p.UpdateTime,
	}

	return po
}

// ExposeProjects converts the Project slice to ProjectOutput pointer slice.
func ExposeProjects(_ps []*Project) []*ProjectOutput {
	if len(_ps) == 0 {
		return nil
	}

	pos := make([]*ProjectOutput, len(_ps))
	for i := range _ps {
		pos[i] = ExposeProject(_ps[i])
	}
	return pos
}
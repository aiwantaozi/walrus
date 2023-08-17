// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus". DO NOT EDIT.

package model

import (
	"context"
	stdsql "database/sql"
	"errors"
	"fmt"
	"reflect"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"

	"github.com/seal-io/walrus/pkg/dao/model/internal"
	"github.com/seal-io/walrus/pkg/dao/model/predicate"
	"github.com/seal-io/walrus/pkg/dao/model/template"
	"github.com/seal-io/walrus/pkg/dao/model/templateversion"
	"github.com/seal-io/walrus/pkg/dao/types/object"
	"github.com/seal-io/walrus/pkg/dao/types/status"
)

// TemplateUpdate is the builder for updating Template entities.
type TemplateUpdate struct {
	config
	hooks     []Hook
	mutation  *TemplateMutation
	modifiers []func(*sql.UpdateBuilder)
	object    *Template
}

// Where appends a list predicates to the TemplateUpdate builder.
func (tu *TemplateUpdate) Where(ps ...predicate.Template) *TemplateUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetDescription sets the "description" field.
func (tu *TemplateUpdate) SetDescription(s string) *TemplateUpdate {
	tu.mutation.SetDescription(s)
	return tu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (tu *TemplateUpdate) SetNillableDescription(s *string) *TemplateUpdate {
	if s != nil {
		tu.SetDescription(*s)
	}
	return tu
}

// ClearDescription clears the value of the "description" field.
func (tu *TemplateUpdate) ClearDescription() *TemplateUpdate {
	tu.mutation.ClearDescription()
	return tu
}

// SetLabels sets the "labels" field.
func (tu *TemplateUpdate) SetLabels(m map[string]string) *TemplateUpdate {
	tu.mutation.SetLabels(m)
	return tu
}

// ClearLabels clears the value of the "labels" field.
func (tu *TemplateUpdate) ClearLabels() *TemplateUpdate {
	tu.mutation.ClearLabels()
	return tu
}

// SetUpdateTime sets the "update_time" field.
func (tu *TemplateUpdate) SetUpdateTime(t time.Time) *TemplateUpdate {
	tu.mutation.SetUpdateTime(t)
	return tu
}

// SetStatus sets the "status" field.
func (tu *TemplateUpdate) SetStatus(s status.Status) *TemplateUpdate {
	tu.mutation.SetStatus(s)
	return tu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (tu *TemplateUpdate) SetNillableStatus(s *status.Status) *TemplateUpdate {
	if s != nil {
		tu.SetStatus(*s)
	}
	return tu
}

// ClearStatus clears the value of the "status" field.
func (tu *TemplateUpdate) ClearStatus() *TemplateUpdate {
	tu.mutation.ClearStatus()
	return tu
}

// SetIcon sets the "icon" field.
func (tu *TemplateUpdate) SetIcon(s string) *TemplateUpdate {
	tu.mutation.SetIcon(s)
	return tu
}

// SetNillableIcon sets the "icon" field if the given value is not nil.
func (tu *TemplateUpdate) SetNillableIcon(s *string) *TemplateUpdate {
	if s != nil {
		tu.SetIcon(*s)
	}
	return tu
}

// ClearIcon clears the value of the "icon" field.
func (tu *TemplateUpdate) ClearIcon() *TemplateUpdate {
	tu.mutation.ClearIcon()
	return tu
}

// SetSource sets the "source" field.
func (tu *TemplateUpdate) SetSource(s string) *TemplateUpdate {
	tu.mutation.SetSource(s)
	return tu
}

// AddVersionIDs adds the "versions" edge to the TemplateVersion entity by IDs.
func (tu *TemplateUpdate) AddVersionIDs(ids ...object.ID) *TemplateUpdate {
	tu.mutation.AddVersionIDs(ids...)
	return tu
}

// AddVersions adds the "versions" edges to the TemplateVersion entity.
func (tu *TemplateUpdate) AddVersions(t ...*TemplateVersion) *TemplateUpdate {
	ids := make([]object.ID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.AddVersionIDs(ids...)
}

// Mutation returns the TemplateMutation object of the builder.
func (tu *TemplateUpdate) Mutation() *TemplateMutation {
	return tu.mutation
}

// ClearVersions clears all "versions" edges to the TemplateVersion entity.
func (tu *TemplateUpdate) ClearVersions() *TemplateUpdate {
	tu.mutation.ClearVersions()
	return tu
}

// RemoveVersionIDs removes the "versions" edge to TemplateVersion entities by IDs.
func (tu *TemplateUpdate) RemoveVersionIDs(ids ...object.ID) *TemplateUpdate {
	tu.mutation.RemoveVersionIDs(ids...)
	return tu
}

// RemoveVersions removes "versions" edges to TemplateVersion entities.
func (tu *TemplateUpdate) RemoveVersions(t ...*TemplateVersion) *TemplateUpdate {
	ids := make([]object.ID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.RemoveVersionIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TemplateUpdate) Save(ctx context.Context) (int, error) {
	if err := tu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TemplateUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TemplateUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TemplateUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tu *TemplateUpdate) defaults() error {
	if _, ok := tu.mutation.UpdateTime(); !ok {
		if template.UpdateDefaultUpdateTime == nil {
			return fmt.Errorf("model: uninitialized template.UpdateDefaultUpdateTime (forgotten import model/runtime?)")
		}
		v := template.UpdateDefaultUpdateTime()
		tu.mutation.SetUpdateTime(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (tu *TemplateUpdate) check() error {
	if v, ok := tu.mutation.Source(); ok {
		if err := template.SourceValidator(v); err != nil {
			return &ValidationError{Name: "source", err: fmt.Errorf(`model: validator failed for field "Template.source": %w`, err)}
		}
	}
	return nil
}

// Set is different from other Set* methods,
// it sets the value by judging the definition of each field within the entire object.
//
// For default fields, Set calls if the value is not zero.
//
// For no default but required fields, Set calls directly.
//
// For no default but optional fields, Set calls if the value is not zero,
// or clears if the value is zero.
//
// For example:
//
//	## Without Default
//
//	### Required
//
//	db.SetX(obj.X)
//
//	### Optional or Default
//
//	if _is_zero_value_(obj.X) {
//	   db.SetX(obj.X)
//	} else {
//	   db.ClearX()
//	}
//
//	## With Default
//
//	if _is_zero_value_(obj.X) {
//	   db.SetX(obj.X)
//	}
func (tu *TemplateUpdate) Set(obj *Template) *TemplateUpdate {
	// Without Default.
	if obj.Description != "" {
		tu.SetDescription(obj.Description)
	} else {
		tu.ClearDescription()
	}
	if !reflect.ValueOf(obj.Labels).IsZero() {
		tu.SetLabels(obj.Labels)
	} else {
		tu.ClearLabels()
	}
	if !reflect.ValueOf(obj.Status).IsZero() {
		tu.SetStatus(obj.Status)
	}
	if obj.Icon != "" {
		tu.SetIcon(obj.Icon)
	}
	tu.SetSource(obj.Source)

	// With Default.
	if obj.UpdateTime != nil {
		tu.SetUpdateTime(*obj.UpdateTime)
	}

	// Record the given object.
	tu.object = obj

	return tu
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (tu *TemplateUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *TemplateUpdate {
	tu.modifiers = append(tu.modifiers, modifiers...)
	return tu
}

func (tu *TemplateUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := tu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(template.Table, template.Columns, sqlgraph.NewFieldSpec(template.FieldID, field.TypeString))
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.Description(); ok {
		_spec.SetField(template.FieldDescription, field.TypeString, value)
	}
	if tu.mutation.DescriptionCleared() {
		_spec.ClearField(template.FieldDescription, field.TypeString)
	}
	if value, ok := tu.mutation.Labels(); ok {
		_spec.SetField(template.FieldLabels, field.TypeJSON, value)
	}
	if tu.mutation.LabelsCleared() {
		_spec.ClearField(template.FieldLabels, field.TypeJSON)
	}
	if value, ok := tu.mutation.UpdateTime(); ok {
		_spec.SetField(template.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := tu.mutation.Status(); ok {
		_spec.SetField(template.FieldStatus, field.TypeJSON, value)
	}
	if tu.mutation.StatusCleared() {
		_spec.ClearField(template.FieldStatus, field.TypeJSON)
	}
	if value, ok := tu.mutation.Icon(); ok {
		_spec.SetField(template.FieldIcon, field.TypeString, value)
	}
	if tu.mutation.IconCleared() {
		_spec.ClearField(template.FieldIcon, field.TypeString)
	}
	if value, ok := tu.mutation.Source(); ok {
		_spec.SetField(template.FieldSource, field.TypeString, value)
	}
	if tu.mutation.VersionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   template.VersionsTable,
			Columns: []string{template.VersionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(templateversion.FieldID, field.TypeString),
			},
		}
		edge.Schema = tu.schemaConfig.TemplateVersion
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedVersionsIDs(); len(nodes) > 0 && !tu.mutation.VersionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   template.VersionsTable,
			Columns: []string{template.VersionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(templateversion.FieldID, field.TypeString),
			},
		}
		edge.Schema = tu.schemaConfig.TemplateVersion
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.VersionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   template.VersionsTable,
			Columns: []string{template.VersionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(templateversion.FieldID, field.TypeString),
			},
		}
		edge.Schema = tu.schemaConfig.TemplateVersion
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = tu.schemaConfig.Template
	ctx = internal.NewSchemaConfigContext(ctx, tu.schemaConfig)
	_spec.AddModifiers(tu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{template.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// TemplateUpdateOne is the builder for updating a single Template entity.
type TemplateUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *TemplateMutation
	modifiers []func(*sql.UpdateBuilder)
	object    *Template
}

// SetDescription sets the "description" field.
func (tuo *TemplateUpdateOne) SetDescription(s string) *TemplateUpdateOne {
	tuo.mutation.SetDescription(s)
	return tuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (tuo *TemplateUpdateOne) SetNillableDescription(s *string) *TemplateUpdateOne {
	if s != nil {
		tuo.SetDescription(*s)
	}
	return tuo
}

// ClearDescription clears the value of the "description" field.
func (tuo *TemplateUpdateOne) ClearDescription() *TemplateUpdateOne {
	tuo.mutation.ClearDescription()
	return tuo
}

// SetLabels sets the "labels" field.
func (tuo *TemplateUpdateOne) SetLabels(m map[string]string) *TemplateUpdateOne {
	tuo.mutation.SetLabels(m)
	return tuo
}

// ClearLabels clears the value of the "labels" field.
func (tuo *TemplateUpdateOne) ClearLabels() *TemplateUpdateOne {
	tuo.mutation.ClearLabels()
	return tuo
}

// SetUpdateTime sets the "update_time" field.
func (tuo *TemplateUpdateOne) SetUpdateTime(t time.Time) *TemplateUpdateOne {
	tuo.mutation.SetUpdateTime(t)
	return tuo
}

// SetStatus sets the "status" field.
func (tuo *TemplateUpdateOne) SetStatus(s status.Status) *TemplateUpdateOne {
	tuo.mutation.SetStatus(s)
	return tuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (tuo *TemplateUpdateOne) SetNillableStatus(s *status.Status) *TemplateUpdateOne {
	if s != nil {
		tuo.SetStatus(*s)
	}
	return tuo
}

// ClearStatus clears the value of the "status" field.
func (tuo *TemplateUpdateOne) ClearStatus() *TemplateUpdateOne {
	tuo.mutation.ClearStatus()
	return tuo
}

// SetIcon sets the "icon" field.
func (tuo *TemplateUpdateOne) SetIcon(s string) *TemplateUpdateOne {
	tuo.mutation.SetIcon(s)
	return tuo
}

// SetNillableIcon sets the "icon" field if the given value is not nil.
func (tuo *TemplateUpdateOne) SetNillableIcon(s *string) *TemplateUpdateOne {
	if s != nil {
		tuo.SetIcon(*s)
	}
	return tuo
}

// ClearIcon clears the value of the "icon" field.
func (tuo *TemplateUpdateOne) ClearIcon() *TemplateUpdateOne {
	tuo.mutation.ClearIcon()
	return tuo
}

// SetSource sets the "source" field.
func (tuo *TemplateUpdateOne) SetSource(s string) *TemplateUpdateOne {
	tuo.mutation.SetSource(s)
	return tuo
}

// AddVersionIDs adds the "versions" edge to the TemplateVersion entity by IDs.
func (tuo *TemplateUpdateOne) AddVersionIDs(ids ...object.ID) *TemplateUpdateOne {
	tuo.mutation.AddVersionIDs(ids...)
	return tuo
}

// AddVersions adds the "versions" edges to the TemplateVersion entity.
func (tuo *TemplateUpdateOne) AddVersions(t ...*TemplateVersion) *TemplateUpdateOne {
	ids := make([]object.ID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.AddVersionIDs(ids...)
}

// Mutation returns the TemplateMutation object of the builder.
func (tuo *TemplateUpdateOne) Mutation() *TemplateMutation {
	return tuo.mutation
}

// ClearVersions clears all "versions" edges to the TemplateVersion entity.
func (tuo *TemplateUpdateOne) ClearVersions() *TemplateUpdateOne {
	tuo.mutation.ClearVersions()
	return tuo
}

// RemoveVersionIDs removes the "versions" edge to TemplateVersion entities by IDs.
func (tuo *TemplateUpdateOne) RemoveVersionIDs(ids ...object.ID) *TemplateUpdateOne {
	tuo.mutation.RemoveVersionIDs(ids...)
	return tuo
}

// RemoveVersions removes "versions" edges to TemplateVersion entities.
func (tuo *TemplateUpdateOne) RemoveVersions(t ...*TemplateVersion) *TemplateUpdateOne {
	ids := make([]object.ID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.RemoveVersionIDs(ids...)
}

// Where appends a list predicates to the TemplateUpdate builder.
func (tuo *TemplateUpdateOne) Where(ps ...predicate.Template) *TemplateUpdateOne {
	tuo.mutation.Where(ps...)
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TemplateUpdateOne) Select(field string, fields ...string) *TemplateUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Template entity.
func (tuo *TemplateUpdateOne) Save(ctx context.Context) (*Template, error) {
	if err := tuo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TemplateUpdateOne) SaveX(ctx context.Context) *Template {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TemplateUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TemplateUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tuo *TemplateUpdateOne) defaults() error {
	if _, ok := tuo.mutation.UpdateTime(); !ok {
		if template.UpdateDefaultUpdateTime == nil {
			return fmt.Errorf("model: uninitialized template.UpdateDefaultUpdateTime (forgotten import model/runtime?)")
		}
		v := template.UpdateDefaultUpdateTime()
		tuo.mutation.SetUpdateTime(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TemplateUpdateOne) check() error {
	if v, ok := tuo.mutation.Source(); ok {
		if err := template.SourceValidator(v); err != nil {
			return &ValidationError{Name: "source", err: fmt.Errorf(`model: validator failed for field "Template.source": %w`, err)}
		}
	}
	return nil
}

// Set is different from other Set* methods,
// it sets the value by judging the definition of each field within the entire object.
//
// For default fields, Set calls if the value changes from the original.
//
// For no default but required fields, Set calls if the value changes from the original.
//
// For no default but optional fields, Set calls if the value changes from the original,
// or clears if changes to zero.
//
// For example:
//
//	## Without Default
//
//	### Required
//
//	db.SetX(obj.X)
//
//	### Optional or Default
//
//	if _is_zero_value_(obj.X) {
//	   if _is_not_equal_(db.X, obj.X) {
//	      db.SetX(obj.X)
//	   }
//	} else {
//	   db.ClearX()
//	}
//
//	## With Default
//
//	if _is_zero_value_(obj.X) && _is_not_equal_(db.X, obj.X) {
//	   db.SetX(obj.X)
//	}
func (tuo *TemplateUpdateOne) Set(obj *Template) *TemplateUpdateOne {
	h := func(n ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			mt := m.(*TemplateMutation)
			db, err := mt.Client().Template.Get(ctx, *mt.id)
			if err != nil {
				return nil, fmt.Errorf("failed getting Template with id: %v", *mt.id)
			}

			// Without Default.
			if obj.Description != "" {
				if db.Description != obj.Description {
					tuo.SetDescription(obj.Description)
				}
			} else {
				tuo.ClearDescription()
			}
			if !reflect.ValueOf(obj.Labels).IsZero() {
				if !reflect.DeepEqual(db.Labels, obj.Labels) {
					tuo.SetLabels(obj.Labels)
				}
			} else {
				tuo.ClearLabels()
			}
			if !reflect.ValueOf(obj.Status).IsZero() {
				if !db.Status.Equal(obj.Status) {
					tuo.SetStatus(obj.Status)
				}
			}
			if obj.Icon != "" {
				if db.Icon != obj.Icon {
					tuo.SetIcon(obj.Icon)
				}
			}
			if db.Source != obj.Source {
				tuo.SetSource(obj.Source)
			}

			// With Default.
			if (obj.UpdateTime != nil) && (!reflect.DeepEqual(db.UpdateTime, obj.UpdateTime)) {
				tuo.SetUpdateTime(*obj.UpdateTime)
			}

			// Record the given object.
			tuo.object = obj

			return n.Mutate(ctx, m)
		})
	}

	tuo.hooks = append(tuo.hooks, h)

	return tuo
}

// getClientSet returns the ClientSet for the given builder.
func (tuo *TemplateUpdateOne) getClientSet() (mc ClientSet) {
	if _, ok := tuo.config.driver.(*txDriver); ok {
		tx := &Tx{config: tuo.config}
		tx.init()
		mc = tx
	} else {
		cli := &Client{config: tuo.config}
		cli.init()
		mc = cli
	}
	return mc
}

// SaveE calls the given function after updated the Template entity,
// which is always good for cascading update operations.
func (tuo *TemplateUpdateOne) SaveE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *Template) error) (*Template, error) {
	obj, err := tuo.Save(ctx)
	if err != nil &&
		(tuo.object == nil || !errors.Is(err, stdsql.ErrNoRows)) {
		return nil, err
	}

	if len(cbs) == 0 {
		return obj, err
	}

	mc := tuo.getClientSet()

	if obj == nil {
		obj = tuo.object
	} else if x := tuo.object; x != nil {
		if _, set := tuo.mutation.Field(template.FieldDescription); set {
			obj.Description = x.Description
		}
		if _, set := tuo.mutation.Field(template.FieldLabels); set {
			obj.Labels = x.Labels
		}
		if _, set := tuo.mutation.Field(template.FieldStatus); set {
			obj.Status = x.Status
		}
		if _, set := tuo.mutation.Field(template.FieldIcon); set {
			obj.Icon = x.Icon
		}
		if _, set := tuo.mutation.Field(template.FieldSource); set {
			obj.Source = x.Source
		}
		obj.Edges = x.Edges
	}

	for i := range cbs {
		if err = cbs[i](ctx, mc, obj); err != nil {
			return nil, err
		}
	}

	return obj, nil
}

// SaveEX is like SaveE, but panics if an error occurs.
func (tuo *TemplateUpdateOne) SaveEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *Template) error) *Template {
	obj, err := tuo.SaveE(ctx, cbs...)
	if err != nil {
		panic(err)
	}
	return obj
}

// ExecE calls the given function after executed the query,
// which is always good for cascading update operations.
func (tuo *TemplateUpdateOne) ExecE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *Template) error) error {
	_, err := tuo.SaveE(ctx, cbs...)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TemplateUpdateOne) ExecEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *Template) error) {
	if err := tuo.ExecE(ctx, cbs...); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (tuo *TemplateUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *TemplateUpdateOne {
	tuo.modifiers = append(tuo.modifiers, modifiers...)
	return tuo
}

func (tuo *TemplateUpdateOne) sqlSave(ctx context.Context) (_node *Template, err error) {
	if err := tuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(template.Table, template.Columns, sqlgraph.NewFieldSpec(template.FieldID, field.TypeString))
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`model: missing "Template.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, template.FieldID)
		for _, f := range fields {
			if !template.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("model: invalid field %q for query", f)}
			}
			if f != template.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.Description(); ok {
		_spec.SetField(template.FieldDescription, field.TypeString, value)
	}
	if tuo.mutation.DescriptionCleared() {
		_spec.ClearField(template.FieldDescription, field.TypeString)
	}
	if value, ok := tuo.mutation.Labels(); ok {
		_spec.SetField(template.FieldLabels, field.TypeJSON, value)
	}
	if tuo.mutation.LabelsCleared() {
		_spec.ClearField(template.FieldLabels, field.TypeJSON)
	}
	if value, ok := tuo.mutation.UpdateTime(); ok {
		_spec.SetField(template.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := tuo.mutation.Status(); ok {
		_spec.SetField(template.FieldStatus, field.TypeJSON, value)
	}
	if tuo.mutation.StatusCleared() {
		_spec.ClearField(template.FieldStatus, field.TypeJSON)
	}
	if value, ok := tuo.mutation.Icon(); ok {
		_spec.SetField(template.FieldIcon, field.TypeString, value)
	}
	if tuo.mutation.IconCleared() {
		_spec.ClearField(template.FieldIcon, field.TypeString)
	}
	if value, ok := tuo.mutation.Source(); ok {
		_spec.SetField(template.FieldSource, field.TypeString, value)
	}
	if tuo.mutation.VersionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   template.VersionsTable,
			Columns: []string{template.VersionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(templateversion.FieldID, field.TypeString),
			},
		}
		edge.Schema = tuo.schemaConfig.TemplateVersion
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedVersionsIDs(); len(nodes) > 0 && !tuo.mutation.VersionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   template.VersionsTable,
			Columns: []string{template.VersionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(templateversion.FieldID, field.TypeString),
			},
		}
		edge.Schema = tuo.schemaConfig.TemplateVersion
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.VersionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   template.VersionsTable,
			Columns: []string{template.VersionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(templateversion.FieldID, field.TypeString),
			},
		}
		edge.Schema = tuo.schemaConfig.TemplateVersion
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = tuo.schemaConfig.Template
	ctx = internal.NewSchemaConfigContext(ctx, tuo.schemaConfig)
	_spec.AddModifiers(tuo.modifiers...)
	_node = &Template{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{template.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}

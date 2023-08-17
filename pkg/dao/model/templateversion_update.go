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
	"github.com/seal-io/walrus/pkg/dao/model/service"
	"github.com/seal-io/walrus/pkg/dao/model/templateversion"
	"github.com/seal-io/walrus/pkg/dao/types"
	"github.com/seal-io/walrus/pkg/dao/types/object"
)

// TemplateVersionUpdate is the builder for updating TemplateVersion entities.
type TemplateVersionUpdate struct {
	config
	hooks     []Hook
	mutation  *TemplateVersionMutation
	modifiers []func(*sql.UpdateBuilder)
	object    *TemplateVersion
}

// Where appends a list predicates to the TemplateVersionUpdate builder.
func (tvu *TemplateVersionUpdate) Where(ps ...predicate.TemplateVersion) *TemplateVersionUpdate {
	tvu.mutation.Where(ps...)
	return tvu
}

// SetUpdateTime sets the "update_time" field.
func (tvu *TemplateVersionUpdate) SetUpdateTime(t time.Time) *TemplateVersionUpdate {
	tvu.mutation.SetUpdateTime(t)
	return tvu
}

// SetSchema sets the "schema" field.
func (tvu *TemplateVersionUpdate) SetSchema(ts *types.TemplateSchema) *TemplateVersionUpdate {
	tvu.mutation.SetSchema(ts)
	return tvu
}

// AddServiceIDs adds the "services" edge to the Service entity by IDs.
func (tvu *TemplateVersionUpdate) AddServiceIDs(ids ...object.ID) *TemplateVersionUpdate {
	tvu.mutation.AddServiceIDs(ids...)
	return tvu
}

// AddServices adds the "services" edges to the Service entity.
func (tvu *TemplateVersionUpdate) AddServices(s ...*Service) *TemplateVersionUpdate {
	ids := make([]object.ID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return tvu.AddServiceIDs(ids...)
}

// Mutation returns the TemplateVersionMutation object of the builder.
func (tvu *TemplateVersionUpdate) Mutation() *TemplateVersionMutation {
	return tvu.mutation
}

// ClearServices clears all "services" edges to the Service entity.
func (tvu *TemplateVersionUpdate) ClearServices() *TemplateVersionUpdate {
	tvu.mutation.ClearServices()
	return tvu
}

// RemoveServiceIDs removes the "services" edge to Service entities by IDs.
func (tvu *TemplateVersionUpdate) RemoveServiceIDs(ids ...object.ID) *TemplateVersionUpdate {
	tvu.mutation.RemoveServiceIDs(ids...)
	return tvu
}

// RemoveServices removes "services" edges to Service entities.
func (tvu *TemplateVersionUpdate) RemoveServices(s ...*Service) *TemplateVersionUpdate {
	ids := make([]object.ID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return tvu.RemoveServiceIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tvu *TemplateVersionUpdate) Save(ctx context.Context) (int, error) {
	if err := tvu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, tvu.sqlSave, tvu.mutation, tvu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tvu *TemplateVersionUpdate) SaveX(ctx context.Context) int {
	affected, err := tvu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tvu *TemplateVersionUpdate) Exec(ctx context.Context) error {
	_, err := tvu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tvu *TemplateVersionUpdate) ExecX(ctx context.Context) {
	if err := tvu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tvu *TemplateVersionUpdate) defaults() error {
	if _, ok := tvu.mutation.UpdateTime(); !ok {
		if templateversion.UpdateDefaultUpdateTime == nil {
			return fmt.Errorf("model: uninitialized templateversion.UpdateDefaultUpdateTime (forgotten import model/runtime?)")
		}
		v := templateversion.UpdateDefaultUpdateTime()
		tvu.mutation.SetUpdateTime(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (tvu *TemplateVersionUpdate) check() error {
	if _, ok := tvu.mutation.TemplateID(); tvu.mutation.TemplateCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "TemplateVersion.template"`)
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
func (tvu *TemplateVersionUpdate) Set(obj *TemplateVersion) *TemplateVersionUpdate {
	// Without Default.
	tvu.SetSchema(obj.Schema)

	// With Default.
	if obj.UpdateTime != nil {
		tvu.SetUpdateTime(*obj.UpdateTime)
	}

	// Record the given object.
	tvu.object = obj

	return tvu
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (tvu *TemplateVersionUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *TemplateVersionUpdate {
	tvu.modifiers = append(tvu.modifiers, modifiers...)
	return tvu
}

func (tvu *TemplateVersionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := tvu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(templateversion.Table, templateversion.Columns, sqlgraph.NewFieldSpec(templateversion.FieldID, field.TypeString))
	if ps := tvu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tvu.mutation.UpdateTime(); ok {
		_spec.SetField(templateversion.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := tvu.mutation.Schema(); ok {
		_spec.SetField(templateversion.FieldSchema, field.TypeJSON, value)
	}
	if tvu.mutation.ServicesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   templateversion.ServicesTable,
			Columns: []string{templateversion.ServicesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(service.FieldID, field.TypeString),
			},
		}
		edge.Schema = tvu.schemaConfig.Service
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tvu.mutation.RemovedServicesIDs(); len(nodes) > 0 && !tvu.mutation.ServicesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   templateversion.ServicesTable,
			Columns: []string{templateversion.ServicesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(service.FieldID, field.TypeString),
			},
		}
		edge.Schema = tvu.schemaConfig.Service
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tvu.mutation.ServicesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   templateversion.ServicesTable,
			Columns: []string{templateversion.ServicesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(service.FieldID, field.TypeString),
			},
		}
		edge.Schema = tvu.schemaConfig.Service
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = tvu.schemaConfig.TemplateVersion
	ctx = internal.NewSchemaConfigContext(ctx, tvu.schemaConfig)
	_spec.AddModifiers(tvu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, tvu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{templateversion.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tvu.mutation.done = true
	return n, nil
}

// TemplateVersionUpdateOne is the builder for updating a single TemplateVersion entity.
type TemplateVersionUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *TemplateVersionMutation
	modifiers []func(*sql.UpdateBuilder)
	object    *TemplateVersion
}

// SetUpdateTime sets the "update_time" field.
func (tvuo *TemplateVersionUpdateOne) SetUpdateTime(t time.Time) *TemplateVersionUpdateOne {
	tvuo.mutation.SetUpdateTime(t)
	return tvuo
}

// SetSchema sets the "schema" field.
func (tvuo *TemplateVersionUpdateOne) SetSchema(ts *types.TemplateSchema) *TemplateVersionUpdateOne {
	tvuo.mutation.SetSchema(ts)
	return tvuo
}

// AddServiceIDs adds the "services" edge to the Service entity by IDs.
func (tvuo *TemplateVersionUpdateOne) AddServiceIDs(ids ...object.ID) *TemplateVersionUpdateOne {
	tvuo.mutation.AddServiceIDs(ids...)
	return tvuo
}

// AddServices adds the "services" edges to the Service entity.
func (tvuo *TemplateVersionUpdateOne) AddServices(s ...*Service) *TemplateVersionUpdateOne {
	ids := make([]object.ID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return tvuo.AddServiceIDs(ids...)
}

// Mutation returns the TemplateVersionMutation object of the builder.
func (tvuo *TemplateVersionUpdateOne) Mutation() *TemplateVersionMutation {
	return tvuo.mutation
}

// ClearServices clears all "services" edges to the Service entity.
func (tvuo *TemplateVersionUpdateOne) ClearServices() *TemplateVersionUpdateOne {
	tvuo.mutation.ClearServices()
	return tvuo
}

// RemoveServiceIDs removes the "services" edge to Service entities by IDs.
func (tvuo *TemplateVersionUpdateOne) RemoveServiceIDs(ids ...object.ID) *TemplateVersionUpdateOne {
	tvuo.mutation.RemoveServiceIDs(ids...)
	return tvuo
}

// RemoveServices removes "services" edges to Service entities.
func (tvuo *TemplateVersionUpdateOne) RemoveServices(s ...*Service) *TemplateVersionUpdateOne {
	ids := make([]object.ID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return tvuo.RemoveServiceIDs(ids...)
}

// Where appends a list predicates to the TemplateVersionUpdate builder.
func (tvuo *TemplateVersionUpdateOne) Where(ps ...predicate.TemplateVersion) *TemplateVersionUpdateOne {
	tvuo.mutation.Where(ps...)
	return tvuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tvuo *TemplateVersionUpdateOne) Select(field string, fields ...string) *TemplateVersionUpdateOne {
	tvuo.fields = append([]string{field}, fields...)
	return tvuo
}

// Save executes the query and returns the updated TemplateVersion entity.
func (tvuo *TemplateVersionUpdateOne) Save(ctx context.Context) (*TemplateVersion, error) {
	if err := tvuo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, tvuo.sqlSave, tvuo.mutation, tvuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tvuo *TemplateVersionUpdateOne) SaveX(ctx context.Context) *TemplateVersion {
	node, err := tvuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tvuo *TemplateVersionUpdateOne) Exec(ctx context.Context) error {
	_, err := tvuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tvuo *TemplateVersionUpdateOne) ExecX(ctx context.Context) {
	if err := tvuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tvuo *TemplateVersionUpdateOne) defaults() error {
	if _, ok := tvuo.mutation.UpdateTime(); !ok {
		if templateversion.UpdateDefaultUpdateTime == nil {
			return fmt.Errorf("model: uninitialized templateversion.UpdateDefaultUpdateTime (forgotten import model/runtime?)")
		}
		v := templateversion.UpdateDefaultUpdateTime()
		tvuo.mutation.SetUpdateTime(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (tvuo *TemplateVersionUpdateOne) check() error {
	if _, ok := tvuo.mutation.TemplateID(); tvuo.mutation.TemplateCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "TemplateVersion.template"`)
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
func (tvuo *TemplateVersionUpdateOne) Set(obj *TemplateVersion) *TemplateVersionUpdateOne {
	h := func(n ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			mt := m.(*TemplateVersionMutation)
			db, err := mt.Client().TemplateVersion.Get(ctx, *mt.id)
			if err != nil {
				return nil, fmt.Errorf("failed getting TemplateVersion with id: %v", *mt.id)
			}

			// Without Default.
			if !reflect.DeepEqual(db.Schema, obj.Schema) {
				tvuo.SetSchema(obj.Schema)
			}

			// With Default.
			if (obj.UpdateTime != nil) && (!reflect.DeepEqual(db.UpdateTime, obj.UpdateTime)) {
				tvuo.SetUpdateTime(*obj.UpdateTime)
			}

			// Record the given object.
			tvuo.object = obj

			return n.Mutate(ctx, m)
		})
	}

	tvuo.hooks = append(tvuo.hooks, h)

	return tvuo
}

// getClientSet returns the ClientSet for the given builder.
func (tvuo *TemplateVersionUpdateOne) getClientSet() (mc ClientSet) {
	if _, ok := tvuo.config.driver.(*txDriver); ok {
		tx := &Tx{config: tvuo.config}
		tx.init()
		mc = tx
	} else {
		cli := &Client{config: tvuo.config}
		cli.init()
		mc = cli
	}
	return mc
}

// SaveE calls the given function after updated the TemplateVersion entity,
// which is always good for cascading update operations.
func (tvuo *TemplateVersionUpdateOne) SaveE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *TemplateVersion) error) (*TemplateVersion, error) {
	obj, err := tvuo.Save(ctx)
	if err != nil &&
		(tvuo.object == nil || !errors.Is(err, stdsql.ErrNoRows)) {
		return nil, err
	}

	if len(cbs) == 0 {
		return obj, err
	}

	mc := tvuo.getClientSet()

	if obj == nil {
		obj = tvuo.object
	} else if x := tvuo.object; x != nil {
		if _, set := tvuo.mutation.Field(templateversion.FieldSchema); set {
			obj.Schema = x.Schema
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
func (tvuo *TemplateVersionUpdateOne) SaveEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *TemplateVersion) error) *TemplateVersion {
	obj, err := tvuo.SaveE(ctx, cbs...)
	if err != nil {
		panic(err)
	}
	return obj
}

// ExecE calls the given function after executed the query,
// which is always good for cascading update operations.
func (tvuo *TemplateVersionUpdateOne) ExecE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *TemplateVersion) error) error {
	_, err := tvuo.SaveE(ctx, cbs...)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tvuo *TemplateVersionUpdateOne) ExecEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *TemplateVersion) error) {
	if err := tvuo.ExecE(ctx, cbs...); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (tvuo *TemplateVersionUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *TemplateVersionUpdateOne {
	tvuo.modifiers = append(tvuo.modifiers, modifiers...)
	return tvuo
}

func (tvuo *TemplateVersionUpdateOne) sqlSave(ctx context.Context) (_node *TemplateVersion, err error) {
	if err := tvuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(templateversion.Table, templateversion.Columns, sqlgraph.NewFieldSpec(templateversion.FieldID, field.TypeString))
	id, ok := tvuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`model: missing "TemplateVersion.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tvuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, templateversion.FieldID)
		for _, f := range fields {
			if !templateversion.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("model: invalid field %q for query", f)}
			}
			if f != templateversion.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tvuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tvuo.mutation.UpdateTime(); ok {
		_spec.SetField(templateversion.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := tvuo.mutation.Schema(); ok {
		_spec.SetField(templateversion.FieldSchema, field.TypeJSON, value)
	}
	if tvuo.mutation.ServicesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   templateversion.ServicesTable,
			Columns: []string{templateversion.ServicesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(service.FieldID, field.TypeString),
			},
		}
		edge.Schema = tvuo.schemaConfig.Service
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tvuo.mutation.RemovedServicesIDs(); len(nodes) > 0 && !tvuo.mutation.ServicesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   templateversion.ServicesTable,
			Columns: []string{templateversion.ServicesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(service.FieldID, field.TypeString),
			},
		}
		edge.Schema = tvuo.schemaConfig.Service
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tvuo.mutation.ServicesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   templateversion.ServicesTable,
			Columns: []string{templateversion.ServicesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(service.FieldID, field.TypeString),
			},
		}
		edge.Schema = tvuo.schemaConfig.Service
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = tvuo.schemaConfig.TemplateVersion
	ctx = internal.NewSchemaConfigContext(ctx, tvuo.schemaConfig)
	_spec.AddModifiers(tvuo.modifiers...)
	_node = &TemplateVersion{config: tvuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tvuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{templateversion.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tvuo.mutation.done = true
	return _node, nil
}

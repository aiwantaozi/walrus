// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "ent". DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"

	"github.com/seal-io/seal/pkg/dao/model/template"
	"github.com/seal-io/seal/pkg/dao/model/templateversion"
	"github.com/seal-io/seal/pkg/dao/types/oid"
)

// TemplateCreate is the builder for creating a Template entity.
type TemplateCreate struct {
	config
	mutation *TemplateMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetStatus sets the "status" field.
func (tc *TemplateCreate) SetStatus(s string) *TemplateCreate {
	tc.mutation.SetStatus(s)
	return tc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (tc *TemplateCreate) SetNillableStatus(s *string) *TemplateCreate {
	if s != nil {
		tc.SetStatus(*s)
	}
	return tc
}

// SetStatusMessage sets the "statusMessage" field.
func (tc *TemplateCreate) SetStatusMessage(s string) *TemplateCreate {
	tc.mutation.SetStatusMessage(s)
	return tc
}

// SetNillableStatusMessage sets the "statusMessage" field if the given value is not nil.
func (tc *TemplateCreate) SetNillableStatusMessage(s *string) *TemplateCreate {
	if s != nil {
		tc.SetStatusMessage(*s)
	}
	return tc
}

// SetCreateTime sets the "createTime" field.
func (tc *TemplateCreate) SetCreateTime(t time.Time) *TemplateCreate {
	tc.mutation.SetCreateTime(t)
	return tc
}

// SetNillableCreateTime sets the "createTime" field if the given value is not nil.
func (tc *TemplateCreate) SetNillableCreateTime(t *time.Time) *TemplateCreate {
	if t != nil {
		tc.SetCreateTime(*t)
	}
	return tc
}

// SetUpdateTime sets the "updateTime" field.
func (tc *TemplateCreate) SetUpdateTime(t time.Time) *TemplateCreate {
	tc.mutation.SetUpdateTime(t)
	return tc
}

// SetNillableUpdateTime sets the "updateTime" field if the given value is not nil.
func (tc *TemplateCreate) SetNillableUpdateTime(t *time.Time) *TemplateCreate {
	if t != nil {
		tc.SetUpdateTime(*t)
	}
	return tc
}

// SetDescription sets the "description" field.
func (tc *TemplateCreate) SetDescription(s string) *TemplateCreate {
	tc.mutation.SetDescription(s)
	return tc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (tc *TemplateCreate) SetNillableDescription(s *string) *TemplateCreate {
	if s != nil {
		tc.SetDescription(*s)
	}
	return tc
}

// SetIcon sets the "icon" field.
func (tc *TemplateCreate) SetIcon(s string) *TemplateCreate {
	tc.mutation.SetIcon(s)
	return tc
}

// SetNillableIcon sets the "icon" field if the given value is not nil.
func (tc *TemplateCreate) SetNillableIcon(s *string) *TemplateCreate {
	if s != nil {
		tc.SetIcon(*s)
	}
	return tc
}

// SetLabels sets the "labels" field.
func (tc *TemplateCreate) SetLabels(m map[string]string) *TemplateCreate {
	tc.mutation.SetLabels(m)
	return tc
}

// SetSource sets the "source" field.
func (tc *TemplateCreate) SetSource(s string) *TemplateCreate {
	tc.mutation.SetSource(s)
	return tc
}

// SetID sets the "id" field.
func (tc *TemplateCreate) SetID(s string) *TemplateCreate {
	tc.mutation.SetID(s)
	return tc
}

// AddVersionIDs adds the "versions" edge to the TemplateVersion entity by IDs.
func (tc *TemplateCreate) AddVersionIDs(ids ...oid.ID) *TemplateCreate {
	tc.mutation.AddVersionIDs(ids...)
	return tc
}

// AddVersions adds the "versions" edges to the TemplateVersion entity.
func (tc *TemplateCreate) AddVersions(t ...*TemplateVersion) *TemplateCreate {
	ids := make([]oid.ID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tc.AddVersionIDs(ids...)
}

// Mutation returns the TemplateMutation object of the builder.
func (tc *TemplateCreate) Mutation() *TemplateMutation {
	return tc.mutation
}

// Save creates the Template in the database.
func (tc *TemplateCreate) Save(ctx context.Context) (*Template, error) {
	tc.defaults()
	return withHooks[*Template, TemplateMutation](ctx, tc.sqlSave, tc.mutation, tc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TemplateCreate) SaveX(ctx context.Context) *Template {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TemplateCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TemplateCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *TemplateCreate) defaults() {
	if _, ok := tc.mutation.CreateTime(); !ok {
		v := template.DefaultCreateTime()
		tc.mutation.SetCreateTime(v)
	}
	if _, ok := tc.mutation.UpdateTime(); !ok {
		v := template.DefaultUpdateTime()
		tc.mutation.SetUpdateTime(v)
	}
	if _, ok := tc.mutation.Labels(); !ok {
		v := template.DefaultLabels
		tc.mutation.SetLabels(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TemplateCreate) check() error {
	if _, ok := tc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "createTime", err: errors.New(`model: missing required field "Template.createTime"`)}
	}
	if _, ok := tc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "updateTime", err: errors.New(`model: missing required field "Template.updateTime"`)}
	}
	if _, ok := tc.mutation.Labels(); !ok {
		return &ValidationError{Name: "labels", err: errors.New(`model: missing required field "Template.labels"`)}
	}
	if _, ok := tc.mutation.Source(); !ok {
		return &ValidationError{Name: "source", err: errors.New(`model: missing required field "Template.source"`)}
	}
	if v, ok := tc.mutation.Source(); ok {
		if err := template.SourceValidator(v); err != nil {
			return &ValidationError{Name: "source", err: fmt.Errorf(`model: validator failed for field "Template.source": %w`, err)}
		}
	}
	if v, ok := tc.mutation.ID(); ok {
		if err := template.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`model: validator failed for field "Template.id": %w`, err)}
		}
	}
	return nil
}

func (tc *TemplateCreate) sqlSave(ctx context.Context) (*Template, error) {
	if err := tc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Template.ID type: %T", _spec.ID.Value)
		}
	}
	tc.mutation.id = &_node.ID
	tc.mutation.done = true
	return _node, nil
}

func (tc *TemplateCreate) createSpec() (*Template, *sqlgraph.CreateSpec) {
	var (
		_node = &Template{config: tc.config}
		_spec = sqlgraph.NewCreateSpec(template.Table, sqlgraph.NewFieldSpec(template.FieldID, field.TypeString))
	)
	_spec.Schema = tc.schemaConfig.Template
	_spec.OnConflict = tc.conflict
	if id, ok := tc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := tc.mutation.Status(); ok {
		_spec.SetField(template.FieldStatus, field.TypeString, value)
		_node.Status = value
	}
	if value, ok := tc.mutation.StatusMessage(); ok {
		_spec.SetField(template.FieldStatusMessage, field.TypeString, value)
		_node.StatusMessage = value
	}
	if value, ok := tc.mutation.CreateTime(); ok {
		_spec.SetField(template.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = &value
	}
	if value, ok := tc.mutation.UpdateTime(); ok {
		_spec.SetField(template.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = &value
	}
	if value, ok := tc.mutation.Description(); ok {
		_spec.SetField(template.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := tc.mutation.Icon(); ok {
		_spec.SetField(template.FieldIcon, field.TypeString, value)
		_node.Icon = value
	}
	if value, ok := tc.mutation.Labels(); ok {
		_spec.SetField(template.FieldLabels, field.TypeJSON, value)
		_node.Labels = value
	}
	if value, ok := tc.mutation.Source(); ok {
		_spec.SetField(template.FieldSource, field.TypeString, value)
		_node.Source = value
	}
	if nodes := tc.mutation.VersionsIDs(); len(nodes) > 0 {
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
		edge.Schema = tc.schemaConfig.TemplateVersion
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Template.Create().
//		SetStatus(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TemplateUpsert) {
//			SetStatus(v+v).
//		}).
//		Exec(ctx)
func (tc *TemplateCreate) OnConflict(opts ...sql.ConflictOption) *TemplateUpsertOne {
	tc.conflict = opts
	return &TemplateUpsertOne{
		create: tc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Template.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (tc *TemplateCreate) OnConflictColumns(columns ...string) *TemplateUpsertOne {
	tc.conflict = append(tc.conflict, sql.ConflictColumns(columns...))
	return &TemplateUpsertOne{
		create: tc,
	}
}

type (
	// TemplateUpsertOne is the builder for "upsert"-ing
	//  one Template node.
	TemplateUpsertOne struct {
		create *TemplateCreate
	}

	// TemplateUpsert is the "OnConflict" setter.
	TemplateUpsert struct {
		*sql.UpdateSet
	}
)

// SetStatus sets the "status" field.
func (u *TemplateUpsert) SetStatus(v string) *TemplateUpsert {
	u.Set(template.FieldStatus, v)
	return u
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *TemplateUpsert) UpdateStatus() *TemplateUpsert {
	u.SetExcluded(template.FieldStatus)
	return u
}

// ClearStatus clears the value of the "status" field.
func (u *TemplateUpsert) ClearStatus() *TemplateUpsert {
	u.SetNull(template.FieldStatus)
	return u
}

// SetStatusMessage sets the "statusMessage" field.
func (u *TemplateUpsert) SetStatusMessage(v string) *TemplateUpsert {
	u.Set(template.FieldStatusMessage, v)
	return u
}

// UpdateStatusMessage sets the "statusMessage" field to the value that was provided on create.
func (u *TemplateUpsert) UpdateStatusMessage() *TemplateUpsert {
	u.SetExcluded(template.FieldStatusMessage)
	return u
}

// ClearStatusMessage clears the value of the "statusMessage" field.
func (u *TemplateUpsert) ClearStatusMessage() *TemplateUpsert {
	u.SetNull(template.FieldStatusMessage)
	return u
}

// SetUpdateTime sets the "updateTime" field.
func (u *TemplateUpsert) SetUpdateTime(v time.Time) *TemplateUpsert {
	u.Set(template.FieldUpdateTime, v)
	return u
}

// UpdateUpdateTime sets the "updateTime" field to the value that was provided on create.
func (u *TemplateUpsert) UpdateUpdateTime() *TemplateUpsert {
	u.SetExcluded(template.FieldUpdateTime)
	return u
}

// SetDescription sets the "description" field.
func (u *TemplateUpsert) SetDescription(v string) *TemplateUpsert {
	u.Set(template.FieldDescription, v)
	return u
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *TemplateUpsert) UpdateDescription() *TemplateUpsert {
	u.SetExcluded(template.FieldDescription)
	return u
}

// ClearDescription clears the value of the "description" field.
func (u *TemplateUpsert) ClearDescription() *TemplateUpsert {
	u.SetNull(template.FieldDescription)
	return u
}

// SetIcon sets the "icon" field.
func (u *TemplateUpsert) SetIcon(v string) *TemplateUpsert {
	u.Set(template.FieldIcon, v)
	return u
}

// UpdateIcon sets the "icon" field to the value that was provided on create.
func (u *TemplateUpsert) UpdateIcon() *TemplateUpsert {
	u.SetExcluded(template.FieldIcon)
	return u
}

// ClearIcon clears the value of the "icon" field.
func (u *TemplateUpsert) ClearIcon() *TemplateUpsert {
	u.SetNull(template.FieldIcon)
	return u
}

// SetLabels sets the "labels" field.
func (u *TemplateUpsert) SetLabels(v map[string]string) *TemplateUpsert {
	u.Set(template.FieldLabels, v)
	return u
}

// UpdateLabels sets the "labels" field to the value that was provided on create.
func (u *TemplateUpsert) UpdateLabels() *TemplateUpsert {
	u.SetExcluded(template.FieldLabels)
	return u
}

// SetSource sets the "source" field.
func (u *TemplateUpsert) SetSource(v string) *TemplateUpsert {
	u.Set(template.FieldSource, v)
	return u
}

// UpdateSource sets the "source" field to the value that was provided on create.
func (u *TemplateUpsert) UpdateSource() *TemplateUpsert {
	u.SetExcluded(template.FieldSource)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Template.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(template.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *TemplateUpsertOne) UpdateNewValues() *TemplateUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(template.FieldID)
		}
		if _, exists := u.create.mutation.CreateTime(); exists {
			s.SetIgnore(template.FieldCreateTime)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Template.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *TemplateUpsertOne) Ignore() *TemplateUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TemplateUpsertOne) DoNothing() *TemplateUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TemplateCreate.OnConflict
// documentation for more info.
func (u *TemplateUpsertOne) Update(set func(*TemplateUpsert)) *TemplateUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TemplateUpsert{UpdateSet: update})
	}))
	return u
}

// SetStatus sets the "status" field.
func (u *TemplateUpsertOne) SetStatus(v string) *TemplateUpsertOne {
	return u.Update(func(s *TemplateUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *TemplateUpsertOne) UpdateStatus() *TemplateUpsertOne {
	return u.Update(func(s *TemplateUpsert) {
		s.UpdateStatus()
	})
}

// ClearStatus clears the value of the "status" field.
func (u *TemplateUpsertOne) ClearStatus() *TemplateUpsertOne {
	return u.Update(func(s *TemplateUpsert) {
		s.ClearStatus()
	})
}

// SetStatusMessage sets the "statusMessage" field.
func (u *TemplateUpsertOne) SetStatusMessage(v string) *TemplateUpsertOne {
	return u.Update(func(s *TemplateUpsert) {
		s.SetStatusMessage(v)
	})
}

// UpdateStatusMessage sets the "statusMessage" field to the value that was provided on create.
func (u *TemplateUpsertOne) UpdateStatusMessage() *TemplateUpsertOne {
	return u.Update(func(s *TemplateUpsert) {
		s.UpdateStatusMessage()
	})
}

// ClearStatusMessage clears the value of the "statusMessage" field.
func (u *TemplateUpsertOne) ClearStatusMessage() *TemplateUpsertOne {
	return u.Update(func(s *TemplateUpsert) {
		s.ClearStatusMessage()
	})
}

// SetUpdateTime sets the "updateTime" field.
func (u *TemplateUpsertOne) SetUpdateTime(v time.Time) *TemplateUpsertOne {
	return u.Update(func(s *TemplateUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "updateTime" field to the value that was provided on create.
func (u *TemplateUpsertOne) UpdateUpdateTime() *TemplateUpsertOne {
	return u.Update(func(s *TemplateUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetDescription sets the "description" field.
func (u *TemplateUpsertOne) SetDescription(v string) *TemplateUpsertOne {
	return u.Update(func(s *TemplateUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *TemplateUpsertOne) UpdateDescription() *TemplateUpsertOne {
	return u.Update(func(s *TemplateUpsert) {
		s.UpdateDescription()
	})
}

// ClearDescription clears the value of the "description" field.
func (u *TemplateUpsertOne) ClearDescription() *TemplateUpsertOne {
	return u.Update(func(s *TemplateUpsert) {
		s.ClearDescription()
	})
}

// SetIcon sets the "icon" field.
func (u *TemplateUpsertOne) SetIcon(v string) *TemplateUpsertOne {
	return u.Update(func(s *TemplateUpsert) {
		s.SetIcon(v)
	})
}

// UpdateIcon sets the "icon" field to the value that was provided on create.
func (u *TemplateUpsertOne) UpdateIcon() *TemplateUpsertOne {
	return u.Update(func(s *TemplateUpsert) {
		s.UpdateIcon()
	})
}

// ClearIcon clears the value of the "icon" field.
func (u *TemplateUpsertOne) ClearIcon() *TemplateUpsertOne {
	return u.Update(func(s *TemplateUpsert) {
		s.ClearIcon()
	})
}

// SetLabels sets the "labels" field.
func (u *TemplateUpsertOne) SetLabels(v map[string]string) *TemplateUpsertOne {
	return u.Update(func(s *TemplateUpsert) {
		s.SetLabels(v)
	})
}

// UpdateLabels sets the "labels" field to the value that was provided on create.
func (u *TemplateUpsertOne) UpdateLabels() *TemplateUpsertOne {
	return u.Update(func(s *TemplateUpsert) {
		s.UpdateLabels()
	})
}

// SetSource sets the "source" field.
func (u *TemplateUpsertOne) SetSource(v string) *TemplateUpsertOne {
	return u.Update(func(s *TemplateUpsert) {
		s.SetSource(v)
	})
}

// UpdateSource sets the "source" field to the value that was provided on create.
func (u *TemplateUpsertOne) UpdateSource() *TemplateUpsertOne {
	return u.Update(func(s *TemplateUpsert) {
		s.UpdateSource()
	})
}

// Exec executes the query.
func (u *TemplateUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("model: missing options for TemplateCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TemplateUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *TemplateUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("model: TemplateUpsertOne.ID is not supported by MySQL driver. Use TemplateUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *TemplateUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// TemplateCreateBulk is the builder for creating many Template entities in bulk.
type TemplateCreateBulk struct {
	config
	builders []*TemplateCreate
	conflict []sql.ConflictOption
}

// Save creates the Template entities in the database.
func (tcb *TemplateCreateBulk) Save(ctx context.Context) ([]*Template, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Template, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TemplateMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = tcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TemplateCreateBulk) SaveX(ctx context.Context) []*Template {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TemplateCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TemplateCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Template.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TemplateUpsert) {
//			SetStatus(v+v).
//		}).
//		Exec(ctx)
func (tcb *TemplateCreateBulk) OnConflict(opts ...sql.ConflictOption) *TemplateUpsertBulk {
	tcb.conflict = opts
	return &TemplateUpsertBulk{
		create: tcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Template.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (tcb *TemplateCreateBulk) OnConflictColumns(columns ...string) *TemplateUpsertBulk {
	tcb.conflict = append(tcb.conflict, sql.ConflictColumns(columns...))
	return &TemplateUpsertBulk{
		create: tcb,
	}
}

// TemplateUpsertBulk is the builder for "upsert"-ing
// a bulk of Template nodes.
type TemplateUpsertBulk struct {
	create *TemplateCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Template.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(template.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *TemplateUpsertBulk) UpdateNewValues() *TemplateUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(template.FieldID)
			}
			if _, exists := b.mutation.CreateTime(); exists {
				s.SetIgnore(template.FieldCreateTime)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Template.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *TemplateUpsertBulk) Ignore() *TemplateUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TemplateUpsertBulk) DoNothing() *TemplateUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TemplateCreateBulk.OnConflict
// documentation for more info.
func (u *TemplateUpsertBulk) Update(set func(*TemplateUpsert)) *TemplateUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TemplateUpsert{UpdateSet: update})
	}))
	return u
}

// SetStatus sets the "status" field.
func (u *TemplateUpsertBulk) SetStatus(v string) *TemplateUpsertBulk {
	return u.Update(func(s *TemplateUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *TemplateUpsertBulk) UpdateStatus() *TemplateUpsertBulk {
	return u.Update(func(s *TemplateUpsert) {
		s.UpdateStatus()
	})
}

// ClearStatus clears the value of the "status" field.
func (u *TemplateUpsertBulk) ClearStatus() *TemplateUpsertBulk {
	return u.Update(func(s *TemplateUpsert) {
		s.ClearStatus()
	})
}

// SetStatusMessage sets the "statusMessage" field.
func (u *TemplateUpsertBulk) SetStatusMessage(v string) *TemplateUpsertBulk {
	return u.Update(func(s *TemplateUpsert) {
		s.SetStatusMessage(v)
	})
}

// UpdateStatusMessage sets the "statusMessage" field to the value that was provided on create.
func (u *TemplateUpsertBulk) UpdateStatusMessage() *TemplateUpsertBulk {
	return u.Update(func(s *TemplateUpsert) {
		s.UpdateStatusMessage()
	})
}

// ClearStatusMessage clears the value of the "statusMessage" field.
func (u *TemplateUpsertBulk) ClearStatusMessage() *TemplateUpsertBulk {
	return u.Update(func(s *TemplateUpsert) {
		s.ClearStatusMessage()
	})
}

// SetUpdateTime sets the "updateTime" field.
func (u *TemplateUpsertBulk) SetUpdateTime(v time.Time) *TemplateUpsertBulk {
	return u.Update(func(s *TemplateUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "updateTime" field to the value that was provided on create.
func (u *TemplateUpsertBulk) UpdateUpdateTime() *TemplateUpsertBulk {
	return u.Update(func(s *TemplateUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetDescription sets the "description" field.
func (u *TemplateUpsertBulk) SetDescription(v string) *TemplateUpsertBulk {
	return u.Update(func(s *TemplateUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *TemplateUpsertBulk) UpdateDescription() *TemplateUpsertBulk {
	return u.Update(func(s *TemplateUpsert) {
		s.UpdateDescription()
	})
}

// ClearDescription clears the value of the "description" field.
func (u *TemplateUpsertBulk) ClearDescription() *TemplateUpsertBulk {
	return u.Update(func(s *TemplateUpsert) {
		s.ClearDescription()
	})
}

// SetIcon sets the "icon" field.
func (u *TemplateUpsertBulk) SetIcon(v string) *TemplateUpsertBulk {
	return u.Update(func(s *TemplateUpsert) {
		s.SetIcon(v)
	})
}

// UpdateIcon sets the "icon" field to the value that was provided on create.
func (u *TemplateUpsertBulk) UpdateIcon() *TemplateUpsertBulk {
	return u.Update(func(s *TemplateUpsert) {
		s.UpdateIcon()
	})
}

// ClearIcon clears the value of the "icon" field.
func (u *TemplateUpsertBulk) ClearIcon() *TemplateUpsertBulk {
	return u.Update(func(s *TemplateUpsert) {
		s.ClearIcon()
	})
}

// SetLabels sets the "labels" field.
func (u *TemplateUpsertBulk) SetLabels(v map[string]string) *TemplateUpsertBulk {
	return u.Update(func(s *TemplateUpsert) {
		s.SetLabels(v)
	})
}

// UpdateLabels sets the "labels" field to the value that was provided on create.
func (u *TemplateUpsertBulk) UpdateLabels() *TemplateUpsertBulk {
	return u.Update(func(s *TemplateUpsert) {
		s.UpdateLabels()
	})
}

// SetSource sets the "source" field.
func (u *TemplateUpsertBulk) SetSource(v string) *TemplateUpsertBulk {
	return u.Update(func(s *TemplateUpsert) {
		s.SetSource(v)
	})
}

// UpdateSource sets the "source" field to the value that was provided on create.
func (u *TemplateUpsertBulk) UpdateSource() *TemplateUpsertBulk {
	return u.Update(func(s *TemplateUpsert) {
		s.UpdateSource()
	})
}

// Exec executes the query.
func (u *TemplateUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("model: OnConflict was set for builder %d. Set it on the TemplateCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("model: missing options for TemplateCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TemplateUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
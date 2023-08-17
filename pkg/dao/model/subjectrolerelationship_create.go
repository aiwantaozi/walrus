// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus". DO NOT EDIT.

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

	"github.com/seal-io/walrus/pkg/dao/model/project"
	"github.com/seal-io/walrus/pkg/dao/model/role"
	"github.com/seal-io/walrus/pkg/dao/model/subject"
	"github.com/seal-io/walrus/pkg/dao/model/subjectrolerelationship"
	"github.com/seal-io/walrus/pkg/dao/types/object"
)

// SubjectRoleRelationshipCreate is the builder for creating a SubjectRoleRelationship entity.
type SubjectRoleRelationshipCreate struct {
	config
	mutation   *SubjectRoleRelationshipMutation
	hooks      []Hook
	conflict   []sql.ConflictOption
	object     *SubjectRoleRelationship
	fromUpsert bool
}

// SetCreateTime sets the "create_time" field.
func (srrc *SubjectRoleRelationshipCreate) SetCreateTime(t time.Time) *SubjectRoleRelationshipCreate {
	srrc.mutation.SetCreateTime(t)
	return srrc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (srrc *SubjectRoleRelationshipCreate) SetNillableCreateTime(t *time.Time) *SubjectRoleRelationshipCreate {
	if t != nil {
		srrc.SetCreateTime(*t)
	}
	return srrc
}

// SetProjectID sets the "project_id" field.
func (srrc *SubjectRoleRelationshipCreate) SetProjectID(o object.ID) *SubjectRoleRelationshipCreate {
	srrc.mutation.SetProjectID(o)
	return srrc
}

// SetNillableProjectID sets the "project_id" field if the given value is not nil.
func (srrc *SubjectRoleRelationshipCreate) SetNillableProjectID(o *object.ID) *SubjectRoleRelationshipCreate {
	if o != nil {
		srrc.SetProjectID(*o)
	}
	return srrc
}

// SetSubjectID sets the "subject_id" field.
func (srrc *SubjectRoleRelationshipCreate) SetSubjectID(o object.ID) *SubjectRoleRelationshipCreate {
	srrc.mutation.SetSubjectID(o)
	return srrc
}

// SetRoleID sets the "role_id" field.
func (srrc *SubjectRoleRelationshipCreate) SetRoleID(s string) *SubjectRoleRelationshipCreate {
	srrc.mutation.SetRoleID(s)
	return srrc
}

// SetID sets the "id" field.
func (srrc *SubjectRoleRelationshipCreate) SetID(o object.ID) *SubjectRoleRelationshipCreate {
	srrc.mutation.SetID(o)
	return srrc
}

// SetProject sets the "project" edge to the Project entity.
func (srrc *SubjectRoleRelationshipCreate) SetProject(p *Project) *SubjectRoleRelationshipCreate {
	return srrc.SetProjectID(p.ID)
}

// SetSubject sets the "subject" edge to the Subject entity.
func (srrc *SubjectRoleRelationshipCreate) SetSubject(s *Subject) *SubjectRoleRelationshipCreate {
	return srrc.SetSubjectID(s.ID)
}

// SetRole sets the "role" edge to the Role entity.
func (srrc *SubjectRoleRelationshipCreate) SetRole(r *Role) *SubjectRoleRelationshipCreate {
	return srrc.SetRoleID(r.ID)
}

// Mutation returns the SubjectRoleRelationshipMutation object of the builder.
func (srrc *SubjectRoleRelationshipCreate) Mutation() *SubjectRoleRelationshipMutation {
	return srrc.mutation
}

// Save creates the SubjectRoleRelationship in the database.
func (srrc *SubjectRoleRelationshipCreate) Save(ctx context.Context) (*SubjectRoleRelationship, error) {
	if err := srrc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, srrc.sqlSave, srrc.mutation, srrc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (srrc *SubjectRoleRelationshipCreate) SaveX(ctx context.Context) *SubjectRoleRelationship {
	v, err := srrc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (srrc *SubjectRoleRelationshipCreate) Exec(ctx context.Context) error {
	_, err := srrc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (srrc *SubjectRoleRelationshipCreate) ExecX(ctx context.Context) {
	if err := srrc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (srrc *SubjectRoleRelationshipCreate) defaults() error {
	if _, ok := srrc.mutation.CreateTime(); !ok {
		if subjectrolerelationship.DefaultCreateTime == nil {
			return fmt.Errorf("model: uninitialized subjectrolerelationship.DefaultCreateTime (forgotten import model/runtime?)")
		}
		v := subjectrolerelationship.DefaultCreateTime()
		srrc.mutation.SetCreateTime(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (srrc *SubjectRoleRelationshipCreate) check() error {
	if _, ok := srrc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`model: missing required field "SubjectRoleRelationship.create_time"`)}
	}
	if _, ok := srrc.mutation.SubjectID(); !ok {
		return &ValidationError{Name: "subject_id", err: errors.New(`model: missing required field "SubjectRoleRelationship.subject_id"`)}
	}
	if v, ok := srrc.mutation.SubjectID(); ok {
		if err := subjectrolerelationship.SubjectIDValidator(string(v)); err != nil {
			return &ValidationError{Name: "subject_id", err: fmt.Errorf(`model: validator failed for field "SubjectRoleRelationship.subject_id": %w`, err)}
		}
	}
	if _, ok := srrc.mutation.RoleID(); !ok {
		return &ValidationError{Name: "role_id", err: errors.New(`model: missing required field "SubjectRoleRelationship.role_id"`)}
	}
	if v, ok := srrc.mutation.RoleID(); ok {
		if err := subjectrolerelationship.RoleIDValidator(v); err != nil {
			return &ValidationError{Name: "role_id", err: fmt.Errorf(`model: validator failed for field "SubjectRoleRelationship.role_id": %w`, err)}
		}
	}
	if _, ok := srrc.mutation.SubjectID(); !ok {
		return &ValidationError{Name: "subject", err: errors.New(`model: missing required edge "SubjectRoleRelationship.subject"`)}
	}
	if _, ok := srrc.mutation.RoleID(); !ok {
		return &ValidationError{Name: "role", err: errors.New(`model: missing required edge "SubjectRoleRelationship.role"`)}
	}
	return nil
}

func (srrc *SubjectRoleRelationshipCreate) sqlSave(ctx context.Context) (*SubjectRoleRelationship, error) {
	if err := srrc.check(); err != nil {
		return nil, err
	}
	_node, _spec := srrc.createSpec()
	if err := sqlgraph.CreateNode(ctx, srrc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*object.ID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	srrc.mutation.id = &_node.ID
	srrc.mutation.done = true
	return _node, nil
}

func (srrc *SubjectRoleRelationshipCreate) createSpec() (*SubjectRoleRelationship, *sqlgraph.CreateSpec) {
	var (
		_node = &SubjectRoleRelationship{config: srrc.config}
		_spec = sqlgraph.NewCreateSpec(subjectrolerelationship.Table, sqlgraph.NewFieldSpec(subjectrolerelationship.FieldID, field.TypeString))
	)
	_spec.Schema = srrc.schemaConfig.SubjectRoleRelationship
	_spec.OnConflict = srrc.conflict
	if id, ok := srrc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := srrc.mutation.CreateTime(); ok {
		_spec.SetField(subjectrolerelationship.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = &value
	}
	if nodes := srrc.mutation.ProjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   subjectrolerelationship.ProjectTable,
			Columns: []string{subjectrolerelationship.ProjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(project.FieldID, field.TypeString),
			},
		}
		edge.Schema = srrc.schemaConfig.SubjectRoleRelationship
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ProjectID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := srrc.mutation.SubjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   subjectrolerelationship.SubjectTable,
			Columns: []string{subjectrolerelationship.SubjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(subject.FieldID, field.TypeString),
			},
		}
		edge.Schema = srrc.schemaConfig.SubjectRoleRelationship
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.SubjectID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := srrc.mutation.RoleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   subjectrolerelationship.RoleTable,
			Columns: []string{subjectrolerelationship.RoleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeString),
			},
		}
		edge.Schema = srrc.schemaConfig.SubjectRoleRelationship
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.RoleID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// Set is different from other Set* methods,
// it sets the value by judging the definition of each field within the entire object.
//
// For required fields, Set calls directly.
//
// For optional fields, Set calls if the value is not zero.
//
// For example:
//
//	## Required
//
//	db.SetX(obj.X)
//
//	## Optional or Default
//
//	if _is_zero_value_(obj.X) {
//	   db.SetX(obj.X)
//	}
func (srrc *SubjectRoleRelationshipCreate) Set(obj *SubjectRoleRelationship) *SubjectRoleRelationshipCreate {
	// Required.
	srrc.SetSubjectID(obj.SubjectID)
	srrc.SetRoleID(obj.RoleID)

	// Optional.
	if obj.CreateTime != nil {
		srrc.SetCreateTime(*obj.CreateTime)
	}
	if obj.ProjectID != "" {
		srrc.SetProjectID(obj.ProjectID)
	}

	// Record the given object.
	srrc.object = obj

	return srrc
}

// getClientSet returns the ClientSet for the given builder.
func (srrc *SubjectRoleRelationshipCreate) getClientSet() (mc ClientSet) {
	if _, ok := srrc.config.driver.(*txDriver); ok {
		tx := &Tx{config: srrc.config}
		tx.init()
		mc = tx
	} else {
		cli := &Client{config: srrc.config}
		cli.init()
		mc = cli
	}
	return mc
}

// SaveE calls the given function after created the SubjectRoleRelationship entity,
// which is always good for cascading create operations.
func (srrc *SubjectRoleRelationshipCreate) SaveE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *SubjectRoleRelationship) error) (*SubjectRoleRelationship, error) {
	obj, err := srrc.Save(ctx)
	if err != nil {
		return nil, err
	}

	if len(cbs) == 0 {
		return obj, err
	}

	mc := srrc.getClientSet()
	if srrc.fromUpsert {
		q := mc.SubjectRoleRelationships().Query().
			Where(
				subjectrolerelationship.SubjectID(obj.SubjectID),
				subjectrolerelationship.RoleID(obj.RoleID),
			)
		if obj.ProjectID != "" {
			q.Where(subjectrolerelationship.ProjectID(obj.ProjectID))
		} else {
			q.Where(subjectrolerelationship.ProjectIDIsNil())
		}
		obj.ID, err = q.OnlyID(ctx)
		if err != nil {
			return nil, fmt.Errorf("model: failed to query id of SubjectRoleRelationship entity: %w", err)
		}
	}

	if x := srrc.object; x != nil {
		if _, set := srrc.mutation.Field(subjectrolerelationship.FieldProjectID); set {
			obj.ProjectID = x.ProjectID
		}
		if _, set := srrc.mutation.Field(subjectrolerelationship.FieldSubjectID); set {
			obj.SubjectID = x.SubjectID
		}
		if _, set := srrc.mutation.Field(subjectrolerelationship.FieldRoleID); set {
			obj.RoleID = x.RoleID
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
func (srrc *SubjectRoleRelationshipCreate) SaveEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *SubjectRoleRelationship) error) *SubjectRoleRelationship {
	obj, err := srrc.SaveE(ctx, cbs...)
	if err != nil {
		panic(err)
	}
	return obj
}

// ExecE calls the given function after executed the query,
// which is always good for cascading create operations.
func (srrc *SubjectRoleRelationshipCreate) ExecE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *SubjectRoleRelationship) error) error {
	_, err := srrc.SaveE(ctx, cbs...)
	return err
}

// ExecEX is like ExecE, but panics if an error occurs.
func (srrc *SubjectRoleRelationshipCreate) ExecEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *SubjectRoleRelationship) error) {
	if err := srrc.ExecE(ctx, cbs...); err != nil {
		panic(err)
	}
}

// Set leverages the SubjectRoleRelationshipCreate Set method,
// it sets the value by judging the definition of each field within the entire item of the given list.
//
// For required fields, Set calls directly.
//
// For optional fields, Set calls if the value is not zero.
//
// For example:
//
//	## Required
//
//	db.SetX(obj.X)
//
//	## Optional or Default
//
//	if _is_zero_value_(obj.X) {
//	   db.SetX(obj.X)
//	}
func (srrcb *SubjectRoleRelationshipCreateBulk) Set(objs ...*SubjectRoleRelationship) *SubjectRoleRelationshipCreateBulk {
	if len(objs) != 0 {
		client := NewSubjectRoleRelationshipClient(srrcb.config)

		srrcb.builders = make([]*SubjectRoleRelationshipCreate, len(objs))
		for i := range objs {
			srrcb.builders[i] = client.Create().Set(objs[i])
		}

		// Record the given objects.
		srrcb.objects = objs
	}

	return srrcb
}

// getClientSet returns the ClientSet for the given builder.
func (srrcb *SubjectRoleRelationshipCreateBulk) getClientSet() (mc ClientSet) {
	if _, ok := srrcb.config.driver.(*txDriver); ok {
		tx := &Tx{config: srrcb.config}
		tx.init()
		mc = tx
	} else {
		cli := &Client{config: srrcb.config}
		cli.init()
		mc = cli
	}
	return mc
}

// SaveE calls the given function after created the SubjectRoleRelationship entities,
// which is always good for cascading create operations.
func (srrcb *SubjectRoleRelationshipCreateBulk) SaveE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *SubjectRoleRelationship) error) ([]*SubjectRoleRelationship, error) {
	objs, err := srrcb.Save(ctx)
	if err != nil {
		return nil, err
	}

	if len(cbs) == 0 {
		return objs, err
	}

	mc := srrcb.getClientSet()
	if srrcb.fromUpsert {
		for i := range objs {
			obj := objs[i]
			q := mc.SubjectRoleRelationships().Query().
				Where(
					subjectrolerelationship.SubjectID(obj.SubjectID),
					subjectrolerelationship.RoleID(obj.RoleID),
				)
			if obj.ProjectID != "" {
				q.Where(subjectrolerelationship.ProjectID(obj.ProjectID))
			} else {
				q.Where(subjectrolerelationship.ProjectIDIsNil())
			}
			objs[i].ID, err = q.OnlyID(ctx)
			if err != nil {
				return nil, fmt.Errorf("model: failed to query id of SubjectRoleRelationship entity: %w", err)
			}
		}
	}

	if x := srrcb.objects; x != nil {
		for i := range x {
			if _, set := srrcb.builders[i].mutation.Field(subjectrolerelationship.FieldProjectID); set {
				objs[i].ProjectID = x[i].ProjectID
			}
			if _, set := srrcb.builders[i].mutation.Field(subjectrolerelationship.FieldSubjectID); set {
				objs[i].SubjectID = x[i].SubjectID
			}
			if _, set := srrcb.builders[i].mutation.Field(subjectrolerelationship.FieldRoleID); set {
				objs[i].RoleID = x[i].RoleID
			}
			objs[i].Edges = x[i].Edges
		}
	}

	for i := range objs {
		for j := range cbs {
			if err = cbs[j](ctx, mc, objs[i]); err != nil {
				return nil, err
			}
		}
	}

	return objs, nil
}

// SaveEX is like SaveE, but panics if an error occurs.
func (srrcb *SubjectRoleRelationshipCreateBulk) SaveEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *SubjectRoleRelationship) error) []*SubjectRoleRelationship {
	objs, err := srrcb.SaveE(ctx, cbs...)
	if err != nil {
		panic(err)
	}
	return objs
}

// ExecE calls the given function after executed the query,
// which is always good for cascading create operations.
func (srrcb *SubjectRoleRelationshipCreateBulk) ExecE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *SubjectRoleRelationship) error) error {
	_, err := srrcb.SaveE(ctx, cbs...)
	return err
}

// ExecEX is like ExecE, but panics if an error occurs.
func (srrcb *SubjectRoleRelationshipCreateBulk) ExecEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *SubjectRoleRelationship) error) {
	if err := srrcb.ExecE(ctx, cbs...); err != nil {
		panic(err)
	}
}

// ExecE calls the given function after executed the query,
// which is always good for cascading create operations.
func (u *SubjectRoleRelationshipUpsertOne) ExecE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *SubjectRoleRelationship) error) error {
	if len(u.create.conflict) == 0 {
		return errors.New("model: missing options for SubjectRoleRelationshipUpsertOne.OnConflict")
	}
	u.create.fromUpsert = true
	return u.create.ExecE(ctx, cbs...)
}

// ExecEX is like ExecE, but panics if an error occurs.
func (u *SubjectRoleRelationshipUpsertOne) ExecEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *SubjectRoleRelationship) error) {
	if err := u.ExecE(ctx, cbs...); err != nil {
		panic(err)
	}
}

// ExecE calls the given function after executed the query,
// which is always good for cascading create operations.
func (u *SubjectRoleRelationshipUpsertBulk) ExecE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *SubjectRoleRelationship) error) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("model: OnConflict was set for builder %d. Set it on the SubjectRoleRelationshipUpsertBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("model: missing options for SubjectRoleRelationshipUpsertBulk.OnConflict")
	}
	u.create.fromUpsert = true
	return u.create.ExecE(ctx, cbs...)
}

// ExecEX is like ExecE, but panics if an error occurs.
func (u *SubjectRoleRelationshipUpsertBulk) ExecEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *SubjectRoleRelationship) error) {
	if err := u.ExecE(ctx, cbs...); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.SubjectRoleRelationship.Create().
//		SetCreateTime(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SubjectRoleRelationshipUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (srrc *SubjectRoleRelationshipCreate) OnConflict(opts ...sql.ConflictOption) *SubjectRoleRelationshipUpsertOne {
	srrc.conflict = opts
	return &SubjectRoleRelationshipUpsertOne{
		create: srrc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.SubjectRoleRelationship.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (srrc *SubjectRoleRelationshipCreate) OnConflictColumns(columns ...string) *SubjectRoleRelationshipUpsertOne {
	srrc.conflict = append(srrc.conflict, sql.ConflictColumns(columns...))
	return &SubjectRoleRelationshipUpsertOne{
		create: srrc,
	}
}

type (
	// SubjectRoleRelationshipUpsertOne is the builder for "upsert"-ing
	//  one SubjectRoleRelationship node.
	SubjectRoleRelationshipUpsertOne struct {
		create *SubjectRoleRelationshipCreate
	}

	// SubjectRoleRelationshipUpsert is the "OnConflict" setter.
	SubjectRoleRelationshipUpsert struct {
		*sql.UpdateSet
	}
)

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.SubjectRoleRelationship.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(subjectrolerelationship.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *SubjectRoleRelationshipUpsertOne) UpdateNewValues() *SubjectRoleRelationshipUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(subjectrolerelationship.FieldID)
		}
		if _, exists := u.create.mutation.CreateTime(); exists {
			s.SetIgnore(subjectrolerelationship.FieldCreateTime)
		}
		if _, exists := u.create.mutation.ProjectID(); exists {
			s.SetIgnore(subjectrolerelationship.FieldProjectID)
		}
		if _, exists := u.create.mutation.SubjectID(); exists {
			s.SetIgnore(subjectrolerelationship.FieldSubjectID)
		}
		if _, exists := u.create.mutation.RoleID(); exists {
			s.SetIgnore(subjectrolerelationship.FieldRoleID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.SubjectRoleRelationship.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *SubjectRoleRelationshipUpsertOne) Ignore() *SubjectRoleRelationshipUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SubjectRoleRelationshipUpsertOne) DoNothing() *SubjectRoleRelationshipUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SubjectRoleRelationshipCreate.OnConflict
// documentation for more info.
func (u *SubjectRoleRelationshipUpsertOne) Update(set func(*SubjectRoleRelationshipUpsert)) *SubjectRoleRelationshipUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SubjectRoleRelationshipUpsert{UpdateSet: update})
	}))
	return u
}

// Exec executes the query.
func (u *SubjectRoleRelationshipUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("model: missing options for SubjectRoleRelationshipCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SubjectRoleRelationshipUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *SubjectRoleRelationshipUpsertOne) ID(ctx context.Context) (id object.ID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("model: SubjectRoleRelationshipUpsertOne.ID is not supported by MySQL driver. Use SubjectRoleRelationshipUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *SubjectRoleRelationshipUpsertOne) IDX(ctx context.Context) object.ID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// SubjectRoleRelationshipCreateBulk is the builder for creating many SubjectRoleRelationship entities in bulk.
type SubjectRoleRelationshipCreateBulk struct {
	config
	builders   []*SubjectRoleRelationshipCreate
	conflict   []sql.ConflictOption
	objects    []*SubjectRoleRelationship
	fromUpsert bool
}

// Save creates the SubjectRoleRelationship entities in the database.
func (srrcb *SubjectRoleRelationshipCreateBulk) Save(ctx context.Context) ([]*SubjectRoleRelationship, error) {
	specs := make([]*sqlgraph.CreateSpec, len(srrcb.builders))
	nodes := make([]*SubjectRoleRelationship, len(srrcb.builders))
	mutators := make([]Mutator, len(srrcb.builders))
	for i := range srrcb.builders {
		func(i int, root context.Context) {
			builder := srrcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SubjectRoleRelationshipMutation)
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
					_, err = mutators[i+1].Mutate(root, srrcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = srrcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, srrcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, srrcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (srrcb *SubjectRoleRelationshipCreateBulk) SaveX(ctx context.Context) []*SubjectRoleRelationship {
	v, err := srrcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (srrcb *SubjectRoleRelationshipCreateBulk) Exec(ctx context.Context) error {
	_, err := srrcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (srrcb *SubjectRoleRelationshipCreateBulk) ExecX(ctx context.Context) {
	if err := srrcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.SubjectRoleRelationship.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SubjectRoleRelationshipUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (srrcb *SubjectRoleRelationshipCreateBulk) OnConflict(opts ...sql.ConflictOption) *SubjectRoleRelationshipUpsertBulk {
	srrcb.conflict = opts
	return &SubjectRoleRelationshipUpsertBulk{
		create: srrcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.SubjectRoleRelationship.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (srrcb *SubjectRoleRelationshipCreateBulk) OnConflictColumns(columns ...string) *SubjectRoleRelationshipUpsertBulk {
	srrcb.conflict = append(srrcb.conflict, sql.ConflictColumns(columns...))
	return &SubjectRoleRelationshipUpsertBulk{
		create: srrcb,
	}
}

// SubjectRoleRelationshipUpsertBulk is the builder for "upsert"-ing
// a bulk of SubjectRoleRelationship nodes.
type SubjectRoleRelationshipUpsertBulk struct {
	create *SubjectRoleRelationshipCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.SubjectRoleRelationship.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(subjectrolerelationship.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *SubjectRoleRelationshipUpsertBulk) UpdateNewValues() *SubjectRoleRelationshipUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(subjectrolerelationship.FieldID)
			}
			if _, exists := b.mutation.CreateTime(); exists {
				s.SetIgnore(subjectrolerelationship.FieldCreateTime)
			}
			if _, exists := b.mutation.ProjectID(); exists {
				s.SetIgnore(subjectrolerelationship.FieldProjectID)
			}
			if _, exists := b.mutation.SubjectID(); exists {
				s.SetIgnore(subjectrolerelationship.FieldSubjectID)
			}
			if _, exists := b.mutation.RoleID(); exists {
				s.SetIgnore(subjectrolerelationship.FieldRoleID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.SubjectRoleRelationship.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *SubjectRoleRelationshipUpsertBulk) Ignore() *SubjectRoleRelationshipUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SubjectRoleRelationshipUpsertBulk) DoNothing() *SubjectRoleRelationshipUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SubjectRoleRelationshipCreateBulk.OnConflict
// documentation for more info.
func (u *SubjectRoleRelationshipUpsertBulk) Update(set func(*SubjectRoleRelationshipUpsert)) *SubjectRoleRelationshipUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SubjectRoleRelationshipUpsert{UpdateSet: update})
	}))
	return u
}

// Exec executes the query.
func (u *SubjectRoleRelationshipUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("model: OnConflict was set for builder %d. Set it on the SubjectRoleRelationshipCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("model: missing options for SubjectRoleRelationshipCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SubjectRoleRelationshipUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

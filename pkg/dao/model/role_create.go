// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus". DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"

	"github.com/seal-io/walrus/pkg/dao/model/role"
	"github.com/seal-io/walrus/pkg/dao/model/subjectrolerelationship"
	"github.com/seal-io/walrus/pkg/dao/types"
	"github.com/seal-io/walrus/pkg/dao/types/object"
)

// RoleCreate is the builder for creating a Role entity.
type RoleCreate struct {
	config
	mutation   *RoleMutation
	hooks      []Hook
	conflict   []sql.ConflictOption
	object     *Role
	fromUpsert bool
}

// SetCreateTime sets the "create_time" field.
func (rc *RoleCreate) SetCreateTime(t time.Time) *RoleCreate {
	rc.mutation.SetCreateTime(t)
	return rc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (rc *RoleCreate) SetNillableCreateTime(t *time.Time) *RoleCreate {
	if t != nil {
		rc.SetCreateTime(*t)
	}
	return rc
}

// SetUpdateTime sets the "update_time" field.
func (rc *RoleCreate) SetUpdateTime(t time.Time) *RoleCreate {
	rc.mutation.SetUpdateTime(t)
	return rc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (rc *RoleCreate) SetNillableUpdateTime(t *time.Time) *RoleCreate {
	if t != nil {
		rc.SetUpdateTime(*t)
	}
	return rc
}

// SetKind sets the "kind" field.
func (rc *RoleCreate) SetKind(s string) *RoleCreate {
	rc.mutation.SetKind(s)
	return rc
}

// SetNillableKind sets the "kind" field if the given value is not nil.
func (rc *RoleCreate) SetNillableKind(s *string) *RoleCreate {
	if s != nil {
		rc.SetKind(*s)
	}
	return rc
}

// SetDescription sets the "description" field.
func (rc *RoleCreate) SetDescription(s string) *RoleCreate {
	rc.mutation.SetDescription(s)
	return rc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (rc *RoleCreate) SetNillableDescription(s *string) *RoleCreate {
	if s != nil {
		rc.SetDescription(*s)
	}
	return rc
}

// SetPolicies sets the "policies" field.
func (rc *RoleCreate) SetPolicies(tp types.RolePolicies) *RoleCreate {
	rc.mutation.SetPolicies(tp)
	return rc
}

// SetApplicableEnvironmentTypes sets the "applicable_environment_types" field.
func (rc *RoleCreate) SetApplicableEnvironmentTypes(s []string) *RoleCreate {
	rc.mutation.SetApplicableEnvironmentTypes(s)
	return rc
}

// SetSession sets the "session" field.
func (rc *RoleCreate) SetSession(b bool) *RoleCreate {
	rc.mutation.SetSession(b)
	return rc
}

// SetNillableSession sets the "session" field if the given value is not nil.
func (rc *RoleCreate) SetNillableSession(b *bool) *RoleCreate {
	if b != nil {
		rc.SetSession(*b)
	}
	return rc
}

// SetBuiltin sets the "builtin" field.
func (rc *RoleCreate) SetBuiltin(b bool) *RoleCreate {
	rc.mutation.SetBuiltin(b)
	return rc
}

// SetNillableBuiltin sets the "builtin" field if the given value is not nil.
func (rc *RoleCreate) SetNillableBuiltin(b *bool) *RoleCreate {
	if b != nil {
		rc.SetBuiltin(*b)
	}
	return rc
}

// SetID sets the "id" field.
func (rc *RoleCreate) SetID(s string) *RoleCreate {
	rc.mutation.SetID(s)
	return rc
}

// AddSubjectIDs adds the "subjects" edge to the SubjectRoleRelationship entity by IDs.
func (rc *RoleCreate) AddSubjectIDs(ids ...object.ID) *RoleCreate {
	rc.mutation.AddSubjectIDs(ids...)
	return rc
}

// AddSubjects adds the "subjects" edges to the SubjectRoleRelationship entity.
func (rc *RoleCreate) AddSubjects(s ...*SubjectRoleRelationship) *RoleCreate {
	ids := make([]object.ID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return rc.AddSubjectIDs(ids...)
}

// Mutation returns the RoleMutation object of the builder.
func (rc *RoleCreate) Mutation() *RoleMutation {
	return rc.mutation
}

// Save creates the Role in the database.
func (rc *RoleCreate) Save(ctx context.Context) (*Role, error) {
	if err := rc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, rc.sqlSave, rc.mutation, rc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rc *RoleCreate) SaveX(ctx context.Context) *Role {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *RoleCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *RoleCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rc *RoleCreate) defaults() error {
	if _, ok := rc.mutation.CreateTime(); !ok {
		if role.DefaultCreateTime == nil {
			return fmt.Errorf("model: uninitialized role.DefaultCreateTime (forgotten import model/runtime?)")
		}
		v := role.DefaultCreateTime()
		rc.mutation.SetCreateTime(v)
	}
	if _, ok := rc.mutation.UpdateTime(); !ok {
		if role.DefaultUpdateTime == nil {
			return fmt.Errorf("model: uninitialized role.DefaultUpdateTime (forgotten import model/runtime?)")
		}
		v := role.DefaultUpdateTime()
		rc.mutation.SetUpdateTime(v)
	}
	if _, ok := rc.mutation.Kind(); !ok {
		v := role.DefaultKind
		rc.mutation.SetKind(v)
	}
	if _, ok := rc.mutation.Policies(); !ok {
		v := role.DefaultPolicies
		rc.mutation.SetPolicies(v)
	}
	if _, ok := rc.mutation.ApplicableEnvironmentTypes(); !ok {
		v := role.DefaultApplicableEnvironmentTypes
		rc.mutation.SetApplicableEnvironmentTypes(v)
	}
	if _, ok := rc.mutation.Session(); !ok {
		v := role.DefaultSession
		rc.mutation.SetSession(v)
	}
	if _, ok := rc.mutation.Builtin(); !ok {
		v := role.DefaultBuiltin
		rc.mutation.SetBuiltin(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (rc *RoleCreate) check() error {
	if _, ok := rc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`model: missing required field "Role.create_time"`)}
	}
	if _, ok := rc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`model: missing required field "Role.update_time"`)}
	}
	if _, ok := rc.mutation.Kind(); !ok {
		return &ValidationError{Name: "kind", err: errors.New(`model: missing required field "Role.kind"`)}
	}
	if _, ok := rc.mutation.Policies(); !ok {
		return &ValidationError{Name: "policies", err: errors.New(`model: missing required field "Role.policies"`)}
	}
	if _, ok := rc.mutation.Session(); !ok {
		return &ValidationError{Name: "session", err: errors.New(`model: missing required field "Role.session"`)}
	}
	if _, ok := rc.mutation.Builtin(); !ok {
		return &ValidationError{Name: "builtin", err: errors.New(`model: missing required field "Role.builtin"`)}
	}
	if v, ok := rc.mutation.ID(); ok {
		if err := role.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`model: validator failed for field "Role.id": %w`, err)}
		}
	}
	return nil
}

func (rc *RoleCreate) sqlSave(ctx context.Context) (*Role, error) {
	if err := rc.check(); err != nil {
		return nil, err
	}
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Role.ID type: %T", _spec.ID.Value)
		}
	}
	rc.mutation.id = &_node.ID
	rc.mutation.done = true
	return _node, nil
}

func (rc *RoleCreate) createSpec() (*Role, *sqlgraph.CreateSpec) {
	var (
		_node = &Role{config: rc.config}
		_spec = sqlgraph.NewCreateSpec(role.Table, sqlgraph.NewFieldSpec(role.FieldID, field.TypeString))
	)
	_spec.Schema = rc.schemaConfig.Role
	_spec.OnConflict = rc.conflict
	if id, ok := rc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := rc.mutation.CreateTime(); ok {
		_spec.SetField(role.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = &value
	}
	if value, ok := rc.mutation.UpdateTime(); ok {
		_spec.SetField(role.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = &value
	}
	if value, ok := rc.mutation.Kind(); ok {
		_spec.SetField(role.FieldKind, field.TypeString, value)
		_node.Kind = value
	}
	if value, ok := rc.mutation.Description(); ok {
		_spec.SetField(role.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := rc.mutation.Policies(); ok {
		_spec.SetField(role.FieldPolicies, field.TypeJSON, value)
		_node.Policies = value
	}
	if value, ok := rc.mutation.ApplicableEnvironmentTypes(); ok {
		_spec.SetField(role.FieldApplicableEnvironmentTypes, field.TypeJSON, value)
		_node.ApplicableEnvironmentTypes = value
	}
	if value, ok := rc.mutation.Session(); ok {
		_spec.SetField(role.FieldSession, field.TypeBool, value)
		_node.Session = value
	}
	if value, ok := rc.mutation.Builtin(); ok {
		_spec.SetField(role.FieldBuiltin, field.TypeBool, value)
		_node.Builtin = value
	}
	if nodes := rc.mutation.SubjectsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   role.SubjectsTable,
			Columns: []string{role.SubjectsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(subjectrolerelationship.FieldID, field.TypeString),
			},
		}
		edge.Schema = rc.schemaConfig.SubjectRoleRelationship
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
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
func (rc *RoleCreate) Set(obj *Role) *RoleCreate {
	// Required.
	rc.SetID(obj.ID)
	rc.SetKind(obj.Kind)
	rc.SetPolicies(obj.Policies)
	rc.SetSession(obj.Session)
	rc.SetBuiltin(obj.Builtin)

	// Optional.
	if obj.CreateTime != nil {
		rc.SetCreateTime(*obj.CreateTime)
	}
	if obj.UpdateTime != nil {
		rc.SetUpdateTime(*obj.UpdateTime)
	}
	if obj.Description != "" {
		rc.SetDescription(obj.Description)
	}
	if !reflect.ValueOf(obj.ApplicableEnvironmentTypes).IsZero() {
		rc.SetApplicableEnvironmentTypes(obj.ApplicableEnvironmentTypes)
	}

	// Record the given object.
	rc.object = obj

	return rc
}

// getClientSet returns the ClientSet for the given builder.
func (rc *RoleCreate) getClientSet() (mc ClientSet) {
	if _, ok := rc.config.driver.(*txDriver); ok {
		tx := &Tx{config: rc.config}
		tx.init()
		mc = tx
	} else {
		cli := &Client{config: rc.config}
		cli.init()
		mc = cli
	}
	return mc
}

// SaveE calls the given function after created the Role entity,
// which is always good for cascading create operations.
func (rc *RoleCreate) SaveE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *Role) error) (*Role, error) {
	obj, err := rc.Save(ctx)
	if err != nil {
		return nil, err
	}

	if len(cbs) == 0 {
		return obj, err
	}

	mc := rc.getClientSet()

	if x := rc.object; x != nil {
		if _, set := rc.mutation.Field(role.FieldDescription); set {
			obj.Description = x.Description
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
func (rc *RoleCreate) SaveEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *Role) error) *Role {
	obj, err := rc.SaveE(ctx, cbs...)
	if err != nil {
		panic(err)
	}
	return obj
}

// ExecE calls the given function after executed the query,
// which is always good for cascading create operations.
func (rc *RoleCreate) ExecE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *Role) error) error {
	_, err := rc.SaveE(ctx, cbs...)
	return err
}

// ExecEX is like ExecE, but panics if an error occurs.
func (rc *RoleCreate) ExecEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *Role) error) {
	if err := rc.ExecE(ctx, cbs...); err != nil {
		panic(err)
	}
}

// Set leverages the RoleCreate Set method,
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
func (rcb *RoleCreateBulk) Set(objs ...*Role) *RoleCreateBulk {
	if len(objs) != 0 {
		client := NewRoleClient(rcb.config)

		rcb.builders = make([]*RoleCreate, len(objs))
		for i := range objs {
			rcb.builders[i] = client.Create().Set(objs[i])
		}

		// Record the given objects.
		rcb.objects = objs
	}

	return rcb
}

// getClientSet returns the ClientSet for the given builder.
func (rcb *RoleCreateBulk) getClientSet() (mc ClientSet) {
	if _, ok := rcb.config.driver.(*txDriver); ok {
		tx := &Tx{config: rcb.config}
		tx.init()
		mc = tx
	} else {
		cli := &Client{config: rcb.config}
		cli.init()
		mc = cli
	}
	return mc
}

// SaveE calls the given function after created the Role entities,
// which is always good for cascading create operations.
func (rcb *RoleCreateBulk) SaveE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *Role) error) ([]*Role, error) {
	objs, err := rcb.Save(ctx)
	if err != nil {
		return nil, err
	}

	if len(cbs) == 0 {
		return objs, err
	}

	mc := rcb.getClientSet()

	if x := rcb.objects; x != nil {
		for i := range x {
			if _, set := rcb.builders[i].mutation.Field(role.FieldDescription); set {
				objs[i].Description = x[i].Description
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
func (rcb *RoleCreateBulk) SaveEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *Role) error) []*Role {
	objs, err := rcb.SaveE(ctx, cbs...)
	if err != nil {
		panic(err)
	}
	return objs
}

// ExecE calls the given function after executed the query,
// which is always good for cascading create operations.
func (rcb *RoleCreateBulk) ExecE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *Role) error) error {
	_, err := rcb.SaveE(ctx, cbs...)
	return err
}

// ExecEX is like ExecE, but panics if an error occurs.
func (rcb *RoleCreateBulk) ExecEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *Role) error) {
	if err := rcb.ExecE(ctx, cbs...); err != nil {
		panic(err)
	}
}

// ExecE calls the given function after executed the query,
// which is always good for cascading create operations.
func (u *RoleUpsertOne) ExecE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *Role) error) error {
	if len(u.create.conflict) == 0 {
		return errors.New("model: missing options for RoleUpsertOne.OnConflict")
	}
	u.create.fromUpsert = true
	return u.create.ExecE(ctx, cbs...)
}

// ExecEX is like ExecE, but panics if an error occurs.
func (u *RoleUpsertOne) ExecEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *Role) error) {
	if err := u.ExecE(ctx, cbs...); err != nil {
		panic(err)
	}
}

// ExecE calls the given function after executed the query,
// which is always good for cascading create operations.
func (u *RoleUpsertBulk) ExecE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *Role) error) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("model: OnConflict was set for builder %d. Set it on the RoleUpsertBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("model: missing options for RoleUpsertBulk.OnConflict")
	}
	u.create.fromUpsert = true
	return u.create.ExecE(ctx, cbs...)
}

// ExecEX is like ExecE, but panics if an error occurs.
func (u *RoleUpsertBulk) ExecEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *Role) error) {
	if err := u.ExecE(ctx, cbs...); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Role.Create().
//		SetCreateTime(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.RoleUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (rc *RoleCreate) OnConflict(opts ...sql.ConflictOption) *RoleUpsertOne {
	rc.conflict = opts
	return &RoleUpsertOne{
		create: rc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Role.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (rc *RoleCreate) OnConflictColumns(columns ...string) *RoleUpsertOne {
	rc.conflict = append(rc.conflict, sql.ConflictColumns(columns...))
	return &RoleUpsertOne{
		create: rc,
	}
}

type (
	// RoleUpsertOne is the builder for "upsert"-ing
	//  one Role node.
	RoleUpsertOne struct {
		create *RoleCreate
	}

	// RoleUpsert is the "OnConflict" setter.
	RoleUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdateTime sets the "update_time" field.
func (u *RoleUpsert) SetUpdateTime(v time.Time) *RoleUpsert {
	u.Set(role.FieldUpdateTime, v)
	return u
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *RoleUpsert) UpdateUpdateTime() *RoleUpsert {
	u.SetExcluded(role.FieldUpdateTime)
	return u
}

// SetDescription sets the "description" field.
func (u *RoleUpsert) SetDescription(v string) *RoleUpsert {
	u.Set(role.FieldDescription, v)
	return u
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *RoleUpsert) UpdateDescription() *RoleUpsert {
	u.SetExcluded(role.FieldDescription)
	return u
}

// ClearDescription clears the value of the "description" field.
func (u *RoleUpsert) ClearDescription() *RoleUpsert {
	u.SetNull(role.FieldDescription)
	return u
}

// SetPolicies sets the "policies" field.
func (u *RoleUpsert) SetPolicies(v types.RolePolicies) *RoleUpsert {
	u.Set(role.FieldPolicies, v)
	return u
}

// UpdatePolicies sets the "policies" field to the value that was provided on create.
func (u *RoleUpsert) UpdatePolicies() *RoleUpsert {
	u.SetExcluded(role.FieldPolicies)
	return u
}

// SetApplicableEnvironmentTypes sets the "applicable_environment_types" field.
func (u *RoleUpsert) SetApplicableEnvironmentTypes(v []string) *RoleUpsert {
	u.Set(role.FieldApplicableEnvironmentTypes, v)
	return u
}

// UpdateApplicableEnvironmentTypes sets the "applicable_environment_types" field to the value that was provided on create.
func (u *RoleUpsert) UpdateApplicableEnvironmentTypes() *RoleUpsert {
	u.SetExcluded(role.FieldApplicableEnvironmentTypes)
	return u
}

// ClearApplicableEnvironmentTypes clears the value of the "applicable_environment_types" field.
func (u *RoleUpsert) ClearApplicableEnvironmentTypes() *RoleUpsert {
	u.SetNull(role.FieldApplicableEnvironmentTypes)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Role.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(role.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *RoleUpsertOne) UpdateNewValues() *RoleUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(role.FieldID)
		}
		if _, exists := u.create.mutation.CreateTime(); exists {
			s.SetIgnore(role.FieldCreateTime)
		}
		if _, exists := u.create.mutation.Kind(); exists {
			s.SetIgnore(role.FieldKind)
		}
		if _, exists := u.create.mutation.Session(); exists {
			s.SetIgnore(role.FieldSession)
		}
		if _, exists := u.create.mutation.Builtin(); exists {
			s.SetIgnore(role.FieldBuiltin)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Role.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *RoleUpsertOne) Ignore() *RoleUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *RoleUpsertOne) DoNothing() *RoleUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the RoleCreate.OnConflict
// documentation for more info.
func (u *RoleUpsertOne) Update(set func(*RoleUpsert)) *RoleUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&RoleUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *RoleUpsertOne) SetUpdateTime(v time.Time) *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *RoleUpsertOne) UpdateUpdateTime() *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetDescription sets the "description" field.
func (u *RoleUpsertOne) SetDescription(v string) *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *RoleUpsertOne) UpdateDescription() *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.UpdateDescription()
	})
}

// ClearDescription clears the value of the "description" field.
func (u *RoleUpsertOne) ClearDescription() *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.ClearDescription()
	})
}

// SetPolicies sets the "policies" field.
func (u *RoleUpsertOne) SetPolicies(v types.RolePolicies) *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.SetPolicies(v)
	})
}

// UpdatePolicies sets the "policies" field to the value that was provided on create.
func (u *RoleUpsertOne) UpdatePolicies() *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.UpdatePolicies()
	})
}

// SetApplicableEnvironmentTypes sets the "applicable_environment_types" field.
func (u *RoleUpsertOne) SetApplicableEnvironmentTypes(v []string) *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.SetApplicableEnvironmentTypes(v)
	})
}

// UpdateApplicableEnvironmentTypes sets the "applicable_environment_types" field to the value that was provided on create.
func (u *RoleUpsertOne) UpdateApplicableEnvironmentTypes() *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.UpdateApplicableEnvironmentTypes()
	})
}

// ClearApplicableEnvironmentTypes clears the value of the "applicable_environment_types" field.
func (u *RoleUpsertOne) ClearApplicableEnvironmentTypes() *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.ClearApplicableEnvironmentTypes()
	})
}

// Exec executes the query.
func (u *RoleUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("model: missing options for RoleCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *RoleUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *RoleUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("model: RoleUpsertOne.ID is not supported by MySQL driver. Use RoleUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *RoleUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// RoleCreateBulk is the builder for creating many Role entities in bulk.
type RoleCreateBulk struct {
	config
	builders   []*RoleCreate
	conflict   []sql.ConflictOption
	objects    []*Role
	fromUpsert bool
}

// Save creates the Role entities in the database.
func (rcb *RoleCreateBulk) Save(ctx context.Context) ([]*Role, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Role, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RoleMutation)
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
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = rcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *RoleCreateBulk) SaveX(ctx context.Context) []*Role {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *RoleCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *RoleCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Role.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.RoleUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (rcb *RoleCreateBulk) OnConflict(opts ...sql.ConflictOption) *RoleUpsertBulk {
	rcb.conflict = opts
	return &RoleUpsertBulk{
		create: rcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Role.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (rcb *RoleCreateBulk) OnConflictColumns(columns ...string) *RoleUpsertBulk {
	rcb.conflict = append(rcb.conflict, sql.ConflictColumns(columns...))
	return &RoleUpsertBulk{
		create: rcb,
	}
}

// RoleUpsertBulk is the builder for "upsert"-ing
// a bulk of Role nodes.
type RoleUpsertBulk struct {
	create *RoleCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Role.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(role.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *RoleUpsertBulk) UpdateNewValues() *RoleUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(role.FieldID)
			}
			if _, exists := b.mutation.CreateTime(); exists {
				s.SetIgnore(role.FieldCreateTime)
			}
			if _, exists := b.mutation.Kind(); exists {
				s.SetIgnore(role.FieldKind)
			}
			if _, exists := b.mutation.Session(); exists {
				s.SetIgnore(role.FieldSession)
			}
			if _, exists := b.mutation.Builtin(); exists {
				s.SetIgnore(role.FieldBuiltin)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Role.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *RoleUpsertBulk) Ignore() *RoleUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *RoleUpsertBulk) DoNothing() *RoleUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the RoleCreateBulk.OnConflict
// documentation for more info.
func (u *RoleUpsertBulk) Update(set func(*RoleUpsert)) *RoleUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&RoleUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *RoleUpsertBulk) SetUpdateTime(v time.Time) *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *RoleUpsertBulk) UpdateUpdateTime() *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetDescription sets the "description" field.
func (u *RoleUpsertBulk) SetDescription(v string) *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *RoleUpsertBulk) UpdateDescription() *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.UpdateDescription()
	})
}

// ClearDescription clears the value of the "description" field.
func (u *RoleUpsertBulk) ClearDescription() *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.ClearDescription()
	})
}

// SetPolicies sets the "policies" field.
func (u *RoleUpsertBulk) SetPolicies(v types.RolePolicies) *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.SetPolicies(v)
	})
}

// UpdatePolicies sets the "policies" field to the value that was provided on create.
func (u *RoleUpsertBulk) UpdatePolicies() *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.UpdatePolicies()
	})
}

// SetApplicableEnvironmentTypes sets the "applicable_environment_types" field.
func (u *RoleUpsertBulk) SetApplicableEnvironmentTypes(v []string) *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.SetApplicableEnvironmentTypes(v)
	})
}

// UpdateApplicableEnvironmentTypes sets the "applicable_environment_types" field to the value that was provided on create.
func (u *RoleUpsertBulk) UpdateApplicableEnvironmentTypes() *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.UpdateApplicableEnvironmentTypes()
	})
}

// ClearApplicableEnvironmentTypes clears the value of the "applicable_environment_types" field.
func (u *RoleUpsertBulk) ClearApplicableEnvironmentTypes() *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.ClearApplicableEnvironmentTypes()
	})
}

// Exec executes the query.
func (u *RoleUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("model: OnConflict was set for builder %d. Set it on the RoleCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("model: missing options for RoleCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *RoleUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

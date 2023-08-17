// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus". DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"

	"github.com/seal-io/walrus/pkg/dao/model/internal"
	"github.com/seal-io/walrus/pkg/dao/model/predicate"
	"github.com/seal-io/walrus/pkg/dao/model/servicerelationship"
)

// ServiceRelationshipUpdate is the builder for updating ServiceRelationship entities.
type ServiceRelationshipUpdate struct {
	config
	hooks     []Hook
	mutation  *ServiceRelationshipMutation
	modifiers []func(*sql.UpdateBuilder)
	object    *ServiceRelationship
}

// Where appends a list predicates to the ServiceRelationshipUpdate builder.
func (sru *ServiceRelationshipUpdate) Where(ps ...predicate.ServiceRelationship) *ServiceRelationshipUpdate {
	sru.mutation.Where(ps...)
	return sru
}

// Mutation returns the ServiceRelationshipMutation object of the builder.
func (sru *ServiceRelationshipUpdate) Mutation() *ServiceRelationshipMutation {
	return sru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (sru *ServiceRelationshipUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, sru.sqlSave, sru.mutation, sru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (sru *ServiceRelationshipUpdate) SaveX(ctx context.Context) int {
	affected, err := sru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (sru *ServiceRelationshipUpdate) Exec(ctx context.Context) error {
	_, err := sru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sru *ServiceRelationshipUpdate) ExecX(ctx context.Context) {
	if err := sru.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sru *ServiceRelationshipUpdate) check() error {
	if _, ok := sru.mutation.ServiceID(); sru.mutation.ServiceCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "ServiceRelationship.service"`)
	}
	if _, ok := sru.mutation.DependencyID(); sru.mutation.DependencyCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "ServiceRelationship.dependency"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (sru *ServiceRelationshipUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ServiceRelationshipUpdate {
	sru.modifiers = append(sru.modifiers, modifiers...)
	return sru
}

func (sru *ServiceRelationshipUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := sru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(servicerelationship.Table, servicerelationship.Columns, sqlgraph.NewFieldSpec(servicerelationship.FieldID, field.TypeString))
	if ps := sru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	_spec.Node.Schema = sru.schemaConfig.ServiceRelationship
	ctx = internal.NewSchemaConfigContext(ctx, sru.schemaConfig)
	_spec.AddModifiers(sru.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, sru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{servicerelationship.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	sru.mutation.done = true
	return n, nil
}

// ServiceRelationshipUpdateOne is the builder for updating a single ServiceRelationship entity.
type ServiceRelationshipUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *ServiceRelationshipMutation
	modifiers []func(*sql.UpdateBuilder)
	object    *ServiceRelationship
}

// Mutation returns the ServiceRelationshipMutation object of the builder.
func (sruo *ServiceRelationshipUpdateOne) Mutation() *ServiceRelationshipMutation {
	return sruo.mutation
}

// Where appends a list predicates to the ServiceRelationshipUpdate builder.
func (sruo *ServiceRelationshipUpdateOne) Where(ps ...predicate.ServiceRelationship) *ServiceRelationshipUpdateOne {
	sruo.mutation.Where(ps...)
	return sruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (sruo *ServiceRelationshipUpdateOne) Select(field string, fields ...string) *ServiceRelationshipUpdateOne {
	sruo.fields = append([]string{field}, fields...)
	return sruo
}

// Save executes the query and returns the updated ServiceRelationship entity.
func (sruo *ServiceRelationshipUpdateOne) Save(ctx context.Context) (*ServiceRelationship, error) {
	return withHooks(ctx, sruo.sqlSave, sruo.mutation, sruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (sruo *ServiceRelationshipUpdateOne) SaveX(ctx context.Context) *ServiceRelationship {
	node, err := sruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (sruo *ServiceRelationshipUpdateOne) Exec(ctx context.Context) error {
	_, err := sruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sruo *ServiceRelationshipUpdateOne) ExecX(ctx context.Context) {
	if err := sruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sruo *ServiceRelationshipUpdateOne) check() error {
	if _, ok := sruo.mutation.ServiceID(); sruo.mutation.ServiceCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "ServiceRelationship.service"`)
	}
	if _, ok := sruo.mutation.DependencyID(); sruo.mutation.DependencyCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "ServiceRelationship.dependency"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (sruo *ServiceRelationshipUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ServiceRelationshipUpdateOne {
	sruo.modifiers = append(sruo.modifiers, modifiers...)
	return sruo
}

func (sruo *ServiceRelationshipUpdateOne) sqlSave(ctx context.Context) (_node *ServiceRelationship, err error) {
	if err := sruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(servicerelationship.Table, servicerelationship.Columns, sqlgraph.NewFieldSpec(servicerelationship.FieldID, field.TypeString))
	id, ok := sruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`model: missing "ServiceRelationship.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := sruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, servicerelationship.FieldID)
		for _, f := range fields {
			if !servicerelationship.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("model: invalid field %q for query", f)}
			}
			if f != servicerelationship.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := sruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	_spec.Node.Schema = sruo.schemaConfig.ServiceRelationship
	ctx = internal.NewSchemaConfigContext(ctx, sruo.schemaConfig)
	_spec.AddModifiers(sruo.modifiers...)
	_node = &ServiceRelationship{config: sruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, sruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{servicerelationship.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	sruo.mutation.done = true
	return _node, nil
}

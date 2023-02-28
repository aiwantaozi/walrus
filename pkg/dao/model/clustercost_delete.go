// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// GENERATED, DO NOT EDIT.

package model

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"

	"github.com/seal-io/seal/pkg/dao/model/clustercost"
	"github.com/seal-io/seal/pkg/dao/model/internal"
	"github.com/seal-io/seal/pkg/dao/model/predicate"
)

// ClusterCostDelete is the builder for deleting a ClusterCost entity.
type ClusterCostDelete struct {
	config
	hooks    []Hook
	mutation *ClusterCostMutation
}

// Where appends a list predicates to the ClusterCostDelete builder.
func (ccd *ClusterCostDelete) Where(ps ...predicate.ClusterCost) *ClusterCostDelete {
	ccd.mutation.Where(ps...)
	return ccd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ccd *ClusterCostDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, ClusterCostMutation](ctx, ccd.sqlExec, ccd.mutation, ccd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ccd *ClusterCostDelete) ExecX(ctx context.Context) int {
	n, err := ccd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ccd *ClusterCostDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: clustercost.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: clustercost.FieldID,
			},
		},
	}
	_spec.Node.Schema = ccd.schemaConfig.ClusterCost
	ctx = internal.NewSchemaConfigContext(ctx, ccd.schemaConfig)
	if ps := ccd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ccd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ccd.mutation.done = true
	return affected, err
}

// ClusterCostDeleteOne is the builder for deleting a single ClusterCost entity.
type ClusterCostDeleteOne struct {
	ccd *ClusterCostDelete
}

// Where appends a list predicates to the ClusterCostDelete builder.
func (ccdo *ClusterCostDeleteOne) Where(ps ...predicate.ClusterCost) *ClusterCostDeleteOne {
	ccdo.ccd.mutation.Where(ps...)
	return ccdo
}

// Exec executes the deletion query.
func (ccdo *ClusterCostDeleteOne) Exec(ctx context.Context) error {
	n, err := ccdo.ccd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{clustercost.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ccdo *ClusterCostDeleteOne) ExecX(ctx context.Context) {
	if err := ccdo.Exec(ctx); err != nil {
		panic(err)
	}
}
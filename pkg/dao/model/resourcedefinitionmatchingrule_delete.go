// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus". DO NOT EDIT.

package model

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"

	"github.com/seal-io/walrus/pkg/dao/model/internal"
	"github.com/seal-io/walrus/pkg/dao/model/predicate"
	"github.com/seal-io/walrus/pkg/dao/model/resourcedefinitionmatchingrule"
)

// ResourceDefinitionMatchingRuleDelete is the builder for deleting a ResourceDefinitionMatchingRule entity.
type ResourceDefinitionMatchingRuleDelete struct {
	config
	hooks    []Hook
	mutation *ResourceDefinitionMatchingRuleMutation
}

// Where appends a list predicates to the ResourceDefinitionMatchingRuleDelete builder.
func (rdmrd *ResourceDefinitionMatchingRuleDelete) Where(ps ...predicate.ResourceDefinitionMatchingRule) *ResourceDefinitionMatchingRuleDelete {
	rdmrd.mutation.Where(ps...)
	return rdmrd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (rdmrd *ResourceDefinitionMatchingRuleDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, rdmrd.sqlExec, rdmrd.mutation, rdmrd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (rdmrd *ResourceDefinitionMatchingRuleDelete) ExecX(ctx context.Context) int {
	n, err := rdmrd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (rdmrd *ResourceDefinitionMatchingRuleDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(resourcedefinitionmatchingrule.Table, sqlgraph.NewFieldSpec(resourcedefinitionmatchingrule.FieldID, field.TypeString))
	_spec.Node.Schema = rdmrd.schemaConfig.ResourceDefinitionMatchingRule
	ctx = internal.NewSchemaConfigContext(ctx, rdmrd.schemaConfig)
	if ps := rdmrd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, rdmrd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	rdmrd.mutation.done = true
	return affected, err
}

// ResourceDefinitionMatchingRuleDeleteOne is the builder for deleting a single ResourceDefinitionMatchingRule entity.
type ResourceDefinitionMatchingRuleDeleteOne struct {
	rdmrd *ResourceDefinitionMatchingRuleDelete
}

// Where appends a list predicates to the ResourceDefinitionMatchingRuleDelete builder.
func (rdmrdo *ResourceDefinitionMatchingRuleDeleteOne) Where(ps ...predicate.ResourceDefinitionMatchingRule) *ResourceDefinitionMatchingRuleDeleteOne {
	rdmrdo.rdmrd.mutation.Where(ps...)
	return rdmrdo
}

// Exec executes the deletion query.
func (rdmrdo *ResourceDefinitionMatchingRuleDeleteOne) Exec(ctx context.Context) error {
	n, err := rdmrdo.rdmrd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{resourcedefinitionmatchingrule.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (rdmrdo *ResourceDefinitionMatchingRuleDeleteOne) ExecX(ctx context.Context) {
	if err := rdmrdo.Exec(ctx); err != nil {
		panic(err)
	}
}

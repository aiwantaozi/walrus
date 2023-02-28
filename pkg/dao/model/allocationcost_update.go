// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// GENERATED, DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"

	"github.com/seal-io/seal/pkg/dao/model/allocationcost"
	"github.com/seal-io/seal/pkg/dao/model/internal"
	"github.com/seal-io/seal/pkg/dao/model/predicate"
)

// AllocationCostUpdate is the builder for updating AllocationCost entities.
type AllocationCostUpdate struct {
	config
	hooks     []Hook
	mutation  *AllocationCostMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the AllocationCostUpdate builder.
func (acu *AllocationCostUpdate) Where(ps ...predicate.AllocationCost) *AllocationCostUpdate {
	acu.mutation.Where(ps...)
	return acu
}

// SetTotalCost sets the "totalCost" field.
func (acu *AllocationCostUpdate) SetTotalCost(f float64) *AllocationCostUpdate {
	acu.mutation.ResetTotalCost()
	acu.mutation.SetTotalCost(f)
	return acu
}

// SetNillableTotalCost sets the "totalCost" field if the given value is not nil.
func (acu *AllocationCostUpdate) SetNillableTotalCost(f *float64) *AllocationCostUpdate {
	if f != nil {
		acu.SetTotalCost(*f)
	}
	return acu
}

// AddTotalCost adds f to the "totalCost" field.
func (acu *AllocationCostUpdate) AddTotalCost(f float64) *AllocationCostUpdate {
	acu.mutation.AddTotalCost(f)
	return acu
}

// SetCurrency sets the "currency" field.
func (acu *AllocationCostUpdate) SetCurrency(i int) *AllocationCostUpdate {
	acu.mutation.ResetCurrency()
	acu.mutation.SetCurrency(i)
	return acu
}

// SetNillableCurrency sets the "currency" field if the given value is not nil.
func (acu *AllocationCostUpdate) SetNillableCurrency(i *int) *AllocationCostUpdate {
	if i != nil {
		acu.SetCurrency(*i)
	}
	return acu
}

// AddCurrency adds i to the "currency" field.
func (acu *AllocationCostUpdate) AddCurrency(i int) *AllocationCostUpdate {
	acu.mutation.AddCurrency(i)
	return acu
}

// ClearCurrency clears the value of the "currency" field.
func (acu *AllocationCostUpdate) ClearCurrency() *AllocationCostUpdate {
	acu.mutation.ClearCurrency()
	return acu
}

// SetCpuCost sets the "cpuCost" field.
func (acu *AllocationCostUpdate) SetCpuCost(f float64) *AllocationCostUpdate {
	acu.mutation.ResetCpuCost()
	acu.mutation.SetCpuCost(f)
	return acu
}

// SetNillableCpuCost sets the "cpuCost" field if the given value is not nil.
func (acu *AllocationCostUpdate) SetNillableCpuCost(f *float64) *AllocationCostUpdate {
	if f != nil {
		acu.SetCpuCost(*f)
	}
	return acu
}

// AddCpuCost adds f to the "cpuCost" field.
func (acu *AllocationCostUpdate) AddCpuCost(f float64) *AllocationCostUpdate {
	acu.mutation.AddCpuCost(f)
	return acu
}

// SetGpuCost sets the "gpuCost" field.
func (acu *AllocationCostUpdate) SetGpuCost(f float64) *AllocationCostUpdate {
	acu.mutation.ResetGpuCost()
	acu.mutation.SetGpuCost(f)
	return acu
}

// SetNillableGpuCost sets the "gpuCost" field if the given value is not nil.
func (acu *AllocationCostUpdate) SetNillableGpuCost(f *float64) *AllocationCostUpdate {
	if f != nil {
		acu.SetGpuCost(*f)
	}
	return acu
}

// AddGpuCost adds f to the "gpuCost" field.
func (acu *AllocationCostUpdate) AddGpuCost(f float64) *AllocationCostUpdate {
	acu.mutation.AddGpuCost(f)
	return acu
}

// SetRamCost sets the "ramCost" field.
func (acu *AllocationCostUpdate) SetRamCost(f float64) *AllocationCostUpdate {
	acu.mutation.ResetRamCost()
	acu.mutation.SetRamCost(f)
	return acu
}

// SetNillableRamCost sets the "ramCost" field if the given value is not nil.
func (acu *AllocationCostUpdate) SetNillableRamCost(f *float64) *AllocationCostUpdate {
	if f != nil {
		acu.SetRamCost(*f)
	}
	return acu
}

// AddRamCost adds f to the "ramCost" field.
func (acu *AllocationCostUpdate) AddRamCost(f float64) *AllocationCostUpdate {
	acu.mutation.AddRamCost(f)
	return acu
}

// SetPvCost sets the "pvCost" field.
func (acu *AllocationCostUpdate) SetPvCost(f float64) *AllocationCostUpdate {
	acu.mutation.ResetPvCost()
	acu.mutation.SetPvCost(f)
	return acu
}

// SetNillablePvCost sets the "pvCost" field if the given value is not nil.
func (acu *AllocationCostUpdate) SetNillablePvCost(f *float64) *AllocationCostUpdate {
	if f != nil {
		acu.SetPvCost(*f)
	}
	return acu
}

// AddPvCost adds f to the "pvCost" field.
func (acu *AllocationCostUpdate) AddPvCost(f float64) *AllocationCostUpdate {
	acu.mutation.AddPvCost(f)
	return acu
}

// SetPvBytes sets the "pvBytes" field.
func (acu *AllocationCostUpdate) SetPvBytes(f float64) *AllocationCostUpdate {
	acu.mutation.ResetPvBytes()
	acu.mutation.SetPvBytes(f)
	return acu
}

// SetNillablePvBytes sets the "pvBytes" field if the given value is not nil.
func (acu *AllocationCostUpdate) SetNillablePvBytes(f *float64) *AllocationCostUpdate {
	if f != nil {
		acu.SetPvBytes(*f)
	}
	return acu
}

// AddPvBytes adds f to the "pvBytes" field.
func (acu *AllocationCostUpdate) AddPvBytes(f float64) *AllocationCostUpdate {
	acu.mutation.AddPvBytes(f)
	return acu
}

// SetCpuCoreUsageAverage sets the "cpuCoreUsageAverage" field.
func (acu *AllocationCostUpdate) SetCpuCoreUsageAverage(f float64) *AllocationCostUpdate {
	acu.mutation.ResetCpuCoreUsageAverage()
	acu.mutation.SetCpuCoreUsageAverage(f)
	return acu
}

// SetNillableCpuCoreUsageAverage sets the "cpuCoreUsageAverage" field if the given value is not nil.
func (acu *AllocationCostUpdate) SetNillableCpuCoreUsageAverage(f *float64) *AllocationCostUpdate {
	if f != nil {
		acu.SetCpuCoreUsageAverage(*f)
	}
	return acu
}

// AddCpuCoreUsageAverage adds f to the "cpuCoreUsageAverage" field.
func (acu *AllocationCostUpdate) AddCpuCoreUsageAverage(f float64) *AllocationCostUpdate {
	acu.mutation.AddCpuCoreUsageAverage(f)
	return acu
}

// SetCpuCoreUsageMax sets the "cpuCoreUsageMax" field.
func (acu *AllocationCostUpdate) SetCpuCoreUsageMax(f float64) *AllocationCostUpdate {
	acu.mutation.ResetCpuCoreUsageMax()
	acu.mutation.SetCpuCoreUsageMax(f)
	return acu
}

// SetNillableCpuCoreUsageMax sets the "cpuCoreUsageMax" field if the given value is not nil.
func (acu *AllocationCostUpdate) SetNillableCpuCoreUsageMax(f *float64) *AllocationCostUpdate {
	if f != nil {
		acu.SetCpuCoreUsageMax(*f)
	}
	return acu
}

// AddCpuCoreUsageMax adds f to the "cpuCoreUsageMax" field.
func (acu *AllocationCostUpdate) AddCpuCoreUsageMax(f float64) *AllocationCostUpdate {
	acu.mutation.AddCpuCoreUsageMax(f)
	return acu
}

// SetRamByteUsageAverage sets the "ramByteUsageAverage" field.
func (acu *AllocationCostUpdate) SetRamByteUsageAverage(f float64) *AllocationCostUpdate {
	acu.mutation.ResetRamByteUsageAverage()
	acu.mutation.SetRamByteUsageAverage(f)
	return acu
}

// SetNillableRamByteUsageAverage sets the "ramByteUsageAverage" field if the given value is not nil.
func (acu *AllocationCostUpdate) SetNillableRamByteUsageAverage(f *float64) *AllocationCostUpdate {
	if f != nil {
		acu.SetRamByteUsageAverage(*f)
	}
	return acu
}

// AddRamByteUsageAverage adds f to the "ramByteUsageAverage" field.
func (acu *AllocationCostUpdate) AddRamByteUsageAverage(f float64) *AllocationCostUpdate {
	acu.mutation.AddRamByteUsageAverage(f)
	return acu
}

// SetRamByteUsageMax sets the "ramByteUsageMax" field.
func (acu *AllocationCostUpdate) SetRamByteUsageMax(f float64) *AllocationCostUpdate {
	acu.mutation.ResetRamByteUsageMax()
	acu.mutation.SetRamByteUsageMax(f)
	return acu
}

// SetNillableRamByteUsageMax sets the "ramByteUsageMax" field if the given value is not nil.
func (acu *AllocationCostUpdate) SetNillableRamByteUsageMax(f *float64) *AllocationCostUpdate {
	if f != nil {
		acu.SetRamByteUsageMax(*f)
	}
	return acu
}

// AddRamByteUsageMax adds f to the "ramByteUsageMax" field.
func (acu *AllocationCostUpdate) AddRamByteUsageMax(f float64) *AllocationCostUpdate {
	acu.mutation.AddRamByteUsageMax(f)
	return acu
}

// Mutation returns the AllocationCostMutation object of the builder.
func (acu *AllocationCostUpdate) Mutation() *AllocationCostMutation {
	return acu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (acu *AllocationCostUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, AllocationCostMutation](ctx, acu.sqlSave, acu.mutation, acu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (acu *AllocationCostUpdate) SaveX(ctx context.Context) int {
	affected, err := acu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (acu *AllocationCostUpdate) Exec(ctx context.Context) error {
	_, err := acu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acu *AllocationCostUpdate) ExecX(ctx context.Context) {
	if err := acu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (acu *AllocationCostUpdate) check() error {
	if v, ok := acu.mutation.TotalCost(); ok {
		if err := allocationcost.TotalCostValidator(v); err != nil {
			return &ValidationError{Name: "totalCost", err: fmt.Errorf(`model: validator failed for field "AllocationCost.totalCost": %w`, err)}
		}
	}
	if v, ok := acu.mutation.CpuCost(); ok {
		if err := allocationcost.CpuCostValidator(v); err != nil {
			return &ValidationError{Name: "cpuCost", err: fmt.Errorf(`model: validator failed for field "AllocationCost.cpuCost": %w`, err)}
		}
	}
	if v, ok := acu.mutation.GpuCost(); ok {
		if err := allocationcost.GpuCostValidator(v); err != nil {
			return &ValidationError{Name: "gpuCost", err: fmt.Errorf(`model: validator failed for field "AllocationCost.gpuCost": %w`, err)}
		}
	}
	if v, ok := acu.mutation.RamCost(); ok {
		if err := allocationcost.RamCostValidator(v); err != nil {
			return &ValidationError{Name: "ramCost", err: fmt.Errorf(`model: validator failed for field "AllocationCost.ramCost": %w`, err)}
		}
	}
	if v, ok := acu.mutation.PvCost(); ok {
		if err := allocationcost.PvCostValidator(v); err != nil {
			return &ValidationError{Name: "pvCost", err: fmt.Errorf(`model: validator failed for field "AllocationCost.pvCost": %w`, err)}
		}
	}
	if v, ok := acu.mutation.PvBytes(); ok {
		if err := allocationcost.PvBytesValidator(v); err != nil {
			return &ValidationError{Name: "pvBytes", err: fmt.Errorf(`model: validator failed for field "AllocationCost.pvBytes": %w`, err)}
		}
	}
	if v, ok := acu.mutation.CpuCoreUsageAverage(); ok {
		if err := allocationcost.CpuCoreUsageAverageValidator(v); err != nil {
			return &ValidationError{Name: "cpuCoreUsageAverage", err: fmt.Errorf(`model: validator failed for field "AllocationCost.cpuCoreUsageAverage": %w`, err)}
		}
	}
	if v, ok := acu.mutation.CpuCoreUsageMax(); ok {
		if err := allocationcost.CpuCoreUsageMaxValidator(v); err != nil {
			return &ValidationError{Name: "cpuCoreUsageMax", err: fmt.Errorf(`model: validator failed for field "AllocationCost.cpuCoreUsageMax": %w`, err)}
		}
	}
	if v, ok := acu.mutation.RamByteUsageAverage(); ok {
		if err := allocationcost.RamByteUsageAverageValidator(v); err != nil {
			return &ValidationError{Name: "ramByteUsageAverage", err: fmt.Errorf(`model: validator failed for field "AllocationCost.ramByteUsageAverage": %w`, err)}
		}
	}
	if v, ok := acu.mutation.RamByteUsageMax(); ok {
		if err := allocationcost.RamByteUsageMaxValidator(v); err != nil {
			return &ValidationError{Name: "ramByteUsageMax", err: fmt.Errorf(`model: validator failed for field "AllocationCost.ramByteUsageMax": %w`, err)}
		}
	}
	if _, ok := acu.mutation.ConnectorID(); acu.mutation.ConnectorCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "AllocationCost.connector"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (acu *AllocationCostUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *AllocationCostUpdate {
	acu.modifiers = append(acu.modifiers, modifiers...)
	return acu
}

func (acu *AllocationCostUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := acu.check(); err != nil {
		return n, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   allocationcost.Table,
			Columns: allocationcost.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: allocationcost.FieldID,
			},
		},
	}
	if ps := acu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if acu.mutation.ClusterNameCleared() {
		_spec.ClearField(allocationcost.FieldClusterName, field.TypeString)
	}
	if acu.mutation.NamespaceCleared() {
		_spec.ClearField(allocationcost.FieldNamespace, field.TypeString)
	}
	if acu.mutation.NodeCleared() {
		_spec.ClearField(allocationcost.FieldNode, field.TypeString)
	}
	if acu.mutation.ControllerCleared() {
		_spec.ClearField(allocationcost.FieldController, field.TypeString)
	}
	if acu.mutation.ControllerKindCleared() {
		_spec.ClearField(allocationcost.FieldControllerKind, field.TypeString)
	}
	if acu.mutation.PodCleared() {
		_spec.ClearField(allocationcost.FieldPod, field.TypeString)
	}
	if acu.mutation.ContainerCleared() {
		_spec.ClearField(allocationcost.FieldContainer, field.TypeString)
	}
	if value, ok := acu.mutation.TotalCost(); ok {
		_spec.SetField(allocationcost.FieldTotalCost, field.TypeFloat64, value)
	}
	if value, ok := acu.mutation.AddedTotalCost(); ok {
		_spec.AddField(allocationcost.FieldTotalCost, field.TypeFloat64, value)
	}
	if value, ok := acu.mutation.Currency(); ok {
		_spec.SetField(allocationcost.FieldCurrency, field.TypeInt, value)
	}
	if value, ok := acu.mutation.AddedCurrency(); ok {
		_spec.AddField(allocationcost.FieldCurrency, field.TypeInt, value)
	}
	if acu.mutation.CurrencyCleared() {
		_spec.ClearField(allocationcost.FieldCurrency, field.TypeInt)
	}
	if value, ok := acu.mutation.CpuCost(); ok {
		_spec.SetField(allocationcost.FieldCpuCost, field.TypeFloat64, value)
	}
	if value, ok := acu.mutation.AddedCpuCost(); ok {
		_spec.AddField(allocationcost.FieldCpuCost, field.TypeFloat64, value)
	}
	if value, ok := acu.mutation.GpuCost(); ok {
		_spec.SetField(allocationcost.FieldGpuCost, field.TypeFloat64, value)
	}
	if value, ok := acu.mutation.AddedGpuCost(); ok {
		_spec.AddField(allocationcost.FieldGpuCost, field.TypeFloat64, value)
	}
	if value, ok := acu.mutation.RamCost(); ok {
		_spec.SetField(allocationcost.FieldRamCost, field.TypeFloat64, value)
	}
	if value, ok := acu.mutation.AddedRamCost(); ok {
		_spec.AddField(allocationcost.FieldRamCost, field.TypeFloat64, value)
	}
	if value, ok := acu.mutation.PvCost(); ok {
		_spec.SetField(allocationcost.FieldPvCost, field.TypeFloat64, value)
	}
	if value, ok := acu.mutation.AddedPvCost(); ok {
		_spec.AddField(allocationcost.FieldPvCost, field.TypeFloat64, value)
	}
	if value, ok := acu.mutation.PvBytes(); ok {
		_spec.SetField(allocationcost.FieldPvBytes, field.TypeFloat64, value)
	}
	if value, ok := acu.mutation.AddedPvBytes(); ok {
		_spec.AddField(allocationcost.FieldPvBytes, field.TypeFloat64, value)
	}
	if value, ok := acu.mutation.CpuCoreUsageAverage(); ok {
		_spec.SetField(allocationcost.FieldCpuCoreUsageAverage, field.TypeFloat64, value)
	}
	if value, ok := acu.mutation.AddedCpuCoreUsageAverage(); ok {
		_spec.AddField(allocationcost.FieldCpuCoreUsageAverage, field.TypeFloat64, value)
	}
	if value, ok := acu.mutation.CpuCoreUsageMax(); ok {
		_spec.SetField(allocationcost.FieldCpuCoreUsageMax, field.TypeFloat64, value)
	}
	if value, ok := acu.mutation.AddedCpuCoreUsageMax(); ok {
		_spec.AddField(allocationcost.FieldCpuCoreUsageMax, field.TypeFloat64, value)
	}
	if value, ok := acu.mutation.RamByteUsageAverage(); ok {
		_spec.SetField(allocationcost.FieldRamByteUsageAverage, field.TypeFloat64, value)
	}
	if value, ok := acu.mutation.AddedRamByteUsageAverage(); ok {
		_spec.AddField(allocationcost.FieldRamByteUsageAverage, field.TypeFloat64, value)
	}
	if value, ok := acu.mutation.RamByteUsageMax(); ok {
		_spec.SetField(allocationcost.FieldRamByteUsageMax, field.TypeFloat64, value)
	}
	if value, ok := acu.mutation.AddedRamByteUsageMax(); ok {
		_spec.AddField(allocationcost.FieldRamByteUsageMax, field.TypeFloat64, value)
	}
	_spec.Node.Schema = acu.schemaConfig.AllocationCost
	ctx = internal.NewSchemaConfigContext(ctx, acu.schemaConfig)
	_spec.AddModifiers(acu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, acu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{allocationcost.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	acu.mutation.done = true
	return n, nil
}

// AllocationCostUpdateOne is the builder for updating a single AllocationCost entity.
type AllocationCostUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *AllocationCostMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetTotalCost sets the "totalCost" field.
func (acuo *AllocationCostUpdateOne) SetTotalCost(f float64) *AllocationCostUpdateOne {
	acuo.mutation.ResetTotalCost()
	acuo.mutation.SetTotalCost(f)
	return acuo
}

// SetNillableTotalCost sets the "totalCost" field if the given value is not nil.
func (acuo *AllocationCostUpdateOne) SetNillableTotalCost(f *float64) *AllocationCostUpdateOne {
	if f != nil {
		acuo.SetTotalCost(*f)
	}
	return acuo
}

// AddTotalCost adds f to the "totalCost" field.
func (acuo *AllocationCostUpdateOne) AddTotalCost(f float64) *AllocationCostUpdateOne {
	acuo.mutation.AddTotalCost(f)
	return acuo
}

// SetCurrency sets the "currency" field.
func (acuo *AllocationCostUpdateOne) SetCurrency(i int) *AllocationCostUpdateOne {
	acuo.mutation.ResetCurrency()
	acuo.mutation.SetCurrency(i)
	return acuo
}

// SetNillableCurrency sets the "currency" field if the given value is not nil.
func (acuo *AllocationCostUpdateOne) SetNillableCurrency(i *int) *AllocationCostUpdateOne {
	if i != nil {
		acuo.SetCurrency(*i)
	}
	return acuo
}

// AddCurrency adds i to the "currency" field.
func (acuo *AllocationCostUpdateOne) AddCurrency(i int) *AllocationCostUpdateOne {
	acuo.mutation.AddCurrency(i)
	return acuo
}

// ClearCurrency clears the value of the "currency" field.
func (acuo *AllocationCostUpdateOne) ClearCurrency() *AllocationCostUpdateOne {
	acuo.mutation.ClearCurrency()
	return acuo
}

// SetCpuCost sets the "cpuCost" field.
func (acuo *AllocationCostUpdateOne) SetCpuCost(f float64) *AllocationCostUpdateOne {
	acuo.mutation.ResetCpuCost()
	acuo.mutation.SetCpuCost(f)
	return acuo
}

// SetNillableCpuCost sets the "cpuCost" field if the given value is not nil.
func (acuo *AllocationCostUpdateOne) SetNillableCpuCost(f *float64) *AllocationCostUpdateOne {
	if f != nil {
		acuo.SetCpuCost(*f)
	}
	return acuo
}

// AddCpuCost adds f to the "cpuCost" field.
func (acuo *AllocationCostUpdateOne) AddCpuCost(f float64) *AllocationCostUpdateOne {
	acuo.mutation.AddCpuCost(f)
	return acuo
}

// SetGpuCost sets the "gpuCost" field.
func (acuo *AllocationCostUpdateOne) SetGpuCost(f float64) *AllocationCostUpdateOne {
	acuo.mutation.ResetGpuCost()
	acuo.mutation.SetGpuCost(f)
	return acuo
}

// SetNillableGpuCost sets the "gpuCost" field if the given value is not nil.
func (acuo *AllocationCostUpdateOne) SetNillableGpuCost(f *float64) *AllocationCostUpdateOne {
	if f != nil {
		acuo.SetGpuCost(*f)
	}
	return acuo
}

// AddGpuCost adds f to the "gpuCost" field.
func (acuo *AllocationCostUpdateOne) AddGpuCost(f float64) *AllocationCostUpdateOne {
	acuo.mutation.AddGpuCost(f)
	return acuo
}

// SetRamCost sets the "ramCost" field.
func (acuo *AllocationCostUpdateOne) SetRamCost(f float64) *AllocationCostUpdateOne {
	acuo.mutation.ResetRamCost()
	acuo.mutation.SetRamCost(f)
	return acuo
}

// SetNillableRamCost sets the "ramCost" field if the given value is not nil.
func (acuo *AllocationCostUpdateOne) SetNillableRamCost(f *float64) *AllocationCostUpdateOne {
	if f != nil {
		acuo.SetRamCost(*f)
	}
	return acuo
}

// AddRamCost adds f to the "ramCost" field.
func (acuo *AllocationCostUpdateOne) AddRamCost(f float64) *AllocationCostUpdateOne {
	acuo.mutation.AddRamCost(f)
	return acuo
}

// SetPvCost sets the "pvCost" field.
func (acuo *AllocationCostUpdateOne) SetPvCost(f float64) *AllocationCostUpdateOne {
	acuo.mutation.ResetPvCost()
	acuo.mutation.SetPvCost(f)
	return acuo
}

// SetNillablePvCost sets the "pvCost" field if the given value is not nil.
func (acuo *AllocationCostUpdateOne) SetNillablePvCost(f *float64) *AllocationCostUpdateOne {
	if f != nil {
		acuo.SetPvCost(*f)
	}
	return acuo
}

// AddPvCost adds f to the "pvCost" field.
func (acuo *AllocationCostUpdateOne) AddPvCost(f float64) *AllocationCostUpdateOne {
	acuo.mutation.AddPvCost(f)
	return acuo
}

// SetPvBytes sets the "pvBytes" field.
func (acuo *AllocationCostUpdateOne) SetPvBytes(f float64) *AllocationCostUpdateOne {
	acuo.mutation.ResetPvBytes()
	acuo.mutation.SetPvBytes(f)
	return acuo
}

// SetNillablePvBytes sets the "pvBytes" field if the given value is not nil.
func (acuo *AllocationCostUpdateOne) SetNillablePvBytes(f *float64) *AllocationCostUpdateOne {
	if f != nil {
		acuo.SetPvBytes(*f)
	}
	return acuo
}

// AddPvBytes adds f to the "pvBytes" field.
func (acuo *AllocationCostUpdateOne) AddPvBytes(f float64) *AllocationCostUpdateOne {
	acuo.mutation.AddPvBytes(f)
	return acuo
}

// SetCpuCoreUsageAverage sets the "cpuCoreUsageAverage" field.
func (acuo *AllocationCostUpdateOne) SetCpuCoreUsageAverage(f float64) *AllocationCostUpdateOne {
	acuo.mutation.ResetCpuCoreUsageAverage()
	acuo.mutation.SetCpuCoreUsageAverage(f)
	return acuo
}

// SetNillableCpuCoreUsageAverage sets the "cpuCoreUsageAverage" field if the given value is not nil.
func (acuo *AllocationCostUpdateOne) SetNillableCpuCoreUsageAverage(f *float64) *AllocationCostUpdateOne {
	if f != nil {
		acuo.SetCpuCoreUsageAverage(*f)
	}
	return acuo
}

// AddCpuCoreUsageAverage adds f to the "cpuCoreUsageAverage" field.
func (acuo *AllocationCostUpdateOne) AddCpuCoreUsageAverage(f float64) *AllocationCostUpdateOne {
	acuo.mutation.AddCpuCoreUsageAverage(f)
	return acuo
}

// SetCpuCoreUsageMax sets the "cpuCoreUsageMax" field.
func (acuo *AllocationCostUpdateOne) SetCpuCoreUsageMax(f float64) *AllocationCostUpdateOne {
	acuo.mutation.ResetCpuCoreUsageMax()
	acuo.mutation.SetCpuCoreUsageMax(f)
	return acuo
}

// SetNillableCpuCoreUsageMax sets the "cpuCoreUsageMax" field if the given value is not nil.
func (acuo *AllocationCostUpdateOne) SetNillableCpuCoreUsageMax(f *float64) *AllocationCostUpdateOne {
	if f != nil {
		acuo.SetCpuCoreUsageMax(*f)
	}
	return acuo
}

// AddCpuCoreUsageMax adds f to the "cpuCoreUsageMax" field.
func (acuo *AllocationCostUpdateOne) AddCpuCoreUsageMax(f float64) *AllocationCostUpdateOne {
	acuo.mutation.AddCpuCoreUsageMax(f)
	return acuo
}

// SetRamByteUsageAverage sets the "ramByteUsageAverage" field.
func (acuo *AllocationCostUpdateOne) SetRamByteUsageAverage(f float64) *AllocationCostUpdateOne {
	acuo.mutation.ResetRamByteUsageAverage()
	acuo.mutation.SetRamByteUsageAverage(f)
	return acuo
}

// SetNillableRamByteUsageAverage sets the "ramByteUsageAverage" field if the given value is not nil.
func (acuo *AllocationCostUpdateOne) SetNillableRamByteUsageAverage(f *float64) *AllocationCostUpdateOne {
	if f != nil {
		acuo.SetRamByteUsageAverage(*f)
	}
	return acuo
}

// AddRamByteUsageAverage adds f to the "ramByteUsageAverage" field.
func (acuo *AllocationCostUpdateOne) AddRamByteUsageAverage(f float64) *AllocationCostUpdateOne {
	acuo.mutation.AddRamByteUsageAverage(f)
	return acuo
}

// SetRamByteUsageMax sets the "ramByteUsageMax" field.
func (acuo *AllocationCostUpdateOne) SetRamByteUsageMax(f float64) *AllocationCostUpdateOne {
	acuo.mutation.ResetRamByteUsageMax()
	acuo.mutation.SetRamByteUsageMax(f)
	return acuo
}

// SetNillableRamByteUsageMax sets the "ramByteUsageMax" field if the given value is not nil.
func (acuo *AllocationCostUpdateOne) SetNillableRamByteUsageMax(f *float64) *AllocationCostUpdateOne {
	if f != nil {
		acuo.SetRamByteUsageMax(*f)
	}
	return acuo
}

// AddRamByteUsageMax adds f to the "ramByteUsageMax" field.
func (acuo *AllocationCostUpdateOne) AddRamByteUsageMax(f float64) *AllocationCostUpdateOne {
	acuo.mutation.AddRamByteUsageMax(f)
	return acuo
}

// Mutation returns the AllocationCostMutation object of the builder.
func (acuo *AllocationCostUpdateOne) Mutation() *AllocationCostMutation {
	return acuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (acuo *AllocationCostUpdateOne) Select(field string, fields ...string) *AllocationCostUpdateOne {
	acuo.fields = append([]string{field}, fields...)
	return acuo
}

// Save executes the query and returns the updated AllocationCost entity.
func (acuo *AllocationCostUpdateOne) Save(ctx context.Context) (*AllocationCost, error) {
	return withHooks[*AllocationCost, AllocationCostMutation](ctx, acuo.sqlSave, acuo.mutation, acuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (acuo *AllocationCostUpdateOne) SaveX(ctx context.Context) *AllocationCost {
	node, err := acuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (acuo *AllocationCostUpdateOne) Exec(ctx context.Context) error {
	_, err := acuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acuo *AllocationCostUpdateOne) ExecX(ctx context.Context) {
	if err := acuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (acuo *AllocationCostUpdateOne) check() error {
	if v, ok := acuo.mutation.TotalCost(); ok {
		if err := allocationcost.TotalCostValidator(v); err != nil {
			return &ValidationError{Name: "totalCost", err: fmt.Errorf(`model: validator failed for field "AllocationCost.totalCost": %w`, err)}
		}
	}
	if v, ok := acuo.mutation.CpuCost(); ok {
		if err := allocationcost.CpuCostValidator(v); err != nil {
			return &ValidationError{Name: "cpuCost", err: fmt.Errorf(`model: validator failed for field "AllocationCost.cpuCost": %w`, err)}
		}
	}
	if v, ok := acuo.mutation.GpuCost(); ok {
		if err := allocationcost.GpuCostValidator(v); err != nil {
			return &ValidationError{Name: "gpuCost", err: fmt.Errorf(`model: validator failed for field "AllocationCost.gpuCost": %w`, err)}
		}
	}
	if v, ok := acuo.mutation.RamCost(); ok {
		if err := allocationcost.RamCostValidator(v); err != nil {
			return &ValidationError{Name: "ramCost", err: fmt.Errorf(`model: validator failed for field "AllocationCost.ramCost": %w`, err)}
		}
	}
	if v, ok := acuo.mutation.PvCost(); ok {
		if err := allocationcost.PvCostValidator(v); err != nil {
			return &ValidationError{Name: "pvCost", err: fmt.Errorf(`model: validator failed for field "AllocationCost.pvCost": %w`, err)}
		}
	}
	if v, ok := acuo.mutation.PvBytes(); ok {
		if err := allocationcost.PvBytesValidator(v); err != nil {
			return &ValidationError{Name: "pvBytes", err: fmt.Errorf(`model: validator failed for field "AllocationCost.pvBytes": %w`, err)}
		}
	}
	if v, ok := acuo.mutation.CpuCoreUsageAverage(); ok {
		if err := allocationcost.CpuCoreUsageAverageValidator(v); err != nil {
			return &ValidationError{Name: "cpuCoreUsageAverage", err: fmt.Errorf(`model: validator failed for field "AllocationCost.cpuCoreUsageAverage": %w`, err)}
		}
	}
	if v, ok := acuo.mutation.CpuCoreUsageMax(); ok {
		if err := allocationcost.CpuCoreUsageMaxValidator(v); err != nil {
			return &ValidationError{Name: "cpuCoreUsageMax", err: fmt.Errorf(`model: validator failed for field "AllocationCost.cpuCoreUsageMax": %w`, err)}
		}
	}
	if v, ok := acuo.mutation.RamByteUsageAverage(); ok {
		if err := allocationcost.RamByteUsageAverageValidator(v); err != nil {
			return &ValidationError{Name: "ramByteUsageAverage", err: fmt.Errorf(`model: validator failed for field "AllocationCost.ramByteUsageAverage": %w`, err)}
		}
	}
	if v, ok := acuo.mutation.RamByteUsageMax(); ok {
		if err := allocationcost.RamByteUsageMaxValidator(v); err != nil {
			return &ValidationError{Name: "ramByteUsageMax", err: fmt.Errorf(`model: validator failed for field "AllocationCost.ramByteUsageMax": %w`, err)}
		}
	}
	if _, ok := acuo.mutation.ConnectorID(); acuo.mutation.ConnectorCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "AllocationCost.connector"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (acuo *AllocationCostUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *AllocationCostUpdateOne {
	acuo.modifiers = append(acuo.modifiers, modifiers...)
	return acuo
}

func (acuo *AllocationCostUpdateOne) sqlSave(ctx context.Context) (_node *AllocationCost, err error) {
	if err := acuo.check(); err != nil {
		return _node, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   allocationcost.Table,
			Columns: allocationcost.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: allocationcost.FieldID,
			},
		},
	}
	id, ok := acuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`model: missing "AllocationCost.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := acuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, allocationcost.FieldID)
		for _, f := range fields {
			if !allocationcost.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("model: invalid field %q for query", f)}
			}
			if f != allocationcost.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := acuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if acuo.mutation.ClusterNameCleared() {
		_spec.ClearField(allocationcost.FieldClusterName, field.TypeString)
	}
	if acuo.mutation.NamespaceCleared() {
		_spec.ClearField(allocationcost.FieldNamespace, field.TypeString)
	}
	if acuo.mutation.NodeCleared() {
		_spec.ClearField(allocationcost.FieldNode, field.TypeString)
	}
	if acuo.mutation.ControllerCleared() {
		_spec.ClearField(allocationcost.FieldController, field.TypeString)
	}
	if acuo.mutation.ControllerKindCleared() {
		_spec.ClearField(allocationcost.FieldControllerKind, field.TypeString)
	}
	if acuo.mutation.PodCleared() {
		_spec.ClearField(allocationcost.FieldPod, field.TypeString)
	}
	if acuo.mutation.ContainerCleared() {
		_spec.ClearField(allocationcost.FieldContainer, field.TypeString)
	}
	if value, ok := acuo.mutation.TotalCost(); ok {
		_spec.SetField(allocationcost.FieldTotalCost, field.TypeFloat64, value)
	}
	if value, ok := acuo.mutation.AddedTotalCost(); ok {
		_spec.AddField(allocationcost.FieldTotalCost, field.TypeFloat64, value)
	}
	if value, ok := acuo.mutation.Currency(); ok {
		_spec.SetField(allocationcost.FieldCurrency, field.TypeInt, value)
	}
	if value, ok := acuo.mutation.AddedCurrency(); ok {
		_spec.AddField(allocationcost.FieldCurrency, field.TypeInt, value)
	}
	if acuo.mutation.CurrencyCleared() {
		_spec.ClearField(allocationcost.FieldCurrency, field.TypeInt)
	}
	if value, ok := acuo.mutation.CpuCost(); ok {
		_spec.SetField(allocationcost.FieldCpuCost, field.TypeFloat64, value)
	}
	if value, ok := acuo.mutation.AddedCpuCost(); ok {
		_spec.AddField(allocationcost.FieldCpuCost, field.TypeFloat64, value)
	}
	if value, ok := acuo.mutation.GpuCost(); ok {
		_spec.SetField(allocationcost.FieldGpuCost, field.TypeFloat64, value)
	}
	if value, ok := acuo.mutation.AddedGpuCost(); ok {
		_spec.AddField(allocationcost.FieldGpuCost, field.TypeFloat64, value)
	}
	if value, ok := acuo.mutation.RamCost(); ok {
		_spec.SetField(allocationcost.FieldRamCost, field.TypeFloat64, value)
	}
	if value, ok := acuo.mutation.AddedRamCost(); ok {
		_spec.AddField(allocationcost.FieldRamCost, field.TypeFloat64, value)
	}
	if value, ok := acuo.mutation.PvCost(); ok {
		_spec.SetField(allocationcost.FieldPvCost, field.TypeFloat64, value)
	}
	if value, ok := acuo.mutation.AddedPvCost(); ok {
		_spec.AddField(allocationcost.FieldPvCost, field.TypeFloat64, value)
	}
	if value, ok := acuo.mutation.PvBytes(); ok {
		_spec.SetField(allocationcost.FieldPvBytes, field.TypeFloat64, value)
	}
	if value, ok := acuo.mutation.AddedPvBytes(); ok {
		_spec.AddField(allocationcost.FieldPvBytes, field.TypeFloat64, value)
	}
	if value, ok := acuo.mutation.CpuCoreUsageAverage(); ok {
		_spec.SetField(allocationcost.FieldCpuCoreUsageAverage, field.TypeFloat64, value)
	}
	if value, ok := acuo.mutation.AddedCpuCoreUsageAverage(); ok {
		_spec.AddField(allocationcost.FieldCpuCoreUsageAverage, field.TypeFloat64, value)
	}
	if value, ok := acuo.mutation.CpuCoreUsageMax(); ok {
		_spec.SetField(allocationcost.FieldCpuCoreUsageMax, field.TypeFloat64, value)
	}
	if value, ok := acuo.mutation.AddedCpuCoreUsageMax(); ok {
		_spec.AddField(allocationcost.FieldCpuCoreUsageMax, field.TypeFloat64, value)
	}
	if value, ok := acuo.mutation.RamByteUsageAverage(); ok {
		_spec.SetField(allocationcost.FieldRamByteUsageAverage, field.TypeFloat64, value)
	}
	if value, ok := acuo.mutation.AddedRamByteUsageAverage(); ok {
		_spec.AddField(allocationcost.FieldRamByteUsageAverage, field.TypeFloat64, value)
	}
	if value, ok := acuo.mutation.RamByteUsageMax(); ok {
		_spec.SetField(allocationcost.FieldRamByteUsageMax, field.TypeFloat64, value)
	}
	if value, ok := acuo.mutation.AddedRamByteUsageMax(); ok {
		_spec.AddField(allocationcost.FieldRamByteUsageMax, field.TypeFloat64, value)
	}
	_spec.Node.Schema = acuo.schemaConfig.AllocationCost
	ctx = internal.NewSchemaConfigContext(ctx, acuo.schemaConfig)
	_spec.AddModifiers(acuo.modifiers...)
	_node = &AllocationCost{config: acuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, acuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{allocationcost.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	acuo.mutation.done = true
	return _node, nil
}
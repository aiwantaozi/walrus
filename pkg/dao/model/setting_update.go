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
	"github.com/seal-io/walrus/pkg/dao/model/setting"
	"github.com/seal-io/walrus/pkg/dao/types/crypto"
)

// SettingUpdate is the builder for updating Setting entities.
type SettingUpdate struct {
	config
	hooks     []Hook
	mutation  *SettingMutation
	modifiers []func(*sql.UpdateBuilder)
	object    *Setting
}

// Where appends a list predicates to the SettingUpdate builder.
func (su *SettingUpdate) Where(ps ...predicate.Setting) *SettingUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetUpdateTime sets the "update_time" field.
func (su *SettingUpdate) SetUpdateTime(t time.Time) *SettingUpdate {
	su.mutation.SetUpdateTime(t)
	return su
}

// SetValue sets the "value" field.
func (su *SettingUpdate) SetValue(c crypto.String) *SettingUpdate {
	su.mutation.SetValue(c)
	return su
}

// SetHidden sets the "hidden" field.
func (su *SettingUpdate) SetHidden(b bool) *SettingUpdate {
	su.mutation.SetHidden(b)
	return su
}

// SetNillableHidden sets the "hidden" field if the given value is not nil.
func (su *SettingUpdate) SetNillableHidden(b *bool) *SettingUpdate {
	if b != nil {
		su.SetHidden(*b)
	}
	return su
}

// ClearHidden clears the value of the "hidden" field.
func (su *SettingUpdate) ClearHidden() *SettingUpdate {
	su.mutation.ClearHidden()
	return su
}

// SetEditable sets the "editable" field.
func (su *SettingUpdate) SetEditable(b bool) *SettingUpdate {
	su.mutation.SetEditable(b)
	return su
}

// SetNillableEditable sets the "editable" field if the given value is not nil.
func (su *SettingUpdate) SetNillableEditable(b *bool) *SettingUpdate {
	if b != nil {
		su.SetEditable(*b)
	}
	return su
}

// ClearEditable clears the value of the "editable" field.
func (su *SettingUpdate) ClearEditable() *SettingUpdate {
	su.mutation.ClearEditable()
	return su
}

// SetSensitive sets the "sensitive" field.
func (su *SettingUpdate) SetSensitive(b bool) *SettingUpdate {
	su.mutation.SetSensitive(b)
	return su
}

// SetNillableSensitive sets the "sensitive" field if the given value is not nil.
func (su *SettingUpdate) SetNillableSensitive(b *bool) *SettingUpdate {
	if b != nil {
		su.SetSensitive(*b)
	}
	return su
}

// ClearSensitive clears the value of the "sensitive" field.
func (su *SettingUpdate) ClearSensitive() *SettingUpdate {
	su.mutation.ClearSensitive()
	return su
}

// SetPrivate sets the "private" field.
func (su *SettingUpdate) SetPrivate(b bool) *SettingUpdate {
	su.mutation.SetPrivate(b)
	return su
}

// SetNillablePrivate sets the "private" field if the given value is not nil.
func (su *SettingUpdate) SetNillablePrivate(b *bool) *SettingUpdate {
	if b != nil {
		su.SetPrivate(*b)
	}
	return su
}

// ClearPrivate clears the value of the "private" field.
func (su *SettingUpdate) ClearPrivate() *SettingUpdate {
	su.mutation.ClearPrivate()
	return su
}

// Mutation returns the SettingMutation object of the builder.
func (su *SettingUpdate) Mutation() *SettingMutation {
	return su.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SettingUpdate) Save(ctx context.Context) (int, error) {
	if err := su.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *SettingUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SettingUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SettingUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (su *SettingUpdate) defaults() error {
	if _, ok := su.mutation.UpdateTime(); !ok {
		if setting.UpdateDefaultUpdateTime == nil {
			return fmt.Errorf("model: uninitialized setting.UpdateDefaultUpdateTime (forgotten import model/runtime?)")
		}
		v := setting.UpdateDefaultUpdateTime()
		su.mutation.SetUpdateTime(v)
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
func (su *SettingUpdate) Set(obj *Setting) *SettingUpdate {
	// Without Default.
	su.SetValue(obj.Value)
	if obj.Hidden != nil {
		su.SetHidden(*obj.Hidden)
	}
	if obj.Editable != nil {
		su.SetEditable(*obj.Editable)
	}
	if obj.Sensitive != nil {
		su.SetSensitive(*obj.Sensitive)
	}
	if obj.Private != nil {
		su.SetPrivate(*obj.Private)
	}

	// With Default.
	if obj.UpdateTime != nil {
		su.SetUpdateTime(*obj.UpdateTime)
	}

	// Record the given object.
	su.object = obj

	return su
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (su *SettingUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *SettingUpdate {
	su.modifiers = append(su.modifiers, modifiers...)
	return su
}

func (su *SettingUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(setting.Table, setting.Columns, sqlgraph.NewFieldSpec(setting.FieldID, field.TypeString))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.UpdateTime(); ok {
		_spec.SetField(setting.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := su.mutation.Value(); ok {
		_spec.SetField(setting.FieldValue, field.TypeString, value)
	}
	if value, ok := su.mutation.Hidden(); ok {
		_spec.SetField(setting.FieldHidden, field.TypeBool, value)
	}
	if su.mutation.HiddenCleared() {
		_spec.ClearField(setting.FieldHidden, field.TypeBool)
	}
	if value, ok := su.mutation.Editable(); ok {
		_spec.SetField(setting.FieldEditable, field.TypeBool, value)
	}
	if su.mutation.EditableCleared() {
		_spec.ClearField(setting.FieldEditable, field.TypeBool)
	}
	if value, ok := su.mutation.Sensitive(); ok {
		_spec.SetField(setting.FieldSensitive, field.TypeBool, value)
	}
	if su.mutation.SensitiveCleared() {
		_spec.ClearField(setting.FieldSensitive, field.TypeBool)
	}
	if value, ok := su.mutation.Private(); ok {
		_spec.SetField(setting.FieldPrivate, field.TypeBool, value)
	}
	if su.mutation.PrivateCleared() {
		_spec.ClearField(setting.FieldPrivate, field.TypeBool)
	}
	_spec.Node.Schema = su.schemaConfig.Setting
	ctx = internal.NewSchemaConfigContext(ctx, su.schemaConfig)
	_spec.AddModifiers(su.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{setting.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// SettingUpdateOne is the builder for updating a single Setting entity.
type SettingUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *SettingMutation
	modifiers []func(*sql.UpdateBuilder)
	object    *Setting
}

// SetUpdateTime sets the "update_time" field.
func (suo *SettingUpdateOne) SetUpdateTime(t time.Time) *SettingUpdateOne {
	suo.mutation.SetUpdateTime(t)
	return suo
}

// SetValue sets the "value" field.
func (suo *SettingUpdateOne) SetValue(c crypto.String) *SettingUpdateOne {
	suo.mutation.SetValue(c)
	return suo
}

// SetHidden sets the "hidden" field.
func (suo *SettingUpdateOne) SetHidden(b bool) *SettingUpdateOne {
	suo.mutation.SetHidden(b)
	return suo
}

// SetNillableHidden sets the "hidden" field if the given value is not nil.
func (suo *SettingUpdateOne) SetNillableHidden(b *bool) *SettingUpdateOne {
	if b != nil {
		suo.SetHidden(*b)
	}
	return suo
}

// ClearHidden clears the value of the "hidden" field.
func (suo *SettingUpdateOne) ClearHidden() *SettingUpdateOne {
	suo.mutation.ClearHidden()
	return suo
}

// SetEditable sets the "editable" field.
func (suo *SettingUpdateOne) SetEditable(b bool) *SettingUpdateOne {
	suo.mutation.SetEditable(b)
	return suo
}

// SetNillableEditable sets the "editable" field if the given value is not nil.
func (suo *SettingUpdateOne) SetNillableEditable(b *bool) *SettingUpdateOne {
	if b != nil {
		suo.SetEditable(*b)
	}
	return suo
}

// ClearEditable clears the value of the "editable" field.
func (suo *SettingUpdateOne) ClearEditable() *SettingUpdateOne {
	suo.mutation.ClearEditable()
	return suo
}

// SetSensitive sets the "sensitive" field.
func (suo *SettingUpdateOne) SetSensitive(b bool) *SettingUpdateOne {
	suo.mutation.SetSensitive(b)
	return suo
}

// SetNillableSensitive sets the "sensitive" field if the given value is not nil.
func (suo *SettingUpdateOne) SetNillableSensitive(b *bool) *SettingUpdateOne {
	if b != nil {
		suo.SetSensitive(*b)
	}
	return suo
}

// ClearSensitive clears the value of the "sensitive" field.
func (suo *SettingUpdateOne) ClearSensitive() *SettingUpdateOne {
	suo.mutation.ClearSensitive()
	return suo
}

// SetPrivate sets the "private" field.
func (suo *SettingUpdateOne) SetPrivate(b bool) *SettingUpdateOne {
	suo.mutation.SetPrivate(b)
	return suo
}

// SetNillablePrivate sets the "private" field if the given value is not nil.
func (suo *SettingUpdateOne) SetNillablePrivate(b *bool) *SettingUpdateOne {
	if b != nil {
		suo.SetPrivate(*b)
	}
	return suo
}

// ClearPrivate clears the value of the "private" field.
func (suo *SettingUpdateOne) ClearPrivate() *SettingUpdateOne {
	suo.mutation.ClearPrivate()
	return suo
}

// Mutation returns the SettingMutation object of the builder.
func (suo *SettingUpdateOne) Mutation() *SettingMutation {
	return suo.mutation
}

// Where appends a list predicates to the SettingUpdate builder.
func (suo *SettingUpdateOne) Where(ps ...predicate.Setting) *SettingUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SettingUpdateOne) Select(field string, fields ...string) *SettingUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Setting entity.
func (suo *SettingUpdateOne) Save(ctx context.Context) (*Setting, error) {
	if err := suo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SettingUpdateOne) SaveX(ctx context.Context) *Setting {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SettingUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SettingUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (suo *SettingUpdateOne) defaults() error {
	if _, ok := suo.mutation.UpdateTime(); !ok {
		if setting.UpdateDefaultUpdateTime == nil {
			return fmt.Errorf("model: uninitialized setting.UpdateDefaultUpdateTime (forgotten import model/runtime?)")
		}
		v := setting.UpdateDefaultUpdateTime()
		suo.mutation.SetUpdateTime(v)
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
func (suo *SettingUpdateOne) Set(obj *Setting) *SettingUpdateOne {
	h := func(n ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			mt := m.(*SettingMutation)
			db, err := mt.Client().Setting.Get(ctx, *mt.id)
			if err != nil {
				return nil, fmt.Errorf("failed getting Setting with id: %v", *mt.id)
			}

			// Without Default.
			if db.Value != obj.Value {
				suo.SetValue(obj.Value)
			}
			if obj.Hidden != nil {
				if !reflect.DeepEqual(db.Hidden, obj.Hidden) {
					suo.SetHidden(*obj.Hidden)
				}
			}
			if obj.Editable != nil {
				if !reflect.DeepEqual(db.Editable, obj.Editable) {
					suo.SetEditable(*obj.Editable)
				}
			}
			if obj.Sensitive != nil {
				if !reflect.DeepEqual(db.Sensitive, obj.Sensitive) {
					suo.SetSensitive(*obj.Sensitive)
				}
			}
			if obj.Private != nil {
				if !reflect.DeepEqual(db.Private, obj.Private) {
					suo.SetPrivate(*obj.Private)
				}
			}

			// With Default.
			if (obj.UpdateTime != nil) && (!reflect.DeepEqual(db.UpdateTime, obj.UpdateTime)) {
				suo.SetUpdateTime(*obj.UpdateTime)
			}

			// Record the given object.
			suo.object = obj

			return n.Mutate(ctx, m)
		})
	}

	suo.hooks = append(suo.hooks, h)

	return suo
}

// getClientSet returns the ClientSet for the given builder.
func (suo *SettingUpdateOne) getClientSet() (mc ClientSet) {
	if _, ok := suo.config.driver.(*txDriver); ok {
		tx := &Tx{config: suo.config}
		tx.init()
		mc = tx
	} else {
		cli := &Client{config: suo.config}
		cli.init()
		mc = cli
	}
	return mc
}

// SaveE calls the given function after updated the Setting entity,
// which is always good for cascading update operations.
func (suo *SettingUpdateOne) SaveE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *Setting) error) (*Setting, error) {
	obj, err := suo.Save(ctx)
	if err != nil &&
		(suo.object == nil || !errors.Is(err, stdsql.ErrNoRows)) {
		return nil, err
	}

	if len(cbs) == 0 {
		return obj, err
	}

	mc := suo.getClientSet()

	if obj == nil {
		obj = suo.object
	} else if x := suo.object; x != nil {
		if _, set := suo.mutation.Field(setting.FieldValue); set {
			obj.Value = x.Value
		}
		if _, set := suo.mutation.Field(setting.FieldHidden); set {
			obj.Hidden = x.Hidden
		}
		if _, set := suo.mutation.Field(setting.FieldEditable); set {
			obj.Editable = x.Editable
		}
		if _, set := suo.mutation.Field(setting.FieldSensitive); set {
			obj.Sensitive = x.Sensitive
		}
		if _, set := suo.mutation.Field(setting.FieldPrivate); set {
			obj.Private = x.Private
		}
	}

	for i := range cbs {
		if err = cbs[i](ctx, mc, obj); err != nil {
			return nil, err
		}
	}

	return obj, nil
}

// SaveEX is like SaveE, but panics if an error occurs.
func (suo *SettingUpdateOne) SaveEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *Setting) error) *Setting {
	obj, err := suo.SaveE(ctx, cbs...)
	if err != nil {
		panic(err)
	}
	return obj
}

// ExecE calls the given function after executed the query,
// which is always good for cascading update operations.
func (suo *SettingUpdateOne) ExecE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *Setting) error) error {
	_, err := suo.SaveE(ctx, cbs...)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SettingUpdateOne) ExecEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *Setting) error) {
	if err := suo.ExecE(ctx, cbs...); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (suo *SettingUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *SettingUpdateOne {
	suo.modifiers = append(suo.modifiers, modifiers...)
	return suo
}

func (suo *SettingUpdateOne) sqlSave(ctx context.Context) (_node *Setting, err error) {
	_spec := sqlgraph.NewUpdateSpec(setting.Table, setting.Columns, sqlgraph.NewFieldSpec(setting.FieldID, field.TypeString))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`model: missing "Setting.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, setting.FieldID)
		for _, f := range fields {
			if !setting.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("model: invalid field %q for query", f)}
			}
			if f != setting.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.UpdateTime(); ok {
		_spec.SetField(setting.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := suo.mutation.Value(); ok {
		_spec.SetField(setting.FieldValue, field.TypeString, value)
	}
	if value, ok := suo.mutation.Hidden(); ok {
		_spec.SetField(setting.FieldHidden, field.TypeBool, value)
	}
	if suo.mutation.HiddenCleared() {
		_spec.ClearField(setting.FieldHidden, field.TypeBool)
	}
	if value, ok := suo.mutation.Editable(); ok {
		_spec.SetField(setting.FieldEditable, field.TypeBool, value)
	}
	if suo.mutation.EditableCleared() {
		_spec.ClearField(setting.FieldEditable, field.TypeBool)
	}
	if value, ok := suo.mutation.Sensitive(); ok {
		_spec.SetField(setting.FieldSensitive, field.TypeBool, value)
	}
	if suo.mutation.SensitiveCleared() {
		_spec.ClearField(setting.FieldSensitive, field.TypeBool)
	}
	if value, ok := suo.mutation.Private(); ok {
		_spec.SetField(setting.FieldPrivate, field.TypeBool, value)
	}
	if suo.mutation.PrivateCleared() {
		_spec.ClearField(setting.FieldPrivate, field.TypeBool)
	}
	_spec.Node.Schema = suo.schemaConfig.Setting
	ctx = internal.NewSchemaConfigContext(ctx, suo.schemaConfig)
	_spec.AddModifiers(suo.modifiers...)
	_node = &Setting{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{setting.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}

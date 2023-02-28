// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// GENERATED, DO NOT EDIT.

package perspective

import (
	"time"

	"entgo.io/ent"

	"github.com/seal-io/seal/pkg/dao/types"
)

const (
	// Label holds the string label denoting the perspective type in the database.
	Label = "perspective"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the createtime field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the updatetime field in the database.
	FieldUpdateTime = "update_time"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldStartTime holds the string denoting the starttime field in the database.
	FieldStartTime = "start_time"
	// FieldEndTime holds the string denoting the endtime field in the database.
	FieldEndTime = "end_time"
	// FieldBuiltin holds the string denoting the builtin field in the database.
	FieldBuiltin = "builtin"
	// FieldAllocationQueries holds the string denoting the allocationqueries field in the database.
	FieldAllocationQueries = "allocation_queries"
	// Table holds the table name of the perspective in the database.
	Table = "perspectives"
)

// Columns holds all SQL columns for perspective fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldName,
	FieldStartTime,
	FieldEndTime,
	FieldBuiltin,
	FieldAllocationQueries,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/seal-io/seal/pkg/dao/model/runtime"
var (
	Hooks [1]ent.Hook
	// DefaultCreateTime holds the default value on creation for the "createTime" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "updateTime" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "updateTime" field.
	UpdateDefaultUpdateTime func() time.Time
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// StartTimeValidator is a validator for the "startTime" field. It is called by the builders before save.
	StartTimeValidator func(string) error
	// EndTimeValidator is a validator for the "endTime" field. It is called by the builders before save.
	EndTimeValidator func(string) error
	// DefaultBuiltin holds the default value on creation for the "builtin" field.
	DefaultBuiltin bool
	// DefaultAllocationQueries holds the default value on creation for the "allocationQueries" field.
	DefaultAllocationQueries []types.QueryCondition
)

// WithoutFields returns the fields ignored the given list.
func WithoutFields(ignores ...string) []string {
	if len(ignores) == 0 {
		return Columns
	}

	var s = make(map[string]bool, len(ignores))
	for i := range ignores {
		s[ignores[i]] = true
	}

	var r = make([]string, 0, len(Columns)-len(s))
	for i := range Columns {
		if s[Columns[i]] {
			continue
		}
		r = append(r, Columns[i])
	}
	return r
}
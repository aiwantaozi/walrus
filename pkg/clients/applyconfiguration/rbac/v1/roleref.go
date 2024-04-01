// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

package v1

// RoleRefApplyConfiguration represents an declarative configuration of the RoleRef type for use
// with apply.
type RoleRefApplyConfiguration struct {
	APIGroup *string `json:"apiGroup,omitempty"`
	Kind     *string `json:"kind,omitempty"`
	Name     *string `json:"name,omitempty"`
}

// RoleRefApplyConfiguration constructs an declarative configuration of the RoleRef type for use with
// apply.
func RoleRef() *RoleRefApplyConfiguration {
	return &RoleRefApplyConfiguration{}
}

// WithAPIGroup sets the APIGroup field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIGroup field is set to the value of the last call.
func (b *RoleRefApplyConfiguration) WithAPIGroup(value string) *RoleRefApplyConfiguration {
	b.APIGroup = &value
	return b
}

// WithKind sets the Kind field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Kind field is set to the value of the last call.
func (b *RoleRefApplyConfiguration) WithKind(value string) *RoleRefApplyConfiguration {
	b.Kind = &value
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *RoleRefApplyConfiguration) WithName(value string) *RoleRefApplyConfiguration {
	b.Name = &value
	return b
}
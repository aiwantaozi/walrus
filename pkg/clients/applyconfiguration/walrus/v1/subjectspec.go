// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

package v1

import (
	v1 "github.com/seal-io/walrus/pkg/apis/walrus/v1"
)

// SubjectSpecApplyConfiguration represents an declarative configuration of the SubjectSpec type for use
// with apply.
type SubjectSpecApplyConfiguration struct {
	Provider    *string         `json:"provider,omitempty"`
	Role        *v1.SubjectRole `json:"role,omitempty"`
	DisplayName *string         `json:"displayName,omitempty"`
	Description *string         `json:"description,omitempty"`
	Email       *string         `json:"email,omitempty"`
	Groups      []string        `json:"groups,omitempty"`
	Credential  *string         `json:"credential,omitempty"`
}

// SubjectSpecApplyConfiguration constructs an declarative configuration of the SubjectSpec type for use with
// apply.
func SubjectSpec() *SubjectSpecApplyConfiguration {
	return &SubjectSpecApplyConfiguration{}
}

// WithProvider sets the Provider field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Provider field is set to the value of the last call.
func (b *SubjectSpecApplyConfiguration) WithProvider(value string) *SubjectSpecApplyConfiguration {
	b.Provider = &value
	return b
}

// WithRole sets the Role field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Role field is set to the value of the last call.
func (b *SubjectSpecApplyConfiguration) WithRole(value v1.SubjectRole) *SubjectSpecApplyConfiguration {
	b.Role = &value
	return b
}

// WithDisplayName sets the DisplayName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DisplayName field is set to the value of the last call.
func (b *SubjectSpecApplyConfiguration) WithDisplayName(value string) *SubjectSpecApplyConfiguration {
	b.DisplayName = &value
	return b
}

// WithDescription sets the Description field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Description field is set to the value of the last call.
func (b *SubjectSpecApplyConfiguration) WithDescription(value string) *SubjectSpecApplyConfiguration {
	b.Description = &value
	return b
}

// WithEmail sets the Email field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Email field is set to the value of the last call.
func (b *SubjectSpecApplyConfiguration) WithEmail(value string) *SubjectSpecApplyConfiguration {
	b.Email = &value
	return b
}

// WithGroups adds the given value to the Groups field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Groups field.
func (b *SubjectSpecApplyConfiguration) WithGroups(values ...string) *SubjectSpecApplyConfiguration {
	for i := range values {
		b.Groups = append(b.Groups, values[i])
	}
	return b
}

// WithCredential sets the Credential field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Credential field is set to the value of the last call.
func (b *SubjectSpecApplyConfiguration) WithCredential(value string) *SubjectSpecApplyConfiguration {
	b.Credential = &value
	return b
}
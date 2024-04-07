// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

package v1

// SubjectProviderLdapUserSearchApplyConfiguration represents an declarative configuration of the SubjectProviderLdapUserSearch type for use
// with apply.
type SubjectProviderLdapUserSearchApplyConfiguration struct {
	BaseDN               *string `json:"baseDN,omitempty"`
	Filter               *string `json:"filter,omitempty"`
	NameAttribute        *string `json:"nameAttribute,omitempty"`
	DisplayNameAttribute *string `json:"displayNameAttribute,omitempty"`
	EmailAttribute       *string `json:"emailAttribute,omitempty"`
}

// SubjectProviderLdapUserSearchApplyConfiguration constructs an declarative configuration of the SubjectProviderLdapUserSearch type for use with
// apply.
func SubjectProviderLdapUserSearch() *SubjectProviderLdapUserSearchApplyConfiguration {
	return &SubjectProviderLdapUserSearchApplyConfiguration{}
}

// WithBaseDN sets the BaseDN field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the BaseDN field is set to the value of the last call.
func (b *SubjectProviderLdapUserSearchApplyConfiguration) WithBaseDN(value string) *SubjectProviderLdapUserSearchApplyConfiguration {
	b.BaseDN = &value
	return b
}

// WithFilter sets the Filter field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Filter field is set to the value of the last call.
func (b *SubjectProviderLdapUserSearchApplyConfiguration) WithFilter(value string) *SubjectProviderLdapUserSearchApplyConfiguration {
	b.Filter = &value
	return b
}

// WithNameAttribute sets the NameAttribute field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NameAttribute field is set to the value of the last call.
func (b *SubjectProviderLdapUserSearchApplyConfiguration) WithNameAttribute(value string) *SubjectProviderLdapUserSearchApplyConfiguration {
	b.NameAttribute = &value
	return b
}

// WithDisplayNameAttribute sets the DisplayNameAttribute field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DisplayNameAttribute field is set to the value of the last call.
func (b *SubjectProviderLdapUserSearchApplyConfiguration) WithDisplayNameAttribute(value string) *SubjectProviderLdapUserSearchApplyConfiguration {
	b.DisplayNameAttribute = &value
	return b
}

// WithEmailAttribute sets the EmailAttribute field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the EmailAttribute field is set to the value of the last call.
func (b *SubjectProviderLdapUserSearchApplyConfiguration) WithEmailAttribute(value string) *SubjectProviderLdapUserSearchApplyConfiguration {
	b.EmailAttribute = &value
	return b
}
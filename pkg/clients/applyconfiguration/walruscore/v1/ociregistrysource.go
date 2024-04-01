// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

package v1

// OCIRegistrySourceApplyConfiguration represents an declarative configuration of the OCIRegistrySource type for use
// with apply.
type OCIRegistrySourceApplyConfiguration struct {
	BasicAuthApplyConfiguration `json:",inline"`
	Domain                      *string `json:"domain,omitempty"`
}

// OCIRegistrySourceApplyConfiguration constructs an declarative configuration of the OCIRegistrySource type for use with
// apply.
func OCIRegistrySource() *OCIRegistrySourceApplyConfiguration {
	return &OCIRegistrySourceApplyConfiguration{}
}

// WithUsername sets the Username field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Username field is set to the value of the last call.
func (b *OCIRegistrySourceApplyConfiguration) WithUsername(value string) *OCIRegistrySourceApplyConfiguration {
	b.Username = &value
	return b
}

// WithPassword sets the Password field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Password field is set to the value of the last call.
func (b *OCIRegistrySourceApplyConfiguration) WithPassword(value string) *OCIRegistrySourceApplyConfiguration {
	b.Password = &value
	return b
}

// WithSecretRef sets the SecretRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SecretRef field is set to the value of the last call.
func (b *OCIRegistrySourceApplyConfiguration) WithSecretRef(value *LocalObjectReferenceApplyConfiguration) *OCIRegistrySourceApplyConfiguration {
	b.SecretRef = value
	return b
}

// WithDomain sets the Domain field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Domain field is set to the value of the last call.
func (b *OCIRegistrySourceApplyConfiguration) WithDomain(value string) *OCIRegistrySourceApplyConfiguration {
	b.Domain = &value
	return b
}

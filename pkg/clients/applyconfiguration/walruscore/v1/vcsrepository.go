// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

package v1

// VCSRepositoryApplyConfiguration represents an declarative configuration of the VCSRepository type for use
// with apply.
type VCSRepositoryApplyConfiguration struct {
	TokenAuthApplyConfiguration `json:",inline"`
	URL                         *string `json:"url,omitempty"`
}

// VCSRepositoryApplyConfiguration constructs an declarative configuration of the VCSRepository type for use with
// apply.
func VCSRepository() *VCSRepositoryApplyConfiguration {
	return &VCSRepositoryApplyConfiguration{}
}

// WithToken sets the Token field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Token field is set to the value of the last call.
func (b *VCSRepositoryApplyConfiguration) WithToken(value string) *VCSRepositoryApplyConfiguration {
	b.Token = &value
	return b
}

// WithSecretRef sets the SecretRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SecretRef field is set to the value of the last call.
func (b *VCSRepositoryApplyConfiguration) WithSecretRef(value *LocalObjectReferenceApplyConfiguration) *VCSRepositoryApplyConfiguration {
	b.SecretRef = value
	return b
}

// WithURL sets the URL field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the URL field is set to the value of the last call.
func (b *VCSRepositoryApplyConfiguration) WithURL(value string) *VCSRepositoryApplyConfiguration {
	b.URL = &value
	return b
}

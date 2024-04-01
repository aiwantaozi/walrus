// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

package v1

// TokenAuthApplyConfiguration represents an declarative configuration of the TokenAuth type for use
// with apply.
type TokenAuthApplyConfiguration struct {
	Token     *string                                 `json:"token,omitempty"`
	SecretRef *TokenObjectReferenceApplyConfiguration `json:"secretRef,omitempty"`
}

// TokenAuthApplyConfiguration constructs an declarative configuration of the TokenAuth type for use with
// apply.
func TokenAuth() *TokenAuthApplyConfiguration {
	return &TokenAuthApplyConfiguration{}
}

// WithToken sets the Token field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Token field is set to the value of the last call.
func (b *TokenAuthApplyConfiguration) WithToken(value string) *TokenAuthApplyConfiguration {
	b.Token = &value
	return b
}

// WithSecretRef sets the SecretRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SecretRef field is set to the value of the last call.
func (b *TokenAuthApplyConfiguration) WithSecretRef(value *TokenObjectReferenceApplyConfiguration) *TokenAuthApplyConfiguration {
	b.SecretRef = value
	return b
}

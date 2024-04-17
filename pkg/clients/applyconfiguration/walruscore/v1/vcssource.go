// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

package v1

import (
	v1 "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
)

// VCSSourceApplyConfiguration represents an declarative configuration of the VCSSource type for use
// with apply.
type VCSSourceApplyConfiguration struct {
	Platform *v1.VCSPlatform `json:"platform,omitempty"`
	URL      *string         `json:"url,omitempty"`
}

// VCSSourceApplyConfiguration constructs an declarative configuration of the VCSSource type for use with
// apply.
func VCSSource() *VCSSourceApplyConfiguration {
	return &VCSSourceApplyConfiguration{}
}

// WithPlatform sets the Platform field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Platform field is set to the value of the last call.
func (b *VCSSourceApplyConfiguration) WithPlatform(value v1.VCSPlatform) *VCSSourceApplyConfiguration {
	b.Platform = &value
	return b
}

// WithURL sets the URL field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the URL field is set to the value of the last call.
func (b *VCSSourceApplyConfiguration) WithURL(value string) *VCSSourceApplyConfiguration {
	b.URL = &value
	return b
}
// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

package v1

// HelmChartApplyConfiguration represents an declarative configuration of the HelmChart type for use
// with apply.
type HelmChartApplyConfiguration struct {
	BasicAuthApplyConfiguration `json:",inline"`
	URL                         *string `json:"url,omitempty"`
}

// HelmChartApplyConfiguration constructs an declarative configuration of the HelmChart type for use with
// apply.
func HelmChart() *HelmChartApplyConfiguration {
	return &HelmChartApplyConfiguration{}
}

// WithUsername sets the Username field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Username field is set to the value of the last call.
func (b *HelmChartApplyConfiguration) WithUsername(value string) *HelmChartApplyConfiguration {
	b.Username = &value
	return b
}

// WithPassword sets the Password field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Password field is set to the value of the last call.
func (b *HelmChartApplyConfiguration) WithPassword(value string) *HelmChartApplyConfiguration {
	b.Password = &value
	return b
}

// WithSecretRef sets the SecretRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SecretRef field is set to the value of the last call.
func (b *HelmChartApplyConfiguration) WithSecretRef(value *LocalObjectReferenceApplyConfiguration) *HelmChartApplyConfiguration {
	b.SecretRef = value
	return b
}

// WithURL sets the URL field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the URL field is set to the value of the last call.
func (b *HelmChartApplyConfiguration) WithURL(value string) *HelmChartApplyConfiguration {
	b.URL = &value
	return b
}

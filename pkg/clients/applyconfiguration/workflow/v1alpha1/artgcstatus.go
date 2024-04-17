// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1"
)

// ArtGCStatusApplyConfiguration represents an declarative configuration of the ArtGCStatus type for use
// with apply.
type ArtGCStatusApplyConfiguration struct {
	StrategiesProcessed map[v1alpha1.ArtifactGCStrategy]bool `json:"strategiesProcessed,omitempty"`
	PodsRecouped        map[string]bool                      `json:"podsRecouped,omitempty"`
	NotSpecified        *bool                                `json:"notSpecified,omitempty"`
}

// ArtGCStatusApplyConfiguration constructs an declarative configuration of the ArtGCStatus type for use with
// apply.
func ArtGCStatus() *ArtGCStatusApplyConfiguration {
	return &ArtGCStatusApplyConfiguration{}
}

// WithStrategiesProcessed puts the entries into the StrategiesProcessed field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the StrategiesProcessed field,
// overwriting an existing map entries in StrategiesProcessed field with the same key.
func (b *ArtGCStatusApplyConfiguration) WithStrategiesProcessed(entries map[v1alpha1.ArtifactGCStrategy]bool) *ArtGCStatusApplyConfiguration {
	if b.StrategiesProcessed == nil && len(entries) > 0 {
		b.StrategiesProcessed = make(map[v1alpha1.ArtifactGCStrategy]bool, len(entries))
	}
	for k, v := range entries {
		b.StrategiesProcessed[k] = v
	}
	return b
}

// WithPodsRecouped puts the entries into the PodsRecouped field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the PodsRecouped field,
// overwriting an existing map entries in PodsRecouped field with the same key.
func (b *ArtGCStatusApplyConfiguration) WithPodsRecouped(entries map[string]bool) *ArtGCStatusApplyConfiguration {
	if b.PodsRecouped == nil && len(entries) > 0 {
		b.PodsRecouped = make(map[string]bool, len(entries))
	}
	for k, v := range entries {
		b.PodsRecouped[k] = v
	}
	return b
}

// WithNotSpecified sets the NotSpecified field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NotSpecified field is set to the value of the last call.
func (b *ArtGCStatusApplyConfiguration) WithNotSpecified(value bool) *ArtGCStatusApplyConfiguration {
	b.NotSpecified = &value
	return b
}
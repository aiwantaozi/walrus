// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1"
)

// PrometheusApplyConfiguration represents an declarative configuration of the Prometheus type for use
// with apply.
type PrometheusApplyConfiguration struct {
	Name      *string                      `json:"name,omitempty"`
	Labels    []*v1alpha1.MetricLabel      `json:"labels,omitempty"`
	Help      *string                      `json:"help,omitempty"`
	When      *string                      `json:"when,omitempty"`
	Gauge     *GaugeApplyConfiguration     `json:"gauge,omitempty"`
	Histogram *HistogramApplyConfiguration `json:"histogram,omitempty"`
	Counter   *CounterApplyConfiguration   `json:"counter,omitempty"`
}

// PrometheusApplyConfiguration constructs an declarative configuration of the Prometheus type for use with
// apply.
func Prometheus() *PrometheusApplyConfiguration {
	return &PrometheusApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *PrometheusApplyConfiguration) WithName(value string) *PrometheusApplyConfiguration {
	b.Name = &value
	return b
}

// WithLabels adds the given value to the Labels field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Labels field.
func (b *PrometheusApplyConfiguration) WithLabels(values ...**v1alpha1.MetricLabel) *PrometheusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithLabels")
		}
		b.Labels = append(b.Labels, *values[i])
	}
	return b
}

// WithHelp sets the Help field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Help field is set to the value of the last call.
func (b *PrometheusApplyConfiguration) WithHelp(value string) *PrometheusApplyConfiguration {
	b.Help = &value
	return b
}

// WithWhen sets the When field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the When field is set to the value of the last call.
func (b *PrometheusApplyConfiguration) WithWhen(value string) *PrometheusApplyConfiguration {
	b.When = &value
	return b
}

// WithGauge sets the Gauge field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Gauge field is set to the value of the last call.
func (b *PrometheusApplyConfiguration) WithGauge(value *GaugeApplyConfiguration) *PrometheusApplyConfiguration {
	b.Gauge = value
	return b
}

// WithHistogram sets the Histogram field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Histogram field is set to the value of the last call.
func (b *PrometheusApplyConfiguration) WithHistogram(value *HistogramApplyConfiguration) *PrometheusApplyConfiguration {
	b.Histogram = value
	return b
}

// WithCounter sets the Counter field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Counter field is set to the value of the last call.
func (b *PrometheusApplyConfiguration) WithCounter(value *CounterApplyConfiguration) *PrometheusApplyConfiguration {
	b.Counter = value
	return b
}
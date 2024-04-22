// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

package v1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ScheduleTaskStatusApplyConfiguration represents an declarative configuration of the ScheduleTaskStatus type for use
// with apply.
type ScheduleTaskStatusApplyConfiguration struct {
	Schedule              *string  `json:"schedule,omitempty"`
	ActiveTaskName        *string  `json:"activeTaskName,omitempty"`
	ActiveTaskTimeoutTime *v1.Time `json:"activeTaskTimeoutTime,omitempty"`
	LastScheduleTime      *v1.Time `json:"lastScheduleTime,omitempty"`
	LastSuccessfulTime    *v1.Time `json:"lastSuccessfulTime,omitempty"`
	LastFailedTime        *v1.Time `json:"lastFailedTime,omitempty"`
	LastFailedMessage     *string  `json:"lastFailedMessage,omitempty"`
}

// ScheduleTaskStatusApplyConfiguration constructs an declarative configuration of the ScheduleTaskStatus type for use with
// apply.
func ScheduleTaskStatus() *ScheduleTaskStatusApplyConfiguration {
	return &ScheduleTaskStatusApplyConfiguration{}
}

// WithSchedule sets the Schedule field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Schedule field is set to the value of the last call.
func (b *ScheduleTaskStatusApplyConfiguration) WithSchedule(value string) *ScheduleTaskStatusApplyConfiguration {
	b.Schedule = &value
	return b
}

// WithActiveTaskName sets the ActiveTaskName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ActiveTaskName field is set to the value of the last call.
func (b *ScheduleTaskStatusApplyConfiguration) WithActiveTaskName(value string) *ScheduleTaskStatusApplyConfiguration {
	b.ActiveTaskName = &value
	return b
}

// WithActiveTaskTimeoutTime sets the ActiveTaskTimeoutTime field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ActiveTaskTimeoutTime field is set to the value of the last call.
func (b *ScheduleTaskStatusApplyConfiguration) WithActiveTaskTimeoutTime(value v1.Time) *ScheduleTaskStatusApplyConfiguration {
	b.ActiveTaskTimeoutTime = &value
	return b
}

// WithLastScheduleTime sets the LastScheduleTime field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LastScheduleTime field is set to the value of the last call.
func (b *ScheduleTaskStatusApplyConfiguration) WithLastScheduleTime(value v1.Time) *ScheduleTaskStatusApplyConfiguration {
	b.LastScheduleTime = &value
	return b
}

// WithLastSuccessfulTime sets the LastSuccessfulTime field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LastSuccessfulTime field is set to the value of the last call.
func (b *ScheduleTaskStatusApplyConfiguration) WithLastSuccessfulTime(value v1.Time) *ScheduleTaskStatusApplyConfiguration {
	b.LastSuccessfulTime = &value
	return b
}

// WithLastFailedTime sets the LastFailedTime field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LastFailedTime field is set to the value of the last call.
func (b *ScheduleTaskStatusApplyConfiguration) WithLastFailedTime(value v1.Time) *ScheduleTaskStatusApplyConfiguration {
	b.LastFailedTime = &value
	return b
}

// WithLastFailedMessage sets the LastFailedMessage field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LastFailedMessage field is set to the value of the last call.
func (b *ScheduleTaskStatusApplyConfiguration) WithLastFailedMessage(value string) *ScheduleTaskStatusApplyConfiguration {
	b.LastFailedMessage = &value
	return b
}

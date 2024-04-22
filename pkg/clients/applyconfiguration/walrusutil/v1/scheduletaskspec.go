// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

package v1

// ScheduleTaskSpecApplyConfiguration represents an declarative configuration of the ScheduleTaskSpec type for use
// with apply.
type ScheduleTaskSpecApplyConfiguration struct {
	Schedule                  *string `json:"schedule,omitempty"`
	Suspend                   *bool   `json:"suspend,omitempty"`
	Await                     *bool   `json:"await,omitempty"`
	TimeZone                  *string `json:"timeZone,omitempty"`
	ActiveTaskDeadlineSeconds *int64  `json:"activeTaskDeadlineSeconds,omitempty"`
}

// ScheduleTaskSpecApplyConfiguration constructs an declarative configuration of the ScheduleTaskSpec type for use with
// apply.
func ScheduleTaskSpec() *ScheduleTaskSpecApplyConfiguration {
	return &ScheduleTaskSpecApplyConfiguration{}
}

// WithSchedule sets the Schedule field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Schedule field is set to the value of the last call.
func (b *ScheduleTaskSpecApplyConfiguration) WithSchedule(value string) *ScheduleTaskSpecApplyConfiguration {
	b.Schedule = &value
	return b
}

// WithSuspend sets the Suspend field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Suspend field is set to the value of the last call.
func (b *ScheduleTaskSpecApplyConfiguration) WithSuspend(value bool) *ScheduleTaskSpecApplyConfiguration {
	b.Suspend = &value
	return b
}

// WithAwait sets the Await field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Await field is set to the value of the last call.
func (b *ScheduleTaskSpecApplyConfiguration) WithAwait(value bool) *ScheduleTaskSpecApplyConfiguration {
	b.Await = &value
	return b
}

// WithTimeZone sets the TimeZone field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TimeZone field is set to the value of the last call.
func (b *ScheduleTaskSpecApplyConfiguration) WithTimeZone(value string) *ScheduleTaskSpecApplyConfiguration {
	b.TimeZone = &value
	return b
}

// WithActiveTaskDeadlineSeconds sets the ActiveTaskDeadlineSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ActiveTaskDeadlineSeconds field is set to the value of the last call.
func (b *ScheduleTaskSpecApplyConfiguration) WithActiveTaskDeadlineSeconds(value int64) *ScheduleTaskSpecApplyConfiguration {
	b.ActiveTaskDeadlineSeconds = &value
	return b
}

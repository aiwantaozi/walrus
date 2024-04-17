// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

package v1alpha1

// MemoizeApplyConfiguration represents an declarative configuration of the Memoize type for use
// with apply.
type MemoizeApplyConfiguration struct {
	Key    *string                  `json:"key,omitempty"`
	Cache  *CacheApplyConfiguration `json:"cache,omitempty"`
	MaxAge *string                  `json:"maxAge,omitempty"`
}

// MemoizeApplyConfiguration constructs an declarative configuration of the Memoize type for use with
// apply.
func Memoize() *MemoizeApplyConfiguration {
	return &MemoizeApplyConfiguration{}
}

// WithKey sets the Key field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Key field is set to the value of the last call.
func (b *MemoizeApplyConfiguration) WithKey(value string) *MemoizeApplyConfiguration {
	b.Key = &value
	return b
}

// WithCache sets the Cache field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Cache field is set to the value of the last call.
func (b *MemoizeApplyConfiguration) WithCache(value *CacheApplyConfiguration) *MemoizeApplyConfiguration {
	b.Cache = value
	return b
}

// WithMaxAge sets the MaxAge field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MaxAge field is set to the value of the last call.
func (b *MemoizeApplyConfiguration) WithMaxAge(value string) *MemoizeApplyConfiguration {
	b.MaxAge = &value
	return b
}
// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

package v1

// AWSElasticBlockStoreVolumeSourceApplyConfiguration represents an declarative configuration of the AWSElasticBlockStoreVolumeSource type for use
// with apply.
type AWSElasticBlockStoreVolumeSourceApplyConfiguration struct {
	VolumeID  *string `json:"volumeID,omitempty"`
	FSType    *string `json:"fsType,omitempty"`
	Partition *int32  `json:"partition,omitempty"`
	ReadOnly  *bool   `json:"readOnly,omitempty"`
}

// AWSElasticBlockStoreVolumeSourceApplyConfiguration constructs an declarative configuration of the AWSElasticBlockStoreVolumeSource type for use with
// apply.
func AWSElasticBlockStoreVolumeSource() *AWSElasticBlockStoreVolumeSourceApplyConfiguration {
	return &AWSElasticBlockStoreVolumeSourceApplyConfiguration{}
}

// WithVolumeID sets the VolumeID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the VolumeID field is set to the value of the last call.
func (b *AWSElasticBlockStoreVolumeSourceApplyConfiguration) WithVolumeID(value string) *AWSElasticBlockStoreVolumeSourceApplyConfiguration {
	b.VolumeID = &value
	return b
}

// WithFSType sets the FSType field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the FSType field is set to the value of the last call.
func (b *AWSElasticBlockStoreVolumeSourceApplyConfiguration) WithFSType(value string) *AWSElasticBlockStoreVolumeSourceApplyConfiguration {
	b.FSType = &value
	return b
}

// WithPartition sets the Partition field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Partition field is set to the value of the last call.
func (b *AWSElasticBlockStoreVolumeSourceApplyConfiguration) WithPartition(value int32) *AWSElasticBlockStoreVolumeSourceApplyConfiguration {
	b.Partition = &value
	return b
}

// WithReadOnly sets the ReadOnly field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ReadOnly field is set to the value of the last call.
func (b *AWSElasticBlockStoreVolumeSourceApplyConfiguration) WithReadOnly(value bool) *AWSElasticBlockStoreVolumeSourceApplyConfiguration {
	b.ReadOnly = &value
	return b
}
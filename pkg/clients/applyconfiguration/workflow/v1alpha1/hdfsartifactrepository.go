// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

package v1alpha1

import (
	v1 "github.com/seal-io/walrus/pkg/clients/applyconfiguration/core/v1"
)

// HDFSArtifactRepositoryApplyConfiguration represents an declarative configuration of the HDFSArtifactRepository type for use
// with apply.
type HDFSArtifactRepositoryApplyConfiguration struct {
	HDFSConfigApplyConfiguration `json:",inline"`
	PathFormat                   *string `json:"pathFormat,omitempty"`
	Force                        *bool   `json:"force,omitempty"`
}

// HDFSArtifactRepositoryApplyConfiguration constructs an declarative configuration of the HDFSArtifactRepository type for use with
// apply.
func HDFSArtifactRepository() *HDFSArtifactRepositoryApplyConfiguration {
	return &HDFSArtifactRepositoryApplyConfiguration{}
}

// WithKrbCCacheSecret sets the KrbCCacheSecret field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the KrbCCacheSecret field is set to the value of the last call.
func (b *HDFSArtifactRepositoryApplyConfiguration) WithKrbCCacheSecret(value *v1.SecretKeySelectorApplyConfiguration) *HDFSArtifactRepositoryApplyConfiguration {
	b.KrbCCacheSecret = value
	return b
}

// WithKrbKeytabSecret sets the KrbKeytabSecret field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the KrbKeytabSecret field is set to the value of the last call.
func (b *HDFSArtifactRepositoryApplyConfiguration) WithKrbKeytabSecret(value *v1.SecretKeySelectorApplyConfiguration) *HDFSArtifactRepositoryApplyConfiguration {
	b.KrbKeytabSecret = value
	return b
}

// WithKrbUsername sets the KrbUsername field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the KrbUsername field is set to the value of the last call.
func (b *HDFSArtifactRepositoryApplyConfiguration) WithKrbUsername(value string) *HDFSArtifactRepositoryApplyConfiguration {
	b.KrbUsername = &value
	return b
}

// WithKrbRealm sets the KrbRealm field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the KrbRealm field is set to the value of the last call.
func (b *HDFSArtifactRepositoryApplyConfiguration) WithKrbRealm(value string) *HDFSArtifactRepositoryApplyConfiguration {
	b.KrbRealm = &value
	return b
}

// WithKrbConfigConfigMap sets the KrbConfigConfigMap field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the KrbConfigConfigMap field is set to the value of the last call.
func (b *HDFSArtifactRepositoryApplyConfiguration) WithKrbConfigConfigMap(value *v1.ConfigMapKeySelectorApplyConfiguration) *HDFSArtifactRepositoryApplyConfiguration {
	b.KrbConfigConfigMap = value
	return b
}

// WithKrbServicePrincipalName sets the KrbServicePrincipalName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the KrbServicePrincipalName field is set to the value of the last call.
func (b *HDFSArtifactRepositoryApplyConfiguration) WithKrbServicePrincipalName(value string) *HDFSArtifactRepositoryApplyConfiguration {
	b.KrbServicePrincipalName = &value
	return b
}

// WithAddresses adds the given value to the Addresses field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Addresses field.
func (b *HDFSArtifactRepositoryApplyConfiguration) WithAddresses(values ...string) *HDFSArtifactRepositoryApplyConfiguration {
	for i := range values {
		b.Addresses = append(b.Addresses, values[i])
	}
	return b
}

// WithHDFSUser sets the HDFSUser field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the HDFSUser field is set to the value of the last call.
func (b *HDFSArtifactRepositoryApplyConfiguration) WithHDFSUser(value string) *HDFSArtifactRepositoryApplyConfiguration {
	b.HDFSUser = &value
	return b
}

// WithPathFormat sets the PathFormat field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PathFormat field is set to the value of the last call.
func (b *HDFSArtifactRepositoryApplyConfiguration) WithPathFormat(value string) *HDFSArtifactRepositoryApplyConfiguration {
	b.PathFormat = &value
	return b
}

// WithForce sets the Force field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Force field is set to the value of the last call.
func (b *HDFSArtifactRepositoryApplyConfiguration) WithForce(value bool) *HDFSArtifactRepositoryApplyConfiguration {
	b.Force = &value
	return b
}
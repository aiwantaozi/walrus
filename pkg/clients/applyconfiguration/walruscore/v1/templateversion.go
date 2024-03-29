// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

package v1

// TemplateVersionApplyConfiguration represents an declarative configuration of the TemplateVersion type for use
// with apply.
type TemplateVersionApplyConfiguration struct {
	Version             *string                                 `json:"version,omitempty"`
	URL                 *string                                 `json:"url,omitempty"`
	SchemaRef           *LocalObjectReferenceApplyConfiguration `json:"schemaRef,omitempty"`
	OriginalUISchemaRef *LocalObjectReferenceApplyConfiguration `json:"originalUISchemaRef,omitempty"`
	UISchemaRef         *LocalObjectReferenceApplyConfiguration `json:"uiSchemaRef,omitempty"`
}

// TemplateVersionApplyConfiguration constructs an declarative configuration of the TemplateVersion type for use with
// apply.
func TemplateVersion() *TemplateVersionApplyConfiguration {
	return &TemplateVersionApplyConfiguration{}
}

// WithVersion sets the Version field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Version field is set to the value of the last call.
func (b *TemplateVersionApplyConfiguration) WithVersion(value string) *TemplateVersionApplyConfiguration {
	b.Version = &value
	return b
}

// WithURL sets the URL field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the URL field is set to the value of the last call.
func (b *TemplateVersionApplyConfiguration) WithURL(value string) *TemplateVersionApplyConfiguration {
	b.URL = &value
	return b
}

// WithSchemaRef sets the SchemaRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SchemaRef field is set to the value of the last call.
func (b *TemplateVersionApplyConfiguration) WithSchemaRef(value *LocalObjectReferenceApplyConfiguration) *TemplateVersionApplyConfiguration {
	b.SchemaRef = value
	return b
}

// WithOriginalUISchemaRef sets the OriginalUISchemaRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the OriginalUISchemaRef field is set to the value of the last call.
func (b *TemplateVersionApplyConfiguration) WithOriginalUISchemaRef(value *LocalObjectReferenceApplyConfiguration) *TemplateVersionApplyConfiguration {
	b.OriginalUISchemaRef = value
	return b
}

// WithUISchemaRef sets the UISchemaRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UISchemaRef field is set to the value of the last call.
func (b *TemplateVersionApplyConfiguration) WithUISchemaRef(value *LocalObjectReferenceApplyConfiguration) *TemplateVersionApplyConfiguration {
	b.UISchemaRef = value
	return b
}

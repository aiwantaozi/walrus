package types

import (
	core "k8s.io/api/core/v1"
)

const (
	AuthTypeToken = "token"
	AuthTypeBasic = "basic"
)

const (
	TemplateFormatHelm      string = "helm"
	TemplateFormatTerraform string = "terraform"
)

// TemplateVersion defines the version of Template.
type TemplateVersion struct {
	Version string `json:"version,omitempty"`

	// URL of downloading the version.
	// host.
	// +required
	URL string `json:"url"`

	// Schema holds the schema for the template version.
	Schema core.LocalObjectReference `json:"schema,omitempty"`

	// OriginalUISchema holds the original UI schema for the template version.
	OriginalUISchema core.LocalObjectReference `json:"originalUISchema,omitempty"`

	// UISchema holds the UI schema for the template version.
	UISchema core.LocalObjectReference `json:"uiSchema,omitempty"`
}

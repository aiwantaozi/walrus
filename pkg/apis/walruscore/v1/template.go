package v1

import (
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// Template is the schema for the templates API.
//
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:crd-gen:resource:scope="Namespaced",subResources=["status"]
type Template struct {
	meta.TypeMeta   `json:",inline"`
	meta.ObjectMeta `json:"metadata,omitempty"`

	Spec   TemplateSpec   `json:"spec,omitempty"`
	Status TemplateStatus `json:"status,omitempty"`
}

var _ runtime.Object = (*Template)(nil)

type TemplateSpec struct {
}

type TemplateStatus struct {
}

//// TemplateSpec defines the desired state of Template.
//type TemplateSpec struct {
//	// TemplateFormat of the content.
//	// +required
//	TemplateFormat string `json:"templateFormat"`
//
//	// Description of the template.
//	// +optional
//	Description string `json:"description,omitempty"`
//
//	// VCSRepository specifies the vcs repository configure.
//	// +optional
//	VCSRepository *VCSRepository `json:"vcsRepository,omitempty"`
//
//	// HelmOCIChart specifies the OCI format helm chart configure.
//	// +optional
//	HelmOCIChart *HelmOCIChart `json:"helmOCIChart,omitempty"`
//
//	// HelmChart specifies the tgz format helm chart configure.
//	// +optional
//	HelmChart *HelmChart `json:"helmChart,omitempty"`
//}

// TemplateStatus defines the observed state of Template.
//type TemplateStatus struct {
//	// The original name of the template, name generate from chart.yaml, terraform git repo name.
//	OriginalName string `json:"originalName,omitempty"`
//
//	// A URL to an SVG or PNG image to be used as an icon.
//	Icon string `json:"icon,omitempty"`
//
//	// Versions holds the versions for the template.
//	Versions []TemplateVersion `json:"versions,omitempty"`
//
//	// Conditions holds the conditions for the template.
//	Conditions []meta.Condition `json:"conditions,omitempty"`
//}

// TemplateList holds the list of Template.
//
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type TemplateList struct {
	meta.TypeMeta `json:",inline"`
	meta.ListMeta `json:"metadata,omitempty"`

	Items []Template `json:"items"`
}

var _ runtime.Object = (*TemplateList)(nil)

// VCSRepository specifies the vcs repository of the template.
type VCSRepository struct {
	TokenAuth `json:",inline"`

	// URL of download the template from vsc repository with ref and subpath, e.g. https://github.com/walrus-catalog/terraform-static-redis?ref=main.
	// +required
	URL string `json:"url"`
}

// HelmOCIChart specifies the OCI format helm chart configure.
type HelmOCIChart struct {
	// BasicAuth containing authentication credentials for oci chart.
	BasicAuth `json:",inline"`

	// URL of download the template from oci image with ref, e.g. oci://registry-1.docker.io/bitnamicharts/mysql?ref=10.1.0.
	// +required
	URL string `json:"url"`
}

// HelmChart specifies the tgz format helm chart configure.
type HelmChart struct {
	// BasicAuth containing authentication credentials for helm chart.
	BasicAuth `json:",inline"`

	// URL of download the template from chart tgz address, e.g. https://charts.bitnami.com/bitnami/phpbb-16.5.0.tgz.
	// +required
	URL string `json:"url"`
}

// TemplateVersion defines the version of Template.
type TemplateVersion struct {
	Version string `json:"version,omitempty"`

	// URL of downloading the version.
	// host.
	// +required
	URL string `json:"url"`

	// Schema holds the schema for the template version.
	Schema *ObjectReference `json:"schema,omitempty"`

	// OriginalUISchema holds the original UI schema for the template version.
	OriginalUISchema *ObjectReference `json:"originalUISchema,omitempty"`

	// UISchema holds the UI schema for the template version.
	UISchema *ObjectReference `json:"uiSchema,omitempty"`
}

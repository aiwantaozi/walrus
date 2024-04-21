package v1

import (
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
)

// Template is the schema for the templates API.
//
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:crd-gen:resource:scope="Namespaced",subResources=["status"]
type Template struct {
	meta.TypeMeta   `json:",inline"`
	meta.ObjectMeta `json:"metadata,omitempty"`

	Spec   TemplateSpec   `json:"spec"`
	Status TemplateStatus `json:"status,omitempty"`
}

var _ runtime.Object = (*Template)(nil)

// TemplateReference is a reference to a template.
type TemplateReference struct {
	// Namespace is the namespace of the template.
	Namespace string `json:"namespace"`
	// Name is the name of the template.
	Name string `json:"name"`
}

func (in *TemplateReference) ToNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Namespace: in.Namespace,
		Name:      in.Name,
	}
}

// TemplateSpec defines the desired state of Template.
type TemplateSpec struct {
	// TemplateFormat of the content.
	//
	// +k8s:validation:cel[0]:rule="oldSelf == self"
	// +k8s:validation:cel[0]:message="immutable field"
	TemplateFormat string `json:"templateFormat"`

	// Description of the template.
	Description string `json:"description,omitempty"`

	// VCSRepository specifies the configuration for the VCS repository.
	VCSRepository *VCSRepository `json:"vcsRepository"`
}

// TemplateStatus defines the observed state of Template.
type TemplateStatus struct {
	// StatusDescriptor defines the status of the catalog.
	StatusDescriptor `json:",inline"`

	// Project is the project that the template belongs to.
	Project string `json:"project,omitempty"`

	// URL of the template.
	URL string `json:"url,omitempty"`

	// A URL to an SVG or PNG image to be used as an icon.
	Icon string `json:"icon,omitempty"`

	// Versions contains the versions for the template.
	Versions []TemplateVersion `json:"versions,omitempty"`

	// LastSuccessfulSyncTime record the last sync time the template was synchronized successfully.
	LastSuccessfulSyncTime *meta.Time `json:"lastSuccessfulSyncTime,omitempty"`
}

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
	// Platform of the vcs repository.
	//
	// +k8s:validation:enum=["GitHub","GitLab","Gitee"]
	// +k8s:validation:cel[0]:rule="oldSelf == self"
	// +k8s:validation:cel[0]:message="immutable field"
	Platform VCSPlatform `json:"platform"`

	// URL of download the template from vsc repository, may include reference and subpath.
	// e.g. https://github.com/walrus-catalog/terraform-static-redis.
	//
	// +k8s:validation:cel[0]:rule="oldSelf == self"
	// +k8s:validation:cel[0]:message="immutable field"
	URL string `json:"url"`
}

// TemplateVersion defines the version of Template.
type TemplateVersion struct {
	// Version of the template.
	Version string `json:"version,omitempty"`

	// URL of downloading the template version with ref and subpath.
	URL string `json:"url"`

	// TemplateSchemaName holds the template schema name for the template version.
	TemplateSchemaName *string `json:"templateSchemaName,omitempty"`

	// OriginalUISchemaName holds the original UI schema name for the template version.
	OriginalUISchemaName *string `json:"originalUISchemaName,omitempty"`

	// UISchemaName holds the UI schema name for the template version.
	UISchemaName *string `json:"uiSchemaName,omitempty"`

	// Removed indicate the template version is removed.
	Removed bool `json:"removed,omitempty"`
}

// TemplateReferenceWithVersion is a reference to a template with a specific version.
type TemplateReferenceWithVersion struct {
	TemplateReference `json:",inline"`

	// Version is a specific version of the template.
	Version string `json:"version,omitempty"`
}

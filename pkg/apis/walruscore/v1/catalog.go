package v1

import (
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// Catalog is the schema for the catalogs API.
//
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:crd-gen:resource:scope="Namespaced",subResources=["status"]
type Catalog struct {
	meta.TypeMeta   `json:",inline"`
	meta.ObjectMeta `json:"metadata,omitempty"`

	Spec   CatalogSpec   `json:"spec,omitempty"`
	Status CatalogStatus `json:"status,omitempty"`
}

var _ runtime.Object = (*Catalog)(nil)

// CatalogSpec defines the desired state of Catalog.
type CatalogSpec struct {
	// TemplateFormat of the catalog.
	// +required
	TemplateFormat string `json:"templateFormat"`

	// Description of the catalog.
	// +optional
	Description string `json:"description,omitempty"`

	// Filtering specifies the filtering rules for the catalog.
	// +optional
	Filtering *Filtering `json:"filtering,omitempty"`

	// VCSSource specifies the vcs source configure.
	// +optional
	VCSSource *VCSSource `json:"vcsSource,omitempty"`

	// HelmRepositorySource specifies the helm repository configure.
	// +optional
	HelmRepositorySource *HelmRepositorySource `json:"helmRepositorySource,omitempty"`

	// OCIRegistrySource specifies the OCI registry configure.
	// +optional
	OCIRegistrySource *OCIRegistrySource `json:"ociRegistrySource,omitempty"`
}

// CatalogStatus defines the observed state of Catalog.
type CatalogStatus struct {
	// LastSyncTime record the last auto/manual sync catalog time.
	LastSyncTime meta.Time `json:"lastSyncTime"`

	// TemplateCount include template count.
	TemplateCount int64 `json:"templateCount"`

	// Conditions holds the conditions for the catalog.
	// +optional
	Conditions []meta.Condition `json:"conditions,omitempty"`
}

// CatalogList holds the list of Catalog.
//
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type CatalogList struct {
	meta.TypeMeta `json:",inline"`
	meta.ListMeta `json:"metadata,omitempty"`

	Items []Catalog `json:"items"`
}

var _ runtime.Object = (*CatalogList)(nil)

// Filtering specifies the filtering rules for the catalog's template.
type Filtering struct {
	// IncludeFilter specifies the regular expression for matching the include template name.
	// +optional
	IncludeFilter string `json:"includeFilter,omitempty"`

	// ExcludeFilter specifies the regular expression for matching the exclude template name.
	// +optional
	ExcludeFilter string `json:"excludeFilter,omitempty"`
}

// VCSSource specifies the vcs source configure.
type VCSSource struct {
	// TokenAuth containing authentication credentials for vcs source.
	TokenAuth `json:",inline"`

	// +k8s:validation:enum=["github","gitlab","gitee"]
	// +required
	Platform string `json:"platform"`

	// URL of the source address, a valid URL contains at least a protocol and host.
	// +required
	URL string `json:"url"`
}

// OCIRegistrySource specifies the OCI registry configure.
type OCIRegistrySource struct {
	// BasicAuth containing authentication credentials for OCI registry.
	BasicAuth `json:",inline"`

	// Domain for OCI registry.
	// +optional
	Domain string `json:"domain,omitempty"`
}

// HelmRepositorySource specifies the helm repository configure.
type HelmRepositorySource struct {
	// BasicAuth containing authentication credentials for helm repository.
	BasicAuth `json:",inline"`

	// URL of the source address, a valid URL contains at least a protocol and host.
	// +required
	URL string `json:"url"`
}

type BasicAuth struct {
	// Username for OCI registry.
	// +optional
	Username string `json:"username,omitempty"`

	// Password for OCI registry.
	// +optional
	Password string `json:"password,omitempty"`

	// SecretRef specifies the Secret containing authentication credentials for helm repository.
	// For HTTP/S basic auth the secret must contain 'username' and 'password' fields.
	// +optional
	SecretRef *BasicAuthObjectReference `json:"secretRef,omitempty"`
}

type TokenAuth struct {
	// Token for HTTP/S bear token.
	// +optional
	Token string `json:"token,omitempty"`

	// SecretRef specifies the Secret containing authentication credentials.
	// For HTTP/S bear token the secret must contain 'token' field.
	// +optional
	SecretRef *TokenObjectReference `json:"secretRef,omitempty"`
}

// TODO:(michelia)
type LocalObjectReference struct {
	// Name of the object.
	Name string `json:"name"`
}

type TokenObjectReference struct {
	LocalObjectReference `json:",inline"`
}

// TODO:(michelia)
type BasicAuthObjectReference struct {
	LocalObjectReference `json:",inline"`
}

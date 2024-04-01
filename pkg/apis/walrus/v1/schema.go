package v1

import (
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// Schema API for the template's version.
//
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:apireg-gen:resource:scope="Namespaced",categories=["walrus"],subResources=["status"]
type Schema struct {
	meta.TypeMeta   `json:",inline"`
	meta.ObjectMeta `json:"metadata,omitempty"`

	Spec   SchemaSpec   `json:"spec,omitempty"`
	Status SchemaStatus `json:"status,omitempty"`
}

var _ runtime.Object = (*Setting)(nil)

// SchemaSpec defines the desired state of Setting.
type SchemaSpec struct {
	// Value contains the user input schema data,
	// it is provided as a write-only input field.
	Value []byte `json:"value,omitempty"`
}

// SchemaStatus defines the template version's schema.
type SchemaStatus struct {

	// Value is the current value of the schema.
	Value []byte `json:"value"`

	// Conditions holds the conditions for the catalog.
	// +optional
	Conditions []meta.Condition `json:"conditions,omitempty"`
}

// SchemaList holds the list of Schema.
//
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type SchemaList struct {
	meta.TypeMeta `json:",inline"`
	meta.ListMeta `json:"metadata,omitempty"`

	Items []Schema `json:"items"`
}

var _ runtime.Object = (*SchemaList)(nil)

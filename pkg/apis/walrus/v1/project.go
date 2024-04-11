package v1

import (
	"reflect"

	core "k8s.io/api/core/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// Project is the schema for the projects API.
//
// +genclient
// +genclient:onlyVerbs=create,get,list,watch,apply,update,patch,delete,deleteCollection
// +genclient:method=GetSubjects,verb=get,subresource=subjects,result=ProjectSubjects
// +genclient:method=UpdateSubjects,verb=update,subresource=subjects,input=ProjectSubjects,result=ProjectSubjects
// +genclient:method=PatchSubjects,verb=update,subresource=subjects,input=ProjectSubjects,result=ProjectSubjects
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:apireg-gen:resource:scope="Namespaced",categories=["walrus"],shortName=["proj"]
type Project struct {
	meta.TypeMeta   `json:",inline"`
	meta.ObjectMeta `json:"metadata,omitempty"`

	Spec   ProjectSpec   `json:"spec,omitempty"`
	Status ProjectStatus `json:"status,omitempty"`
}

var _ runtime.Object = (*Project)(nil)

// ProjectSpec defines the desired state of Project.
type ProjectSpec struct {
	// DisplayName is the display name of the project.
	DisplayName string `json:"displayName,omitempty"`

	// Description is the description of the project.
	Description string `json:"description,omitempty"`
}

// ProjectStatus defines the observed state of Project.
type ProjectStatus struct {
	// Phase is the current phase of the project.
	Phase core.NamespacePhase `json:"phase"`
}

func (in *Project) Equal(in2 *Project) bool {
	return reflect.DeepEqual(in.Spec, in2.Spec) &&
		in.Status.Phase == in2.Status.Phase
}

// ProjectList holds the list of Project.
//
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ProjectList struct {
	meta.TypeMeta `json:",inline"`
	meta.ListMeta `json:"metadata,omitempty"`

	Items []Project `json:"items"`
}

var _ runtime.Object = (*ProjectList)(nil)

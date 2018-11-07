package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Stage defines a workflow stage. Only one of Spec and Template should be specified.
type Stage struct {
	// Metadata for the resource, like kind and apiversion
	metav1.TypeMeta `json:",inline"`
	// Metadata for the particular object, including name, namespace, labels, etc
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// Stage specification without using template
	Spec StageSpec `json:"spec,omitempty"`
	// Use stage template
	Template TemplateRef `json:"template,omitempty"`
}

// StageSpec defines stage specification.
// Exact one workload should be specified.
type StageSpec struct {
	// Pod kind workload
	Pod PodWorkload `json:"pod,omitempty"`
	// CRD kind workload
	CRD CRDWorkload `json:"crd:omitempty"`
}

// TemplateRef refers to a stage template and defines necessary arguments.
type TemplateRef struct {
	// Template name
	Name string `json:"name"`
	// Arguments passed to the template
	Arguments []ArgumentItem `json:"arguments"`
}
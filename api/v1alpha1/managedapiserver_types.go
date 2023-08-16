/*
Copyright 2023 TrooperT.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ManagedApiServerSpec defines the desired state of ManagedApiServer
type ManagedApiServerSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of ManagedApiServer. Edit managedapiserver_types.go to remove/update
	Foo string `json:"foo,omitempty"`
}

// ManagedApiServerStatus defines the observed state of ManagedApiServer
type ManagedApiServerStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ManagedApiServer is the Schema for the managedapiservers API
type ManagedApiServer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ManagedApiServerSpec   `json:"spec,omitempty"`
	Status ManagedApiServerStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ManagedApiServerList contains a list of ManagedApiServer
type ManagedApiServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ManagedApiServer `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ManagedApiServer{}, &ManagedApiServerList{})
}

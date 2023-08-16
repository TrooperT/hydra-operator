/*
Copyright 2023 TrooperT.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ManagedControllerManagerSpec defines the desired state of ManagedControllerManager
type ManagedControllerManagerSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of ManagedControllerManager. Edit managedcontrollermanager_types.go to remove/update
	Foo string `json:"foo,omitempty"`
}

// ManagedControllerManagerStatus defines the observed state of ManagedControllerManager
type ManagedControllerManagerStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ManagedControllerManager is the Schema for the managedcontrollermanagers API
type ManagedControllerManager struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ManagedControllerManagerSpec   `json:"spec,omitempty"`
	Status ManagedControllerManagerStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ManagedControllerManagerList contains a list of ManagedControllerManager
type ManagedControllerManagerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ManagedControllerManager `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ManagedControllerManager{}, &ManagedControllerManagerList{})
}

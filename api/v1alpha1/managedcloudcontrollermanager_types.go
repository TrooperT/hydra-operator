/*
Copyright 2023 TrooperT.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ManagedCloudControllerManagerSpec defines the desired state of ManagedCloudControllerManager
type ManagedCloudControllerManagerSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of ManagedCloudControllerManager. Edit managedcloudcontrollermanager_types.go to remove/update
	Foo string `json:"foo,omitempty"`
}

// ManagedCloudControllerManagerStatus defines the observed state of ManagedCloudControllerManager
type ManagedCloudControllerManagerStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ManagedCloudControllerManager is the Schema for the managedcloudcontrollermanagers API
type ManagedCloudControllerManager struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ManagedCloudControllerManagerSpec   `json:"spec,omitempty"`
	Status ManagedCloudControllerManagerStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ManagedCloudControllerManagerList contains a list of ManagedCloudControllerManager
type ManagedCloudControllerManagerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ManagedCloudControllerManager `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ManagedCloudControllerManager{}, &ManagedCloudControllerManagerList{})
}

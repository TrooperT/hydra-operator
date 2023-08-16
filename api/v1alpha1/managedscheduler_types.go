/*
Copyright 2023 TrooperT.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ManagedSchedulerSpec defines the desired state of ManagedScheduler
type ManagedSchedulerSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of ManagedScheduler. Edit managedscheduler_types.go to remove/update
	Foo string `json:"foo,omitempty"`
}

// ManagedSchedulerStatus defines the observed state of ManagedScheduler
type ManagedSchedulerStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ManagedScheduler is the Schema for the managedschedulers API
type ManagedScheduler struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ManagedSchedulerSpec   `json:"spec,omitempty"`
	Status ManagedSchedulerStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ManagedSchedulerList contains a list of ManagedScheduler
type ManagedSchedulerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ManagedScheduler `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ManagedScheduler{}, &ManagedSchedulerList{})
}

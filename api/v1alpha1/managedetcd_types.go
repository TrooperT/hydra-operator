/*
Copyright 2023 TrooperT.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ManagedEtcdSpec defines the desired state of ManagedEtcd
type ManagedEtcdSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Version holds the desired version of the ManagedEtcd.
	Version string `json:"version"`
}

// ManagedEtcdStatus defines the observed state of ManagedEtcd
type ManagedEtcdStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Version holds the observed version of the ManagedEtcd.
	// While an upgrade is in progress this value will be the
	// version of the ManagedEtcd when the upgrade began.
	// +optional
	Version string `json:"version"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ManagedEtcd is the Schema for the managedetcds API
type ManagedEtcd struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ManagedEtcdSpec   `json:"spec,omitempty"`
	Status ManagedEtcdStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ManagedEtcdList contains a list of ManagedEtcd
type ManagedEtcdList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ManagedEtcd `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ManagedEtcd{}, &ManagedEtcdList{})
}

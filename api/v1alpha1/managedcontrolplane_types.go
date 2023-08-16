/*
Copyright 2023 TrooperT.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ManagedControlPlanePhase is a type for the ManagedControlPlane's
// phase constants.
type ManagedControlPlanePhase string

const (
	// ManagedControlPlanePhaseCreating indicates the ManagedControlPlane
	// is under creation.
	ManagedControlPlanePhaseCreating = ManagedControlPlanePhase("creating")

	// ManagedControlPlanePhaseFailed indicates ManagedControlPlane creation
	// failed. The system likely requires user intervention.
	ManagedControlPlanePhaseFailed = ManagedControlPlanePhase("failed")

	// ManagedControlPlanePhaseRunning indicates the ManagedControlPlane
	// is ready.
	ManagedControlPlanePhaseRunning = ManagedControlPlanePhase("running")

	// ManagedControlPlanePhaseUnhealthy indicates the ManagedControlPlane
	// was up and running, but is unhealthy now.
	// The system likely requires user intervention.
	ManagedControlPlanePhaseUnhealthy = ManagedControlPlanePhase("unhealthy")

	// ManagedControlPlanePhaseUpdating indicates the ManagedControlPlane
	// is in the process of rolling update
	ManagedControlPlanePhaseUpdating = ManagedControlPlanePhase("updating")

	// ManagedControlPlanePhaseUpdateFailed indicates the ManagedControlPlane's
	// rolling update failed and likely requires user intervention.
	ManagedControlPlanePhaseUpdateFailed = ManagedControlPlanePhase("updateFailed")

	// ManagedControlPlanePhaseDeleting indicates that the ManagedControlPlane
	// is being deleted.
	ManagedControlPlanePhaseDeleting = ManagedControlPlanePhase("deleting")

	// ManagedControlPlanePhaseEmpty is useful for the initial reconcile,
	// before we even state the phase as creating.
	ManagedControlPlanePhaseEmpty = ManagedControlPlanePhase("")
)

// ManagedControlPlaneSpec defines the desired state of ManagedControlPlane
type ManagedControlPlaneSpec struct {
	// Version holds the desired version of the ManagedControlPlane.
	Version string `json:"version"`
	// Unique Tenant ID
	// Uses github.com/rs/xid
	TenantID string `json:"tenantID"`
	// Should this ManagedControlPlane be Highly Available
	EnableHA bool `json:"enableHA"`
}

// ManagedControlPlaneStatus defines the observed state of ManagedControlPlane
type ManagedControlPlaneStatus struct {
	// APIEndpoints represents the endpoints to communicate with the
	// ManagedControlPlane.
	// +optional
	APIEndpoints []APIEndpoint `json:"apiEndpoints,"`

	// Version holds the observed version of the ManagedControlPlane.
	// While an upgrade is in progress this value will be the
	// version of the ManagedControlPlane when the upgrade began.
	// +optional
	Version string `json:"version"`
}

// APIEndpoint represents a reachable Kubernetes API endpoint.
type APIEndpoint struct {
	// The hostname on which the API server is serving.
	Host string `json:"host"`

	// The port on which the API server is serving.
	Port int `json:"port"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:resource:scope=Cluster

// ManagedControlPlane is the Schema for the managedcontrolplanes API
type ManagedControlPlane struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ManagedControlPlaneSpec   `json:"spec,omitempty"`
	Status ManagedControlPlaneStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ManagedControlPlaneList contains a list of ManagedControlPlane
type ManagedControlPlaneList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ManagedControlPlane `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ManagedControlPlane{}, &ManagedControlPlaneList{})
}

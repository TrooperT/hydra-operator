/*
Copyright 2023 TrooperT.
*/

package controllers

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	iaasv1alpha1 "github.com/TrooperT/hydra-operator/api/v1alpha1"
)

// ManagedApiServerReconciler reconciles a ManagedApiServer object
type ManagedApiServerReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=iaas.homelabs.io,resources=managedapiservers,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=iaas.homelabs.io,resources=managedapiservers/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=iaas.homelabs.io,resources=managedapiservers/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the ManagedApiServer object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.14.1/pkg/reconcile
func (r *ManagedApiServerReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	// TODO(user): your logic here

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ManagedApiServerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&iaasv1alpha1.ManagedApiServer{}).
		Complete(r)
}

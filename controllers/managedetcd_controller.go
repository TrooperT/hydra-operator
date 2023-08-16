/*
Copyright 2023 TrooperT.
*/

package controllers

import (
	"context"
	// "fmt"
	// "time"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	// "k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/klog/v2"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	iaasv1alpha1 "github.com/TrooperT/hydra-operator/api/v1alpha1"
	// certmgrv1 "github.com/cert-manager/cert-manager/pkg/apis/certmanager/v1"
)

// ManagedEtcdReconciler reconciles a ManagedEtcd object
type ManagedEtcdReconciler struct {
	client.Client
	Scheme *runtime.Scheme
	Log    logr.Logger
}

//+kubebuilder:rbac:groups=iaas.homelabs.io,resources=managedetcds,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=iaas.homelabs.io,resources=managedetcds/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=iaas.homelabs.io,resources=managedetcds/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the ManagedEtcd object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.14.1/pkg/reconcile
func (r *ManagedEtcdReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := klog.FromContext(ctx).WithName("ManagedEtcdReconciler")
	log.Info("Begin Reconcile")

	metcd := &iaasv1alpha1.ManagedEtcd{}
	err := r.Get(ctx, req.NamespacedName, metcd)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			log.Info("ManagedEtcd resource not found. Ignoring since object must be deleted")
			return ctrl.Result{}, nil
		}
		// Error reading the object - requeue the request.
		log.Error(err, "Failed to get ManagedEtcd. Requeuing")
		return ctrl.Result{}, err
	}
	return ctrl.Result{}, nil
}

func (r *ManagedEtcdReconciler) peerServiceForManagedEtcdReconciler(metcd *iaasv1alpha1.ManagedEtcd) *corev1.Service {
	svc := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      ServiceEtcd,
			Namespace: metcd.Namespace,
			Labels: map[string]string{
				"app":       AppManagedControlPlane,
				"component": ComponentEtcdPeer,
			},
		},
		Spec: corev1.ServiceSpec{
			Type:                     corev1.ServiceTypeClusterIP,
			ClusterIP:                corev1.ClusterIPNone,
			PublishNotReadyAddresses: true,
			Selector: map[string]string{
				"app":       AppManagedControlPlane,
				"component": ComponentEtcd,
			},
			Ports: []corev1.ServicePort{
				{
					Name:       "client",
					Protocol:   "TCP",
					Port:       2379,
					TargetPort: intstr.FromInt(EtcdClientPort),
				},
				{
					Name:       "peer",
					Protocol:   "TCP",
					Port:       2380,
					TargetPort: intstr.FromInt(EtcdPeerPort),
				},
			},
		},
	}

	ctrl.SetControllerReference(metcd, svc, r.Scheme)
	return svc
}

func (r *ManagedEtcdReconciler) clientServiceForManagedEtcdReconciler(metcd *iaasv1alpha1.ManagedEtcd) *corev1.Service {
	svc := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      ServiceEtcdClient,
			Namespace: metcd.Namespace,
			Labels: map[string]string{
				"app":       AppManagedControlPlane,
				"component": ComponentEtcdClient,
			},
		},
		Spec: corev1.ServiceSpec{
			Type:            corev1.ServiceTypeClusterIP,
			SessionAffinity: corev1.ServiceAffinityClientIP,
			Selector: map[string]string{
				"app":       AppManagedControlPlane,
				"component": ComponentEtcd,
			},
			Ports: []corev1.ServicePort{
				{
					Name:       "client",
					Protocol:   "TCP",
					Port:       EtcdClientPort,
					TargetPort: intstr.FromInt(EtcdClientPort),
				},
			},
		},
	}

	ctrl.SetControllerReference(metcd, svc, r.Scheme)
	return svc
}

func (r *ManagedEtcdReconciler) statefulSetForManagedEtcdReconciler(metcd *iaasv1alpha1.ManagedEtcd) *appsv1.StatefulSet {
	sts := &appsv1.StatefulSet{
		ObjectMeta: metav1.ObjectMeta{
			Name:      StatefulSetEtcd,
			Namespace: metcd.Namespace,
			Labels: map[string]string{
				"app":       AppManagedControlPlane,
				"component": ComponentEtcd,
			},
		},
		Spec: appsv1.StatefulSetSpec{
			ServiceName: ServiceEtcd,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app":       AppManagedControlPlane,
				},
			},
		},
	}

	ctrl.SetControllerReference(metcd, sts, r.Scheme)
	return sts

}

// SetupWithManager sets up the controller with the Manager.
func (r *ManagedEtcdReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&iaasv1alpha1.ManagedEtcd{}).
		Complete(r)
}

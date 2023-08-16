/*
Copyright 2023 TrooperT.
*/

package controllers

import (
	"context"
	"fmt"
	"time"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/klog/v2"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	iaasv1alpha1 "github.com/TrooperT/hydra-operator/api/v1alpha1"
	certmgrv1 "github.com/cert-manager/cert-manager/pkg/apis/certmanager/v1"

	"github.com/rs/xid"
)

// ManagedControlPlaneReconciler reconciles a ManagedControlPlane object
type ManagedControlPlaneReconciler struct {
	client.Client
	Scheme *runtime.Scheme
	Log    logr.Logger
}

//+kubebuilder:rbac:groups=iaas.homelabs.io,resources=managedcontrolplanes,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=iaas.homelabs.io,resources=managedcontrolplanes/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=iaas.homelabs.io,resources=managedcontrolplanes/finalizers,verbs=update

// generate rbac to get,list, and watch pods
// +kubebuilder:rbac:groups=core,resources=namespaces,verbs=get;list;watch;create;update;patch;delete

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the ManagedControlPlane object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.14.1/pkg/reconcile
func (r *ManagedControlPlaneReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	// _ = klog.FromContext(ctx)
	log := klog.FromContext(ctx).WithName("ManagedControlPlaneReconciler")
	log.Info("Begin Reconcile")
	// TODO(user): your logic here
	mcp := &iaasv1alpha1.ManagedControlPlane{}
	// log.V(4).Info("Req Values", "Namespace", req.Namespace, "Name", req.Name, "NamespacedName", req.NamespacedName)
	err := r.Get(ctx, req.NamespacedName, mcp)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			log.Info("ManagedControlPlane resource not found. Ignoring since object must be deleted")
			return ctrl.Result{}, nil
		}
		// Error reading the object - requeue the request.
		log.Error(err, "Failed to get ManagedControlPlane. Requeuing")
		return ctrl.Result{}, err
	}
	log.V(4).Info("ManagedControlPlane Values", "Version", mcp.Spec.Version, "TenantID", mcp.Spec.TenantID, "EnableHA", mcp.Spec.EnableHA)
	if mcp.Spec.TenantID == "" {
		id := xid.New()
		mcp.Spec.TenantID = id.String()
		if err = r.Update(ctx, mcp); err != nil {
			log.Error(err, "Failed to update ManagedControlPlane", "ManagedControlPlane.Namespace", mcp.Namespace, "ManagedControlPlane.Name", mcp.Name)
			return ctrl.Result{}, err
		}
		log.V(4).Info("ManagedControlPlane Set TenantID")
		return ctrl.Result{}, nil
	}

	// Check if the tenant Namespace already exists, if not create a new one
	foundNS := &corev1.Namespace{}
	err = r.Get(ctx, types.NamespacedName{
		Name: fmt.Sprintf("tenant-%s", mcp.Spec.TenantID),
	}, foundNS)
	// log.Info("Found Results", "Found", foundNS)
	if err != nil && errors.IsNotFound(err) {
		// Define a new Namespace
		ns := r.nameSpaceForManagedControlPlane(mcp)
		log.Info("Creating new Namespace", "Namespace.Name", ns.Name)
		err = r.Create(ctx, ns)
		if err != nil {
			log.Error(err, "Failed to create new Namespace", "Namespace.Name", ns.Name)
			return ctrl.Result{}, err
		}
		// Namespace created successfully - return and requeue
		return ctrl.Result{RequeueAfter: 5 * time.Second}, nil
	} else if err != nil {
		log.Error(err, "Failed to get Namespace")
		return ctrl.Result{}, err
	}
	// Check if the tenant Namespace is Active, if not requeue
	if foundNS.Status.Phase != "Active" {
		log.V(4).Info("ManagedControlPlane Tenant Namespace Not Ready. Requeuing")
		return ctrl.Result{RequeueAfter: 5 * time.Second}, nil
	}

	// Check if the tenant External APIEndpoint already exists, if not create a new one
	foundSVC := &corev1.Service{}
	err = r.Get(ctx, types.NamespacedName{
		Namespace: foundNS.Name,
		Name:      "control-plane-external",
	}, foundSVC)
	if err != nil && errors.IsNotFound(err) {
		svc := r.loadBalancerForManagedControlPlane(mcp)
		log.Info("Creating new APIEndpoint", "Service", svc)
		err = r.Create(ctx, svc)
		if err != nil {
			log.Error(err, "Failed to create new APIEndpoint", "Service", svc)
			return ctrl.Result{}, err
		}
		// APIEndpoint created successfully - return and requeue
		return ctrl.Result{RequeueAfter: 5 * time.Second}, nil
	}
	// Check if the Loadbalancer service has been assigned an IP.
	// If service has more IPs than ManagedControlPlane, then update and requeue
	// TODO: Restructure this bit to support both:
	// 			external APIEndpoint (Users, etc)
	//			internal APIEndpoint (nodes)
	if len(foundSVC.Status.LoadBalancer.Ingress) > len(mcp.Status.APIEndpoints) {
		eps := []iaasv1alpha1.APIEndpoint{}
		for i := 0; i < len(foundSVC.Status.LoadBalancer.Ingress); i++ {
			eps = append(eps, iaasv1alpha1.APIEndpoint{
				// TODO: logic to support DNS Based Loadbalancers e.g. AWS
				Host: foundSVC.Status.LoadBalancer.Ingress[i].IP,
				Port: int(foundSVC.Spec.Ports[0].Port),
			})
		}
		mcp.Status.APIEndpoints = eps
		if err = r.Status().Update(ctx, mcp); err != nil {
			log.Error(err, "Failed to update ManagedControlPlane", "ManagedControlPlane.Namespace", mcp.Namespace, "ManagedControlPlane.Name", mcp.Name)
			return ctrl.Result{}, err
		}
		log.V(4).Info("ManagedControlPlane Set APIEndpoints")
		return ctrl.Result{}, nil

	}

	log.V(4).Info("End Reconcile")
	return ctrl.Result{}, nil
}

func (r *ManagedControlPlaneReconciler) nameSpaceForManagedControlPlane(mcp *iaasv1alpha1.ManagedControlPlane) *corev1.Namespace {
	tenantNS := &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: fmt.Sprintf("tenant-%s", mcp.Spec.TenantID),
		},
	}
	ctrl.SetControllerReference(mcp, tenantNS, r.Scheme)
	return tenantNS
}

func (r *ManagedControlPlaneReconciler) loadBalancerForManagedControlPlane(mcp *iaasv1alpha1.ManagedControlPlane) *corev1.Service {

	svc := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      ServiceAPIServerExternal,
			Namespace: fmt.Sprintf("tenant-%s", mcp.Spec.TenantID),
		},
		Spec: corev1.ServiceSpec{
			Type:            ServiceAPIServerExternalType,
			SessionAffinity: corev1.ServiceAffinityClientIP,
			Selector: map[string]string{
				"app":       AppManagedControlPlane,
				"component": ComponentAPIServer,
			},
			Ports: []corev1.ServicePort{
				{
					Name:       "apiserver",
					Protocol:   "TCP",
					Port:       APIServerPort,
					TargetPort: intstr.FromInt(APIServerPort),
				},
			},
		},
	}

	ctrl.SetControllerReference(mcp, svc, r.Scheme)
	return svc
}

// SECTION: Cert-manager resources

func (r *ManagedControlPlaneReconciler) pkiSelfSignedIssuerForManagedControlPlane(mcp *iaasv1alpha1.ManagedControlPlane) *certmgrv1.Issuer {

	issuer := &certmgrv1.Issuer{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "",
			Namespace: "",
		},
		Spec: certmgrv1.IssuerSpec{
			IssuerConfig: certmgrv1.IssuerConfig{
				SelfSigned: &certmgrv1.SelfSignedIssuer{},
			},
		},
	}

	ctrl.SetControllerReference(mcp, issuer, r.Scheme)
	return issuer
}

func (r *ManagedControlPlaneReconciler) pkiSelfSignedRootCACertForManagedControlPlane(mcp *iaasv1alpha1.ManagedControlPlane) *certmgrv1.Certificate {

	cert := &certmgrv1.Certificate{}

	ctrl.SetControllerReference(mcp, cert, r.Scheme)
	return cert
}

func (r *ManagedControlPlaneReconciler) pkiRootCAIssuerForManagedControlPlane(mcp *iaasv1alpha1.ManagedControlPlane) *certmgrv1.Issuer {

	issuer := &certmgrv1.Issuer{}

	ctrl.SetControllerReference(mcp, issuer, r.Scheme)
	return issuer
}

func (r *ManagedControlPlaneReconciler) pkiSubordinateCAForETCDCluster(mcp *iaasv1alpha1.ManagedControlPlane) *certmgrv1.Issuer {

	issuer := &certmgrv1.Issuer{}

	ctrl.SetControllerReference(mcp, issuer, r.Scheme)
	return issuer
}

// SetupWithManager sets up the controller with the Manager.
func (r *ManagedControlPlaneReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&iaasv1alpha1.ManagedControlPlane{}).
		Owns(&corev1.Namespace{}).
		Owns(&corev1.Service{}).
		Owns(&iaasv1alpha1.ManagedEtcd{}).
		Owns(&iaasv1alpha1.ManagedApiServer{}).
		Owns(&iaasv1alpha1.ManagedScheduler{}).
		Owns(&iaasv1alpha1.ManagedControllerManager{}).
		Owns(&iaasv1alpha1.ManagedCloudControllerManager{}).
		Complete(r)
}

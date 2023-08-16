*   ManagedControlPlane 
    *   Status  ManagedControlPlaneStatus
        *   APIEndpoints    []APIEndpoint
            *   Host    string
            *   Port    int
        *   Phase   ManagedControlPlanePhase
    *   Spec    ManagedControlPlaneSpec
        *   Version string
---
Manages/Dependencies:
*   Service
    *   Type Loadbalancer
    *   Enables HA and Horizontal Scaling
*   Cert-Manager:
    *   Delegates MGMT of TLS for EtcdCluster and ManagedControlPlane 
    *   Setup ManagedControlPlane as an subordinate CA
*   External-DNS
    *   Delegates MGMT of DNS registration for each ManagedControlPlane's APIEndpoints
*   Etcd-Operator
    *   Delegates MGMT of etcd-cluster for each ManagedControlPlane 

---
# Tree
*   ManagedControlPlaneReconciler
    *   For
        *   iaasv1alpha1.ManagedControlPlane
    *   Manages
        *   corev1.Namespace
        *   corev1.Service
        *   iaasv1alpha1.ManagedEtcd
        *   iaasv1alpha1.ManagedApiServer
        *   iaasv1alpha1.ManagedScheduler
        *   iaasv1alpha1.ManagedControllerManager
        *   iaasv1alpha1.ManagedCloudControllerManager
*   ManagedEtcdReconciler
    *   For
        *   iaasv1alpha1.ManagedEtcd
    *   Manages
        *   appsv1.StatefulSet
        *   corev1.Service [Peers]
        *   corev1.Service [Clients]


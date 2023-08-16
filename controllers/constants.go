/*
Copyright 2023 TrooperT.
*/

package controllers

import (
	corev1 "k8s.io/api/core/v1"
)

const (
	AppManagedControlPlane				= "managed-control-plane"
	ComponentEtcdPeer					= "etcd-peer"
	ComponentEtcdClient					= "etcd-client"
	ComponentEtcd						= "etcd-cluster"
	ComponentAPIServer					= "api-server"
	ComponentScheduler					= "scheduler"
	ComponentControllerManager			= "controller-manager"
	ComponentCloudControllerManager		= "cloud-controller-manager"

	ServiceEtcdClient					= "etcd-client"
	ServiceEtcd							= "etcd"
	StatefulSetEtcd						= "etcd"
	EtcdPeerPort						= 2380
	EtcdClientPort						= 2379

	ServiceAPIServerExternal			= "control-plane-external"
	ServiceAPIServerExternalType		= corev1.ServiceTypeLoadBalancer
	ServiceAPIServerInternal			= "control-plane-internal"
	ServiceAPIServerInternalType		= corev1.ServiceTypeLoadBalancer
	DeploymentAPIServer					= "api-server"
	APIServerPort						= 6443
	// APIServer							= 
	// APIServer							= 
	// APIServer							= 
	// APIServer							= 
	// APIServer							= 
	// APIServer							= 
	// APIServer							= 
	// APIServer							= 
	// APIServer							= 
	// APIServer							= 
	// APIServer							= 
	// APIServer							= 
)
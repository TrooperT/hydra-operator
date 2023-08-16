/*
Copyright 2023 TrooperT.
*/

package controllers

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
	
)
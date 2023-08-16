#!/bin/sh

if [ ! -f $HOME/.kube/config ]
then
mkdir -p $HOME/.kube
ln -s $(git rev-parse --show-toplevel)/.private/00-tanzu/kubeconfig ~/.kube/config
else
export KUBECONFIG=$(git rev-parse --show-toplevel)/.private/00-tanzu/kubeconfig
fi
export PATH=$PATH:$(git rev-parse --show-toplevel)/bin:$(git rev-parse --show-toplevel)/.private/bin
source <(kubectl completion bash)
kubectl get node

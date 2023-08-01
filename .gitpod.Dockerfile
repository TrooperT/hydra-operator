ARG WORKSPACE_GO_VERSION
FROM gitpod/workspace-go:${WORKSPACE_GO_VERSION:-latest}

ARG ARCH
ENV ARCH=${ARCH:-amd64}
ENV OS=linux
ARG OPERATOR_SDK_VERSION
ENV OPERATOR_SDK_VERSION=${OPERATOR_SDK_VERSION:-v1.31.0}
ENV OPERATOR_SDK_DL_URL=https://github.com/operator-framework/operator-sdk/releases/download/${OPERATOR_SDK_VERSION}

# VARIABLES
ARG TANZU_SUPERVISOR_IP
ENV TANZU_SUPERVISOR_IP=${TANZU_SUPERVISOR_IP:-172.16.18.201}
ARG HELM_VERSION
ENV HELM_VERSION=${HELM_VERSION:-v3.11.3}
ARG VMCA_IP
ENV VMCA_IP=${VMCA_IP:-172.16.11.249}

RUN mkdir -p /tmp/osdk && curl -L ${OPERATOR_SDK_DL_URL}/operator-sdk_${OS}_${ARCH} -o /tmp/osdk/operator-sdk_${OS}_${ARCH}
RUN curl -L ${OPERATOR_SDK_DL_URL}/checksums.txt -o /tmp/osdk/checksums.txt
RUN curl -L ${OPERATOR_SDK_DL_URL}/checksums.txt.asc -o /tmp/osdk/checksums.txt.asc
RUN gpg --keyserver keyserver.ubuntu.com --recv-keys 052996E2A20B5C7E
RUN cd /tmp/osdk && gpg -u "Operator SDK (release) <cncf-operator-sdk@cncf.io>" --verify checksums.txt.asc
RUN cd /tmp/osdk && grep operator-sdk_${OS}_${ARCH} checksums.txt | sha256sum -c -

# Download/Unpack kubectl vsphere plugin from WCP Supervisor Cluster
RUN curl -Lkv https://$TANZU_SUPERVISOR_IP/wcp/plugin/linux-amd64/vsphere-plugin.zip -o /tmp/vsphere-plugin.zip
RUN mkdir -p /tmp/vsphere-plugin && unzip /tmp/vsphere-plugin.zip -d /tmp/vsphere-plugin
# Download/Unpack helm
RUN curl -Lk https://get.helm.sh/helm-$HELM_VERSION-linux-amd64.tar.gz -o /tmp/helm.tgz
RUN mkdir -p /tmp/helm && tar xvzf /tmp/helm.tgz -C /tmp/helm

# Install k8s mgmt tools
RUN sudo install --verbose --mode=0755 /tmp/vsphere-plugin/bin/kubectl /usr/bin
RUN sudo install --verbose --mode=0755 /tmp/vsphere-plugin/bin/kubectl-vsphere /usr/bin
RUN sudo install --verbose --mode=0755 /tmp/helm/linux-amd64/helm /usr/bin/helm
RUN sudo ln -s /usr/bin/helm /usr/bin/helm3
# Install Operator SDK
RUN sudo install --verbose --mode=0755 /tmp/osdk/operator-sdk_${OS}_${ARCH} /usr/local/bin/operator-sdk

# Download/Unpack/Install vsphere root CA
RUN curl -Lkv https://$VMCA_IP/certs/download.zip -o /tmp/certs.zip
RUN unzip /tmp/certs.zip -d /tmp
RUN sudo cp -av /tmp/certs/lin/5e4fb928.0 /usr/local/share/ca-certificates/5e4fb928.crt
RUN sudo update-ca-certificates

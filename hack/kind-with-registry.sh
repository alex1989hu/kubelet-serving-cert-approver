#!/bin/sh
#
# Adapted from:
# https://github.com/kubernetes-sigs/kind/commits/master/site/static/examples/kind-with-registry.sh
#
# Copyright 2020 The Kubernetes Project
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -o errexit

# desired cluster name; default is "kind"
KIND_CLUSTER_NAME="${KIND_CLUSTER_NAME:-kind}"

kind_version=$(kind version)
kind_network='kind'
reg_name='kind-registry'
reg_port='5000'
case "${kind_version}" in
  "kind v0.7."* | "kind v0.6."* | "kind v0.5."*)
    kind_network='bridge'
    ;;
esac

# create registry container unless it already exists
running="$(docker inspect -f '{{.State.Running}}' "${reg_name}" 2>/dev/null || true)"
if [ "${running}" != 'true' ]; then
  docker run \
    -d --restart=always -p "127.0.0.1:${reg_port}:5000" --name "${reg_name}" \
    registry:2
fi

reg_host="${reg_name}"
if [ "${kind_network}" = "bridge" ]; then
    reg_host="$(docker inspect -f '{{.NetworkSettings.IPAddress}}' "${reg_name}")"
fi
echo "Registry Host: ${reg_host}"

# create a cluster with the local registry enabled in containerd
cat <<EOF | kind create cluster --name "${KIND_CLUSTER_NAME}" --config=-
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
containerdConfigPatches:
- |-
  [plugins."io.containerd.grpc.v1.cri".registry.mirrors."localhost:${reg_port}"]
    endpoint = ["http://${reg_host}:5000"]
nodes:
- role: control-plane
  image: kindest/node:v1.33.1
  kubeadmConfigPatches:
  - |-
    kind: InitConfiguration
    nodeRegistration:
      kubeletExtraArgs:
        "rotate-server-certificates": "true"
- role: worker
  image: kindest/node:v1.33.1
  kubeadmConfigPatches:
  - |-
    kind: JoinConfiguration
    nodeRegistration:
      kubeletExtraArgs:
        "rotate-server-certificates": "true"
- role: worker
  image: kindest/node:v1.30.0
  kubeadmConfigPatches:
  - |-
    kind: JoinConfiguration
    nodeRegistration:
      kubeletExtraArgs:
        "rotate-server-certificates": "true"
kubeadmConfigPatches:
- |-
  kind: ClusterConfiguration
  apiServer:
    extraArgs:
      # Increase duration of Event Time To Leave (TTL)
      "event-ttl": "8h0m0s"
  controllerManager:
    extraArgs:
      # Disable csrcleaner contoller to avoid removal of Certificate Signing Request; keep default KiND options
      "controllers": "*,bootstrapsigner,tokencleaner,-csrcleaner"
EOF

cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: ConfigMap
metadata:
  name: local-registry-hosting
  namespace: kube-public
data:
  localRegistryHosting.v1: |
    host: "localhost:${reg_port}"
    help: "https://kind.sigs.k8s.io/docs/user/local-registry/"
EOF

if [ "${kind_network}" != "bridge" ]; then
  containers=$(docker network inspect ${kind_network} -f "{{range .Containers}}{{.Name}} {{end}}")
  needs_connect="true"
  for c in $containers; do
    if [ "$c" = "${reg_name}" ]; then
      needs_connect="false"
    fi
  done
  if [ "${needs_connect}" = "true" ]; then
    docker network connect "${kind_network}" "${reg_name}" || true
  fi
fi

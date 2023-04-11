#!/bin/bash

# Copyright AppsCode Inc. and Contributors
#
# Licensed under the AppsCode Community License 1.0.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Community-1.0.0.md
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -e

install_crossplane() {
    helm repo add crossplane-stable https://charts.crossplane.io/stable && helm repo update
    helm upgrade -i crossplane \
        crossplane-stable/crossplane \
        --namespace crossplane-system \
        --create-namespace

    kubectl wait --for=condition=ready pods --all -n crossplane-system --timeout=5m

    cat <<EOF | kubectl apply -f -
  apiVersion: pkg.crossplane.io/v1
  kind: Provider
  metadata:
    name: upbound-provider-aws
  spec:
    package: xpkg.upbound.io/upbound/provider-aws:v0.27.0
EOF

    cat >aws-credentials.txt <<EOF
[default]
aws_access_key_id = ${AWS_ACCESS_KEY_ID}
aws_secret_access_key = ${AWS_SECRET_ACCESS_KEY}
EOF

    kubectl create secret generic aws-secret \
        -n crossplane-system \
        --from-file=creds=./aws-credentials.txt

    helm upgrade -i aws-peering ./capa-vpc-peering-operator -n kubedb-managed --create-namespace

    cat <<EOF | kubectl apply -f -
apiVersion: aws.upbound.io/v1beta1
kind: ProviderConfig
metadata:
  name: default
spec:
  credentials:
    source: Secret
    secretRef:
      namespace: crossplane-system
      name: aws-secret
      key: creds
EOF
}

install_crossplane

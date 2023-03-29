#!/bin/bash

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

  cat > aws-credentials.txt <<EOF
[default]
aws_access_key_id = ${AWS_ACCESS_KEY_ID}
aws_secret_access_key = ${AWS_SECRET_ACCESS_KEY}
EOF

  kubectl create secret generic aws-secret \
  -n crossplane-system \
  --from-file=creds=./aws-credentials.txt

  helm upgrade -i aws-peering ./aws-peering-connection-operator -n kubedb-managed --create-namespace
}

install_crossplane

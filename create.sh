#!/bin/bash -ex

if [[ $# -eq 0 ]]; then
    echo usage: $0 resourcegroup
    exit 1
fi

RESOURCEGROUP=$1

set +e
HAS_GOBINDATA=$(command -v go-bindata)
HAS_GOYACC=$(command -v go-yacc)
set -e

# GOPATH/bin needs to be set in $PATH
if [[ -z "${HAS_GOBINDATA:-}" ]]; then
	go get github.com/go-bindata/go-bindata/...
fi
if [[ -z "${HAS_GOYACC:-}" ]]; then
	go get github.com/cznic/goyacc
fi

rm -rf _out
mkdir _out
go generate ./...

go run cmd/create/create.go <<EOF
TenantID: $AZURE_TENANT_ID
SubscriptionID: $AZURE_SUBSCRIPTION_ID
ClientID: $AZURE_CLIENT_ID
ClientSecret: $AZURE_CLIENT_SECRET
Location: eastus
ResourceGroup: $RESOURCEGROUP
VMSize: Standard_D4s_v3
ComputeCount: 1
InfraCount: 1
ImageResourceGroup: images
ImageResourceName: centos7-3.10-201806231427
EOF

helm template pkg/helm/chart -f _out/values.yaml --output-dir _out

# poor man's helm create (without tiller running)
oc delete -Rf _out/osa/templates || true
oc create -Rf _out/osa/templates

az group create -n $RESOURCEGROUP -l eastus
az group deployment create -g $RESOURCEGROUP --template-file _out/azuredeploy.json

# will eventually run as an HCP pod, for development run it locally
KUBECONFIG=_out/admin.kubeconfig go run cmd/sync/sync.go

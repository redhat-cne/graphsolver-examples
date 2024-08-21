#!/bin/bash
set -o nounset -o pipefail

# Creates kind cluster
cat > kind-config.yaml <<EOF
# three node (two workers) cluster config
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
nodes:
- role: control-plane
- role: worker
- role: worker
EOF

go mod vendor
make build 

kind delete cluster --name test-cluster || true
kind create cluster --name test-cluster --config kind-config.yaml

./simplegraphsolver 2>&1|sed 's/veth[a-zA-Z0-9]\+/veth/g'|sed 's/time="[^"]*" //'  > test/actual.txt 

diff test/actual.txt test/expected.txt
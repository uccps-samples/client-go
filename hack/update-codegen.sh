#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

SCRIPT_ROOT=$(dirname ${BASH_SOURCE})/..
CODEGEN_PKG=${CODEGEN_PKG:-$(cd ${SCRIPT_ROOT}; ls -d -1 ./vendor/k8s.io/code-generator 2>/dev/null || echo ../../../k8s.io/code-generator)}

verify="${VERIFY:-}"

set -x
# Because go mod sux, we have to fake the vendor for generator in order to be able to build it...
mv ${CODEGEN_PKG}/generate-groups.sh ${CODEGEN_PKG}/generate-groups.sh.orig
sed 's/go install/#GO111MODULE=on go install/g' ${CODEGEN_PKG}/generate-groups.sh.orig > ${CODEGEN_PKG}/generate-groups.sh
function cleanup {
  mv ${CODEGEN_PKG}/generate-groups.sh.orig ${CODEGEN_PKG}/generate-groups.sh
}
trap cleanup EXIT

go install ./${CODEGEN_PKG}/cmd/{defaulter-gen,client-gen,lister-gen,informer-gen,deepcopy-gen}

for group in apiserver apps authorization build cloudnetwork config image imageregistry network oauth project quota route samples security securityinternal template user; do
  bash ${CODEGEN_PKG}/generate-groups.sh "client,lister,informer" \
    github.com/uccps-samples/client-go/${group} \
    github.com/uccps-samples/api \
    "${group}:v1" \
    --go-header-file ${SCRIPT_ROOT}/hack/boilerplate.txt \
    --plural-exceptions=DNS:DNSes,DNSList:DNSList,Endpoints:Endpoints,Features:Features,FeaturesList:FeaturesList,SecurityContextConstraints:SecurityContextConstraints \
    ${verify}
done

for group in console; do
  bash ${CODEGEN_PKG}/generate-groups.sh "client,lister,informer" \
    github.com/uccps-samples/client-go/${group} \
    github.com/uccps-samples/api \
    "${group}:v1,v1alpha1" \
    --go-header-file ${SCRIPT_ROOT}/hack/boilerplate.txt \
    --plural-exceptions=DNS:DNSes,DNSList:DNSList,Endpoints:Endpoints,Features:Features,FeaturesList:FeaturesList,SecurityContextConstraints:SecurityContextConstraints \
    ${verify}
done

for group in servicecertsigner operatorcontrolplane sharedresource; do
  bash ${CODEGEN_PKG}/generate-groups.sh "client,lister,informer" \
    github.com/uccps-samples/client-go/${group} \
    github.com/uccps-samples/api \
    "${group}:v1alpha1" \
    --go-header-file ${SCRIPT_ROOT}/hack/boilerplate.txt \
    --plural-exceptions=DNS:DNSes,DNSList:DNSList,Endpoints:Endpoints,Features:Features,FeaturesList:FeaturesList,SecurityContextConstraints:SecurityContextConstraints \
    ${verify}
done

for group in operator; do
  bash ${CODEGEN_PKG}/generate-groups.sh "client,lister,informer" \
    github.com/uccps-samples/client-go/${group} \
    github.com/uccps-samples/api \
    "${group}:v1,v1alpha1" \
    --go-header-file ${SCRIPT_ROOT}/hack/boilerplate.txt \
    --plural-exceptions=DNS:DNSes,DNSList:DNSList,Endpoints:Endpoints,Features:Features,FeaturesList:FeaturesList,SecurityContextConstraints:SecurityContextConstraints \
    ${verify}
done

for group in helm machine; do
  bash ${CODEGEN_PKG}/generate-groups.sh "client,lister,informer" \
    github.com/uccps-samples/client-go/${group} \
    github.com/uccps-samples/api \
    "${group}:v1beta1" \
    --go-header-file ${SCRIPT_ROOT}/hack/boilerplate.txt \
    --plural-exceptions=DNS:DNSes,DNSList:DNSList,Endpoints:Endpoints,Features:Features,FeaturesList:FeaturesList,SecurityContextConstraints:SecurityContextConstraints \
    ${verify}
done

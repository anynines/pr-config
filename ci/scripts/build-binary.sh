#!/bin/bash

set -o errexit # Exit immediately if a simple command exits with a non-zero status
set -o nounset # Report the usage of uninitialized variables

version="$(cat version/version)"
binary_dir=${PWD}/${GOOS}-binary

echo -n "Prepare sources..."
mkdir -p /go/src/github.com/anynines
ln -s ${PWD}/pr-config /go/src/github.com/anynines/pr-config
echo "done"

cd /go/src/github.com/anynines/pr-config
echo -n "Build binary for ${GOOS} ${GOARCH}..."
go build \
  -o ${binary_dir}/pr-config-${GOOS}-${GOARCH}-${version} \
  main.go
echo "done"

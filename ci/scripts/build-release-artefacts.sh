#!/bin/bash

set -o errexit # Exit immediately if a simple command exits with a non-zero status
set -o nounset # Report the usage of uninitialized variables

release_artefacts=${PWD}/release-artefacts
release_version=$(cat version/version)

if [ -z "${release_version}" ]; then
  echo "Missing version number" >&2
  exit 1
fi

echo -n "Build release name..."
echo "v${release_version}" > ${release_artefacts}/release-name
echo "done"

echo -n "Build release tag..."
echo "v${release_version}" > ${release_artefacts}/tag
echo "done"

echo -n "Build release commitish..."
git -C pr-config rev-parse HEAD > ${release_artefacts}/commitish
echo "done"

echo -n "Build release notes..."
cat > ${release_artefacts}/notes.md <<EOF
EOF
echo "done"

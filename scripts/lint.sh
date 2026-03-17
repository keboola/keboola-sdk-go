#!/bin/bash

set -o errexit          # Exit on most errors (see the manual)
set -o errtrace         # Make sure any error trap is inherited
set -o nounset          # Disallow expansion of unset variables
set -o pipefail         # Use last non-zero exit code in a pipeline
#set -o xtrace          # Trace the execution of the script (debug)

# Check the most important problems first
echo "Running go vet ..."
if ! go vet $(go list ./pkg/... | grep -v 'pkg/keboola/management'); then
    echo "Please fix ^^^ errors. You can try run \"task fix\"."
    echo
    exit 1
fi

# Check modules
echo "Running go mod tidy/verify ..."
GOWORK=off go mod tidy
git diff --exit-code -- go.mod go.sum
GOWORK=off go mod verify
echo "Ok. go.mod and go.sum are valid."
echo


# Run linters
echo "Running golangci-lint ..."
if golangci-lint run -c "./build/ci/golangci.yml" "$@"; then
    echo "Ok. The code looks good."
    echo
else
    echo "Please fix ^^^ errors. You can try run \"task fix\"."
    echo
    exit 1
fi

# --- upload module ---
echo "Running go vet (upload) ..."
if ! (cd upload && go vet ./...); then
    echo "Please fix ^^^ errors."
    exit 1
fi

echo "Running go mod tidy/verify (upload) ..."
(cd upload && GOWORK=off go mod tidy)
git diff --exit-code -- upload/go.mod upload/go.sum
(cd upload && GOWORK=off go mod verify)
echo "Ok. upload/go.mod and upload/go.sum are valid."

echo "Running golangci-lint (upload) ..."
if (cd upload && golangci-lint run -c "../build/ci/golangci.yml"); then
    echo "Ok. Upload module looks good."
else
    echo "Please fix ^^^ errors. You can try run \"task upload-fix\"."
    exit 1
fi

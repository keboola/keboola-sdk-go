#!/bin/bash

set -o errexit          # Exit on most errors (see the manual)
set -o errtrace         # Make sure any error trap is inherited
set -o nounset          # Disallow expansion of unset variables
set -o pipefail         # Use last non-zero exit code in a pipeline
#set -o xtrace          # Trace the execution of the script (debug)

SCRIPT_DIR="$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"
cd "$SCRIPT_DIR"

# Version to install
GOLANGCI_LINT_VERSION="v2.15.0"

# golangci-lint
# Built from source by the Go toolchain of the project: the module is verified by the Go checksum database, and the
# linter is never built with an older Go than the code it checks (it cannot load packages of a newer Go version).
# The install.sh of golangci-lint does not work since v2.15.0: its checksum lookup matches also the .sbom.json asset
# of the archive and the verification of a correct download fails.
if ! command -v golangci-lint &> /dev/null
then
    go install "github.com/golangci/golangci-lint/v2/cmd/golangci-lint@${GOLANGCI_LINT_VERSION}"
fi

# Verify installation
echo "Verifying golangci-lint installation..."
golangci-lint --version

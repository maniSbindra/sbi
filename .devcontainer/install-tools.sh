#!/usr/bin/env bash
set -euo pipefail

# Installer SHAs — update these when bumping tool versions
SYFT_VERSION="v1.42.3"
SYFT_INSTALLER_SHA="860126c650c2d05b63b83a3895e41268162315a3" # v1.42.3 anchore/syft install.sh

TRIVY_VERSION="v0.69.3"
TRIVY_INSTALLER_SHA="6fb20c8edd70745d6b34bff0387b53b03c8a760a" # v0.69.3 aquasecurity/trivy contrib/install.sh

GOLANGCI_LINT_VERSION="v2.11.4"
GOLANGCI_LINT_INSTALLER_SHA="8f3b0c7ed018e57905fbd873c697e0b1ede605a5" # v2.11.4 golangci/golangci-lint install.sh

echo "=== Installing scanning tools ==="

# Syft (SBOM generator)
echo "Installing Syft..."
curl -sSfL "https://raw.githubusercontent.com/anchore/syft/${SYFT_INSTALLER_SHA}/install.sh" | sudo sh -s -- -b /usr/local/bin "${SYFT_VERSION}"

# Trivy (vulnerability scanner)
echo "Installing Trivy..."
curl -sSfL "https://raw.githubusercontent.com/aquasecurity/trivy/${TRIVY_INSTALLER_SHA}/contrib/install.sh" | sudo sh -s -- -b /usr/local/bin "${TRIVY_VERSION}"

# golangci-lint
echo "Installing golangci-lint..."
curl -sSfL "https://raw.githubusercontent.com/golangci/golangci-lint/${GOLANGCI_LINT_INSTALLER_SHA}/install.sh" | sudo sh -s -- -b /usr/local/bin "${GOLANGCI_LINT_VERSION}"

echo "=== Tools installed ==="
syft version
trivy version
golangci-lint version

#!/usr/bin/env bash
set -euo pipefail

echo "=== Installing scanning tools ==="

# Syft (SBOM generator)
echo "Installing Syft..."
curl -sSfL https://raw.githubusercontent.com/anchore/syft/main/install.sh | sh -s -- -b /usr/local/bin

# Trivy (vulnerability scanner)
echo "Installing Trivy..."
curl -sSfL https://raw.githubusercontent.com/aquasecurity/trivy/main/contrib/install.sh | sh -s -- -b /usr/local/bin

# golangci-lint
echo "Installing golangci-lint..."
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s -- -b /usr/local/bin

echo "=== Tools installed ==="
syft version
trivy version
golangci-lint version

#!/usr/bin/env bash
set -euo pipefail

echo "=== Installing scanning tools ==="

# Syft (SBOM generator)
echo "Installing Syft..."
curl -sSfL https://raw.githubusercontent.com/anchore/syft/860126c650c2d05b63b83a3895e41268162315a3/install.sh | sh -s -- -b /usr/local/bin v1.42.3

# Trivy (vulnerability scanner)
echo "Installing Trivy..."
curl -sSfL https://raw.githubusercontent.com/aquasecurity/trivy/6fb20c8edd70745d6b34bff0387b53b03c8a760a/contrib/install.sh | sh -s -- -b /usr/local/bin v0.69.3

# golangci-lint
echo "Installing golangci-lint..."
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/8f3b0c7ed018e57905fbd873c697e0b1ede605a5/install.sh | sh -s -- -b /usr/local/bin v2.11.4

echo "=== Tools installed ==="
syft version
trivy version
golangci-lint version

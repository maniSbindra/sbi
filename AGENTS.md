# AGENTS.md — LLM Context for SBI

> **SBI** (Secure Base Images) is a CLI tool that performs nightly vulnerability scanning of container base images from MCR (Microsoft Container Registry) and Docker Hub, then generates ranked recommendation reports.

## Project Overview

- **Module:** `github.com/manisbindra/sbi`
- **Language:** Go 1.26, built with `task build`
- **CLI framework:** Cobra (`github.com/spf13/cobra`)
- **Database:** SQLite via `modernc.org/sqlite` (pure Go, no CGO)
- **Binary name:** `daily-recommendations`

## Repository Structure

```text
cmd/                          # Cobra CLI commands (scan, report, reset-db)
pkg/
  domain/                     # Data models (ImageRecord, Language, Vulnerability, etc.)
  infrastructure/
    database/                 # SQLite schema, persistence, ranked queries
      schema.go               # Table definitions (images, languages, vulnerabilities, etc.)
      repository.go           # InsertImage (upsert), QueryTopImages, QueryLanguages
    scanner/
      analyzer.go             # Orchestrates: pull -> inspect -> syft -> trivy -> verify
      syft.go                 # SBOM parsing, language detection from packages
      trivy.go                # Vulnerability scanning
      registry.go             # Tag discovery from container registries
      docker.go               # Docker pull, inspect, run commands in containers
    report/
      markdown.go             # Markdown report generation
      json.go                 # JSON report generation
  usecase/
    pipeline.go               # Top-level orchestration: config -> discover -> scan -> report
config/
  repositories.json           # Image sources for nightly scans
  smoke-test/repositories.json # Minimal config for CI smoke tests
docs/
  daily_recommendations.md    # Generated nightly report (committed to repo)
  daily_recommendations.json  # Generated JSON report
```

## Build, Test, Lint

All commands use [Task](https://taskfile.dev/) (see `Taskfile.yml`).
**Always run `task lint` before creating a PR** — CI runs linting on all PRs and will fail on errors.

```bash
task build          # Build binary to bin/{OS}-{ARCH}/daily-recommendations
task test           # Unit tests with coverage
task test:short     # Fast tests without coverage
task lint           # All linters: golangci-lint, markdown, YAML, govulncheck
task lint:go        # golangci-lint with auto-fix
task lint:md        # Markdown linting (uses markdownlint-cli)
task lint:yaml      # YAML validation
task all            # deps -> lint -> test -> build
task scan           # Build + run full nightly scan
task report         # Generate reports from existing database
task clean          # Remove build/coverage artifacts
```

## CLI Usage

```bash
# Scan images and generate report
daily-recommendations scan \
  --database azure_linux_images.db \
  --config-dir config \
  --output docs/daily_recommendations.md \
  --max-tags 10 \
  --top-n 10 \
  --comprehensive \
  --verbose

# Generate report from existing DB
daily-recommendations report --database azure_linux_images.db

# Reset database
daily-recommendations reset-db --database azure_linux_images.db
```

**Key scan flags:**

- `--max-tags N` — limit tags per repository (0 = all)
- `--comprehensive` — enable secrets + misconfiguration scanning
- `--update-existing` — rescan images already in the database
- `--no-cleanup` — keep Docker images after scanning

## Scan Pipeline Flow

```text
config/repositories.json
  |
  |-- Repository entries (no ':') -> GetTags() from registry API -> FilterTags() -> LimitTags()
  |-- Single image entries (has ':') -> scan directly
  |
  v
For each image:
  1. docker pull
  2. docker inspect          -> size, layers, digest, created date
  3. syft <image> -o json    -> SBOM: languages, packages, capabilities
  4. trivy image <image>     -> vulnerabilities (+ secrets/misconfig if --comprehensive)
  5. mergeLanguages()        -> combine Syft results + image-name detection
  6. verifyRuntimes()        -> run version commands inside container
  7. applyRuntimeVersions()  -> update with verified versions
  8. InsertImage() (upsert)  -> store in SQLite
  |
  v
Generate reports:
  - docs/daily_recommendations.md
  - docs/daily_recommendations.json
```

## Image Configuration Format

`config/repositories.json` supports two entry types:

```json
{
  "defaults": { "registry": "mcr.microsoft.com", "maxTags": 0 },
  "tagFilter": {
    "skipExact": ["latest", "dev", "nightly"],
    "excludeKeywords": ["debug", "test", "experimental", "arm", "amd", "azl"],
    "excludePatterns": ["(?i)[-.]?(alpha|beta|rc|preview)..."],
    "requireDigit": true
  },
  "repositories": [
    {
      "description": "Tag discovery repos",
      "images": ["azurelinux/base/python"]
    },
    {
      "description": "Pinned single images",
      "images": ["docker.io/library/python:3-slim"]
    }
  ]
}
```

- **No colon** -> repository: tags are discovered via `GET /v2/{repo}/tags/list`, then filtered
- **Has colon** -> single image: scanned directly, no tag discovery

## Language Detection (3-stage pipeline)

### Stage 1: SBOM Package Detection (`syft.go`)

Syft scans all packages in the image. Package **names** are matched against regex patterns:

| Language | Regex | Matches |
|----------|-------|---------|
| Python | `(?i)(python&#124;cpython)[\d.-]*$` | `python3`, `cpython` |
| Node.js | `(?i)(node&#124;nodejs)[\d.-]*$` | `nodejs`, `node18` |
| Java | `(?i)(java&#124;openjdk&#124;jdk&#124;jre)[\d.-]*$` | `openjdk-21`, `jre` |
| Go | `(?i)^golang$&#124;^go[\d.]` | `golang`, `go1.21` |
| Ruby | `(?i)^ruby[\d.-]*$` | `ruby3.2` |
| PHP | `(?i)^php[\d.-]*$` | `php8.2` |
| .NET | `(?i)(dotnet&#124;aspnet&#124;\.net)[\d.-]*$` | `dotnet-runtime-8.0` |
| Rust | `(?i)^rust[\d.-]*$` | `rust` |
| Lua | `(?i)^lua[\d.-]*$` | `lua5.4` |

**Excluded packages** (not treated as runtimes): `python-pip`, `python3-pip`, `nodejs-npm`, `java-common`, `python-setuptools`

**Priority system** — when multiple packages match the same language, the highest-priority match wins. This ensures the actual runtime package (e.g., `python3`) is chosen over support libraries (e.g., `python3-libs`) which may report different versions:

- Priority 100: Exact runtime packages (`python3`, `nodejs`, `openjdk-*`, `dotnet-runtime-*`)
- Priority 90: SDK packages (`dotnet-sdk`)
- Priority 50: Other matches
- +5 bonus: RPM/DEB packages over other Syft-detected types

**Version cleaning:** `^(\d+(?:\.\d+)*)` strips build suffixes — `3.12.9-8.azl3` -> `3.12.9`

**Important limitation:** Detection matches against package **name** only, not package **type**. Syft's `dotnet` type packages (e.g., `Microsoft.NETCore.App.Runtime.linux-arm64`) are NOT detected by Stage 1 because the name doesn't match the .NET regex. These are caught by Stage 2.

### Stage 2: Image-Name Detection (`analyzer.go`)

Fallback for .NET distroless images where the runtime is baked in without a matching system package name. Currently only .NET is detected this way:

| Pattern | Matches | Detects |
|---------|---------|---------|
| `(?i)mcr\.microsoft\.com/dotnet/(?:aspnet&#124;runtime):(\d+\.\d+)` | `mcr.microsoft.com/dotnet/aspnet:8.0` | .NET 8.0 |

Only triggers if .NET was **not** already found by Stage 1. This is the primary detection path for .NET Azure Linux distroless images where Syft reports packages with type `dotnet` but names like `Microsoft.NETCore.App.Runtime.linux-arm64` that don't match the name-based regex.

Note: The SDK pattern (`dotnet/sdk:*`) is **not** matched — SDK detection relies on Stage 1 finding a `dotnet-sdk-*` system package.

### Stage 3: Runtime Verification (`analyzer.go`)

For each detected language that has a defined runtime verification command (Python, Node.js, Java, Go, Ruby, PHP, .NET, Rust), runs the actual runtime command **inside the container**:

| Language | Command | Version Pattern |
|----------|---------|-----------------|
| Python | `python3 --version` | `Python\s+(\d+\.\d+\.\d+)` |
| Node.js | `node --version` | `v(\d+\.\d+\.\d+)` |
| Java | `java -version` | `version\s+"?(\d+[\d.]*)` |
| Go | `go version` | `go(\d+\.\d+[\d.]*)` |
| Ruby | `ruby --version` | `ruby\s+(\d+\.\d+\.\d+)` |
| PHP | `php --version` | `PHP\s+(\d+\.\d+\.\d+)` |
| .NET | `dotnet --info` | `Version:\s+(\d+\.\d+[\d.]*)` |
| Rust | `rustc --version` | `rustc\s+(\d+\.\d+\.\d+)` |

Runtime verification overrides the Syft/image-name version with the precise version and sets `Verified=true`. Note: runtime verification only runs for languages already detected by Stage 1 or 2 — it does not independently discover new languages.

## Known Edge Cases

1. **Go build images** — Go is tarball-installed, not RPM-packaged. Syft won't detect `golang` as a system package. Go detection from image names is not yet implemented (planned for a future PR).

2. **JDK images include Python** — Azure Linux JDK images ship `python3` as a system dependency. The image will appear in **both** Java and Python rankings. This is intentional — Python CVEs in the image are real security concerns.

3. **Go modules in non-Go images** — Syft detects `go-module` entries (compiled Go binaries embedded in images). These are **not** treated as "Go runtime" — the language regex only matches system package names (`golang`, `go1.x`), not `go-module` type entries.

4. **.NET distroless images** — Syft reports .NET packages with type `dotnet` but names like `Microsoft.NETCore.App.Runtime.linux-arm64` which don't match the name-based regex in Stage 1. The `dotnetImagePattern` in Stage 2 detects .NET from the image name (`dotnet/aspnet:8.0`) as a fallback, then `dotnet --info` in Stage 3 verifies the exact patch version.

5. **.NET SDK images** — The SDK image name pattern (`dotnet/sdk:10.0-azurelinux3.0`) is **not** matched by `dotnetImagePattern` which only looks for `aspnet` or `runtime`. SDK detection relies on Stage 1 finding a `dotnet-sdk-*` package.

## Report Ranking

Images are ranked **per language** using this sort order (ascending = fewer/smaller is better):

1. **Critical vulnerabilities** (ascending)
2. **High vulnerabilities** (ascending)
3. **Total vulnerabilities** (ascending)
4. **Image size in bytes** (ascending)

The report shows the top N images per language (default: 10) with columns: Rank, Image, Version, Critical, High, Total, Size, Digest, Pinned Reference.

## Database Schema

| Table | Purpose |
|-------|---------|
| `images` | Image metadata, vulnerability counts, scan timestamps |
| `languages` | Detected languages per image (language, version, major_minor, verified) |
| `vulnerabilities` | Individual CVEs (severity, CVSS score, fixed_version) |
| `system_packages` | OS-level packages (RPM/DEB/APK) |
| `package_managers` | Detected package managers (pip, npm, etc.) |
| `capabilities` | Container capabilities |
| `security_findings` | Secrets and misconfigurations (from --comprehensive) |

`InsertImage()` performs an **upsert** (`ON CONFLICT(name) DO UPDATE`) and clears related child tables on re-scan.

## CI/CD Workflows

| Workflow | Trigger | Purpose |
|----------|---------|---------|
| `ci.yml` | Push/PR to main | Build + unit tests + govulncheck |
| `lint.yml` | Push/PR to main | golangci-lint, markdown lint, YAML lint |
| `nightly-scan.yml` | Daily 02:00 UTC | Full scan -> commit reports to main |
| `scan-smoke-test.yml` | PR to main | Minimal scan to validate pipeline |

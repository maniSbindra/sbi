# Daily Recommendations by Language Report

A Go CLI tool that scans container base images from registries (MCR, Docker Hub, etc.), analyzes them for vulnerabilities and language runtimes, and generates a daily markdown report ranking the most secure images per programming language.

## Overview

This project replicates and extends the nightly recommendations pipeline from [secure-container-base-image-recommender](https://github.com/maniSbindra/secure-container-base-image-recommender) (Python) in Go.

### Pipeline

1. **Discover** — Enumerate image tags from configured container registries (MCR API)
2. **Pull & Analyze** — Pull images, generate SBOM with [Syft](https://github.com/anchore/syft), detect languages and packages
3. **Scan** — Run [Trivy](https://github.com/aquasecurity/trivy) vulnerability scanning (with optional secrets/misconfig detection)
4. **Verify** — Runtime verification of detected languages inside containers
5. **Store** — Persist all results in a SQLite database
6. **Report** — Generate a ranked markdown report per language (critical → high → total vulns → size)

## Prerequisites

- **Go** 1.22+
- **Docker** (for image pulling and analysis)
- **Syft** ([install](https://github.com/anchore/syft#installation))
- **Trivy** ([install](https://github.com/aquasecurity/trivy#get-trivy))

## Quick Start

```bash
# Build
go build -o bin/daily-recommendations ./cmd

# Scan all configured repositories and generate report
./bin/daily-recommendations scan --verbose

# Generate report from existing database
./bin/daily-recommendations report

# Clear the database
./bin/daily-recommendations reset-db
```

## CLI Commands

### `scan`

Scans configured registries, analyzes images, and generates the report.

```bash
daily-recommendations scan [flags]
```

| Flag | Default | Description |
| ---- | ------- | ----------- |
| `--max-tags` | 5 | Maximum tags per repository (0 = all) |
| `--comprehensive` | false | Enable comprehensive scanning (secrets + misconfigs) |
| `--no-cleanup` | false | Keep Docker images after scanning |
| `--update-existing` | false | Rescan images already in the database |

### `report`

Generates the markdown report from existing database data.

```bash
daily-recommendations report [flags]
```

### `reset-db`

Clears all data from the database.

```bash
daily-recommendations reset-db
```

### Global Flags

| Flag | Default | Description |
| ---- | ------- | ----------- |
| `--database` | `azure_linux_images.db` | Path to SQLite database |
| `--config-dir` | `config` | Path to configuration directory |
| `--output` | `docs/daily_recommendations.md` | Path to output report |
| `--top-n` | 10 | Number of top images per language |
| `--verbose, -v` | false | Enable verbose output |
| `--debug, -d` | false | Enable debug output |

## Configuration

Image sources are defined in `config/repositories.txt`:

```text
# MCR repository (all tags enumerated)
azurelinux/base/python

# Single image reference (exact tag)
docker.io/library/python:3.12-slim
mcr.microsoft.com/dotnet/aspnet:8.0
```

## Development

Requires [Task](https://taskfile.dev/) for build automation:

```bash
# Install dependencies
task deps

# Run linter
task lint

# Run tests
task test

# Build
task build

# Run everything
task all
```

## Project Structure

```text
cmd/                           # CLI entry point and cobra commands
pkg/
  domain/                      # Domain models (ImageRecord, Language, etc.)
  infrastructure/
    database/                  # SQLite schema and repository
    scanner/                   # Registry, Docker, Syft, Trivy integration
    report/                    # Markdown report generation
  usecase/                     # Pipeline orchestration
config/                        # Configuration files
docs/                          # Generated reports
```

## License

MIT

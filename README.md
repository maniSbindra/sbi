# SBI — Secure Base Image Recommendations

Every night, this project scans container base images for vulnerabilities and generates ranked security recommendations per programming language.

## 📊 Daily Reports

| Format   | Link                                                               |
|----------|--------------------------------------------------------------------|
| Markdown | [docs/daily_recommendations.md](docs/daily_recommendations.md)     |
| JSON     | [docs/daily_recommendations.json](docs/daily_recommendations.json) |

Reports are regenerated nightly at 02:00 UTC via [GitHub Actions](.github/workflows/nightly-scan.yml) and committed automatically. Images are ranked per language by: fewest critical → fewest high → fewest total vulnerabilities → smallest size.

## How It Works

A nightly [GitHub Actions workflow](.github/workflows/nightly-scan.yml) runs the full pipeline:

1. **Discover** — Enumerate image tags from configured container registries (MCR, Docker Hub)
2. **Pull & Analyze** — Pull images, generate SBOM with [Syft](https://github.com/anchore/syft), detect language runtimes
3. **Scan** — Run [Trivy](https://github.com/aquasecurity/trivy) vulnerability scanning
4. **Verify** — Runtime verification of detected languages inside containers
5. **Store** — Persist results in a SQLite database (tracked via Git LFS)
6. **Report** — Generate ranked markdown and JSON reports, commit and push to this repo

### What Gets Scanned

Image sources and tag filtering rules are configured in [`config/repositories.json`](config/repositories.json). Currently scans Azure Linux base/distroless images, Docker Hub official images, .NET, and OpenJDK images.

## Running Locally

### Prerequisites

- **Go** 1.26+, **Docker**, **[Syft](https://github.com/anchore/syft#installation)**, **[Trivy](https://github.com/aquasecurity/trivy#get-trivy)**

### Quick Start

```bash
# Build
task build

# Scan all configured repositories and generate reports
./bin/daily-recommendations scan --verbose

# Regenerate reports from existing database
./bin/daily-recommendations report

# Clear the database
./bin/daily-recommendations reset-db
```

### CLI Flags

| Flag | Default | Description |
| ---- | ------- | ----------- |
| `--database` | `azure_linux_images.db` | Path to SQLite database |
| `--config-dir` | `config` | Path to configuration directory |
| `--output` | `docs/daily_recommendations.md` | Path to output report |
| `--top-n` | 10 | Number of top images per language |
| `--max-tags` | 5 | Maximum tags per repository (0 = all) |
| `--comprehensive` | false | Enable secrets + misconfig scanning |
| `--update-existing` | false | Rescan images already in the database |
| `--verbose, -v` | false | Enable verbose output |

## Configuration

Image sources and tag filtering rules are defined in [`config/repositories.json`](config/repositories.json):

```json
{
  "defaults": {
    "registry": "mcr.microsoft.com",
    "maxTags": 0
  },
  "tagFilter": {
    "skipExact": ["latest", "dev", "nightly", "edge"],
    "excludeKeywords": ["debug", "test", "arm", "amd"],
    "excludePatterns": ["(?i)[-.]?(alpha|beta|rc|preview)[\\d.]*$"],
    "requireDigit": true
  },
  "repositories": [
    {
      "description": "Azure Linux base images",
      "images": ["azurelinux/base/python", "azurelinux/base/nodejs"]
    }
  ]
}
```

## Development

Requires [Task](https://taskfile.dev/) for build automation:

```bash
task build        # Build binary
task test         # Run tests
task lint         # Run all linters (go, markdown, yaml)
task vulncheck    # Run Go vulnerability check
task all          # Build + test + lint
```

## Project Structure

```text
cmd/                           # CLI entry point and cobra commands
pkg/
  domain/                      # Domain models (ImageRecord, Language, etc.)
  infrastructure/
    database/                  # SQLite schema and repository
    scanner/                   # Registry, Docker, Syft, Trivy integration
    report/                    # Markdown and JSON report generation
  usecase/                     # Pipeline orchestration
config/                        # Image sources and tag filter config
docs/                          # Generated daily reports
```

## Background

This project replicates and extends the nightly recommendations pipeline from [secure-container-base-image-recommender](https://github.com/maniSbindra/secure-container-base-image-recommender) (Python) in Go.

## License

MIT

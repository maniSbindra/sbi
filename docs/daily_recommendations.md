# Daily Recommended Images by Language

_Generated: 2026-02-24T10:32:16Z. Criteria: lowest critical → high → total vulnerabilities → size. Top 10 per language._

**Note:** Image sizes are based on Linux amd64 platform as reported by `docker images` on GitHub runners. Actual sizes may vary significantly on other platforms (macOS, Windows, etc.).

## Scanned Repositories and Images

This report includes analysis from **28 configured sources** across 5 groups (see [repositories.json](../config/repositories.json)):

**Azure Linux base images:**

- `azurelinux/base/core`
- `azurelinux/base/python`
- `azurelinux/base/nodejs`

**Azure Linux distroless images:**

- `azurelinux/distroless/base`
- `azurelinux/distroless/minimal`
- `azurelinux/distroless/python`
- `azurelinux/distroless/nodejs`

**Docker Hub images (for comparison):**

- `docker.io/library/python:3-slim`
- `docker.io/library/python:3.12-slim`
- `docker.io/library/node:20.0-slim`
- `docker.io/library/node:22-bookworm-slim`

**.NET images:**

- `mcr.microsoft.com/dotnet/aspnet:8.0`
- `mcr.microsoft.com/dotnet/runtime:8.0`
- `mcr.microsoft.com/dotnet/sdk:8.0`
- `mcr.microsoft.com/dotnet/aspnet:10.0-azurelinux3.0-distroless`
- `mcr.microsoft.com/dotnet/aspnet:9.0-azurelinux3.0-distroless`
- `mcr.microsoft.com/dotnet/aspnet:8.0-azurelinux3.0-distroless`
- `mcr.microsoft.com/dotnet/runtime:10.0-azurelinux3.0-distroless`
- `mcr.microsoft.com/dotnet/runtime:9.0-azurelinux3.0-distroless`
- `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0`
- `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0`

**OpenJDK images:**

- `mcr.microsoft.com/openjdk/jdk:25-azurelinux`
- `mcr.microsoft.com/openjdk/jdk:25-distroless`
- `mcr.microsoft.com/openjdk/jdk:21-azurelinux`
- `mcr.microsoft.com/openjdk/jdk:21-distroless`
- `mcr.microsoft.com/openjdk/jdk:21-ubuntu`
- `mcr.microsoft.com/openjdk/jdk:17-distroless`
- `mcr.microsoft.com/openjdk/jdk:11-distroless`

## Dotnet

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:9.0-azurelinux3.0-distroless` | 9.0.13 | 0 | 0 | 0 | 105.0 MB | `sha256:3995507d446e` | `mcr.microsoft.com/dotnet/runtime:9.0-azurelinux3.0-distroless@sha256:3995507d446ee3ce1fd278962e7a6749495811edce2519ecee31274030113afa` |
| 2 | `mcr.microsoft.com/dotnet/runtime:10.0-azurelinux3.0-distroless` | 10.0.3 | 0 | 0 | 0 | 110.0 MB | `sha256:9d28d4f5ccb4` | `mcr.microsoft.com/dotnet/runtime:10.0-azurelinux3.0-distroless@sha256:9d28d4f5ccb4f54ceddfb54597a2bb991a93a07f3e2f6a61845e3e98ec040892` |
| 3 | `mcr.microsoft.com/dotnet/aspnet:8.0-azurelinux3.0-distroless` | 8.0.24 | 0 | 0 | 0 | 124.0 MB | `sha256:e91f943b47d5` | `mcr.microsoft.com/dotnet/aspnet:8.0-azurelinux3.0-distroless@sha256:e91f943b47d5c2e2d6ac7cdd979e414c82aeb11980c8ef52e409fcf2afa63913` |
| 4 | `mcr.microsoft.com/dotnet/aspnet:9.0-azurelinux3.0-distroless` | 9.0.13 | 0 | 0 | 0 | 130.0 MB | `sha256:9f856e8de88e` | `mcr.microsoft.com/dotnet/aspnet:9.0-azurelinux3.0-distroless@sha256:9f856e8de88e2ca03687bca28e8575f2e3255ea43df1896ffbd76023e7b07f41` |
| 5 | `mcr.microsoft.com/dotnet/aspnet:10.0-azurelinux3.0-distroless` | 10.0.3 | 0 | 0 | 0 | 137.0 MB | `sha256:973ac891bc21` | `mcr.microsoft.com/dotnet/aspnet:10.0-azurelinux3.0-distroless@sha256:973ac891bc21916cb4f579ed3cd5737fac0a1452d30b11a25493df65eefd4786` |
| 6 | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0` | 9.0.311 | 0 | 0 | 0 | 879.0 MB | `sha256:ef9d3354c106` | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0@sha256:ef9d3354c106e4efa0946c79dca08038c28cd03da2eef678d6b9a640028218fe` |
| 7 | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0` | 10.0.103 | 0 | 0 | 0 | 928.0 MB | `sha256:774b85a4de4b` | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0@sha256:774b85a4de4b63b5974d0a599e7166b0ff20f3848d3cc67333beb733405d027b` |
| 8 | `mcr.microsoft.com/dotnet/runtime:8.0` | 8.0.24 | 1 | 2 | 88 | 193.0 MB | `sha256:b969aeab7f6b` | `mcr.microsoft.com/dotnet/runtime:8.0@sha256:b969aeab7f6b14fcf00f55f420d730df228a93b4731402ca2f8eeb41da5d39c8` |
| 9 | `mcr.microsoft.com/dotnet/aspnet:8.0` | 8.0.24 | 1 | 2 | 88 | 218.0 MB | `sha256:0d6e2e245f18` | `mcr.microsoft.com/dotnet/aspnet:8.0@sha256:0d6e2e245f180ef785f51aab639c8d5d29afc3b43b95c0ee6dfaf5b84895cd6a` |
| 10 | `mcr.microsoft.com/dotnet/sdk:8.0` | 8.0.418 | 1 | 15 | 170 | 850.0 MB | `sha256:58359d0b8fe8` | `mcr.microsoft.com/dotnet/sdk:8.0@sha256:58359d0b8fe8237be1d63ac335ca378e2857f976c7f92791f7c84b3c117037f5` |

## Java

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:21-azurelinux` | 21.0.10 | 0 | 0 | 0 | 482.0 MB | `sha256:25e3fc34980e` | `mcr.microsoft.com/openjdk/jdk:21-azurelinux@sha256:25e3fc34980ec368bd35b82fa1ef178378b829e6ba56c7917705227548494ec3` |
| 2 | `mcr.microsoft.com/openjdk/jdk:25-azurelinux` | 25.0.2 | 0 | 0 | 0 | 527.0 MB | `sha256:0fc3d50b7d5f` | `mcr.microsoft.com/openjdk/jdk:25-azurelinux@sha256:0fc3d50b7d5f45204add224062115a357c807faa510354e98c5698af61eedc2a` |
| 3 | `mcr.microsoft.com/openjdk/jdk:11-distroless` | 11.0.30 | 1 | 0 | 1 | 320.0 MB | `sha256:5731cde41949` | `mcr.microsoft.com/openjdk/jdk:11-distroless@sha256:5731cde419498b0efb3bd9df567c6bbc9add7d84347a0cbe3e2984d683802205` |
| 4 | `mcr.microsoft.com/openjdk/jdk:17-distroless` | 17.0.18 | 1 | 0 | 1 | 323.0 MB | `sha256:70a054029a40` | `mcr.microsoft.com/openjdk/jdk:17-distroless@sha256:70a054029a4096c428fd92f85769a9212a1e118b322b0d1893728a810c4a5a9d` |
| 5 | `mcr.microsoft.com/openjdk/jdk:21-distroless` | 21.0.10 | 1 | 0 | 1 | 351.0 MB | `sha256:c13850904419` | `mcr.microsoft.com/openjdk/jdk:21-distroless@sha256:c138509044194768cdb020dcd7f83ca08419e02267ccaafa72c57f1d8570e869` |
| 6 | `mcr.microsoft.com/openjdk/jdk:25-distroless` | 25.0.2 | 1 | 0 | 1 | 396.0 MB | `sha256:cc6e9d3c8852` | `mcr.microsoft.com/openjdk/jdk:25-distroless@sha256:cc6e9d3c8852cba4f5862186b684787582d9f3dd6c6d547352c0dfc542789d90` |
| 7 | `mcr.microsoft.com/openjdk/jdk:21-ubuntu` | 21.0.10 | 1 | 0 | 64 | 478.0 MB | `sha256:95786a5e9219` | `mcr.microsoft.com/openjdk/jdk:21-ubuntu@sha256:95786a5e921982d10b1a03352e049daaeab63d83bf0ebf1e3049cf327b091392` |

## Node

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/base/nodejs:20.14` | 20.14.0 | 0 | 0 | 0 | 146.0 MB | `sha256:59a1e9373259` | `mcr.microsoft.com/azurelinux/base/nodejs:20.14@sha256:59a1e937325941dbd3bb722be4cef2f2b4b78bb8befa089c8b4915fd389d8662` |
| 2 | `mcr.microsoft.com/azurelinux/base/nodejs:20` | 20.14.0 | 0 | 0 | 0 | 146.0 MB | `sha256:59a1e9373259` | `mcr.microsoft.com/azurelinux/base/nodejs:20@sha256:59a1e937325941dbd3bb722be4cef2f2b4b78bb8befa089c8b4915fd389d8662` |
| 3 | `mcr.microsoft.com/azurelinux/distroless/nodejs:20.14-nonroot` | 20.14.0 | 0 | 7 | 9 | 106.0 MB | `sha256:1dccae1ebad2` | `mcr.microsoft.com/azurelinux/distroless/nodejs:20.14-nonroot@sha256:1dccae1ebad222f36adc1207edfd8b0da2c06a71a6ccc7ceced54516f4af349a` |
| 4 | `mcr.microsoft.com/azurelinux/distroless/nodejs:20.14` | 20.14.0 | 0 | 7 | 9 | 106.0 MB | `sha256:6643a8667737` | `mcr.microsoft.com/azurelinux/distroless/nodejs:20.14@sha256:6643a866773721bea665e29606618a55d28d42fdefb83e38b2e462fbb763a356` |
| 5 | `mcr.microsoft.com/azurelinux/distroless/nodejs:20-nonroot` | 20.14.0 | 0 | 7 | 9 | 106.0 MB | `sha256:1dccae1ebad2` | `mcr.microsoft.com/azurelinux/distroless/nodejs:20-nonroot@sha256:1dccae1ebad222f36adc1207edfd8b0da2c06a71a6ccc7ceced54516f4af349a` |
| 6 | `mcr.microsoft.com/azurelinux/distroless/nodejs:20` | 20.14.0 | 0 | 7 | 9 | 106.0 MB | `sha256:6643a8667737` | `mcr.microsoft.com/azurelinux/distroless/nodejs:20@sha256:6643a866773721bea665e29606618a55d28d42fdefb83e38b2e462fbb763a356` |
| 7 | `docker.io/library/node:22-bookworm-slim` | 22.22.0 | 1 | 16 | 101 | 216.2 MB | `sha256:5373f1906319` | `docker.io/library/node:22-bookworm-slim@sha256:5373f1906319b3a1f291da5d102f4ce5c77ccbe29eb637f072b6c7b70443fc36` |
| 8 | `docker.io/library/node:20.0-slim` | 20.0.0 | 6 | 54 | 252 | 238.7 MB | `sha256:702d475af4b8` | `docker.io/library/node:20.0-slim@sha256:702d475af4b8045f701afd10ea865f4454e9aab4fe3e92d50e8f36906055f456` |

## Python

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot` | 3.12.9 | 0 | 0 | 0 | 83.6 MB | `sha256:66bdc517a7fe` | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot@sha256:66bdc517a7fe1fefa119516a7874881d763d72d234697c5f6702ba18d2d72cff` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/python:3.12` | 3.12.9 | 0 | 0 | 0 | 83.6 MB | `sha256:358a31b2fed4` | `mcr.microsoft.com/azurelinux/distroless/python:3.12@sha256:358a31b2fed412c12d864b69f58374653bd1fc6c7140ade660d60d4acbb089c3` |
| 3 | `mcr.microsoft.com/azurelinux/distroless/python:3-nonroot` | 3.12.9 | 0 | 0 | 0 | 83.6 MB | `sha256:66bdc517a7fe` | `mcr.microsoft.com/azurelinux/distroless/python:3-nonroot@sha256:66bdc517a7fe1fefa119516a7874881d763d72d234697c5f6702ba18d2d72cff` |
| 4 | `mcr.microsoft.com/azurelinux/distroless/python:3` | 3.12.9 | 0 | 0 | 0 | 83.6 MB | `sha256:358a31b2fed4` | `mcr.microsoft.com/azurelinux/distroless/python:3@sha256:358a31b2fed412c12d864b69f58374653bd1fc6c7140ade660d60d4acbb089c3` |
| 5 | `mcr.microsoft.com/azurelinux/base/python:3.12` | 3.12.9 | 0 | 0 | 0 | 139.0 MB | `sha256:7be8b46a4dfa` | `mcr.microsoft.com/azurelinux/base/python:3.12@sha256:7be8b46a4dfad5acc5851ecf76ea800cd2818eaec40832baf5787849645461fd` |
| 6 | `mcr.microsoft.com/azurelinux/base/python:3` | 3.12.9 | 0 | 0 | 0 | 139.0 MB | `sha256:7be8b46a4dfa` | `mcr.microsoft.com/azurelinux/base/python:3@sha256:7be8b46a4dfad5acc5851ecf76ea800cd2818eaec40832baf5787849645461fd` |
| 7 | `mcr.microsoft.com/openjdk/jdk:21-azurelinux` | 3.12.9 | 0 | 0 | 0 | 482.0 MB | `sha256:25e3fc34980e` | `mcr.microsoft.com/openjdk/jdk:21-azurelinux@sha256:25e3fc34980ec368bd35b82fa1ef178378b829e6ba56c7917705227548494ec3` |
| 8 | `mcr.microsoft.com/openjdk/jdk:25-azurelinux` | 3.12.9 | 0 | 0 | 0 | 527.0 MB | `sha256:0fc3d50b7d5f` | `mcr.microsoft.com/openjdk/jdk:25-azurelinux@sha256:0fc3d50b7d5f45204add224062115a357c807faa510354e98c5698af61eedc2a` |
| 9 | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0` | 3.12.9 | 0 | 0 | 0 | 879.0 MB | `sha256:ef9d3354c106` | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0@sha256:ef9d3354c106e4efa0946c79dca08038c28cd03da2eef678d6b9a640028218fe` |
| 10 | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0` | 3.12.9 | 0 | 0 | 0 | 928.0 MB | `sha256:774b85a4de4b` | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0@sha256:774b85a4de4b63b5974d0a599e7166b0ff20f3848d3cc67333beb733405d027b` |


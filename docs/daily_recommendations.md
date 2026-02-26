# Daily Recommended Images by Language

_Generated: 2026-02-26T04:33:17Z. Criteria: lowest critical → high → total vulnerabilities → size. Top 10 per language._

**Note:** Image sizes are based on Linux amd64 platform as reported by `docker images` on GitHub runners. Actual sizes may vary significantly on other platforms (macOS, Windows, etc.).

## Scanned Repositories and Images

This report includes analysis from **30 configured sources** across 6 groups (see [repositories.json](../config/repositories.json)):

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

**Go images:**

- `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0`
- `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0`

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
| 8 | `mcr.microsoft.com/dotnet/runtime:8.0` | 8.0.24 | 1 | 2 | 86 | 193.0 MB | `sha256:d304745fcb83` | `mcr.microsoft.com/dotnet/runtime:8.0@sha256:d304745fcb833f4dd78a0be39812f594722c598303cf0ed3feefd1cbb5d8cc48` |
| 9 | `mcr.microsoft.com/dotnet/aspnet:8.0` | 8.0.24 | 1 | 2 | 86 | 218.0 MB | `sha256:88c86e5469c2` | `mcr.microsoft.com/dotnet/aspnet:8.0@sha256:88c86e5469c2b69969e07343193fd3dbe4ff6936d95b994dbfa27348299af111` |
| 10 | `mcr.microsoft.com/dotnet/sdk:8.0` | 8.0.418 | 1 | 15 | 168 | 850.0 MB | `sha256:bfb6ed602caa` | `mcr.microsoft.com/dotnet/sdk:8.0@sha256:bfb6ed602caa605141700aea7dc7d42574b74b704368e67d683c71a002123808` |

## Go

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0` | 1.25.7 | 0 | 0 | 0 | 809.0 MB | `sha256:6e92b12751fd` | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0@sha256:6e92b12751fdc385f6c415e0a0fe6f630c59df631281deb295f465035320cd8e` |
| 2 | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0` | 1.26.0 | 0 | 0 | 0 | 841.0 MB | `sha256:7df9ca08df20` | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0@sha256:7df9ca08df200e793e1bd35ff92128c965c5cf43815b7beb78278b7dcabb9da3` |

## Java

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:11-distroless` | 11.0.30 | 0 | 0 | 0 | 320.0 MB | `sha256:529ff38afc06` | `mcr.microsoft.com/openjdk/jdk:11-distroless@sha256:529ff38afc063b5cd6e179a38c025d5281d17634fc46fb71f2c1d4a3445b7b31` |
| 2 | `mcr.microsoft.com/openjdk/jdk:17-distroless` | 17.0.18 | 0 | 0 | 0 | 323.0 MB | `sha256:8d45f659fb57` | `mcr.microsoft.com/openjdk/jdk:17-distroless@sha256:8d45f659fb5769c25bbd26f954b0c726a2dc0815f890e99cd13f185a13bfb149` |
| 3 | `mcr.microsoft.com/openjdk/jdk:21-distroless` | 21.0.10 | 0 | 0 | 0 | 351.0 MB | `sha256:8966499f3a21` | `mcr.microsoft.com/openjdk/jdk:21-distroless@sha256:8966499f3a211b0e21bf8f094ee471c946369d827394f9bdd8dd6686d666f5aa` |
| 4 | `mcr.microsoft.com/openjdk/jdk:25-distroless` | 25.0.2 | 0 | 0 | 0 | 396.0 MB | `sha256:e64831f7a293` | `mcr.microsoft.com/openjdk/jdk:25-distroless@sha256:e64831f7a293eb69e703284610d75729df2419c1683fe1678ad768c16b55580c` |
| 5 | `mcr.microsoft.com/openjdk/jdk:21-azurelinux` | 21.0.10 | 0 | 0 | 0 | 482.0 MB | `sha256:4c48d44c4450` | `mcr.microsoft.com/openjdk/jdk:21-azurelinux@sha256:4c48d44c4450a27482974bd839e1593405f41cdfd263d3e4f9647e7bb8097372` |
| 6 | `mcr.microsoft.com/openjdk/jdk:25-azurelinux` | 25.0.2 | 0 | 0 | 0 | 527.0 MB | `sha256:e30b2d29bf11` | `mcr.microsoft.com/openjdk/jdk:25-azurelinux@sha256:e30b2d29bf11dca58cfcb4ecdd823e42af6b3e226f2d5fbf4eef60e11223ce13` |
| 7 | `mcr.microsoft.com/openjdk/jdk:21-ubuntu` | 21.0.10 | 0 | 0 | 63 | 457.0 MB | `sha256:8418bd6a661d` | `mcr.microsoft.com/openjdk/jdk:21-ubuntu@sha256:8418bd6a661de40d996fe45df202168fc7392b11651864459fe6442aff9b57cd` |

## Node

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/base/nodejs:20.14` | 20.14.0 | 0 | 0 | 0 | 146.0 MB | `sha256:59a1e9373259` | `mcr.microsoft.com/azurelinux/base/nodejs:20.14@sha256:59a1e937325941dbd3bb722be4cef2f2b4b78bb8befa089c8b4915fd389d8662` |
| 2 | `mcr.microsoft.com/azurelinux/base/nodejs:20` | 20.14.0 | 0 | 0 | 0 | 146.0 MB | `sha256:59a1e9373259` | `mcr.microsoft.com/azurelinux/base/nodejs:20@sha256:59a1e937325941dbd3bb722be4cef2f2b4b78bb8befa089c8b4915fd389d8662` |
| 3 | `mcr.microsoft.com/azurelinux/distroless/nodejs:20.14-nonroot` | 20.14.0 | 0 | 7 | 9 | 106.0 MB | `sha256:1dccae1ebad2` | `mcr.microsoft.com/azurelinux/distroless/nodejs:20.14-nonroot@sha256:1dccae1ebad222f36adc1207edfd8b0da2c06a71a6ccc7ceced54516f4af349a` |
| 4 | `mcr.microsoft.com/azurelinux/distroless/nodejs:20.14` | 20.14.0 | 0 | 7 | 9 | 106.0 MB | `sha256:6643a8667737` | `mcr.microsoft.com/azurelinux/distroless/nodejs:20.14@sha256:6643a866773721bea665e29606618a55d28d42fdefb83e38b2e462fbb763a356` |
| 5 | `mcr.microsoft.com/azurelinux/distroless/nodejs:20-nonroot` | 20.14.0 | 0 | 7 | 9 | 106.0 MB | `sha256:1dccae1ebad2` | `mcr.microsoft.com/azurelinux/distroless/nodejs:20-nonroot@sha256:1dccae1ebad222f36adc1207edfd8b0da2c06a71a6ccc7ceced54516f4af349a` |
| 6 | `mcr.microsoft.com/azurelinux/distroless/nodejs:20` | 20.14.0 | 0 | 7 | 9 | 106.0 MB | `sha256:6643a8667737` | `mcr.microsoft.com/azurelinux/distroless/nodejs:20@sha256:6643a866773721bea665e29606618a55d28d42fdefb83e38b2e462fbb763a356` |
| 7 | `docker.io/library/node:22-bookworm-slim` | 22.22.0 | 1 | 16 | 99 | 216.2 MB | `sha256:dd9d21971ec4` | `docker.io/library/node:22-bookworm-slim@sha256:dd9d21971ec4395903fa6143c2b9267d048ae01ca6d3ea96f16cb30df6187d94` |
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
| 7 | `mcr.microsoft.com/openjdk/jdk:21-azurelinux` | 3.12.9 | 0 | 0 | 0 | 482.0 MB | `sha256:4c48d44c4450` | `mcr.microsoft.com/openjdk/jdk:21-azurelinux@sha256:4c48d44c4450a27482974bd839e1593405f41cdfd263d3e4f9647e7bb8097372` |
| 8 | `mcr.microsoft.com/openjdk/jdk:25-azurelinux` | 3.12.9 | 0 | 0 | 0 | 527.0 MB | `sha256:e30b2d29bf11` | `mcr.microsoft.com/openjdk/jdk:25-azurelinux@sha256:e30b2d29bf11dca58cfcb4ecdd823e42af6b3e226f2d5fbf4eef60e11223ce13` |
| 9 | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0` | 3.12.9 | 0 | 0 | 0 | 809.0 MB | `sha256:6e92b12751fd` | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0@sha256:6e92b12751fdc385f6c415e0a0fe6f630c59df631281deb295f465035320cd8e` |
| 10 | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0` | 3.12.9 | 0 | 0 | 0 | 841.0 MB | `sha256:7df9ca08df20` | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0@sha256:7df9ca08df200e793e1bd35ff92128c965c5cf43815b7beb78278b7dcabb9da3` |


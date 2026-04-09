# Daily Recommended Images by Language

_Generated: 2026-04-09T07:50:17Z. Criteria: lowest critical → high → total vulnerabilities → size. Top 10 per language per base OS._

**Note:** Image sizes are based on Linux amd64 platform as reported by `docker images` on GitHub runners. Actual sizes may vary significantly on other platforms (macOS, Windows, etc.).

## Scanned Repositories and Images

This report includes analysis from **37 configured sources** across 8 groups (see [repositories.json](../config/repositories.json)):

**Base / minimal images (no runtime):**

- `mcr.microsoft.com/azurelinux/base/core:3.0`
- `mcr.microsoft.com/azurelinux/distroless/base:3.0`
- `mcr.microsoft.com/azurelinux/distroless/minimal:3.0`

**Azure Linux Python and Node.js images:**

- `azurelinux/base/python`
- `azurelinux/base/nodejs`

**Azure Linux distroless runtime images:**

- `azurelinux/distroless/python`
- `azurelinux/distroless/nodejs`

**.NET Azure Linux images:**

- `mcr.microsoft.com/dotnet/aspnet:10.0-azurelinux3.0-distroless`
- `mcr.microsoft.com/dotnet/aspnet:9.0-azurelinux3.0-distroless`
- `mcr.microsoft.com/dotnet/aspnet:8.0-azurelinux3.0-distroless`
- `mcr.microsoft.com/dotnet/runtime:10.0-azurelinux3.0-distroless`
- `mcr.microsoft.com/dotnet/runtime:9.0-azurelinux3.0-distroless`
- `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0`
- `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0`

**.NET Ubuntu (Noble) images:**

- `mcr.microsoft.com/dotnet/aspnet:10.0-noble`
- `mcr.microsoft.com/dotnet/aspnet:9.0-noble`
- `mcr.microsoft.com/dotnet/aspnet:8.0-noble`
- `mcr.microsoft.com/dotnet/runtime:10.0-noble`
- `mcr.microsoft.com/dotnet/runtime:9.0-noble`
- `mcr.microsoft.com/dotnet/runtime:8.0-noble`
- `mcr.microsoft.com/dotnet/sdk:10.0-noble`
- `mcr.microsoft.com/dotnet/sdk:9.0-noble`
- `mcr.microsoft.com/dotnet/sdk:8.0-noble`

**.NET Debian images:**

- `mcr.microsoft.com/dotnet/aspnet:8.0`
- `mcr.microsoft.com/dotnet/runtime:8.0`
- `mcr.microsoft.com/dotnet/sdk:8.0`

**Go images:**

- `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0`
- `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0`

**OpenJDK images:**

- `mcr.microsoft.com/openjdk/jdk:25-azurelinux`
- `mcr.microsoft.com/openjdk/jdk:25-distroless`
- `mcr.microsoft.com/openjdk/jdk:25-ubuntu`
- `mcr.microsoft.com/openjdk/jdk:21-azurelinux`
- `mcr.microsoft.com/openjdk/jdk:21-distroless`
- `mcr.microsoft.com/openjdk/jdk:21-ubuntu`
- `mcr.microsoft.com/openjdk/jdk:17-distroless`
- `mcr.microsoft.com/openjdk/jdk:17-ubuntu`
- `mcr.microsoft.com/openjdk/jdk:11-distroless`

## Dotnet

### Azure Linux

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:9.0-azurelinux3.0-distroless` | 9.0.14 | 0 | 0 | 0 | 107.0 MB | `sha256:277fa93ff08f` | `mcr.microsoft.com/dotnet/runtime:9.0-azurelinux3.0-distroless@sha256:277fa93ff08fe17a5f865f4406cf57986233226ba4d57998088ce997c6e212b2` |
| 2 | `mcr.microsoft.com/dotnet/runtime:10.0-azurelinux3.0-distroless` | 10.0.5 | 0 | 0 | 0 | 112.0 MB | `sha256:f89ca4cdb258` | `mcr.microsoft.com/dotnet/runtime:10.0-azurelinux3.0-distroless@sha256:f89ca4cdb2589520a0c42133618dd85f4169daadf30839fb5fa5960708ef16c4` |
| 3 | `mcr.microsoft.com/dotnet/aspnet:8.0-azurelinux3.0-distroless` | 8.0.25 | 0 | 0 | 0 | 126.0 MB | `sha256:d0b8969dd9ec` | `mcr.microsoft.com/dotnet/aspnet:8.0-azurelinux3.0-distroless@sha256:d0b8969dd9ec92fcf7d37a652653ae0b07b457e90ea04a10bb7fd67185bf301c` |
| 4 | `mcr.microsoft.com/dotnet/aspnet:9.0-azurelinux3.0-distroless` | 9.0.14 | 0 | 0 | 0 | 132.0 MB | `sha256:8c422651b58c` | `mcr.microsoft.com/dotnet/aspnet:9.0-azurelinux3.0-distroless@sha256:8c422651b58cad627d85815f54ceeab28171dbb5fc633ec9e783708072613177` |
| 5 | `mcr.microsoft.com/dotnet/aspnet:10.0-azurelinux3.0-distroless` | 10.0.5 | 0 | 0 | 0 | 139.0 MB | `sha256:4926af2df5bb` | `mcr.microsoft.com/dotnet/aspnet:10.0-azurelinux3.0-distroless@sha256:4926af2df5bb2c23934ad25d7dc7c08910e8affd38327f99e8ee2d770d8ce145` |
| 6 | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0` | 9.0.312 | 0 | 0 | 0 | 880.0 MB | `sha256:f98093e58066` | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0@sha256:f98093e58066a09403630d3a4b0d5af82329a4eb39b76d41b6384be832e8fc32` |
| 7 | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0` | 10.0.201 | 0 | 0 | 0 | 933.0 MB | `sha256:52066c5f2566` | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0@sha256:52066c5f2566b1e251f0a7ff6c9380d655054d560d6f88a724d23b72448b9a3a` |

### Debian

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0` | 8.0.25 | 1 | 7 | 135 | 193.0 MB | `sha256:e66b853161cd` | `mcr.microsoft.com/dotnet/runtime:8.0@sha256:e66b853161cdde692f10863a7f5166048e16f7d6022d1eba03cc62367a540a1a` |
| 2 | `mcr.microsoft.com/dotnet/aspnet:8.0` | 8.0.25 | 1 | 7 | 135 | 218.0 MB | `sha256:f24d74e8185b` | `mcr.microsoft.com/dotnet/aspnet:8.0@sha256:f24d74e8185b14aba4c7563e059e0f526699d4730998a83515d3af03058e1266` |
| 3 | `mcr.microsoft.com/dotnet/sdk:8.0` | 8.0.419 | 1 | 18 | 222 | 851.0 MB | `sha256:dd09bcce84d9` | `mcr.microsoft.com/dotnet/sdk:8.0@sha256:dd09bcce84d9130e7f3e85c83a8ce9709e1d95e45c48c722a0d7923b38d8024c` |

### Ubuntu

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0-noble` | 8.0.25 | 0 | 0 | 17 | 193.0 MB | `sha256:919d1e7429df` | `mcr.microsoft.com/dotnet/runtime:8.0-noble@sha256:919d1e7429dfaca99487338465276987870f3b2a06422a8871bcbd3f44bc35d4` |
| 2 | `mcr.microsoft.com/dotnet/runtime:9.0-noble` | 9.0.14 | 0 | 0 | 17 | 198.0 MB | `sha256:12147f660897` | `mcr.microsoft.com/dotnet/runtime:9.0-noble@sha256:12147f66089752f72e215806e352d9b2336cda9b46dc72b9d425dddfecfa7d8b` |
| 3 | `mcr.microsoft.com/dotnet/runtime:10.0-noble` | 10.0.5 | 0 | 0 | 17 | 203.0 MB | `sha256:260c4d478b25` | `mcr.microsoft.com/dotnet/runtime:10.0-noble@sha256:260c4d478b25ccb5a84be74dc5ea5a772fb4fb520e9226ca8f1cb57bf1425fb1` |
| 4 | `mcr.microsoft.com/dotnet/aspnet:8.0-noble` | 8.0.25 | 0 | 0 | 17 | 217.0 MB | `sha256:3c08a6c357ba` | `mcr.microsoft.com/dotnet/aspnet:8.0-noble@sha256:3c08a6c357bac63ba74e23b54db4fd54996662cd2260ca332e69b03393e2e352` |
| 5 | `mcr.microsoft.com/dotnet/aspnet:9.0-noble` | 9.0.14 | 0 | 0 | 17 | 223.0 MB | `sha256:729acff37e08` | `mcr.microsoft.com/dotnet/aspnet:9.0-noble@sha256:729acff37e0859d9a69ad19cb43145b2306743c235397220fd6cb301d059e0e5` |
| 6 | `mcr.microsoft.com/dotnet/aspnet:10.0-noble` | 10.0.5 | 0 | 0 | 17 | 230.0 MB | `sha256:ccdca44cd4f2` | `mcr.microsoft.com/dotnet/aspnet:10.0-noble@sha256:ccdca44cd4f256d50187f920dc8ccc2a9ea7a8a4597ac1d51e08fddb2e3b3205` |
| 7 | `mcr.microsoft.com/dotnet/sdk:9.0-noble` | 9.0.312 | 0 | 0 | 21 | 837.0 MB | `sha256:446c032fc8de` | `mcr.microsoft.com/dotnet/sdk:9.0-noble@sha256:446c032fc8de02409c2aa1f34cbcceffae22bb6ab9a93bf19fa99a578cd066f3` |
| 8 | `mcr.microsoft.com/dotnet/sdk:10.0-noble` | 10.0.201 | 0 | 0 | 21 | 891.0 MB | `sha256:f061e5a7532b` | `mcr.microsoft.com/dotnet/sdk:10.0-noble@sha256:f061e5a7532b36fa1d1b684857fe1f504ba92115b9934f154643266613c44c62` |
| 9 | `mcr.microsoft.com/dotnet/sdk:8.0-noble` | 8.0.419 | 0 | 8 | 29 | 837.0 MB | `sha256:67f7a8d541a1` | `mcr.microsoft.com/dotnet/sdk:8.0-noble@sha256:67f7a8d541a18505a74326085ed997eb4409cb228d429c55d38ccaea3d486dce` |

## Go

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0` | 1.25.9 | 0 | 0 | 0 | 818.0 MB | `sha256:3efee8683ba8` | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0@sha256:3efee8683ba87f9b3e0155cd932f9b79b16354e11410bffaa91a974ad13b9954` |
| 2 | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0` | 1.26.2 | 0 | 0 | 0 | 850.0 MB | `sha256:33d887d995c3` | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0@sha256:33d887d995c314d0efbc0e8693a6821495f7b1b92516316fd9c8b43f74579dbd` |

## Java

### Azure Linux

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:21-azurelinux` | 21.0.10 | 0 | 0 | 0 | 483.0 MB | `sha256:6d25fbd43b93` | `mcr.microsoft.com/openjdk/jdk:21-azurelinux@sha256:6d25fbd43b9306e7d983ee56fafd85970cbd464925a7a067b03145bd69c7d2d9` |
| 2 | `mcr.microsoft.com/openjdk/jdk:25-azurelinux` | 25.0.2 | 0 | 0 | 0 | 527.0 MB | `sha256:1922a2189166` | `mcr.microsoft.com/openjdk/jdk:25-azurelinux@sha256:1922a2189166ff13d84c74b72c3d3195940cd6c08436b6c987b79ff4beb58616` |
| 3 | `mcr.microsoft.com/openjdk/jdk:11-distroless` | 11.0.30 | 0 | 0 | 6 | 320.0 MB | `sha256:8dee6fd212bb` | `mcr.microsoft.com/openjdk/jdk:11-distroless@sha256:8dee6fd212bb0a17aee5d3e98df9b39492c06f5410c9e4be35b1ae6ab1d246d2` |
| 4 | `mcr.microsoft.com/openjdk/jdk:17-distroless` | 17.0.18 | 0 | 0 | 6 | 324.0 MB | `sha256:ffa65843530d` | `mcr.microsoft.com/openjdk/jdk:17-distroless@sha256:ffa65843530d6d1fb04c3fc3778bb37867ab5b34f76a647650d49bab6ddb8f8b` |
| 5 | `mcr.microsoft.com/openjdk/jdk:21-distroless` | 21.0.10 | 0 | 0 | 6 | 351.0 MB | `sha256:bcab1bc3c7e2` | `mcr.microsoft.com/openjdk/jdk:21-distroless@sha256:bcab1bc3c7e2e82cfe54f8fe4b6e3f74959c2687a4fb4559de50ae4dc5126c3f` |
| 6 | `mcr.microsoft.com/openjdk/jdk:25-distroless` | 25.0.2 | 0 | 0 | 6 | 396.0 MB | `sha256:94729242b07b` | `mcr.microsoft.com/openjdk/jdk:25-distroless@sha256:94729242b07b77438e5c8a181e041f3395148c16dd4740f0cd653b25d901221c` |

### Ubuntu

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:17-ubuntu` | 17.0.18 | 0 | 0 | 74 | 428.0 MB | `sha256:3be0c6cf5878` | `mcr.microsoft.com/openjdk/jdk:17-ubuntu@sha256:3be0c6cf5878bd9a7ff56bcbb5e4f0d00a7799dfc39125cabd2be2314405d297` |
| 2 | `mcr.microsoft.com/openjdk/jdk:21-ubuntu` | 21.0.10 | 0 | 0 | 74 | 455.0 MB | `sha256:231f77efd3a8` | `mcr.microsoft.com/openjdk/jdk:21-ubuntu@sha256:231f77efd3a8ead114ac48a69874bd71a20b52d82ccb2de26579dbb28375ba57` |
| 3 | `mcr.microsoft.com/openjdk/jdk:25-ubuntu` | 25.0.2 | 0 | 0 | 74 | 500.0 MB | `sha256:29f3ba4dbb2b` | `mcr.microsoft.com/openjdk/jdk:25-ubuntu@sha256:29f3ba4dbb2ba3b1b69dbe71f204749cfbebac6866fa2dae826eece2d9476b5f` |

## Node

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/base/nodejs:24.13` | 24.13.0 | 0 | 0 | 0 | 163.0 MB | `sha256:2d1ed7ab9e4b` | `mcr.microsoft.com/azurelinux/base/nodejs:24.13@sha256:2d1ed7ab9e4b143703eefbbc1ce5d03bd4475da0a335f157e51b313a1948414f` |
| 2 | `mcr.microsoft.com/azurelinux/base/nodejs:24` | 24.13.0 | 0 | 0 | 0 | 163.0 MB | `sha256:2d1ed7ab9e4b` | `mcr.microsoft.com/azurelinux/base/nodejs:24@sha256:2d1ed7ab9e4b143703eefbbc1ce5d03bd4475da0a335f157e51b313a1948414f` |
| 3 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.13-nonroot` | 24.13.0 | 0 | 19 | 23 | 123.0 MB | `sha256:e25fe7994e0a` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.13-nonroot@sha256:e25fe7994e0a31cf3d6523d8b53a651f7390cfafd83cfe1e8706bc80f6988d34` |
| 4 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.13` | 24.13.0 | 0 | 19 | 23 | 123.0 MB | `sha256:9602d6864048` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.13@sha256:9602d686404823a802b0a858f5b65793ae2d6bd8900bcd64bb423d71f7897501` |
| 5 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24-nonroot` | 24.13.0 | 0 | 19 | 23 | 123.0 MB | `sha256:e25fe7994e0a` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24-nonroot@sha256:e25fe7994e0a31cf3d6523d8b53a651f7390cfafd83cfe1e8706bc80f6988d34` |
| 6 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24` | 24.13.0 | 0 | 19 | 23 | 123.0 MB | `sha256:9602d6864048` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24@sha256:9602d686404823a802b0a858f5b65793ae2d6bd8900bcd64bb423d71f7897501` |
| 7 | `mcr.microsoft.com/azurelinux/base/nodejs:20.14` | 20.14.0 | 1 | 4 | 8 | 146.0 MB | `sha256:b8a48ba88fdf` | `mcr.microsoft.com/azurelinux/base/nodejs:20.14@sha256:b8a48ba88fdff68c20a2895cb002d547ea47779e37e2fdcb6cd64e4a875f8a71` |
| 8 | `mcr.microsoft.com/azurelinux/base/nodejs:20` | 20.14.0 | 1 | 4 | 8 | 146.0 MB | `sha256:b8a48ba88fdf` | `mcr.microsoft.com/azurelinux/base/nodejs:20@sha256:b8a48ba88fdff68c20a2895cb002d547ea47779e37e2fdcb6cd64e4a875f8a71` |
| 9 | `mcr.microsoft.com/azurelinux/distroless/nodejs:20.14-nonroot` | 20.14.0 | 1 | 13 | 22 | 106.0 MB | `sha256:60773809112a` | `mcr.microsoft.com/azurelinux/distroless/nodejs:20.14-nonroot@sha256:60773809112a278016c356118123621ddf763c80f970faa36e6366bdaf794d50` |
| 10 | `mcr.microsoft.com/azurelinux/distroless/nodejs:20.14` | 20.14.0 | 1 | 13 | 22 | 106.0 MB | `sha256:dfc185c7bb4a` | `mcr.microsoft.com/azurelinux/distroless/nodejs:20.14@sha256:dfc185c7bb4ac4f740b6581ee50d2a5f868dd12c7626b40a1fc5698c980c12e6` |

## Python

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot` | 3.12.9 | 0 | 0 | 0 | 83.7 MB | `sha256:ddef666511ee` | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot@sha256:ddef666511eefea84adc3654905d55619eddfb89fe045bbcadc385357a5883c6` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/python:3.12` | 3.12.9 | 0 | 0 | 0 | 83.7 MB | `sha256:2cb276ad2888` | `mcr.microsoft.com/azurelinux/distroless/python:3.12@sha256:2cb276ad288899c783733abc4f0932814ce1d9c8658effef4782ce3b4c1a8c49` |
| 3 | `mcr.microsoft.com/azurelinux/distroless/python:3-nonroot` | 3.12.9 | 0 | 0 | 0 | 83.7 MB | `sha256:ddef666511ee` | `mcr.microsoft.com/azurelinux/distroless/python:3-nonroot@sha256:ddef666511eefea84adc3654905d55619eddfb89fe045bbcadc385357a5883c6` |
| 4 | `mcr.microsoft.com/azurelinux/distroless/python:3` | 3.12.9 | 0 | 0 | 0 | 83.7 MB | `sha256:2cb276ad2888` | `mcr.microsoft.com/azurelinux/distroless/python:3@sha256:2cb276ad288899c783733abc4f0932814ce1d9c8658effef4782ce3b4c1a8c49` |
| 5 | `mcr.microsoft.com/azurelinux/base/python:3.12` | 3.12.9 | 0 | 0 | 0 | 139.0 MB | `sha256:8300706a0e64` | `mcr.microsoft.com/azurelinux/base/python:3.12@sha256:8300706a0e644a7e260b6a464a1ae3967181ee9057a4accceaaf0ecdedb1d4a2` |
| 6 | `mcr.microsoft.com/azurelinux/base/python:3` | 3.12.9 | 0 | 0 | 0 | 139.0 MB | `sha256:8300706a0e64` | `mcr.microsoft.com/azurelinux/base/python:3@sha256:8300706a0e644a7e260b6a464a1ae3967181ee9057a4accceaaf0ecdedb1d4a2` |
| 7 | `mcr.microsoft.com/openjdk/jdk:21-azurelinux` | 3.12.9 | 0 | 0 | 0 | 483.0 MB | `sha256:6d25fbd43b93` | `mcr.microsoft.com/openjdk/jdk:21-azurelinux@sha256:6d25fbd43b9306e7d983ee56fafd85970cbd464925a7a067b03145bd69c7d2d9` |
| 8 | `mcr.microsoft.com/openjdk/jdk:25-azurelinux` | 3.12.9 | 0 | 0 | 0 | 527.0 MB | `sha256:1922a2189166` | `mcr.microsoft.com/openjdk/jdk:25-azurelinux@sha256:1922a2189166ff13d84c74b72c3d3195940cd6c08436b6c987b79ff4beb58616` |
| 9 | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0` | 3.12.9 | 0 | 0 | 0 | 818.0 MB | `sha256:3efee8683ba8` | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0@sha256:3efee8683ba87f9b3e0155cd932f9b79b16354e11410bffaa91a974ad13b9954` |
| 10 | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0` | 3.12.9 | 0 | 0 | 0 | 850.0 MB | `sha256:33d887d995c3` | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0@sha256:33d887d995c314d0efbc0e8693a6821495f7b1b92516316fd9c8b43f74579dbd` |

## Base / No Runtime

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0` | 3.0 | 0 | 0 | 0 | 3.7 MB | `sha256:5a66f9f16ac6` | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0@sha256:5a66f9f16ac675db2a8229dac72d83811b73b502d6ad192d8b374c7f3be498af` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/base:3.0` | 3.0 | 0 | 0 | 0 | 34.1 MB | `sha256:32820d2cf20e` | `mcr.microsoft.com/azurelinux/distroless/base:3.0@sha256:32820d2cf20e896aa9111742dd683dd0ccff370f742e256889bb3bb50320c0d4` |
| 3 | `mcr.microsoft.com/azurelinux/base/core:3.0` | 3.0 | 0 | 0 | 0 | 75.3 MB | `sha256:35149ae8dd17` | `mcr.microsoft.com/azurelinux/base/core:3.0@sha256:35149ae8dd179684f969944f54a337c665a64e702486154eb44253fb39c2505b` |

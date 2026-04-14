# Daily Recommended Images by Language

_Generated: 2026-04-14T04:57:08Z. Criteria: lowest critical → high → total vulnerabilities → size. Top 10 per language per base OS._

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
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0` | 8.0.25 | 1 | 9 | 137 | 193.0 MB | `sha256:e66b853161cd` | `mcr.microsoft.com/dotnet/runtime:8.0@sha256:e66b853161cdde692f10863a7f5166048e16f7d6022d1eba03cc62367a540a1a` |
| 2 | `mcr.microsoft.com/dotnet/aspnet:8.0` | 8.0.25 | 1 | 9 | 137 | 218.0 MB | `sha256:f24d74e8185b` | `mcr.microsoft.com/dotnet/aspnet:8.0@sha256:f24d74e8185b14aba4c7563e059e0f526699d4730998a83515d3af03058e1266` |
| 3 | `mcr.microsoft.com/dotnet/sdk:8.0` | 8.0.419 | 1 | 20 | 224 | 851.0 MB | `sha256:dd09bcce84d9` | `mcr.microsoft.com/dotnet/sdk:8.0@sha256:dd09bcce84d9130e7f3e85c83a8ce9709e1d95e45c48c722a0d7923b38d8024c` |

### Ubuntu

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0-noble` | 8.0.25 | 0 | 0 | 5 | 193.0 MB | `sha256:daac7d1a3912` | `mcr.microsoft.com/dotnet/runtime:8.0-noble@sha256:daac7d1a3912aa872708ef996c7c76a1c33f2ea62faa4366ce4006c95f89c751` |
| 2 | `mcr.microsoft.com/dotnet/runtime:9.0-noble` | 9.0.14 | 0 | 0 | 5 | 198.0 MB | `sha256:efd18e356994` | `mcr.microsoft.com/dotnet/runtime:9.0-noble@sha256:efd18e3569944e5729ffe78f8ae221e38fa2c3e6cd47e5a1b169f6258c10cae4` |
| 3 | `mcr.microsoft.com/dotnet/runtime:10.0-noble` | 10.0.5 | 0 | 0 | 5 | 203.0 MB | `sha256:d899417078f6` | `mcr.microsoft.com/dotnet/runtime:10.0-noble@sha256:d899417078f6f2ace195c70bd63c4851f4c2e29c38f50fcfb46f2be5f7e3638f` |
| 4 | `mcr.microsoft.com/dotnet/aspnet:8.0-noble` | 8.0.25 | 0 | 0 | 5 | 217.0 MB | `sha256:0f586d14e376` | `mcr.microsoft.com/dotnet/aspnet:8.0-noble@sha256:0f586d14e376db2d8fef58594f770fe81d74dbd1605fe07c22d871e88c626739` |
| 5 | `mcr.microsoft.com/dotnet/aspnet:9.0-noble` | 9.0.14 | 0 | 0 | 5 | 223.0 MB | `sha256:f666852c32c7` | `mcr.microsoft.com/dotnet/aspnet:9.0-noble@sha256:f666852c32c7f81d1cc5ee9d5fa85bcd6e33207814994d119cc4ddc102607f49` |
| 6 | `mcr.microsoft.com/dotnet/aspnet:10.0-noble` | 10.0.5 | 0 | 0 | 5 | 230.0 MB | `sha256:c433886fdfe3` | `mcr.microsoft.com/dotnet/aspnet:10.0-noble@sha256:c433886fdfe33c6427966a412328867b2be9a64f540a105d08943c2dc6fba39b` |
| 7 | `mcr.microsoft.com/dotnet/sdk:9.0-noble` | 9.0.312 | 0 | 0 | 9 | 837.0 MB | `sha256:9355633cd511` | `mcr.microsoft.com/dotnet/sdk:9.0-noble@sha256:9355633cd511585611cca51effa906e3bf8078cb7bc18880686a38af7b54aacc` |
| 8 | `mcr.microsoft.com/dotnet/sdk:10.0-noble` | 10.0.201 | 0 | 0 | 9 | 891.0 MB | `sha256:127d7d4d601a` | `mcr.microsoft.com/dotnet/sdk:10.0-noble@sha256:127d7d4d601ae26b8e04c54efb37e9ce8766931bded0ee59fcd799afd21d6850` |
| 9 | `mcr.microsoft.com/dotnet/sdk:8.0-noble` | 8.0.419 | 0 | 8 | 17 | 837.0 MB | `sha256:7b912fde4456` | `mcr.microsoft.com/dotnet/sdk:8.0-noble@sha256:7b912fde44560911ff5accca61bf9c2f59c2484ed11e45645ff23eceaea0b733` |

## Go

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0` | 1.25.9 | 0 | 0 | 0 | 818.0 MB | `sha256:12078713878e` | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0@sha256:12078713878ec827997a05d29de4430056f7e7e982deae74f96bd305ba4e2a9e` |
| 2 | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0` | 1.26.2 | 0 | 0 | 0 | 850.0 MB | `sha256:d8b025ce76cc` | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0@sha256:d8b025ce76cc35ed03f1c94b0ccdde04286e83a876b873540390e7dd3e7bb888` |

## Java

### Azure Linux

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:11-distroless` | 11.0.30 | 0 | 0 | 0 | 320.0 MB | `sha256:1f8f2142a0b3` | `mcr.microsoft.com/openjdk/jdk:11-distroless@sha256:1f8f2142a0b3264ab1f6f1f74bedaff41f5ce7e3c1ae24839f98412e5af439d9` |
| 2 | `mcr.microsoft.com/openjdk/jdk:17-distroless` | 17.0.18 | 0 | 0 | 0 | 324.0 MB | `sha256:4f7b81257a6f` | `mcr.microsoft.com/openjdk/jdk:17-distroless@sha256:4f7b81257a6f91464199ff59f67bae5b1c7d5be51de51999a12b08604ba6939c` |
| 3 | `mcr.microsoft.com/openjdk/jdk:21-distroless` | 21.0.10 | 0 | 0 | 0 | 351.0 MB | `sha256:d0a1e6199138` | `mcr.microsoft.com/openjdk/jdk:21-distroless@sha256:d0a1e6199138394e412f34c25a2742f6155df7b5bcc917baade5bd7fed648302` |
| 4 | `mcr.microsoft.com/openjdk/jdk:25-distroless` | 25.0.2 | 0 | 0 | 0 | 396.0 MB | `sha256:d2ac87450c84` | `mcr.microsoft.com/openjdk/jdk:25-distroless@sha256:d2ac87450c84c688d14de166b78924ee979030131626f3953eb3143eebc36c6c` |
| 5 | `mcr.microsoft.com/openjdk/jdk:21-azurelinux` | 21.0.10 | 0 | 0 | 0 | 483.0 MB | `sha256:7db3552814e0` | `mcr.microsoft.com/openjdk/jdk:21-azurelinux@sha256:7db3552814e03af80b1a6762476e9e977ab4fcd61231efab68031125a0c809c1` |
| 6 | `mcr.microsoft.com/openjdk/jdk:25-azurelinux` | 25.0.2 | 0 | 0 | 0 | 528.0 MB | `sha256:0187b4323c77` | `mcr.microsoft.com/openjdk/jdk:25-azurelinux@sha256:0187b4323c778eebb15f07e73284ac2d941a91bac0114d124930e70099a7f601` |

### Ubuntu

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:17-ubuntu` | 17.0.18 | 0 | 0 | 74 | 427.0 MB | `sha256:0086684a6d57` | `mcr.microsoft.com/openjdk/jdk:17-ubuntu@sha256:0086684a6d57806c558b53957f6b76480d95fbea96a6c3e40e6d5b214fde162d` |
| 2 | `mcr.microsoft.com/openjdk/jdk:21-ubuntu` | 21.0.10 | 0 | 0 | 74 | 454.0 MB | `sha256:ca691fefdb30` | `mcr.microsoft.com/openjdk/jdk:21-ubuntu@sha256:ca691fefdb3015f9b21b95fa481fae6e836ccdde1263be2d44934c702d0c7c23` |
| 3 | `mcr.microsoft.com/openjdk/jdk:25-ubuntu` | 25.0.2 | 0 | 0 | 74 | 499.0 MB | `sha256:2960c2ba70d8` | `mcr.microsoft.com/openjdk/jdk:25-ubuntu@sha256:2960c2ba70d831f732a57f7f2e4dbe63dae903abd4e7b716b3bf9228710233ec` |

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
| 7 | `mcr.microsoft.com/openjdk/jdk:21-azurelinux` | 3.12.9 | 0 | 0 | 0 | 483.0 MB | `sha256:7db3552814e0` | `mcr.microsoft.com/openjdk/jdk:21-azurelinux@sha256:7db3552814e03af80b1a6762476e9e977ab4fcd61231efab68031125a0c809c1` |
| 8 | `mcr.microsoft.com/openjdk/jdk:25-azurelinux` | 3.12.9 | 0 | 0 | 0 | 528.0 MB | `sha256:0187b4323c77` | `mcr.microsoft.com/openjdk/jdk:25-azurelinux@sha256:0187b4323c778eebb15f07e73284ac2d941a91bac0114d124930e70099a7f601` |
| 9 | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0` | 3.12.9 | 0 | 0 | 0 | 818.0 MB | `sha256:12078713878e` | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0@sha256:12078713878ec827997a05d29de4430056f7e7e982deae74f96bd305ba4e2a9e` |
| 10 | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0` | 3.12.9 | 0 | 0 | 0 | 850.0 MB | `sha256:d8b025ce76cc` | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0@sha256:d8b025ce76cc35ed03f1c94b0ccdde04286e83a876b873540390e7dd3e7bb888` |

## Base / No Runtime

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0` | 3.0 | 0 | 0 | 0 | 3.7 MB | `sha256:5a66f9f16ac6` | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0@sha256:5a66f9f16ac675db2a8229dac72d83811b73b502d6ad192d8b374c7f3be498af` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/base:3.0` | 3.0 | 0 | 0 | 0 | 34.1 MB | `sha256:32820d2cf20e` | `mcr.microsoft.com/azurelinux/distroless/base:3.0@sha256:32820d2cf20e896aa9111742dd683dd0ccff370f742e256889bb3bb50320c0d4` |
| 3 | `mcr.microsoft.com/azurelinux/base/core:3.0` | 3.0 | 0 | 0 | 0 | 75.3 MB | `sha256:35149ae8dd17` | `mcr.microsoft.com/azurelinux/base/core:3.0@sha256:35149ae8dd179684f969944f54a337c665a64e702486154eb44253fb39c2505b` |

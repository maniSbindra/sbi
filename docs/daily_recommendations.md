# Daily Recommended Images by Language

_Generated: 2026-03-12T04:26:06Z. Criteria: lowest critical → high → total vulnerabilities → size. Top 10 per language per base OS._

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
| 1 | `mcr.microsoft.com/dotnet/runtime:9.0-azurelinux3.0-distroless` | 9.0.14 | 0 | 0 | 0 | 105.0 MB | `sha256:98b136128d9d` | `mcr.microsoft.com/dotnet/runtime:9.0-azurelinux3.0-distroless@sha256:98b136128d9d6d61a76075895d946cbc32d1e9af50cf9ad95f85990221fb6338` |
| 2 | `mcr.microsoft.com/dotnet/runtime:10.0-azurelinux3.0-distroless` | 10.0.4 | 0 | 0 | 0 | 110.0 MB | `sha256:55a80f6bd60c` | `mcr.microsoft.com/dotnet/runtime:10.0-azurelinux3.0-distroless@sha256:55a80f6bd60cd1d31daaf7a0505e27807498405ae7489a90d3886cdcff957d7c` |
| 3 | `mcr.microsoft.com/dotnet/aspnet:8.0-azurelinux3.0-distroless` | 8.0.25 | 0 | 0 | 0 | 124.0 MB | `sha256:432ed2f83fd2` | `mcr.microsoft.com/dotnet/aspnet:8.0-azurelinux3.0-distroless@sha256:432ed2f83fd2ed8628baf643fc4252250938ef868da6864e71271e6829fbd2a6` |
| 4 | `mcr.microsoft.com/dotnet/aspnet:9.0-azurelinux3.0-distroless` | 9.0.14 | 0 | 0 | 0 | 130.0 MB | `sha256:1a3bfd111f1c` | `mcr.microsoft.com/dotnet/aspnet:9.0-azurelinux3.0-distroless@sha256:1a3bfd111f1c99effb95cc80ad9c08b87b34c21f2d46ee58688c44587ea2bd0b` |
| 5 | `mcr.microsoft.com/dotnet/aspnet:10.0-azurelinux3.0-distroless` | 10.0.4 | 0 | 0 | 0 | 137.0 MB | `sha256:d6e7d0163ea0` | `mcr.microsoft.com/dotnet/aspnet:10.0-azurelinux3.0-distroless@sha256:d6e7d0163ea0c12f81d415ffbb366ba664152528fbd9176f22bacc745ac108f8` |
| 6 | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0` | 9.0.312 | 0 | 0 | 0 | 879.0 MB | `sha256:73837c6a200c` | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0@sha256:73837c6a200ce49db381d84c855d146fafe59e4f1e13227ebaeac48e5190deaf` |
| 7 | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0` | 10.0.200 | 0 | 0 | 0 | 933.0 MB | `sha256:478302acc4bd` | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0@sha256:478302acc4bda2d4f0179a7077b3b04670ce888311fc874cf3596ba9f5809680` |

### Debian

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0` | 8.0.25 | 1 | 2 | 95 | 193.0 MB | `sha256:73b3de5b2f82` | `mcr.microsoft.com/dotnet/runtime:8.0@sha256:73b3de5b2f826f48236674089e3a3af492d9eac4e714b4af8c8d3c7fe0d3f41c` |
| 2 | `mcr.microsoft.com/dotnet/aspnet:8.0` | 8.0.25 | 1 | 2 | 95 | 218.0 MB | `sha256:a1a6ed414bde` | `mcr.microsoft.com/dotnet/aspnet:8.0@sha256:a1a6ed414bde21fce19c1f84e7b4c17d8020d9a82dad55a8adb3775ae5be644d` |
| 3 | `mcr.microsoft.com/dotnet/sdk:8.0` | 8.0.419 | 1 | 16 | 186 | 850.0 MB | `sha256:312aeb1357b2` | `mcr.microsoft.com/dotnet/sdk:8.0@sha256:312aeb1357b2461f5a729b5ca38a54b8381c8e7880d0672585e8613cdc972096` |

### Ubuntu

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0-noble` | 8.0.25 | 0 | 0 | 12 | 193.0 MB | `sha256:4afe8693b45c` | `mcr.microsoft.com/dotnet/runtime:8.0-noble@sha256:4afe8693b45ca972c98fd8b04bbe3d898111f1a9d1754f1c8fabae04da0c93aa` |
| 2 | `mcr.microsoft.com/dotnet/runtime:9.0-noble` | 9.0.14 | 0 | 0 | 12 | 198.0 MB | `sha256:ff70733f9ded` | `mcr.microsoft.com/dotnet/runtime:9.0-noble@sha256:ff70733f9ded21a73a0e0b17eee7a7c50b732f5ebc79e710ede6491982cc4fa0` |
| 3 | `mcr.microsoft.com/dotnet/runtime:10.0-noble` | 10.0.4 | 0 | 0 | 12 | 203.0 MB | `sha256:b1fc29ba8c7c` | `mcr.microsoft.com/dotnet/runtime:10.0-noble@sha256:b1fc29ba8c7cf56eacc4aaf078dbbfa70b533130968957356d0fef95d3ac62ac` |
| 4 | `mcr.microsoft.com/dotnet/aspnet:8.0-noble` | 8.0.25 | 0 | 0 | 12 | 217.0 MB | `sha256:69277bbf9a95` | `mcr.microsoft.com/dotnet/aspnet:8.0-noble@sha256:69277bbf9a95acea6be093f7e546e5ab9ee55768f12a9f32c3d543cd350d039e` |
| 5 | `mcr.microsoft.com/dotnet/aspnet:9.0-noble` | 9.0.14 | 0 | 0 | 12 | 223.0 MB | `sha256:281c970f6d43` | `mcr.microsoft.com/dotnet/aspnet:9.0-noble@sha256:281c970f6d43a3dc577d61c635275faf08e3ceeef4533462fe7e30884a23c4e6` |
| 6 | `mcr.microsoft.com/dotnet/aspnet:10.0-noble` | 10.0.4 | 0 | 0 | 12 | 230.0 MB | `sha256:8b75cdf59a50` | `mcr.microsoft.com/dotnet/aspnet:10.0-noble@sha256:8b75cdf59a5068d9adfd8a6d202cc7671b2dc8f5f46c51e3b88a0a632e8fad1f` |
| 7 | `mcr.microsoft.com/dotnet/sdk:9.0-noble` | 9.0.312 | 0 | 0 | 29 | 836.0 MB | `sha256:1c1583362df1` | `mcr.microsoft.com/dotnet/sdk:9.0-noble@sha256:1c1583362df1cd0caa017418d5ada26b6db140418fddb6c2718c41eb0b36d3ad` |
| 8 | `mcr.microsoft.com/dotnet/sdk:10.0-noble` | 10.0.200 | 0 | 0 | 29 | 890.0 MB | `sha256:ea13841f10c6` | `mcr.microsoft.com/dotnet/sdk:10.0-noble@sha256:ea13841f10c6c410a6df63c6c97ab549dfd2b5fcfff1c00186531ba30208117d` |
| 9 | `mcr.microsoft.com/dotnet/sdk:8.0-noble` | 8.0.419 | 0 | 8 | 37 | 836.0 MB | `sha256:d64fbc29eff1` | `mcr.microsoft.com/dotnet/sdk:8.0-noble@sha256:d64fbc29eff1787a23da76a6a580d4f2bcc8c6a8d3d432d74723fcd401f6b22a` |

## Go

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0` | 1.25.8 | 0 | 0 | 0 | 809.0 MB | `sha256:8deb7086a5fd` | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0@sha256:8deb7086a5fd7fe8b4258ca4c3431141b3d0daafc514c7da09d01da43a6e6ad0` |
| 2 | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0` | 1.26.1 | 0 | 0 | 0 | 841.0 MB | `sha256:1c76136f9648` | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0@sha256:1c76136f964815a0eeea177a348fc7a7ca19c9c1a25e6b31a4666dbb5b210cc4` |

## Java

### Azure Linux

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:21-azurelinux` | 21.0.10 | 0 | 0 | 0 | 491.0 MB | `sha256:4306b0bf5a90` | `mcr.microsoft.com/openjdk/jdk:21-azurelinux@sha256:4306b0bf5a90920168351a616f826636d5afe9365ea60f2bcb372d69e0647a64` |
| 2 | `mcr.microsoft.com/openjdk/jdk:25-azurelinux` | 25.0.2 | 0 | 0 | 0 | 536.0 MB | `sha256:507a1984974e` | `mcr.microsoft.com/openjdk/jdk:25-azurelinux@sha256:507a1984974efbbeb7ec71580e91ca12b70dfa35a880f2e58670c5461de231e3` |
| 3 | `mcr.microsoft.com/openjdk/jdk:11-distroless` | 11.0.30 | 0 | 0 | 3 | 320.0 MB | `sha256:6a8ef29c095a` | `mcr.microsoft.com/openjdk/jdk:11-distroless@sha256:6a8ef29c095a18f0ede670977df3c25cd90e4a9fd7eb37793a0456f72dbdd834` |
| 4 | `mcr.microsoft.com/openjdk/jdk:17-distroless` | 17.0.18 | 0 | 0 | 3 | 323.0 MB | `sha256:b364e47d93d7` | `mcr.microsoft.com/openjdk/jdk:17-distroless@sha256:b364e47d93d7abdb1d99ff6be822eb2b4513899c080eaf295130f0839296836c` |
| 5 | `mcr.microsoft.com/openjdk/jdk:21-distroless` | 21.0.10 | 0 | 0 | 3 | 351.0 MB | `sha256:c74db69a9bf4` | `mcr.microsoft.com/openjdk/jdk:21-distroless@sha256:c74db69a9bf40dc88bea347746f6f6e199e14adebe7a72c433de92284ac0c445` |
| 6 | `mcr.microsoft.com/openjdk/jdk:25-distroless` | 25.0.2 | 0 | 0 | 3 | 396.0 MB | `sha256:18b5c6afb498` | `mcr.microsoft.com/openjdk/jdk:25-distroless@sha256:18b5c6afb498df7a6fc0528602809ff48d3eaabd6f20237f1bf7c0c60d400254` |

### Ubuntu

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:17-ubuntu` | 17.0.18 | 0 | 0 | 66 | 430.0 MB | `sha256:f03c5fc956d1` | `mcr.microsoft.com/openjdk/jdk:17-ubuntu@sha256:f03c5fc956d13f0907b489221b42d79d0ea1942f9b975314424e1c9dfd147537` |
| 2 | `mcr.microsoft.com/openjdk/jdk:21-ubuntu` | 21.0.10 | 0 | 0 | 66 | 457.0 MB | `sha256:861d770bc07c` | `mcr.microsoft.com/openjdk/jdk:21-ubuntu@sha256:861d770bc07cfa0fb2dec123e82c052469156dcf5cd0b26a1a1c685c6332123c` |
| 3 | `mcr.microsoft.com/openjdk/jdk:25-ubuntu` | 25.0.2 | 0 | 0 | 66 | 502.0 MB | `sha256:dc63a92a778d` | `mcr.microsoft.com/openjdk/jdk:25-ubuntu@sha256:dc63a92a778d19fb98edba828a8b5174d99cfcdbd6ba71ae90f95ed6dbb1d228` |

## Node

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/base/nodejs:20.14` | 20.14.0 | 0 | 0 | 0 | 146.0 MB | `sha256:b8a48ba88fdf` | `mcr.microsoft.com/azurelinux/base/nodejs:20.14@sha256:b8a48ba88fdff68c20a2895cb002d547ea47779e37e2fdcb6cd64e4a875f8a71` |
| 2 | `mcr.microsoft.com/azurelinux/base/nodejs:20` | 20.14.0 | 0 | 0 | 0 | 146.0 MB | `sha256:b8a48ba88fdf` | `mcr.microsoft.com/azurelinux/base/nodejs:20@sha256:b8a48ba88fdff68c20a2895cb002d547ea47779e37e2fdcb6cd64e4a875f8a71` |
| 3 | `mcr.microsoft.com/azurelinux/distroless/nodejs:20.14-nonroot` | 20.14.0 | 0 | 11 | 13 | 106.0 MB | `sha256:60773809112a` | `mcr.microsoft.com/azurelinux/distroless/nodejs:20.14-nonroot@sha256:60773809112a278016c356118123621ddf763c80f970faa36e6366bdaf794d50` |
| 4 | `mcr.microsoft.com/azurelinux/distroless/nodejs:20.14` | 20.14.0 | 0 | 11 | 13 | 106.0 MB | `sha256:dfc185c7bb4a` | `mcr.microsoft.com/azurelinux/distroless/nodejs:20.14@sha256:dfc185c7bb4ac4f740b6581ee50d2a5f868dd12c7626b40a1fc5698c980c12e6` |
| 5 | `mcr.microsoft.com/azurelinux/distroless/nodejs:20-nonroot` | 20.14.0 | 0 | 11 | 13 | 106.0 MB | `sha256:60773809112a` | `mcr.microsoft.com/azurelinux/distroless/nodejs:20-nonroot@sha256:60773809112a278016c356118123621ddf763c80f970faa36e6366bdaf794d50` |
| 6 | `mcr.microsoft.com/azurelinux/distroless/nodejs:20` | 20.14.0 | 0 | 11 | 13 | 106.0 MB | `sha256:dfc185c7bb4a` | `mcr.microsoft.com/azurelinux/distroless/nodejs:20@sha256:dfc185c7bb4ac4f740b6581ee50d2a5f868dd12c7626b40a1fc5698c980c12e6` |

## Python

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot` | 3.12.9 | 0 | 0 | 0 | 83.6 MB | `sha256:b422aa585030` | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot@sha256:b422aa5850309ddd74c2be3cd329fb8f3d8318c7106d0dbc96075def3153ae7d` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/python:3.12` | 3.12.9 | 0 | 0 | 0 | 83.6 MB | `sha256:3db9d5eeb703` | `mcr.microsoft.com/azurelinux/distroless/python:3.12@sha256:3db9d5eeb703dc18befbc396d4b4f2b9dd30c0dd82ecd49e12edc50f9b3131ec` |
| 3 | `mcr.microsoft.com/azurelinux/distroless/python:3-nonroot` | 3.12.9 | 0 | 0 | 0 | 83.6 MB | `sha256:b422aa585030` | `mcr.microsoft.com/azurelinux/distroless/python:3-nonroot@sha256:b422aa5850309ddd74c2be3cd329fb8f3d8318c7106d0dbc96075def3153ae7d` |
| 4 | `mcr.microsoft.com/azurelinux/distroless/python:3` | 3.12.9 | 0 | 0 | 0 | 83.6 MB | `sha256:3db9d5eeb703` | `mcr.microsoft.com/azurelinux/distroless/python:3@sha256:3db9d5eeb703dc18befbc396d4b4f2b9dd30c0dd82ecd49e12edc50f9b3131ec` |
| 5 | `mcr.microsoft.com/azurelinux/base/python:3.12` | 3.12.9 | 0 | 0 | 0 | 139.0 MB | `sha256:8798301c18bf` | `mcr.microsoft.com/azurelinux/base/python:3.12@sha256:8798301c18bf675d2a4c6fe45f01aa2fe8ab469f05e571a07579f63f865a4809` |
| 6 | `mcr.microsoft.com/azurelinux/base/python:3` | 3.12.9 | 0 | 0 | 0 | 139.0 MB | `sha256:8798301c18bf` | `mcr.microsoft.com/azurelinux/base/python:3@sha256:8798301c18bf675d2a4c6fe45f01aa2fe8ab469f05e571a07579f63f865a4809` |
| 7 | `mcr.microsoft.com/openjdk/jdk:21-azurelinux` | 3.12.9 | 0 | 0 | 0 | 491.0 MB | `sha256:4306b0bf5a90` | `mcr.microsoft.com/openjdk/jdk:21-azurelinux@sha256:4306b0bf5a90920168351a616f826636d5afe9365ea60f2bcb372d69e0647a64` |
| 8 | `mcr.microsoft.com/openjdk/jdk:25-azurelinux` | 3.12.9 | 0 | 0 | 0 | 536.0 MB | `sha256:507a1984974e` | `mcr.microsoft.com/openjdk/jdk:25-azurelinux@sha256:507a1984974efbbeb7ec71580e91ca12b70dfa35a880f2e58670c5461de231e3` |
| 9 | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0` | 3.12.9 | 0 | 0 | 0 | 809.0 MB | `sha256:8deb7086a5fd` | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0@sha256:8deb7086a5fd7fe8b4258ca4c3431141b3d0daafc514c7da09d01da43a6e6ad0` |
| 10 | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0` | 3.12.9 | 0 | 0 | 0 | 841.0 MB | `sha256:1c76136f9648` | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0@sha256:1c76136f964815a0eeea177a348fc7a7ca19c9c1a25e6b31a4666dbb5b210cc4` |

## Base / No Runtime

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0` | 3.0 | 0 | 0 | 0 | 3.6 MB | `sha256:22810fd97d6a` | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0@sha256:22810fd97d6ad5ec7d5bdd5b00233a3050be01d9e26b47b16cb6f1a7f178834b` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/base:3.0` | 3.0 | 0 | 0 | 0 | 34.0 MB | `sha256:09b4ed711653` | `mcr.microsoft.com/azurelinux/distroless/base:3.0@sha256:09b4ed711653fc770cbc35c3b9088d4d4c8c835187b01ddf7b2ed88de5fa5366` |
| 3 | `mcr.microsoft.com/azurelinux/base/core:3.0` | 3.0 | 0 | 0 | 0 | 75.3 MB | `sha256:a452d39c9157` | `mcr.microsoft.com/azurelinux/base/core:3.0@sha256:a452d39c91576f5a2c983c7d3b62521fabd08e16b4a7237e24bf2be3b06e1651` |

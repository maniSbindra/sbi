# Daily Recommended Images by Language

_Generated: 2026-04-01T04:57:35Z. Criteria: lowest critical → high → total vulnerabilities → size. Top 10 per language per base OS._

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
| 1 | `mcr.microsoft.com/dotnet/runtime:9.0-azurelinux3.0-distroless` | 9.0.14 | 0 | 0 | 0 | 105.0 MB | `sha256:fc4b11480aae` | `mcr.microsoft.com/dotnet/runtime:9.0-azurelinux3.0-distroless@sha256:fc4b11480aae92c558bd6eecd41190bb098607ab4a6a7c00ad99506a795c65b8` |
| 2 | `mcr.microsoft.com/dotnet/runtime:10.0-azurelinux3.0-distroless` | 10.0.5 | 0 | 0 | 0 | 110.0 MB | `sha256:685d02669c6a` | `mcr.microsoft.com/dotnet/runtime:10.0-azurelinux3.0-distroless@sha256:685d02669c6a61a5caabe52ea1fae2428fcc93e8feaf7ff7978d6555d6047bba` |
| 3 | `mcr.microsoft.com/dotnet/aspnet:8.0-azurelinux3.0-distroless` | 8.0.25 | 0 | 0 | 0 | 124.0 MB | `sha256:86629c303fa4` | `mcr.microsoft.com/dotnet/aspnet:8.0-azurelinux3.0-distroless@sha256:86629c303fa47d08c888eace6f9b837ae273834594966c5a97a25917b23a96b9` |
| 4 | `mcr.microsoft.com/dotnet/aspnet:9.0-azurelinux3.0-distroless` | 9.0.14 | 0 | 0 | 0 | 130.0 MB | `sha256:8d88def665d8` | `mcr.microsoft.com/dotnet/aspnet:9.0-azurelinux3.0-distroless@sha256:8d88def665d82658a9d6ed6808df73592039bd39439fbb0dfcc35dd2f22c68f7` |
| 5 | `mcr.microsoft.com/dotnet/aspnet:10.0-azurelinux3.0-distroless` | 10.0.5 | 0 | 0 | 0 | 137.0 MB | `sha256:0fd8551e4e86` | `mcr.microsoft.com/dotnet/aspnet:10.0-azurelinux3.0-distroless@sha256:0fd8551e4e86acab6697e49caef6ccb44edbd3584b05882caab3e482ff94692a` |
| 6 | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0` | 9.0.312 | 0 | 4 | 4 | 879.0 MB | `sha256:f1c4c2210b8d` | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0@sha256:f1c4c2210b8d225c03ef11d79be475e996daf5646d3606a2cbf03abc95919d99` |
| 7 | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0` | 10.0.201 | 0 | 4 | 4 | 933.0 MB | `sha256:0647d36f93ae` | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0@sha256:0647d36f93ae07e9ca2e9f01a1a09f2f0b7d93cd82fe8d4857dec8bb0359b7e3` |

### Debian

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0` | 8.0.25 | 1 | 9 | 108 | 193.0 MB | `sha256:b473b5d596c0` | `mcr.microsoft.com/dotnet/runtime:8.0@sha256:b473b5d596c09fd55bb41f615623d349712013c1c6f4dfff9db741ac8dff36ec` |
| 2 | `mcr.microsoft.com/dotnet/aspnet:8.0` | 8.0.25 | 1 | 9 | 108 | 218.0 MB | `sha256:d4d80bf500f4` | `mcr.microsoft.com/dotnet/aspnet:8.0@sha256:d4d80bf500f4c8307e5c44bf61eb58aec027da07c4d1c40816846fe5eef3f34d` |
| 3 | `mcr.microsoft.com/dotnet/sdk:8.0` | 8.0.419 | 1 | 20 | 195 | 850.0 MB | `sha256:a9330090b730` | `mcr.microsoft.com/dotnet/sdk:8.0@sha256:a9330090b730c2abde8f3f43b3d1e24e4b8cba028de7bab1a7fdcd50cb7a3b7e` |

### Ubuntu

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0-noble` | 8.0.25 | 0 | 0 | 13 | 193.0 MB | `sha256:10f9ea5f7642` | `mcr.microsoft.com/dotnet/runtime:8.0-noble@sha256:10f9ea5f764239a862eac6b1dca4c77d421f00f9da4fe5c68d5db73d9cd3cc35` |
| 2 | `mcr.microsoft.com/dotnet/runtime:9.0-noble` | 9.0.14 | 0 | 0 | 13 | 198.0 MB | `sha256:f2a1ce9eee20` | `mcr.microsoft.com/dotnet/runtime:9.0-noble@sha256:f2a1ce9eee2097b4ae3ad865bf18a8d304b67e62992a3eb5fc20038b1d92fd7e` |
| 3 | `mcr.microsoft.com/dotnet/runtime:10.0-noble` | 10.0.5 | 0 | 0 | 13 | 203.0 MB | `sha256:6e5657812fa3` | `mcr.microsoft.com/dotnet/runtime:10.0-noble@sha256:6e5657812fa342bfdf34b17d01a644c0a9cc9ff560a1f4d93134f5a8087b5ca8` |
| 4 | `mcr.microsoft.com/dotnet/aspnet:8.0-noble` | 8.0.25 | 0 | 0 | 13 | 217.0 MB | `sha256:f787fd89b89f` | `mcr.microsoft.com/dotnet/aspnet:8.0-noble@sha256:f787fd89b89f72a425873009dd53dcacf5e59c324abd1bc1561e3b5801c16dad` |
| 5 | `mcr.microsoft.com/dotnet/aspnet:9.0-noble` | 9.0.14 | 0 | 0 | 13 | 223.0 MB | `sha256:c05a77cfc605` | `mcr.microsoft.com/dotnet/aspnet:9.0-noble@sha256:c05a77cfc6055ff4c4a07da8a4baaa992189f14f1c02447616fd4dc9e57c96f3` |
| 6 | `mcr.microsoft.com/dotnet/aspnet:10.0-noble` | 10.0.5 | 0 | 0 | 13 | 230.0 MB | `sha256:a04d1c1d2d26` | `mcr.microsoft.com/dotnet/aspnet:10.0-noble@sha256:a04d1c1d2d26119049494057d80ea6cda25bbd8aef7c444a1fc1ef874fd3955b` |
| 7 | `mcr.microsoft.com/dotnet/sdk:9.0-noble` | 9.0.312 | 0 | 0 | 18 | 836.0 MB | `sha256:2630db661965` | `mcr.microsoft.com/dotnet/sdk:9.0-noble@sha256:2630db6619659d2a213872c65b65ff851f866e3b1a6f218472aa8c1b42f9917e` |
| 8 | `mcr.microsoft.com/dotnet/sdk:10.0-noble` | 10.0.201 | 0 | 0 | 18 | 890.0 MB | `sha256:478b9038d187` | `mcr.microsoft.com/dotnet/sdk:10.0-noble@sha256:478b9038d187e5b5c29bfa8173ded5d29e864b5ad06102a12106380ee01e2e49` |
| 9 | `mcr.microsoft.com/dotnet/sdk:8.0-noble` | 8.0.419 | 0 | 8 | 26 | 836.0 MB | `sha256:ad8217434511` | `mcr.microsoft.com/dotnet/sdk:8.0-noble@sha256:ad8217434511581f7a902bb1ff90779e099b176b8880cd15dd97e930be0cb6a3` |

## Java

### Azure Linux

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:11-distroless` | 11.0.30 | 0 | 0 | 0 | 320.0 MB | `sha256:1c017a20039c` | `mcr.microsoft.com/openjdk/jdk:11-distroless@sha256:1c017a20039c7fe3a5695b881cb8d4119fb44c965efdb3baddde5a3878a4636c` |
| 2 | `mcr.microsoft.com/openjdk/jdk:17-distroless` | 17.0.18 | 0 | 0 | 0 | 323.0 MB | `sha256:f151b083f12d` | `mcr.microsoft.com/openjdk/jdk:17-distroless@sha256:f151b083f12d43117040fd0850890b157e6c4db4f2c6c71479812bf0c91bac6f` |
| 3 | `mcr.microsoft.com/openjdk/jdk:21-distroless` | 21.0.10 | 0 | 0 | 0 | 351.0 MB | `sha256:1945f10141b4` | `mcr.microsoft.com/openjdk/jdk:21-distroless@sha256:1945f10141b492d1ef3915b0e80be19c67a32badb398ca2e152b1ccfbdb3a525` |
| 4 | `mcr.microsoft.com/openjdk/jdk:25-distroless` | 25.0.2 | 0 | 0 | 0 | 396.0 MB | `sha256:2df342fbd8e0` | `mcr.microsoft.com/openjdk/jdk:25-distroless@sha256:2df342fbd8e02fe0697962dc42e6305df0d8f57eab668ad2625cc6704aa4bde9` |
| 5 | `mcr.microsoft.com/openjdk/jdk:21-azurelinux` | 21.0.10 | 0 | 2 | 2 | 484.0 MB | `sha256:281f988de5c9` | `mcr.microsoft.com/openjdk/jdk:21-azurelinux@sha256:281f988de5c9e6898f799ee0bcdfefc21500feb8ce20ed98121ac6f783003ce9` |
| 6 | `mcr.microsoft.com/openjdk/jdk:25-azurelinux` | 25.0.2 | 0 | 2 | 2 | 528.0 MB | `sha256:6996e4bdd4ec` | `mcr.microsoft.com/openjdk/jdk:25-azurelinux@sha256:6996e4bdd4ec1957f85145daa50754d89bc4583ddba549c401bb5a9ce1e01333` |

### Ubuntu

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:17-ubuntu` | 17.0.18 | 0 | 0 | 62 | 428.0 MB | `sha256:9dc95b6aa040` | `mcr.microsoft.com/openjdk/jdk:17-ubuntu@sha256:9dc95b6aa040b985f3673dc29db5abc8c5bc293a8bdea5752085d539858479ee` |
| 2 | `mcr.microsoft.com/openjdk/jdk:21-ubuntu` | 21.0.10 | 0 | 0 | 62 | 455.0 MB | `sha256:c8f24edeec38` | `mcr.microsoft.com/openjdk/jdk:21-ubuntu@sha256:c8f24edeec38c3f2b6ffd540ce76feab111c1197e83892cb789e86c94821b2d3` |
| 3 | `mcr.microsoft.com/openjdk/jdk:25-ubuntu` | 25.0.2 | 0 | 0 | 62 | 500.0 MB | `sha256:d7c9c527dbd8` | `mcr.microsoft.com/openjdk/jdk:25-ubuntu@sha256:d7c9c527dbd80511fd2f01caf4ab030ccac7e7e9f28d22897085b20bc244f59a` |

## Go

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0` | 1.25.8 | 0 | 2 | 2 | 810.0 MB | `sha256:67527a2c916c` | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0@sha256:67527a2c916cd5a5622b717d9894724198aee5e5ffee07d209c4d37ab3ae4f75` |
| 2 | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0` | 1.26.1 | 0 | 2 | 2 | 843.0 MB | `sha256:cc01cc579e88` | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0@sha256:cc01cc579e881305ed09a8cd24827e96588643dc2c09758fab97018180fcdf8f` |

## Python

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot` | 3.12.9 | 0 | 2 | 2 | 83.6 MB | `sha256:b422aa585030` | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot@sha256:b422aa5850309ddd74c2be3cd329fb8f3d8318c7106d0dbc96075def3153ae7d` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/python:3.12` | 3.12.9 | 0 | 2 | 2 | 83.6 MB | `sha256:3db9d5eeb703` | `mcr.microsoft.com/azurelinux/distroless/python:3.12@sha256:3db9d5eeb703dc18befbc396d4b4f2b9dd30c0dd82ecd49e12edc50f9b3131ec` |
| 3 | `mcr.microsoft.com/azurelinux/distroless/python:3-nonroot` | 3.12.9 | 0 | 2 | 2 | 83.6 MB | `sha256:b422aa585030` | `mcr.microsoft.com/azurelinux/distroless/python:3-nonroot@sha256:b422aa5850309ddd74c2be3cd329fb8f3d8318c7106d0dbc96075def3153ae7d` |
| 4 | `mcr.microsoft.com/azurelinux/distroless/python:3` | 3.12.9 | 0 | 2 | 2 | 83.6 MB | `sha256:3db9d5eeb703` | `mcr.microsoft.com/azurelinux/distroless/python:3@sha256:3db9d5eeb703dc18befbc396d4b4f2b9dd30c0dd82ecd49e12edc50f9b3131ec` |
| 5 | `mcr.microsoft.com/openjdk/jdk:21-azurelinux` | 3.12.9 | 0 | 2 | 2 | 484.0 MB | `sha256:281f988de5c9` | `mcr.microsoft.com/openjdk/jdk:21-azurelinux@sha256:281f988de5c9e6898f799ee0bcdfefc21500feb8ce20ed98121ac6f783003ce9` |
| 6 | `mcr.microsoft.com/openjdk/jdk:25-azurelinux` | 3.12.9 | 0 | 2 | 2 | 528.0 MB | `sha256:6996e4bdd4ec` | `mcr.microsoft.com/openjdk/jdk:25-azurelinux@sha256:6996e4bdd4ec1957f85145daa50754d89bc4583ddba549c401bb5a9ce1e01333` |
| 7 | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0` | 3.12.9 | 0 | 2 | 2 | 810.0 MB | `sha256:67527a2c916c` | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0@sha256:67527a2c916cd5a5622b717d9894724198aee5e5ffee07d209c4d37ab3ae4f75` |
| 8 | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0` | 3.12.9 | 0 | 2 | 2 | 843.0 MB | `sha256:cc01cc579e88` | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0@sha256:cc01cc579e881305ed09a8cd24827e96588643dc2c09758fab97018180fcdf8f` |
| 9 | `mcr.microsoft.com/azurelinux/base/python:3.12` | 3.12.9 | 0 | 4 | 4 | 139.0 MB | `sha256:8798301c18bf` | `mcr.microsoft.com/azurelinux/base/python:3.12@sha256:8798301c18bf675d2a4c6fe45f01aa2fe8ab469f05e571a07579f63f865a4809` |
| 10 | `mcr.microsoft.com/azurelinux/base/python:3` | 3.12.9 | 0 | 4 | 4 | 139.0 MB | `sha256:8798301c18bf` | `mcr.microsoft.com/azurelinux/base/python:3@sha256:8798301c18bf675d2a4c6fe45f01aa2fe8ab469f05e571a07579f63f865a4809` |

## Node

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/base/nodejs:20.14` | 20.14.0 | 0 | 4 | 4 | 146.0 MB | `sha256:b8a48ba88fdf` | `mcr.microsoft.com/azurelinux/base/nodejs:20.14@sha256:b8a48ba88fdff68c20a2895cb002d547ea47779e37e2fdcb6cd64e4a875f8a71` |
| 2 | `mcr.microsoft.com/azurelinux/base/nodejs:20` | 20.14.0 | 0 | 4 | 4 | 146.0 MB | `sha256:b8a48ba88fdf` | `mcr.microsoft.com/azurelinux/base/nodejs:20@sha256:b8a48ba88fdff68c20a2895cb002d547ea47779e37e2fdcb6cd64e4a875f8a71` |
| 3 | `mcr.microsoft.com/azurelinux/distroless/nodejs:20.14-nonroot` | 20.14.0 | 0 | 13 | 16 | 106.0 MB | `sha256:60773809112a` | `mcr.microsoft.com/azurelinux/distroless/nodejs:20.14-nonroot@sha256:60773809112a278016c356118123621ddf763c80f970faa36e6366bdaf794d50` |
| 4 | `mcr.microsoft.com/azurelinux/distroless/nodejs:20.14` | 20.14.0 | 0 | 13 | 16 | 106.0 MB | `sha256:dfc185c7bb4a` | `mcr.microsoft.com/azurelinux/distroless/nodejs:20.14@sha256:dfc185c7bb4ac4f740b6581ee50d2a5f868dd12c7626b40a1fc5698c980c12e6` |
| 5 | `mcr.microsoft.com/azurelinux/distroless/nodejs:20-nonroot` | 20.14.0 | 0 | 13 | 16 | 106.0 MB | `sha256:60773809112a` | `mcr.microsoft.com/azurelinux/distroless/nodejs:20-nonroot@sha256:60773809112a278016c356118123621ddf763c80f970faa36e6366bdaf794d50` |
| 6 | `mcr.microsoft.com/azurelinux/distroless/nodejs:20` | 20.14.0 | 0 | 13 | 16 | 106.0 MB | `sha256:dfc185c7bb4a` | `mcr.microsoft.com/azurelinux/distroless/nodejs:20@sha256:dfc185c7bb4ac4f740b6581ee50d2a5f868dd12c7626b40a1fc5698c980c12e6` |

## Base / No Runtime

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0` | 3.0 | 0 | 0 | 0 | 3.6 MB | `sha256:22810fd97d6a` | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0@sha256:22810fd97d6ad5ec7d5bdd5b00233a3050be01d9e26b47b16cb6f1a7f178834b` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/base:3.0` | 3.0 | 0 | 0 | 0 | 34.0 MB | `sha256:09b4ed711653` | `mcr.microsoft.com/azurelinux/distroless/base:3.0@sha256:09b4ed711653fc770cbc35c3b9088d4d4c8c835187b01ddf7b2ed88de5fa5366` |
| 3 | `mcr.microsoft.com/azurelinux/base/core:3.0` | 3.0 | 0 | 2 | 2 | 75.3 MB | `sha256:a452d39c9157` | `mcr.microsoft.com/azurelinux/base/core:3.0@sha256:a452d39c91576f5a2c983c7d3b62521fabd08e16b4a7237e24bf2be3b06e1651` |

# Daily Recommended Images by Language

_Generated: 2026-03-20T04:26:02Z. Criteria: lowest critical → high → total vulnerabilities → size. Top 10 per language per base OS._

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
| 6 | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0` | 9.0.312 | 0 | 0 | 0 | 879.0 MB | `sha256:f1c4c2210b8d` | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0@sha256:f1c4c2210b8d225c03ef11d79be475e996daf5646d3606a2cbf03abc95919d99` |
| 7 | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0` | 10.0.201 | 0 | 0 | 0 | 933.0 MB | `sha256:0647d36f93ae` | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0@sha256:0647d36f93ae07e9ca2e9f01a1a09f2f0b7d93cd82fe8d4857dec8bb0359b7e3` |

### Debian

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0` | 8.0.25 | 1 | 2 | 97 | 193.0 MB | `sha256:b473b5d596c0` | `mcr.microsoft.com/dotnet/runtime:8.0@sha256:b473b5d596c09fd55bb41f615623d349712013c1c6f4dfff9db741ac8dff36ec` |
| 2 | `mcr.microsoft.com/dotnet/aspnet:8.0` | 8.0.25 | 1 | 2 | 97 | 218.0 MB | `sha256:d4d80bf500f4` | `mcr.microsoft.com/dotnet/aspnet:8.0@sha256:d4d80bf500f4c8307e5c44bf61eb58aec027da07c4d1c40816846fe5eef3f34d` |
| 3 | `mcr.microsoft.com/dotnet/sdk:8.0` | 8.0.419 | 1 | 16 | 192 | 850.0 MB | `sha256:a9330090b730` | `mcr.microsoft.com/dotnet/sdk:8.0@sha256:a9330090b730c2abde8f3f43b3d1e24e4b8cba028de7bab1a7fdcd50cb7a3b7e` |

### Ubuntu

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0-noble` | 8.0.25 | 0 | 0 | 12 | 193.0 MB | `sha256:10f9ea5f7642` | `mcr.microsoft.com/dotnet/runtime:8.0-noble@sha256:10f9ea5f764239a862eac6b1dca4c77d421f00f9da4fe5c68d5db73d9cd3cc35` |
| 2 | `mcr.microsoft.com/dotnet/runtime:9.0-noble` | 9.0.14 | 0 | 0 | 12 | 198.0 MB | `sha256:f2a1ce9eee20` | `mcr.microsoft.com/dotnet/runtime:9.0-noble@sha256:f2a1ce9eee2097b4ae3ad865bf18a8d304b67e62992a3eb5fc20038b1d92fd7e` |
| 3 | `mcr.microsoft.com/dotnet/runtime:10.0-noble` | 10.0.5 | 0 | 0 | 12 | 203.0 MB | `sha256:6e5657812fa3` | `mcr.microsoft.com/dotnet/runtime:10.0-noble@sha256:6e5657812fa342bfdf34b17d01a644c0a9cc9ff560a1f4d93134f5a8087b5ca8` |
| 4 | `mcr.microsoft.com/dotnet/aspnet:8.0-noble` | 8.0.25 | 0 | 0 | 12 | 217.0 MB | `sha256:f787fd89b89f` | `mcr.microsoft.com/dotnet/aspnet:8.0-noble@sha256:f787fd89b89f72a425873009dd53dcacf5e59c324abd1bc1561e3b5801c16dad` |
| 5 | `mcr.microsoft.com/dotnet/aspnet:9.0-noble` | 9.0.14 | 0 | 0 | 12 | 223.0 MB | `sha256:c05a77cfc605` | `mcr.microsoft.com/dotnet/aspnet:9.0-noble@sha256:c05a77cfc6055ff4c4a07da8a4baaa992189f14f1c02447616fd4dc9e57c96f3` |
| 6 | `mcr.microsoft.com/dotnet/aspnet:10.0-noble` | 10.0.5 | 0 | 0 | 12 | 230.0 MB | `sha256:a04d1c1d2d26` | `mcr.microsoft.com/dotnet/aspnet:10.0-noble@sha256:a04d1c1d2d26119049494057d80ea6cda25bbd8aef7c444a1fc1ef874fd3955b` |
| 7 | `mcr.microsoft.com/dotnet/sdk:9.0-noble` | 9.0.312 | 0 | 0 | 17 | 836.0 MB | `sha256:2630db661965` | `mcr.microsoft.com/dotnet/sdk:9.0-noble@sha256:2630db6619659d2a213872c65b65ff851f866e3b1a6f218472aa8c1b42f9917e` |
| 8 | `mcr.microsoft.com/dotnet/sdk:10.0-noble` | 10.0.201 | 0 | 0 | 17 | 890.0 MB | `sha256:478b9038d187` | `mcr.microsoft.com/dotnet/sdk:10.0-noble@sha256:478b9038d187e5b5c29bfa8173ded5d29e864b5ad06102a12106380ee01e2e49` |
| 9 | `mcr.microsoft.com/dotnet/sdk:8.0-noble` | 8.0.419 | 0 | 8 | 25 | 836.0 MB | `sha256:ad8217434511` | `mcr.microsoft.com/dotnet/sdk:8.0-noble@sha256:ad8217434511581f7a902bb1ff90779e099b176b8880cd15dd97e930be0cb6a3` |

## Go

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0` | 1.25.8 | 0 | 0 | 0 | 809.0 MB | `sha256:c64816802225` | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0@sha256:c64816802225eab888fbc31638f929f9c9441e60b51b814da53c843d4a78f334` |
| 2 | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0` | 1.26.1 | 0 | 0 | 0 | 841.0 MB | `sha256:11adc6e97c00` | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0@sha256:11adc6e97c00c27eadbab59d43dfc3ecd183ac26f3dff86725be612d9d55b6b4` |

## Java

### Azure Linux

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:21-azurelinux` | 21.0.10 | 0 | 0 | 0 | 482.0 MB | `sha256:bb4e06946b22` | `mcr.microsoft.com/openjdk/jdk:21-azurelinux@sha256:bb4e06946b22d3e73a404d68b374e3be0889b93b4cfe1815fbc13302a9593b7d` |
| 2 | `mcr.microsoft.com/openjdk/jdk:25-azurelinux` | 25.0.2 | 0 | 0 | 0 | 527.0 MB | `sha256:e1d716aca7ea` | `mcr.microsoft.com/openjdk/jdk:25-azurelinux@sha256:e1d716aca7ea670ce3433db5b8983252d41841a3093619049065b701c495a63f` |
| 3 | `mcr.microsoft.com/openjdk/jdk:11-distroless` | 11.0.30 | 0 | 1 | 3 | 320.0 MB | `sha256:b51499120f93` | `mcr.microsoft.com/openjdk/jdk:11-distroless@sha256:b51499120f932f6b6c377b69a2bb91732323e4620e3fe4680bfb6c7187c676ae` |
| 4 | `mcr.microsoft.com/openjdk/jdk:17-distroless` | 17.0.18 | 0 | 1 | 3 | 323.0 MB | `sha256:3371073545c9` | `mcr.microsoft.com/openjdk/jdk:17-distroless@sha256:3371073545c937049840c588f84657b5981077d79441ba7af400577141720d9b` |
| 5 | `mcr.microsoft.com/openjdk/jdk:21-distroless` | 21.0.10 | 0 | 1 | 3 | 351.0 MB | `sha256:ac00578dac53` | `mcr.microsoft.com/openjdk/jdk:21-distroless@sha256:ac00578dac535973bfa8e00e24f4a120d4e9349cd49897d65e9c574e3feee5fd` |
| 6 | `mcr.microsoft.com/openjdk/jdk:25-distroless` | 25.0.2 | 0 | 1 | 3 | 396.0 MB | `sha256:4c50e1fde31a` | `mcr.microsoft.com/openjdk/jdk:25-distroless@sha256:4c50e1fde31ac0d3d96d8a364afc4c5afed3ce15d3f6cc7818904c6822b973cd` |

### Ubuntu

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:17-ubuntu` | 17.0.18 | 0 | 1 | 66 | 440.0 MB | `sha256:bb600bfff83d` | `mcr.microsoft.com/openjdk/jdk:17-ubuntu@sha256:bb600bfff83d507c5f285d141e6482c2592a4628d7ade1c12ad318fdad307a48` |
| 2 | `mcr.microsoft.com/openjdk/jdk:21-ubuntu` | 21.0.10 | 0 | 1 | 66 | 467.0 MB | `sha256:3e9650137bd4` | `mcr.microsoft.com/openjdk/jdk:21-ubuntu@sha256:3e9650137bd485e5714422085aec3520cfe79b735e5c8a71590f0fe83d3dcaa1` |
| 3 | `mcr.microsoft.com/openjdk/jdk:25-ubuntu` | 25.0.2 | 0 | 1 | 66 | 512.0 MB | `sha256:2fa6e2fc2610` | `mcr.microsoft.com/openjdk/jdk:25-ubuntu@sha256:2fa6e2fc2610873a78ae82199c8e80a11535cab02b01a1e7e35265f0dca5dbeb` |

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
| 7 | `mcr.microsoft.com/openjdk/jdk:21-azurelinux` | 3.12.9 | 0 | 0 | 0 | 482.0 MB | `sha256:bb4e06946b22` | `mcr.microsoft.com/openjdk/jdk:21-azurelinux@sha256:bb4e06946b22d3e73a404d68b374e3be0889b93b4cfe1815fbc13302a9593b7d` |
| 8 | `mcr.microsoft.com/openjdk/jdk:25-azurelinux` | 3.12.9 | 0 | 0 | 0 | 527.0 MB | `sha256:e1d716aca7ea` | `mcr.microsoft.com/openjdk/jdk:25-azurelinux@sha256:e1d716aca7ea670ce3433db5b8983252d41841a3093619049065b701c495a63f` |
| 9 | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0` | 3.12.9 | 0 | 0 | 0 | 809.0 MB | `sha256:c64816802225` | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0@sha256:c64816802225eab888fbc31638f929f9c9441e60b51b814da53c843d4a78f334` |
| 10 | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0` | 3.12.9 | 0 | 0 | 0 | 841.0 MB | `sha256:11adc6e97c00` | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0@sha256:11adc6e97c00c27eadbab59d43dfc3ecd183ac26f3dff86725be612d9d55b6b4` |

## Base / No Runtime

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0` | 3.0 | 0 | 0 | 0 | 3.6 MB | `sha256:22810fd97d6a` | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0@sha256:22810fd97d6ad5ec7d5bdd5b00233a3050be01d9e26b47b16cb6f1a7f178834b` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/base:3.0` | 3.0 | 0 | 0 | 0 | 34.0 MB | `sha256:09b4ed711653` | `mcr.microsoft.com/azurelinux/distroless/base:3.0@sha256:09b4ed711653fc770cbc35c3b9088d4d4c8c835187b01ddf7b2ed88de5fa5366` |
| 3 | `mcr.microsoft.com/azurelinux/base/core:3.0` | 3.0 | 0 | 0 | 0 | 75.3 MB | `sha256:a452d39c9157` | `mcr.microsoft.com/azurelinux/base/core:3.0@sha256:a452d39c91576f5a2c983c7d3b62521fabd08e16b4a7237e24bf2be3b06e1651` |

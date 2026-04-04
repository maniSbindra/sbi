# Daily Recommended Images by Language

_Generated: 2026-04-04T04:29:31Z. Criteria: lowest critical → high → total vulnerabilities → size. Top 10 per language per base OS._

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
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0` | 8.0.25 | 1 | 7 | 117 | 193.0 MB | `sha256:b473b5d596c0` | `mcr.microsoft.com/dotnet/runtime:8.0@sha256:b473b5d596c09fd55bb41f615623d349712013c1c6f4dfff9db741ac8dff36ec` |
| 2 | `mcr.microsoft.com/dotnet/aspnet:8.0` | 8.0.25 | 1 | 7 | 117 | 218.0 MB | `sha256:d4d80bf500f4` | `mcr.microsoft.com/dotnet/aspnet:8.0@sha256:d4d80bf500f4c8307e5c44bf61eb58aec027da07c4d1c40816846fe5eef3f34d` |
| 3 | `mcr.microsoft.com/dotnet/sdk:8.0` | 8.0.419 | 1 | 18 | 204 | 850.0 MB | `sha256:a9330090b730` | `mcr.microsoft.com/dotnet/sdk:8.0@sha256:a9330090b730c2abde8f3f43b3d1e24e4b8cba028de7bab1a7fdcd50cb7a3b7e` |

### Ubuntu

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0-noble` | 8.0.25 | 0 | 0 | 5 | 193.0 MB | `sha256:919d1e7429df` | `mcr.microsoft.com/dotnet/runtime:8.0-noble@sha256:919d1e7429dfaca99487338465276987870f3b2a06422a8871bcbd3f44bc35d4` |
| 2 | `mcr.microsoft.com/dotnet/runtime:9.0-noble` | 9.0.14 | 0 | 0 | 5 | 198.0 MB | `sha256:12147f660897` | `mcr.microsoft.com/dotnet/runtime:9.0-noble@sha256:12147f66089752f72e215806e352d9b2336cda9b46dc72b9d425dddfecfa7d8b` |
| 3 | `mcr.microsoft.com/dotnet/runtime:10.0-noble` | 10.0.5 | 0 | 0 | 5 | 203.0 MB | `sha256:260c4d478b25` | `mcr.microsoft.com/dotnet/runtime:10.0-noble@sha256:260c4d478b25ccb5a84be74dc5ea5a772fb4fb520e9226ca8f1cb57bf1425fb1` |
| 4 | `mcr.microsoft.com/dotnet/aspnet:8.0-noble` | 8.0.25 | 0 | 0 | 5 | 217.0 MB | `sha256:3c08a6c357ba` | `mcr.microsoft.com/dotnet/aspnet:8.0-noble@sha256:3c08a6c357bac63ba74e23b54db4fd54996662cd2260ca332e69b03393e2e352` |
| 5 | `mcr.microsoft.com/dotnet/aspnet:9.0-noble` | 9.0.14 | 0 | 0 | 5 | 223.0 MB | `sha256:729acff37e08` | `mcr.microsoft.com/dotnet/aspnet:9.0-noble@sha256:729acff37e0859d9a69ad19cb43145b2306743c235397220fd6cb301d059e0e5` |
| 6 | `mcr.microsoft.com/dotnet/aspnet:10.0-noble` | 10.0.5 | 0 | 0 | 5 | 230.0 MB | `sha256:ccdca44cd4f2` | `mcr.microsoft.com/dotnet/aspnet:10.0-noble@sha256:ccdca44cd4f256d50187f920dc8ccc2a9ea7a8a4597ac1d51e08fddb2e3b3205` |
| 7 | `mcr.microsoft.com/dotnet/sdk:9.0-noble` | 9.0.312 | 0 | 0 | 9 | 837.0 MB | `sha256:446c032fc8de` | `mcr.microsoft.com/dotnet/sdk:9.0-noble@sha256:446c032fc8de02409c2aa1f34cbcceffae22bb6ab9a93bf19fa99a578cd066f3` |
| 8 | `mcr.microsoft.com/dotnet/sdk:10.0-noble` | 10.0.201 | 0 | 0 | 9 | 891.0 MB | `sha256:f061e5a7532b` | `mcr.microsoft.com/dotnet/sdk:10.0-noble@sha256:f061e5a7532b36fa1d1b684857fe1f504ba92115b9934f154643266613c44c62` |
| 9 | `mcr.microsoft.com/dotnet/sdk:8.0-noble` | 8.0.419 | 0 | 8 | 17 | 837.0 MB | `sha256:67f7a8d541a1` | `mcr.microsoft.com/dotnet/sdk:8.0-noble@sha256:67f7a8d541a18505a74326085ed997eb4409cb228d429c55d38ccaea3d486dce` |

## Go

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0` | 1.25.8 | 0 | 0 | 0 | 810.0 MB | `sha256:42739e5cb970` | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0@sha256:42739e5cb9703e3e0908d0e393ef9e0acde9d7e2ef09edf7b910eda2ebde8eaf` |
| 2 | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0` | 1.26.1 | 0 | 0 | 0 | 843.0 MB | `sha256:7fbb3de047ac` | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0@sha256:7fbb3de047ac3df460e7f24cacf8ef70eb120b515dce2628e6ad8cc9ea2ce2be` |

## Java

### Azure Linux

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:11-distroless` | 11.0.30 | 0 | 0 | 0 | 320.0 MB | `sha256:7d4ee82a4313` | `mcr.microsoft.com/openjdk/jdk:11-distroless@sha256:7d4ee82a431314885fbd55b6c69238d25d5fbaa11023d9088400a71630e9f39a` |
| 2 | `mcr.microsoft.com/openjdk/jdk:17-distroless` | 17.0.18 | 0 | 0 | 0 | 323.0 MB | `sha256:8d196604357d` | `mcr.microsoft.com/openjdk/jdk:17-distroless@sha256:8d196604357dbc00ae5f7de53ef214b700d6ea0fb3a35c6e51df8a6d9618c435` |
| 3 | `mcr.microsoft.com/openjdk/jdk:21-distroless` | 21.0.10 | 0 | 0 | 0 | 351.0 MB | `sha256:210f90fa3383` | `mcr.microsoft.com/openjdk/jdk:21-distroless@sha256:210f90fa33831e8ba690bd7117477e1c1bc9bd9b8c0a5dfd8ec98cf5b5f3c96e` |
| 4 | `mcr.microsoft.com/openjdk/jdk:25-distroless` | 25.0.2 | 0 | 0 | 0 | 396.0 MB | `sha256:38e1bed79149` | `mcr.microsoft.com/openjdk/jdk:25-distroless@sha256:38e1bed79149c690d77c57b9cb0d03cab4a06cf3197328f511f95ad635d03fe9` |
| 5 | `mcr.microsoft.com/openjdk/jdk:21-azurelinux` | 21.0.10 | 0 | 0 | 0 | 484.0 MB | `sha256:6c41ab33d752` | `mcr.microsoft.com/openjdk/jdk:21-azurelinux@sha256:6c41ab33d7524396dd775562677dda8d3ca65f3d5035b7660e6a263e6571c990` |
| 6 | `mcr.microsoft.com/openjdk/jdk:25-azurelinux` | 25.0.2 | 0 | 0 | 0 | 528.0 MB | `sha256:5bb142763529` | `mcr.microsoft.com/openjdk/jdk:25-azurelinux@sha256:5bb142763529403f7bf0a0437d170ce581df7a4dfa8589904a0cc68e270d5aa8` |

### Ubuntu

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:17-ubuntu` | 17.0.18 | 0 | 0 | 56 | 428.0 MB | `sha256:1c2ad2a2970f` | `mcr.microsoft.com/openjdk/jdk:17-ubuntu@sha256:1c2ad2a2970ff909e3076e60650b1def135570c988fcbc0ca490b0c8ea74eb00` |
| 2 | `mcr.microsoft.com/openjdk/jdk:21-ubuntu` | 21.0.10 | 0 | 0 | 56 | 455.0 MB | `sha256:02d32ccecf91` | `mcr.microsoft.com/openjdk/jdk:21-ubuntu@sha256:02d32ccecf91a79a9677f8dead6e038be597119482246be3e8698d70db6abc23` |
| 3 | `mcr.microsoft.com/openjdk/jdk:25-ubuntu` | 25.0.2 | 0 | 0 | 56 | 500.0 MB | `sha256:f964f484c728` | `mcr.microsoft.com/openjdk/jdk:25-ubuntu@sha256:f964f484c728d63c539d13601ad6751148b634dab35148ead9b089322f6106a2` |

## Python

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:21-azurelinux` | 3.12.9 | 0 | 0 | 0 | 484.0 MB | `sha256:6c41ab33d752` | `mcr.microsoft.com/openjdk/jdk:21-azurelinux@sha256:6c41ab33d7524396dd775562677dda8d3ca65f3d5035b7660e6a263e6571c990` |
| 2 | `mcr.microsoft.com/openjdk/jdk:25-azurelinux` | 3.12.9 | 0 | 0 | 0 | 528.0 MB | `sha256:5bb142763529` | `mcr.microsoft.com/openjdk/jdk:25-azurelinux@sha256:5bb142763529403f7bf0a0437d170ce581df7a4dfa8589904a0cc68e270d5aa8` |
| 3 | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0` | 3.12.9 | 0 | 0 | 0 | 810.0 MB | `sha256:42739e5cb970` | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0@sha256:42739e5cb9703e3e0908d0e393ef9e0acde9d7e2ef09edf7b910eda2ebde8eaf` |
| 4 | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0` | 3.12.9 | 0 | 0 | 0 | 843.0 MB | `sha256:7fbb3de047ac` | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0@sha256:7fbb3de047ac3df460e7f24cacf8ef70eb120b515dce2628e6ad8cc9ea2ce2be` |
| 5 | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot` | 3.12.9 | 0 | 2 | 2 | 83.6 MB | `sha256:b422aa585030` | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot@sha256:b422aa5850309ddd74c2be3cd329fb8f3d8318c7106d0dbc96075def3153ae7d` |
| 6 | `mcr.microsoft.com/azurelinux/distroless/python:3.12` | 3.12.9 | 0 | 2 | 2 | 83.6 MB | `sha256:3db9d5eeb703` | `mcr.microsoft.com/azurelinux/distroless/python:3.12@sha256:3db9d5eeb703dc18befbc396d4b4f2b9dd30c0dd82ecd49e12edc50f9b3131ec` |
| 7 | `mcr.microsoft.com/azurelinux/distroless/python:3-nonroot` | 3.12.9 | 0 | 2 | 2 | 83.6 MB | `sha256:b422aa585030` | `mcr.microsoft.com/azurelinux/distroless/python:3-nonroot@sha256:b422aa5850309ddd74c2be3cd329fb8f3d8318c7106d0dbc96075def3153ae7d` |
| 8 | `mcr.microsoft.com/azurelinux/distroless/python:3` | 3.12.9 | 0 | 2 | 2 | 83.6 MB | `sha256:3db9d5eeb703` | `mcr.microsoft.com/azurelinux/distroless/python:3@sha256:3db9d5eeb703dc18befbc396d4b4f2b9dd30c0dd82ecd49e12edc50f9b3131ec` |
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

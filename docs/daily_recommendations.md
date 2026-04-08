# Daily Recommended Images by Language

_Generated: 2026-04-08T04:49:03Z. Criteria: lowest critical → high → total vulnerabilities → size. Top 10 per language per base OS._

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
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0` | 8.0.25 | 1 | 7 | 135 | 193.0 MB | `sha256:b473b5d596c0` | `mcr.microsoft.com/dotnet/runtime:8.0@sha256:b473b5d596c09fd55bb41f615623d349712013c1c6f4dfff9db741ac8dff36ec` |
| 2 | `mcr.microsoft.com/dotnet/aspnet:8.0` | 8.0.25 | 1 | 7 | 135 | 218.0 MB | `sha256:d4d80bf500f4` | `mcr.microsoft.com/dotnet/aspnet:8.0@sha256:d4d80bf500f4c8307e5c44bf61eb58aec027da07c4d1c40816846fe5eef3f34d` |
| 3 | `mcr.microsoft.com/dotnet/sdk:8.0` | 8.0.419 | 1 | 18 | 222 | 850.0 MB | `sha256:a9330090b730` | `mcr.microsoft.com/dotnet/sdk:8.0@sha256:a9330090b730c2abde8f3f43b3d1e24e4b8cba028de7bab1a7fdcd50cb7a3b7e` |

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

## Java

### Azure Linux

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:21-azurelinux` | 21.0.10 | 0 | 0 | 0 | 484.0 MB | `sha256:f086dc87c4c6` | `mcr.microsoft.com/openjdk/jdk:21-azurelinux@sha256:f086dc87c4c666f755b6a5a6e9cfb46cbf993da3610cbdcba148deddae7b6fee` |
| 2 | `mcr.microsoft.com/openjdk/jdk:25-azurelinux` | 25.0.2 | 0 | 0 | 0 | 528.0 MB | `sha256:c8d1d53d54c6` | `mcr.microsoft.com/openjdk/jdk:25-azurelinux@sha256:c8d1d53d54c6795df81e6b532814a5da670783ed7c2bf5d1a048b5fca86c2c20` |
| 3 | `mcr.microsoft.com/openjdk/jdk:11-distroless` | 11.0.30 | 0 | 0 | 6 | 320.0 MB | `sha256:d11164660516` | `mcr.microsoft.com/openjdk/jdk:11-distroless@sha256:d11164660516264e41adf402516e114cd3fe7050ba4c3acff5008c93d199b44c` |
| 4 | `mcr.microsoft.com/openjdk/jdk:17-distroless` | 17.0.18 | 0 | 0 | 6 | 323.0 MB | `sha256:12cb01c7fe6b` | `mcr.microsoft.com/openjdk/jdk:17-distroless@sha256:12cb01c7fe6b8721a5d57499d958e5fb823107c846a3d7a044f142592e62a339` |
| 5 | `mcr.microsoft.com/openjdk/jdk:21-distroless` | 21.0.10 | 0 | 0 | 6 | 351.0 MB | `sha256:2f4ebc8cf9f5` | `mcr.microsoft.com/openjdk/jdk:21-distroless@sha256:2f4ebc8cf9f53281b32ea11f7ad7aa26e3f4c354c02353adf791719c4d097a3b` |
| 6 | `mcr.microsoft.com/openjdk/jdk:25-distroless` | 25.0.2 | 0 | 0 | 6 | 396.0 MB | `sha256:5f69d762c406` | `mcr.microsoft.com/openjdk/jdk:25-distroless@sha256:5f69d762c406658b2413d62c5bf67565cbe7b998e5756556b132340b6d781a19` |

### Ubuntu

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:17-ubuntu` | 17.0.18 | 0 | 0 | 62 | 428.0 MB | `sha256:aba01212eada` | `mcr.microsoft.com/openjdk/jdk:17-ubuntu@sha256:aba01212eada70b9a373c6ee5bbec2753f612887a375310c43baebb8d5f2fdd8` |
| 2 | `mcr.microsoft.com/openjdk/jdk:21-ubuntu` | 21.0.10 | 0 | 0 | 62 | 455.0 MB | `sha256:2ff286b36bfe` | `mcr.microsoft.com/openjdk/jdk:21-ubuntu@sha256:2ff286b36bfe404eb152d729011c0c8c99b3cfb3acbede820bb89021b9baaf84` |
| 3 | `mcr.microsoft.com/openjdk/jdk:25-ubuntu` | 25.0.2 | 0 | 0 | 62 | 500.0 MB | `sha256:242b476cf2f2` | `mcr.microsoft.com/openjdk/jdk:25-ubuntu@sha256:242b476cf2f2af97440600922b00b7143ee134b2b0bd7402135d17eb95ceee3d` |

## Node

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/base/nodejs:24.13` | 24.13.0 | 0 | 0 | 0 | 163.0 MB | `sha256:2d1ed7ab9e4b` | `mcr.microsoft.com/azurelinux/base/nodejs:24.13@sha256:2d1ed7ab9e4b143703eefbbc1ce5d03bd4475da0a335f157e51b313a1948414f` |
| 2 | `mcr.microsoft.com/azurelinux/base/nodejs:24` | 24.13.0 | 0 | 0 | 0 | 163.0 MB | `sha256:2d1ed7ab9e4b` | `mcr.microsoft.com/azurelinux/base/nodejs:24@sha256:2d1ed7ab9e4b143703eefbbc1ce5d03bd4475da0a335f157e51b313a1948414f` |
| 3 | `mcr.microsoft.com/azurelinux/base/nodejs:20.14` | 20.14.0 | 0 | 4 | 4 | 146.0 MB | `sha256:b8a48ba88fdf` | `mcr.microsoft.com/azurelinux/base/nodejs:20.14@sha256:b8a48ba88fdff68c20a2895cb002d547ea47779e37e2fdcb6cd64e4a875f8a71` |
| 4 | `mcr.microsoft.com/azurelinux/base/nodejs:20` | 20.14.0 | 0 | 4 | 4 | 146.0 MB | `sha256:b8a48ba88fdf` | `mcr.microsoft.com/azurelinux/base/nodejs:20@sha256:b8a48ba88fdff68c20a2895cb002d547ea47779e37e2fdcb6cd64e4a875f8a71` |
| 5 | `mcr.microsoft.com/azurelinux/distroless/nodejs:20.14-nonroot` | 20.14.0 | 0 | 13 | 16 | 106.0 MB | `sha256:60773809112a` | `mcr.microsoft.com/azurelinux/distroless/nodejs:20.14-nonroot@sha256:60773809112a278016c356118123621ddf763c80f970faa36e6366bdaf794d50` |
| 6 | `mcr.microsoft.com/azurelinux/distroless/nodejs:20.14` | 20.14.0 | 0 | 13 | 16 | 106.0 MB | `sha256:dfc185c7bb4a` | `mcr.microsoft.com/azurelinux/distroless/nodejs:20.14@sha256:dfc185c7bb4ac4f740b6581ee50d2a5f868dd12c7626b40a1fc5698c980c12e6` |
| 7 | `mcr.microsoft.com/azurelinux/distroless/nodejs:20-nonroot` | 20.14.0 | 0 | 13 | 16 | 106.0 MB | `sha256:60773809112a` | `mcr.microsoft.com/azurelinux/distroless/nodejs:20-nonroot@sha256:60773809112a278016c356118123621ddf763c80f970faa36e6366bdaf794d50` |
| 8 | `mcr.microsoft.com/azurelinux/distroless/nodejs:20` | 20.14.0 | 0 | 13 | 16 | 106.0 MB | `sha256:dfc185c7bb4a` | `mcr.microsoft.com/azurelinux/distroless/nodejs:20@sha256:dfc185c7bb4ac4f740b6581ee50d2a5f868dd12c7626b40a1fc5698c980c12e6` |
| 9 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.13-nonroot` | 24.13.0 | 0 | 19 | 23 | 123.0 MB | `sha256:e25fe7994e0a` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.13-nonroot@sha256:e25fe7994e0a31cf3d6523d8b53a651f7390cfafd83cfe1e8706bc80f6988d34` |
| 10 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.13` | 24.13.0 | 0 | 19 | 23 | 123.0 MB | `sha256:9602d6864048` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.13@sha256:9602d686404823a802b0a858f5b65793ae2d6bd8900bcd64bb423d71f7897501` |

## Python

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot` | 3.12.9 | 0 | 0 | 0 | 83.7 MB | `sha256:ddef666511ee` | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot@sha256:ddef666511eefea84adc3654905d55619eddfb89fe045bbcadc385357a5883c6` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/python:3.12` | 3.12.9 | 0 | 0 | 0 | 83.7 MB | `sha256:2cb276ad2888` | `mcr.microsoft.com/azurelinux/distroless/python:3.12@sha256:2cb276ad288899c783733abc4f0932814ce1d9c8658effef4782ce3b4c1a8c49` |
| 3 | `mcr.microsoft.com/azurelinux/distroless/python:3-nonroot` | 3.12.9 | 0 | 0 | 0 | 83.7 MB | `sha256:ddef666511ee` | `mcr.microsoft.com/azurelinux/distroless/python:3-nonroot@sha256:ddef666511eefea84adc3654905d55619eddfb89fe045bbcadc385357a5883c6` |
| 4 | `mcr.microsoft.com/azurelinux/distroless/python:3` | 3.12.9 | 0 | 0 | 0 | 83.7 MB | `sha256:2cb276ad2888` | `mcr.microsoft.com/azurelinux/distroless/python:3@sha256:2cb276ad288899c783733abc4f0932814ce1d9c8658effef4782ce3b4c1a8c49` |
| 5 | `mcr.microsoft.com/azurelinux/base/python:3.12` | 3.12.9 | 0 | 0 | 0 | 139.0 MB | `sha256:8300706a0e64` | `mcr.microsoft.com/azurelinux/base/python:3.12@sha256:8300706a0e644a7e260b6a464a1ae3967181ee9057a4accceaaf0ecdedb1d4a2` |
| 6 | `mcr.microsoft.com/azurelinux/base/python:3` | 3.12.9 | 0 | 0 | 0 | 139.0 MB | `sha256:8300706a0e64` | `mcr.microsoft.com/azurelinux/base/python:3@sha256:8300706a0e644a7e260b6a464a1ae3967181ee9057a4accceaaf0ecdedb1d4a2` |
| 7 | `mcr.microsoft.com/openjdk/jdk:21-azurelinux` | 3.12.9 | 0 | 0 | 0 | 484.0 MB | `sha256:f086dc87c4c6` | `mcr.microsoft.com/openjdk/jdk:21-azurelinux@sha256:f086dc87c4c666f755b6a5a6e9cfb46cbf993da3610cbdcba148deddae7b6fee` |
| 8 | `mcr.microsoft.com/openjdk/jdk:25-azurelinux` | 3.12.9 | 0 | 0 | 0 | 528.0 MB | `sha256:c8d1d53d54c6` | `mcr.microsoft.com/openjdk/jdk:25-azurelinux@sha256:c8d1d53d54c6795df81e6b532814a5da670783ed7c2bf5d1a048b5fca86c2c20` |
| 9 | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0` | 3.12.9 | 0 | 0 | 54 | 810.0 MB | `sha256:42739e5cb970` | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0@sha256:42739e5cb9703e3e0908d0e393ef9e0acde9d7e2ef09edf7b910eda2ebde8eaf` |
| 10 | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0` | 3.12.9 | 0 | 0 | 70 | 843.0 MB | `sha256:7fbb3de047ac` | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0@sha256:7fbb3de047ac3df460e7f24cacf8ef70eb120b515dce2628e6ad8cc9ea2ce2be` |

## Go

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0` | 1.25.8 | 0 | 0 | 54 | 810.0 MB | `sha256:42739e5cb970` | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0@sha256:42739e5cb9703e3e0908d0e393ef9e0acde9d7e2ef09edf7b910eda2ebde8eaf` |
| 2 | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0` | 1.26.1 | 0 | 0 | 70 | 843.0 MB | `sha256:7fbb3de047ac` | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0@sha256:7fbb3de047ac3df460e7f24cacf8ef70eb120b515dce2628e6ad8cc9ea2ce2be` |

## Base / No Runtime

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0` | 3.0 | 0 | 0 | 0 | 3.7 MB | `sha256:5a66f9f16ac6` | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0@sha256:5a66f9f16ac675db2a8229dac72d83811b73b502d6ad192d8b374c7f3be498af` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/base:3.0` | 3.0 | 0 | 0 | 0 | 34.1 MB | `sha256:32820d2cf20e` | `mcr.microsoft.com/azurelinux/distroless/base:3.0@sha256:32820d2cf20e896aa9111742dd683dd0ccff370f742e256889bb3bb50320c0d4` |
| 3 | `mcr.microsoft.com/azurelinux/base/core:3.0` | 3.0 | 0 | 0 | 0 | 75.3 MB | `sha256:35149ae8dd17` | `mcr.microsoft.com/azurelinux/base/core:3.0@sha256:35149ae8dd179684f969944f54a337c665a64e702486154eb44253fb39c2505b` |

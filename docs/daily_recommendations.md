# Daily Recommended Images by Language

_Generated: 2026-04-19T05:02:03Z. Criteria: lowest critical → high → total vulnerabilities → size. Top 10 per language per base OS._

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
| 1 | `mcr.microsoft.com/dotnet/runtime:9.0-azurelinux3.0-distroless` | 9.0.15 | 0 | 0 | 0 | 107.0 MB | `sha256:f57b292ed7d6` | `mcr.microsoft.com/dotnet/runtime:9.0-azurelinux3.0-distroless@sha256:f57b292ed7d6f7c88d3b53addfcc02e520aae273eddb57cf0ea49f9396cb2424` |
| 2 | `mcr.microsoft.com/dotnet/runtime:10.0-azurelinux3.0-distroless` | 10.0.6 | 0 | 0 | 0 | 112.0 MB | `sha256:ad413e07e19e` | `mcr.microsoft.com/dotnet/runtime:10.0-azurelinux3.0-distroless@sha256:ad413e07e19e0223a3cf63ee12c9ccea84d252f3216348213e3be168ca9a589f` |
| 3 | `mcr.microsoft.com/dotnet/aspnet:8.0-azurelinux3.0-distroless` | 8.0.26 | 0 | 0 | 0 | 126.0 MB | `sha256:742fe85d9ef9` | `mcr.microsoft.com/dotnet/aspnet:8.0-azurelinux3.0-distroless@sha256:742fe85d9ef99cd3d43fc11d4e97e1fb3fe5d30d02d183c7d52aa6238f05904c` |
| 4 | `mcr.microsoft.com/dotnet/aspnet:9.0-azurelinux3.0-distroless` | 9.0.15 | 0 | 0 | 0 | 132.0 MB | `sha256:22b1bd549a20` | `mcr.microsoft.com/dotnet/aspnet:9.0-azurelinux3.0-distroless@sha256:22b1bd549a200249bea512b6b2fce8885a47d98c346d0d33511b2bf70653d207` |
| 5 | `mcr.microsoft.com/dotnet/aspnet:10.0-azurelinux3.0-distroless` | 10.0.6 | 0 | 0 | 0 | 139.0 MB | `sha256:f51debfacc77` | `mcr.microsoft.com/dotnet/aspnet:10.0-azurelinux3.0-distroless@sha256:f51debfacc772b1a39f90c07cfefabc5b5026d66f17b91bd5a6b97c815d595c4` |
| 6 | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0` | 9.0.313 | 0 | 2 | 12 | 888.0 MB | `sha256:13b64ec155e9` | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0@sha256:13b64ec155e9aefa220263957b6d86f1add651106fc5df3cbecacb7a74eb3b06` |
| 7 | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0` | 10.0.202 | 0 | 2 | 14 | 943.0 MB | `sha256:5d0b4af94fe2` | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0@sha256:5d0b4af94fe23bcdc6e5df135d460a2fee03aec6979e63c669d2bef399327184` |

### Debian

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0` | 8.0.26 | 1 | 7 | 125 | 193.0 MB | `sha256:d0a2313ac5fb` | `mcr.microsoft.com/dotnet/runtime:8.0@sha256:d0a2313ac5fbf1dfda2bf57e63a431010ae030dc3b9d7fc0fd1cff163b49b0e6` |
| 2 | `mcr.microsoft.com/dotnet/aspnet:8.0` | 8.0.26 | 1 | 7 | 125 | 218.0 MB | `sha256:8c7de8ca26ba` | `mcr.microsoft.com/dotnet/aspnet:8.0@sha256:8c7de8ca26baa56530cdc3d55b5563420ee2a302619ad808c53ec64613143771` |
| 3 | `mcr.microsoft.com/dotnet/sdk:8.0` | 8.0.420 | 1 | 22 | 226 | 850.0 MB | `sha256:ec47373b40fc` | `mcr.microsoft.com/dotnet/sdk:8.0@sha256:ec47373b40fc22a31c891c5be8df83f8cbadc24e785b10899545e33a1a1c2183` |

### Ubuntu

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0-noble` | 8.0.26 | 0 | 0 | 5 | 193.0 MB | `sha256:f1994beca66a` | `mcr.microsoft.com/dotnet/runtime:8.0-noble@sha256:f1994beca66a16766311f5234fbfc9a0ce6efb1fa400fa0a3cd95eb1fd6f2a71` |
| 2 | `mcr.microsoft.com/dotnet/runtime:9.0-noble` | 9.0.15 | 0 | 0 | 5 | 198.0 MB | `sha256:08ec9b342063` | `mcr.microsoft.com/dotnet/runtime:9.0-noble@sha256:08ec9b342063d70cf18e7ee10a3f95f22180be321af77e03a8ed2151093ef65f` |
| 3 | `mcr.microsoft.com/dotnet/runtime:10.0-noble` | 10.0.6 | 0 | 0 | 5 | 203.0 MB | `sha256:3cba07eb3d4d` | `mcr.microsoft.com/dotnet/runtime:10.0-noble@sha256:3cba07eb3d4d656c24e4767ed4bc0dc4734e2bc35d3dc2d691df58847a90ad4d` |
| 4 | `mcr.microsoft.com/dotnet/aspnet:8.0-noble` | 8.0.26 | 0 | 0 | 5 | 217.0 MB | `sha256:b0811426a260` | `mcr.microsoft.com/dotnet/aspnet:8.0-noble@sha256:b0811426a2602ce0be3d1f605257c601b38c3a26e7b49283c9330de01c67cefd` |
| 5 | `mcr.microsoft.com/dotnet/aspnet:9.0-noble` | 9.0.15 | 0 | 0 | 5 | 223.0 MB | `sha256:4180d760569f` | `mcr.microsoft.com/dotnet/aspnet:9.0-noble@sha256:4180d760569fdf8f639cc79d62de294356c166eeebe3e4ab19d4f012960b5be8` |
| 6 | `mcr.microsoft.com/dotnet/aspnet:10.0-noble` | 10.0.6 | 0 | 0 | 5 | 230.0 MB | `sha256:9271888dde1f` | `mcr.microsoft.com/dotnet/aspnet:10.0-noble@sha256:9271888dde1f4408dfb7ce75bc0f513c903d20ff2f1287ab153b641d4588ec7d` |
| 7 | `mcr.microsoft.com/dotnet/sdk:9.0-noble` | 9.0.313 | 0 | 2 | 21 | 837.0 MB | `sha256:c4d10394421c` | `mcr.microsoft.com/dotnet/sdk:9.0-noble@sha256:c4d10394421c0e7c116b6f8866b7c80fa74352ea294629cda1fe4f2acd8f8f6f` |
| 8 | `mcr.microsoft.com/dotnet/sdk:10.0-noble` | 10.0.202 | 0 | 2 | 23 | 892.0 MB | `sha256:adc02be8b879` | `mcr.microsoft.com/dotnet/sdk:10.0-noble@sha256:adc02be8b87957d07208a4a3e51775935b33bad3317de8c45b1e67357b4c073b` |
| 9 | `mcr.microsoft.com/dotnet/sdk:8.0-noble` | 8.0.420 | 0 | 12 | 31 | 836.0 MB | `sha256:a91a987babf7` | `mcr.microsoft.com/dotnet/sdk:8.0-noble@sha256:a91a987babf792bfe535ad5d85e94c492f66ffc2efc25aea343432a5bd606e93` |

## Go

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0` | 1.25.9 | 0 | 0 | 0 | 818.0 MB | `sha256:39b419f2bf42` | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0@sha256:39b419f2bf4203008aced1cb7195b96558fa216f15874e15c4f3fa974555e498` |
| 2 | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0` | 1.26.2 | 0 | 0 | 0 | 850.0 MB | `sha256:767bcb01a5e1` | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0@sha256:767bcb01a5e1cf0a5bd261f7be24c383f788e241f94a9b6f8003984996790905` |

## Java

### Azure Linux

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:21-azurelinux` | 21.0.10 | 0 | 0 | 0 | 484.0 MB | `sha256:4ef48cfb80ec` | `mcr.microsoft.com/openjdk/jdk:21-azurelinux@sha256:4ef48cfb80ec0f75b07aed715c7bf78ef1a8505e093a5ed7398bf799462992ad` |
| 2 | `mcr.microsoft.com/openjdk/jdk:25-azurelinux` | 25.0.2 | 0 | 0 | 0 | 529.0 MB | `sha256:66ea6780e6ba` | `mcr.microsoft.com/openjdk/jdk:25-azurelinux@sha256:66ea6780e6bae0e6b8e55308a835461eeba7059e6e5dc00e0d84b8b85cd2630b` |
| 3 | `mcr.microsoft.com/openjdk/jdk:11-distroless` | 11.0.30 | 0 | 0 | 10 | 320.0 MB | `sha256:f52f500b831d` | `mcr.microsoft.com/openjdk/jdk:11-distroless@sha256:f52f500b831dce00e08b82fdafa95b0103c3983b1fe0ba3d4fa8dd52e73232fa` |
| 4 | `mcr.microsoft.com/openjdk/jdk:17-distroless` | 17.0.18 | 0 | 0 | 10 | 324.0 MB | `sha256:3c704fee447b` | `mcr.microsoft.com/openjdk/jdk:17-distroless@sha256:3c704fee447bba6696c096a87c4821c864b61cfa2e15c8eb4f10efbe97c36a71` |
| 5 | `mcr.microsoft.com/openjdk/jdk:21-distroless` | 21.0.10 | 0 | 0 | 10 | 351.0 MB | `sha256:065acded60ff` | `mcr.microsoft.com/openjdk/jdk:21-distroless@sha256:065acded60ffac2a532ba091388b3a60691bb855ce395456020756247fd6efa6` |
| 6 | `mcr.microsoft.com/openjdk/jdk:25-distroless` | 25.0.2 | 0 | 0 | 10 | 396.0 MB | `sha256:7ac3e73a174c` | `mcr.microsoft.com/openjdk/jdk:25-distroless@sha256:7ac3e73a174c6e8c7af456dc31667f458f8e2be3351497feb2ecd6e6cd01e52b` |

### Ubuntu

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:17-ubuntu` | 17.0.18 | 0 | 0 | 74 | 427.0 MB | `sha256:995efb386928` | `mcr.microsoft.com/openjdk/jdk:17-ubuntu@sha256:995efb386928e143ce9ee3d354050e88244c432e356214276fa07118757d17dd` |
| 2 | `mcr.microsoft.com/openjdk/jdk:21-ubuntu` | 21.0.10 | 0 | 0 | 74 | 454.0 MB | `sha256:2431a93da3a8` | `mcr.microsoft.com/openjdk/jdk:21-ubuntu@sha256:2431a93da3a825c69fc082ceffbe87b2b1ba4df29325d5342a0e2c5dec0af3db` |
| 3 | `mcr.microsoft.com/openjdk/jdk:25-ubuntu` | 25.0.2 | 0 | 0 | 74 | 499.0 MB | `sha256:49f46d95119f` | `mcr.microsoft.com/openjdk/jdk:25-ubuntu@sha256:49f46d95119ff7baac3a92bf1e7adfe87b104ddb39f9523c901e33cb902b6599` |

## Python

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:21-azurelinux` | 3.12.9 | 0 | 0 | 0 | 484.0 MB | `sha256:4ef48cfb80ec` | `mcr.microsoft.com/openjdk/jdk:21-azurelinux@sha256:4ef48cfb80ec0f75b07aed715c7bf78ef1a8505e093a5ed7398bf799462992ad` |
| 2 | `mcr.microsoft.com/openjdk/jdk:25-azurelinux` | 3.12.9 | 0 | 0 | 0 | 529.0 MB | `sha256:66ea6780e6ba` | `mcr.microsoft.com/openjdk/jdk:25-azurelinux@sha256:66ea6780e6bae0e6b8e55308a835461eeba7059e6e5dc00e0d84b8b85cd2630b` |
| 3 | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0` | 3.12.9 | 0 | 0 | 0 | 818.0 MB | `sha256:39b419f2bf42` | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0@sha256:39b419f2bf4203008aced1cb7195b96558fa216f15874e15c4f3fa974555e498` |
| 4 | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0` | 3.12.9 | 0 | 0 | 0 | 850.0 MB | `sha256:767bcb01a5e1` | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0@sha256:767bcb01a5e1cf0a5bd261f7be24c383f788e241f94a9b6f8003984996790905` |
| 5 | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot` | 3.12.9 | 0 | 0 | 10 | 83.7 MB | `sha256:ddef666511ee` | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot@sha256:ddef666511eefea84adc3654905d55619eddfb89fe045bbcadc385357a5883c6` |
| 6 | `mcr.microsoft.com/azurelinux/distroless/python:3.12` | 3.12.9 | 0 | 0 | 10 | 83.7 MB | `sha256:2cb276ad2888` | `mcr.microsoft.com/azurelinux/distroless/python:3.12@sha256:2cb276ad288899c783733abc4f0932814ce1d9c8658effef4782ce3b4c1a8c49` |
| 7 | `mcr.microsoft.com/azurelinux/distroless/python:3-nonroot` | 3.12.9 | 0 | 0 | 10 | 83.7 MB | `sha256:ddef666511ee` | `mcr.microsoft.com/azurelinux/distroless/python:3-nonroot@sha256:ddef666511eefea84adc3654905d55619eddfb89fe045bbcadc385357a5883c6` |
| 8 | `mcr.microsoft.com/azurelinux/distroless/python:3` | 3.12.9 | 0 | 0 | 10 | 83.7 MB | `sha256:2cb276ad2888` | `mcr.microsoft.com/azurelinux/distroless/python:3@sha256:2cb276ad288899c783733abc4f0932814ce1d9c8658effef4782ce3b4c1a8c49` |
| 9 | `mcr.microsoft.com/azurelinux/base/python:3.12` | 3.12.9 | 0 | 0 | 10 | 139.0 MB | `sha256:8300706a0e64` | `mcr.microsoft.com/azurelinux/base/python:3.12@sha256:8300706a0e644a7e260b6a464a1ae3967181ee9057a4accceaaf0ecdedb1d4a2` |
| 10 | `mcr.microsoft.com/azurelinux/base/python:3` | 3.12.9 | 0 | 0 | 10 | 139.0 MB | `sha256:8300706a0e64` | `mcr.microsoft.com/azurelinux/base/python:3@sha256:8300706a0e644a7e260b6a464a1ae3967181ee9057a4accceaaf0ecdedb1d4a2` |

## Node

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/base/nodejs:24.13` | 24.13.0 | 0 | 0 | 10 | 163.0 MB | `sha256:2d1ed7ab9e4b` | `mcr.microsoft.com/azurelinux/base/nodejs:24.13@sha256:2d1ed7ab9e4b143703eefbbc1ce5d03bd4475da0a335f157e51b313a1948414f` |
| 2 | `mcr.microsoft.com/azurelinux/base/nodejs:24` | 24.13.0 | 0 | 0 | 10 | 163.0 MB | `sha256:2d1ed7ab9e4b` | `mcr.microsoft.com/azurelinux/base/nodejs:24@sha256:2d1ed7ab9e4b143703eefbbc1ce5d03bd4475da0a335f157e51b313a1948414f` |
| 3 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.13-nonroot` | 24.13.0 | 0 | 19 | 33 | 123.0 MB | `sha256:e25fe7994e0a` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.13-nonroot@sha256:e25fe7994e0a31cf3d6523d8b53a651f7390cfafd83cfe1e8706bc80f6988d34` |
| 4 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.13` | 24.13.0 | 0 | 19 | 33 | 123.0 MB | `sha256:9602d6864048` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24.13@sha256:9602d686404823a802b0a858f5b65793ae2d6bd8900bcd64bb423d71f7897501` |
| 5 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24-nonroot` | 24.13.0 | 0 | 19 | 33 | 123.0 MB | `sha256:e25fe7994e0a` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24-nonroot@sha256:e25fe7994e0a31cf3d6523d8b53a651f7390cfafd83cfe1e8706bc80f6988d34` |
| 6 | `mcr.microsoft.com/azurelinux/distroless/nodejs:24` | 24.13.0 | 0 | 19 | 33 | 123.0 MB | `sha256:9602d6864048` | `mcr.microsoft.com/azurelinux/distroless/nodejs:24@sha256:9602d686404823a802b0a858f5b65793ae2d6bd8900bcd64bb423d71f7897501` |
| 7 | `mcr.microsoft.com/azurelinux/base/nodejs:20.14` | 20.14.0 | 1 | 6 | 28 | 146.0 MB | `sha256:b8a48ba88fdf` | `mcr.microsoft.com/azurelinux/base/nodejs:20.14@sha256:b8a48ba88fdff68c20a2895cb002d547ea47779e37e2fdcb6cd64e4a875f8a71` |
| 8 | `mcr.microsoft.com/azurelinux/base/nodejs:20` | 20.14.0 | 1 | 6 | 28 | 146.0 MB | `sha256:b8a48ba88fdf` | `mcr.microsoft.com/azurelinux/base/nodejs:20@sha256:b8a48ba88fdff68c20a2895cb002d547ea47779e37e2fdcb6cd64e4a875f8a71` |
| 9 | `mcr.microsoft.com/azurelinux/distroless/nodejs:20.14-nonroot` | 20.14.0 | 1 | 15 | 42 | 106.0 MB | `sha256:60773809112a` | `mcr.microsoft.com/azurelinux/distroless/nodejs:20.14-nonroot@sha256:60773809112a278016c356118123621ddf763c80f970faa36e6366bdaf794d50` |
| 10 | `mcr.microsoft.com/azurelinux/distroless/nodejs:20.14` | 20.14.0 | 1 | 15 | 42 | 106.0 MB | `sha256:dfc185c7bb4a` | `mcr.microsoft.com/azurelinux/distroless/nodejs:20.14@sha256:dfc185c7bb4ac4f740b6581ee50d2a5f868dd12c7626b40a1fc5698c980c12e6` |

## Base / No Runtime

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0` | 3.0 | 0 | 0 | 0 | 3.7 MB | `sha256:5a66f9f16ac6` | `mcr.microsoft.com/azurelinux/distroless/minimal:3.0@sha256:5a66f9f16ac675db2a8229dac72d83811b73b502d6ad192d8b374c7f3be498af` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/base:3.0` | 3.0 | 0 | 0 | 10 | 34.1 MB | `sha256:32820d2cf20e` | `mcr.microsoft.com/azurelinux/distroless/base:3.0@sha256:32820d2cf20e896aa9111742dd683dd0ccff370f742e256889bb3bb50320c0d4` |
| 3 | `mcr.microsoft.com/azurelinux/base/core:3.0` | 3.0 | 0 | 0 | 10 | 75.3 MB | `sha256:35149ae8dd17` | `mcr.microsoft.com/azurelinux/base/core:3.0@sha256:35149ae8dd179684f969944f54a337c665a64e702486154eb44253fb39c2505b` |

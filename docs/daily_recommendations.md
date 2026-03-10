# Daily Recommended Images by Language

_Generated: 2026-03-10T16:05:01Z. Criteria: lowest critical → high → total vulnerabilities → size. Top 10 per language per base OS._

**Note:** Image sizes are based on Linux amd64 platform as reported by `docker images` on GitHub runners. Actual sizes may vary significantly on other platforms (macOS, Windows, etc.).

## Scanned Repositories and Images

This report includes analysis from **37 configured sources** across 7 groups (see [repositories.json](../config/repositories.json)):

**Azure Linux base images:**

- `azurelinux/base/core`
- `azurelinux/base/python`
- `azurelinux/base/nodejs`

**Azure Linux distroless images:**

- `azurelinux/distroless/base`
- `azurelinux/distroless/minimal`
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
| 1 | `mcr.microsoft.com/dotnet/runtime:9.0-azurelinux3.0-distroless` | 9.0.13 | 0 | 0 | 0 | 105.0 MB | `sha256:3995507d446e` | `mcr.microsoft.com/dotnet/runtime:9.0-azurelinux3.0-distroless@sha256:3995507d446ee3ce1fd278962e7a6749495811edce2519ecee31274030113afa` |
| 2 | `mcr.microsoft.com/dotnet/runtime:10.0-azurelinux3.0-distroless` | 10.0.3 | 0 | 0 | 0 | 110.0 MB | `sha256:9d28d4f5ccb4` | `mcr.microsoft.com/dotnet/runtime:10.0-azurelinux3.0-distroless@sha256:9d28d4f5ccb4f54ceddfb54597a2bb991a93a07f3e2f6a61845e3e98ec040892` |
| 3 | `mcr.microsoft.com/dotnet/aspnet:8.0-azurelinux3.0-distroless` | 8.0.24 | 0 | 0 | 0 | 124.0 MB | `sha256:e91f943b47d5` | `mcr.microsoft.com/dotnet/aspnet:8.0-azurelinux3.0-distroless@sha256:e91f943b47d5c2e2d6ac7cdd979e414c82aeb11980c8ef52e409fcf2afa63913` |
| 4 | `mcr.microsoft.com/dotnet/aspnet:9.0-azurelinux3.0-distroless` | 9.0.13 | 0 | 0 | 0 | 130.0 MB | `sha256:9f856e8de88e` | `mcr.microsoft.com/dotnet/aspnet:9.0-azurelinux3.0-distroless@sha256:9f856e8de88e2ca03687bca28e8575f2e3255ea43df1896ffbd76023e7b07f41` |
| 5 | `mcr.microsoft.com/dotnet/aspnet:10.0-azurelinux3.0-distroless` | 10.0.3 | 0 | 0 | 0 | 137.0 MB | `sha256:973ac891bc21` | `mcr.microsoft.com/dotnet/aspnet:10.0-azurelinux3.0-distroless@sha256:973ac891bc21916cb4f579ed3cd5737fac0a1452d30b11a25493df65eefd4786` |
| 6 | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0` | 9.0.311 | 0 | 0 | 0 | 879.0 MB | `sha256:ef9d3354c106` | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0@sha256:ef9d3354c106e4efa0946c79dca08038c28cd03da2eef678d6b9a640028218fe` |
| 7 | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0` | 10.0.103 | 0 | 0 | 0 | 928.0 MB | `sha256:774b85a4de4b` | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0@sha256:774b85a4de4b63b5974d0a599e7166b0ff20f3848d3cc67333beb733405d027b` |

### Debian

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0` | 8.0.24 | 1 | 2 | 95 | 193.0 MB | `sha256:d304745fcb83` | `mcr.microsoft.com/dotnet/runtime:8.0@sha256:d304745fcb833f4dd78a0be39812f594722c598303cf0ed3feefd1cbb5d8cc48` |
| 2 | `mcr.microsoft.com/dotnet/aspnet:8.0` | 8.0.24 | 1 | 2 | 95 | 218.0 MB | `sha256:88c86e5469c2` | `mcr.microsoft.com/dotnet/aspnet:8.0@sha256:88c86e5469c2b69969e07343193fd3dbe4ff6936d95b994dbfa27348299af111` |
| 3 | `mcr.microsoft.com/dotnet/sdk:8.0` | 8.0.418 | 1 | 15 | 177 | 850.0 MB | `sha256:bfb6ed602caa` | `mcr.microsoft.com/dotnet/sdk:8.0@sha256:bfb6ed602caa605141700aea7dc7d42574b74b704368e67d683c71a002123808` |

### Ubuntu

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0-noble` | 8.0.24 | 0 | 0 | 12 | 193.0 MB | `sha256:6c946d640ec5` | `mcr.microsoft.com/dotnet/runtime:8.0-noble@sha256:6c946d640ec5fb8ca9e2c578903287a12ac395b3e4806703022b253174aed3c6` |
| 2 | `mcr.microsoft.com/dotnet/runtime:9.0-noble` | 9.0.13 | 0 | 0 | 12 | 198.0 MB | `sha256:a66ee75449f9` | `mcr.microsoft.com/dotnet/runtime:9.0-noble@sha256:a66ee75449f9f5158fa789252491b8d74ddb53660226860f8261ac475db0a0ff` |
| 3 | `mcr.microsoft.com/dotnet/runtime:10.0-noble` | 10.0.3 | 0 | 0 | 12 | 203.0 MB | `sha256:20c8b96cccff` | `mcr.microsoft.com/dotnet/runtime:10.0-noble@sha256:20c8b96cccff81d140c32ada512417f1c42248b64f54a80a319311b813a9dfb4` |
| 4 | `mcr.microsoft.com/dotnet/aspnet:8.0-noble` | 8.0.24 | 0 | 0 | 12 | 217.0 MB | `sha256:5ee04807992d` | `mcr.microsoft.com/dotnet/aspnet:8.0-noble@sha256:5ee04807992d78c8aadecd6c551750d3010f1300b94328ecb05e29d35ae10f04` |
| 5 | `mcr.microsoft.com/dotnet/aspnet:9.0-noble` | 9.0.13 | 0 | 0 | 12 | 223.0 MB | `sha256:06cf1b820062` | `mcr.microsoft.com/dotnet/aspnet:9.0-noble@sha256:06cf1b8200622a8b1ec766118e476ec7ffe5d919199522c87326662cccc58421` |
| 6 | `mcr.microsoft.com/dotnet/aspnet:10.0-noble` | 10.0.3 | 0 | 0 | 12 | 230.0 MB | `sha256:aec87aa74ddf` | `mcr.microsoft.com/dotnet/aspnet:10.0-noble@sha256:aec87aa74ddf129da573fa69f42f229a23c953a1c6fdecedea1aa6b1fe147d76` |
| 7 | `mcr.microsoft.com/dotnet/sdk:9.0-noble` | 9.0.311 | 0 | 0 | 37 | 836.0 MB | `sha256:45d815f7ad35` | `mcr.microsoft.com/dotnet/sdk:9.0-noble@sha256:45d815f7ad357e2cb10ac9c5247d42af972dd3c66a34f64bdf34a5e39bde698a` |
| 8 | `mcr.microsoft.com/dotnet/sdk:10.0-noble` | 10.0.103 | 0 | 0 | 37 | 884.0 MB | `sha256:e362a8dbcd69` | `mcr.microsoft.com/dotnet/sdk:10.0-noble@sha256:e362a8dbcd691522456da26a5198b8f3ca1d7641c95624fadc5e3e82678bd08a` |
| 9 | `mcr.microsoft.com/dotnet/sdk:8.0-noble` | 8.0.418 | 0 | 8 | 45 | 836.0 MB | `sha256:d582c7052a55` | `mcr.microsoft.com/dotnet/sdk:8.0-noble@sha256:d582c7052a557c953858a3e2f77d4dd6fc184dcf617b14cbcc426153fcd450d8` |

## Go

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0` | 1.25.8 | 0 | 0 | 0 | 809.0 MB | `sha256:8deb7086a5fd` | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0@sha256:8deb7086a5fd7fe8b4258ca4c3431141b3d0daafc514c7da09d01da43a6e6ad0` |
| 2 | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0` | 1.26.1 | 0 | 0 | 0 | 841.0 MB | `sha256:1c76136f9648` | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0@sha256:1c76136f964815a0eeea177a348fc7a7ca19c9c1a25e6b31a4666dbb5b210cc4` |

## Java

### Azure Linux

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:21-azurelinux` | 21.0.10 | 0 | 0 | 0 | 482.0 MB | `sha256:b5b331195224` | `mcr.microsoft.com/openjdk/jdk:21-azurelinux@sha256:b5b33119522439771a6d9086b41ca70ccf4de373f7c5325b61e8c7016020b02e` |
| 2 | `mcr.microsoft.com/openjdk/jdk:25-azurelinux` | 25.0.2 | 0 | 0 | 0 | 527.0 MB | `sha256:e84b14873c27` | `mcr.microsoft.com/openjdk/jdk:25-azurelinux@sha256:e84b14873c2736485e066b54773ab00e6d7eef0a8a6f89e0bf0284935314fafb` |
| 3 | `mcr.microsoft.com/openjdk/jdk:11-distroless` | 11.0.30 | 0 | 0 | 3 | 320.0 MB | `sha256:30969baf5c45` | `mcr.microsoft.com/openjdk/jdk:11-distroless@sha256:30969baf5c45efa2013037faf9bbc2d7974a1feb2a17e9f9b8ec9a6b9d8843f7` |
| 4 | `mcr.microsoft.com/openjdk/jdk:17-distroless` | 17.0.18 | 0 | 0 | 3 | 323.0 MB | `sha256:b66af1bbbd55` | `mcr.microsoft.com/openjdk/jdk:17-distroless@sha256:b66af1bbbd55a24b7bc8c47bf2fa11d2e4fdd4a16cfb059bab3fffe4d759a394` |
| 5 | `mcr.microsoft.com/openjdk/jdk:21-distroless` | 21.0.10 | 0 | 0 | 3 | 351.0 MB | `sha256:281220e4307c` | `mcr.microsoft.com/openjdk/jdk:21-distroless@sha256:281220e4307c485fddddfc175bdf14cd62abef49da50204eda9afccc50221343` |
| 6 | `mcr.microsoft.com/openjdk/jdk:25-distroless` | 25.0.2 | 0 | 0 | 3 | 396.0 MB | `sha256:c190efe034ba` | `mcr.microsoft.com/openjdk/jdk:25-distroless@sha256:c190efe034ba6730d52f5875f50f76a3f44fe1d6c83b57a41683a2dd00dac606` |

### Ubuntu

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:17-ubuntu` | 17.0.18 | 0 | 0 | 66 | 430.0 MB | `sha256:0f3b0e116eab` | `mcr.microsoft.com/openjdk/jdk:17-ubuntu@sha256:0f3b0e116eabef304ae058d61c63a492d48212c44fcd18dfe2d6fb6c3d0af5d7` |
| 2 | `mcr.microsoft.com/openjdk/jdk:21-ubuntu` | 21.0.10 | 0 | 0 | 66 | 457.0 MB | `sha256:84e0bf61cbc6` | `mcr.microsoft.com/openjdk/jdk:21-ubuntu@sha256:84e0bf61cbc6f619f74cbbb4b41b077b0535f16792c158e15f0ed9b166554f3b` |
| 3 | `mcr.microsoft.com/openjdk/jdk:25-ubuntu` | 25.0.2 | 0 | 0 | 66 | 502.0 MB | `sha256:180f372724db` | `mcr.microsoft.com/openjdk/jdk:25-ubuntu@sha256:180f372724db62715cf6d1dafeb168d20e66b13dad493d33d1d24c07d0cc55c0` |

## Node

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/base/nodejs:20.14` | 20.14.0 | 0 | 0 | 0 | 146.0 MB | `sha256:59a1e9373259` | `mcr.microsoft.com/azurelinux/base/nodejs:20.14@sha256:59a1e937325941dbd3bb722be4cef2f2b4b78bb8befa089c8b4915fd389d8662` |
| 2 | `mcr.microsoft.com/azurelinux/base/nodejs:20` | 20.14.0 | 0 | 0 | 0 | 146.0 MB | `sha256:59a1e9373259` | `mcr.microsoft.com/azurelinux/base/nodejs:20@sha256:59a1e937325941dbd3bb722be4cef2f2b4b78bb8befa089c8b4915fd389d8662` |
| 3 | `mcr.microsoft.com/azurelinux/distroless/nodejs:20.14-nonroot` | 20.14.0 | 0 | 10 | 12 | 106.0 MB | `sha256:1dccae1ebad2` | `mcr.microsoft.com/azurelinux/distroless/nodejs:20.14-nonroot@sha256:1dccae1ebad222f36adc1207edfd8b0da2c06a71a6ccc7ceced54516f4af349a` |
| 4 | `mcr.microsoft.com/azurelinux/distroless/nodejs:20.14` | 20.14.0 | 0 | 10 | 12 | 106.0 MB | `sha256:6643a8667737` | `mcr.microsoft.com/azurelinux/distroless/nodejs:20.14@sha256:6643a866773721bea665e29606618a55d28d42fdefb83e38b2e462fbb763a356` |
| 5 | `mcr.microsoft.com/azurelinux/distroless/nodejs:20-nonroot` | 20.14.0 | 0 | 10 | 12 | 106.0 MB | `sha256:1dccae1ebad2` | `mcr.microsoft.com/azurelinux/distroless/nodejs:20-nonroot@sha256:1dccae1ebad222f36adc1207edfd8b0da2c06a71a6ccc7ceced54516f4af349a` |
| 6 | `mcr.microsoft.com/azurelinux/distroless/nodejs:20` | 20.14.0 | 0 | 10 | 12 | 106.0 MB | `sha256:6643a8667737` | `mcr.microsoft.com/azurelinux/distroless/nodejs:20@sha256:6643a866773721bea665e29606618a55d28d42fdefb83e38b2e462fbb763a356` |

## Python

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot` | 3.12.9 | 0 | 0 | 0 | 83.6 MB | `sha256:66bdc517a7fe` | `mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot@sha256:66bdc517a7fe1fefa119516a7874881d763d72d234697c5f6702ba18d2d72cff` |
| 2 | `mcr.microsoft.com/azurelinux/distroless/python:3.12` | 3.12.9 | 0 | 0 | 0 | 83.6 MB | `sha256:358a31b2fed4` | `mcr.microsoft.com/azurelinux/distroless/python:3.12@sha256:358a31b2fed412c12d864b69f58374653bd1fc6c7140ade660d60d4acbb089c3` |
| 3 | `mcr.microsoft.com/azurelinux/distroless/python:3-nonroot` | 3.12.9 | 0 | 0 | 0 | 83.6 MB | `sha256:66bdc517a7fe` | `mcr.microsoft.com/azurelinux/distroless/python:3-nonroot@sha256:66bdc517a7fe1fefa119516a7874881d763d72d234697c5f6702ba18d2d72cff` |
| 4 | `mcr.microsoft.com/azurelinux/distroless/python:3` | 3.12.9 | 0 | 0 | 0 | 83.6 MB | `sha256:358a31b2fed4` | `mcr.microsoft.com/azurelinux/distroless/python:3@sha256:358a31b2fed412c12d864b69f58374653bd1fc6c7140ade660d60d4acbb089c3` |
| 5 | `mcr.microsoft.com/azurelinux/base/python:3.12` | 3.12.9 | 0 | 0 | 0 | 139.0 MB | `sha256:7be8b46a4dfa` | `mcr.microsoft.com/azurelinux/base/python:3.12@sha256:7be8b46a4dfad5acc5851ecf76ea800cd2818eaec40832baf5787849645461fd` |
| 6 | `mcr.microsoft.com/azurelinux/base/python:3` | 3.12.9 | 0 | 0 | 0 | 139.0 MB | `sha256:7be8b46a4dfa` | `mcr.microsoft.com/azurelinux/base/python:3@sha256:7be8b46a4dfad5acc5851ecf76ea800cd2818eaec40832baf5787849645461fd` |
| 7 | `mcr.microsoft.com/openjdk/jdk:21-azurelinux` | 3.12.9 | 0 | 0 | 0 | 482.0 MB | `sha256:b5b331195224` | `mcr.microsoft.com/openjdk/jdk:21-azurelinux@sha256:b5b33119522439771a6d9086b41ca70ccf4de373f7c5325b61e8c7016020b02e` |
| 8 | `mcr.microsoft.com/openjdk/jdk:25-azurelinux` | 3.12.9 | 0 | 0 | 0 | 527.0 MB | `sha256:e84b14873c27` | `mcr.microsoft.com/openjdk/jdk:25-azurelinux@sha256:e84b14873c2736485e066b54773ab00e6d7eef0a8a6f89e0bf0284935314fafb` |
| 9 | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0` | 3.12.9 | 0 | 0 | 0 | 809.0 MB | `sha256:8deb7086a5fd` | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0@sha256:8deb7086a5fd7fe8b4258ca4c3431141b3d0daafc514c7da09d01da43a6e6ad0` |
| 10 | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0` | 3.12.9 | 0 | 0 | 0 | 841.0 MB | `sha256:1c76136f9648` | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0@sha256:1c76136f964815a0eeea177a348fc7a7ca19c9c1a25e6b31a4666dbb5b210cc4` |

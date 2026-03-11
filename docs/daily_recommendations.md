# Daily Recommended Images by Language

_Generated: 2026-03-11T04:26:58Z. Criteria: lowest critical → high → total vulnerabilities → size. Top 10 per language per base OS._

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
| 3 | `mcr.microsoft.com/dotnet/sdk:8.0` | 8.0.419 | 1 | 16 | 177 | 850.0 MB | `sha256:312aeb1357b2` | `mcr.microsoft.com/dotnet/sdk:8.0@sha256:312aeb1357b2461f5a729b5ca38a54b8381c8e7880d0672585e8613cdc972096` |

### Ubuntu

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:8.0-noble` | 8.0.25 | 0 | 0 | 12 | 193.0 MB | `sha256:4afe8693b45c` | `mcr.microsoft.com/dotnet/runtime:8.0-noble@sha256:4afe8693b45ca972c98fd8b04bbe3d898111f1a9d1754f1c8fabae04da0c93aa` |
| 2 | `mcr.microsoft.com/dotnet/runtime:9.0-noble` | 9.0.14 | 0 | 0 | 12 | 198.0 MB | `sha256:ff70733f9ded` | `mcr.microsoft.com/dotnet/runtime:9.0-noble@sha256:ff70733f9ded21a73a0e0b17eee7a7c50b732f5ebc79e710ede6491982cc4fa0` |
| 3 | `mcr.microsoft.com/dotnet/runtime:10.0-noble` | 10.0.4 | 0 | 0 | 12 | 203.0 MB | `sha256:b1fc29ba8c7c` | `mcr.microsoft.com/dotnet/runtime:10.0-noble@sha256:b1fc29ba8c7cf56eacc4aaf078dbbfa70b533130968957356d0fef95d3ac62ac` |
| 4 | `mcr.microsoft.com/dotnet/aspnet:8.0-noble` | 8.0.25 | 0 | 0 | 12 | 217.0 MB | `sha256:69277bbf9a95` | `mcr.microsoft.com/dotnet/aspnet:8.0-noble@sha256:69277bbf9a95acea6be093f7e546e5ab9ee55768f12a9f32c3d543cd350d039e` |
| 5 | `mcr.microsoft.com/dotnet/aspnet:9.0-noble` | 9.0.14 | 0 | 0 | 12 | 223.0 MB | `sha256:281c970f6d43` | `mcr.microsoft.com/dotnet/aspnet:9.0-noble@sha256:281c970f6d43a3dc577d61c635275faf08e3ceeef4533462fe7e30884a23c4e6` |
| 6 | `mcr.microsoft.com/dotnet/aspnet:10.0-noble` | 10.0.4 | 0 | 0 | 12 | 230.0 MB | `sha256:8b75cdf59a50` | `mcr.microsoft.com/dotnet/aspnet:10.0-noble@sha256:8b75cdf59a5068d9adfd8a6d202cc7671b2dc8f5f46c51e3b88a0a632e8fad1f` |
| 7 | `mcr.microsoft.com/dotnet/sdk:9.0-noble` | 9.0.312 | 0 | 0 | 19 | 836.0 MB | `sha256:1c1583362df1` | `mcr.microsoft.com/dotnet/sdk:9.0-noble@sha256:1c1583362df1cd0caa017418d5ada26b6db140418fddb6c2718c41eb0b36d3ad` |
| 8 | `mcr.microsoft.com/dotnet/sdk:10.0-noble` | 10.0.200 | 0 | 0 | 19 | 890.0 MB | `sha256:ea13841f10c6` | `mcr.microsoft.com/dotnet/sdk:10.0-noble@sha256:ea13841f10c6c410a6df63c6c97ab549dfd2b5fcfff1c00186531ba30208117d` |
| 9 | `mcr.microsoft.com/dotnet/sdk:8.0-noble` | 8.0.419 | 0 | 8 | 27 | 836.0 MB | `sha256:d64fbc29eff1` | `mcr.microsoft.com/dotnet/sdk:8.0-noble@sha256:d64fbc29eff1787a23da76a6a580d4f2bcc8c6a8d3d432d74723fcd401f6b22a` |

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
| 3 | `mcr.microsoft.com/azurelinux/distroless/nodejs:20.14-nonroot` | 20.14.0 | 0 | 11 | 13 | 106.0 MB | `sha256:1dccae1ebad2` | `mcr.microsoft.com/azurelinux/distroless/nodejs:20.14-nonroot@sha256:1dccae1ebad222f36adc1207edfd8b0da2c06a71a6ccc7ceced54516f4af349a` |
| 4 | `mcr.microsoft.com/azurelinux/distroless/nodejs:20.14` | 20.14.0 | 0 | 11 | 13 | 106.0 MB | `sha256:6643a8667737` | `mcr.microsoft.com/azurelinux/distroless/nodejs:20.14@sha256:6643a866773721bea665e29606618a55d28d42fdefb83e38b2e462fbb763a356` |
| 5 | `mcr.microsoft.com/azurelinux/distroless/nodejs:20-nonroot` | 20.14.0 | 0 | 11 | 13 | 106.0 MB | `sha256:1dccae1ebad2` | `mcr.microsoft.com/azurelinux/distroless/nodejs:20-nonroot@sha256:1dccae1ebad222f36adc1207edfd8b0da2c06a71a6ccc7ceced54516f4af349a` |
| 6 | `mcr.microsoft.com/azurelinux/distroless/nodejs:20` | 20.14.0 | 0 | 11 | 13 | 106.0 MB | `sha256:6643a8667737` | `mcr.microsoft.com/azurelinux/distroless/nodejs:20@sha256:6643a866773721bea665e29606618a55d28d42fdefb83e38b2e462fbb763a356` |

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

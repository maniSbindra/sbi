# Daily Recommended Images by Language

_Generated: 2026-03-07T04:14:31Z. Criteria: lowest critical → high → total vulnerabilities → size. Top 10 per language._

**Note:** Image sizes are based on Linux amd64 platform as reported by `docker images` on GitHub runners. Actual sizes may vary significantly on other platforms (macOS, Windows, etc.).

## Scanned Repositories and Images

This report includes analysis from **26 configured sources** across 5 groups (see [repositories.json](../config/repositories.json)):

**Azure Linux base images:**

- `azurelinux/base/core`
- `azurelinux/base/python`
- `azurelinux/base/nodejs`

**Azure Linux distroless images:**

- `azurelinux/distroless/base`
- `azurelinux/distroless/minimal`
- `azurelinux/distroless/python`
- `azurelinux/distroless/nodejs`

**.NET images:**

- `mcr.microsoft.com/dotnet/aspnet:8.0`
- `mcr.microsoft.com/dotnet/runtime:8.0`
- `mcr.microsoft.com/dotnet/sdk:8.0`
- `mcr.microsoft.com/dotnet/aspnet:10.0-azurelinux3.0-distroless`
- `mcr.microsoft.com/dotnet/aspnet:9.0-azurelinux3.0-distroless`
- `mcr.microsoft.com/dotnet/aspnet:8.0-azurelinux3.0-distroless`
- `mcr.microsoft.com/dotnet/runtime:10.0-azurelinux3.0-distroless`
- `mcr.microsoft.com/dotnet/runtime:9.0-azurelinux3.0-distroless`
- `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0`
- `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0`

**Go images:**

- `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0`
- `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0`

**OpenJDK images:**

- `mcr.microsoft.com/openjdk/jdk:25-azurelinux`
- `mcr.microsoft.com/openjdk/jdk:25-distroless`
- `mcr.microsoft.com/openjdk/jdk:21-azurelinux`
- `mcr.microsoft.com/openjdk/jdk:21-distroless`
- `mcr.microsoft.com/openjdk/jdk:21-ubuntu`
- `mcr.microsoft.com/openjdk/jdk:17-distroless`
- `mcr.microsoft.com/openjdk/jdk:11-distroless`

## Dotnet

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/dotnet/runtime:9.0-azurelinux3.0-distroless` | 9.0.13 | 0 | 0 | 0 | 105.0 MB | `sha256:3995507d446e` | `mcr.microsoft.com/dotnet/runtime:9.0-azurelinux3.0-distroless@sha256:3995507d446ee3ce1fd278962e7a6749495811edce2519ecee31274030113afa` |
| 2 | `mcr.microsoft.com/dotnet/runtime:10.0-azurelinux3.0-distroless` | 10.0.3 | 0 | 0 | 0 | 110.0 MB | `sha256:9d28d4f5ccb4` | `mcr.microsoft.com/dotnet/runtime:10.0-azurelinux3.0-distroless@sha256:9d28d4f5ccb4f54ceddfb54597a2bb991a93a07f3e2f6a61845e3e98ec040892` |
| 3 | `mcr.microsoft.com/dotnet/aspnet:8.0-azurelinux3.0-distroless` | 8.0.24 | 0 | 0 | 0 | 124.0 MB | `sha256:e91f943b47d5` | `mcr.microsoft.com/dotnet/aspnet:8.0-azurelinux3.0-distroless@sha256:e91f943b47d5c2e2d6ac7cdd979e414c82aeb11980c8ef52e409fcf2afa63913` |
| 4 | `mcr.microsoft.com/dotnet/aspnet:9.0-azurelinux3.0-distroless` | 9.0.13 | 0 | 0 | 0 | 130.0 MB | `sha256:9f856e8de88e` | `mcr.microsoft.com/dotnet/aspnet:9.0-azurelinux3.0-distroless@sha256:9f856e8de88e2ca03687bca28e8575f2e3255ea43df1896ffbd76023e7b07f41` |
| 5 | `mcr.microsoft.com/dotnet/aspnet:10.0-azurelinux3.0-distroless` | 10.0.3 | 0 | 0 | 0 | 137.0 MB | `sha256:973ac891bc21` | `mcr.microsoft.com/dotnet/aspnet:10.0-azurelinux3.0-distroless@sha256:973ac891bc21916cb4f579ed3cd5737fac0a1452d30b11a25493df65eefd4786` |
| 6 | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0` | 9.0.311 | 0 | 0 | 0 | 879.0 MB | `sha256:ef9d3354c106` | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0@sha256:ef9d3354c106e4efa0946c79dca08038c28cd03da2eef678d6b9a640028218fe` |
| 7 | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0` | 10.0.103 | 0 | 0 | 0 | 928.0 MB | `sha256:774b85a4de4b` | `mcr.microsoft.com/dotnet/sdk:10.0-azurelinux3.0@sha256:774b85a4de4b63b5974d0a599e7166b0ff20f3848d3cc67333beb733405d027b` |
| 8 | `mcr.microsoft.com/dotnet/runtime:8.0` | 8.0.24 | 1 | 2 | 95 | 193.0 MB | `sha256:d304745fcb83` | `mcr.microsoft.com/dotnet/runtime:8.0@sha256:d304745fcb833f4dd78a0be39812f594722c598303cf0ed3feefd1cbb5d8cc48` |
| 9 | `mcr.microsoft.com/dotnet/aspnet:8.0` | 8.0.24 | 1 | 2 | 95 | 218.0 MB | `sha256:88c86e5469c2` | `mcr.microsoft.com/dotnet/aspnet:8.0@sha256:88c86e5469c2b69969e07343193fd3dbe4ff6936d95b994dbfa27348299af111` |
| 10 | `mcr.microsoft.com/dotnet/sdk:8.0` | 8.0.418 | 1 | 15 | 177 | 850.0 MB | `sha256:bfb6ed602caa` | `mcr.microsoft.com/dotnet/sdk:8.0@sha256:bfb6ed602caa605141700aea7dc7d42574b74b704368e67d683c71a002123808` |

## Go

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0` | 1.26.1 | 0 | 0 | 0 | 841.0 MB | `sha256:80bf045513f1` | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0@sha256:80bf045513f114e7d9190df22d932effe19106c49b6c75c35de7c110fe124869` |
| 2 | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0` | 1.25.8 | 0 | 0 | 18 | 809.0 MB | `sha256:38282b1b0ad9` | `mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0@sha256:38282b1b0ad90ea79719d26e340392af9ee91d23948673d70ecb4a296960ebc7` |

## Java

| Rank | Image | Version | Crit | High | Total | Size | Digest | Pinned Reference |
|------|-------|---------|------|------|-------|------|--------|------------------|
| 1 | `mcr.microsoft.com/openjdk/jdk:21-azurelinux` | 21.0.10 | 0 | 0 | 0 | 482.0 MB | `sha256:6612f4edbfee` | `mcr.microsoft.com/openjdk/jdk:21-azurelinux@sha256:6612f4edbfeef81f814fea1e38f2ebb2b34807485ac9da68136e2038dfbeb647` |
| 2 | `mcr.microsoft.com/openjdk/jdk:25-azurelinux` | 25.0.2 | 0 | 0 | 0 | 527.0 MB | `sha256:46807e55432f` | `mcr.microsoft.com/openjdk/jdk:25-azurelinux@sha256:46807e55432ff5690086add25209f6922d4b02cdcfcc07b8d5daa4a764c86f47` |
| 3 | `mcr.microsoft.com/openjdk/jdk:11-distroless` | 11.0.30 | 0 | 0 | 5 | 320.0 MB | `sha256:e8f7191d5267` | `mcr.microsoft.com/openjdk/jdk:11-distroless@sha256:e8f7191d5267cd5020f5756e8621b2608fb3224054b4f84d6b7c6f5d031d7363` |
| 4 | `mcr.microsoft.com/openjdk/jdk:17-distroless` | 17.0.18 | 0 | 0 | 5 | 323.0 MB | `sha256:5d16efd09a4b` | `mcr.microsoft.com/openjdk/jdk:17-distroless@sha256:5d16efd09a4baec50193d81dcadc051fccc06fa815d91b473aaaf7688655fabb` |
| 5 | `mcr.microsoft.com/openjdk/jdk:21-distroless` | 21.0.10 | 0 | 0 | 5 | 351.0 MB | `sha256:f2efd77c7bab` | `mcr.microsoft.com/openjdk/jdk:21-distroless@sha256:f2efd77c7babdcd983afa3e83eb1649ee6837e5867de147ebb5e3481e12488e4` |
| 6 | `mcr.microsoft.com/openjdk/jdk:25-distroless` | 25.0.2 | 0 | 0 | 5 | 396.0 MB | `sha256:280e7f3ae5e2` | `mcr.microsoft.com/openjdk/jdk:25-distroless@sha256:280e7f3ae5e25a056a71fa21b97ccfb741dfa40194235edd299c33a26c0935a7` |
| 7 | `mcr.microsoft.com/openjdk/jdk:21-ubuntu` | 21.0.10 | 0 | 0 | 68 | 457.0 MB | `sha256:a62ef8b27d0c` | `mcr.microsoft.com/openjdk/jdk:21-ubuntu@sha256:a62ef8b27d0c87720389bb8f953495ccbf0d70a38791bcc0bb853e49507f3a3e` |

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
| 7 | `mcr.microsoft.com/openjdk/jdk:21-azurelinux` | 3.12.9 | 0 | 0 | 0 | 482.0 MB | `sha256:6612f4edbfee` | `mcr.microsoft.com/openjdk/jdk:21-azurelinux@sha256:6612f4edbfeef81f814fea1e38f2ebb2b34807485ac9da68136e2038dfbeb647` |
| 8 | `mcr.microsoft.com/openjdk/jdk:25-azurelinux` | 3.12.9 | 0 | 0 | 0 | 527.0 MB | `sha256:46807e55432f` | `mcr.microsoft.com/openjdk/jdk:25-azurelinux@sha256:46807e55432ff5690086add25209f6922d4b02cdcfcc07b8d5daa4a764c86f47` |
| 9 | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0` | 3.12.9 | 0 | 0 | 0 | 841.0 MB | `sha256:80bf045513f1` | `mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0@sha256:80bf045513f114e7d9190df22d932effe19106c49b6c75c35de7c110fe124869` |
| 10 | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0` | 3.12.9 | 0 | 0 | 0 | 879.0 MB | `sha256:ef9d3354c106` | `mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0@sha256:ef9d3354c106e4efa0946c79dca08038c28cd03da2eef678d6b9a640028218fe` |


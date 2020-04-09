# OCI Distribution Spec Conformance Results

Static HTML reports of [OCI Distribution Spec](https://github.com/opencontainers/distribution-spec) conformance tests run against various registries.

Note: a test result is considered "verified" if a registry provider has submitted results to [oci-conformance](https://github.com/opencontainers/oci-conformance).

## Registry Conformance Summary

### Open Source

| Registry | Status | Verified? | Notes | Links |
| -------- | -------- | -------- | -------- | -------- |
| [anuvu/zot](https://github.com/anuvu/zot) | [![](https://github.com/opencontainers/oci-conformance/workflows/zot/badge.svg)](https://github.com/opencontainers/oci-conformance/actions?query=workflow%3Azot) | [under review](https://github.com/opencontainers/oci-conformance/pull/38) | | [job definition](https://github.com/opencontainers/oci-conformance/blob/master/.github/workflows/zot.yml) &#x7c; [test report](https://oci-distribution-conformance-results.s3.amazonaws.com/zot.html) |
| [containersolutions/trow](https://github.com/containersolutions/trow) | [![](https://github.com/opencontainers/oci-conformance/workflows/trow/badge.svg)](https://github.com/opencontainers/oci-conformance/actions?query=workflow%3Atrow) | [under review](https://github.com/opencontainers/oci-conformance/pull/41) | | [job definition](https://github.com/opencontainers/oci-conformance/blob/master/.github/workflows/trow.yml) &#x7c; [test report](https://oci-distribution-conformance-results.s3.amazonaws.com/trow.html) |
| [docker/distribution](https://github.com/docker/distribution) | [![](https://github.com/opencontainers/oci-conformance/workflows/distribution/badge.svg)](https://github.com/opencontainers/oci-conformance/actions?query=workflow%3Adistribution) | no | Known issues: [#1936](https://github.com/docker/distribution/issues/1936) [#3141](https://github.com/docker/distribution/issues/3141) | [job definition](https://github.com/opencontainers/oci-conformance/blob/master/.github/workflows/distribution.yml) &#x7c; [test report](https://oci-distribution-conformance-results.s3.amazonaws.com/distribution.html) |
| [goharbor/harbor](https://github.com/goharbor/harbor) | | no | In Progress | |
| [quay/quay](https://github.com/quay/quay) | | no | TODO | |
| [sonatype/nexus3](https://github.com/sonatype/docker-nexus3) | | no | TODO | |

### Commercial

| Registry | Vendor | Status | Verified? | Notes | Links |
| -------- | -------- | -------- | -------- | -------- | -------- |
| [ACR](https://azure.microsoft.com/en-us/services/container-registry/) | Microsoft | [![](https://github.com/opencontainers/oci-conformance/workflows/acr/badge.svg)](https://github.com/opencontainers/oci-conformance/actions?query=workflow%3Aacr) | no | | [job definition](https://github.com/opencontainers/oci-conformance/blob/master/.github/workflows/acr.yml) &#x7c; [test report](https://oci-distribution-conformance-results.s3.amazonaws.com/acr.html) |
| [Alibaba](https://www.alibabacloud.com/product/container-registry) | Alibaba | [![](https://github.com/opencontainers/oci-conformance/workflows/alibaba/badge.svg)](https://github.com/opencontainers/oci-conformance/actions?query=workflow%3Aalibaba) | no | | [job definition](https://github.com/opencontainers/oci-conformance/blob/master/.github/workflows/alibaba.yml) &#x7c; [test report](https://oci-distribution-conformance-results.s3.amazonaws.com/alibaba.html) |
| [Docker Hub](https://hub.docker.com/) | Docker | [![](https://github.com/opencontainers/oci-conformance/workflows/dockerhub/badge.svg)](https://github.com/opencontainers/oci-conformance/actions?query=workflow%3Adockerhub) | no | | [job definition](https://github.com/opencontainers/oci-conformance/blob/master/.github/workflows/dockerhub.yml) &#x7c; [test report](https://oci-distribution-conformance-results.s3.amazonaws.com/dockerhub.html) |
| [ECR](https://aws.amazon.com/ecr/) | Amazon | [![](https://github.com/opencontainers/oci-conformance/workflows/ecr/badge.svg)](https://github.com/opencontainers/oci-conformance/actions?query=workflow%3Aecr) | no | | [job definition](https://github.com/opencontainers/oci-conformance/blob/master/.github/workflows/ecr.yml) &#x7c; [test report](https://oci-distribution-conformance-results.s3.amazonaws.com/ecr.html) |
| [GCR](https://cloud.google.com/container-registry/) | Google | [![](https://github.com/opencontainers/oci-conformance/workflows/gcr/badge.svg)](https://github.com/opencontainers/oci-conformance/actions?query=workflow%3Agcr) | no | | [job definition](https://github.com/opencontainers/oci-conformance/blob/master/.github/workflows/gcr.yml) &#x7c; [test report](https://oci-distribution-conformance-results.s3.amazonaws.com/gcr.html) |
| [GitHub](https://github.com/features/packages) | GitHub | [![](https://github.com/opencontainers/oci-conformance/workflows/github/badge.svg)](https://github.com/opencontainers/oci-conformance/actions?query=workflow%3Agithub) | no | | [job definition](https://github.com/opencontainers/oci-conformance/blob/master/.github/workflows/github.yml) &#x7c; [test report](https://oci-distribution-conformance-results.s3.amazonaws.com/github.html) |
| [JCR](https://jfrog.com/container-registry/) | JFrog | [![](https://github.com/opencontainers/oci-conformance/workflows/jcr/badge.svg)](https://github.com/opencontainers/oci-conformance/actions?query=workflow%3Ajcr) | no | | [job definition](https://github.com/opencontainers/oci-conformance/blob/master/.github/workflows/jcr.yml) &#x7c; [test report](https://oci-distribution-conformance-results.s3.amazonaws.com/jcr.html) |
| [Quay.io](https://quay.io/repository/) | Red Hat | [![](https://github.com/opencontainers/oci-conformance/workflows/quay/badge.svg)](https://github.com/opencontainers/oci-conformance/actions?query=workflow%3Aquay) | no | | [job definition](https://github.com/opencontainers/oci-conformance/blob/master/.github/workflows/quay.yml) &#x7c; [test report](https://oci-distribution-conformance-results.s3.amazonaws.com/quay.html) |

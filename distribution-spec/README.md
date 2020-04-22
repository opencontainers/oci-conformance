# OCI Distribution Spec Conformance Results

---
<p><strong style="color: #bfac1d">⚠️ WARNING</strong></p>

The test results found on this page do not indicate OCI conformance (or lack thereof) for any of the listed registries.
This page is meant for informational purposes only.

---

Static HTML reports of [OCI Distribution Spec](https://github.com/bloodorangeio/distribution-spec) conformance tests run against various registries.

See any bad or missing information? You are encouraged to [submit a pull request](https://github.com/bloodorangeio/oci-conformance/blob/master/distribution-spec/README.md).

Test results are considered "verified" if a registry provider has submitted results to [oci-conformance](https://github.com/bloodorangeio/oci-conformance).

__Table of Contents:__

- [Open Source](#open-source)
  - [anuvu/zot](#anuvu/zot)
  - [containersolutions/trow](#containersolutionstrow)
  - [docker/distribution](#dockerdistribution)
  - [goharbor/harbor](#goharborharbor)
  - [quay/quay](#quayquay)
  - [sonatype/nexus3](#sonatypenexus3)
- [Commercial](#commercial)
  - [ACR](#acr)
  - [Alibaba](#alibaba)
  - [Docker Hub](#docker-hub)
  - [ECR](#ecr)
  - [GCR](#gcr)
  - [GitHub](#github)
  - [JCR](#jcr)
  - [Quay.io](#quayio)

## Open Source

### anuvu/zot

__Homepage:__ [https://github.com/anuvu/zot](https://github.com/anuvu/zot)

__Verified:__ [under review](https://github.com/bloodorangeio/oci-conformance/pull/38)

__Test Results:__

| Workflow           | Status                                                                                                                                                               | Test Report                                                                                           | Job Definition                                                                                    |
| ------------------ | -------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- |
| Pull               | [![](https://github.com/bloodorangeio/oci-conformance/workflows/zot-1/badge.svg)](https://github.com/bloodorangeio/oci-conformance/actions?query=workflow%3Azot-1) | [Link](https://oci-conformance.s3.amazonaws.com/distribution-spec/zot/pull/report.html)               | [Link](https://github.com/bloodorangeio/oci-conformance/blob/master/.github/workflows/zot_1.yml) |
| Push               | [![](https://github.com/bloodorangeio/oci-conformance/workflows/zot-2/badge.svg)](https://github.com/bloodorangeio/oci-conformance/actions?query=workflow%3Azot-2) | [Link](https://oci-conformance.s3.amazonaws.com/distribution-spec/zot/push/report.html)               | [Link](https://github.com/bloodorangeio/oci-conformance/blob/master/.github/workflows/zot_2.yml) |
| Content Discovery  | [![](https://github.com/bloodorangeio/oci-conformance/workflows/zot-3/badge.svg)](https://github.com/bloodorangeio/oci-conformance/actions?query=workflow%3Azot-3) | [Link](https://oci-conformance.s3.amazonaws.com/distribution-spec/zot/content-discovery/report.html)  | [Link](https://github.com/bloodorangeio/oci-conformance/blob/master/.github/workflows/zot_3.yml) |
| Content Management | [![](https://github.com/bloodorangeio/oci-conformance/workflows/zot-4/badge.svg)](https://github.com/bloodorangeio/oci-conformance/actions?query=workflow%3Azot-4) | [Link](https://oci-conformance.s3.amazonaws.com/distribution-spec/zot/content-management/report.html) | [Link](https://github.com/bloodorangeio/oci-conformance/blob/master/.github/workflows/zot_4.yml) |

### containersolutions/trow

__Homepage:__ [https://github.com/containersolutions/trow](https://github.com/containersolutions/trow)

__Verified:__ [under review](https://github.com/bloodorangeio/oci-conformance/pull/41)

__Test Results:__

| Workflow           | Status                                                                                                                                                                 | Test Report                                                                                            | Job Definition                                                                                     |
| ------------------ | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------ | -------------------------------------------------------------------------------------------------- |
| Pull               | [![](https://github.com/bloodorangeio/oci-conformance/workflows/trow-1/badge.svg)](https://github.com/bloodorangeio/oci-conformance/actions?query=workflow%3Atrow-1) | [Link](https://oci-conformance.s3.amazonaws.com/distribution-spec/trow/pull/report.html)               | [Link](https://github.com/bloodorangeio/oci-conformance/blob/master/.github/workflows/trow_1.yml) |
| Push               | [![](https://github.com/bloodorangeio/oci-conformance/workflows/trow-2/badge.svg)](https://github.com/bloodorangeio/oci-conformance/actions?query=workflow%3Atrow-2) | [Link](https://oci-conformance.s3.amazonaws.com/distribution-spec/trow/push/report.html)               | [Link](https://github.com/bloodorangeio/oci-conformance/blob/master/.github/workflows/trow_2.yml) |
| Content Discovery  | [![](https://github.com/bloodorangeio/oci-conformance/workflows/trow-3/badge.svg)](https://github.com/bloodorangeio/oci-conformance/actions?query=workflow%3Atrow-3) | [Link](https://oci-conformance.s3.amazonaws.com/distribution-spec/trow/content-discovery/report.html)  | [Link](https://github.com/bloodorangeio/oci-conformance/blob/master/.github/workflows/trow_3.yml) |
| Content Management | [![](https://github.com/bloodorangeio/oci-conformance/workflows/trow-4/badge.svg)](https://github.com/bloodorangeio/oci-conformance/actions?query=workflow%3Atrow-4) | [Link](https://oci-conformance.s3.amazonaws.com/distribution-spec/trow/content-management/report.html) | [Link](https://github.com/bloodorangeio/oci-conformance/blob/master/.github/workflows/trow_4.yml) |

### docker/distribution

__Homepage:__ [https://github.com/docker/distribution](https://github.com/docker/distribution)

__Verified:__ no

__Notes:__ Known issues: [#1936](https://github.com/docker/distribution/issues/1936) [#3141](https://github.com/docker/distribution/issues/3141) 

__Test Results:__

| Workflow           | Status                                                                                                                                                                                 | Test Report                                                                                                    | Job Definition                                                                                             |
| ------------------ | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| Pull               | [![](https://github.com/bloodorangeio/oci-conformance/workflows/distribution-1/badge.svg)](https://github.com/bloodorangeio/oci-conformance/actions?query=workflow%3Adistribution-1) | [Link](https://oci-conformance.s3.amazonaws.com/distribution-spec/distribution/pull/report.html)               | [Link](https://github.com/bloodorangeio/oci-conformance/blob/master/.github/workflows/distribution_1.yml) |
| Push               | [![](https://github.com/bloodorangeio/oci-conformance/workflows/distribution-2/badge.svg)](https://github.com/bloodorangeio/oci-conformance/actions?query=workflow%3Adistribution-2) | [Link](https://oci-conformance.s3.amazonaws.com/distribution-spec/distribution/push/report.html)               | [Link](https://github.com/bloodorangeio/oci-conformance/blob/master/.github/workflows/distribution_2.yml) |
| Content Discovery  | [![](https://github.com/bloodorangeio/oci-conformance/workflows/distribution-3/badge.svg)](https://github.com/bloodorangeio/oci-conformance/actions?query=workflow%3Adistribution-3) | [Link](https://oci-conformance.s3.amazonaws.com/distribution-spec/distribution/content-discovery/report.html)  | [Link](https://github.com/bloodorangeio/oci-conformance/blob/master/.github/workflows/distribution_3.yml) |
| Content Management | [![](https://github.com/bloodorangeio/oci-conformance/workflows/distribution-4/badge.svg)](https://github.com/bloodorangeio/oci-conformance/actions?query=workflow%3Adistribution-4) | [Link](https://oci-conformance.s3.amazonaws.com/distribution-spec/distribution/content-management/report.html) | [Link](https://github.com/bloodorangeio/oci-conformance/blob/master/.github/workflows/distribution_4.yml) |

### goharbor/harbor

__Homepage:__ [https://github.com/goharbor/harbor](https://github.com/goharbor/harbor)

__Verified:__ no

__Test Results:__

| Workflow           | Status                                                                                                                                                                   | Test Report                                                                                              | Job Definition                                                                                      |
| ------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | -------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------|
| Pull               | [![](https://github.com/bloodorangeio/oci-conformance/workflows/harbor-1/badge.svg)](https://github.com/bloodorangeio/oci-conformance/actions?query=workflow%3Aharbor-1) | [Link](https://oci-conformance.s3.amazonaws.com/distribution-spec/harbor/pull/report.html)               | [Link](https://github.com/bloodorangeio/oci-conformance/blob/master/.github/workflows/harbor_1.yml) |
| Push               | [![](https://github.com/bloodorangeio/oci-conformance/workflows/harbor-2/badge.svg)](https://github.com/bloodorangeio/oci-conformance/actions?query=workflow%3Aharbor-2) | [Link](https://oci-conformance.s3.amazonaws.com/distribution-spec/harbor/push/report.html)               | [Link](https://github.com/bloodorangeio/oci-conformance/blob/master/.github/workflows/harbor_2.yml) |
| Content Discovery  | [![](https://github.com/bloodorangeio/oci-conformance/workflows/harbor-3/badge.svg)](https://github.com/bloodorangeio/oci-conformance/actions?query=workflow%3Aharbor-3) | [Link](https://oci-conformance.s3.amazonaws.com/distribution-spec/harbor/content-discovery/report.html)  | [Link](https://github.com/bloodorangeio/oci-conformance/blob/master/.github/workflows/harbor_3.yml) |
| Content Management | [![](https://github.com/bloodorangeio/oci-conformance/workflows/harbor-4/badge.svg)](https://github.com/bloodorangeio/oci-conformance/actions?query=workflow%3Aharbor-4) | [Link](https://oci-conformance.s3.amazonaws.com/distribution-spec/harbor/content-management/report.html) | [Link](https://github.com/bloodorangeio/oci-conformance/blob/master/.github/workflows/harbor_4.yml) |

### quay/quay

__Homepage:__ [https://github.com/quay/quay](https://github.com/quay/quay)

_TODO_

### sonatype/nexus3

__Homepage:__ [https://github.com/sonatype/docker-nexus3](https://github.com/sonatype/docker-nexus3)

_TODO_

## Commercial

### ACR

__Homepage:__ [https://azure.microsoft.com/en-us/services/container-registry/](https://azure.microsoft.com/en-us/services/container-registry/)

__Vendor:__ Microsoft

__Verified:__ no

__Test Results:__

| Workflow           | Status                                                                                                                                                               | Test Report                                                                                           | Job Definition                                                                                    |
| ------------------ | -------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- |
| Pull               | [![](https://github.com/bloodorangeio/oci-conformance/workflows/acr-1/badge.svg)](https://github.com/bloodorangeio/oci-conformance/actions?query=workflow%3Aacr-1) | [Link](https://oci-conformance.s3.amazonaws.com/distribution-spec/acr/pull/report.html)               | [Link](https://github.com/bloodorangeio/oci-conformance/blob/master/.github/workflows/acr_1.yml) |
| Push               | [![](https://github.com/bloodorangeio/oci-conformance/workflows/acr-2/badge.svg)](https://github.com/bloodorangeio/oci-conformance/actions?query=workflow%3Aacr-2) | [Link](https://oci-conformance.s3.amazonaws.com/distribution-spec/acr/push/report.html)               | [Link](https://github.com/bloodorangeio/oci-conformance/blob/master/.github/workflows/acr_2.yml) |
| Content Discovery  | [![](https://github.com/bloodorangeio/oci-conformance/workflows/acr-3/badge.svg)](https://github.com/bloodorangeio/oci-conformance/actions?query=workflow%3Aacr-3) | [Link](https://oci-conformance.s3.amazonaws.com/distribution-spec/acr/content-discovery/report.html)  | [Link](https://github.com/bloodorangeio/oci-conformance/blob/master/.github/workflows/acr_3.yml) |
| Content Management | [![](https://github.com/bloodorangeio/oci-conformance/workflows/acr-4/badge.svg)](https://github.com/bloodorangeio/oci-conformance/actions?query=workflow%3Aacr-4) | [Link](https://oci-conformance.s3.amazonaws.com/distribution-spec/acr/content-management/report.html) | [Link](https://github.com/bloodorangeio/oci-conformance/blob/master/.github/workflows/acr_4.yml) |

### Alibaba

__Homepage:__ [https://www.alibabacloud.com/product/container-registry](https://www.alibabacloud.com/product/container-registry)

__Vendor:__ Alibaba

__Verified:__ no

__Test Results:__

| Workflow           | Status                                                                                                                                                                       | Test Report                                                                                               | Job Definition                                                                                        |
| ------------------ | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- |
| Pull               | [![](https://github.com/bloodorangeio/oci-conformance/workflows/alibaba-1/badge.svg)](https://github.com/bloodorangeio/oci-conformance/actions?query=workflow%3Aalibaba-1) | [Link](https://oci-conformance.s3.amazonaws.com/distribution-spec/alibaba/pull/report.html)               | [Link](https://github.com/bloodorangeio/oci-conformance/blob/master/.github/workflows/alibaba_1.yml) |
| Push               | [![](https://github.com/bloodorangeio/oci-conformance/workflows/alibaba-2/badge.svg)](https://github.com/bloodorangeio/oci-conformance/actions?query=workflow%3Aalibaba-2) | [Link](https://oci-conformance.s3.amazonaws.com/distribution-spec/alibaba/push/report.html)               | [Link](https://github.com/bloodorangeio/oci-conformance/blob/master/.github/workflows/alibaba_2.yml) |
| Content Discovery  | [![](https://github.com/bloodorangeio/oci-conformance/workflows/alibaba-3/badge.svg)](https://github.com/bloodorangeio/oci-conformance/actions?query=workflow%3Aalibaba-3) | [Link](https://oci-conformance.s3.amazonaws.com/distribution-spec/alibaba/content-discovery/report.html)  | [Link](https://github.com/bloodorangeio/oci-conformance/blob/master/.github/workflows/alibaba_3.yml) |
| Content Management | [![](https://github.com/bloodorangeio/oci-conformance/workflows/alibaba-4/badge.svg)](https://github.com/bloodorangeio/oci-conformance/actions?query=workflow%3Aalibaba-4) | [Link](https://oci-conformance.s3.amazonaws.com/distribution-spec/alibaba/content-management/report.html) | [Link](https://github.com/bloodorangeio/oci-conformance/blob/master/.github/workflows/alibaba_4.yml) |

### Docker Hub

__Homepage:__ [https://hub.docker.com](https://hub.docker.com)

__Vendor:__ Docker

__Verified:__ no

__Test Results:__

| Workflow           | Status                                                                                                                                                                           | Test Report                                                                                                 | Job Definition                                                                                          |
| ------------------ | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- |
| Pull               | [![](https://github.com/bloodorangeio/oci-conformance/workflows/dockerhub-1/badge.svg)](https://github.com/bloodorangeio/oci-conformance/actions?query=workflow%3Adockerhub-1) | [Link](https://oci-conformance.s3.amazonaws.com/distribution-spec/dockerhub/pull/report.html)               | [Link](https://github.com/bloodorangeio/oci-conformance/blob/master/.github/workflows/dockerhub_1.yml) |
| Push               | [![](https://github.com/bloodorangeio/oci-conformance/workflows/dockerhub-2/badge.svg)](https://github.com/bloodorangeio/oci-conformance/actions?query=workflow%3Adockerhub-2) | [Link](https://oci-conformance.s3.amazonaws.com/distribution-spec/dockerhub/push/report.html)               | [Link](https://github.com/bloodorangeio/oci-conformance/blob/master/.github/workflows/dockerhub_2.yml) |
| Content Discovery  | [![](https://github.com/bloodorangeio/oci-conformance/workflows/dockerhub-3/badge.svg)](https://github.com/bloodorangeio/oci-conformance/actions?query=workflow%3Adockerhub-3) | [Link](https://oci-conformance.s3.amazonaws.com/distribution-spec/dockerhub/content-discovery/report.html)  | [Link](https://github.com/bloodorangeio/oci-conformance/blob/master/.github/workflows/dockerhub_3.yml) |
| Content Management | [![](https://github.com/bloodorangeio/oci-conformance/workflows/dockerhub-4/badge.svg)](https://github.com/bloodorangeio/oci-conformance/actions?query=workflow%3Adockerhub-4) | [Link](https://oci-conformance.s3.amazonaws.com/distribution-spec/dockerhub/content-management/report.html) | [Link](https://github.com/bloodorangeio/oci-conformance/blob/master/.github/workflows/dockerhub_4.yml) |

### ECR

__Homepage:__ [https://aws.amazon.com/ecr/](https://aws.amazon.com/ecr/)

__Vendor:__ Amazon

__Verified:__ no

__Test Results:__

| Workflow           | Status                                                                                                                                                               | Test Report                                                                                           | Job Definition                                                                                    |
| ------------------ | -------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- |
| Pull               | [![](https://github.com/bloodorangeio/oci-conformance/workflows/ecr-1/badge.svg)](https://github.com/bloodorangeio/oci-conformance/actions?query=workflow%3Aecr-1) | [Link](https://oci-conformance.s3.amazonaws.com/distribution-spec/ecr/pull/report.html)               | [Link](https://github.com/bloodorangeio/oci-conformance/blob/master/.github/workflows/ecr_1.yml) |
| Push               | [![](https://github.com/bloodorangeio/oci-conformance/workflows/ecr-2/badge.svg)](https://github.com/bloodorangeio/oci-conformance/actions?query=workflow%3Aecr-2) | [Link](https://oci-conformance.s3.amazonaws.com/distribution-spec/ecr/push/report.html)               | [Link](https://github.com/bloodorangeio/oci-conformance/blob/master/.github/workflows/ecr_2.yml) |
| Content Discovery  | [![](https://github.com/bloodorangeio/oci-conformance/workflows/ecr-3/badge.svg)](https://github.com/bloodorangeio/oci-conformance/actions?query=workflow%3Aecr-3) | [Link](https://oci-conformance.s3.amazonaws.com/distribution-spec/ecr/content-discovery/report.html)  | [Link](https://github.com/bloodorangeio/oci-conformance/blob/master/.github/workflows/ecr_3.yml) |
| Content Management | [![](https://github.com/bloodorangeio/oci-conformance/workflows/ecr-4/badge.svg)](https://github.com/bloodorangeio/oci-conformance/actions?query=workflow%3Aecr-4) | [Link](https://oci-conformance.s3.amazonaws.com/distribution-spec/ecr/content-management/report.html) | [Link](https://github.com/bloodorangeio/oci-conformance/blob/master/.github/workflows/ecr_4.yml) |

### GCR

__Homepage:__ [https://cloud.google.com/container-registry/](https://cloud.google.com/container-registry/)

__Vendor:__ Google

__Verified:__ no

__Test Results:__

| Workflow           | Status                                                                                                                                                               | Test Report                                                                                           | Job Definition                                                                                    |
| ------------------ | -------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- |
| Pull               | [![](https://github.com/bloodorangeio/oci-conformance/workflows/gcr-1/badge.svg)](https://github.com/bloodorangeio/oci-conformance/actions?query=workflow%3Agcr-1) | [Link](https://oci-conformance.s3.amazonaws.com/distribution-spec/gcr/pull/report.html)               | [Link](https://github.com/bloodorangeio/oci-conformance/blob/master/.github/workflows/gcr_1.yml) |
| Push               | [![](https://github.com/bloodorangeio/oci-conformance/workflows/gcr-2/badge.svg)](https://github.com/bloodorangeio/oci-conformance/actions?query=workflow%3Agcr-2) | [Link](https://oci-conformance.s3.amazonaws.com/distribution-spec/gcr/push/report.html)               | [Link](https://github.com/bloodorangeio/oci-conformance/blob/master/.github/workflows/gcr_2.yml) |
| Content Discovery  | [![](https://github.com/bloodorangeio/oci-conformance/workflows/gcr-3/badge.svg)](https://github.com/bloodorangeio/oci-conformance/actions?query=workflow%3Agcr-3) | [Link](https://oci-conformance.s3.amazonaws.com/distribution-spec/gcr/content-discovery/report.html)  | [Link](https://github.com/bloodorangeio/oci-conformance/blob/master/.github/workflows/gcr_3.yml) |
| Content Management | [![](https://github.com/bloodorangeio/oci-conformance/workflows/gcr-4/badge.svg)](https://github.com/bloodorangeio/oci-conformance/actions?query=workflow%3Agcr-4) | [Link](https://oci-conformance.s3.amazonaws.com/distribution-spec/gcr/content-management/report.html) | [Link](https://github.com/bloodorangeio/oci-conformance/blob/master/.github/workflows/gcr_4.yml) |

### GitHub

__Homepage:__ [https://github.com/features/packages](https://github.com/features/packages)

__Vendor:__ GitHub

__Verified:__ no

__Test Results:__

| Workflow           | Status                                                                                                                                                                     | Test Report                                                                                              | Job Definition                                                                                       |
| ------------------ | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| Pull               | [![](https://github.com/bloodorangeio/oci-conformance/workflows/github-1/badge.svg)](https://github.com/bloodorangeio/oci-conformance/actions?query=workflow%3Agithub-1) | [Link](https://oci-conformance.s3.amazonaws.com/distribution-spec/github/pull/report.html)               | [Link](https://github.com/bloodorangeio/oci-conformance/blob/master/.github/workflows/github_1.yml) |
| Push               | [![](https://github.com/bloodorangeio/oci-conformance/workflows/github-2/badge.svg)](https://github.com/bloodorangeio/oci-conformance/actions?query=workflow%3Agithub-2) | [Link](https://oci-conformance.s3.amazonaws.com/distribution-spec/github/push/report.html)               | [Link](https://github.com/bloodorangeio/oci-conformance/blob/master/.github/workflows/github_2.yml) |
| Content Discovery  | [![](https://github.com/bloodorangeio/oci-conformance/workflows/github-3/badge.svg)](https://github.com/bloodorangeio/oci-conformance/actions?query=workflow%3Agithub-3) | [Link](https://oci-conformance.s3.amazonaws.com/distribution-spec/github/content-discovery/report.html)  | [Link](https://github.com/bloodorangeio/oci-conformance/blob/master/.github/workflows/github_3.yml) |
| Content Management | [![](https://github.com/bloodorangeio/oci-conformance/workflows/github-4/badge.svg)](https://github.com/bloodorangeio/oci-conformance/actions?query=workflow%3Agithub-4) | [Link](https://oci-conformance.s3.amazonaws.com/distribution-spec/github/content-management/report.html) | [Link](https://github.com/bloodorangeio/oci-conformance/blob/master/.github/workflows/github_4.yml) |

### JCR

__Homepage:__ [https://jfrog.com/container-registry/](https://jfrog.com/container-registry/)

__Vendor:__ JFrog

__Verified:__ no

__Test Results:__

| Workflow           | Status                                                                                                                                                               | Test Report                                                                                           | Job Definition                                                                                    |
| ------------------ | -------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- |
| Pull               | [![](https://github.com/bloodorangeio/oci-conformance/workflows/jcr-1/badge.svg)](https://github.com/bloodorangeio/oci-conformance/actions?query=workflow%3Ajcr-1) | [Link](https://oci-conformance.s3.amazonaws.com/distribution-spec/jcr/pull/report.html)               | [Link](https://github.com/bloodorangeio/oci-conformance/blob/master/.github/workflows/jcr_1.yml) |
| Push               | [![](https://github.com/bloodorangeio/oci-conformance/workflows/jcr-2/badge.svg)](https://github.com/bloodorangeio/oci-conformance/actions?query=workflow%3Ajcr-2) | [Link](https://oci-conformance.s3.amazonaws.com/distribution-spec/jcr/push/report.html)               | [Link](https://github.com/bloodorangeio/oci-conformance/blob/master/.github/workflows/jcr_2.yml) |
| Content Discovery  | [![](https://github.com/bloodorangeio/oci-conformance/workflows/jcr-3/badge.svg)](https://github.com/bloodorangeio/oci-conformance/actions?query=workflow%3Ajcr-3) | [Link](https://oci-conformance.s3.amazonaws.com/distribution-spec/jcr/content-discovery/report.html)  | [Link](https://github.com/bloodorangeio/oci-conformance/blob/master/.github/workflows/jcr_3.yml) |
| Content Management | [![](https://github.com/bloodorangeio/oci-conformance/workflows/jcr-4/badge.svg)](https://github.com/bloodorangeio/oci-conformance/actions?query=workflow%3Ajcr-4) | [Link](https://oci-conformance.s3.amazonaws.com/distribution-spec/jcr/content-management/report.html) | [Link](https://github.com/bloodorangeio/oci-conformance/blob/master/.github/workflows/jcr_4.yml) |

### Quay.io

__Homepage:__ [https://quay.io/repository/](https://quay.io/repository/)

__Vendor:__ Red Hat

__Verified:__ no

__Test Results:__

| Workflow           | Status                                                                                                                                                                     | Test Report                                                                                              | Job Definition                                                                                       |
| ------------------ | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| Pull               | [![](https://github.com/bloodorangeio/oci-conformance/workflows/quayio-1/badge.svg)](https://github.com/bloodorangeio/oci-conformance/actions?query=workflow%3Aquayio-1) | [Link](https://oci-conformance.s3.amazonaws.com/distribution-spec/quayio/pull/report.html)               | [Link](https://github.com/bloodorangeio/oci-conformance/blob/master/.github/workflows/quayio_1.yml) |
| Push               | [![](https://github.com/bloodorangeio/oci-conformance/workflows/quayio-2/badge.svg)](https://github.com/bloodorangeio/oci-conformance/actions?query=workflow%3Aquayio-2) | [Link](https://oci-conformance.s3.amazonaws.com/distribution-spec/quayio/push/report.html)               | [Link](https://github.com/bloodorangeio/oci-conformance/blob/master/.github/workflows/quayio_2.yml) |
| Content Discovery  | [![](https://github.com/bloodorangeio/oci-conformance/workflows/quayio-3/badge.svg)](https://github.com/bloodorangeio/oci-conformance/actions?query=workflow%3Aquayio-3) | [Link](https://oci-conformance.s3.amazonaws.com/distribution-spec/quayio/content-discovery/report.html)  | [Link](https://github.com/bloodorangeio/oci-conformance/blob/master/.github/workflows/quayio_3.yml) |
| Content Management | [![](https://github.com/bloodorangeio/oci-conformance/workflows/quayio-4/badge.svg)](https://github.com/bloodorangeio/oci-conformance/actions?query=workflow%3Aquayio-4) | [Link](https://oci-conformance.s3.amazonaws.com/distribution-spec/quayio/content-management/report.html) | [Link](https://github.com/bloodorangeio/oci-conformance/blob/master/.github/workflows/quayio_4.yml) |

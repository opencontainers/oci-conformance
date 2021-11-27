## Introduction
OpenRegistry is an open source, decentralized container registry which is fully compliant with [OCI Container Distribution Specification](https://github.com/opencontainers/distribution-spec/blob/main/spec.md).
The specification provides similar capabilities as that of the Docker Registry HTTP API V2 protocol.

## Why OpenRegistry?
For the longest time, we have relied on DockerHub to host and distribute our container images (both private and public). OpenRegistry tries to provide a decentralized alternative to that by running a community driven container registry, for People by People.

OpenRegitry uses AkashNetwork as it's compute layer and SkyNet for storage. Since AkashNetwork provides a spot like compute market, fault tolerance, Scalability and Resiliency are our priorities from day one.

## Getting Started
Working with OpenRegistry is no different than working with any other container registry. Following are the steps to get started:

### Sign-up:
Head over to [OpenRegistry](https://openregistry.dev) and sign yourself up. The sign process is essential as pushing to container repositories is a restricted operation and requires proper authorization.
> Currently we're only accepting registrations for a closed Beta program, Kindly register for Beta [here](https://parachute.openregistry.dev)

### Push an Image:
When using Docker CLI, the images are pushed to DockerHub by default. For Pushing images to OpenRegistry instead, follow the below steps:
* change the name of your image, e.g if you have an image named alpine:latest, change it like so:
```bash
docker tag alpine:latest openregistry.dev/janedoe/alipne:latest
docker push openregistry.dev/janedoe/alpine:latest
```

### Pull an Image:
Assuming you've pushed an image using the above method:
```bash
docker pull openregistry.dev/janedoe/alpine:latest
```

### How to Run this project locally:
OpenRegistry is not Go Gettable right now because of a dependency issue with Go-Skynet. To build this project locally, please use the following method:
```bash
git clone https://github.com/containerish/OpenRegistry.git
make mod-fix
go build
```

### Recreating the Results: 
https://raw.githubusercontent.com/containerish/OpenRegistry/master/scripts/recreating.sh

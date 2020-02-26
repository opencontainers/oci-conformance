# Trow OCI Conformance Tests

## What is Trow

Trow is a container registry and image management solution for Kubernetes. It is designed to be a
secure and fast way to distribute containers within a Kubernetes cluster. Planned features include
support for advanced auditing and authentication, plus ahead-of-time distribution of images.

Trow is written in Rust for speed and reliability.

## Recreating the Results

To verify the conformance tests, run the below bash script. 

The script will checkout the open source Trow registry from GitHub, and build a Docker image. The
image is then run in the background. The script will then checkout and build the OCI distribution
tests, again as a Docker image. The tests are then run against Trow registry and the results written
out.

The script requires that bash, docker and git are installed and available.

Adrian Mouat

```
#!/bin/bash

# This script will build and run the Trow registry, then build and run
# the OCI conformance tests against it. 
# It depends on bash, docker and git being available.

set -eu 

spec_name=distribution-spec
spec_version=v1.0
prod_name=example
trow_version=v0.2.0

# check out trow repo
rm -rf trow-tmp && git clone --branch "$trow_version" git://github.com/ContainerSolutions/trow.git trow-tmp
(cd trow-tmp/docker && ./build.sh) # Builds containersol/trow:default image
rm -rf trow-tmp
# start trow
docker network create trow-conf-test || true
docker stop trow || true
docker rm trow || true
docker run --net trow-conf-test --name trow -d containersol/trow:default --no-tls

# check out conformance repo
rm -rf conf-tmp && git clone https://github.com/opencontainers/${spec_name}.git conf-tmp
(cd conf-tmp && docker build -t conformance:latest -f Dockerfile .)
rm -rf conf-tmp 
# Would delete results as well, but they're owned by root :(
docker run --rm \
  --net trow-conf-test \
  -v $(pwd)/results:/results \
  -w /results \
  -e OCI_ROOT_URL="http://trow:8000" \
  -e OCI_NAMESPACE="myorg/myrepo" \
  -e OCI_DEBUG="true" \
  conformance:latest
cp results/report.html ./
cp results/junit.xml ./
docker stop trow && docker rm trow
docker network rm trow-conf-test
```

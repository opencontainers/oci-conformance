# Trow OCI Conformance Tests

## What is Trow

Trow is a container registry and image management solution for Kubernetes. It is designed to be a
secure and fast way to distribute containers within a Kubernetes cluster. Planned features include
support for advanced auditing and authentication, plus ahead-of-time distribution of images.

Trow is written in Rust for speed and reliability.

## Recreating the Results

To verify the conformance tests, please run the script `run-conformance.sh` that is included in this
directory. 

The script will checkout the open source Trow registry from GitHub, and build a Docker image. The
image is then run in the background. The script will then checkout and build the OCI distribution
tests, again as a Docker image. The tests are then run against Trow registry and the results written
out.

The script requires that bash, docker and git are installed and available.

Adrian Mouat

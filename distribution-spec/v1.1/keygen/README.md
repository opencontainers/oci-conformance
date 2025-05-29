# Keygen

Keygen implements a read-only OCI distribution engine, for both public and
licensed container images. Below are instructions on how to run the OCI
conformance test against Keygen's API.

## Running the test

Clone OCI distribution spec repository:

```
git clone https://github.com/opencontainers/distribution-spec -b v1.1.0
```

Build conformance test binary:

```sh
cd distribution-spec/conformance

go test -c
```

Set test environment variables:

```sh
export OCI_ROOT_URL=https://oci.pkg.keygen.sh
export OCI_USERNAME=license
export OCI_PASSWORD=9EBD41-A4BCBF-4667C5-286307-C29748-V3
export OCI_NAMESPACE=oci/conformance
export OCI_MANIFEST_DIGEST=sha256:6c455a830fd8b80063189e4007b2d0d83119945802f8559f10a1c19fdcea9c50
export OCI_BLOB_DIGEST=sha256:a948904f2f0f479b8f8197694b30184b0d2ed1c1cd2a1ec0fb85d299a192a44
export OCI_TAG_NAME=latest
export OCI_TEST_PULL=1
export OCI_TEST_PUSH=0
export OCI_TEST_CONTENT_DISCOVERY=0
export OCI_TEST_CONTENT_MANAGEMENT=0
```

Run the conformance test:

```sh
./conformance.test
```

![Cloudsmith Logo](https://cloudsmith.com/assets/brand/cloudsmith-square-logo.svg)
# Cloudsmith

## Documentation

To learn how to store your OCI artifacts in Cloudsmith see our [documentation for OCI](https://help.cloudsmith.io/docs/oci-repository).

## Running the tests

To run the tests go to [Cloudsmith](https://cloudsmith.com/) and sign up for an acccount.

Create an organization and a repository to run the test against

Clone this project:
```bash
git clone https://github.com/opencontainers/distribution-spec -b v1.1.0
```

Set the following environment variables
```bash
export OCI_ROOT_URL=https://docker.cloudsmith.io
export OCI_USERNAME=<your-cloudsmith-user>
export OCI_PASSWORD=<your-cloudsmith-token>
export OCI_NAMESPACE=<your-cloudsmith-org>/<your-cloudsmith-repo>/oci-conformance-testing
export OCI_CROSSMOUNT_NAMESPACE=<your-cloudsmith-org>/<your-cloudsmith-repo>/oci-testing-crossmount
export OCI_TEST_PULL=1
export OCI_TEST_PUSH=1
export OCI_TEST_CONTENT_DISCOVERY=1
export OCI_TEST_CONTENT_MANAGEMENT=0
```

Run the tests:
```bash
./conformance.test
```

This should show you the following result:
```bash
Ran 64 of 79 Specs in 126.355 seconds
SUCCESS! -- 64 Passed | 0 Failed | 0 Pending | 15 Skipped
PASS
```
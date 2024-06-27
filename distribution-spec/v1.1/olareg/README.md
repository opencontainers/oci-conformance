# olareg

## Running olareg as a Container

```shell
docker run -d --rm -p 127.0.0.1:5000:5000 ghcr.io/olareg/olareg:edge serve
```

## Run the Conformance Tests

```shell
cd distribution-spec
make conformance-binary
(
  [ -d conformance/results ] || mkdir -p conformance/results
  cd conformance/results
  OCI_ROOT_URL="http://localhost:5000" \
  OCI_NAMESPACE="myorg/myrepo" \
  OCI_CROSSMOUNT_NAMESPACE="myorg/other" \
  OCI_TEST_PULL=1 \
  OCI_TEST_PUSH=1 \
  OCI_TEST_CONTENT_DISCOVERY=1 \
  OCI_TEST_CONTENT_MANAGEMENT=1 \
  OCI_HIDE_SKIPPED_WORKFLOWS=0 \
  OCI_DEBUG=0 \
  OCI_DELETE_MANIFEST_BEFORE_BLOBS=1 \
    ../../output/conformance.test
)
```

This should produce the following results:

```text
Ran 74 of 80 Specs in 0.278 seconds
SUCCESS! -- 74 Passed | 0 Failed | 0 Pending | 6 Skipped
PASS
```

More details can be found in the resulting `report.html`.

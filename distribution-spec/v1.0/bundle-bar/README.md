Sign up for an account at https://bundle.bar/.

Create a new token in your account, and copy it into
your clipboard.

Set your Bundle Bar username and token as
environment variables:

```
BB_USER="<username>"
BB_TOKEN="<token>"
```

Download the OCI Distribution Specification
`v1.0.1` source:

```
git clone https://github.com/opencontainers/distribution-spec -b v1.0.1
```

Build conformance binary (requires Go):

```
cd distribution-spec/
make conformance-binary
```

Set various environment variables:

```
export OCI_ROOT_URL="https://bundle.bar"
export OCI_USERNAME="${BB_USER}"
export OCI_PASSWORD="${BB_TOKEN}"
export OCI_NAMESPACE="u/${BB_USER}/conformance/primary"
export OCI_CROSSMOUNT_NAMESPACE="u/${BB_USER}/conformance/secondary"
export OCI_TEST_PULL=1
export OCI_TEST_PUSH=1
export OCI_TEST_CONTENT_DISCOVERY=1
export OCI_TEST_CONTENT_MANAGEMENT=1
```

Run the conformance tests:

```
output/conformance.test
```

This should produce the following results:

```
Ran 59 of 62 Specs in 36.562 seconds
SUCCESS! -- 59 Passed | 0 Failed | 0 Pending | 3 Skipped
```

More details can be found in the resulting `report.html`.

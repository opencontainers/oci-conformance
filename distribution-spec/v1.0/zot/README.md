Download the Zot `v1.4.1` source:

```
git clone https://github.com/project-zot/zot.git -b v1.4.1
```

Build Zot binary (requires Go):

```
cd zot/
make binary
```

Create valid configuration file:

```
cat > config.yml <<EOL
http:
  address: 127.0.0.1
  port: 5000
storage:
  rootDirectory: "${PWD}/registry"
  gc: false
  dedupe: false
EOL
```

Run Zot server:

```
bin/zot serve config.yml
```

In another terminal, download the OCI Distribution
Specification `v1.0.1` source:

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
export OCI_ROOT_URL="http://127.0.0.1:5000"
export OCI_NAMESPACE="conformance/primary"
export OCI_CROSSMOUNT_NAMESPACE="conformance/secondary"
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
Ran 59 of 62 Specs in 0.058 seconds
SUCCESS! -- 59 Passed | 0 Failed | 0 Pending | 3 Skipped
```

More details can be found in the resulting `report.html`.

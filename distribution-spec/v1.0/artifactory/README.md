![Artifactory Logo](https://media.jfrog.com/wp-content/uploads/2022/05/24134044/artifactory-Icon.png)
# jFrog Artifactory



### Get Artifactory
#### * Try out Artifactory - start [here](https://jfrog.com/start-free/)
#### * Create a local OCI repository called oci-local

```bash
docker login <artifactory-url>
```
---

## Run the conformance test

---


#### Clone OCI Distribution Spec

```bash
git clone https://github.com/opencontainers/distribution-spec
```

#### Build conformance binary
```bash
cd distribution-spec/conformance
go test -c
```

#### Set various environment variables
```bash
export OCI_ROOT_URL=<artifactory-url>
export OCI_USERNAME=<artifactory-user>
export OCI_PASSWORD=<artifactory-token>
export OCI_NAMESPACE=oci-local
export OCI_TEST_PULL=1
export OCI_TEST_PUSH=1
export OCI_TEST_CONTENT_DISCOVERY=1
export OCI_TEST_CONTENT_MANAGEMENT=1
```

#### Run conformance test
```bash
./conformance.test
```

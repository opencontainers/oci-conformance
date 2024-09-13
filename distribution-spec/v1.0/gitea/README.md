# Gitea

![Gitea Logo](https://raw.githubusercontent.com/go-gitea/gitea/main/assets/logo.svg)

## Configure Gitea test instance

Start Gitea server as docker container:

```sh
docker run --rm -e USER_UID=1000 -e USER_GID=1000 -p 3000:3000 gitea/gitea:1.22
```

In browser open <http://localhost:3000> and enter administrator username and password and complete setup wizard.

Create `test-org` organization.

## Run the conformance test

### Clone OCI Distribution Spec

```bash
git clone https://github.com/opencontainers/distribution-spec -b v1.0.1
```

### Build conformance binary

```bash
cd distribution-spec
make conformance-binary
```

### Set various environment variables

```bash
export OCI_ROOT_URL=http://localhost:3000
export OCI_USERNAME=<username>
export OCI_PASSWORD=<password>
export OCI_NAMESPACE=test-org/test
export OCI_CROSSMOUNT_NAMESPACE=test-org/test-cross-mount
export OCI_TEST_PULL=1
export OCI_TEST_PUSH=1
export OCI_TEST_CONTENT_DISCOVERY=1
export OCI_TEST_CONTENT_MANAGEMENT=1
```

### Run conformance test

```bash
./conformance.test
```

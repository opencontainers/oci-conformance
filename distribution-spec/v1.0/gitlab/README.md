![logo](https://about.gitlab.com/images/press/logo/svg/gitlab-logo-100.svg)

# GitLab Container Registry

## Documentation

To learn how to use the GitLab Container Registry, see the [user documentation](https://docs.gitlab.com/ee/user/packages/container_registry/). [Administration instructions](https://docs.gitlab.com/ee/administration/packages/container_registry.html) are also available.

## Running the Conformance Tests

GitLab runs the OCI Distribution conformance tests against the GitLab.com Container Registry daily.
See the [latest results](https://gitlab.com/gitlab-com/registry-oci-conformance). To replicate the tests using GitLab CI/CD, either:

- Clone this project.
- Use the provided `gitlab-ci.yml` file.

For self-managed instances, or if you want to execute the tests against GitLab.com from a local machine, please follow the instructions below.

### Prerequisites and Assumptions

For the instructions presented in this document, we assume that:

- You are using GitLab SaaS. If you are using a self-managed instance, please update the registry address (`registry.gitlab.com`) accordingly.

- You have a GitLab project against which you can execute the tests. We will assume the project path is `my-group/my-project`.

- A user with a [Developer+](https://docs.gitlab.com/ee/user/permissions.html#project-members-permissions) role in the project. We will assume the username is `myuser`.

- A valid authentication token with the `api` scope. See the [authentication documentation](https://docs.gitlab.com/ee/user/packages/container_registry/authenticate_with_container_registry.html) for more information. We will assume the token value is `gltoken`.

### Instructions

1. Clone the OCI Distribution Spec repository:

   ```bash
   git clone https://github.com/opencontainers/distribution-spec -b v1.0.1
   ```

2. Build the conformance test binary:

   ```bash
   cd distribution-spec/conformance
   go test -c
   ```

3. Set environment variables:

   ```bash
   export OCI_ROOT_URL=https://registry.gitlab.com
   export OCI_NAMESPACE=my-group/my-project/repo-a
   export OCI_CROSSMOUNT_NAMESPACE=my-group/my-project/repo-b
   export OCI_USERNAME=myuser
   export OCI_PASSWORD=gltoken
   export OCI_TEST_PULL=1
   export OCI_TEST_PUSH=1
   export OCI_TEST_CONTENT_DISCOVERY=1
   export OCI_TEST_CONTENT_MANAGEMENT=1
   export OCI_AUTOMATIC_CROSSMOUNT=1
   export OCI_DELETE_MANIFEST_BEFORE_BLOBS=1
   ```

4. Run the tests:

   ```bash
   ./conformance.test
   ```

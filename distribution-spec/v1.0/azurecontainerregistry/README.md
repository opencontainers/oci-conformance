Prerequisites: Install the [Azure CLI](https://docs.microsoft.com/en-us/cli/azure/) and create an Azure subscription.

Create a resource group using the `az group create --name myResourceGroup --location eastus` command.

Create a registry by running `az acr create --resource-group myResourceGroup --name <testacr> --sku Basic`.

Enable basic credentials `az acr update -n <testacr> --admin-enabled true`.

Using `az acr credential show -n <testacr>` take note of the username and one of the admin credentials.

Get into the [distribution-spec folder](https://github.com/opencontainers/distribution-spec/tree/main/conformance), by running `cd distribution-spec/conformance`.
And build the test binary `go test -c`. 


Set the required environment variables using the previously obtained credentials.

```
export OCI_ROOT_URL="https://<testacr>.azurecr.io"
export OCI_USERNAME="<username>"
export OCI_PASSWORD="<password>"
export OCI_NAMESPACE="conformance"
export OCI_TEST_PULL=1
export OCI_TEST_PUSH=1
export OCI_TEST_CONTENT_DISCOVERY=1
export OCI_TEST_CONTENT_MANAGEMENT=1
```

And finally run the conformance tests by running `./conformance.test`.

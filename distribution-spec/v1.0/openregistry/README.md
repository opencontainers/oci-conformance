
#### Clone OpenRegistry repo
```bash
git clone git@github.com:containerish/OpenRegistry.git
```

#### Build and run OpenRegistry binary(required Go)
```bash
cd OpenRegistry
export $(cat env-vars.example |xargs)
make mod-fix
go build -o OpenRegistry && ./OpenRegistry
```

#### In another terminal, clone OCI Distribution Spec ```v1.0.1```

```bash
git clone https://github.com/opencontainers/distribution-spec -b v1.0.1
```

#### Build conformance binary
```bash
cd distribution-spec
make conformance-binary
```

#### Set various environment variables
```bash
export OCI_ROOT_URL=http://0.0.0.0:5000
export OCI_NAMESPACE=johndoe/dummmy-server
export OCI_CROSSMOUNT_NAMESPACE=johndoe/dummmy-server-cross-mount
export OCI_USERNAME=johndoe
export OCI_PASSWORD=Qwerty@123
export OCI_TEST_PULL=1
export OCI_TEST_PUSH=1
export OCI_TEST_CONTENT_DISCOVERY=1
export OCI_TEST_CONTENT_MANAGEMENT=1
```

#### Create an OpenRegistry user since PUSH required authorisation
```bash 
curl -XPOST -d '{"email":"johndoe@example.com","username":"johndoe","password":"Qwerty@123"}' "http://0.0.0.0:5000/auth/signup"
```
#### Run conformance test
```bash
output/conformance.test
```
#### This should produce the following results
```bash
Ran 59 of 62 Specs in 53.443 seconds
SUCCESS! -- 59 Passed | 0 Failed | 0 Pending | 3 Skipped
PASS
```
more details can be found in the resulting ```report.html```.

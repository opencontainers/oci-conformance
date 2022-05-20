# Keppel

## Prepare

Clone the source code.

```ShellSession
$ git clone https://github.com/sapcc/keppel.git
Cloning into 'keppel'...
remote: Enumerating objects: 14513, done.
remote: Counting objects: 100% (3578/3578), done.
remote: Compressing objects: 100% (1996/1996), done.
remote: Total 14513 (delta 2418), reused 2297 (delta 1487), pack-reused 10935
Receiving objects: 100% (14513/14513), 12.45 MiB | 4.99 MiB/s, done.
Resolving deltas: 100% (8670/8670), done.
```

Start Postgresql database, build and start Keppel.
This requires Golang 1.16+, postgres and openssl to be installed.

```ShellSession
$ cd keppel
$ bash testing/with-postgres-db.sh make run-api-for-conformance-test
>> Configuring PostgreSQL...
>> Starting PostgreSQL...
pg_ctl: another server might be running; trying to start server anyway
>> Running command: make run-api-for-conformance-test...
go build -mod vendor -ldflags '-s -w -X github.com/sapcc/keppel/internal/keppel.Version=1.0.0-dev' -o build/keppel .
openssl genrsa -out testing/conformance-test/privkey.pem 4096
Generating RSA private key, 4096 bit long modulus (2 primes)
.........++++
...............................................................................................++++
e is 65537 (0x010001)
Ready to run conformance test
set -euo pipefail && source testing/conformance-test/env.sh && /home/sandro/src/github.com/sapcc/keppel/build/keppel server api
2022/05/13 11:51:24 INFO: starting keppel-api 1.0.0-dev
2022/05/13 11:51:25 INFO: listening on :8080
```

In another terminal, download the OCI Distribution Specification source and build it.

```ShellSession
$ git clone https://github.com/opencontainers/distribution-spec
Cloning into 'distribution-spec'...
remote: Enumerating objects: 1294, done.
remote: Counting objects: 100% (331/331), done.
remote: Compressing objects: 100% (201/201), done.
remote: Total 1294 (delta 166), reused 251 (delta 119), pack-reused 963
Receiving objects: 100% (1294/1294), 761.38 KiB | 2.15 MiB/s, done.
Resolving deltas: 100% (591/591), done.
$ cd distribution-spec
$ make conformance-binary
cd conformance && \
        CGO_ENABLED=0 go test -c -o /home/user/src/github.com/opencontainers/distribution-spec/output//conformance.test \
                --ldflags="-X github.com/opencontainers/distribution-spec/conformance.Version=7db04aef5a9ab17763d6ac2f9bade915edeef3e1"
go: downloading github.com/bloodorangeio/reggie v0.5.0
go: downloading github.com/onsi/ginkgo v1.16.4
go: downloading github.com/onsi/gomega v1.16.0
go: downloading github.com/go-resty/resty/v2 v2.1.0
go: downloading github.com/mitchellh/mapstructure v1.1.2
go: downloading golang.org/x/net v0.0.0-20210428140749-89ef3d95e781
go: downloading github.com/nxadm/tail v1.4.8
go: downloading golang.org/x/sys v0.0.0-20210423082822-04245dca01da
go: downloading github.com/fsnotify/fsnotify v1.4.9
```

Set environment variables

```ShellSession
export OCI_ROOT_URL="http://localhost:8080"
export OCI_NAMESPACE="conformance-test/oci"
export OCI_USERNAME="johndoe"
export OCI_PASSWORD="SuperSecret"
export OCI_TEST_PULL=1
export OCI_TEST_PUSH=1
export OCI_TEST_CONTENT_DISCOVERY=1
export OCI_TEST_CONTENT_MANAGEMENT=1
export OCI_CROSSMOUNT_NAMESPACE="$OCI_NAMESPACE"
export OCI_HIDE_SKIPPED_WORKFLOWS=0
export OCI_DELETE_MANIFEST_BEFORE_BLOBS=1
```

Run the conformance test

```ShellSession
$ output/conformance.test
...
```

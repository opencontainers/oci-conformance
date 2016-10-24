# OCI Certification Program (DRAFT)

This is a draft process for the OCI Certification Program. There are two official certified trademarks:

1. **OCI Certified Runtime**
2. **OCI Certified Image**

## Process

### (1) Apply

The first step is to apply and contact the OCI Certification Program.

### (2) Test

The second step is to test your OCI Certified solution, whether it's an OCI Certified Runtime or OCI Certified Image.

### (3) Certify

If your solution passes testing, you need to publish your test process/results on GitHub and the OCI Certification Program will issue an official certificate of approval.

### (4) Promote

Promote your OCI Certified solution using official OCI marks on your marketing material.

## OCI Certified Runtime (Runtime Specification)

An OCI Certified Runtime has three main requirements to fulfill: A) bundle B) lifecycle C) configuration:

* create a container by an ‘oci bundle’ directly: `create <container-id> <path-to-bundle>`
* start a related container and get the EXPECTED running status
* kill/delete a related container AS EXPECTED

The OCI runtime tool provides an "oci testing bundle" and it looks like this:
* "oci testing bundle"
  * config.json
  * rootfs
    * runtimetest
    * config.json
    * [other necessary binaries from busybox rootfs]

Strictly speaking, to test an OCI Certified Runtime (RUNTIME-X) we MUST verify:

1. If RUNTIME-X creates a container ‘oci testing bundle’, it means RUNTIME-X **complaint to requirement A (and part of B)**.
2. If RUNTIME-x could then *start/kill/delete* this container, it means RUNTIME-X is **compliant to requirement B**.
3. In ‘oci testing bundle’, config.json and rootfs/config.json are the same.
4. In config.json, the main process is `runtimetest`, so when a RUNTIME-X start this bundle, ‘runtimetest’ will run.
5. What `runtimetest` does is parsing configurations in rootfs/config.json and comparing the values with its runtime environment.
6. If all the values match, it means RUNTIME-X **compliant to requirement C**.
 
Thus we confirm that RUNTIME-X is an OCI Certified Runtime.

More information on how this works via a CLI command is here (see test_runtime.sh):
https://github.com/opencontainers/runtime-tools#testing-oci-runtimes

## OCI Certified Image (Image Specification)

To validate an image, you can use the **oci-image-validate** command:

```sh
$ oci-image-validate --type imageLayout --ref latest <IMAGE-X>
<IMAGE-X>: OK
```

Thus we confirm that IMAGE-X is an OCI Certified Image.

## Contact

To reach out to the OCI Certification team, email [certification@opencontainers.org](mailto:certification@opencontainers.org)

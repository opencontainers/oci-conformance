# OCI Certification Program (DRAFT)

This is a draft process for the OCI Certification Program. 
There are types of products that can be certified:

1. **OCI Certified Runtime**
2. **OCI Certified Image**

Organizations or individuals wishing to use any of the below registered trademarks, trademarks pending registration, or trademarks in use must complete the appropriate certification program.
* OCI
* OCI Certified
* Open Container Initiative
* Open Container Format


## Process

An organization or individual may apply for OCI Certification for their product regardless of their membership in the Open Container Initiative organization.  

### (1) Apply

The first step is to apply and contact the OCI Certification Program.

Information expected in an application will include:
* Organization name
* Organization website (if public)
* Intended certification program
* Product name
* Product website (if public)
* Primary contact name, email, and phone
* Secondary contact name, email, and phone
* Approximate intended date to publish test results and steps for peer review

### (2) Test

The second step is to test your product, whether it is intended to be an OCI Certified Runtime or an OCI Certified Image.

### (3) Publish Results

The next step is to publish your step-by-step test process and results.  
The current location for publishing test process and results is GitHub.

### (4) Peer Verification

One or more members of the OCI Certification Program Working Group (Cert WG) may attempt to reproduce your test results and contact you with questions.

### (5) Certify

If your product passes testing and peer review, the OCI Certification Program will issue an official certificate of approval.

### (6) Promote

Promote your OCI Certified product using official OCI marks on your marketing material.


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


## Versions

Each certification program will have versioning aligned with the major (and possibly minor) versions of the appropriate specification.
For example, 
* There may be an OCI Certified Runtime **v1.0** program associated with the OCI Runtime specification v1.0.x
* There may be an OCI Certified Runtime **v2.0** program associated with the OCI Runtime specification v2.0.x
* There may be an OCI Certified Image **v1.1** program associated with the OCI Runtime specification v1.1.x

To maintain real interoperability, at the discretion of the OCI Trademark Board, older certification program verions may be discontinued.
Products that were certified under older versions may be asked to certifiy on a supported version or discontinue using OCI marks.


## Contact

To reach out to the OCI Certification team, email [certification@opencontainers.org](mailto:certification@opencontainers.org)

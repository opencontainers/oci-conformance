# OCI Certification Program (DRAFT)

This is a *draft process* for the OCI Certification Program. The OCI Certifcation Working Group (CertWG) is responsible for drafting the process around OCI Certification. 

## Membership

The CertWG membership is currently comprised of:

* Amazon: Alex Talsma
* CoreOS: Alex Polvi
* Docker: Brian Goff, Stephen Walli
* Google: Sarah Novotny
* Fujitsu: Kamezawa Hiroyuki
* Huawei: Liang Chenye, Nan Zhou
* IBM: **Jeff Borek (Cert WG Co-Chair)**, Phil Estes
* Intel: David Lyle
* Microsoft: **Rob Dolin (Cert WG Co-Chair)**
* Red Hat: Mrunal Patel
* Replicated: Mark Campbell

## Certification

There are currently two types of products that can be certified:

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
* Intended certification program (ex: OCI Certified Image v1.0)
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

One or more members of the OCI Certification Program Working Group (Cert WG) may attempt to reproduce the test results for your product and may contact you with questions. No company can review its own results.

### (5) Certify

If your product passes testing and peer review, the OCI Certification Program will issue an official certificate of approval.

### (6) Promote

Promote your OCI Certified product using official OCI marks on your marketing material.

## OCI Certified Runtime (Runtime Specification)

https://github.com/opencontainers/runtime-tools

## OCI Certified Image (Image Specification)

https://github.com/opencontainers/image-tools

## Versions

Each certification program will have versioning aligned with the major (and possibly minor) versions of the **appropriate specification**.
For example, 
* There may be an OCI Certified Runtime **v1.0** program associated with the OCI Runtime specification v1.0.x
* There may be an OCI Certified Runtime **v2.0** program associated with the OCI Runtime specification v2.0.x
* There may be an OCI Certified Image **v1.1** program associated with the OCI Image Format specification v1.1.x

## Deprecation
To maintain real interoperability, at the discretion of the OCI Trademark Board, older certification program verions MAY be discontinued.
If a certification program version is discontinued, products that were certified under older versions MAY be asked to either:
* Certifiy on a supported version
* Discontinue using OCI marks

Before a certification program version is discontinued, the below criteria SHOULD be met:
* A replacement certification program is in-place (i.e. OCI Certified Image v1.1 replaced with OCI Certified Image v2.0)
* An emailed notice is provided to contacts for all known certified products
* A minimum of six months has transpired since the replacement certification program became available

At the discretion of the OCI Trademark Board, multiple versions of certification programs MAY operate simultaneously for extended periods of time.  
For example, OCI Certified Runtime v1.1 and OCI Certified Runtime v2.0 could both be deemed valuable.

## Contact

To reach out to the OCI Certification team, email [certification@opencontainers.org](mailto:certification@opencontainers.org)

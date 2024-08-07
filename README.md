# OCI Certified

[![Netlify Status](https://api.netlify.com/api/v1/badges/3e16a8a1-e338-4050-a39d-426b33f5292a/deploy-status)](https://app.netlify.com/sites/oci-conformance/deploys)

This repo contains everything related to certifying that your product or service
is fully compliant with the specs defined by the [Open Container Initiative](https://www.opencontainers.org/) (OCI).

The following specifications are currently tested for compliance:

- [OCI Distribution Specification](https://github.com/opencontainers/distribution-spec)

## OCI Certified Conformance Program

* [Instructions](instructions.md)

* [Terms and Conditions](./terms-conditions/OCI_Certified_Terms.md)

* [Participation Form](./participation-form/OCI_Certified_Form.md)

* [Brand Guidelines](OCI-certified-brand-guide-v1.pdf)

* [Marks](https://github.com/opencontainers/artwork/tree/master/certified)

* [Frequently Asked Questions](faq.md)

## Development

*Requires Docker*

To run the site locally, build and run from the provided [Dockerfile](./Dockerfile):
```
docker build -t oci-conformance . && \
    docker run --rm -p 8080:8080 oci-conformance
```

The site can then be accessed at [http://localhost:8080/](http://localhost:8080/)

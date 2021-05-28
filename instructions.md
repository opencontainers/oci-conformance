# How to submit conformance results

## Running the tests

Each spec will provide its own testing instructions, and each will produce 
the following files that contain the test results and output:
- `report.html`
- `junit.xml`

### OCI Distribution Specification

Please see instructions [here](https://github.com/opencontainers/distribution-spec/blob/main/conformance/README.md).

## Uploading

Prepare a PR to
[https://github.com/opencontainers/oci-conformance](https://github.com/opencontainers/oci-conformance).
Here are [directions](https://help.github.com/en/articles/creating-a-pull-request-from-a-fork) to
prepare a pull request from a fork.
In the descriptions below, `$spec.X.Y` refers to the spec and its major and minor
version, and `$dir` is a short subdirectory name to hold the results for your
product.  Examples would be `gcr` or `dockerhub`.

Description: `Conformance results for $spec/vX.Y/$dir`

### Contents of the PR

For simplicity you can submit the tarball or extract the relevant information from the tarball to compose your submission. 

```
$spec/vX.Y/$dir/README.md: Description of how to reproduce your results.
$spec/vX.Y/$dir/report.html: Human-readable HTML test report.
$spec/vX.Y/$dir/junit.xml: Machine-readable JUnit test report.
$spec/vX.Y/$dir/PRODUCT.yaml: See below.
```

#### PRODUCT.yaml

This file describes your product. It is YAML formatted with the following root-level fields. Please fill in as appropriate.

| Field               | Description |
| ------------------- | ----------- |
| `vendor`            | Name of the legal entity that is certifying. This entity must have a signed participation form on file with the OCI  |
| `name`              | Name of the product being certified. |
| `version`           | The version of the product being certified (not the version of OCI spec). |
| `website_url`       | URL to the product information website |
| `repo_url`          | If your product is open source, this field is necessary to point to the primary GitHub repo containing the source. It's OK if this is a mirror. OPTIONAL  |
| `documentation_url` | URL to the product documentation |
| `product_logo_url`  | URL to the product's logo, (must be in SVG, AI or EPS format -- not a PNG -- and include the product name). OPTIONAL. If not supplied, we'll use your company logo. Please see logo guidelines (TODO: link) |
| `type`              | Is your product a distribution, hosted platform, or installer (see [definitions](https://github.com/opencontainers/oci-conformance/blob/main/faq.md#what-is-a-distribution-and-what-is-a-platform)) |
| `description` | One sentence description of your offering |

Examples below are for a fictional OCI implementation called _Turbo
Encabulator_ produced by a company named _Yoyodyne_.

```yaml
vendor: Yoyodyne
name: Turbo Encabulator
version: v1.7.4
website_url: https://yoyo.dyne/turbo-encabulator
repo_url: https://github.com/yoyo.dyne/turbo-encabulator
documentation_url: https://yoyo.dyne/turbo-encabulator/docs
product_logo_url: https://yoyo.dyne/assets/turbo-encabulator.svg
type: distribution
description: 'The Yoyodyne Turbo Encabulator is a superb OCI distribution for all of your Encabulating needs.'
```

## Amendment for Private Review

If you need a private review for an unreleased product, please email a zip file containing what you would otherwise submit
as a pull request to certification@opencontainers.org. We'll review and confirm that you are ready to be OCI Certified
as soon as you open the pull request. We can then often arrange to accept your pull request soon after you make it, at which point you become OCI Certified.

## Review

A reviewer will shortly comment on and/or accept your pull request, following this [process](reviewing.md).
If you don't see a response within 3 business days, please contact certification@opencontainers.org.

## Example Script

Combining the steps provided here, the process looks like this (Example: `distribution-spec/v1.0`):

```
spec_name=distribution-spec
spec_version=v1.0
prod_name=example

rm -rf tmp && git clone https://github.com/opencontainers/${spec_name}.git tmp
(cd tmp && docker build -t conformance:latest -f Dockerfile.conformance .)
rm -rf tmp results
docker run --rm \
  -v $(pwd)/results:/results \
  -w /results \
  -e OCI_ROOT_URL="https://r.myreg.io" \
  -e OCI_NAMESPACE="myorg/myrepo" \
  -e OCI_USERNAME="myuser" \
  -e OCI_PASSWORD="mypass" \
  -e OCI_DEBUG="true" \
  conformance:latest

mkdir -p ./${spec_name}/${spec_version}/${prod_name}
cp ./results/* ./${spec_name}/${spec_version}/${prod_name}/
rm -rf results

cat << EOF > ./${spec_name}/${spec_version}/${prod_name}/PRODUCT.yaml
vendor: Yoyodyne
name: Turbo Encabulator
version: v1.7.4
website_url: https://yoyo.dyne/turbo-encabulator
repo_url: https://github.com/yoyo.dyne/turbo-encabulator
documentation_url: https://yoyo.dyne/turbo-encabulator/docs
product_logo_url: https://yoyo.dyne/assets/turbo-encabulator.svg
type: distribution
description: 'The Yoyodyne Turbo Encabulator is a superb OCI distribution for all of your Encabulating needs.'
EOF
```

## Issues

If you have problems certifying that you feel are an issue with the conformance
program itself (and not just your own implementation), you can file an issue in
the [repository](https://github.com/opencontainers/oci-conformance).

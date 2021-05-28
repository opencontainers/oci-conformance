# Frequently Asked Questions

## What is the cost of certification?

For members (TODO: link) of the Open Container Initiative, there is no charge for
certification. There is also no charge for non-profit organizations. For commercial organizations that wish to
certify but don't want to become an OCI member, the fee is the same as joining (TODO: link) the OCI.

## What about community distributions?

A community distribution does not have a company behind it.
For the purposes of certification, we treat community distributions as non-profit organizations, so
there is no charge. However, we do require an individual to complete the certification
[agreement](./participation-form/OCI_Certified_Form.md) so
that we have an official contact (or multiple contacts) if your software falls out of compliance.

## What versions of OCI specs can be certified?

You can certify the currently released version of each spec and the two versions before that. 
The currently released version is the latest stable tag in each spec repo. 
Already certified implementations remain certified as long as a newer version is certified at least once a year after the initial certification.

## What is a distribution, hosted platform, and an installer?

From the bottom of the OCI Distributions & Platforms spreadsheet (TODO: link):

* A **vendor** is an organization providing an OCI distribution, hosted platform, or installer.
* A **product** is a distribution, hosted platform, or installer provided by a vendor.	
* A **distribution** is software based on OCI that can be installed by an end user on to a public cloud or bare metal and includes patches to the upstream codebase.	
* A **hosted** platform is an OCI service provided and managed by a vendor.	
* An **installer** downloads and then installs vanilla upstream OCI.	

## Do I need to re-submit results if we rebrand our product?

No. If the software is the same, and just the name has changed, you just need to submit a revised Participation Form available at https://github.com/opencontainers/oci-conformance/blob/main/participation-form/OCI_Certified_Form.md that includes the new name. Please also open a pull request to update the name in your PRODUCT.yaml file. We do ask that you send us the new Participation Form **prior** to announcing the name change. You can submit the pull request after the announcement, if necessary.

## Is a participation form required per company or per product?

Per product. Each separate product (i.e., different product name) from your company requires a different [participation form](https://github.com/opencontainers/oci-conformance/blob/main/participation-form/OCI_Certified_Form.md). We don't need a new form for new versions of an existing product.

## Are there any limitations regarding the logo?

Yes. First, we need it in SVG, AI, or EPS format. If you don't have a product-specific logo, it is fine to reuse your company one. However, we have a set of requirements (TODO: link) including that it incorporate the name of the product or company in English (it's fine to also include text in other languages) and that the logo not be reversed.

Second, your logo can't be derivative of or include portions of the OCI logo (TODO: link), in order to preserve the value of the OCI logo. Of course, under the terms of the OCI Certified agreement, you're welcome and encouraged to use the [OCI Certified logo](https://github.com/opencontainers/artwork/tree/master/certified) to indicate that your product is OCI Certified.

## Can I provide a link to the installation directions?

No. We need the Readme of your conformance submission to contain more detailed directions on how to install and configure your OCI implementation. The reason is that we need your users to be able to replicate the installation so that they can re-run the tests in the same environment and confirm that they get the same conformance results. We're trying to minimize link rot (i.e., the installation directions moving or changing) and we're requesting that you include sufficient detail (particularly around flags and permissions) such that an informed user will be able to replicate your results.

## How can I track issues that have been opened regarding certification?

When a product fails certification, the issue could be in the implementation or in the conformance
tests. We use a tracking issue (TODO: link) to record issues with the
tests.

## Can I certify my private cloud that will not be available outside of our company?

You can, but it requires membership in OCI. Instead, you may be able to accomplish your goal of ensuring conformance
simply by [running](instructions.md) the conformance tests on your private cloud. As long as you pass, your
implementation is conformant. It can't be certified unless you complete the participation form, but certification
(and the ability to use the OCI Certified mark) is probably unnecessary for an internal-only product.

## I still have questions. Can you help?

Yes. Please email us at certification@opencontainers.org.

# How to review conformance results

## Technical Requirements

1. Verify that the list of files matches the
[expected list](https://github.com/opencontainers/oci-conformance/blob/main/instructions.md#contents-of-the-pr).

2. Note the `$spec/vX.Y` subdirectory that the PR is in, this is the spec and 
version of OCI for which conformance is being claimed, referenced as the
"Conformance Version" from hereon.

3. Verify that the Conformance Version is the current or previous two versions of the OCI spec.

4. Look at `report.html`.  Verify that the `major.minor` version defined
in the report summary both exactly match the Conformance Version.
The patch version does not matter.

5. Look at `report.html`. Verify that the summary contains the text "All tests passed".

## Policy Requirements

1. Confirm that the vendor is currently a
member (TODO: link) of OCI. If they are not, request
that they reach out to OCI (TODO: email) to become a member in order to
complete their certification. (Companies can alternatively pay a certfication
fee equal to the cost of membership if, for whatever reason, they don't wish to
become an OCI member.) Alternatively, non-profit organizations (including community
distributions like Debian) can certify at no cost.

Review the `PRODUCT.yaml` file, and:

2. Verify that there is a Participation Form on file for the `vendor`, and that
the vendor is in good standing in the program.

3. Verify the product `name` and `website_url` are listed in the
"Qualifying Offerings" section of the vendor's Participation Form, i.e. that
the `name` and `website_url` are listed.

4. Review the `name`. The name should not contain the term "OCI".

If the submission doesn't meet all policy requirements, reply with a message indicating "Signed participation form needed", "Files missing from PR", "Membership in OCI or confirmation of non-profit status needed", etc.

## Tasks to Complete After Review

1. Update the OCI Distributions & Platforms spreadsheet (TODO: link) to reflect the vendor's certified offering.

2. Add the vendor's information to the OCI landscape (TODO: link), which also causes them to appear on (TODO: link) and (TODO: link).

3. Add a comment saying "You are now OCI Certified" and merge the PR.

# Image Tools Interface

The below document describes an interface for the OCI Image Tools project as desired by the OCI Certification Program working group (Cert WG.)

Goals:
* An engineer can verify a container image complies with the OCI Image Spec (https://github.com/opencontainers/image-spec/releases) 

Inputs:
* (REQUIRED): Path or reference to a container image
* (OPTIONAL): If the user would like verbose output
* (OPTIONAL): Which version of the OCI Image Spec to validate against 
(IF the OCI Image Tool supports multiple OCI Image Spec versions; Default to highest supported)

Outputs:
* Number of cases run
* Number of cases passed
* Number of cases failed
* List of cases failed
* Version of OCI Image Tools
* Version of OCI Image Spec
* (IF VERBOSE) List of cases passed

## Possible Example

INPUT: 
`oci-image-tools image validate <image>`

OUTPUT (default): 
```
OCI Image Tools vX.Y.Z validating <image> with OCI Image Spec vA.B.C:
Case 12: FAILED: Image using reserved field key: https://github.com/opencontainers/image-spec/blob/master/descriptor.md#reserved 
Results: 
* Ran: 24 / 24 (100%) validation cases
* Passed: 23 / 24 (96%)
* Failed: 1 / 24 (4%)
```

OUTPUT (verbose):
```
OCI Image Tools vX.Y.Z validating <image> with OCI Image Spec vA.B.C:
Case 1: PASSED: …
Case 2: PASSED: …
…
Case 12: FAILED: Image using reserved field key: https://github.com/opencontainers/image-spec/blob/master/descriptor.md#reserved
Case 13: PASSED: …
…
Case 24: PASSED: …
Results: 
* Ran: 24 / 24 (100%) validation cases
* Passed: 23 / 24 (96%)
* Failed: 1 / 24 (4%)
```

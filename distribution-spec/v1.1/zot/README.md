# zot registry

## Running zot registry

```shell
zot-linux-amd64 serve config-conformance.json
```

with config-conformance.json as:
```json
{
  "distSpecVersion": "1.1.1",
  "storage": {
    "rootDirectory": "/tmp/zot",
    "gc": true,
    "dedupe": false
  },
  "http": {
    "address": "0.0.0.0",
    "port": "8080"
  },
  "log": {
    "level": "debug"
  }
}
```

## Run the Conformance Tests

```shell
cd distribution-spec/conformance
(
  export OCI_VERSION="1.1"
  export OCI_REGISTRY="127.0.0.1:8080"
  export OCI_TLS="disabled"
  export OCI_REPO1="oci-conformance/distribution-test"
  export OCI_REPO2="oci-conformance/crossmount-test"
  export OCI_RESULTS_DIR="."
  go run -buildvcs=true .
)
```

This should produce the following results:
```text
Test results
OCI Conformance Test: Pass
  OCI Conformance Test/ping: Pass
  OCI Conformance Test/empty: Pass
    OCI Conformance Test/empty/tag list: Pass
    OCI Conformance Test/empty/referrers: Pass
  OCI Conformance Test/sha256 blobs: Pass
    OCI Conformance Test/sha256 blobs/get-missing: Pass
    OCI Conformance Test/sha256 blobs/post only: Pass
      OCI Conformance Test/sha256 blobs/post only/blob-post-only: Pass
      OCI Conformance Test/sha256 blobs/post only/blob-head: Pass
      OCI Conformance Test/sha256 blobs/post only/blob-get: Pass
      OCI Conformance Test/sha256 blobs/post only/blob-delete: Pass
    OCI Conformance Test/sha256 blobs/post+put: Pass
      OCI Conformance Test/sha256 blobs/post+put/blob-post-put: Pass
      OCI Conformance Test/sha256 blobs/post+put/blob-head: Pass
      OCI Conformance Test/sha256 blobs/post+put/blob-get: Pass
      OCI Conformance Test/sha256 blobs/post+put/blob-delete: Pass
    OCI Conformance Test/sha256 blobs/chunked single: Pass
      OCI Conformance Test/sha256 blobs/chunked single/blob-patch-chunked: Pass
      OCI Conformance Test/sha256 blobs/chunked single/blob-head: Pass
      OCI Conformance Test/sha256 blobs/chunked single/blob-get: Pass
      OCI Conformance Test/sha256 blobs/chunked single/blob-delete: Pass
    OCI Conformance Test/sha256 blobs/stream: Pass
      OCI Conformance Test/sha256 blobs/stream/blob-patch-stream: Pass
      OCI Conformance Test/sha256 blobs/stream/blob-head: Pass
      OCI Conformance Test/sha256 blobs/stream/blob-get: Pass
      OCI Conformance Test/sha256 blobs/stream/blob-delete: Pass
    OCI Conformance Test/sha256 blobs/mount: Pass
      OCI Conformance Test/sha256 blobs/mount/blob-post-put: Pass
      OCI Conformance Test/sha256 blobs/mount/blob-mount: Skip
       - registry returned status 202, fell back to blob POST+PUT
      OCI Conformance Test/sha256 blobs/mount/blob-delete: Pass
      OCI Conformance Test/sha256 blobs/mount/blob-delete: Pass
    OCI Conformance Test/sha256 blobs/mount anonymous: Pass
      OCI Conformance Test/sha256 blobs/mount anonymous/blob-post-put: Pass
      OCI Conformance Test/sha256 blobs/mount anonymous/blob-mount-anonymous: Skip
       - registry returned status 202, fell back to blob POST+PUT
      OCI Conformance Test/sha256 blobs/mount anonymous/blob-delete: Pass
      OCI Conformance Test/sha256 blobs/mount anonymous/blob-delete: Pass
    OCI Conformance Test/sha256 blobs/mount missing: Pass
      OCI Conformance Test/sha256 blobs/mount missing/blob-mount: Pass
      OCI Conformance Test/sha256 blobs/mount missing/blob-head: Pass
      OCI Conformance Test/sha256 blobs/mount missing/blob-get: Pass
      OCI Conformance Test/sha256 blobs/mount missing/blob-delete: Pass
    OCI Conformance Test/sha256 blobs/post cancel: Pass
      OCI Conformance Test/sha256 blobs/post cancel/blob-post-cancel: Disabled
      OCI Conformance Test/sha256 blobs/post cancel/blob-delete: Pass
    OCI Conformance Test/sha256 blobs/chunked multi: Pass
      OCI Conformance Test/sha256 blobs/chunked multi/blob-patch-chunked: Pass
      OCI Conformance Test/sha256 blobs/chunked multi/blob-head: Pass
      OCI Conformance Test/sha256 blobs/chunked multi/blob-get: Pass
      OCI Conformance Test/sha256 blobs/chunked multi/blob-delete: Pass
    OCI Conformance Test/sha256 blobs/chunked multi and put chunk: Pass
      OCI Conformance Test/sha256 blobs/chunked multi and put chunk/blob-patch-chunked: Pass
      OCI Conformance Test/sha256 blobs/chunked multi and put chunk/blob-head: Pass
      OCI Conformance Test/sha256 blobs/chunked multi and put chunk/blob-get: Pass
      OCI Conformance Test/sha256 blobs/chunked multi and put chunk/blob-delete: Pass
    OCI Conformance Test/sha256 blobs/chunked out-of-order: Pass
      OCI Conformance Test/sha256 blobs/chunked out-of-order/blob-patch-chunked: Pass
      OCI Conformance Test/sha256 blobs/chunked out-of-order/blob-head: Pass
      OCI Conformance Test/sha256 blobs/chunked out-of-order/blob-get: Pass
      OCI Conformance Test/sha256 blobs/chunked out-of-order/blob-delete: Pass
    OCI Conformance Test/sha256 blobs/chunked out-of-order and put chunk: Pass
      OCI Conformance Test/sha256 blobs/chunked out-of-order and put chunk/blob-patch-chunked: Pass
      OCI Conformance Test/sha256 blobs/chunked out-of-order and put chunk/blob-head: Pass
      OCI Conformance Test/sha256 blobs/chunked out-of-order and put chunk/blob-get: Pass
      OCI Conformance Test/sha256 blobs/chunked out-of-order and put chunk/blob-delete: Pass
    OCI Conformance Test/sha256 blobs/range requests: Pass
      OCI Conformance Test/sha256 blobs/range requests/blob-post-put: Pass
      OCI Conformance Test/sha256 blobs/range requests/range 500-1499: Pass
      OCI Conformance Test/sha256 blobs/range requests/range 500-: Pass
      OCI Conformance Test/sha256 blobs/range requests/range -500: Pass
      OCI Conformance Test/sha256 blobs/range requests/range 2000-5000: Pass
      OCI Conformance Test/sha256 blobs/range requests/range 500-0: Pass
      OCI Conformance Test/sha256 blobs/range requests/range 5000-10000: Pass
      OCI Conformance Test/sha256 blobs/range requests/blob-delete: Pass
    OCI Conformance Test/sha256 blobs/empty: Pass
      OCI Conformance Test/sha256 blobs/empty/blob-post-put: Pass
      OCI Conformance Test/sha256 blobs/empty/blob-head: Pass
      OCI Conformance Test/sha256 blobs/empty/blob-get: Pass
      OCI Conformance Test/sha256 blobs/empty/blob-delete: Pass
    OCI Conformance Test/sha256 blobs/emptyJSON: Pass
      OCI Conformance Test/sha256 blobs/emptyJSON/blob-post-put: Pass
      OCI Conformance Test/sha256 blobs/emptyJSON/blob-head: Pass
      OCI Conformance Test/sha256 blobs/emptyJSON/blob-get: Pass
      OCI Conformance Test/sha256 blobs/emptyJSON/blob-delete: Pass
    OCI Conformance Test/sha256 blobs/bad digest post only: Pass
      OCI Conformance Test/sha256 blobs/bad digest post only/blob-post-only: Pass
    OCI Conformance Test/sha256 blobs/bad digest post+put: Pass
      OCI Conformance Test/sha256 blobs/bad digest post+put/blob-post-put: Pass
    OCI Conformance Test/sha256 blobs/bad digest chunked: Pass
      OCI Conformance Test/sha256 blobs/bad digest chunked/blob-patch-chunked: Pass
    OCI Conformance Test/sha256 blobs/bad digest chunked and put chunk: Pass
      OCI Conformance Test/sha256 blobs/bad digest chunked and put chunk/blob-patch-chunked: Pass
    OCI Conformance Test/sha256 blobs/bad digest stream: Pass
      OCI Conformance Test/sha256 blobs/bad digest stream/blob-patch-stream: Pass
  OCI Conformance Test/sha512 blobs: Pass
    OCI Conformance Test/sha512 blobs/get-missing: Pass
    OCI Conformance Test/sha512 blobs/post only: Pass
      OCI Conformance Test/sha512 blobs/post only/blob-post-only: Pass
      OCI Conformance Test/sha512 blobs/post only/blob-head: Pass
      OCI Conformance Test/sha512 blobs/post only/blob-get: Pass
      OCI Conformance Test/sha512 blobs/post only/blob-delete: Pass
    OCI Conformance Test/sha512 blobs/post+put: Pass
      OCI Conformance Test/sha512 blobs/post+put/blob-post-put: Pass
      OCI Conformance Test/sha512 blobs/post+put/blob-head: Pass
      OCI Conformance Test/sha512 blobs/post+put/blob-get: Pass
      OCI Conformance Test/sha512 blobs/post+put/blob-delete: Pass
    OCI Conformance Test/sha512 blobs/chunked single: Pass
      OCI Conformance Test/sha512 blobs/chunked single/blob-patch-chunked: Pass
      OCI Conformance Test/sha512 blobs/chunked single/blob-head: Pass
      OCI Conformance Test/sha512 blobs/chunked single/blob-get: Pass
      OCI Conformance Test/sha512 blobs/chunked single/blob-delete: Pass
    OCI Conformance Test/sha512 blobs/stream: Pass
      OCI Conformance Test/sha512 blobs/stream/blob-patch-stream: Pass
      OCI Conformance Test/sha512 blobs/stream/blob-head: Pass
      OCI Conformance Test/sha512 blobs/stream/blob-get: Pass
      OCI Conformance Test/sha512 blobs/stream/blob-delete: Pass
    OCI Conformance Test/sha512 blobs/mount: Pass
      OCI Conformance Test/sha512 blobs/mount/blob-post-put: Pass
      OCI Conformance Test/sha512 blobs/mount/blob-mount: Skip
       - registry returned status 202, fell back to blob POST+PUT
      OCI Conformance Test/sha512 blobs/mount/blob-delete: Pass
      OCI Conformance Test/sha512 blobs/mount/blob-delete: Pass
    OCI Conformance Test/sha512 blobs/mount anonymous: Pass
      OCI Conformance Test/sha512 blobs/mount anonymous/blob-post-put: Pass
      OCI Conformance Test/sha512 blobs/mount anonymous/blob-mount-anonymous: Skip
       - registry returned status 202, fell back to blob POST+PUT
      OCI Conformance Test/sha512 blobs/mount anonymous/blob-delete: Pass
      OCI Conformance Test/sha512 blobs/mount anonymous/blob-delete: Pass
    OCI Conformance Test/sha512 blobs/mount missing: Pass
      OCI Conformance Test/sha512 blobs/mount missing/blob-mount: Pass
      OCI Conformance Test/sha512 blobs/mount missing/blob-head: Pass
      OCI Conformance Test/sha512 blobs/mount missing/blob-get: Pass
      OCI Conformance Test/sha512 blobs/mount missing/blob-delete: Pass
    OCI Conformance Test/sha512 blobs/post cancel: Pass
      OCI Conformance Test/sha512 blobs/post cancel/blob-post-cancel: Disabled
      OCI Conformance Test/sha512 blobs/post cancel/blob-delete: Pass
    OCI Conformance Test/sha512 blobs/chunked multi: Pass
      OCI Conformance Test/sha512 blobs/chunked multi/blob-patch-chunked: Pass
      OCI Conformance Test/sha512 blobs/chunked multi/blob-head: Pass
      OCI Conformance Test/sha512 blobs/chunked multi/blob-get: Pass
      OCI Conformance Test/sha512 blobs/chunked multi/blob-delete: Pass
    OCI Conformance Test/sha512 blobs/chunked multi and put chunk: Pass
      OCI Conformance Test/sha512 blobs/chunked multi and put chunk/blob-patch-chunked: Pass
      OCI Conformance Test/sha512 blobs/chunked multi and put chunk/blob-head: Pass
      OCI Conformance Test/sha512 blobs/chunked multi and put chunk/blob-get: Pass
      OCI Conformance Test/sha512 blobs/chunked multi and put chunk/blob-delete: Pass
    OCI Conformance Test/sha512 blobs/chunked out-of-order: Pass
      OCI Conformance Test/sha512 blobs/chunked out-of-order/blob-patch-chunked: Pass
      OCI Conformance Test/sha512 blobs/chunked out-of-order/blob-head: Pass
      OCI Conformance Test/sha512 blobs/chunked out-of-order/blob-get: Pass
      OCI Conformance Test/sha512 blobs/chunked out-of-order/blob-delete: Pass
    OCI Conformance Test/sha512 blobs/chunked out-of-order and put chunk: Pass
      OCI Conformance Test/sha512 blobs/chunked out-of-order and put chunk/blob-patch-chunked: Pass
      OCI Conformance Test/sha512 blobs/chunked out-of-order and put chunk/blob-head: Pass
      OCI Conformance Test/sha512 blobs/chunked out-of-order and put chunk/blob-get: Pass
      OCI Conformance Test/sha512 blobs/chunked out-of-order and put chunk/blob-delete: Pass
    OCI Conformance Test/sha512 blobs/range requests: Pass
      OCI Conformance Test/sha512 blobs/range requests/blob-post-put: Pass
      OCI Conformance Test/sha512 blobs/range requests/range 500-1499: Pass
      OCI Conformance Test/sha512 blobs/range requests/range 500-: Pass
      OCI Conformance Test/sha512 blobs/range requests/range -500: Pass
      OCI Conformance Test/sha512 blobs/range requests/range 2000-5000: Pass
      OCI Conformance Test/sha512 blobs/range requests/range 500-0: Pass
      OCI Conformance Test/sha512 blobs/range requests/range 5000-10000: Pass
      OCI Conformance Test/sha512 blobs/range requests/blob-delete: Pass
    OCI Conformance Test/sha512 blobs/emptyJSON: Pass
      OCI Conformance Test/sha512 blobs/emptyJSON/blob-post-put: Pass
      OCI Conformance Test/sha512 blobs/emptyJSON/blob-head: Pass
      OCI Conformance Test/sha512 blobs/emptyJSON/blob-get: Pass
      OCI Conformance Test/sha512 blobs/emptyJSON/blob-delete: Pass
    OCI Conformance Test/sha512 blobs/empty: Pass
      OCI Conformance Test/sha512 blobs/empty/blob-post-put: Pass
      OCI Conformance Test/sha512 blobs/empty/blob-head: Pass
      OCI Conformance Test/sha512 blobs/empty/blob-get: Pass
      OCI Conformance Test/sha512 blobs/empty/blob-delete: Pass
    OCI Conformance Test/sha512 blobs/bad digest post only: Pass
      OCI Conformance Test/sha512 blobs/bad digest post only/blob-post-only: Pass
    OCI Conformance Test/sha512 blobs/bad digest post+put: Pass
      OCI Conformance Test/sha512 blobs/bad digest post+put/blob-post-put: Pass
    OCI Conformance Test/sha512 blobs/bad digest chunked: Pass
      OCI Conformance Test/sha512 blobs/bad digest chunked/blob-patch-chunked: Pass
    OCI Conformance Test/sha512 blobs/bad digest chunked and put chunk: Pass
      OCI Conformance Test/sha512 blobs/bad digest chunked and put chunk/blob-patch-chunked: Pass
    OCI Conformance Test/sha512 blobs/bad digest stream: Pass
      OCI Conformance Test/sha512 blobs/bad digest stream/blob-patch-stream: Pass
  OCI Conformance Test/image: Pass
    OCI Conformance Test/image/push: Pass
      OCI Conformance Test/image/push/blob-post-put: Pass
      OCI Conformance Test/image/push/blob-post-put: Pass
      OCI Conformance Test/image/push/blob-post-put: Pass
      OCI Conformance Test/image/push/manifest-by-digest: Pass
      OCI Conformance Test/image/push/manifest-by-tag: Pass
    OCI Conformance Test/image/tag-list: Pass
    OCI Conformance Test/image/head: Pass
      OCI Conformance Test/image/head/manifest-head-by-tag: Pass
      OCI Conformance Test/image/head/manifest-head-by-digest: Pass
      OCI Conformance Test/image/head/blob-head: Pass
      OCI Conformance Test/image/head/blob-head: Pass
      OCI Conformance Test/image/head/blob-head: Pass
    OCI Conformance Test/image/pull: Pass
      OCI Conformance Test/image/pull/manifest-by-tag: Pass
      OCI Conformance Test/image/pull/manifest-by-digest: Pass
      OCI Conformance Test/image/pull/blob-get: Pass
      OCI Conformance Test/image/pull/blob-get: Pass
      OCI Conformance Test/image/pull/blob-get: Pass
    OCI Conformance Test/image/delete: Pass
      OCI Conformance Test/image/delete/tag-delete: Pass
      OCI Conformance Test/image/delete/manifest-delete: Pass
      OCI Conformance Test/image/delete/blob-delete: Pass
      OCI Conformance Test/image/delete/blob-delete: Pass
      OCI Conformance Test/image/delete/blob-delete: Pass
  OCI Conformance Test/image-uncompressed: Pass
    OCI Conformance Test/image-uncompressed/push: Pass
      OCI Conformance Test/image-uncompressed/push/blob-post-put: Pass
      OCI Conformance Test/image-uncompressed/push/blob-post-put: Pass
      OCI Conformance Test/image-uncompressed/push/blob-post-put: Pass
      OCI Conformance Test/image-uncompressed/push/manifest-by-digest: Pass
      OCI Conformance Test/image-uncompressed/push/manifest-by-tag: Pass
    OCI Conformance Test/image-uncompressed/tag-list: Pass
    OCI Conformance Test/image-uncompressed/head: Pass
      OCI Conformance Test/image-uncompressed/head/manifest-head-by-tag: Pass
      OCI Conformance Test/image-uncompressed/head/manifest-head-by-digest: Pass
      OCI Conformance Test/image-uncompressed/head/blob-head: Pass
      OCI Conformance Test/image-uncompressed/head/blob-head: Pass
      OCI Conformance Test/image-uncompressed/head/blob-head: Pass
    OCI Conformance Test/image-uncompressed/pull: Pass
      OCI Conformance Test/image-uncompressed/pull/manifest-by-tag: Pass
      OCI Conformance Test/image-uncompressed/pull/manifest-by-digest: Pass
      OCI Conformance Test/image-uncompressed/pull/blob-get: Pass
      OCI Conformance Test/image-uncompressed/pull/blob-get: Pass
      OCI Conformance Test/image-uncompressed/pull/blob-get: Pass
    OCI Conformance Test/image-uncompressed/delete: Pass
      OCI Conformance Test/image-uncompressed/delete/tag-delete: Pass
      OCI Conformance Test/image-uncompressed/delete/manifest-delete: Pass
      OCI Conformance Test/image-uncompressed/delete/blob-delete: Pass
      OCI Conformance Test/image-uncompressed/delete/blob-delete: Pass
      OCI Conformance Test/image-uncompressed/delete/blob-delete: Pass
  OCI Conformance Test/large-manifest: Pass
    OCI Conformance Test/large-manifest/push: Pass
      OCI Conformance Test/large-manifest/push/blob-post-put: Pass
      OCI Conformance Test/large-manifest/push/blob-post-put: Pass
      OCI Conformance Test/large-manifest/push/blob-post-put: Pass
      OCI Conformance Test/large-manifest/push/manifest-by-digest: Pass
      OCI Conformance Test/large-manifest/push/manifest-by-tag: Pass
    OCI Conformance Test/large-manifest/tag-list: Pass
    OCI Conformance Test/large-manifest/head: Pass
      OCI Conformance Test/large-manifest/head/manifest-head-by-tag: Pass
      OCI Conformance Test/large-manifest/head/manifest-head-by-digest: Pass
      OCI Conformance Test/large-manifest/head/blob-head: Pass
      OCI Conformance Test/large-manifest/head/blob-head: Pass
      OCI Conformance Test/large-manifest/head/blob-head: Pass
    OCI Conformance Test/large-manifest/pull: Pass
      OCI Conformance Test/large-manifest/pull/manifest-by-tag: Pass
      OCI Conformance Test/large-manifest/pull/manifest-by-digest: Pass
      OCI Conformance Test/large-manifest/pull/blob-get: Pass
      OCI Conformance Test/large-manifest/pull/blob-get: Pass
      OCI Conformance Test/large-manifest/pull/blob-get: Pass
    OCI Conformance Test/large-manifest/delete: Pass
      OCI Conformance Test/large-manifest/delete/tag-delete: Pass
      OCI Conformance Test/large-manifest/delete/manifest-delete: Pass
      OCI Conformance Test/large-manifest/delete/blob-delete: Pass
      OCI Conformance Test/large-manifest/delete/blob-delete: Pass
      OCI Conformance Test/large-manifest/delete/blob-delete: Pass
  OCI Conformance Test/index: Pass
    OCI Conformance Test/index/push: Pass
      OCI Conformance Test/index/push/blob-post-put: Pass
      OCI Conformance Test/index/push/blob-post-put: Pass
      OCI Conformance Test/index/push/blob-post-put: Pass
      OCI Conformance Test/index/push/blob-post-put: Pass
      OCI Conformance Test/index/push/blob-post-put: Pass
      OCI Conformance Test/index/push/blob-post-put: Pass
      OCI Conformance Test/index/push/manifest-by-digest: Pass
      OCI Conformance Test/index/push/manifest-by-digest: Pass
      OCI Conformance Test/index/push/manifest-by-digest: Pass
      OCI Conformance Test/index/push/manifest-by-tag: Pass
    OCI Conformance Test/index/tag-list: Pass
    OCI Conformance Test/index/head: Pass
      OCI Conformance Test/index/head/manifest-head-by-tag: Pass
      OCI Conformance Test/index/head/manifest-head-by-digest: Pass
      OCI Conformance Test/index/head/manifest-head-by-digest: Pass
      OCI Conformance Test/index/head/manifest-head-by-digest: Pass
      OCI Conformance Test/index/head/blob-head: Pass
      OCI Conformance Test/index/head/blob-head: Pass
      OCI Conformance Test/index/head/blob-head: Pass
      OCI Conformance Test/index/head/blob-head: Pass
      OCI Conformance Test/index/head/blob-head: Pass
      OCI Conformance Test/index/head/blob-head: Pass
    OCI Conformance Test/index/pull: Pass
      OCI Conformance Test/index/pull/manifest-by-tag: Pass
      OCI Conformance Test/index/pull/manifest-by-digest: Pass
      OCI Conformance Test/index/pull/manifest-by-digest: Pass
      OCI Conformance Test/index/pull/manifest-by-digest: Pass
      OCI Conformance Test/index/pull/blob-get: Pass
      OCI Conformance Test/index/pull/blob-get: Pass
      OCI Conformance Test/index/pull/blob-get: Pass
      OCI Conformance Test/index/pull/blob-get: Pass
      OCI Conformance Test/index/pull/blob-get: Pass
      OCI Conformance Test/index/pull/blob-get: Pass
    OCI Conformance Test/index/delete: Pass
      OCI Conformance Test/index/delete/tag-delete: Pass
      OCI Conformance Test/index/delete/manifest-delete: Pass
      OCI Conformance Test/index/delete/manifest-delete: Pass
      OCI Conformance Test/index/delete/manifest-delete: Pass
      OCI Conformance Test/index/delete/blob-delete: Pass
      OCI Conformance Test/index/delete/blob-delete: Pass
      OCI Conformance Test/index/delete/blob-delete: Pass
      OCI Conformance Test/index/delete/blob-delete: Pass
      OCI Conformance Test/index/delete/blob-delete: Pass
      OCI Conformance Test/index/delete/blob-delete: Pass
  OCI Conformance Test/nested-index: Pass
    OCI Conformance Test/nested-index/push: Pass
      OCI Conformance Test/nested-index/push/blob-post-put: Pass
      OCI Conformance Test/nested-index/push/blob-post-put: Pass
      OCI Conformance Test/nested-index/push/blob-post-put: Pass
      OCI Conformance Test/nested-index/push/blob-post-put: Pass
      OCI Conformance Test/nested-index/push/blob-post-put: Pass
      OCI Conformance Test/nested-index/push/blob-post-put: Pass
      OCI Conformance Test/nested-index/push/blob-post-put: Pass
      OCI Conformance Test/nested-index/push/blob-post-put: Pass
      OCI Conformance Test/nested-index/push/blob-post-put: Pass
      OCI Conformance Test/nested-index/push/blob-post-put: Pass
      OCI Conformance Test/nested-index/push/manifest-by-digest: Pass
      OCI Conformance Test/nested-index/push/manifest-by-digest: Pass
      OCI Conformance Test/nested-index/push/manifest-by-digest: Pass
      OCI Conformance Test/nested-index/push/manifest-by-digest: Pass
      OCI Conformance Test/nested-index/push/manifest-by-digest: Pass
      OCI Conformance Test/nested-index/push/manifest-by-digest: Pass
      OCI Conformance Test/nested-index/push/manifest-by-digest: Pass
      OCI Conformance Test/nested-index/push/manifest-by-tag: Pass
    OCI Conformance Test/nested-index/tag-list: Pass
    OCI Conformance Test/nested-index/head: Pass
      OCI Conformance Test/nested-index/head/manifest-head-by-tag: Pass
      OCI Conformance Test/nested-index/head/manifest-head-by-digest: Pass
      OCI Conformance Test/nested-index/head/manifest-head-by-digest: Pass
      OCI Conformance Test/nested-index/head/manifest-head-by-digest: Pass
      OCI Conformance Test/nested-index/head/manifest-head-by-digest: Pass
      OCI Conformance Test/nested-index/head/manifest-head-by-digest: Pass
      OCI Conformance Test/nested-index/head/manifest-head-by-digest: Pass
      OCI Conformance Test/nested-index/head/manifest-head-by-digest: Pass
      OCI Conformance Test/nested-index/head/blob-head: Pass
      OCI Conformance Test/nested-index/head/blob-head: Pass
      OCI Conformance Test/nested-index/head/blob-head: Pass
      OCI Conformance Test/nested-index/head/blob-head: Pass
      OCI Conformance Test/nested-index/head/blob-head: Pass
      OCI Conformance Test/nested-index/head/blob-head: Pass
      OCI Conformance Test/nested-index/head/blob-head: Pass
      OCI Conformance Test/nested-index/head/blob-head: Pass
      OCI Conformance Test/nested-index/head/blob-head: Pass
      OCI Conformance Test/nested-index/head/blob-head: Pass
    OCI Conformance Test/nested-index/pull: Pass
      OCI Conformance Test/nested-index/pull/manifest-by-tag: Pass
      OCI Conformance Test/nested-index/pull/manifest-by-digest: Pass
      OCI Conformance Test/nested-index/pull/manifest-by-digest: Pass
      OCI Conformance Test/nested-index/pull/manifest-by-digest: Pass
      OCI Conformance Test/nested-index/pull/manifest-by-digest: Pass
      OCI Conformance Test/nested-index/pull/manifest-by-digest: Pass
      OCI Conformance Test/nested-index/pull/manifest-by-digest: Pass
      OCI Conformance Test/nested-index/pull/manifest-by-digest: Pass
      OCI Conformance Test/nested-index/pull/blob-get: Pass
      OCI Conformance Test/nested-index/pull/blob-get: Pass
      OCI Conformance Test/nested-index/pull/blob-get: Pass
      OCI Conformance Test/nested-index/pull/blob-get: Pass
      OCI Conformance Test/nested-index/pull/blob-get: Pass
      OCI Conformance Test/nested-index/pull/blob-get: Pass
      OCI Conformance Test/nested-index/pull/blob-get: Pass
      OCI Conformance Test/nested-index/pull/blob-get: Pass
      OCI Conformance Test/nested-index/pull/blob-get: Pass
      OCI Conformance Test/nested-index/pull/blob-get: Pass
    OCI Conformance Test/nested-index/delete: Pass
      OCI Conformance Test/nested-index/delete/tag-delete: Pass
      OCI Conformance Test/nested-index/delete/manifest-delete: Pass
      OCI Conformance Test/nested-index/delete/manifest-delete: Pass
      OCI Conformance Test/nested-index/delete/manifest-delete: Pass
      OCI Conformance Test/nested-index/delete/manifest-delete: Pass
      OCI Conformance Test/nested-index/delete/manifest-delete: Pass
      OCI Conformance Test/nested-index/delete/manifest-delete: Pass
      OCI Conformance Test/nested-index/delete/manifest-delete: Pass
      OCI Conformance Test/nested-index/delete/blob-delete: Pass
      OCI Conformance Test/nested-index/delete/blob-delete: Pass
      OCI Conformance Test/nested-index/delete/blob-delete: Pass
      OCI Conformance Test/nested-index/delete/blob-delete: Pass
      OCI Conformance Test/nested-index/delete/blob-delete: Pass
      OCI Conformance Test/nested-index/delete/blob-delete: Pass
      OCI Conformance Test/nested-index/delete/blob-delete: Pass
      OCI Conformance Test/nested-index/delete/blob-delete: Pass
      OCI Conformance Test/nested-index/delete/blob-delete: Pass
      OCI Conformance Test/nested-index/delete/blob-delete: Pass
  OCI Conformance Test/empty-index: Pass
    OCI Conformance Test/empty-index/push: Pass
      OCI Conformance Test/empty-index/push/manifest-by-digest: Pass
      OCI Conformance Test/empty-index/push/manifest-by-tag: Pass
    OCI Conformance Test/empty-index/tag-list: Pass
    OCI Conformance Test/empty-index/head: Pass
      OCI Conformance Test/empty-index/head/manifest-head-by-tag: Pass
      OCI Conformance Test/empty-index/head/manifest-head-by-digest: Pass
    OCI Conformance Test/empty-index/pull: Pass
      OCI Conformance Test/empty-index/pull/manifest-by-tag: Pass
      OCI Conformance Test/empty-index/pull/manifest-by-digest: Pass
    OCI Conformance Test/empty-index/delete: Pass
      OCI Conformance Test/empty-index/delete/tag-delete: Pass
      OCI Conformance Test/empty-index/delete/manifest-delete: Pass
  OCI Conformance Test/artifact: Pass
    OCI Conformance Test/artifact/push: Pass
      OCI Conformance Test/artifact/push/blob-post-put: Pass
      OCI Conformance Test/artifact/push/blob-post-put: Pass
      OCI Conformance Test/artifact/push/manifest-by-digest: Pass
      OCI Conformance Test/artifact/push/manifest-by-tag: Pass
    OCI Conformance Test/artifact/tag-list: Pass
    OCI Conformance Test/artifact/head: Pass
      OCI Conformance Test/artifact/head/manifest-head-by-tag: Pass
      OCI Conformance Test/artifact/head/manifest-head-by-digest: Pass
      OCI Conformance Test/artifact/head/blob-head: Pass
      OCI Conformance Test/artifact/head/blob-head: Pass
    OCI Conformance Test/artifact/pull: Pass
      OCI Conformance Test/artifact/pull/manifest-by-tag: Pass
      OCI Conformance Test/artifact/pull/manifest-by-digest: Pass
      OCI Conformance Test/artifact/pull/blob-get: Pass
      OCI Conformance Test/artifact/pull/blob-get: Pass
    OCI Conformance Test/artifact/delete: Pass
      OCI Conformance Test/artifact/delete/tag-delete: Pass
      OCI Conformance Test/artifact/delete/manifest-delete: Pass
      OCI Conformance Test/artifact/delete/blob-delete: Pass
      OCI Conformance Test/artifact/delete/blob-delete: Pass
  OCI Conformance Test/artifact-index: Pass
    OCI Conformance Test/artifact-index/push: Pass
      OCI Conformance Test/artifact-index/push/blob-post-put: Pass
      OCI Conformance Test/artifact-index/push/blob-post-put: Pass
      OCI Conformance Test/artifact-index/push/blob-post-put: Pass
      OCI Conformance Test/artifact-index/push/manifest-by-digest: Pass
      OCI Conformance Test/artifact-index/push/manifest-by-digest: Pass
      OCI Conformance Test/artifact-index/push/manifest-by-digest: Pass
      OCI Conformance Test/artifact-index/push/manifest-by-tag: Pass
    OCI Conformance Test/artifact-index/tag-list: Pass
    OCI Conformance Test/artifact-index/head: Pass
      OCI Conformance Test/artifact-index/head/manifest-head-by-tag: Pass
      OCI Conformance Test/artifact-index/head/manifest-head-by-digest: Pass
      OCI Conformance Test/artifact-index/head/manifest-head-by-digest: Pass
      OCI Conformance Test/artifact-index/head/manifest-head-by-digest: Pass
      OCI Conformance Test/artifact-index/head/blob-head: Pass
      OCI Conformance Test/artifact-index/head/blob-head: Pass
      OCI Conformance Test/artifact-index/head/blob-head: Pass
    OCI Conformance Test/artifact-index/pull: Pass
      OCI Conformance Test/artifact-index/pull/manifest-by-tag: Pass
      OCI Conformance Test/artifact-index/pull/manifest-by-digest: Pass
      OCI Conformance Test/artifact-index/pull/manifest-by-digest: Pass
      OCI Conformance Test/artifact-index/pull/manifest-by-digest: Pass
      OCI Conformance Test/artifact-index/pull/blob-get: Pass
      OCI Conformance Test/artifact-index/pull/blob-get: Pass
      OCI Conformance Test/artifact-index/pull/blob-get: Pass
    OCI Conformance Test/artifact-index/delete: Pass
      OCI Conformance Test/artifact-index/delete/tag-delete: Pass
      OCI Conformance Test/artifact-index/delete/manifest-delete: Pass
      OCI Conformance Test/artifact-index/delete/manifest-delete: Pass
      OCI Conformance Test/artifact-index/delete/manifest-delete: Pass
      OCI Conformance Test/artifact-index/delete/blob-delete: Pass
      OCI Conformance Test/artifact-index/delete/blob-delete: Pass
      OCI Conformance Test/artifact-index/delete/blob-delete: Pass
  OCI Conformance Test/artifact-without-layers: Pass
    OCI Conformance Test/artifact-without-layers/push: Pass
      OCI Conformance Test/artifact-without-layers/push/blob-post-put: Pass
      OCI Conformance Test/artifact-without-layers/push/manifest-by-digest: Pass
      OCI Conformance Test/artifact-without-layers/push/manifest-by-tag: Pass
    OCI Conformance Test/artifact-without-layers/tag-list: Pass
    OCI Conformance Test/artifact-without-layers/head: Pass
      OCI Conformance Test/artifact-without-layers/head/manifest-head-by-tag: Pass
      OCI Conformance Test/artifact-without-layers/head/manifest-head-by-digest: Pass
      OCI Conformance Test/artifact-without-layers/head/blob-head: Pass
    OCI Conformance Test/artifact-without-layers/pull: Pass
      OCI Conformance Test/artifact-without-layers/pull/manifest-by-tag: Pass
      OCI Conformance Test/artifact-without-layers/pull/manifest-by-digest: Pass
      OCI Conformance Test/artifact-without-layers/pull/blob-get: Pass
    OCI Conformance Test/artifact-without-layers/delete: Pass
      OCI Conformance Test/artifact-without-layers/delete/tag-delete: Pass
      OCI Conformance Test/artifact-without-layers/delete/manifest-delete: Pass
      OCI Conformance Test/artifact-without-layers/delete/blob-delete: Pass
  OCI Conformance Test/artifacts-with-subject: Pass
    OCI Conformance Test/artifacts-with-subject/push: Pass
      OCI Conformance Test/artifacts-with-subject/push/blob-post-put: Pass
      OCI Conformance Test/artifacts-with-subject/push/blob-post-put: Pass
      OCI Conformance Test/artifacts-with-subject/push/blob-post-put: Pass
      OCI Conformance Test/artifacts-with-subject/push/blob-post-put: Pass
      OCI Conformance Test/artifacts-with-subject/push/blob-post-put: Pass
      OCI Conformance Test/artifacts-with-subject/push/blob-post-put: Pass
      OCI Conformance Test/artifacts-with-subject/push/blob-post-put: Pass
      OCI Conformance Test/artifacts-with-subject/push/blob-post-put: Pass
      OCI Conformance Test/artifacts-with-subject/push/manifest-by-digest: Pass
      OCI Conformance Test/artifacts-with-subject/push/manifest-by-digest: Pass
      OCI Conformance Test/artifacts-with-subject/push/manifest-by-digest: Pass
      OCI Conformance Test/artifacts-with-subject/push/manifest-by-digest: Pass
      OCI Conformance Test/artifacts-with-subject/push/manifest-by-tag: Pass
      OCI Conformance Test/artifacts-with-subject/push/manifest-by-tag: Pass
      OCI Conformance Test/artifacts-with-subject/push/manifest-by-tag: Pass
    OCI Conformance Test/artifacts-with-subject/tag-list: Pass
    OCI Conformance Test/artifacts-with-subject/head: Pass
      OCI Conformance Test/artifacts-with-subject/head/manifest-head-by-tag: Pass
      OCI Conformance Test/artifacts-with-subject/head/manifest-head-by-tag: Pass
      OCI Conformance Test/artifacts-with-subject/head/manifest-head-by-tag: Pass
      OCI Conformance Test/artifacts-with-subject/head/manifest-head-by-digest: Pass
      OCI Conformance Test/artifacts-with-subject/head/manifest-head-by-digest: Pass
      OCI Conformance Test/artifacts-with-subject/head/manifest-head-by-digest: Pass
      OCI Conformance Test/artifacts-with-subject/head/manifest-head-by-digest: Pass
      OCI Conformance Test/artifacts-with-subject/head/blob-head: Pass
      OCI Conformance Test/artifacts-with-subject/head/blob-head: Pass
      OCI Conformance Test/artifacts-with-subject/head/blob-head: Pass
      OCI Conformance Test/artifacts-with-subject/head/blob-head: Pass
      OCI Conformance Test/artifacts-with-subject/head/blob-head: Pass
      OCI Conformance Test/artifacts-with-subject/head/blob-head: Pass
      OCI Conformance Test/artifacts-with-subject/head/blob-head: Pass
      OCI Conformance Test/artifacts-with-subject/head/blob-head: Pass
    OCI Conformance Test/artifacts-with-subject/pull: Pass
      OCI Conformance Test/artifacts-with-subject/pull/manifest-by-tag: Pass
      OCI Conformance Test/artifacts-with-subject/pull/manifest-by-tag: Pass
      OCI Conformance Test/artifacts-with-subject/pull/manifest-by-tag: Pass
      OCI Conformance Test/artifacts-with-subject/pull/manifest-by-digest: Pass
      OCI Conformance Test/artifacts-with-subject/pull/manifest-by-digest: Pass
      OCI Conformance Test/artifacts-with-subject/pull/manifest-by-digest: Pass
      OCI Conformance Test/artifacts-with-subject/pull/manifest-by-digest: Pass
      OCI Conformance Test/artifacts-with-subject/pull/blob-get: Pass
      OCI Conformance Test/artifacts-with-subject/pull/blob-get: Pass
+ conformance_rc=0
+ set -e
      OCI Conformance Test/artifacts-with-subject/pull/blob-get: Pass
+ [[ -f report.html ]]
      OCI Conformance Test/artifacts-with-subject/pull/blob-get: Pass
+ echo 'Found report.html.'
+ echo has-report=true
+ echo 'Conformance return code: 0'
+ exit 0
      OCI Conformance Test/artifacts-with-subject/pull/blob-get: Pass
      OCI Conformance Test/artifacts-with-subject/pull/blob-get: Pass
      OCI Conformance Test/artifacts-with-subject/pull/blob-get: Pass
      OCI Conformance Test/artifacts-with-subject/pull/blob-get: Pass
    OCI Conformance Test/artifacts-with-subject/referrers: Pass
    OCI Conformance Test/artifacts-with-subject/delete: Pass
      OCI Conformance Test/artifacts-with-subject/delete/tag-delete: Pass
      OCI Conformance Test/artifacts-with-subject/delete/tag-delete: Pass
      OCI Conformance Test/artifacts-with-subject/delete/tag-delete: Pass
      OCI Conformance Test/artifacts-with-subject/delete/manifest-delete: Pass
      OCI Conformance Test/artifacts-with-subject/delete/manifest-delete: Pass
      OCI Conformance Test/artifacts-with-subject/delete/manifest-delete: Pass
      OCI Conformance Test/artifacts-with-subject/delete/manifest-delete: Pass
      OCI Conformance Test/artifacts-with-subject/delete/blob-delete: Pass
      OCI Conformance Test/artifacts-with-subject/delete/blob-delete: Pass
      OCI Conformance Test/artifacts-with-subject/delete/blob-delete: Pass
      OCI Conformance Test/artifacts-with-subject/delete/blob-delete: Pass
      OCI Conformance Test/artifacts-with-subject/delete/blob-delete: Pass
      OCI Conformance Test/artifacts-with-subject/delete/blob-delete: Pass
      OCI Conformance Test/artifacts-with-subject/delete/blob-delete: Pass
      OCI Conformance Test/artifacts-with-subject/delete/blob-delete: Pass
  OCI Conformance Test/index-with-subject: Pass
    OCI Conformance Test/index-with-subject/push: Pass
      OCI Conformance Test/index-with-subject/push/blob-post-put: Pass
      OCI Conformance Test/index-with-subject/push/blob-post-put: Pass
      OCI Conformance Test/index-with-subject/push/blob-post-put: Pass
      OCI Conformance Test/index-with-subject/push/blob-post-put: Pass
      OCI Conformance Test/index-with-subject/push/blob-post-put: Pass
      OCI Conformance Test/index-with-subject/push/blob-post-put: Pass
      OCI Conformance Test/index-with-subject/push/blob-post-put: Pass
      OCI Conformance Test/index-with-subject/push/blob-post-put: Pass
      OCI Conformance Test/index-with-subject/push/blob-post-put: Pass
      OCI Conformance Test/index-with-subject/push/manifest-by-digest: Pass
      OCI Conformance Test/index-with-subject/push/manifest-by-digest: Pass
      OCI Conformance Test/index-with-subject/push/manifest-by-digest: Pass
      OCI Conformance Test/index-with-subject/push/manifest-by-digest: Pass
      OCI Conformance Test/index-with-subject/push/manifest-by-digest: Pass
      OCI Conformance Test/index-with-subject/push/manifest-by-digest: Pass
      OCI Conformance Test/index-with-subject/push/manifest-by-tag: Pass
      OCI Conformance Test/index-with-subject/push/manifest-by-tag: Pass
    OCI Conformance Test/index-with-subject/tag-list: Pass
    OCI Conformance Test/index-with-subject/head: Pass
      OCI Conformance Test/index-with-subject/head/manifest-head-by-tag: Pass
      OCI Conformance Test/index-with-subject/head/manifest-head-by-tag: Pass
      OCI Conformance Test/index-with-subject/head/manifest-head-by-digest: Pass
      OCI Conformance Test/index-with-subject/head/manifest-head-by-digest: Pass
      OCI Conformance Test/index-with-subject/head/manifest-head-by-digest: Pass
      OCI Conformance Test/index-with-subject/head/manifest-head-by-digest: Pass
      OCI Conformance Test/index-with-subject/head/manifest-head-by-digest: Pass
      OCI Conformance Test/index-with-subject/head/manifest-head-by-digest: Pass
      OCI Conformance Test/index-with-subject/head/blob-head: Pass
      OCI Conformance Test/index-with-subject/head/blob-head: Pass
      OCI Conformance Test/index-with-subject/head/blob-head: Pass
      OCI Conformance Test/index-with-subject/head/blob-head: Pass
      OCI Conformance Test/index-with-subject/head/blob-head: Pass
      OCI Conformance Test/index-with-subject/head/blob-head: Pass
      OCI Conformance Test/index-with-subject/head/blob-head: Pass
      OCI Conformance Test/index-with-subject/head/blob-head: Pass
      OCI Conformance Test/index-with-subject/head/blob-head: Pass
    OCI Conformance Test/index-with-subject/pull: Pass
      OCI Conformance Test/index-with-subject/pull/manifest-by-tag: Pass
      OCI Conformance Test/index-with-subject/pull/manifest-by-tag: Pass
      OCI Conformance Test/index-with-subject/pull/manifest-by-digest: Pass
      OCI Conformance Test/index-with-subject/pull/manifest-by-digest: Pass
      OCI Conformance Test/index-with-subject/pull/manifest-by-digest: Pass
      OCI Conformance Test/index-with-subject/pull/manifest-by-digest: Pass
      OCI Conformance Test/index-with-subject/pull/manifest-by-digest: Pass
      OCI Conformance Test/index-with-subject/pull/manifest-by-digest: Pass
      OCI Conformance Test/index-with-subject/pull/blob-get: Pass
      OCI Conformance Test/index-with-subject/pull/blob-get: Pass
      OCI Conformance Test/index-with-subject/pull/blob-get: Pass
      OCI Conformance Test/index-with-subject/pull/blob-get: Pass
      OCI Conformance Test/index-with-subject/pull/blob-get: Pass
      OCI Conformance Test/index-with-subject/pull/blob-get: Pass
      OCI Conformance Test/index-with-subject/pull/blob-get: Pass
      OCI Conformance Test/index-with-subject/pull/blob-get: Pass
      OCI Conformance Test/index-with-subject/pull/blob-get: Pass
    OCI Conformance Test/index-with-subject/referrers: Pass
    OCI Conformance Test/index-with-subject/delete: Pass
      OCI Conformance Test/index-with-subject/delete/tag-delete: Pass
      OCI Conformance Test/index-with-subject/delete/tag-delete: Pass
      OCI Conformance Test/index-with-subject/delete/manifest-delete: Pass
      OCI Conformance Test/index-with-subject/delete/manifest-delete: Pass
      OCI Conformance Test/index-with-subject/delete/manifest-delete: Pass
      OCI Conformance Test/index-with-subject/delete/manifest-delete: Pass
      OCI Conformance Test/index-with-subject/delete/manifest-delete: Pass
      OCI Conformance Test/index-with-subject/delete/manifest-delete: Pass
      OCI Conformance Test/index-with-subject/delete/blob-delete: Pass
      OCI Conformance Test/index-with-subject/delete/blob-delete: Pass
      OCI Conformance Test/index-with-subject/delete/blob-delete: Pass
      OCI Conformance Test/index-with-subject/delete/blob-delete: Pass
      OCI Conformance Test/index-with-subject/delete/blob-delete: Pass
      OCI Conformance Test/index-with-subject/delete/blob-delete: Pass
      OCI Conformance Test/index-with-subject/delete/blob-delete: Pass
      OCI Conformance Test/index-with-subject/delete/blob-delete: Pass
      OCI Conformance Test/index-with-subject/delete/blob-delete: Pass
  OCI Conformance Test/missing-subject: Pass
    OCI Conformance Test/missing-subject/push: Pass
      OCI Conformance Test/missing-subject/push/blob-post-put: Pass
      OCI Conformance Test/missing-subject/push/blob-post-put: Pass
      OCI Conformance Test/missing-subject/push/manifest-by-digest: Pass
      OCI Conformance Test/missing-subject/push/manifest-by-tag: Pass
    OCI Conformance Test/missing-subject/tag-list: Pass
    OCI Conformance Test/missing-subject/head: Pass
      OCI Conformance Test/missing-subject/head/manifest-head-by-tag: Pass
      OCI Conformance Test/missing-subject/head/manifest-head-by-digest: Pass
      OCI Conformance Test/missing-subject/head/blob-head: Pass
      OCI Conformance Test/missing-subject/head/blob-head: Pass
    OCI Conformance Test/missing-subject/pull: Pass
      OCI Conformance Test/missing-subject/pull/manifest-by-tag: Pass
      OCI Conformance Test/missing-subject/pull/manifest-by-digest: Pass
      OCI Conformance Test/missing-subject/pull/blob-get: Pass
      OCI Conformance Test/missing-subject/pull/blob-get: Pass
    OCI Conformance Test/missing-subject/referrers: Pass
    OCI Conformance Test/missing-subject/delete: Pass
      OCI Conformance Test/missing-subject/delete/tag-delete: Pass
      OCI Conformance Test/missing-subject/delete/manifest-delete: Pass
      OCI Conformance Test/missing-subject/delete/blob-delete: Pass
      OCI Conformance Test/missing-subject/delete/blob-delete: Pass
  OCI Conformance Test/data-field: Pass
    OCI Conformance Test/data-field/push: Pass
      OCI Conformance Test/data-field/push/blob-post-put: Pass
      OCI Conformance Test/data-field/push/blob-post-put: Pass
      OCI Conformance Test/data-field/push/blob-post-put: Pass
      OCI Conformance Test/data-field/push/manifest-by-digest: Pass
      OCI Conformance Test/data-field/push/manifest-by-tag: Pass
    OCI Conformance Test/data-field/tag-list: Pass
    OCI Conformance Test/data-field/head: Pass
      OCI Conformance Test/data-field/head/manifest-head-by-tag: Pass
      OCI Conformance Test/data-field/head/manifest-head-by-digest: Pass
      OCI Conformance Test/data-field/head/blob-head: Pass
      OCI Conformance Test/data-field/head/blob-head: Pass
      OCI Conformance Test/data-field/head/blob-head: Pass
    OCI Conformance Test/data-field/pull: Pass
      OCI Conformance Test/data-field/pull/manifest-by-tag: Pass
      OCI Conformance Test/data-field/pull/manifest-by-digest: Pass
      OCI Conformance Test/data-field/pull/blob-get: Pass
      OCI Conformance Test/data-field/pull/blob-get: Pass
      OCI Conformance Test/data-field/pull/blob-get: Pass
    OCI Conformance Test/data-field/delete: Pass
      OCI Conformance Test/data-field/delete/tag-delete: Pass
      OCI Conformance Test/data-field/delete/manifest-delete: Pass
      OCI Conformance Test/data-field/delete/blob-delete: Pass
      OCI Conformance Test/data-field/delete/blob-delete: Pass
      OCI Conformance Test/data-field/delete/blob-delete: Pass
  OCI Conformance Test/non-distributable-layers: Pass
    OCI Conformance Test/non-distributable-layers/push: Pass
      OCI Conformance Test/non-distributable-layers/push/blob-post-put: Pass
      OCI Conformance Test/non-distributable-layers/push/blob-post-put: Pass
      OCI Conformance Test/non-distributable-layers/push/manifest-by-digest: Pass
      OCI Conformance Test/non-distributable-layers/push/manifest-by-tag: Pass
    OCI Conformance Test/non-distributable-layers/tag-list: Pass
    OCI Conformance Test/non-distributable-layers/head: Pass
      OCI Conformance Test/non-distributable-layers/head/manifest-head-by-tag: Pass
      OCI Conformance Test/non-distributable-layers/head/manifest-head-by-digest: Pass
      OCI Conformance Test/non-distributable-layers/head/blob-head: Pass
      OCI Conformance Test/non-distributable-layers/head/blob-head: Pass
    OCI Conformance Test/non-distributable-layers/pull: Pass
      OCI Conformance Test/non-distributable-layers/pull/manifest-by-tag: Pass
      OCI Conformance Test/non-distributable-layers/pull/manifest-by-digest: Pass
      OCI Conformance Test/non-distributable-layers/pull/blob-get: Pass
      OCI Conformance Test/non-distributable-layers/pull/blob-get: Pass
    OCI Conformance Test/non-distributable-layers/delete: Pass
      OCI Conformance Test/non-distributable-layers/delete/tag-delete: Pass
      OCI Conformance Test/non-distributable-layers/delete/manifest-delete: Pass
      OCI Conformance Test/non-distributable-layers/delete/blob-delete: Pass
      OCI Conformance Test/non-distributable-layers/delete/blob-delete: Pass
  OCI Conformance Test/custom-fields: Pass
    OCI Conformance Test/custom-fields/push: Pass
      OCI Conformance Test/custom-fields/push/blob-post-put: Pass
      OCI Conformance Test/custom-fields/push/blob-post-put: Pass
      OCI Conformance Test/custom-fields/push/blob-post-put: Pass
      OCI Conformance Test/custom-fields/push/blob-post-put: Pass
      OCI Conformance Test/custom-fields/push/blob-post-put: Pass
      OCI Conformance Test/custom-fields/push/blob-post-put: Pass
      OCI Conformance Test/custom-fields/push/manifest-by-digest: Pass
      OCI Conformance Test/custom-fields/push/manifest-by-digest: Pass
      OCI Conformance Test/custom-fields/push/manifest-by-digest: Pass
      OCI Conformance Test/custom-fields/push/manifest-by-tag: Pass
    OCI Conformance Test/custom-fields/tag-list: Pass
    OCI Conformance Test/custom-fields/head: Pass
      OCI Conformance Test/custom-fields/head/manifest-head-by-tag: Pass
      OCI Conformance Test/custom-fields/head/manifest-head-by-digest: Pass
      OCI Conformance Test/custom-fields/head/manifest-head-by-digest: Pass
      OCI Conformance Test/custom-fields/head/manifest-head-by-digest: Pass
      OCI Conformance Test/custom-fields/head/blob-head: Pass
      OCI Conformance Test/custom-fields/head/blob-head: Pass
      OCI Conformance Test/custom-fields/head/blob-head: Pass
      OCI Conformance Test/custom-fields/head/blob-head: Pass
      OCI Conformance Test/custom-fields/head/blob-head: Pass
      OCI Conformance Test/custom-fields/head/blob-head: Pass
    OCI Conformance Test/custom-fields/pull: Pass
      OCI Conformance Test/custom-fields/pull/manifest-by-tag: Pass
      OCI Conformance Test/custom-fields/pull/manifest-by-digest: Pass
      OCI Conformance Test/custom-fields/pull/manifest-by-digest: Pass
      OCI Conformance Test/custom-fields/pull/manifest-by-digest: Pass
      OCI Conformance Test/custom-fields/pull/blob-get: Pass
      OCI Conformance Test/custom-fields/pull/blob-get: Pass
      OCI Conformance Test/custom-fields/pull/blob-get: Pass
      OCI Conformance Test/custom-fields/pull/blob-get: Pass
      OCI Conformance Test/custom-fields/pull/blob-get: Pass
      OCI Conformance Test/custom-fields/pull/blob-get: Pass
    OCI Conformance Test/custom-fields/delete: Pass
      OCI Conformance Test/custom-fields/delete/tag-delete: Pass
      OCI Conformance Test/custom-fields/delete/manifest-delete: Pass
      OCI Conformance Test/custom-fields/delete/manifest-delete: Pass
      OCI Conformance Test/custom-fields/delete/manifest-delete: Pass
      OCI Conformance Test/custom-fields/delete/blob-delete: Pass
      OCI Conformance Test/custom-fields/delete/blob-delete: Pass
      OCI Conformance Test/custom-fields/delete/blob-delete: Pass
      OCI Conformance Test/custom-fields/delete/blob-delete: Pass
      OCI Conformance Test/custom-fields/delete/blob-delete: Pass
      OCI Conformance Test/custom-fields/delete/blob-delete: Pass
  OCI Conformance Test/no-layers: Pass
    OCI Conformance Test/no-layers/push: Pass
      OCI Conformance Test/no-layers/push/blob-post-put: Pass
      OCI Conformance Test/no-layers/push/manifest-by-digest: Pass
      OCI Conformance Test/no-layers/push/manifest-by-tag: Pass
    OCI Conformance Test/no-layers/tag-list: Pass
    OCI Conformance Test/no-layers/head: Pass
      OCI Conformance Test/no-layers/head/manifest-head-by-tag: Pass
      OCI Conformance Test/no-layers/head/manifest-head-by-digest: Pass
      OCI Conformance Test/no-layers/head/blob-head: Pass
    OCI Conformance Test/no-layers/pull: Pass
      OCI Conformance Test/no-layers/pull/manifest-by-tag: Pass
      OCI Conformance Test/no-layers/pull/manifest-by-digest: Pass
      OCI Conformance Test/no-layers/pull/blob-get: Pass
    OCI Conformance Test/no-layers/delete: Pass
      OCI Conformance Test/no-layers/delete/tag-delete: Pass
      OCI Conformance Test/no-layers/delete/manifest-delete: Pass
      OCI Conformance Test/no-layers/delete/blob-delete: Pass
  OCI Conformance Test/sha512: Pass
    OCI Conformance Test/sha512/push: Pass
      OCI Conformance Test/sha512/push/blob-post-put: Pass
      OCI Conformance Test/sha512/push/blob-post-put: Pass
      OCI Conformance Test/sha512/push/blob-post-put: Pass
      OCI Conformance Test/sha512/push/blob-post-put: Pass
      OCI Conformance Test/sha512/push/blob-post-put: Pass
      OCI Conformance Test/sha512/push/blob-post-put: Pass
      OCI Conformance Test/sha512/push/manifest-by-digest: Pass
      OCI Conformance Test/sha512/push/manifest-by-digest: Pass
      OCI Conformance Test/sha512/push/manifest-by-digest: Pass
    OCI Conformance Test/sha512/tag-list: Pass
    OCI Conformance Test/sha512/head: Pass
      OCI Conformance Test/sha512/head/manifest-head-by-digest: Pass
      OCI Conformance Test/sha512/head/manifest-head-by-digest: Pass
      OCI Conformance Test/sha512/head/manifest-head-by-digest: Pass
      OCI Conformance Test/sha512/head/blob-head: Pass
      OCI Conformance Test/sha512/head/blob-head: Pass
      OCI Conformance Test/sha512/head/blob-head: Pass
      OCI Conformance Test/sha512/head/blob-head: Pass
      OCI Conformance Test/sha512/head/blob-head: Pass
      OCI Conformance Test/sha512/head/blob-head: Pass
    OCI Conformance Test/sha512/pull: Pass
      OCI Conformance Test/sha512/pull/manifest-by-digest: Pass
      OCI Conformance Test/sha512/pull/manifest-by-digest: Pass
      OCI Conformance Test/sha512/pull/manifest-by-digest: Pass
      OCI Conformance Test/sha512/pull/blob-get: Pass
      OCI Conformance Test/sha512/pull/blob-get: Pass
      OCI Conformance Test/sha512/pull/blob-get: Pass
      OCI Conformance Test/sha512/pull/blob-get: Pass
      OCI Conformance Test/sha512/pull/blob-get: Pass
      OCI Conformance Test/sha512/pull/blob-get: Pass
    OCI Conformance Test/sha512/delete: Pass
      OCI Conformance Test/sha512/delete/manifest-delete: Pass
      OCI Conformance Test/sha512/delete/manifest-delete: Pass
      OCI Conformance Test/sha512/delete/manifest-delete: Pass
      OCI Conformance Test/sha512/delete/blob-delete: Pass
      OCI Conformance Test/sha512/delete/blob-delete: Pass
      OCI Conformance Test/sha512/delete/blob-delete: Pass
      OCI Conformance Test/sha512/delete/blob-delete: Pass
      OCI Conformance Test/sha512/delete/blob-delete: Pass
      OCI Conformance Test/sha512/delete/blob-delete: Pass
  OCI Conformance Test/bad-digest-image: Pass
    OCI Conformance Test/bad-digest-image/push: Pass
      OCI Conformance Test/bad-digest-image/push/blob-post-put: Pass
      OCI Conformance Test/bad-digest-image/push/blob-post-put: Pass
      OCI Conformance Test/bad-digest-image/push/blob-post-put: Pass
      OCI Conformance Test/bad-digest-image/push/manifest-by-digest: Pass
    OCI Conformance Test/bad-digest-image/tag-list: Pass
    OCI Conformance Test/bad-digest-image/head: Pass
      OCI Conformance Test/bad-digest-image/head/blob-head: Pass
      OCI Conformance Test/bad-digest-image/head/blob-head: Pass
      OCI Conformance Test/bad-digest-image/head/blob-head: Pass
    OCI Conformance Test/bad-digest-image/pull: Pass
      OCI Conformance Test/bad-digest-image/pull/blob-get: Pass
      OCI Conformance Test/bad-digest-image/pull/blob-get: Pass
      OCI Conformance Test/bad-digest-image/pull/blob-get: Pass
    OCI Conformance Test/bad-digest-image/delete: Pass
      OCI Conformance Test/bad-digest-image/delete/manifest-delete: Pass
      OCI Conformance Test/bad-digest-image/delete/blob-delete: Pass
      OCI Conformance Test/bad-digest-image/delete/blob-delete: Pass
      OCI Conformance Test/bad-digest-image/delete/blob-delete: Pass
  OCI Conformance Test/missing-manifest: Pass
    OCI Conformance Test/missing-manifest/by-digest: Pass
    OCI Conformance Test/missing-manifest/by-tag: Pass
  OCI Conformance Test/invalid-digest-format: Pass
    OCI Conformance Test/invalid-digest-format/blob-post-put: Pass
    OCI Conformance Test/invalid-digest-format/blob-post-put: Pass
    OCI Conformance Test/invalid-digest-format/manifest-put: Pass
    OCI Conformance Test/invalid-digest-format/manifest-get: Pass
    OCI Conformance Test/invalid-digest-format/delete: Pass
      OCI Conformance Test/invalid-digest-format/delete/manifest-delete: Pass
      OCI Conformance Test/invalid-digest-format/delete/blob-delete: Pass
      OCI Conformance Test/invalid-digest-format/delete/blob-delete: Pass
Configuration:
  registry: 10.1.0.182:8080
  tls: disabled
  repo1: oci-conformance/distribution-test
  repo2: oci-conformance/crossmount-test
  username: ""
  password: ""
  cacheAuth: true
  logging: warn
  apis:
    ping: true
    pull: true
    push: true
    blobs:
      atomic: true
      delete: true
      digestHeader: false
      mountAnonymous: true
      uploadCancel: false
    manifests:
      atomic: true
      delete: true
      digestHeader: false
      tagParam: false
    tags:
      atomic: true
      delete: true
      list: true
    referrer: true
  data:
    image: true
    index: true
    indexList: true
    sparse: false
    artifact: true
    subject: true
    subjectMissing: true
    artifactList: true
    subjectList: true
    dataField: true
    nondistributable: true
    customFields: true
    noLayers: true
    emptyBlob: true
    sha512: true
  roData:
    tags: []
    manifests: []
    blobs: []
    referrers: []
  resultsDir: .
  version: "1.1"
  commit: fcfba1ec55526073f48b2f6d4e3d7eef410ddcbc
  
OCI Conformance Result: Pass
  Disabled......................:          2
  Skip..........................:          4
  Pass..........................:        848
  FAIL..........................:          0
  Error.........................:          0
  Total.........................:        854

API conformance:
  Tag listing...................:       Pass
  Tag delete....................:       Pass
  Tag delete atomic.............:       Pass
  Blob upload cancel............:   Disabled
  Blob push.....................:       Pass
  Blob post only................:       Pass
  Blob post put.................:       Pass
  Blob chunked..................:       Pass
  Blob streaming................:       Pass
  Blob mount....................:       Pass
  Blob anonymous mount..........:       Skip
  Blob get......................:       Pass
  Blob get range................:       Pass
  Blob head.....................:       Pass
  Blob delete...................:       Pass
  Blob delete atomic............:       Pass
  Manifest put by digest........:       Pass
  Manifest put by tag...........:       Pass
  Manifest put with tag params..:   Disabled
  Manifest put with subject.....:       Pass
  Manifest get by digest........:       Pass
  Manifest get by tag...........:       Pass
  Manifest head by digest.......:       Pass
  Manifest head by tag..........:       Pass
  Manifest delete...............:       Pass
  Manifest delete atomic........:       Pass
  Referrers.....................:       Pass
  Ping..........................:       Pass

Data conformance:
  Artifact......................:       Pass
  Artifact Index................:       Pass
  Artifact without Layers.......:       Pass
  Artifacts with Subject........:       Pass
  Bad Digest Image..............:       Pass
  Blobs sha256..................:       Pass
  Blobs sha512..................:       Pass
  Custom Fields.................:       Pass
  Data Field....................:       Pass
  Empty Index...................:       Pass
  Image.........................:       Pass
  Image Uncompressed............:       Pass
  Index.........................:       Pass
  Index with Subject............:       Pass
  Invalid Manifest Digest.......:       Pass
  Image with Large Manifest.....:       Pass
  Missing Subject...............:       Pass
  Nested Index..................:       Pass
  No Layers.....................:       Pass
  Non-distributable Layers......:       Pass
  Digest Algorithm sha512.......:       Pass
  Sparse Manifests..............:   Disabled
  Tag Param.....................:   Disabled
  Tag Param sha512..............:   Disabled
```

More details can be found in the resulting `report.html`.

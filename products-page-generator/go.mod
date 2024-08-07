module github.com/opencontainers/oci-conformance/products-page-generator

go 1.22.6

// See https://github.com/darccio/mergo#important-notes
replace github.com/imdario/mergo => github.com/imdario/mergo v0.3.16

require (
	github.com/Masterminds/semver v1.5.0
	github.com/Masterminds/sprig v2.22.0+incompatible
	github.com/joshdk/go-junit v1.0.0
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/Masterminds/goutils v1.1.1 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/huandu/xstrings v1.5.0 // indirect
	github.com/imdario/mergo v0.0.0-00010101000000-000000000000 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	golang.org/x/crypto v0.26.0 // indirect
)

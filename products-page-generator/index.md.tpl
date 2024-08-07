---
permalink: index.html
---

__Table of Contents:__

- [Distribution Spec](#distribution-spec)
  - [Open Source](#open-source)
{{- with .Submissions -}}
{{- range $key, $value := . }}
{{- if .IsOSS }}
    - [{{ $value.Meta.Name }}](#{{ lower $value.Meta.Name | replace " " "-" }})
{{- end -}}
{{- end -}}
{{ end }}
  - [Hosted](#hosted)
{{- with .Submissions -}}
{{- range $key, $value := . }}
{{- if not .IsOSS }}
    - [{{ $value.Meta.Name }}](#{{ lower $value.Meta.Name | replace " " "-" }})
{{- end }}
{{- end }}
{{ end -}}
- [Image Spec](#image-spec)
- [Runtime Spec](#runtime-spec)

---

## Distribution Spec

Each of the products found below conform to the
[OCI Distribution Specification](https://github.com/opencontainers/distribution-spec),
in varying degrees.

You are encouraged to use this information (as well as other criteria) to make
technical decisions when choosing a product to use as an OCI registry.

### Open Source

{{ with .Submissions -}}
{{- range $key, $value := . }}
{{- if .IsOSS }}
#### {{ $value.Meta.Name }}

<img src="{{ $value.Meta.ProductLogoURL}}" style="max-width:120px"/>

> {{ $value.Meta.Description }}

__Homepage:__ [{{ $value.Meta.WebsiteURL }}]({{ $value.Meta.WebsiteURL }})


{{ if ne $value.Meta.DocumentationURL $value.Meta.RepoURL }}
__Documentation:__ [{{ $value.Meta.DocumentationURL }}]({{ $value.Meta.DocumentationURL }})
{{ end }}

{{ if $value.Meta.RepoURL }}
__Repository:__ [{{ $value.Meta.RepoURL }}]({{ $value.Meta.RepoURL }})
{{ end }}

__Vendor:__ {{ $value.Meta.Vendor }}

__Latest software version tested__: `{{ $value.Meta.Version }}`

__Latest spec version supported__: `{{ $value.LatestVersion }}` <sub><br/>**[Test report](./static/{{ $value.LatestVersion }}/reports/{{ $key }}/)**
| **[How to reproduce](./static/{{ $value.LatestVersion }}/instructions/{{ $key }}/)**</sub>

{{ if $value.BadgesMarkdown }}
__Live results:__ {{ $value.BadgesMarkdown }}
{{ end }}

| Supported workflows: | {{ range $_, $version := $value.AllVersions }}{{ $version }} | {{ end }}
| ---------------------- | ------ |
{{- range $_, $v := (index $value.Workflows $value.LatestVersion) }}
| {{ $v.Name }} {{ if $v.Required }}(required){{ end }} | {{ range $_, $version := $value.AllVersions }}{{ range $_, $workflow := (index $value.Workflows $version) }}{{ if eq $workflow.Name $v.Name }}{{ if $workflow.Supported }} ✅  {{ else }}  ✖️ {{ end }}   | {{ end }}{{ end }}{{ end }}
{{- end }}

{{ end }}
{{- end }}
{{ end }}

### Hosted

{{ with .Submissions -}}
{{- range $key, $value := . }}
{{- if not .IsOSS }}
#### {{ $value.Meta.Name }}

<img src="{{ $value.Meta.ProductLogoURL}}" style="max-width:120px"/>

> {{ $value.Meta.Description }}

__Homepage:__ [{{ $value.Meta.WebsiteURL }}]({{ $value.Meta.WebsiteURL }})

__Documentation:__ [{{ $value.Meta.DocumentationURL }}]({{ $value.Meta.DocumentationURL }})

__Vendor:__ {{ $value.Meta.Vendor }}

__Latest spec version supported__: `{{ $value.LatestVersion }}` <sub><br/>**[Test report](./static/{{ $value.LatestVersion }}/reports/{{ $key }}/)**
| **[How to reproduce](./static/{{ $value.LatestVersion }}/instructions/{{ $key }}/)**</sub>

{{ if $value.BadgesMarkdown }}
__Live results:__ {{ $value.BadgesMarkdown }}
{{ end }}

| Supported workflows: | {{ range $_, $version := $value.AllVersions }}{{ $version }} | {{ end }}
| ---------------------- | ------ |
{{- range $_, $v := (index $value.Workflows $value.LatestVersion) }}
| {{ $v.Name }} {{ if $v.Required }}(required){{ end }} | {{ range $_, $version := $value.AllVersions }}{{ range $_, $workflow := (index $value.Workflows $version) }}{{ if eq $workflow.Name $v.Name }}{{ if $workflow.Supported }} ✅  {{ else }}  ✖️ {{ end }}   | {{ end }}{{ end }}{{ end }}
{{- end }}

{{ end }}
{{- end }}
{{ end }}

---

## Image Spec

*There is not yet a submission process in place for products conforming to the
[OCI Image Format Specification](https://github.com/opencontainers/image-spec).*

---

## Runtime Spec

*There is not yet a submission process in place for products conforming to the
[OCI Runtime Specification](https://github.com/opencontainers/runtime-spec).*

---
<br/>
Last generated: `{{ now | date "Jan 2 15:04:05 MST" }}`

Git commit: `{{ .GitCommit }}`

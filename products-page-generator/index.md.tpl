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

{{- with .Submissions -}}
{{- range $key, $value := . }}
{{- if .IsOSS }}

#### {{ $value.Meta.Name }}

<img src="{{ $value.Meta.ProductLogoURL}}" style="max-width:120px"/>

> {{ $value.Meta.Description }}

__Homepage:__ [{{ $value.Meta.WebsiteURL }}]({{ $value.Meta.WebsiteURL }})

{{- if ne $value.Meta.DocumentationURL $value.Meta.RepoURL }}

__Documentation:__ [{{ $value.Meta.DocumentationURL }}]({{ $value.Meta.DocumentationURL }})
{{- end }}

{{- if $value.Meta.RepoURL }}

__Repository:__ [{{ $value.Meta.RepoURL }}]({{ $value.Meta.RepoURL }})
{{- end }}

__Vendor:__ {{ $value.Meta.Vendor }}

__Latest software version tested__: {{ $value.Meta.Version }}

__Latest spec version supported__: {{ $value.LatestVersion }}

{{- if $value.BadgesMarkdown }}

__Live results:__ {{ $value.BadgesMarkdown }}
{{- end }}

{{- range $_, $version := $value.AllVersions }}

<div class="version" style="padding: .25em 0 .25em .75em; border: 2px solid #ddd; border-radius: 2px;">
<details {{- if eq $value.LatestVersion $version }} open="true" {{- end}}>
<summary><strong>Version: {{ $version }}</strong></summary>

<ul>
<li><a href="./static/{{ $version }}/reports/{{ $key }}/">Test report</a></li>
<li><a href="./static/{{ $version }}/instructions/{{ $key }}/">How to reproduce</a></li>
</ul>

<div class="workflows" style="margin: .25em; padding: .25em; background: #eee;">
<strong>Workflows:</strong><br/>
{{- range $_, $workflow := (index $value.Workflows $version) }}
{{- if or (eq $workflow.Supported.String "Skip") (eq $workflow.Supported.String "Disabled") }}
✖️ {{ $workflow.Name }}<br/>
{{- else if ne $workflow.Supported.String "Pass" }}
❌ {{ $workflow.Name }}<br/>
{{- end }}{{/* non-passing workflows */}}
{{- end }}{{/* range over Workflows */}}

{{- if gt (index (index $value.Summary $version).Counts "Pass") 4 }}<details>
<summary><strong>✅✅...✅ {{- (index (index $value.Summary $version).Counts "Pass") }} of {{ len (index $value.Workflows $version) }} workflows passed</strong></summary>{{- end }}
{{- range $_, $workflow := (index $value.Workflows $version) }}
{{- if eq $workflow.Supported.String "Pass" }}
✅ {{ $workflow.Name }}<br/>
{{- end }}{{/* passing workflows */}}
{{- end }}{{/* range over Workflows */}}
{{- if gt (index (index $value.Summary $version).Counts "Pass") 4 }}</details>{{- end }}
</div>

</details>
</div>
{{- end }}{{/* range over AllVersions */}}
<br/>
{{- end }}{{/* if OSS */}}
{{- end }}{{/* range over submissions */}}
{{- end }}{{/* with .Submissions */}}

### Hosted

{{- with .Submissions -}}
{{- range $key, $value := . }}
{{- if not .IsOSS }}

#### {{ $value.Meta.Name }}

<img src="{{ $value.Meta.ProductLogoURL}}" style="max-width:120px"/>

> {{ $value.Meta.Description }}

__Homepage:__ [{{ $value.Meta.WebsiteURL }}]({{ $value.Meta.WebsiteURL }})

{{- if ne $value.Meta.DocumentationURL $value.Meta.RepoURL }}

__Documentation:__ [{{ $value.Meta.DocumentationURL }}]({{ $value.Meta.DocumentationURL }})
{{- end }}

{{- if $value.Meta.RepoURL }}

__Repository:__ [{{ $value.Meta.RepoURL }}]({{ $value.Meta.RepoURL }})
{{- end }}

__Vendor:__ {{ $value.Meta.Vendor }}

__Latest software version tested__: {{ $value.Meta.Version }}

__Latest spec version supported__: {{ $value.LatestVersion }}

{{- if $value.BadgesMarkdown }}

__Live results:__ {{ $value.BadgesMarkdown }}
{{- end }}

{{- range $_, $version := $value.AllVersions }}

<div class="version" style="padding: .25em 0 .25em .75em; border: 2px solid #ddd; border-radius: 2px;">
<details {{- if eq $value.LatestVersion $version }} open="true" {{- end}}>
<summary><strong>Version: {{ $version }}</strong></summary>

<ul>
<li><a href="./static/{{ $version }}/reports/{{ $key }}/">Test report</a></li>
<li><a href="./static/{{ $version }}/instructions/{{ $key }}/">How to reproduce</a></li>
</ul>

<div class="workflows" style="margin: .25em; padding: .25em; background: #eee;">
<strong>Workflows:</strong><br/>
{{- range $_, $workflow := (index $value.Workflows $version) }}
{{- if or (eq $workflow.Supported.String "Skip") (eq $workflow.Supported.String "Disabled") }}
✖️ {{ $workflow.Name }}<br/>
{{- else if ne $workflow.Supported.String "Pass" }}
❌ {{ $workflow.Name }}<br/>
{{- end }}{{/* non-passing workflows */}}
{{- end }}{{/* range over Workflows */}}

{{- if gt (index (index $value.Summary $version).Counts "Pass") 4 }}<details>
<summary><strong>✅✅...✅ {{- (index (index $value.Summary $version).Counts "Pass") }} of {{ len (index $value.Workflows $version) }} workflows passed</strong></summary>{{- end }}
{{- range $_, $workflow := (index $value.Workflows $version) }}
{{- if eq $workflow.Supported.String "Pass" }}
✅ {{ $workflow.Name }}<br/>
{{- end }}{{/* passing workflows */}}
{{- end }}{{/* range over Workflows */}}
{{- if gt (index (index $value.Summary $version).Counts "Pass") 4 }}</details>{{- end }}
</div>

</details>
</div>
{{- end }}{{/* range over AllVersions */}}
<br/>
{{- end }}{{/* if OSS */}}
{{- end }}{{/* range over submissions */}}
{{- end }}{{/* with .Submissions */}}

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

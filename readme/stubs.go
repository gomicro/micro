package readme

const (
	stub = `# {{ .Name | Title }}{{ if .Installable }}

## Installation

` + "```" + `
go get {{ .Source }}
` + "```" + `{{ end }}

## Versioning

The project will be versioned in accordance with [Semver 2.0.0](https://semver.org).{{ if (ne .ReleaseURL "") }} See the [releases]({{ .ReleaseURL }}) section for the latest version.{{ end }} Until version 1.0.0 the project is considered to be unstable.

It is always highly recommended to vendor the version you are using.

## License
See [LICENSE.md](./LICENSE.md) for more information.
`
)

# Service

[![GitHub Workflow Status](https://img.shields.io/github/workflow/status/gomicro/service/Build/master)](https://github.com/gomicro/service/actions?query=workflow%3ABuild+branch%3Amaster)
[![Go Reportcard](https://goreportcard.com/badge/github.com/gomicro/service)](https://goreportcard.com/report/github.com/gomicro/service)
[![License](https://img.shields.io/github/license/gomicro/service.svg)](https://github.com/gomicro/service/blob/master/LICENSE.md)
[![Release](https://img.shields.io/github/release/gomicro/service.svg)](https://github.com/gomicro/service/releases/latest)

Service is a bootstrapping generator for creating new microservices.

## Installing
```
go get github.com/gomicro/service
```

## Usage

```
Generate a bootstrap of a new microservice

Usage:
  service [flags]
  service [command]

Available Commands:
  help        Help about any command
  version     Display the version

Flags:
      --db              whether the service will have a database or not
  -h, --help            help for service
      --installable     whether the service will be installable or not
      --name string     service name (default "service")
      --org string      organization name (default "gomicro")
      --source string   source location (default "https://github.com/gomicro/service")

Use "service [command] --help" for more information about a command.
```

## Versioning

The project will be versioned in accordance with [Semver 2.0.0](https://semver.org). See the [releases](https://github.com/gomicro/butcher/releases) section for the latest version. Until version 1.0.0 the project is considered to be unstable.

It is always highly recommended to vendor the version you are using.

## License
See [LICENSE.md](./LICENSE.md) for more information.
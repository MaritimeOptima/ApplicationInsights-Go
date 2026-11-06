# ApplicationInsights-Go (maintained fork)

This is a fork of [Microsoft's Application Insights SDK for Go](https://github.com/microsoft/ApplicationInsights-Go).
Microsoft stopped maintaining the upstream SDK around 2020. This fork keeps its dependencies up to date.

## Usage

The module path is still `github.com/microsoft/ApplicationInsights-Go`, so
existing imports keep working. To use this fork, add a `replace` directive to
your `go.mod`.

## Changes vs upstream

- Updated dependencies: `code.cloudfoundry.org/clock` v1.88.0 and
  `github.com/gofrs/uuid/v5` (upstream used `gofrs/uuid` v3). Dropped the unused indirect dependencies.
- Requires Go 1.27 (upstream: Go 1.12).
- Ran `go fix` over the code, for example `interface{}` is now `any`.
- SDK version bumped from 0.4.4 to 0.4.7.
- Dependabot enabled for dependency updates.
- Removed the schema generator script (`generateschema.ps1`); any changes
to the shcema must be handled manually.
- Removed the upstream Travis CI config and Microsoft's community files
  (Code of Conduct, Security, Contributing, PR template).

## Contributing

Before releasing a new version, update `Version` in `appinsights/package.go`.

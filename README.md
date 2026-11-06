# ApplicationInsights-Go (maintained fork)

This is a fork of [Microsoft's Application Insights SDK for Go](https://github.com/microsoft/ApplicationInsights-Go).
Microsoft stopped maintaining the upstream SDK around 2020. This fork keeps its dependencies up to date.

## Usage

The module path is still `github.com/microsoft/ApplicationInsights-Go`, so
existing imports keep working. To use this fork, add a `replace` directive to
your `go.mod`.

## Changes vs upstream

- Dependencies are kept up to date.
- Requires a new Go version.
- Ran `go fix` over the code, for example `interface{}` is now `any`.
- SDK version is bumped.
- Dependabot enabled for dependency updates.
- Removed the schema generator script (`generateschema.ps1`); any changes
to the schema must be handled manually.
- Removed the Travis CI config and Microsoft's community files
  (Code of Conduct, Security, Contributing, PR template).
- Removed a bunch of lines from .gitignore

## Contributing

Before releasing a new version, update `Version` in `appinsights/package.go`.

# go-appinfo

A module to provide application information, extending the standard values during build time.

This module can be used to compile additional application information into a binary and retrieve it together with the
standard information at runtime. For example, it allows building user agent strings (via `UserAgent()` function).

The following values are determined automatically, but partially require building from within a git repository:
- version (via [Go modules](https://go.dev/ref/mod#versions))
- commit hash
- tree state (dirty/clean), i.e. whether there are uncommitted changes in the working directory
- compiler
- Go version used for compilation
- platform/architecture

Other values can be set during build time via `ldflags`:
- name
- metadata
- build date

If no name is set explicitly, it defaults to the [module path](https://go.dev/ref/mod#module-path)'s part after the last
`/`.

This example shows all settable attributes:

```shell
meta="buildhost:$HOSTNAME"
buildDate=$(date -u +'%Y-%m-%dT%H:%M:%SZ')
go build -ldflags "\
    -X github.com/ionos-cloud/go-appinfo.name=my-app \
    -X github.com/ionos-cloud/go-appinfo.metadata=$meta \
    -X github.com/ionos-cloud/go-appinfo.buildDate=$buildDate \
    " ./cmd
```

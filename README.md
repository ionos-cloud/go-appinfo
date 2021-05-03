# go-version

A package to provide version information settable during build time.

This package can be used to compile the version information into a binary. Useful, for example, when building user agent
strings.

It is possible to set pre-release information either directly or by using the version string only, e.g. if the
pre-release is contained as git tag.

Go build example with magic:

```console
tag=$(git describe --tag)
commit=$(git rev-parse --short HEAD)
meta=builddate-$(date -Iseconds)
go build -ldflags "\
    -X github.com/ionos-cloud/go-version.version=$tag \
    -X github.com/ionos-cloud/go-version.metadata=$meta \
    -X github.com/ionos-cloud/go-version.gitCommit=$commit \
    " ./...
```
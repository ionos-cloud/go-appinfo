# go-appinfo

A package to provide application information settable during build time.

This package can be used to compile the application information into a binary. Useful, for example, when 
building user agent strings (which is actually provided as function). 

This example shows all settable attributes:

```console
tag=$(git describe --tag)
commit=$(git rev-parse --short HEAD)
meta=builddate-$(date -Iseconds)
go build -ldflags "\
    -X github.com/ionos-cloud/go-appinfo.name=my-app \
    -X github.com/ionos-cloud/go-appinfo.version=$tag \
    -X github.com/ionos-cloud/go-appinfo.metadata=$meta \
    -X github.com/ionos-cloud/go-appinfo.gitCommit=$commit \
    " ./...
```
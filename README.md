# go-appinfo

A package to provide application information settable during build time.

This package can be used to compile the application information into a binary. Useful, for example, when 
building user agent strings (which is actually provided as function). 

This package does not make assumptions. It will provide the Compiler, Platform and GoVersion since these
values are unambiguously determinable by the Runtime. Other fields need to be set by the builder with ldflags.

This example shows all settable attributes:

```shell
tag=$(git describe --tag)
commit=$(git rev-parse --short HEAD)
meta="buildhost:$HOSTNAME"
buildDate=$(date -u +'%Y-%m-%dT%H:%M:%SZ')
treeState=$(test -n "`git status --porcelain`" && echo 'dirty' || echo 'clean')
go build -ldflags "\
    -X github.com/ionos-cloud/go-appinfo.name=my-app \
    -X github.com/ionos-cloud/go-appinfo.version=$tag \
    -X github.com/ionos-cloud/go-appinfo.metadata=$meta \
    -X github.com/ionos-cloud/go-appinfo.buildDate=$buildDate \
    -X github.com/ionos-cloud/go-appinfo.gitCommit=$commit \
    -X github.com/ionos-cloud/go-appinfo.gitTreeState=$treeState \
    " ./...
```
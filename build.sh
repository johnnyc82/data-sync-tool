#!/bin/zsh



export GOPRIVATE=github.com/ivystreetweb/config-tool


go get ./...
go mod download
gofmt -w ./..

git describe --tags --abbrev=0 > cmd/hsync/version.txt

cd cmd/hsync/
go install
# go build -v -a

now="$(date)"
printf "build at: %s\n" "$now"
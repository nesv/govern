#!/usr/bin/env bash

function cleanup {
    rm -f govern
}
trap cleanup EXIT

set -e

go test ./...
go build -o govern ./...
./govern --version
./govern facts
./govern facts --output=yaml
./govern --verbose --path="testfiles" play site.yml --inventory=hosts
cleanup

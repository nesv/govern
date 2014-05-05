#!/usr/bin/env bash

function cleanup {
    rm -f govern
}

set -e

trap cleanup SIGINT SIGKILL

go test ./...
go build -o govern ./...
./govern -v=1 -logtostderr -path="testfiles" $@
cleanup

#!/bin/sh

cd "$(dirname "$0")"
exec gofmt -l -w -e -s -tabs *.go $(find src/app -type f -name '*.go')

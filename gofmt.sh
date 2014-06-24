#!/bin/sh

find "$(dirname "$0")" -type f -name '*.go' | xargs -r gofmt -l -w -e -s

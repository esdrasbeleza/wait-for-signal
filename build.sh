#!/bin/sh

VERSION=0.1

if [ -z "${GOOS}" ]; then
    export GOOS=linux
fi

if [ -z "${GOARCH}" ]; then
    export GOARCH=amd64
fi

go build -v -o bin/wait-for-signal-${GOOS}-${GOARCH}-${VERSION} ./cmd/wait-for-signal

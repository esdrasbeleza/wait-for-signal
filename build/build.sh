#!/bin/sh

VERSION=0.1

if [ -z "${GOOS}" ]; then
    export GOOS=linux
fi

if [ -z "${GOARCH}" ]; then
    export GOARCH=amd64
fi

if [ ! -d "bin/" ]; then
    mkdir bin
fi

go build -v -o bin/wait-for-signal-${GOOS}-${GOARCH}-${VERSION} -ldflags="-s -w" ../cmd/wait-for-signal
go build -v -o bin/send-signal-${GOOS}-${GOARCH}-${VERSION} -ldflags="-s -w" ../cmd/send-signal

#!/bin/sh
GOOS=linux GOARCH=amd64 ./build.sh
GOOS=linux GOARCH=386 ./build.sh
GOOS=darwin GOARCH=amd64 ./build.sh

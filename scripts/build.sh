#!/usr/bin/env sh
set -x
NAME=${NAME-3ber}

. ./scripts/version.sh

GOOS=linux GOARCH=amd64 go build -ldflags "${VERSIONFLAGS}" -o bin/${NAME}
GOOS=linux GOARCH=arm64 go build -ldflags "${VERSIONFLAGS}" -o bin/${NAME}-arm64
GOOS=darwin GOARCH=amd64 go build -ldflags "${VERSIONFLAGS}" -o bin/${NAME}-darwin
GOOS=darwin GOARCH=arm64 go build -ldflags "${VERSIONFLAGS}" -o bin/${NAME}-darwin-arm64
GOOS=windows go build -ldflags "${VERSIONFLAGS}" -o bin/${NAME}.exe

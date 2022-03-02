#!/usr/bin/env sh
set -e
set -x
NAME=${NAME-3ber}
CROSS=${CROSS-false}

. ./scripts/version.sh

if [ "$CROSS" == "true" ]; then
	GOOS=linux GOARCH=amd64 go build -ldflags "${VERSIONFLAGS}" -o bin/${NAME}-Linux-x86_64
	GOOS=linux GOARCH=arm64 go build -ldflags "${VERSIONFLAGS}" -o bin/${NAME}-Linux-arm64
	GOOS=darwin GOARCH=amd64 go build -ldflags "${VERSIONFLAGS}" -o bin/${NAME}-Darwin-x86_64
	GOOS=darwin GOARCH=arm64 go build -ldflags "${VERSIONFLAGS}" -o bin/${NAME}-Darwin-arm64
	GOOS=windows go build -ldflags "${VERSIONFLAGS}" -o bin/${NAME}.exe
else
	go build -ldflags "${VERSIONFLAGS}" -o bin/${NAME}
fi

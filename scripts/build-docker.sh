#!/usr/bin/env sh
set -x
NAME=${NAME-3ber}

. ./scripts/version.sh

go build -ldflags "${VERSIONFLAGS}" -o bin/${NAME}

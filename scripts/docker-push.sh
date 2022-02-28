#!/usr/bin/env sh
set -e
set -x
IMAGE=${IMAGE-gcr.io/freshly-docker/3ber}

. ./scripts/version.sh

docker push ${IMAGE}:${VERSION}

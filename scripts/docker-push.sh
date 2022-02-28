#!/usr/bin/env sh
set -x
IMAGE=${IMAGE-gcr.io/freshly-docker/3ber}

. ./scripts/version.sh

docker push ${IMAGE}:${VERSION}

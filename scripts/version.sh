#!/usr/bin/env sh
set -x
REPO=${REPO-github.com/freshly/3ber}
GO=${GO-go}
ARCH=${ARCH:-$("${GO}" env GOARCH)}
SUFFIX="-${ARCH}"
GIT_TAG=
COMMIT=

if [ -d .git ]; then
    if [ -z "$GIT_TAG" ]; then
        GIT_TAG=$(git tag -l --contains HEAD | head -n 1)
    fi
    if [ -n "$(git status --porcelain --untracked-files=no)" ]; then
        DIRTY="-dirty"
    fi

    COMMIT=$(git log -n3 --pretty=format:"%H %ae" | cut -f1 -d\  | head -1)
    if [ -z "${COMMIT}" ]; then
        COMMIT=$(git rev-parse HEAD || true)
    fi
fi

VERSION="${COMMIT:0:8}${DIRTY}"
VERSIONFLAGS="
    -X ${REPO}/pkg/version.Version=${VERSION}
    -X ${REPO}/pkg/version.GitCommit=${COMMIT:0:8}
"
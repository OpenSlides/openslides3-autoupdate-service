#!/bin/bash

VERSION_INFO=$(git log -1 --pretty=format:"%h%x09%an%x09%ad%x09%s")

docker build --tag openslides/openslides-autoupdate \
    --build-arg VERSION_INFO="$VERSION_INFO" \
    .

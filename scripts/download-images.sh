#!/bin/bash

set -euo pipefail

TESTDATA=$(realpath $(dirname $0)/../integration/testdata)
IMAGES_DIR=${TESTDATA}/images/
FIRST=busybox:1.35.0
SECOND=alpine:3.16.2

mkdir -p $IMAGES_DIR

echo "Downloading images..."
docker pull $FIRST
docker pull $SECOND

echo "Saving images..."
docker save $FIRST -o $IMAGES_DIR/single.tar
docker save $FIRST $SECOND -o $IMAGES_DIR/multiple.tar

echo "This is not a tar file..." > $IMAGES_DIR/invalid.tar

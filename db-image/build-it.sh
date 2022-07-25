#!/usr/bin/env bash

# not pushing to registry because I don't want to bother with multi-arch builds
# just build and run it on the same machine locally
docker build . -t query-node-deps-pg-image
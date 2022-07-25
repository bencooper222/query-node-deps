#!/usr/bin/env bash

# run build-it.sh first

docker run -d \
 -p 5432:5432 \
 query-node-deps-pg-image



echo "If you see a meaningless hash above, that's the ID of your docker container (so this script probably worked)."
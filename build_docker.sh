#!/bin/sh

docker build --rm -t trigun117/grabber .
echo "$DOCKER_PASS" | docker login -u "$DOCKER_USER" --password-stdin
docker push trigun117/grabber
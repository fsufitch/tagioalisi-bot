#!/usr/bin/env bash

set -ex

docker build -f Dockerfile.discordbot -t docker.io/fsufitch/tagioalisi-discordbot .
docker build -f Dockerfile.webapp -t docker.io/fsufitch/tagioalisi-webapp .
docker build -f Dockerfile.grpcwebproxy -t docker.io/fsufitch/tagioalisi-grpcwebproxy .


#!/bin/sh

echo "Be sure to run docker login before this script"
echo "docker login "

VER="0.0.4"
echo "building binary v.$VER"
GO111MODULE=on go mod vendor
CGO_ENABLED=0 GOOS=linux go build -o main
echo "Building docker v.$VER"
docker build -t url:v$VER .
echo "Pushing v.$VER"
docker push url:v$VER

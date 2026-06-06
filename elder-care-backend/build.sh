#!/bin/sh

mkdir -p build
cp -R ./migration ./build
cp -R ./script ./build
cp -R ./config ./build
cp cmd/elder_care_backend_config.ini ./build

GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./build/elder_care_backend ./cmd

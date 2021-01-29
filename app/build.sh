#!/bin/bash -e

mkdir -p assets

CGO_ENABLED=0 go build -o assets/main main.go 

docker build -t dirtyci-app .

rm -r assets
#!/usr/bin/env bash

set -e
echo "" > coverage.txt

for d in $(go list ./... | grep -v vendor); do
    echo "go test -covermode=count -coverprofile=profile.out $d\n"
    go test -coverprofile=profile.out -covermode=count $d
    if [ -f profile.out ]; then
        cat profile.out >> coverage.txt
        rm profile.out
    fi
done
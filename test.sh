#!/usr/bin/env bash

set -e
echo "" > coverage.txt

for d in $(go list ./... | grep -v vendor); do
    echo "go test -covermode=atomic -coverprofile=profile.out  $d\n"
    go test -coverprofile=profile.out -covermode=atomic $d
    if [ -f profile.out ]; then
        cat profile.out >> coverage.txt
        rm profile.out
    fi
done

go tool cover -html=coverage.txt -o coverage.html

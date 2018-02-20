#!/usr/bin/env bash

while true; do
    go-bindata -debug template/...
    go run $(find . -type f -name "*.go" -not -name "*_test.go")  
    echo
    echo Reloading app...
    echo
done

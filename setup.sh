#!/usr/bin/env bash

if ! command -v go > /dev/null; then
    echo "go not installed"
    exit 1
fi

exec go build -o jwt-decoder ./cmd/jwt-decoder/main.go 


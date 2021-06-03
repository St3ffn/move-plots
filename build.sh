#!/bin/bash

SCRIPT_PATH=$(dirname "$(realpath -s "$0")")

# git commit id
GIT_COMMIT=$(git rev-list -1 HEAD)
# most recent tag
GIT_VERSION_TAG=$(git describe --tags --abbrev=0)

go build -o move-plots -ldflags="-X 'main.GitVersion=$GIT_VERSION_TAG'" "$SCRIPT_PATH/main.go"
env GOOS=linux GOARCH=amd64 go build -o move-plots-linux-amd64 -ldflags="-X 'main.GitVersion=$GIT_VERSION_TAG'" "$SCRIPT_PATH/main.go"
env GOOS=darwin GOARCH=amd64 go build -o move-plots-darwin-amd64 -ldflags="-X 'main.GitVersion=$GIT_VERSION_TAG'" "$SCRIPT_PATH/main.go"
env GOOS=linux GOARCH=arm64 go build -o move-plots-linux-arm64 -ldflags="-X 'main.GitVersion=$GIT_VERSION_TAG'" "$SCRIPT_PATH/main.go"
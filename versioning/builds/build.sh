#!/usr/bin/env bash

VERSION=$(git describe --always)
LAST_COMMIT_USER="$(tr -d '[:space:]' <<<"$(git log -1 --format=%cn)<$(git log -1 --format=%ce)>")"
LAST_COMMIT_HASH=$(git log -1 --format=%H)
LAST_COMMIT_TIME=$(git log -1 --format=%cd --date=format:'%Y-%m-%d_%H:%M:%S')

go build -ldflags "-X main.appVersion=$VERSION -X main.lastCommitTime=$LAST_COMMIT_TIME -X main.lastCommitHash=$LAST_COMMIT_HASH -X main.lastCommitUser=$LAST_COMMIT_USER -X main.buildTime=$(date -u +%Y-%m-%d_%H:%M:%S)" main.go

./main
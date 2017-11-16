#!/usr/bin/env bash

set -e

basedir=$(dirname "$0")

# Protocol files fetching taken from github.com/cyrus-and/chrome-remote-interface
base='https://chromium.googlesource.com'
curl -s "$base/chromium/src/+/master/third_party/WebKit/Source/core/inspector/browser_protocol.json?format=TEXT" | base64 -D > $basedir/cri-gen/json-proto/browser_protocol.json
curl -s "$base/v8/v8/+/master/src/inspector/js_protocol.json?format=TEXT" | base64 -D > $basedir/cri-gen/json-proto/js_protocol.json

cd $basedir/cri-gen
go run createFiles.go jsonDefinitions.go main.go
cd - > /dev/null 2>&1
gofmt -s -w .

#!/usr/bin/env bash

set -e

basedir=$(dirname "$0")

# Protocol files fetching taken from github.com/cyrus-and/chrome-remote-interface
base='https://raw.githubusercontent.com/ChromeDevTools/devtools-protocol/master/json'
curl -s "$base/browser_protocol.json" > $basedir/cri-gen/json-proto/browser_protocol.json
curl -s "$base/js_protocol.json" > $basedir/cri-gen/json-proto/js_protocol.json

cd $basedir/cri-gen
go run createFiles.go jsonDefinitions.go main.go
cd - > /dev/null 2>&1

branch_name="$(git rev-parse --symbolic-full-name --abbrev-ref HEAD)"
sed -i -e "s/branch=[^) ]*)/branch=$branch_name)/g" README.md
rm README.md-e

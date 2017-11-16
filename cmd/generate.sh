set -e

cd cmd/cri-gen
go run createFiles.go jsonDefinitions.go main.go
gofmt -s -w ../../

#!/bin/bash

echo "go install github.com/golang/protobuf/protoc-gen-go"
(mkdir -p test/tools && cd test/tools && go install github.com/golang/protobuf/protoc-gen-go)

echo "go install cmd/protoc-gen-go-grpc"
(mkdir -p cmd/protoc-gen-go-grpc && cd cmd/protoc-gen-go-grpc && go install .)

mkdir -p proto
protoc echo.proto --go_opt=paths=source_relative --go_out=plugins=grpc:proto

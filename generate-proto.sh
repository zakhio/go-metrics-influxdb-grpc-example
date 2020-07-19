#!/bin/bash

GRPC_GO_VERSION=v1.30.0
GRPC_GO_PATH=tmp/grpc-go

echo "go install github.com/golang/protobuf/protoc-gen-go"
go install github.com/golang/protobuf/protoc-gen-go

echo "go install grpc-go/cmd/protoc-gen-go-grpc"
test -d "tmp/grpc-go" && (cd $GRPC_GO_PATH && git pull origin $GRPC_GO_VERSION && true) || git clone -b $GRPC_GO_VERSION https://github.com/grpc/grpc-go $GRPC_GO_PATH
(cd $GRPC_GO_PATH/cmd/protoc-gen-go-grpc && go install .)

echo "protoc echo.proto --go_opt=paths=source_relative --go_out=plugins=grpc:proto"
mkdir -p proto
protoc echo.proto --go_opt=paths=source_relative --go_out=plugins=grpc:proto

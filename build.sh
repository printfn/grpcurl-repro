#!/usr/bin/env bash
set -euo pipefail
cd "$(dirname "$0")"

PATH="$PATH:$(go env GOPATH)/bin"

# go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
# go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

(cd protobuf-shared && protoc \
	--go_out=sharedpb \
	--go_opt=paths=source_relative \
	--go-grpc_out=sharedpb \
	--go-grpc_opt=paths=source_relative \
	--proto_path=. \
	message.proto)

(cd protobuf-shared && go mod tidy)

(cd protobuf-service && protoc \
	--go_out=servicepb \
	--go_opt=paths=source_relative \
	--go-grpc_out=servicepb \
	--go-grpc_opt=paths=source_relative \
	--proto_path=. \
	--proto_path=servicepb \
	--proto_path=/usr/local/include \
	service.proto)

(cd protobuf-service && go mod tidy)

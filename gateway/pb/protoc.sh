#!/bin/bash
cd "$(dirname "$0")" || exit 1
protoc -I=. --go_out=. --go-grpc_out=. *.proto

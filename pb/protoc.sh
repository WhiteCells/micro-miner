#!/bin/bash

cd "$(dirname "$0")" || exit 1

rm -rf ../gateway/pb/
rm -rf ../services/hiveos/pb/
rm -rf ../services/system/pb/

mkdir ../gateway/pb/
mkdir ../services/hiveos/pb/
mkdir ../services/system/pb/

# protoc -I=./hiveos --go_out=../gateway/pb --go-grpc_out=../gateway/pb ./hiveos/*.proto
# protoc -I=./system --go_out=../gateway/pb --go-grpc_out=../gateway/pb ./system/*.proto
# protoc -I=./hiveos --go_out=../services/hiveos/pb --go-grpc_out=../services/hiveos/pb ./hiveos/*.proto
# protoc -I=./system --go_out=../services/system/pb --go-grpc_out=../services/system/pb ./system/*.proto

# protoc -I=. --go_out=../services/system/pb --go-grpc_out=../services/system/pb system.proto
# protoc -I=. -I=pb --go_out=../services/system/pb --go-grpc_out=../services/system/pb pb/system.proto


protoc -I=. --go_out=../gateway/pb --go-grpc_out=../gateway/pb *.proto
protoc -I=. --go_out=../services/hiveos/pb --go-grpc_out=../services/hiveos/pb hiveos.proto
protoc -I=. --go_out=../services/system/pb --go-grpc_out=../services/system/pb system.proto

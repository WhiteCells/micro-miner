package main

import (
	"log"
	"net"

	"micro-user/userpb"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":50011")
	if err != nil {
		//
		log.Fatal("")
	}

	grpcServer := grpc.NewServer()
	userpb.LoginReq
}

package main

import (
	"log"
	"net"

	"micro-user/grpcs"
	pb "micro-user/pb/userpb"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":50011")
	if err != nil {
		log.Fatalf("micro-user failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &grpcs.UserServiceServer{})
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("micro-user failed to serve: %v", err)
	}
}

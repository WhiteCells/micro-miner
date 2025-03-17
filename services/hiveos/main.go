package main

import (
	"log"
	"micro-hiveos/grpcs"
	"micro-hiveos/pb"
	"net"

	"google.golang.org/grpc"
)

func init() {

}

func main() {
	listener, err := net.Listen("tcp", ":50031")
	if err != nil {
		log.Fatalf("micro-hiveos failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterHiveosServiceServer(grpcServer, &grpcs.HiveosServiceServer{})
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("micro-hiveos failed to serve: %v", err)
	}
}

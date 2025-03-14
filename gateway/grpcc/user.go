package grpcc

import (
	pb "gateway/pb/userpb"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var UserGrpcClient pb.UserServiceClient

func InitUserGrpcClient() {
	client, err := grpc.NewClient(
		"localhost:50011",
		// "dns:///micro-user",
		grpc.WithTransportCredentials((insecure.NewCredentials())),
		// grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		//
		log.Fatalf("failed to init user grpc client: %v", err)
	}
	UserGrpcClient = pb.NewUserServiceClient(client)
}

package grpcc

import (
	"gateway/pb"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var SystemGrpcClient pb.SystemServiceClient

func InitSystemGrpcClient() {
	client, err := grpc.NewClient(
		"localhost:50011",
		// "dns:///micro-system",
		grpc.WithTransportCredentials((insecure.NewCredentials())),
		// grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		//
		log.Fatalf("failed to init user grpc client: %v", err)
	}
	SystemGrpcClient = pb.NewSystemServiceClient(client)
}

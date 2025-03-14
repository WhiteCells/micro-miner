package grpcc

import (
	pb "gateway/pb/userpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewUserGrpcClient() pb.UserServiceClient {
	client, err := grpc.NewClient("localhost:9001", grpc.WithTransportCredentials((insecure.NewCredentials())))
	if err != nil {
		//
		return nil
	}
	return pb.NewUserServiceClient(client)
}

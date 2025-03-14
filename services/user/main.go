package main

import (
	"log"
	"net"

	"micro-user/grpcs"
	pb "micro-user/pb/userpb"
	"micro-user/utils"

	"google.golang.org/grpc"
)

func init() {
	if err := utils.InitConfig("config.yml", "yml"); err != nil {
		log.Fatalf("init config error: %v", err)
	}
	if err := utils.InitDB(); err != nil {
		log.Fatalf("init db error: %v", err)
	}
	utils.NewKafkaProducer(utils.Config.Kafka.Brokers)
	utils.NewKafkaConsumer(utils.Config.Kafka.Brokers)
}

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

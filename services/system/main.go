package main

import (
	"log"
	"net"

	"micro-system/dao/model/migrate"
	"micro-system/grpcs"
	"micro-system/pb"
	"micro-system/utils"

	"google.golang.org/grpc"
)

func init() {
	if err := utils.InitConfig("system.config.yml", "yml"); err != nil {
		log.Fatalf("init config error: %v", err)
	}
	if err := utils.InitDB(); err != nil {
		log.Fatalf("init db error: %v", err)
	}
	migrate.Migrate(utils.DB)
	utils.NewKafkaProducer(utils.Config.Kafka.Brokers)
	utils.NewKafkaConsumer(utils.Config.Kafka.Brokers)
}

func main() {
	listener, err := net.Listen("tcp", ":50011")
	if err != nil {
		log.Fatalf("micro-system failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	server := grpcs.NewSystemServiceServer()
	pb.RegisterSystemServiceServer(grpcServer, server)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("micro-system failed to serve: %v", err)
	}
}

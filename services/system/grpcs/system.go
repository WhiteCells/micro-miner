package grpcs

import (
	"context"
	"micro-system/pb"
)

type SystemServiceServer struct {
	pb.UnimplementedSystemServiceServer
}

func (s *SystemServiceServer) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginRsp, error) {
	return nil, nil
}

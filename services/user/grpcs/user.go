package grpcs

import (
	"context"
	pb "micro-user/pb/userpb"
)

type UserServiceServer struct {
	pb.UnimplementedUserServiceServer
}

func (s *UserServiceServer) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginRsp, error) {
	return nil, nil
}

func (s *UserServiceServer) AddUser(ctx context.Context, req *pb.AddUserReq) (*pb.AddUserRsp, error) {
	return nil, nil
}

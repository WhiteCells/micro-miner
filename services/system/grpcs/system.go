package grpcs

import (
	"context"
	"micro-system/dao"
	"micro-system/dao/model"
	"micro-system/pb"
	"micro-system/utils"
	"time"
)

type SystemServiceServer struct {
	pb.UnimplementedSystemServiceServer
	userDao *dao.UserDao
}

func NewSystemServiceServer() *SystemServiceServer {
	return &SystemServiceServer{
		userDao: dao.NewUserDao(),
	}
}

func (m *SystemServiceServer) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginRsp, error) {
	user, err := m.userDao.GetUserByEmail(req.Email)
	if err != nil {
		return &pb.LoginRsp{Msg: "user not exits"}, err
	}
	if err := utils.ValidPassword(user.Password, req.Password); err != nil {
		return &pb.LoginRsp{Msg: "password error"}, err
	}
	return &pb.LoginRsp{Msg: "login success"}, nil
}

func (m *SystemServiceServer) Register(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterRsp, error) {
	user := &model.User{
		Name:        req.Name,
		Email:       req.Email,
		Password:    req.Password,
		LastLoginAt: time.Now(),
	}
	if err := m.userDao.AddUser(user); err != nil {
		return &pb.RegisterRsp{Msg: "register error"}, err
	}
	return &pb.RegisterRsp{Msg: "register success"}, nil
}

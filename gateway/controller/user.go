package controller

import (
	"gateway/grpcc"
	pb "gateway/pb"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userGrpcClient pb.UserServiceClient
}

func NewUserController() *UserController {
	return &UserController{
		userGrpcClient: grpcc.NewUserGrpcClient(),
	}
}

func (m *UserController) Login(ctx *gin.Context) {

	ctx.JSON(200, gin.H{})
}

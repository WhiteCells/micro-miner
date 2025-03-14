package controller

import (
	"gateway/common/rsp"
	"gateway/grpcc"
	pb "gateway/pb/userpb"
	"net/http"

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
	var in pb.LoginReq
	if err := ctx.ShouldBind(&in); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, err.Error(), "")
	}
	out, err := m.userGrpcClient.Login(ctx, &in)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), "")
	}
	rsp.LoginSuccess(ctx, http.StatusOK, "", out.GetMsg(), out.GetJwt(), out.GetRole())
}

func (m *UserController) Register(ctx *gin.Context) {
	var in pb.AddUserReq
	if err := ctx.ShouldBind(&in); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, err.Error(), "")
	}
	out, err := m.userGrpcClient.AddUser(ctx, &in)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), "")
	}
	rsp.Success(ctx, http.StatusOK, "", out.GetMsg())
}

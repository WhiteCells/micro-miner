package controller

import (
	"gateway/common/rsp"
	"gateway/grpcc"
	pb "gateway/pb/userpb"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	// userGrpcClient pb.UserServiceClient
}

func NewUserController() *UserController {
	return &UserController{
		// userGrpcClient: grpcc.NewUserGrpcClient(),
	}
}

func (UserController) Login(ctx *gin.Context) {
	var in pb.LoginReq
	if err := ctx.ShouldBind(&in); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, err.Error(), "")
		return
	}
	out, err := grpcc.UserGrpcClient.Login(ctx, &in)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), "")
		return
	}
	rsp.LoginSuccess(ctx, http.StatusOK, "", out.GetMsg(), out.GetJwt(), out.GetRole())
}

func (UserController) Register(ctx *gin.Context) {
	var in pb.AddUserReq
	if err := ctx.ShouldBind(&in); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, err.Error(), "")
		return
	}
	out, err := grpcc.UserGrpcClient.AddUser(ctx, &in)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), "")
		return
	}
	rsp.Success(ctx, http.StatusOK, "", out.GetMsg())
}

package controller

import (
	"gateway/common/rsp"
	"gateway/grpcc"
	"gateway/pb"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SystemController struct {
}

func NewSystemController() *SystemController {
	return &SystemController{}
}

func (SystemController) Login(ctx *gin.Context) {
	var in pb.LoginReq
	if err := ctx.ShouldBind(&in); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, err.Error(), "")
		return
	}
	out, err := grpcc.SystemGrpcClient.Login(ctx, &in)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), "")
		return
	}
	rsp.LoginSuccess(ctx, http.StatusOK, "", out.GetMsg(), "out.GetJwt()", "out.GetRole()")
}

// func (SystemController) Register(ctx *gin.Context) {
// 	var in pb.AddUserReq
// 	if err := ctx.ShouldBind(&in); err != nil {
// 		rsp.Error(ctx, http.StatusBadRequest, err.Error(), "")
// 		return
// 	}
// 	out, err := grpcc.SystemGrpcClient.AddUser(ctx, &in)
// 	if err != nil {
// 		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), "")
// 		return
// 	}
// 	rsp.Success(ctx, http.StatusOK, "", out.GetMsg())
// }

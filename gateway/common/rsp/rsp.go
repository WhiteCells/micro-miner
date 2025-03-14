package rsp

import "github.com/gin-gonic/gin"

func LoginSuccess(ctx *gin.Context, statusCode int, msg string, data interface{}, token string, role string) {
	ctx.JSON(statusCode, gin.H{
		"code":         statusCode,
		"data":         data,
		"msg":          msg,
		"access_token": token,
		"role":         role,
	})
}

func Success(ctx *gin.Context, statusCode int, msg string, data interface{}) {
	ctx.JSON(statusCode, gin.H{
		"code": statusCode,
		"data": data,
		"msg":  msg,
	})
}

func Error(ctx *gin.Context, statusCode int, msg string, data interface{}) {
	ctx.JSON(statusCode, gin.H{
		"code": statusCode,
		"data": data,
		"msg":  msg,
	})
}

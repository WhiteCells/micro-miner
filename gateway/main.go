package main

import (
	"gateway/grpcc"
	"gateway/route"
	"gateway/utils"

	"github.com/gin-gonic/gin"
)

func init() {
	utils.InitConfig("config.yml", "yml")
	utils.InitCache()
	grpcc.InitUserGrpcClient()
}

func main() {
	gin.SetMode(gin.DebugMode)
	ctx := gin.Default()

	route.RegisterRoutes(ctx)
	if err := ctx.Run(); err != nil {
		//
	}
}

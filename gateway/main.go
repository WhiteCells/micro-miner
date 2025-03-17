package main

import (
	"gateway/grpcc"
	"gateway/route"
	"gateway/utils"

	"github.com/gin-gonic/gin"
)

func init() {
	utils.InitConfig("gateway.config.yml", "yml")
	utils.InitCache()
	grpcc.InitSystemGrpcClient()
}

func main() {
	gin.SetMode(gin.DebugMode)
	ctx := gin.Default()

	route.RegisterRoutes(ctx)
	if err := ctx.Run(); err != nil {
		//
	}
}

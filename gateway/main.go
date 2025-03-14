package main

import (
	"log"
	"os"

	"gateway/route"
	"gateway/utils"

	"github.com/gin-gonic/gin"
)

func init() {
	if err := utils.InitConfig("config", "yml"); err != nil {
		//
		log.Fatalf("init config error: %v", err)
		os.Exit(1)
	}

}

func main() {
	gin.SetMode(gin.DebugMode)
	ctx := gin.Default()

	route.RegisterRoutes(ctx)
	ctx.Run()
}

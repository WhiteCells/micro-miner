package main

import (
	"gateway/route"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)
	ctx := gin.Default()

	route.RegisterRoutes(ctx)
	ctx.Run()
}

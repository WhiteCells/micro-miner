package route

import (
	"gateway/controller"

	"github.com/gin-gonic/gin"
)

type UserRoute struct {
	userController *controller.SystemController
}

func NewUserRoute() *UserRoute {
	return &UserRoute{
		userController: controller.NewSystemController(),
	}
}

func (m *UserRoute) Routes(r *gin.RouterGroup) {
	user := r.Group("/user")
	{
		user.POST("/login", m.userController.Login)
		user.POST("/register", m.userController.Register)
	}
}

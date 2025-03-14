package route

import (
	"gateway/controller"

	"github.com/gin-gonic/gin"
)

type UserRoute struct {
	userController *controller.UserController
}

func NewUserRoute() *UserRoute {
	return &UserRoute{
		userController: controller.NewUserController(),
	}
}

func (m *UserRoute) RegisterUserRoutes(r *gin.RouterGroup) {
	user := r.Group("/user")
	{
		user.POST("/login", m.userController.Login)
		user.POST("/register", m.userController.Register)
	}
}

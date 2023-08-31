package routes

import (
	"github.com/MasterAbror/users/api/controller"
	"github.com/MasterAbror/users/environment"
)

// UserRoute -> Route for user module
type UserRoute struct {
	Handler    environment.GinRouter
	Controller controller.UserController
}

// NewUserRoute -> initializes new instance of UserRoute
func NewUserRoute(
	controller controller.UserController,
	handler environment.GinRouter,
) UserRoute {
	return UserRoute{
		Handler:    handler,
		Controller: controller,
	}
}

// Setup -> setups user routes
func (u UserRoute) Setup() {
	user := u.Handler.Gin.Group("/auth")
	{
		user.POST("/register", u.Controller.CreateUser)
		user.POST("/login", u.Controller.LoginUser)
	}
}

package routes

import (
	"github.com/MasterAbror/users/api/controller"
	"github.com/MasterAbror/users/environment"
)

type RoleRoute struct {
	Handler    environment.GinRouter
	Controller controller.RoleController
}

func NewRoleRoute(
	controller controller.RoleController,
	handler environment.GinRouter,
) RoleRoute {
	return RoleRoute{
		Handler:    handler,
		Controller: controller,
	}
}

func (u RoleRoute) Setup() {
	role := u.Handler.Gin.Group("/role")
	{
		role.POST("/create", u.Controller.CreateRole)
	}
}

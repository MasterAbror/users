package routes

import (
	"github.com/MasterAbror/users/api/controller"
	"github.com/MasterAbror/users/environment"
)

type LevelRoute struct {
	Handler    environment.GinRouter
	Controller controller.LevelController
}

func NewLevelRoute(
	controller controller.LevelController,
	handler environment.GinRouter,
) LevelRoute {
	return LevelRoute{
		Handler:    handler,
		Controller: controller,
	}
}

func (u LevelRoute) Setup() {
	level := u.Handler.Gin.Group("/level")
	{
		level.POST("/create", u.Controller.CreateLevel)
	}
}

package routes

import (
	"github.com/MasterAbror/users/api/controller"
	"github.com/MasterAbror/users/environment"
)

// UserRoute -> Route for user module
type GroupRoute struct {
	Handler    environment.GinRouter
	Controller controller.GroupController
	User       controller.UserController
}

// NewGroupRoute -> initializes new instance of GroupRoute
func NewGroupRoute(
	controller controller.GroupController,
	handler environment.GinRouter,
) GroupRoute {
	return GroupRoute{
		Handler:    handler,
		Controller: controller,
	}
}

// Setup -> setups Group routes
func (u GroupRoute) Setup() {
	group := u.Handler.Gin.Group("/group")
	{
		group.GET("/all", u.User.Authorization, u.Controller.All)
		group.POST("/create", u.User.Authorization, u.Controller.CreateGroup)
		group.GET("/read/:id", u.User.Authorization, u.Controller.ReadGroup)
		group.PUT("/update", u.User.Authorization, u.Controller.UpdateGroup)
		group.DELETE("/delete/:id", u.User.Authorization, u.Controller.DeleteGroup)
	}
}

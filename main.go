package main

import (
	"github.com/MasterAbror/users/api/controller"
	"github.com/MasterAbror/users/api/repository"
	"github.com/MasterAbror/users/api/routes"
	"github.com/MasterAbror/users/api/service"
	"github.com/MasterAbror/users/database"
	"github.com/MasterAbror/users/environment"
)

func init() {
	environment.LoadEnv()
}

func main() {

	router := environment.NewGinRouter()
	db := database.NewDatabase()

	// add up these
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)
	userRoute := routes.NewUserRoute(userController, router)
	userRoute.Setup()

	groupRepository := repository.NewGroupRepository(db)
	groupService := service.NewGroupService(groupRepository)
	groupController := controller.NewGroupController(groupService)
	groupRoute := routes.NewGroupRoute(groupController, router)
	groupRoute.Setup()

	levelRepository := repository.NewLevelRepository(db)
	levelService := service.NewLevelService(levelRepository)
	levelController := controller.NewLevelController(levelService)
	levelRoute := routes.NewLevelRoute(levelController, router)
	levelRoute.Setup()

	roleRepository := repository.NewRoleRepository(db)
	roleService := service.NewRoleService(roleRepository)
	roleController := controller.NewRoleController(roleService)
	roleRoute := routes.NewRoleRoute(roleController, router)
	roleRoute.Setup()

	// db.DB.AutoMigrate(&models.User{})
	// db.DB.AutoMigrate(&models.Role{})
	// db.DB.AutoMigrate(&models.Group{})
	// db.DB.AutoMigrate(&models.Level{})

	router.Gin.Run(":1606")
}

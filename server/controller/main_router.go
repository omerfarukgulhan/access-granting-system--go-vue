package controller

import (
	"access-granting/persistence"
	"access-granting/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MainRouter struct {
	authController     *AuthController
	userController     *UserController
	roleController     *RoleController
	userRoleController *UserRoleController
}

func NewRouter(authController *AuthController, userController *UserController, roleController *RoleController, userRoleController *UserRoleController) *MainRouter {
	return &MainRouter{
		authController:     authController,
		userController:     userController,
		roleController:     roleController,
		userRoleController: userRoleController,
	}
}

func (mainRouter *MainRouter) RegisterRoutes(server *gin.Engine) {
	mainRouter.authController.RegisterAuthRoutes(server)
	mainRouter.userController.RegisterUserRoutes(server)
	mainRouter.roleController.RegisterRoleRoutes(server)
	mainRouter.userRoleController.RegisterUserRoleRoutes(server)
}

func InitializeRouter(db *gorm.DB, server *gin.Engine) {
	userRepo := persistence.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := NewUserController(userService)

	authService := service.NewAuthService(userService)
	authController := NewAuthController(authService)

	roleRepo := persistence.NewRoleRepository(db)
	roleService := service.NewRoleService(roleRepo)
	roleController := NewRoleController(roleService)

	userRoleRepo := persistence.NewUserRoleRepository(db)
	userRoleService := service.NewUserRoleService(userRoleRepo, userService, roleService)
	userRoleController := NewUserRoleController(userRoleService)

	mainRouter := NewRouter(authController, userController, roleController, userRoleController)
	mainRouter.RegisterRoutes(server)
}

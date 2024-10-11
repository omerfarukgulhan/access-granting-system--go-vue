package controller

import (
	"access-granting/persistence"
	"access-granting/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MainRouter struct {
	authController *AuthController
	userController *UserController
}

func NewRouter(authController *AuthController, userController *UserController) *MainRouter {
	return &MainRouter{
		authController: authController,
		userController: userController,
	}
}

func (mainRouter *MainRouter) RegisterRoutes(server *gin.Engine) {
	mainRouter.authController.RegisterAuthRoutes(server)
	mainRouter.userController.RegisterUserRoutes(server)
}

func InitializeRouter(db *gorm.DB, server *gin.Engine) {
	userRepo := persistence.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := NewUserController(userService)

	authService := service.NewAuthService(userService)
	authController := NewAuthController(authService)

	mainRouter := NewRouter(authController, userController)
	mainRouter.RegisterRoutes(server)
}

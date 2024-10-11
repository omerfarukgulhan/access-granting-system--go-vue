package controller

import (
	"access-granting/common/util/result"
	"access-granting/controller/messages"
	"access-granting/domain/requests"
	"access-granting/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService service.IAuthService
}

func NewAuthController(authService service.IAuthService) *AuthController {
	return &AuthController{authService: authService}
}

func (authController *AuthController) RegisterAuthRoutes(router *gin.Engine) {
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/register", authController.Register)
		authGroup.POST("/login", authController.Login)
	}
}

func (authController *AuthController) Register(ctx *gin.Context) {
	var newUser requests.UserCreateRequest
	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, result.NewResult(false, messages.InvalidRequestData))
		return
	}

	user, err := authController.authService.Register(newUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, result.NewResult(false, err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, result.NewDataResult(true, messages.SuccessfulRegister, user))
}

func (authController *AuthController) Login(ctx *gin.Context) {
	var signInCredentials requests.SignInCredentials
	if err := ctx.ShouldBindJSON(&signInCredentials); err != nil {
		ctx.JSON(http.StatusBadRequest, result.NewResult(false, messages.InvalidRequestData))
		return
	}

	authResponse, err := authController.authService.Login(signInCredentials)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, result.NewResult(false, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, result.NewDataResult(true, messages.SuccessfulLogin, authResponse))
}

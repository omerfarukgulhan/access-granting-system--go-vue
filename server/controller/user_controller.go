package controller

import (
	"access-granting/common/util"
	"access-granting/common/util/result"
	"access-granting/controller/messages"
	"access-granting/controller/middlewares"
	"access-granting/domain/requests"
	"access-granting/domain/responses"
	"access-granting/service"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService service.IUserService
}

func NewUserController(userService service.IUserService) *UserController {
	return &UserController{userService: userService}
}

func (userController *UserController) RegisterUserRoutes(server *gin.Engine) {
	userGroup := server.Group("/users")
	{
		userGroup.GET("", userController.GetUsers)
		userGroup.GET("/:id", userController.GetUserById)
		userGroup.POST("", userController.AddUser)
		userGroup.PUT("/activate/:token", userController.ActivateUser)
		userGroup.PUT("/:id", middlewares.Authenticate, userController.UpdateUser)
		userGroup.PUT("/update-password", middlewares.Authenticate, userController.UpdatePassword)
		userGroup.PUT("/reset-password", userController.ResetPassword)
		userGroup.PUT("/reset-password/verify/:token", userController.SetPassword)
		userGroup.DELETE("/:id", middlewares.Authenticate, userController.DeleteUser)
	}
}

func (userController *UserController) GetUsers(ctx *gin.Context) {
	users, err := userController.userService.GetUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, result.NewResult(false, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, result.NewDataResult(true, messages.DataFetched, users))
}

func (userController *UserController) GetUserById(ctx *gin.Context) {
	userId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, result.NewResult(false, messages.InvalidId))
		return
	}

	user, err := userController.userService.GetUserById(userId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, result.NewResult(false, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, result.NewDataResult(true, messages.DataFetched, user))
}

func (userController *UserController) AddUser(ctx *gin.Context) {
	var userRequest requests.UserCreateRequest
	if err := bindAndValidateUsers(ctx, &userRequest); err != nil {
		return
	}

	userResponse, err := userController.userService.AddUser(userRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, result.NewResult(false, err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, result.NewDataResult(true, messages.DataAdded, userResponse))
}

func (userController *UserController) ActivateUser(ctx *gin.Context) {
	token := ctx.Param("token")
	activatedUser, err := userController.userService.ActivateUser(token)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, result.NewResult(false, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, result.NewDataResult(true, messages.UserActivated, activatedUser))
}

func (userController *UserController) UpdateUser(ctx *gin.Context) {
	authUserId, err := util.GetUserIdFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, result.NewResult(false, messages.NotAuthorized))
		return
	}

	userId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, result.NewResult(false, messages.InvalidId))
		return
	}

	if !checkAuthorization(ctx, authUserId, userId) {
		return
	}

	var userUpdateRequest requests.UserUpdateRequest
	if err := bindAndValidateUsers(ctx, &userUpdateRequest); err != nil {
		return
	}

	currentUser, err := userController.userService.GetUserById(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, result.NewResult(false, err.Error()))
		return
	}

	profileImagePath, err := handleProfileImage(ctx, &currentUser, userUpdateRequest.ProfileImage)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, result.NewResult(false, messages.ImageError))
		return
	}

	updatedUser, err := userController.userService.UpdateUser(userId, requests.UserUpdateServiceRequest{
		Username:     userUpdateRequest.Username,
		ProfileImage: profileImagePath,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, result.NewResult(false, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, result.NewDataResult(true, messages.DataUpdated, updatedUser))
}

func (userController *UserController) UpdatePassword(ctx *gin.Context) {
	authUserId, err := util.GetUserIdFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, result.NewResult(false, messages.NotAuthorized))
		return
	}

	var passwordUpdate requests.UserPasswordUpdateRequest
	if err := bindAndValidateUsers(ctx, &passwordUpdate); err != nil {
		return
	}

	err = userController.userService.UpdatePassword(authUserId, passwordUpdate)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, result.NewResult(false, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, result.NewResult(true, messages.DataUpdated))
}

func (userController *UserController) ResetPassword(ctx *gin.Context) {
	var passwordReset requests.UserPasswordResetRequest
	if err := bindAndValidateUsers(ctx, &passwordReset); err != nil {
		return
	}

	err := userController.userService.ResetPassword(passwordReset)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, result.NewResult(false, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, result.NewResult(true, messages.PasswordResetEmail))
}

func (userController *UserController) SetPassword(ctx *gin.Context) {
	var passwordSet requests.UserPasswordSetRequest
	if err := bindAndValidateUsers(ctx, &passwordSet); err != nil {
		return
	}

	tokenParam := ctx.Param("token")
	err := userController.userService.SetPassword(tokenParam, passwordSet)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, result.NewResult(false, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, result.NewResult(true, messages.DataUpdated))
}

func (userController *UserController) DeleteUser(ctx *gin.Context) {
	authUserId, err := util.GetUserIdFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, result.NewResult(false, messages.NotAuthorized))
		return
	}

	userId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, result.NewResult(false, messages.InvalidId))
		return
	}

	if !checkAuthorization(ctx, authUserId, userId) {
		return
	}

	currentUser, err := userController.userService.GetUserById(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, result.NewResult(false, err.Error()))
		return
	}

	err = util.DeleteProfileImage(currentUser.ProfileImage)
	if err != nil {
		ctx.JSON(http.StatusNotFound, result.NewResult(false, err.Error()))
		return
	}

	err = userController.userService.DeleteUser(userId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, result.NewResult(false, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, result.NewResult(true, messages.DataDeleted))
}

func bindAndValidateUsers(ctx *gin.Context, req interface{}) error {
	if err := ctx.ShouldBind(req); err != nil {
		ctx.JSON(http.StatusBadRequest, result.NewResult(false, messages.InvalidRequestData))
		return err
	}
	return nil
}

func checkAuthorization(ctx *gin.Context, authUserId, userId int64) bool {
	if authUserId != userId {
		ctx.JSON(http.StatusUnauthorized, result.NewResult(false, messages.NotAuthorized))
		return false
	}
	return true
}

func handleProfileImage(ctx *gin.Context, currentUser *responses.UserWithRolesResponse, newProfileImage *multipart.FileHeader) (string, error) {
	var profileImagePath string
	if newProfileImage != nil {
		path, err := util.SaveProfileImage(ctx, newProfileImage)
		if err != nil {
			return "", err
		}
		if currentUser.ProfileImage != "" && currentUser.ProfileImage != "default.jpg" {
			err := util.DeleteProfileImage(currentUser.ProfileImage)
			if err != nil {
				return "", err
			}
		}
		profileImagePath = path
	}
	return profileImagePath, nil
}

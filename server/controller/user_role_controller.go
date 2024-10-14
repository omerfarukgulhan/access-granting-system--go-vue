package controller

import (
	"access-granting/common/util/result"
	"access-granting/controller/messages"
	"access-granting/controller/middlewares"
	"access-granting/domain/requests"
	"access-granting/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserRoleController struct {
	userRoleService service.IUserRoleService
}

func NewUserRoleController(userRoleService service.IUserRoleService) *UserRoleController {
	return &UserRoleController{userRoleService: userRoleService}
}

func (userRoleController *UserRoleController) RegisterUserRoleRoutes(server *gin.Engine) {
	userRoleGroup := server.Group("/user-roles")
	{
		userRoleGroup.Use(middlewares.Authenticate)
		userRoleGroup.Use(middlewares.Authorize("Admin"))
		userRoleGroup.POST("", userRoleController.AddUserRole)
		userRoleGroup.DELETE("/:userId/:roleId", userRoleController.DeleteUserRole)
	}
}

func (userRoleController *UserRoleController) AddUserRole(ctx *gin.Context) {
	var userRoleRequest requests.UserRoleCreateRequest
	if err := bindAndValidateUserRole(ctx, &userRoleRequest); err != nil {
		return
	}

	newUserRole, err := userRoleController.userRoleService.AddUserRole(userRoleRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, result.NewResult(false, err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, result.NewDataResult(true, messages.DataAdded, newUserRole))
}

func (userRoleController *UserRoleController) DeleteUserRole(ctx *gin.Context) {
	userId, err := strconv.ParseInt(ctx.Param("userId"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, result.NewResult(false, messages.InvalidId))
		return
	}

	roleId, err := strconv.ParseInt(ctx.Param("roleId"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, result.NewResult(false, messages.InvalidId))
		return
	}

	err = userRoleController.userRoleService.DeleteUserRole(userId, roleId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, result.NewResult(false, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, result.NewResult(true, messages.DataDeleted))
}

func bindAndValidateUserRole(ctx *gin.Context, req interface{}) error {
	if err := ctx.ShouldBind(req); err != nil {
		ctx.JSON(http.StatusBadRequest, result.NewResult(false, messages.InvalidRequestData))
		return err
	}
	return nil
}

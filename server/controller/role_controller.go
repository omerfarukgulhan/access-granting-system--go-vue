package controller

import (
	"access-granting/common/util/result"
	"access-granting/controller/messages"
	"access-granting/domain/requests"
	"access-granting/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RoleController struct {
	roleService service.IRoleService
}

func NewRoleController(roleService service.IRoleService) *RoleController {
	return &RoleController{roleService: roleService}
}

func (roleController *RoleController) RegisterRoleRoutes(server *gin.Engine) {
	roleGroup := server.Group("/roles")
	{
		roleGroup.GET("", roleController.GetRoles)
		roleGroup.GET("/:id", roleController.GetRoleById)
		roleGroup.POST("", roleController.AddRole)
		roleGroup.PUT("/:id", roleController.UpdateRole)
		roleGroup.DELETE("/:id", roleController.DeleteRole)
	}
}

func (roleController *RoleController) GetRoles(ctx *gin.Context) {
	roles, err := roleController.roleService.GetRoles()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, result.NewResult(false, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, result.NewDataResult(true, messages.DataFetched, roles))
}

func (roleController *RoleController) GetRoleById(ctx *gin.Context) {
	roleId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, result.NewResult(false, "Invalid role ID"))
		return
	}

	role, err := roleController.roleService.GetRoleById(roleId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, result.NewResult(false, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, result.NewDataResult(true, messages.DataFetched, role))
}

func (roleController *RoleController) AddRole(ctx *gin.Context) {
	var roleRequest requests.RoleCreateRequest
	if err := bindAndValidateRoles(ctx, &roleRequest); err != nil {
		return
	}

	roleResponse, err := roleController.roleService.AddRole(roleRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, result.NewResult(false, err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, result.NewDataResult(true, messages.DataAdded, roleResponse))
}

func (roleController *RoleController) UpdateRole(ctx *gin.Context) {
	roleId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, result.NewResult(false, "Invalid role ID"))
		return
	}

	var roleRequest requests.RoleUpdateRequest
	if err := bindAndValidateRoles(ctx, &roleRequest); err != nil {
		return
	}

	roleResponse, err := roleController.roleService.UpdateRole(roleId, roleRequest)
	if err != nil {
		ctx.JSON(http.StatusNotFound, result.NewResult(false, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, result.NewDataResult(true, messages.DataUpdated, roleResponse))
}

func (roleController *RoleController) DeleteRole(ctx *gin.Context) {
	roleId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, result.NewResult(false, "Invalid role ID"))
		return
	}

	if err := roleController.roleService.DeleteRole(roleId); err != nil {
		ctx.JSON(http.StatusNotFound, result.NewResult(false, err.Error()))
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func bindAndValidateRoles(ctx *gin.Context, req interface{}) error {
	if err := ctx.ShouldBind(req); err != nil {
		ctx.JSON(http.StatusBadRequest, result.NewResult(false, messages.InvalidRequestData))
		return err
	}
	return nil
}

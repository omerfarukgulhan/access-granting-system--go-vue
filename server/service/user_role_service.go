package service

import (
	"access-granting/domain/entities"
	"access-granting/domain/requests"
	"access-granting/domain/responses"
	"access-granting/persistence"
)

type IUserRoleService interface {
	AddUserRole(userRole requests.UserRoleCreateRequest) (responses.UserRoleResponse, error)
	DeleteUserRole(userId, roleId int64) error
}

type UserRoleService struct {
	userRoleRepository persistence.IUserRoleRepository
	userService        IUserService
	roleService        IRoleService
}

func NewUserRoleService(userRoleRepository persistence.IUserRoleRepository, userService IUserService, roleService IRoleService) IUserRoleService {
	return &UserRoleService{
		userRoleRepository: userRoleRepository,
		userService:        userService,
		roleService:        roleService,
	}
}

func (userRoleService *UserRoleService) AddUserRole(userRole requests.UserRoleCreateRequest) (responses.UserRoleResponse, error) {
	_, err := userRoleService.userService.GetUserById(userRole.UserId)
	if err != nil {
		return responses.UserRoleResponse{}, err
	}

	_, err = userRoleService.roleService.GetRoleById(userRole.RoleId)
	if err != nil {
		return responses.UserRoleResponse{}, err
	}

	newUserRole := entities.UserRole{
		RoleId: userRole.RoleId,
		UserId: userRole.UserId,
	}
	_, err = userRoleService.userRoleRepository.AddUserRole(newUserRole)
	if err != nil {
		return responses.UserRoleResponse{}, err
	}

	return convertUserRoleToResponse(newUserRole), nil
}

func (userRoleService *UserRoleService) DeleteUserRole(userId, roleId int64) error {
	if _, err := userRoleService.userService.GetUserById(userId); err != nil {
		return err
	}
	if _, err := userRoleService.roleService.GetRoleById(roleId); err != nil {
		return err
	}

	return userRoleService.userRoleRepository.DeleteUserRole(userId, roleId)
}

func convertUserRoleToResponse(userRole entities.UserRole) responses.UserRoleResponse {
	return responses.NewUserRoleResponse(userRole)
}

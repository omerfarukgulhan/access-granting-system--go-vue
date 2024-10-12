package service

import (
	"access-granting/common/util/id"
	"access-granting/domain/entities"
	"access-granting/domain/requests"
	"access-granting/domain/responses"
	"access-granting/persistence"
	"errors"
	"strings"
)

type IRoleService interface {
	GetRoles() ([]responses.RoleResponse, error)
	GetRoleById(roleId int64) (responses.RoleResponse, error)
	AddRole(role requests.RoleCreateRequest) (responses.RoleResponse, error)
	UpdateRole(roleId int64, role requests.RoleUpdateRequest) (responses.RoleResponse, error)
	DeleteRole(roleId int64) error
}

type RoleService struct {
	roleRepository persistence.IRoleRepository
}

func NewRoleService(roleRepository persistence.IRoleRepository) IRoleService {
	return &RoleService{roleRepository: roleRepository}
}

func (roleService *RoleService) GetRoles() ([]responses.RoleResponse, error) {
	roles, err := roleService.roleRepository.GetRoles()
	if err != nil {
		return nil, err
	}
	return convertRolesToResponses(roles), nil
}

func (roleService *RoleService) GetRoleById(roleId int64) (responses.RoleResponse, error) {
	role, err := roleService.roleRepository.GetRoleById(roleId)
	if err != nil {
		return responses.RoleResponse{}, err
	}
	return convertRoleToResponse(role), nil
}

func (roleService *RoleService) AddRole(role requests.RoleCreateRequest) (responses.RoleResponse, error) {
	err := validateRoleName(role.Name)
	if err != nil {
		return responses.RoleResponse{}, err
	}

	uniqueId, err := id.GetUniqueId()
	if err != nil {
		return responses.RoleResponse{}, err
	}

	newRole := entities.Role{
		Id:   uniqueId,
		Name: role.Name,
	}

	createdRole, err := roleService.roleRepository.AddRole(newRole)
	if err != nil {
		return responses.RoleResponse{}, err
	}

	return convertRoleToResponse(createdRole), nil
}

func (roleService *RoleService) UpdateRole(roleId int64, role requests.RoleUpdateRequest) (responses.RoleResponse, error) {
	err := validateRoleName(role.Name)
	if err != nil {
		return responses.RoleResponse{}, err
	}

	updatedRole := entities.Role{
		Name: role.Name,
	}

	updatedRole, err = roleService.roleRepository.UpdateRole(roleId, updatedRole)
	if err != nil {
		return responses.RoleResponse{}, err
	}

	return convertRoleToResponse(updatedRole), nil
}

func (roleService *RoleService) DeleteRole(roleId int64) error {
	return roleService.roleRepository.DeleteRole(roleId)
}

func convertRolesToResponses(roles []entities.Role) []responses.RoleResponse {
	roleResponses := make([]responses.RoleResponse, len(roles))
	for i, role := range roles {
		roleResponses[i] = responses.NewRoleResponse(role)
	}
	return roleResponses
}

func convertRoleToResponse(role entities.Role) responses.RoleResponse {
	return responses.NewRoleResponse(role)
}

func validateRoleName(roleName string) error {
	if strings.TrimSpace(roleName) == "" {
		return errors.New("role name cannot be empty")
	}
	if len(roleName) < 4 {
		return errors.New("role name must be at least 4 characters long")
	}
	return nil
}

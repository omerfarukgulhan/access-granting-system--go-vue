package responses

import "access-granting/domain/entities"

type UserRoleResponse struct {
	UserId int64 `json:"userId"`
	RoleId int64 `json:"roleId"`
}

func NewUserRoleResponse(userRole entities.UserRole) UserRoleResponse {
	return UserRoleResponse{
		UserId: userRole.UserId,
		RoleId: userRole.RoleId,
	}
}

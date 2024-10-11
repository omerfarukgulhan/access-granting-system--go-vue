package responses

import "access-granting/domain/entities"

type RoleResponse struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func NewRoleResponse(role entities.Role) RoleResponse {
	return RoleResponse{
		Id:   role.Id,
		Name: role.Name,
	}
}

package responses

import "access-granting/domain/entities"

type RoleWithUsersResponse struct {
	Id    int64          `json:"id"`
	Name  string         `json:"name"`
	Users []UserResponse `json:"users"`
}

func NewRoleWithUsersResponse(role entities.Role) RoleWithUsersResponse {
	return RoleWithUsersResponse{
		Id:    role.Id,
		Name:  role.Name,
		Users: convertUsersToResponses(role.Users),
	}
}

func convertUsersToResponses(users []entities.User) []UserResponse {
	userResponses := make([]UserResponse, len(users))
	for i, user := range users {
		userResponses[i] = NewUserResponse(user)
	}
	return userResponses
}

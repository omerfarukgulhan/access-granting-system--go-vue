package responses

import "access-granting/domain/entities"

type UserResponse struct {
	Id           int64          `json:"id"`
	Username     string         `json:"username"`
	Email        string         `json:"email"`
	ProfileImage string         `json:"profileImage"`
	Roles        []RoleResponse `json:"roles"`
}

func NewUserResponse(user entities.User) UserResponse {
	return UserResponse{
		Id:           user.Id,
		Username:     user.Username,
		Email:        user.Email,
		ProfileImage: user.ProfileImage,
		Roles:        convertRolesToResponses((user.Roles)),
	}
}

func convertRolesToResponses(roles []entities.Role) []RoleResponse {
	roleResponses := make([]RoleResponse, len(roles))
	for i, role := range roles {
		roleResponses[i] = NewRoleResponse(role)
	}
	return roleResponses
}

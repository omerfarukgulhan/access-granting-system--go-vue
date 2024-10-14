package responses

import "access-granting/domain/entities"

type AuthResponse struct {
	UserResponse UserWithRolesResponse `json:"user"`
	Token        string                `json:"token"`
	Prefix       string                `json:"prefix"`
}

func NewAuthResponseUserEntity(user entities.User, token string) AuthResponse {
	return AuthResponse{
		UserResponse: NewUserWithRolesResponse(user),
		Token:        token,
		Prefix:       "Bearer",
	}
}

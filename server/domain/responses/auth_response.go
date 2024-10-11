package responses

import "access-granting/domain/entities"

type AuthResponse struct {
	UserResponse UserResponse `json:"user"`
	Token        string       `json:"token"`
	Prefix       string       `json:"prefix"`
}

func NewAuthResponseUserEntity(user entities.User, token string) AuthResponse {
	return AuthResponse{
		UserResponse: NewUserResponse(user),
		Token:        token,
		Prefix:       "Bearer",
	}
}

func NewAuthResponseFromUserResponse(user UserResponse, token string) AuthResponse {
	return AuthResponse{
		UserResponse: user,
		Token:        token,
		Prefix:       "Bearer",
	}
}

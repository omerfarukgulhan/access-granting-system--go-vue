package service

import (
	"access-granting/common/security"
	"access-granting/domain/requests"
	"access-granting/domain/responses"
	"errors"
)

type IAuthService interface {
	Register(user requests.UserCreateRequest) (responses.UserResponse, error)
	Login(signInCredentials requests.SignInCredentials) (responses.AuthResponse, error)
}

type AuthService struct {
	userService IUserService
}

func NewAuthService(userService IUserService) IAuthService {
	return &AuthService{userService: userService}
}

func (authService *AuthService) Register(userCreate requests.UserCreateRequest) (responses.UserResponse, error) {
	user, err := authService.userService.AddUser(userCreate)
	if err != nil {
		return responses.UserResponse{}, err
	}
	return user, nil
}

func (authService *AuthService) Login(signInCredentials requests.SignInCredentials) (responses.AuthResponse, error) {
	user, err := authService.userService.GetUserByEmail(signInCredentials.Email)
	if err != nil {
		return responses.AuthResponse{}, err
	}

	if !user.IsActive {
		return responses.AuthResponse{}, errors.New("user is not activated")
	}

	if !security.CheckPasswordHash(signInCredentials.Password, user.Password) {
		return responses.AuthResponse{}, errors.New("incorrect password")
	}

	var roles []string
	for _, role := range user.Roles {
		roles = append(roles, role.Name)
	}

	token, err := security.GenerateToken(user.Id, user.Email, roles)
	if err != nil {
		return responses.AuthResponse{}, err
	}

	return responses.NewAuthResponseUserEntity(user, token), nil
}

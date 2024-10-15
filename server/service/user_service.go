package service

import (
	"access-granting/common/security"
	"access-granting/domain/entities"
	"access-granting/domain/requests"
	"access-granting/domain/responses"
	"access-granting/persistence"
	"errors"
	"regexp"
	"strings"

	"github.com/google/uuid"
)

type IUserService interface {
	GetUsers() ([]responses.UserResponse, error)
	GetUserById(userId int64) (responses.UserWithRolesResponse, error)
	GetUserByEmail(email string) (entities.User, error)
	AddUser(user requests.UserCreateRequest) (responses.UserResponse, error)
	ActivateUser(token string) (responses.UserResponse, error)
	UpdateUser(userId int64, user requests.UserUpdateServiceRequest) (responses.UserWithRolesResponse, error)
	UpdatePassword(userId int64, passwordUpdate requests.UserPasswordUpdateRequest) error
	ResetPassword(passwordReset requests.UserPasswordResetRequest) error
	SetPassword(token string, passwordSet requests.UserPasswordSetRequest) error
	DeleteUser(userId int64) error
}

type UserService struct {
	userRepository persistence.IUserRepository
	hashCost       int
}

func NewUserService(userRepository persistence.IUserRepository) IUserService {
	return &UserService{userRepository: userRepository, hashCost: 10}
}

func (userService *UserService) GetUsers() ([]responses.UserResponse, error) {
	users, err := userService.userRepository.GetUsers()
	if err != nil {
		return nil, err
	}
	return convertUsersToResponses(users), nil
}

func (userService *UserService) GetUserById(userId int64) (responses.UserWithRolesResponse, error) {
	user, err := userService.userRepository.GetUserById(userId)
	if err != nil {
		return responses.UserWithRolesResponse{}, err
	}
	return convertUserWithRolesToResponse(user), nil
}

func (userService *UserService) GetUserByEmail(email string) (entities.User, error) {
	user, err := userService.userRepository.GetUserByEmail(email)
	if err != nil {
		return entities.User{}, err
	}
	return user, nil
}

func (userService *UserService) AddUser(user requests.UserCreateRequest) (responses.UserResponse, error) {
	validationError := validateUserCreate(user)
	if validationError != nil {
		return responses.UserResponse{}, validationError
	}

	hashedPassword, err := security.HashPassword(user.Password, userService.hashCost)
	if err != nil {
		return responses.UserResponse{}, err
	}

	activationToken := uuid.New().String()
	newUser := entities.User{
		Username:        user.Username,
		Email:           user.Email,
		Password:        hashedPassword,
		ActivationToken: activationToken,
	}

	createdUser, err := userService.userRepository.AddUser(newUser)
	if err != nil {
		return responses.UserResponse{}, err
	}

	err = SendActivationEmail(user.Email, activationToken)
	if err != nil {
		return responses.UserResponse{}, err
	}

	return convertUserToResponse(createdUser), nil
}

func (userService *UserService) ActivateUser(token string) (responses.UserResponse, error) {
	if token == "" {
		return responses.UserResponse{}, errors.New("activation token cannot be empty")
	}

	user, err := userService.userRepository.GetUserByActivationToken(token)
	if err != nil {
		return responses.UserResponse{}, err
	}

	if user.IsActive {
		return responses.UserResponse{}, errors.New("user is already activated")
	}

	user.IsActive = true
	user.ActivationToken = " "
	updatedUser, err := userService.userRepository.UpdateUser(user.Id, user)
	if err != nil {
		return responses.UserResponse{}, err
	}

	return convertUserToResponse(updatedUser), nil
}

func (userService *UserService) UpdateUser(userId int64, user requests.UserUpdateServiceRequest) (responses.UserWithRolesResponse, error) {
	err := validateUserUpdate(user)
	if err != nil {
		return responses.UserWithRolesResponse{}, err
	}

	existingUser, err := userService.userRepository.GetUserById(userId)
	if err != nil {
		return responses.UserWithRolesResponse{}, err
	}

	existingUser.Username = user.Username
	existingUser.ProfileImage = user.ProfileImage
	updatedUser, err := userService.userRepository.UpdateUser(userId, existingUser)
	if err != nil {
		return responses.UserWithRolesResponse{}, err
	}

	return convertUserWithRolesToResponse(updatedUser), nil
}

func (userService *UserService) UpdatePassword(userId int64, passwordUpdate requests.UserPasswordUpdateRequest) error {
	err := validatePasswordUpdate(passwordUpdate.OldPassword)
	if err != nil {
		return err
	}

	existingUser, err := userService.userRepository.GetUserById(userId)
	if err != nil {
		return err
	}

	isOldPasswordCorrect := security.CheckPasswordHash(passwordUpdate.OldPassword, existingUser.Password)
	if !isOldPasswordCorrect {
		return errors.New("old password is incorrect")
	}

	newPassword, err := security.HashPassword(passwordUpdate.NewPassword, userService.hashCost)
	if err != nil {
		return err
	}

	existingUser.Password = newPassword
	_, err = userService.userRepository.UpdateUser(userId, existingUser)

	return err
}

func (userService *UserService) ResetPassword(passwordReset requests.UserPasswordResetRequest) error {
	resetToken := uuid.New().String()
	user, err := userService.userRepository.GetUserByEmail(passwordReset.Email)
	if err != nil {
		return err
	}

	user.PasswordResetToken = resetToken
	_, err = userService.userRepository.UpdateUser(user.Id, user)
	if err != nil {
		return err
	}

	return SendPasswordResetEmail(user.Email, resetToken)
}

func (userService *UserService) SetPassword(token string, passwordSet requests.UserPasswordSetRequest) error {
	err := validatePasswordUpdate(passwordSet.Password)
	if err != nil {
		return err
	}

	if token == "" {
		return errors.New("password reset token cannot be empty")
	}

	user, err := userService.userRepository.GetUserByPasswordResetToken(token)
	if err != nil {
		return err
	}

	hashedPassword, err := security.HashPassword(passwordSet.Password, userService.hashCost)
	if err != nil {
		return err
	}

	user.Password = hashedPassword
	user.PasswordResetToken = " "
	_, err = userService.userRepository.UpdateUser(user.Id, user)

	return err
}

func (userService *UserService) DeleteUser(userId int64) error {
	return userService.userRepository.DeleteUser(userId)
}

func convertUsersToResponses(users []entities.User) []responses.UserResponse {
	userResponses := make([]responses.UserResponse, len(users))
	for i, user := range users {
		userResponses[i] = responses.NewUserResponse(user)
	}
	return userResponses
}

func convertUserToResponse(user entities.User) responses.UserResponse {
	return responses.NewUserResponse(user)
}

func convertUserWithRolesToResponse(user entities.User) responses.UserWithRolesResponse {
	return responses.NewUserWithRolesResponse(user)
}

func validateUserCreate(req requests.UserCreateRequest) error {
	if err := validateUsername(req.Username); err != nil {
		return err
	}
	if err := validateEmail(req.Email); err != nil {
		return err
	}
	if err := validatePassword(req.Password); err != nil {
		return err
	}
	return nil
}

func validateUserUpdate(req requests.UserUpdateServiceRequest) error {
	if err := validateUsername(req.Username); err != nil {
		return err
	}
	return nil
}

func validatePasswordUpdate(oldPassword string) error {
	if err := validatePassword(oldPassword); err != nil {
		return err
	}
	return nil
}

func validateUsername(username string) error {
	if strings.TrimSpace(username) == "" {
		return errors.New("username cannot be empty")
	}
	if len(username) < 4 {
		return errors.New("username must be at least 4 characters long")
	}
	return nil
}

func validateEmail(email string) error {
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	if !re.MatchString(email) {
		return errors.New("invalid email format")
	}
	return nil
}

func validatePassword(password string) error {
	if len(password) < 6 {
		return errors.New("password must be at least 6 characters long")
	}

	hasLetter := false
	hasNumber := false
	for _, char := range password {
		switch {
		case 'a' <= char && char <= 'z', 'A' <= char && char <= 'Z':
			hasLetter = true
		case '0' <= char && char <= '9':
			hasNumber = true
		}
	}

	if !hasLetter || !hasNumber {
		return errors.New("password must contain at least one letter and one number")
	}

	return nil
}

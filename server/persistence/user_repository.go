package persistence

import (
	"access-granting/domain/entities"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type IUserRepository interface {
	GetUsers() ([]entities.User, error)
	GetUserById(userId int64) (entities.User, error)
	GetUserByEmail(email string) (entities.User, error)
	GetUserByActivationToken(token string) (entities.User, error)
	GetUserByPasswordResetToken(token string) (entities.User, error)
	AddUser(user entities.User) (entities.User, error)
	UpdateUser(userId int64, user entities.User) (entities.User, error)
	DeleteUser(userId int64) error
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{db: db}
}

func (userRepository *UserRepository) GetUsers() ([]entities.User, error) {
	var users []entities.User
	result := userRepository.db.Where("is_active = ?", true).Find(&users)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrInvalidDB) {
			return nil, fmt.Errorf("database connection error: %w", result.Error)
		}
		return nil, fmt.Errorf("failed to retrieve users")
	}
	if len(users) == 0 {
		return nil, fmt.Errorf("no active users found")
	}
	return users, nil
}

func (userRepository *UserRepository) GetUserById(userId int64) (entities.User, error) {
	var user entities.User
	err := userRepository.db.Preload("Roles").First(&user, userId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.User{}, fmt.Errorf("user with id %d not found", userId)
		}
		if errors.Is(err, gorm.ErrInvalidDB) {
			return entities.User{}, fmt.Errorf("database connection error: %w", err)
		}
		return entities.User{}, fmt.Errorf("failed to retrieve user by id %d", userId)
	}
	return user, nil
}

func (userRepository *UserRepository) GetUserByEmail(email string) (entities.User, error) {
	return userRepository.getUserByField("email", email)
}

func (userRepository *UserRepository) GetUserByActivationToken(token string) (entities.User, error) {
	return userRepository.getUserByField("activation_token", token)
}

func (userRepository *UserRepository) GetUserByPasswordResetToken(token string) (entities.User, error) {
	return userRepository.getUserByField("password_reset_token", token)
}

func (userRepository *UserRepository) getUserByField(field, value string) (entities.User, error) {
	var user entities.User
	err := userRepository.db.Preload("Roles").Where(field+" = ?", value).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.User{}, fmt.Errorf("user with %s %s not found", field, value)
		}
		if errors.Is(err, gorm.ErrInvalidDB) {
			return entities.User{}, fmt.Errorf("database connection error: %w", err)
		}
		return entities.User{}, fmt.Errorf("failed to retrieve user by %s %s", field, value)
	}
	return user, nil
}

func (userRepository *UserRepository) AddUser(user entities.User) (entities.User, error) {
	err := userRepository.db.Create(&user).Error
	if err != nil {
		if isUniqueConstraintError(err, "email") {
			return entities.User{}, fmt.Errorf("email %s is already in use", user.Email)
		}
		if isUniqueConstraintError(err, "username") {
			return entities.User{}, fmt.Errorf("username %s is already taken", user.Username)
		}
		return entities.User{}, fmt.Errorf("failed to add user")
	}
	return user, nil
}

func (userRepository *UserRepository) UpdateUser(userId int64, user entities.User) (entities.User, error) {
	err := userRepository.db.Model(&user).Where("id = ?", userId).Updates(user).Error
	if err != nil {
		if isUniqueConstraintError(err, "email") {
			return entities.User{}, fmt.Errorf("email %s is already in use", user.Email)
		}
		if isUniqueConstraintError(err, "username") {
			return entities.User{}, fmt.Errorf("username %s is already taken", user.Username)
		}
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.User{}, fmt.Errorf("user with id %d not found", userId)
		}
		return entities.User{}, fmt.Errorf("failed to update user with id %d", userId)
	}
	return user, nil
}

func (userRepository *UserRepository) DeleteUser(userId int64) error {
	result := userRepository.db.Delete(&entities.User{}, userId)
	if result.Error != nil {
		return fmt.Errorf("failed to delete user with id %d", userId)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("user with id %d not found", userId)
	}
	return nil
}

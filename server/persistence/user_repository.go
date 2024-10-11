package persistence

import (
	"access-granting/domain/entities"
	"errors"

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
	result := userRepository.db.Preload("Roles").Where("is_active = ?", true).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (userRepository *UserRepository) GetUserById(userId int64) (entities.User, error) {
	var user entities.User
	err := userRepository.db.Preload("Roles").First(&user, userId).Error
	if err != nil {
		return entities.User{}, err
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
		return entities.User{}, err
	}
	return user, nil
}

func (userRepository *UserRepository) AddUser(user entities.User) (entities.User, error) {
	err := userRepository.db.Create(&user).Error
	if err != nil {
		return entities.User{}, err
	}
	return user, nil
}

func (userRepository *UserRepository) UpdateUser(userId int64, user entities.User) (entities.User, error) {
	err := userRepository.db.Model(&user).Where("id = ?", userId).Updates(user).Error
	if err != nil {
		return entities.User{}, err
	}
	return user, nil
}

func (r *UserRepository) DeleteUser(userId int64) error {
	result := r.db.Delete(&entities.User{}, userId)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("user not found")
	}
	return nil
}

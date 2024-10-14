package persistence

import (
	"access-granting/domain/entities"
	"errors"

	"gorm.io/gorm"
)

type IUserRoleRepository interface {
	AddUserRole(userRole entities.UserRole) (entities.UserRole, error)
	DeleteUserRole(userId, roleId int64) error
}

type UserRoleRepository struct {
	db *gorm.DB
}

func NewUserRoleRepository(db *gorm.DB) IUserRoleRepository {
	return &UserRoleRepository{db: db}
}

func (userRoleRepository *UserRoleRepository) AddUserRole(userRole entities.UserRole) (entities.UserRole, error) {
	if err := userRoleRepository.db.Create(&userRole).Error; err != nil {
		return entities.UserRole{}, err
	}
	return userRole, nil
}

func (userRoleRepository *UserRoleRepository) DeleteUserRole(userId, roleId int64) error {
	result := userRoleRepository.db.Where("user_id = ? AND role_id = ?", userId, roleId).Delete(&entities.UserRole{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("user role not found")
	}
	return nil
}

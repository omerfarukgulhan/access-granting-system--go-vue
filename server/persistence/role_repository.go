package persistence

import (
	"access-granting/domain/entities"
	"errors"

	"gorm.io/gorm"
)

type IRoleRepository interface {
	GetRoles() ([]entities.Role, error)
	GetRoleById(roleId int64) (entities.Role, error)
	AddRole(role entities.Role) (entities.Role, error)
	UpdateRole(roleId int64, role entities.Role) (entities.Role, error)
	DeleteRole(roleId int64) error
}

type RoleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) IRoleRepository {
	return &RoleRepository{db: db}
}

func (roleRepository *RoleRepository) GetRoles() ([]entities.Role, error) {
	var roles []entities.Role
	if err := roleRepository.db.Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

func (roleRepository *RoleRepository) GetRoleById(roleId int64) (entities.Role, error) {
	var role entities.Role
	if err := roleRepository.db.First(&role, roleId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return role, errors.New("role not found")
		}
		return role, err
	}
	return role, nil
}

func (roleRepository *RoleRepository) AddRole(role entities.Role) (entities.Role, error) {
	if err := roleRepository.db.Create(&role).Error; err != nil {
		return role, err
	}
	return role, nil
}

func (roleRepository *RoleRepository) UpdateRole(roleId int64, role entities.Role) (entities.Role, error) {
	var existingRole entities.Role
	if err := roleRepository.db.First(&existingRole, roleId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return existingRole, errors.New("role not found")
		}
		return existingRole, err
	}

	if err := roleRepository.db.Model(&existingRole).Updates(role).Error; err != nil {
		return existingRole, err
	}

	return existingRole, nil
}

func (roleRepository *RoleRepository) DeleteRole(roleId int64) error {
	var role entities.Role
	if err := roleRepository.db.First(&role, roleId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("role not found")
		}
		return err
	}

	if err := roleRepository.db.Delete(&role).Error; err != nil {
		return err
	}

	return nil
}

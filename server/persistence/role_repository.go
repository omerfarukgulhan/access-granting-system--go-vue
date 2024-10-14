package persistence

import (
	"access-granting/domain/entities"
	"errors"
	"fmt"

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
		return nil, fmt.Errorf("failed to retrieve roles: %w", err)
	}
	return roles, nil
}

func (roleRepository *RoleRepository) GetRoleById(roleId int64) (entities.Role, error) {
	var role entities.Role
	if err := roleRepository.db.Preload("Users").First(&role, roleId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return role, fmt.Errorf("role with id %d not found", roleId)
		}
		return role, fmt.Errorf("failed to retrieve role with id %d", roleId)
	}
	return role, nil
}

func (roleRepository *RoleRepository) AddRole(role entities.Role) (entities.Role, error) {
	if err := roleRepository.db.Create(&role).Error; err != nil {
		if isUniqueConstraintError(err, "name") {
			return role, fmt.Errorf("role name '%s' already exists", role.Name)
		}
		return role, fmt.Errorf("failed to add role")
	}
	return role, nil
}

func (roleRepository *RoleRepository) UpdateRole(roleId int64, role entities.Role) (entities.Role, error) {
	err := roleRepository.db.Model(&role).Where("id = ?", roleId).Updates(role).Error
	if err != nil {
		if isUniqueConstraintError(err, "name") {
			return entities.Role{}, fmt.Errorf("name %s is already in use", role.Name)
		}
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.Role{}, fmt.Errorf("role with id %d not found", roleId)
		}
		return entities.Role{}, fmt.Errorf("failed to update role with id %d", roleId)
	}
	var updatedRole entities.Role
	err = roleRepository.db.Preload("Users").Where("id = ?", roleId).First(&updatedRole).Error
	if err != nil {
		return entities.Role{}, fmt.Errorf("failed to fetch updated role with related information")
	}

	return updatedRole, nil
}

func (roleRepository *RoleRepository) DeleteRole(roleId int64) error {
	result := roleRepository.db.Delete(&entities.Role{}, roleId)
	if result.Error != nil {
		return fmt.Errorf("failed to delete role with id %d", roleId)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("role with id %d not found", roleId)
	}
	return nil
}

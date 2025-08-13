package services

import (
	"pixel/app/models"

	"github.com/google/uuid"
	"github.com/goravel/framework/facades"
)

type RoleService struct{}

func NewRoleService() *RoleService {
	return &RoleService{}
}

func (s *RoleService) Create(role models.Role) error {
	return facades.Orm().Query().Create(&role)
}

func (s *RoleService) GetAll() ([]models.Role, error) {
	var roles []models.Role
	err := facades.Orm().Query().Find(&roles)
	return roles, err
}

func (s *RoleService) GetByID(id uuid.UUID) (*models.Role, error) {
	var role models.Role
	err := facades.Orm().Query().Where("id", id).First(&role)
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (s *RoleService) Update(id uuid.UUID, updated models.Role) error {
	var role models.Role
	if err := facades.Orm().Query().Where("id", id).First(&role); err != nil {
		return err
	}
	role.Role = updated.Role
	role.Description = updated.Description
	role.IsActive = updated.IsActive
	return facades.Orm().Query().Save(&role)
}

func (s *RoleService) Delete(id uuid.UUID) error {
	var role models.Role
	if err := facades.Orm().Query().Where("id", id).First(&role); err != nil {
		return err
	}
	_, err := facades.Orm().Query().Delete(&role)
	return err
}

func (s *RoleService) GetByRoleType(roleType models.RoleType) (*models.Role, error) {
	var role models.Role
	err := facades.Orm().Query().Where("role = ?", roleType).First(&role)
	return &role, err
}

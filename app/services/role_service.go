package services

import (
	"pixel/app/models"
	"slices"

	"github.com/google/uuid"
	"github.com/goravel/framework/contracts/database/orm"
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

func (s *RoleService) AssignRoleToUser(user *models.User, roleName models.RoleType) error {
	var role models.Role
	err := facades.Orm().Query().
		Where("role = ?", roleName).
		First(&role)

	if err != nil {
		return err
	}

	return facades.Orm().Query().
		Model(user).
		Association("Roles").
		Append(&role)
}

func (s *RoleService) AssignRoleToUserTx(query orm.Query, user *models.User, roleName models.RoleType) error {
	var role models.Role
	err := query.Where("role = ?", roleName).First(&role)

	if err != nil {
		return err
	}

	return query.Model(user).Association("Roles").Append(&role)
}

func (s *RoleService) HasRoleByUserID(userID uuid.UUID, role models.RoleType) (bool, error) {
	var user models.User

	err := facades.Orm().Query().
		With("Roles").
		Where("id = ?", userID).
		First(&user)

	if err != nil {
		return false, err
	}

	for _, userRole := range user.Roles {
		if userRole.Role == role && userRole.IsActive {
			return true, nil
		}
	}

	return false, nil
}

func (s *RoleService) HasRolesByUserID(userID uuid.UUID, roles []models.RoleType) (bool, error) {
	var user models.User

	err := facades.Orm().Query().
		With("Roles").
		Where("id = ?", userID).
		First(&user)

	if err != nil {
		return false, err
	}

	for _, userRole := range user.Roles {
		if !userRole.IsActive {
			continue
		}

		if slices.Contains(roles, userRole.Role) {
			return true, nil
		}
	}

	return false, nil
}

func (s *RoleService) HasRole(user *models.User, role models.RoleType) (bool, error) {
	if len(user.Roles) == 0 {
		err := facades.Orm().Query().
			With("Roles").
			Where("id = ?", user.ID).
			First(user)

		if err != nil {
			return false, err
		}
	}

	for _, userRole := range user.Roles {
		if userRole.Role == role && userRole.IsActive {
			return true, nil
		}
	}

	return false, nil
}

func (s *RoleService) HasRoles(user *models.User, roles []models.RoleType) (bool, error) {
	if len(user.Roles) == 0 {
		err := facades.Orm().Query().
			With("Roles").
			Where("id = ?", user.ID).
			First(user)

		if err != nil {
			return false, err
		}
	}

	for _, userRole := range user.Roles {
		if !userRole.IsActive {
			continue
		}

		if slices.Contains(roles, userRole.Role) {
			return true, nil
		}
	}

	return false, nil
}

func (s *RoleService) GetUserRolesByUserID(userID uuid.UUID) ([]models.Role, error) {
	var user models.User

	err := facades.Orm().Query().
		With("Roles").
		Where("id = ?", userID).
		First(&user)

	if err != nil {
		return nil, err
	}

	var activeRoles []models.Role
	for _, role := range user.Roles {
		if role.IsActive {
			activeRoles = append(activeRoles, role)
		}
	}

	return activeRoles, nil
}

func (s *RoleService) GetUserRoles(user *models.User) ([]models.Role, error) {
	if len(user.Roles) == 0 {
		err := facades.Orm().Query().
			With("Roles").
			Where("id = ?", user.ID).
			First(user)

		if err != nil {
			return nil, err
		}
	}

	var activeRoles []models.Role
	for _, role := range user.Roles {
		if role.IsActive {
			activeRoles = append(activeRoles, role)
		}
	}

	return activeRoles, nil
}

func (s *RoleService) IsSuperAdmin(user *models.User) (bool, error) {
	return s.HasRole(user, models.RoleSuperAdmin)
}

func (s *RoleService) IsAdmin(user *models.User) (bool, error) {
	return s.HasRole(user, models.RoleAdmin)
}

func (s *RoleService) IsManager(user *models.User) (bool, error) {
	return s.HasRole(user, models.RoleManager)
}

func (s *RoleService) IsPartner(user *models.User) (bool, error) {
	return s.HasRole(user, models.RolePartner)
}

func (s *RoleService) IsClient(user *models.User) (bool, error) {
	return s.HasRole(user, models.RoleClient)
}

func (s *RoleService) IsSuperAdminByID(userID uuid.UUID) (bool, error) {
	return s.HasRoleByUserID(userID, models.RoleSuperAdmin)
}

func (s *RoleService) IsAdminByID(userID uuid.UUID) (bool, error) {
	return s.HasRoleByUserID(userID, models.RoleAdmin)
}

func (s *RoleService) IsManagerByID(userID uuid.UUID) (bool, error) {
	return s.HasRoleByUserID(userID, models.RoleManager)
}

func (s *RoleService) IsPartnerByID(userID uuid.UUID) (bool, error) {
	return s.HasRoleByUserID(userID, models.RolePartner)
}

func (s *RoleService) IsClientByID(userID uuid.UUID) (bool, error) {
	return s.HasRoleByUserID(userID, models.RoleClient)
}

package seeders

import (
	"errors"
	"pixel/app/models"

	"github.com/goravel/framework/facades"
	"gorm.io/gorm"
)

type RoleSeeder struct {
}

func (s *RoleSeeder) Signature() string {
	return "RoleSeeder"
}

// Run executes the seeder logic.
func (s *RoleSeeder) Run() error {
	rolesToSeed := []models.Role{
		{Role: models.RoleSuperAdmin, Description: "A super administrator with full system access."},
		{Role: models.RoleAdmin, Description: "An administrator with broad system permissions."},
		{Role: models.RoleManager, Description: "A manager who can oversee specific parts of the system."},
		{Role: models.RolePartner, Description: "A partner with limited, specific access."},
		{Role: models.RoleClient, Description: "A client with access only to their own data."},
	}

	for _, roleData := range rolesToSeed {
		var existingRole models.Role

		err := facades.Orm().Query().Where("role = ?", roleData.Role).First(&existingRole)

		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				newRole := models.Role{
					Role:        roleData.Role,
					Description: roleData.Description,
					IsActive:    true,
				}
				if createErr := facades.Orm().Query().Create(&newRole); createErr != nil {
					return createErr
				}
			} else {
				return err
			}
		}
	}

	return nil
}

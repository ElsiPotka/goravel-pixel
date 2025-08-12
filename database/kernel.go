package database

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/contracts/database/seeder"

	"pixel/database/migrations"
	"pixel/database/seeders"
)

type Kernel struct {
}

func (kernel Kernel) Migrations() []schema.Migration {
	return []schema.Migration{
		&migrations.M20210101000001CreateUsersTable{},
		&migrations.M20210101000002CreateJobsTable{},
		&migrations.M20250812210439CreateRolesTable{},
		&migrations.M20250812211259CreateUsersHasRoleTable{},
		&migrations.M20250812212932CreatePermissionsTable{},
		&migrations.M20250812213219CreateRoleHasPermissionTable{},
	}
}

func (kernel Kernel) Seeders() []seeder.Seeder {
	return []seeder.Seeder{
		&seeders.DatabaseSeeder{},
		&seeders.RoleSeeder{},
	}
}

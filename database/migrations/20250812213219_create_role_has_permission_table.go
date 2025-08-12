package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250812213219CreateRoleHasPermissionTable struct{}

// Signature The unique signature for the migration.
func (r *M20250812213219CreateRoleHasPermissionTable) Signature() string {
	return "20250812213219_create_role_has_permission_table"
}

// Up Run the migrations.
func (r *M20250812213219CreateRoleHasPermissionTable) Up() error {
	if !facades.Schema().HasTable("role_has_permission") {
		return facades.Schema().Create("role_has_permission", func(table schema.Blueprint) {
			table.Uuid("role_id")
			table.Uuid("permission_id")

			table.Primary("role_id", "permission_id")
			table.TimestampsTz()
			table.SoftDeletesTz()

			table.Foreign("role_id").References("id").On("roles")
			table.Foreign("permission_id").References("id").On("permissions")
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250812213219CreateRoleHasPermissionTable) Down() error {
	return facades.Schema().DropIfExists("role_has_permission")
}

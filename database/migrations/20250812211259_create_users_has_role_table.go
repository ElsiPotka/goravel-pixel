package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250812211259CreateUsersHasRoleTable struct{}

// Signature The unique signature for the migration.
func (r *M20250812211259CreateUsersHasRoleTable) Signature() string {
	return "20250812211259_create_users_has_role_table"
}

// Up Run the migrations.
func (r *M20250812211259CreateUsersHasRoleTable) Up() error {
	if !facades.Schema().HasTable("user_has_role") {
		return facades.Schema().Create("user_has_role", func(table schema.Blueprint) {
			table.Uuid("user_id")
			table.Uuid("role_id")
			table.Primary("user_id", "role_id")
			table.TimestampsTz()
			table.SoftDeletesTz()

			table.Foreign("user_id").References("id").On("users")
			table.Foreign("role_id").References("id").On("roles")
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250812211259CreateUsersHasRoleTable) Down() error {
	return facades.Schema().DropIfExists("user_has_role")
}

package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250812210439CreateRolesTable struct{}

// Signature The unique signature for the migration.
func (r *M20250812210439CreateRolesTable) Signature() string {
	return "20250812210439_create_roles_table"
}

// Up Run the migrations.
func (r *M20250812210439CreateRolesTable) Up() error {
	if !facades.Schema().HasTable("roles") {
		return facades.Schema().Create("roles", func(table schema.Blueprint) {
			table.Uuid("id")
			table.Primary("id")
			table.String("role", 50)
			table.Unique("role")
			table.Text("description").Nullable()
			table.Boolean("is_active").Default(true)
			table.Index("is_active")
			table.TimestampsTz()
			table.SoftDeletesTz()
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250812210439CreateRolesTable) Down() error {
	return facades.Schema().DropIfExists("roles")
}

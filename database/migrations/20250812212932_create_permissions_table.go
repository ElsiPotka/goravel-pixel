package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250812212932CreatePermissionsTable struct{}

// Signature The unique signature for the migration.
func (r *M20250812212932CreatePermissionsTable) Signature() string {
	return "20250812212932_create_permissions_table"
}

// Up Run the migrations.
func (r *M20250812212932CreatePermissionsTable) Up() error {
	if !facades.Schema().HasTable("permissions") {
		return facades.Schema().Create("permissions", func(table schema.Blueprint) {
			table.Uuid("id")
			table.Primary("id")
			table.String("permission", 50)
			table.Unique("permission")
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
func (r *M20250812212932CreatePermissionsTable) Down() error {
	return facades.Schema().DropIfExists("permissions")
}

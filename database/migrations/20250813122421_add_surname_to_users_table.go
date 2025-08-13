package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250813122421AddSurnameToUsersTable struct{}

// Signature The unique signature for the migration.
func (r *M20250813122421AddSurnameToUsersTable) Signature() string {
	return "20250813122421_add_surname_to_users_table"
}

// Up Run the migrations.
func (r *M20250813122421AddSurnameToUsersTable) Up() error {
	return facades.Schema().Table("users", func(table schema.Blueprint) {
		table.String("surname", 100).After("name")
	})
}

// Down Reverse the migrations.
func (r *M20250813122421AddSurnameToUsersTable) Down() error {
	return facades.Schema().Table("users", func(table schema.Blueprint) {
		table.DropColumn("surname")
	})
}

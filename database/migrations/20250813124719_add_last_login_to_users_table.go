package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250813124719AddLastLoginToUsersTable struct{}

// Signature The unique signature for the migration.
func (r *M20250813124719AddLastLoginToUsersTable) Signature() string {
	return "20250813124719_add_last_login_to_users_table"
}

// Up Run the migrations.
func (r *M20250813124719AddLastLoginToUsersTable) Up() error {
	return facades.Schema().Table("users", func(table schema.Blueprint) {

		table.TimestampTz("last_login").Nullable().After("password")
	})
}

// Down Reverse the migrations.
func (r *M20250813124719AddLastLoginToUsersTable) Down() error {
	return facades.Schema().Table("users", func(table schema.Blueprint) {
		table.DropColumn("last_login")
	})
}

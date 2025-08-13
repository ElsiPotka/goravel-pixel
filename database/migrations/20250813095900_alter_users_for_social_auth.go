package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250813095900AlterUsersForSocialAuth struct{}

// Signature The unique signature for the migration.
func (r *M20250813095900AlterUsersForSocialAuth) Signature() string {
	return "20250813095900_alter_users_for_social_auth"
}

// Up Run the migrations.
func (r *M20250813095900AlterUsersForSocialAuth) Up() error {
	return facades.Schema().Table("users", func(table schema.Blueprint) {

		table.String("password").Nullable().Change()

		table.String("avatar_url", 255).Nullable()
	})
}

// Down Reverse the migrations.
func (r *M20250813095900AlterUsersForSocialAuth) Down() error {
	return facades.Schema().Table("users", func(table schema.Blueprint) {
		// Revert the password column to be non-nullable.
		//table.String("password").NotNullable().Change()

		// Drop the avatar_url column.
		table.DropColumn("avatar_url")
	})
}

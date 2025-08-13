package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250813095246CreateSocialAccountsTable struct{}

// Signature The unique signature for the migration.
func (r *M20250813095246CreateSocialAccountsTable) Signature() string {
	return "20250813095246_create_social_accounts_table"
}

// Up Run the migrations.
func (r *M20250813095246CreateSocialAccountsTable) Up() error {
	if !facades.Schema().HasTable("social_accounts") {
		return facades.Schema().Create("social_accounts", func(table schema.Blueprint) {
			table.Uuid("id")
			table.Primary("id")
			table.Uuid("user_id")
			table.String("provider_name", 50)
			table.String("provider_id", 255)

			table.Foreign("user_id").References("id").On("users")

			table.Unique("user_id", "provider_name")
			table.Unique("provider_name", "provider_id")

			table.TimestampsTz()
			table.SoftDeletesTz()
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250813095246CreateSocialAccountsTable) Down() error {
	return facades.Schema().DropIfExists("social_accounts")
}

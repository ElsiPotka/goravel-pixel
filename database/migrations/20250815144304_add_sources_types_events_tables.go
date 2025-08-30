package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250815144304AddSourcesTypesEventsTables struct{}

// Signature The unique signature for the migration.
func (r *M20250815144304AddSourcesTypesEventsTables) Signature() string {
	return "20250815144304_add_sources_types_events_tables"
}

// Up Run the migrations.
func (r *M20250815144304AddSourcesTypesEventsTables) Up() error {
	if err := facades.Schema().Create("sources", func(table schema.Blueprint) {
		table.Uuid("id")
		table.Primary("id")
		table.String("name", 255)
		table.Text("description")
		table.Integer("max_audience")
		table.TimestampsTz()
		table.SoftDeletesTz()
	}); err != nil {
		return err
	}

	if err := facades.Schema().Create("types", func(table schema.Blueprint) {
		table.Uuid("id")
		table.Primary("id")
		table.String("name", 255)
		table.Text("description")
		table.Uuid("source_id").Nullable()

		table.Foreign("source_id").References("id").On("sources")

		table.TimestampsTz()
		table.SoftDeletesTz()
	}); err != nil {
		return err
	}

	if err := facades.Schema().Create("events", func(table schema.Blueprint) {
		table.Uuid("id")
		table.Primary("id")
		table.String("name", 255)
		table.Text("description")
		table.Uuid("type_id").Nullable()

		table.Foreign("type_id").References("id").On("types")

		table.TimestampsTz()
		table.SoftDeletesTz()
	}); err != nil {
		return err
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250815144304AddSourcesTypesEventsTables) Down() error {
	if err := facades.Schema().DropIfExists("events"); err != nil {
		return err
	}
	if err := facades.Schema().DropIfExists("types"); err != nil {
		return err
	}
	return facades.Schema().DropIfExists("sources")
}

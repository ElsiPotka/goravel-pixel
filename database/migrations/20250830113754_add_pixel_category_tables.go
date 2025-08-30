package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250830113754AddPixelCategoryTables struct{}

// Signature The unique signature for the migration.
func (r *M20250830113754AddPixelCategoryTables) Signature() string {
	return "20250830113754_add_pixel_category_tables"
}

// Up Run the migrations.
func (r *M20250830113754AddPixelCategoryTables) Up() error {
	if err := facades.Schema().Create("categories", func(table schema.Blueprint) {
		table.Uuid("id")
		table.Primary("id")
		table.String("name", 255)
		table.Text("description")
		table.TimestampsTz()
		table.SoftDeletesTz()
	}); err != nil {
		return err
	}

	if err := facades.Schema().Create("statuses", func(table schema.Blueprint) {
		table.Uuid("id")
		table.Primary("id")
		table.String("name", 255)
		table.Text("description")
		table.String("color", 255)
		table.String("context", 255)
		table.TimestampsTz()
		table.SoftDeletesTz()
	}); err != nil {
		return err
	}

	if err := facades.Schema().Create("pixels", func(table schema.Blueprint) {
		table.Uuid("id")
		table.Primary("id")
		table.String("name", 255)
		table.Text("description")
		table.Integer("audience").Default(0)
		table.String("website_logo", 255)
		table.String("website_url", 255)
		table.Float("price", 10, 2)
		table.String("currency", 10)
		table.String("audience_proof", 255)

		table.Uuid("creator_id").Nullable()
		table.Foreign("creator_id").References("id").On("users")

		table.Uuid("reviewer_id").Nullable()
		table.Foreign("reviewer_id").References("id").On("users")

		table.Uuid("source_id").Nullable()
		table.Foreign("source_id").References("id").On("sources")

		table.Uuid("event_id").Nullable()
		table.Foreign("event_id").References("id").On("events")

		table.Uuid("type_id").Nullable()
		table.Foreign("type_id").References("id").On("types")

		table.Uuid("status_id").Nullable()
		table.Foreign("status_id").References("id").On("statuses")

		table.Uuid("category_id").Nullable()
		table.Foreign("category_id").References("id").On("categories")

		table.TimestampsTz()
		table.SoftDeletesTz()
	}); err != nil {
		return err
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250830113754AddPixelCategoryTables) Down() error {
	if err := facades.Schema().DropIfExists("pixels"); err != nil {
		return err
	}
	if err := facades.Schema().DropIfExists("statuses"); err != nil {
		return err
	}
	return facades.Schema().DropIfExists("categories")
}

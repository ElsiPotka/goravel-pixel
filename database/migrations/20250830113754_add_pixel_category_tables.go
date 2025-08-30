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

	if err := facades.Schema().Create("pixels", func(table schema.Blueprint) {
		table.Uuid("id")
		table.Primary("id")
		table.String("name", 255)
		table.Text("description")
		table.String("website_logo", 255)
		table.String("website_url", 255)
		table.Float("price", 10, 2)
		table.String("currency", 10)
		table.String("audience_proof", 255)

		table.Uuid("source_id").Nullable()
		table.Foreign("source_id").References("id").On("sources")

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
	return facades.Schema().DropIfExists("categories")
}

// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// MoviesColumns holds the columns for the "movies" table.
	MoviesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "title", Type: field.TypeString},
		{Name: "director", Type: field.TypeString},
	}
	// MoviesTable holds the schema information for the "movies" table.
	MoviesTable = &schema.Table{
		Name:       "movies",
		Columns:    MoviesColumns,
		PrimaryKey: []*schema.Column{MoviesColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		MoviesTable,
	}
)

func init() {
}

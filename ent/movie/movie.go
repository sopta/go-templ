// Code generated by ent, DO NOT EDIT.

package movie

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the movie type in the database.
	Label = "movie"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldDirector holds the string denoting the director field in the database.
	FieldDirector = "director"
	// Table holds the table name of the movie in the database.
	Table = "movies"
)

// Columns holds all SQL columns for movie fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldDirector,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the Movie queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByDirector orders the results by the director field.
func ByDirector(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDirector, opts...).ToFunc()
}
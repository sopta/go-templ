package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Movie holds the schema definition for the Movie entity.
type Movie struct {
	ent.Schema
}

// Fields of the Movie.
func (Movie) Fields() []ent.Field {
	return []ent.Field{
		field.String("title"),
		field.String("director"),
	}
}

// Edges of the Movie.
func (Movie) Edges() []ent.Edge {
	return nil
}

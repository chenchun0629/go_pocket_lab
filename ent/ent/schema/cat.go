package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Cat holds the schema definition for the Cat entity.
type Cat struct {
	ent.Schema
}

// Fields of the Cat.
func (Cat) Fields() []ent.Field {
	return []ent.Field{
		field.Int("age").
			Positive(),
		field.String("name").
			Default("unknown"),
	}
}

// Edges of the Cat.
func (Cat) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", Admin.Type).
			Ref("cats").
			Unique(),

		edge.To("cb", Bee.Type),
	}
}

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Bee holds the schema definition for the Bee entity.
type Bee struct {
	ent.Schema
}

// Fields of the Bee.
func (Bee) Fields() []ent.Field {
	return []ent.Field{
		field.Int("age").
			Positive(),
		field.String("name").
			Default("unknown"),
	}
}

// Edges of the Bee.
func (Bee) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", Admin.Type).
			Ref("bees"),

		edge.From("bc", Cat.Type).
			Ref("cb"),
	}
}

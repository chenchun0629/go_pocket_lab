package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// FieldTest holds the schema definition for the FieldTest entity.
type FieldTest struct {
	ent.Schema
}

func (FieldTest) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Charset:   "utf8mb4",
			Collation: "utf8mb4_unicode_ci",
		},
	}
}

type FTConfig struct {
	AppId string
}

// Fields of the FieldTest.
func (FieldTest) Fields() []ent.Field {
	return []ent.Field{
		field.Time("date_d"),
		field.Time("date").SchemaType(map[string]string{
			dialect.MySQL: "date",
		}),
		field.String("title").Annotations(entsql.Annotation{
			Size: 64,
		}),
		field.JSON("j_f", []string{}),
		field.JSON("j_s_f", &FTConfig{}),
		field.String("app_id"),
	}
}

// Edges of the FieldTest.
func (FieldTest) Edges() []ent.Edge {
	return nil
}

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// CasbinRules holds the schema definition for the CasbinRules entity.
type CasbinRules struct {
	ent.Schema
}

// Fields of the CasbinRules.
func (CasbinRules) Fields() []ent.Field {
	return []ent.Field{
		field.String("ptype").Optional().Default(""),
		field.String("v0").Optional().Default(""),
		field.String("v1").Optional().Default(""),
		field.String("v2").Optional().Default(""),
		field.String("v3").Optional().Default(""),
		field.String("v4").Optional().Default(""),
		field.String("v5").Optional().Default(""),
	}
}

// Edges of the CasbinRules.
func (CasbinRules) Edges() []ent.Edge {
	return nil
}

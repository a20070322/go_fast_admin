package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"time"
)

type AuditMixin struct{
	mixin.Schema
}

// Fields of the AuditMixin.
func (AuditMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			//StructTag(`json:"-"`).
			Immutable().
			Default(time.Now),
		field.Time("updated_at").
			//StructTag(`json:"-"`).
			Optional().
			UpdateDefault(time.Now),
		field.Time("deleted_at").
			StructTag(`json:"-"`).
			Optional(),
	}
}
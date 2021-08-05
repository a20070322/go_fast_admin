package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AdminRole holds the schema definition for the AdminRole entity.
type AdminRole struct {
	ent.Schema
}

func (AdminRole) Mixin() []ent.Mixin {
	return []ent.Mixin{
		AuditMixin{},
	}
}
// Fields of the AdminRole.
func (AdminRole) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Unique().
			Comment("角色名称"),
		field.Bool("is_enable").
			Comment("是否启用").
			Default(false).
			Optional(),
	}
}

// Edges of the AdminRole.
func (AdminRole) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", AdminUser.Type),
		edge.To("menu", AdminMenus.Type),
	}
}

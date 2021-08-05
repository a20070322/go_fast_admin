package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// AdminUser holds the schema definition for the AdminUser entity.
type AdminUser struct {
	ent.Schema
}
func (AdminUser) Mixin() []ent.Mixin {
	return []ent.Mixin{
		AuditMixin{},
	}
}
// Fields of the AdminUser.
func (AdminUser) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Comment("id"),
		field.String("username").
			Unique().
			Comment("用户名"),
		field.String("password").
			Sensitive().
			Comment("密码"),
		field.String("avatar").
			Optional().
			Comment("头像"),
		field.String("phone").
			Optional().
			Comment("手机号"),
		field.Bool("is_enable").
			Comment("是否启用").
			Default(false).
			Optional(),
	}
}

// Edges of the AdminUser.
func (AdminUser) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("role", AdminRole.Type).
			Ref("user"),
	}
}

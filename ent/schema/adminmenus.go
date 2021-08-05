package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AdminMenus holds the schema definition for the AdminMenus entity.
type AdminMenus struct {
	ent.Schema
}

func (AdminMenus) Mixin() []ent.Mixin {
	return []ent.Mixin{
		AuditMixin{},
	}
}
// Fields of the AdminMenus.
func (AdminMenus) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Comment("菜单名称"),
		field.String("path").
			Optional().
			Comment("路由地址"),
		field.String("path_action").
			Optional().
			Comment("路由类型"),
		field.String("router_path").
			Optional().
			Comment("页面路由地址"),
		field.String("icon").
			Comment("菜单图标").
			Default("").
			Optional(),
		field.Int8("type").
			Comment("菜单类型 1目录 2菜单 3按钮").
			Default(0).
			Optional(),
		field.String("power_str").
			Comment("权限字符").
			Optional(),
		field.Int("sort").
			Comment("排序").
			Default(0).
			Min(0).
			Optional(),
		field.Int("fid").
			Comment("父级id").
			Optional(),
		field.Bool("is_external_link").
			Comment("是否外链").
			Default(false).
			Optional(),
		field.Bool("is_show").
			Comment("显示状态").
			Default(true).
			Optional(),
		field.Bool("is_enable").
			Comment("是否启用").
			Default(false).
			Optional(),
	}
}

// Edges of the AdminMenus.
func (AdminMenus) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("role", AdminRole.Type).
			Ref("menu"),
	}
}

package schema
import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	)
// AdminDict holds the schema definition for the AdminDict entity.
type AdminDict struct {
	ent.Schema
}

func (AdminDict) Mixin() []ent.Mixin {
    return []ent.Mixin{
        AuditMixin{},
    }
}
// Fields of the AdminDict.
func (AdminDict) Fields() []ent.Field {
	return []ent.Field{
    field.String("dict_type").
        Unique().
        Comment("字典类型"),
    field.String("dict_name").
        Comment("字典名称"),
    field.String("remarks").
        Comment("备注"),
    field.Bool("is_enable").
        Comment("是否启用"),
    }
}

// Edges of the AdminDict.
func (AdminDict) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("key", AdminDictKey.Type),
	}
}


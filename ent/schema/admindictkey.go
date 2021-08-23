package schema
import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field")
// AdminDictKey holds the schema definition for the AdminDictKey entity.
type AdminDictKey struct {
	ent.Schema
}

func (AdminDictKey) Mixin() []ent.Mixin {
    return []ent.Mixin{
        AuditMixin{},
    }
}
// Fields of the AdminDictKey.
func (AdminDictKey) Fields() []ent.Field {
	return []ent.Field{
    field.String("dict_label").
        Comment("字典标签"),
    field.String("dict_code").
        Comment("字典键值"),
    field.Int("sort").
    	Optional().
        Comment("排序"),
    field.String("remarks").
    	Optional().
        Comment("备注"),
    field.Bool("is_enable").
        Comment("状态"),
    }
}

// Edges of the AdminDictKey.
func (AdminDictKey) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("P", AdminDict.Type).
			Ref("key").Unique(),
	}
}


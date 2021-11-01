package schema
import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	)
// AutoUserExample holds the schema definition for the AutoUserExample entity.

type AutoUserExample struct {
	ent.Schema
}

// AutoUserExample  Mixin
func (AutoUserExample) Mixin() []ent.Mixin {
    return []ent.Mixin{
            AuditMixin{},
        }
}

// Fields of the AutoUserExample.
func (AutoUserExample) Fields() []ent.Field {
	return []ent.Field{
    field.String("test_text").
        Default("").
        Optional().
        Comment("文本域测试"),
    field.Bool("test_bool").
        Default(true).
        Optional().
        Comment("布尔值测试"),
    field.Int("test_int").Optional().
        Comment("数字类型测试"),
    }
}


// Edges of the AutoUserExample.
func (AutoUserExample) Edges() []ent.Edge {
	return []ent.Edge{

	}
}
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
)

// Group holds the schema definition for the Group entity.
type Group struct {
	ent.Schema
}

// Fields of the Group.
func (Group) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			SchemaType(map[string]string{
				dialect.MySQL: "varchar(100)",
			}),
		field.Text("description").
			NotEmpty().
			SchemaType(map[string]string{
				dialect.MySQL: "varchar(100)",
			}),
	}
}

// Edges of the Group.
func (Group) Edges() []ent.Edge {
	return nil
}

func (Group) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

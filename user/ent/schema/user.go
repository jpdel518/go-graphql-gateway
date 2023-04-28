package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("first_name").
			NotEmpty().
			SchemaType(map[string]string{
				dialect.MySQL: "varchar(100)",
			}),
		field.Text("last_name").
			NotEmpty().
			SchemaType(map[string]string{
				dialect.MySQL: "varchar(100)",
			}),
		field.Int("age").
			Default(20),
		field.String("address").
			NotEmpty().
			SchemaType(map[string]string{
				dialect.MySQL: "varchar(300)",
			}),
		field.String("email").
			Optional(),
		field.Int("group_id").
			Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

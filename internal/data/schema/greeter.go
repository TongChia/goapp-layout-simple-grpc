package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Greeter holds the schema definition for the User entity.
type Greeter struct {
	ent.Schema
}

// Fields of the Greeter.
func (Greeter) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("title").Unique(),
	}
}

func (Greeter) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

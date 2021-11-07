package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// ApplicationUser holds the schema definition for the ApplicationUser entity.
type ApplicationUser struct {
	ent.Schema
}

// Fields of the ApplicationUser.
func (ApplicationUser) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.String("app_id"),
		field.UUID("user_id", uuid.UUID{}),
		field.Bool("original").Default(true),
		field.Int64("create_at").
			DefaultFunc(func() int64 {
				return time.Now().Unix()
			}),
		field.Int64("delete_at").
			DefaultFunc(func() int64 {
				return 0
			}),
	}
}

// Edges of the ApplicationUser.
func (ApplicationUser) Edges() []ent.Edge {
	return nil
}

func (ApplicationUser) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("app_id", "user_id"),
		index.Fields("app_id"),
		index.Fields("user_id"),
		index.Fields("original"),
	}
}

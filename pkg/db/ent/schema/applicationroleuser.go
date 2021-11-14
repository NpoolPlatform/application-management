package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// ApplicationRoleUser holds the schema definition for the ApplicationRoleUser entity.
type ApplicationRoleUser struct {
	ent.Schema
}

// Fields of the ApplicationRoleUser.
func (ApplicationRoleUser) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.String("app_id"),
		field.UUID("role_id", uuid.UUID{}),
		field.UUID("user_id", uuid.UUID{}),
		field.Uint32("create_at").
			DefaultFunc(func() uint32 {
				return uint32(time.Now().Unix())
			}),
		field.Uint32("delete_at").
			DefaultFunc(func() uint32 {
				return 0
			}),
	}
}

// Edges of the ApplicationRoleUser.
func (ApplicationRoleUser) Edges() []ent.Edge {
	return nil
}

func (ApplicationRoleUser) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("app_id", "user_id"),
		index.Fields("app_id", "role_id"),
	}
}

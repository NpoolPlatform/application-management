package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// ApplicationGroupUser holds the schema definition for the ApplicationGroupUser entity.
type ApplicationGroupUser struct {
	ent.Schema
}

// Fields of the ApplicationGroupUser.
func (ApplicationGroupUser) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.UUID("group_id", uuid.UUID{}),
		field.String("app_id"),
		field.UUID("user_id", uuid.UUID{}),
		field.String("annotation").Optional(),
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

// Edges of the ApplicationGroupUser.
func (ApplicationGroupUser) Edges() []ent.Edge {
	return nil
}

func (ApplicationGroupUser) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("group_id", "app_id"),
		index.Fields("user_id", "app_id"),
		index.Fields("user_id"),
		index.Fields("app_id"),
	}
}

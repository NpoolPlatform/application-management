package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// ApplicationRole holds the schema definition for the ApplicationRole entity.
type ApplicationRole struct {
	ent.Schema
}

// Fields of the ApplicationRole.
func (ApplicationRole) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.UUID("app_id", uuid.UUID{}),
		field.String("role_name"),
		field.UUID("creator", uuid.UUID{}),
		field.Uint32("create_at").
			DefaultFunc(func() uint32 {
				return uint32(time.Now().Unix())
			}),
		field.Uint32("update_at").
			DefaultFunc(func() uint32 {
				return uint32(time.Now().Unix())
			}).
			UpdateDefault(func() uint32 {
				return uint32(time.Now().Unix())
			}),
		field.Uint32("delete_at").
			DefaultFunc(func() uint32 {
				return 0
			}),
		field.String("annotation").Optional(),
	}
}

// Edges of the ApplicationRole.
func (ApplicationRole) Edges() []ent.Edge {
	return nil
}

func (ApplicationRole) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("app_id"),
		index.Fields("creator"),
	}
}

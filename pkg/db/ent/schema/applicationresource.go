package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// ApplicationResource holds the schema definition for the ApplicationResource entity.
type ApplicationResource struct {
	ent.Schema
}

// Fields of the ApplicationResource.
func (ApplicationResource) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.String("app_id"),
		field.String("resource_name"),
		field.String("resource_description").Optional(),
		field.String("type").
			Default("API"),
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
	}
}

// Edges of the ApplicationResource.
func (ApplicationResource) Edges() []ent.Edge {
	return nil
}

func (ApplicationResource) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("app_id"),
	}
}

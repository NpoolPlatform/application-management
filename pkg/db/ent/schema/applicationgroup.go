package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// ApplicationGroup holds the schema definition for the ApplicationGroup entity.
type ApplicationGroup struct {
	ent.Schema
}

// Fields of the ApplicationGroup.
func (ApplicationGroup) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.UUID("app_id", uuid.UUID{}),
		field.String("group_name"),
		field.String("group_logo").Optional(),
		field.UUID("group_owner", uuid.UUID{}),
		field.String("annotation").Optional(),
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

// Edges of the ApplicationGroup.
func (ApplicationGroup) Edges() []ent.Edge {
	return nil
}

func (ApplicationGroup) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("app_id"),
		index.Fields("group_owner"),
	}
}

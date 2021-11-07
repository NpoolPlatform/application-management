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
		field.String("app_id"),
		field.String("group_name"),
		field.String("group_logo"),
		field.UUID("group_owner", uuid.UUID{}),
		field.String("annotation").Optional(),
		field.Int64("create_at").
			DefaultFunc(func() int64 {
				return time.Now().Unix()
			}),
		field.Int64("update_at").
			DefaultFunc(func() int64 {
				return time.Now().Unix()
			}).
			UpdateDefault(func() int64 {
				return time.Now().Unix()
			}),
		field.Int64("delete_at").
			DefaultFunc(func() int64 {
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

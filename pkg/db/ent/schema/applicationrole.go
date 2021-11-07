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
		field.String("app_id"),
		field.String("role_name"),
		field.UUID("creator", uuid.UUID{}),
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

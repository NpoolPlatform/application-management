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
		field.UUID("app_id", uuid.UUID{}),
		field.UUID("user_id", uuid.UUID{}),
		field.Bool("original").Default(true),
		field.Bool("kyc_verify").Default(false),
		field.Bool("ga_verify").Default(false),
		field.Bool("ga_login").Default(false),
		field.Uint32("Login_number").Default(0),
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

package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/thanhpk/randstr"
)

// Empty holds the schema definition for the Empty entity.
type Application struct {
	ent.Schema
}

// Fields of the Empty.
func (Application) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.String("application_name").Unique(),
		field.UUID("application_owner", uuid.UUID{}),
		field.String("homepage_url").Optional(),
		field.String("redirect_url").Optional(),
		field.String("client_secret").Unique().
			DefaultFunc(func() string {
				return randstr.Hex(20)
			}).Sensitive(),
		field.String("application_logo").Optional(),
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

// Edges of the Empty.
func (Application) Edges() []ent.Edge {
	return nil
}

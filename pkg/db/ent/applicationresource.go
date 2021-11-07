// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/application-management/pkg/db/ent/applicationresource"
	"github.com/google/uuid"
)

// ApplicationResource is the model entity for the ApplicationResource schema.
type ApplicationResource struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// AppID holds the value of the "app_id" field.
	AppID string `json:"app_id,omitempty"`
	// ResourceName holds the value of the "resource_name" field.
	ResourceName string `json:"resource_name,omitempty"`
	// ResourceDescription holds the value of the "resource_description" field.
	ResourceDescription string `json:"resource_description,omitempty"`
	// Type holds the value of the "type" field.
	Type string `json:"type,omitempty"`
	// Creator holds the value of the "creator" field.
	Creator uuid.UUID `json:"creator,omitempty"`
	// CreateAt holds the value of the "create_at" field.
	CreateAt int64 `json:"create_at,omitempty"`
	// UpdateAt holds the value of the "update_at" field.
	UpdateAt int64 `json:"update_at,omitempty"`
	// DeleteAt holds the value of the "delete_at" field.
	DeleteAt int64 `json:"delete_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ApplicationResource) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case applicationresource.FieldCreateAt, applicationresource.FieldUpdateAt, applicationresource.FieldDeleteAt:
			values[i] = new(sql.NullInt64)
		case applicationresource.FieldAppID, applicationresource.FieldResourceName, applicationresource.FieldResourceDescription, applicationresource.FieldType:
			values[i] = new(sql.NullString)
		case applicationresource.FieldID, applicationresource.FieldCreator:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type ApplicationResource", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ApplicationResource fields.
func (ar *ApplicationResource) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case applicationresource.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				ar.ID = *value
			}
		case applicationresource.FieldAppID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value.Valid {
				ar.AppID = value.String
			}
		case applicationresource.FieldResourceName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field resource_name", values[i])
			} else if value.Valid {
				ar.ResourceName = value.String
			}
		case applicationresource.FieldResourceDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field resource_description", values[i])
			} else if value.Valid {
				ar.ResourceDescription = value.String
			}
		case applicationresource.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				ar.Type = value.String
			}
		case applicationresource.FieldCreator:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field creator", values[i])
			} else if value != nil {
				ar.Creator = *value
			}
		case applicationresource.FieldCreateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_at", values[i])
			} else if value.Valid {
				ar.CreateAt = value.Int64
			}
		case applicationresource.FieldUpdateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field update_at", values[i])
			} else if value.Valid {
				ar.UpdateAt = value.Int64
			}
		case applicationresource.FieldDeleteAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field delete_at", values[i])
			} else if value.Valid {
				ar.DeleteAt = value.Int64
			}
		}
	}
	return nil
}

// Update returns a builder for updating this ApplicationResource.
// Note that you need to call ApplicationResource.Unwrap() before calling this method if this ApplicationResource
// was returned from a transaction, and the transaction was committed or rolled back.
func (ar *ApplicationResource) Update() *ApplicationResourceUpdateOne {
	return (&ApplicationResourceClient{config: ar.config}).UpdateOne(ar)
}

// Unwrap unwraps the ApplicationResource entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ar *ApplicationResource) Unwrap() *ApplicationResource {
	tx, ok := ar.config.driver.(*txDriver)
	if !ok {
		panic("ent: ApplicationResource is not a transactional entity")
	}
	ar.config.driver = tx.drv
	return ar
}

// String implements the fmt.Stringer.
func (ar *ApplicationResource) String() string {
	var builder strings.Builder
	builder.WriteString("ApplicationResource(")
	builder.WriteString(fmt.Sprintf("id=%v", ar.ID))
	builder.WriteString(", app_id=")
	builder.WriteString(ar.AppID)
	builder.WriteString(", resource_name=")
	builder.WriteString(ar.ResourceName)
	builder.WriteString(", resource_description=")
	builder.WriteString(ar.ResourceDescription)
	builder.WriteString(", type=")
	builder.WriteString(ar.Type)
	builder.WriteString(", creator=")
	builder.WriteString(fmt.Sprintf("%v", ar.Creator))
	builder.WriteString(", create_at=")
	builder.WriteString(fmt.Sprintf("%v", ar.CreateAt))
	builder.WriteString(", update_at=")
	builder.WriteString(fmt.Sprintf("%v", ar.UpdateAt))
	builder.WriteString(", delete_at=")
	builder.WriteString(fmt.Sprintf("%v", ar.DeleteAt))
	builder.WriteByte(')')
	return builder.String()
}

// ApplicationResources is a parsable slice of ApplicationResource.
type ApplicationResources []*ApplicationResource

func (ar ApplicationResources) config(cfg config) {
	for _i := range ar {
		ar[_i].config = cfg
	}
}

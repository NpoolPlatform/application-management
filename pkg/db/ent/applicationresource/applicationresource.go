// Code generated by entc, DO NOT EDIT.

package applicationresource

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the applicationresource type in the database.
	Label = "application_resource"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAppID holds the string denoting the app_id field in the database.
	FieldAppID = "app_id"
	// FieldResourceName holds the string denoting the resource_name field in the database.
	FieldResourceName = "resource_name"
	// FieldResourceDescription holds the string denoting the resource_description field in the database.
	FieldResourceDescription = "resource_description"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldCreator holds the string denoting the creator field in the database.
	FieldCreator = "creator"
	// FieldCreateAt holds the string denoting the create_at field in the database.
	FieldCreateAt = "create_at"
	// FieldUpdateAt holds the string denoting the update_at field in the database.
	FieldUpdateAt = "update_at"
	// FieldDeleteAt holds the string denoting the delete_at field in the database.
	FieldDeleteAt = "delete_at"
	// Table holds the table name of the applicationresource in the database.
	Table = "application_resources"
)

// Columns holds all SQL columns for applicationresource fields.
var Columns = []string{
	FieldID,
	FieldAppID,
	FieldResourceName,
	FieldResourceDescription,
	FieldType,
	FieldCreator,
	FieldCreateAt,
	FieldUpdateAt,
	FieldDeleteAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultType holds the default value on creation for the "type" field.
	DefaultType string
	// DefaultCreateAt holds the default value on creation for the "create_at" field.
	DefaultCreateAt func() uint32
	// DefaultUpdateAt holds the default value on creation for the "update_at" field.
	DefaultUpdateAt func() uint32
	// UpdateDefaultUpdateAt holds the default value on update for the "update_at" field.
	UpdateDefaultUpdateAt func() uint32
	// DefaultDeleteAt holds the default value on creation for the "delete_at" field.
	DefaultDeleteAt func() uint32
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

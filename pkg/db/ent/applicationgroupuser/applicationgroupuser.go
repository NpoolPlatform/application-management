// Code generated by entc, DO NOT EDIT.

package applicationgroupuser

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the applicationgroupuser type in the database.
	Label = "application_group_user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldGroupID holds the string denoting the group_id field in the database.
	FieldGroupID = "group_id"
	// FieldAppID holds the string denoting the app_id field in the database.
	FieldAppID = "app_id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldAnnotation holds the string denoting the annotation field in the database.
	FieldAnnotation = "annotation"
	// FieldCreateAt holds the string denoting the create_at field in the database.
	FieldCreateAt = "create_at"
	// FieldDeleteAt holds the string denoting the delete_at field in the database.
	FieldDeleteAt = "delete_at"
	// Table holds the table name of the applicationgroupuser in the database.
	Table = "application_group_users"
)

// Columns holds all SQL columns for applicationgroupuser fields.
var Columns = []string{
	FieldID,
	FieldGroupID,
	FieldAppID,
	FieldUserID,
	FieldAnnotation,
	FieldCreateAt,
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
	// DefaultCreateAt holds the default value on creation for the "create_at" field.
	DefaultCreateAt func() int64
	// DefaultDeleteAt holds the default value on creation for the "delete_at" field.
	DefaultDeleteAt func() int64
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

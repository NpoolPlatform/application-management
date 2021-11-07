// Code generated by entc, DO NOT EDIT.

package applicationrole

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/application-management/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.ApplicationRole {
	return predicate.ApplicationRole(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.ApplicationRole {
	return predicate.ApplicationRole(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.ApplicationRole {
	return predicate.ApplicationRole(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.ApplicationRole {
	return predicate.ApplicationRole(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.ApplicationRole {
	return predicate.ApplicationRole(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.ApplicationRole {
	return predicate.ApplicationRole(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.ApplicationRole {
	return predicate.ApplicationRole(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.ApplicationRole {
	return predicate.ApplicationRole(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.ApplicationRole {
	return predicate.ApplicationRole(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// AppID applies equality check predicate on the "app_id" field. It's identical to AppIDEQ.
func AppID(v string) predicate.ApplicationRole {
	return predicate.ApplicationRole(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppID), v))
	})
}

// RoleName applies equality check predicate on the "role_name" field. It's identical to RoleNameEQ.
func RoleName(v string) predicate.ApplicationRole {
	return predicate.ApplicationRole(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRoleName), v))
	})
}

// Creator applies equality check predicate on the "creator" field. It's identical to CreatorEQ.
func Creator(v uuid.UUID) predicate.ApplicationRole {
	return predicate.ApplicationRole(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreator), v))
	})
}

// CreateAt applies equality check predicate on the "create_at" field. It's identical to CreateAtEQ.
func CreateAt(v int64) predicate.ApplicationRole {
	return predicate.ApplicationRole(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateAt), v))
	})
}

// UpdateAt applies equality check predicate on the "update_at" field. It's identical to UpdateAtEQ.
func UpdateAt(v int64) predicate.ApplicationRole {
	return predicate.ApplicationRole(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateAt), v))
	})
}

// DeleteAt applies equality check predicate on the "delete_at" field. It's identical to DeleteAtEQ.
func DeleteAt(v int64) predicate.ApplicationRole {
	return predicate.ApplicationRole(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeleteAt), v))
	})
}

// Annotation applies equality check predicate on the "annotation" field. It's identical to AnnotationEQ.
func Annotation(v string) predicate.ApplicationRole {
	return predicate.ApplicationRole(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAnnotation), v))
	})
}

// AppIDEQ applies the EQ predicate on the "app_id" field.
func AppIDEQ(v string) predicate.ApplicationRole {
	return predicate.ApplicationRole(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppID), v))
	})
}

// AppIDNEQ applies the NEQ predicate on the "app_id" field.
func AppIDNEQ(v string) predicate.ApplicationRole {
	return predicate.ApplicationRole(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAppID), v))
	})
}

// AppIDIn applies the In predicate on the "app_id" field.
func AppIDIn(vs ...string) predicate.ApplicationRole {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ApplicationRole(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAppID), v...))
	})
}

// AppIDNotIn applies the NotIn predicate on the "app_id" field.
func AppIDNotIn(vs ...string) predicate.ApplicationRole {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ApplicationRole(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAppID), v...))
	})
}

// AppIDGT applies the GT predicate on the "app_id" field.
func AppIDGT(v string) predicate.ApplicationRole {
	return predicate.ApplicationRole(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAppID), v))
	})
}

// AppIDGTE applies the GTE predicate on the "app_id" field.
func AppIDGTE(v string) predicate.ApplicationRole {
	return predicate.ApplicationRole(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAppID), v))
	})
}

// AppIDLT applies the LT predicate on the "app_id" field.
func AppIDLT(v string) predicate.ApplicationRole {
	return predicate.ApplicationRole(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAppID), v))
	})
}

// AppIDLTE applies the LTE predicate on the "app_id" field.
func AppIDLTE(v string) predicate.ApplicationRole {
	return predicate.ApplicationRole(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAppID), v))
	})
}

// AppIDContains applies the Contains predicate on the "app_id" field.
func AppIDContains(v string) predicate.ApplicationRole {
	return predicate.ApplicationRole(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAppID), v))
	})
}

// AppIDHasPrefix applies the HasPrefix predicate on the "app_id" field.
func AppIDHasPrefix(v string) predicate.ApplicationRole {
	return predicate.ApplicationRole(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAppID), v))
	})
}

// AppIDHasSuffix applies the HasSuffix predicate on the "app_id" field.
func AppIDHasSuffix(v string) predicate.ApplicationRole {
	return predicate.ApplicationRole(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAppID), v))
	})
}

// AppIDEqualFold applies the EqualFold predicate on the "app_id" field.
func AppIDEqualFold(v string) predicate.ApplicationRole {
	return predicate.ApplicationRole(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAppID), v))
	})
}

// AppIDContainsFold applies the ContainsFold predicate on the "app_id" field.
func AppIDContainsFold(v string) predicate.ApplicationRole {
	return predicate.ApplicationRole(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAppID), v))
	})
}

// RoleNameEQ applies the EQ predicate on the "role_name" field.
func RoleNameEQ(v string) predicate.ApplicationRole {
	return predicate.ApplicationRole(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRoleName), v))
	})
}

// RoleNameNEQ applies the NEQ predicate on the "role_name" field.
func RoleNameNEQ(v string) predicate.ApplicationRole {
	return predicate.ApplicationRole(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRoleName), v))
	})
}

// RoleNameIn applies the In predicate on the "role_name" field.
func RoleNameIn(vs ...string) predicate.ApplicationRole {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ApplicationRole(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRoleName), v...))
	})
}

// RoleNameNotIn applies the NotIn predicate on the "role_name" field.
func RoleNameNotIn(vs ...string) predicate.ApplicationRole {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ApplicationRole(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRoleName), v...))
	})
}

// RoleNameGT applies the GT predicate on the "role_name" field.
func RoleNameGT(v string) predicate.ApplicationRole {
	return predicate.ApplicationRole(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRoleName), v))
	})
}

// RoleNameGTE applies the GTE predicate on the "role_name" field.
func RoleNameGTE(v string) predicate.ApplicationRole {
	return predicate.ApplicationRole(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRoleName), v))
	})
}

// RoleNameLT applies the LT predicate on the "role_name" field.
func RoleNameLT(v string) predicate.ApplicationRole {
	return predicate.ApplicationRole(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRoleName), v))
	})
}

// RoleNameLTE applies the LTE predicate on the "role_name" field.
func RoleNameLTE(v string) predicate.ApplicationRole {
	return predicate.ApplicationRole(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRoleName), v))
	})
}

// RoleNameContains applies the Contains predicate on the "role_name" field.
func RoleNameContains(v string) predicate.ApplicationRole {
	return predicate.ApplicationRole(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldRoleName), v))
	})
}

// RoleNameHasPrefix applies the HasPrefix predicate on the "role_name" field.
func RoleNameHasPrefix(v string) predicate.ApplicationRole {
	return predicate.ApplicationRole(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldRoleName), v))
	})
}

// RoleNameHasSuffix applies the HasSuffix predicate on the "role_name" field.
func RoleNameHasSuffix(v string) predicate.ApplicationRole {
	return predicate.ApplicationRole(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldRoleName), v))
	})
}

// RoleNameEqualFold applies the EqualFold predicate on the "role_name" field.
func RoleNameEqualFold(v string) predicate.ApplicationRole {
	return predicate.ApplicationRole(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldRoleName), v))
	})
}

// RoleNameContainsFold applies the ContainsFold predicate on the "role_name" field.
func RoleNameContainsFold(v string) predicate.ApplicationRole {
	return predicate.ApplicationRole(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldRoleName), v))
	})
}

// CreatorEQ applies the EQ predicate on the "creator" field.
func CreatorEQ(v uuid.UUID) predicate.ApplicationRole {
	return predicate.ApplicationRole(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreator), v))
	})
}

// CreatorNEQ applies the NEQ predicate on the "creator" field.
func CreatorNEQ(v uuid.UUID) predicate.ApplicationRole {
	return predicate.ApplicationRole(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreator), v))
	})
}

// CreatorIn applies the In predicate on the "creator" field.
func CreatorIn(vs ...uuid.UUID) predicate.ApplicationRole {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ApplicationRole(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreator), v...))
	})
}

// CreatorNotIn applies the NotIn predicate on the "creator" field.
func CreatorNotIn(vs ...uuid.UUID) predicate.ApplicationRole {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ApplicationRole(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreator), v...))
	})
}

// CreatorGT applies the GT predicate on the "creator" field.
func CreatorGT(v uuid.UUID) predicate.ApplicationRole {
	return predicate.ApplicationRole(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreator), v))
	})
}

// CreatorGTE applies the GTE predicate on the "creator" field.
func CreatorGTE(v uuid.UUID) predicate.ApplicationRole {
	return predicate.ApplicationRole(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreator), v))
	})
}

// CreatorLT applies the LT predicate on the "creator" field.
func CreatorLT(v uuid.UUID) predicate.ApplicationRole {
	return predicate.ApplicationRole(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreator), v))
	})
}

// CreatorLTE applies the LTE predicate on the "creator" field.
func CreatorLTE(v uuid.UUID) predicate.ApplicationRole {
	return predicate.ApplicationRole(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreator), v))
	})
}

// CreateAtEQ applies the EQ predicate on the "create_at" field.
func CreateAtEQ(v int64) predicate.ApplicationRole {
	return predicate.ApplicationRole(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateAt), v))
	})
}

// CreateAtNEQ applies the NEQ predicate on the "create_at" field.
func CreateAtNEQ(v int64) predicate.ApplicationRole {
	return predicate.ApplicationRole(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateAt), v))
	})
}

// CreateAtIn applies the In predicate on the "create_at" field.
func CreateAtIn(vs ...int64) predicate.ApplicationRole {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ApplicationRole(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreateAt), v...))
	})
}

// CreateAtNotIn applies the NotIn predicate on the "create_at" field.
func CreateAtNotIn(vs ...int64) predicate.ApplicationRole {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ApplicationRole(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreateAt), v...))
	})
}

// CreateAtGT applies the GT predicate on the "create_at" field.
func CreateAtGT(v int64) predicate.ApplicationRole {
	return predicate.ApplicationRole(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateAt), v))
	})
}

// CreateAtGTE applies the GTE predicate on the "create_at" field.
func CreateAtGTE(v int64) predicate.ApplicationRole {
	return predicate.ApplicationRole(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateAt), v))
	})
}

// CreateAtLT applies the LT predicate on the "create_at" field.
func CreateAtLT(v int64) predicate.ApplicationRole {
	return predicate.ApplicationRole(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateAt), v))
	})
}

// CreateAtLTE applies the LTE predicate on the "create_at" field.
func CreateAtLTE(v int64) predicate.ApplicationRole {
	return predicate.ApplicationRole(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateAt), v))
	})
}

// UpdateAtEQ applies the EQ predicate on the "update_at" field.
func UpdateAtEQ(v int64) predicate.ApplicationRole {
	return predicate.ApplicationRole(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtNEQ applies the NEQ predicate on the "update_at" field.
func UpdateAtNEQ(v int64) predicate.ApplicationRole {
	return predicate.ApplicationRole(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtIn applies the In predicate on the "update_at" field.
func UpdateAtIn(vs ...int64) predicate.ApplicationRole {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ApplicationRole(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdateAt), v...))
	})
}

// UpdateAtNotIn applies the NotIn predicate on the "update_at" field.
func UpdateAtNotIn(vs ...int64) predicate.ApplicationRole {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ApplicationRole(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdateAt), v...))
	})
}

// UpdateAtGT applies the GT predicate on the "update_at" field.
func UpdateAtGT(v int64) predicate.ApplicationRole {
	return predicate.ApplicationRole(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtGTE applies the GTE predicate on the "update_at" field.
func UpdateAtGTE(v int64) predicate.ApplicationRole {
	return predicate.ApplicationRole(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtLT applies the LT predicate on the "update_at" field.
func UpdateAtLT(v int64) predicate.ApplicationRole {
	return predicate.ApplicationRole(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtLTE applies the LTE predicate on the "update_at" field.
func UpdateAtLTE(v int64) predicate.ApplicationRole {
	return predicate.ApplicationRole(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateAt), v))
	})
}

// DeleteAtEQ applies the EQ predicate on the "delete_at" field.
func DeleteAtEQ(v int64) predicate.ApplicationRole {
	return predicate.ApplicationRole(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeleteAt), v))
	})
}

// DeleteAtNEQ applies the NEQ predicate on the "delete_at" field.
func DeleteAtNEQ(v int64) predicate.ApplicationRole {
	return predicate.ApplicationRole(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeleteAt), v))
	})
}

// DeleteAtIn applies the In predicate on the "delete_at" field.
func DeleteAtIn(vs ...int64) predicate.ApplicationRole {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ApplicationRole(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDeleteAt), v...))
	})
}

// DeleteAtNotIn applies the NotIn predicate on the "delete_at" field.
func DeleteAtNotIn(vs ...int64) predicate.ApplicationRole {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ApplicationRole(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDeleteAt), v...))
	})
}

// DeleteAtGT applies the GT predicate on the "delete_at" field.
func DeleteAtGT(v int64) predicate.ApplicationRole {
	return predicate.ApplicationRole(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeleteAt), v))
	})
}

// DeleteAtGTE applies the GTE predicate on the "delete_at" field.
func DeleteAtGTE(v int64) predicate.ApplicationRole {
	return predicate.ApplicationRole(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeleteAt), v))
	})
}

// DeleteAtLT applies the LT predicate on the "delete_at" field.
func DeleteAtLT(v int64) predicate.ApplicationRole {
	return predicate.ApplicationRole(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeleteAt), v))
	})
}

// DeleteAtLTE applies the LTE predicate on the "delete_at" field.
func DeleteAtLTE(v int64) predicate.ApplicationRole {
	return predicate.ApplicationRole(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeleteAt), v))
	})
}

// AnnotationEQ applies the EQ predicate on the "annotation" field.
func AnnotationEQ(v string) predicate.ApplicationRole {
	return predicate.ApplicationRole(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAnnotation), v))
	})
}

// AnnotationNEQ applies the NEQ predicate on the "annotation" field.
func AnnotationNEQ(v string) predicate.ApplicationRole {
	return predicate.ApplicationRole(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAnnotation), v))
	})
}

// AnnotationIn applies the In predicate on the "annotation" field.
func AnnotationIn(vs ...string) predicate.ApplicationRole {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ApplicationRole(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAnnotation), v...))
	})
}

// AnnotationNotIn applies the NotIn predicate on the "annotation" field.
func AnnotationNotIn(vs ...string) predicate.ApplicationRole {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ApplicationRole(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAnnotation), v...))
	})
}

// AnnotationGT applies the GT predicate on the "annotation" field.
func AnnotationGT(v string) predicate.ApplicationRole {
	return predicate.ApplicationRole(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAnnotation), v))
	})
}

// AnnotationGTE applies the GTE predicate on the "annotation" field.
func AnnotationGTE(v string) predicate.ApplicationRole {
	return predicate.ApplicationRole(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAnnotation), v))
	})
}

// AnnotationLT applies the LT predicate on the "annotation" field.
func AnnotationLT(v string) predicate.ApplicationRole {
	return predicate.ApplicationRole(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAnnotation), v))
	})
}

// AnnotationLTE applies the LTE predicate on the "annotation" field.
func AnnotationLTE(v string) predicate.ApplicationRole {
	return predicate.ApplicationRole(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAnnotation), v))
	})
}

// AnnotationContains applies the Contains predicate on the "annotation" field.
func AnnotationContains(v string) predicate.ApplicationRole {
	return predicate.ApplicationRole(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAnnotation), v))
	})
}

// AnnotationHasPrefix applies the HasPrefix predicate on the "annotation" field.
func AnnotationHasPrefix(v string) predicate.ApplicationRole {
	return predicate.ApplicationRole(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAnnotation), v))
	})
}

// AnnotationHasSuffix applies the HasSuffix predicate on the "annotation" field.
func AnnotationHasSuffix(v string) predicate.ApplicationRole {
	return predicate.ApplicationRole(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAnnotation), v))
	})
}

// AnnotationIsNil applies the IsNil predicate on the "annotation" field.
func AnnotationIsNil() predicate.ApplicationRole {
	return predicate.ApplicationRole(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldAnnotation)))
	})
}

// AnnotationNotNil applies the NotNil predicate on the "annotation" field.
func AnnotationNotNil() predicate.ApplicationRole {
	return predicate.ApplicationRole(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldAnnotation)))
	})
}

// AnnotationEqualFold applies the EqualFold predicate on the "annotation" field.
func AnnotationEqualFold(v string) predicate.ApplicationRole {
	return predicate.ApplicationRole(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAnnotation), v))
	})
}

// AnnotationContainsFold applies the ContainsFold predicate on the "annotation" field.
func AnnotationContainsFold(v string) predicate.ApplicationRole {
	return predicate.ApplicationRole(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAnnotation), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ApplicationRole) predicate.ApplicationRole {
	return predicate.ApplicationRole(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ApplicationRole) predicate.ApplicationRole {
	return predicate.ApplicationRole(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ApplicationRole) predicate.ApplicationRole {
	return predicate.ApplicationRole(func(s *sql.Selector) {
		p(s.Not())
	})
}

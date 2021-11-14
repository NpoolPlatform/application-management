// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/application-management/pkg/db/ent/applicationrole"
	"github.com/NpoolPlatform/application-management/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// ApplicationRoleUpdate is the builder for updating ApplicationRole entities.
type ApplicationRoleUpdate struct {
	config
	hooks    []Hook
	mutation *ApplicationRoleMutation
}

// Where appends a list predicates to the ApplicationRoleUpdate builder.
func (aru *ApplicationRoleUpdate) Where(ps ...predicate.ApplicationRole) *ApplicationRoleUpdate {
	aru.mutation.Where(ps...)
	return aru
}

// SetAppID sets the "app_id" field.
func (aru *ApplicationRoleUpdate) SetAppID(s string) *ApplicationRoleUpdate {
	aru.mutation.SetAppID(s)
	return aru
}

// SetRoleName sets the "role_name" field.
func (aru *ApplicationRoleUpdate) SetRoleName(s string) *ApplicationRoleUpdate {
	aru.mutation.SetRoleName(s)
	return aru
}

// SetCreator sets the "creator" field.
func (aru *ApplicationRoleUpdate) SetCreator(u uuid.UUID) *ApplicationRoleUpdate {
	aru.mutation.SetCreator(u)
	return aru
}

// SetCreateAt sets the "create_at" field.
func (aru *ApplicationRoleUpdate) SetCreateAt(u uint32) *ApplicationRoleUpdate {
	aru.mutation.ResetCreateAt()
	aru.mutation.SetCreateAt(u)
	return aru
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (aru *ApplicationRoleUpdate) SetNillableCreateAt(u *uint32) *ApplicationRoleUpdate {
	if u != nil {
		aru.SetCreateAt(*u)
	}
	return aru
}

// AddCreateAt adds u to the "create_at" field.
func (aru *ApplicationRoleUpdate) AddCreateAt(u uint32) *ApplicationRoleUpdate {
	aru.mutation.AddCreateAt(u)
	return aru
}

// SetUpdateAt sets the "update_at" field.
func (aru *ApplicationRoleUpdate) SetUpdateAt(u uint32) *ApplicationRoleUpdate {
	aru.mutation.ResetUpdateAt()
	aru.mutation.SetUpdateAt(u)
	return aru
}

// AddUpdateAt adds u to the "update_at" field.
func (aru *ApplicationRoleUpdate) AddUpdateAt(u uint32) *ApplicationRoleUpdate {
	aru.mutation.AddUpdateAt(u)
	return aru
}

// SetDeleteAt sets the "delete_at" field.
func (aru *ApplicationRoleUpdate) SetDeleteAt(u uint32) *ApplicationRoleUpdate {
	aru.mutation.ResetDeleteAt()
	aru.mutation.SetDeleteAt(u)
	return aru
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (aru *ApplicationRoleUpdate) SetNillableDeleteAt(u *uint32) *ApplicationRoleUpdate {
	if u != nil {
		aru.SetDeleteAt(*u)
	}
	return aru
}

// AddDeleteAt adds u to the "delete_at" field.
func (aru *ApplicationRoleUpdate) AddDeleteAt(u uint32) *ApplicationRoleUpdate {
	aru.mutation.AddDeleteAt(u)
	return aru
}

// SetAnnotation sets the "annotation" field.
func (aru *ApplicationRoleUpdate) SetAnnotation(s string) *ApplicationRoleUpdate {
	aru.mutation.SetAnnotation(s)
	return aru
}

// SetNillableAnnotation sets the "annotation" field if the given value is not nil.
func (aru *ApplicationRoleUpdate) SetNillableAnnotation(s *string) *ApplicationRoleUpdate {
	if s != nil {
		aru.SetAnnotation(*s)
	}
	return aru
}

// ClearAnnotation clears the value of the "annotation" field.
func (aru *ApplicationRoleUpdate) ClearAnnotation() *ApplicationRoleUpdate {
	aru.mutation.ClearAnnotation()
	return aru
}

// Mutation returns the ApplicationRoleMutation object of the builder.
func (aru *ApplicationRoleUpdate) Mutation() *ApplicationRoleMutation {
	return aru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (aru *ApplicationRoleUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	aru.defaults()
	if len(aru.hooks) == 0 {
		affected, err = aru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ApplicationRoleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			aru.mutation = mutation
			affected, err = aru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(aru.hooks) - 1; i >= 0; i-- {
			if aru.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = aru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, aru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (aru *ApplicationRoleUpdate) SaveX(ctx context.Context) int {
	affected, err := aru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (aru *ApplicationRoleUpdate) Exec(ctx context.Context) error {
	_, err := aru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aru *ApplicationRoleUpdate) ExecX(ctx context.Context) {
	if err := aru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (aru *ApplicationRoleUpdate) defaults() {
	if _, ok := aru.mutation.UpdateAt(); !ok {
		v := applicationrole.UpdateDefaultUpdateAt()
		aru.mutation.SetUpdateAt(v)
	}
}

func (aru *ApplicationRoleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   applicationrole.Table,
			Columns: applicationrole.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: applicationrole.FieldID,
			},
		},
	}
	if ps := aru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aru.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: applicationrole.FieldAppID,
		})
	}
	if value, ok := aru.mutation.RoleName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: applicationrole.FieldRoleName,
		})
	}
	if value, ok := aru.mutation.Creator(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: applicationrole.FieldCreator,
		})
	}
	if value, ok := aru.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: applicationrole.FieldCreateAt,
		})
	}
	if value, ok := aru.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: applicationrole.FieldCreateAt,
		})
	}
	if value, ok := aru.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: applicationrole.FieldUpdateAt,
		})
	}
	if value, ok := aru.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: applicationrole.FieldUpdateAt,
		})
	}
	if value, ok := aru.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: applicationrole.FieldDeleteAt,
		})
	}
	if value, ok := aru.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: applicationrole.FieldDeleteAt,
		})
	}
	if value, ok := aru.mutation.Annotation(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: applicationrole.FieldAnnotation,
		})
	}
	if aru.mutation.AnnotationCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: applicationrole.FieldAnnotation,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, aru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{applicationrole.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ApplicationRoleUpdateOne is the builder for updating a single ApplicationRole entity.
type ApplicationRoleUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ApplicationRoleMutation
}

// SetAppID sets the "app_id" field.
func (aruo *ApplicationRoleUpdateOne) SetAppID(s string) *ApplicationRoleUpdateOne {
	aruo.mutation.SetAppID(s)
	return aruo
}

// SetRoleName sets the "role_name" field.
func (aruo *ApplicationRoleUpdateOne) SetRoleName(s string) *ApplicationRoleUpdateOne {
	aruo.mutation.SetRoleName(s)
	return aruo
}

// SetCreator sets the "creator" field.
func (aruo *ApplicationRoleUpdateOne) SetCreator(u uuid.UUID) *ApplicationRoleUpdateOne {
	aruo.mutation.SetCreator(u)
	return aruo
}

// SetCreateAt sets the "create_at" field.
func (aruo *ApplicationRoleUpdateOne) SetCreateAt(u uint32) *ApplicationRoleUpdateOne {
	aruo.mutation.ResetCreateAt()
	aruo.mutation.SetCreateAt(u)
	return aruo
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (aruo *ApplicationRoleUpdateOne) SetNillableCreateAt(u *uint32) *ApplicationRoleUpdateOne {
	if u != nil {
		aruo.SetCreateAt(*u)
	}
	return aruo
}

// AddCreateAt adds u to the "create_at" field.
func (aruo *ApplicationRoleUpdateOne) AddCreateAt(u uint32) *ApplicationRoleUpdateOne {
	aruo.mutation.AddCreateAt(u)
	return aruo
}

// SetUpdateAt sets the "update_at" field.
func (aruo *ApplicationRoleUpdateOne) SetUpdateAt(u uint32) *ApplicationRoleUpdateOne {
	aruo.mutation.ResetUpdateAt()
	aruo.mutation.SetUpdateAt(u)
	return aruo
}

// AddUpdateAt adds u to the "update_at" field.
func (aruo *ApplicationRoleUpdateOne) AddUpdateAt(u uint32) *ApplicationRoleUpdateOne {
	aruo.mutation.AddUpdateAt(u)
	return aruo
}

// SetDeleteAt sets the "delete_at" field.
func (aruo *ApplicationRoleUpdateOne) SetDeleteAt(u uint32) *ApplicationRoleUpdateOne {
	aruo.mutation.ResetDeleteAt()
	aruo.mutation.SetDeleteAt(u)
	return aruo
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (aruo *ApplicationRoleUpdateOne) SetNillableDeleteAt(u *uint32) *ApplicationRoleUpdateOne {
	if u != nil {
		aruo.SetDeleteAt(*u)
	}
	return aruo
}

// AddDeleteAt adds u to the "delete_at" field.
func (aruo *ApplicationRoleUpdateOne) AddDeleteAt(u uint32) *ApplicationRoleUpdateOne {
	aruo.mutation.AddDeleteAt(u)
	return aruo
}

// SetAnnotation sets the "annotation" field.
func (aruo *ApplicationRoleUpdateOne) SetAnnotation(s string) *ApplicationRoleUpdateOne {
	aruo.mutation.SetAnnotation(s)
	return aruo
}

// SetNillableAnnotation sets the "annotation" field if the given value is not nil.
func (aruo *ApplicationRoleUpdateOne) SetNillableAnnotation(s *string) *ApplicationRoleUpdateOne {
	if s != nil {
		aruo.SetAnnotation(*s)
	}
	return aruo
}

// ClearAnnotation clears the value of the "annotation" field.
func (aruo *ApplicationRoleUpdateOne) ClearAnnotation() *ApplicationRoleUpdateOne {
	aruo.mutation.ClearAnnotation()
	return aruo
}

// Mutation returns the ApplicationRoleMutation object of the builder.
func (aruo *ApplicationRoleUpdateOne) Mutation() *ApplicationRoleMutation {
	return aruo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (aruo *ApplicationRoleUpdateOne) Select(field string, fields ...string) *ApplicationRoleUpdateOne {
	aruo.fields = append([]string{field}, fields...)
	return aruo
}

// Save executes the query and returns the updated ApplicationRole entity.
func (aruo *ApplicationRoleUpdateOne) Save(ctx context.Context) (*ApplicationRole, error) {
	var (
		err  error
		node *ApplicationRole
	)
	aruo.defaults()
	if len(aruo.hooks) == 0 {
		node, err = aruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ApplicationRoleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			aruo.mutation = mutation
			node, err = aruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(aruo.hooks) - 1; i >= 0; i-- {
			if aruo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = aruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, aruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (aruo *ApplicationRoleUpdateOne) SaveX(ctx context.Context) *ApplicationRole {
	node, err := aruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (aruo *ApplicationRoleUpdateOne) Exec(ctx context.Context) error {
	_, err := aruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aruo *ApplicationRoleUpdateOne) ExecX(ctx context.Context) {
	if err := aruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (aruo *ApplicationRoleUpdateOne) defaults() {
	if _, ok := aruo.mutation.UpdateAt(); !ok {
		v := applicationrole.UpdateDefaultUpdateAt()
		aruo.mutation.SetUpdateAt(v)
	}
}

func (aruo *ApplicationRoleUpdateOne) sqlSave(ctx context.Context) (_node *ApplicationRole, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   applicationrole.Table,
			Columns: applicationrole.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: applicationrole.FieldID,
			},
		},
	}
	id, ok := aruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing ApplicationRole.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := aruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, applicationrole.FieldID)
		for _, f := range fields {
			if !applicationrole.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != applicationrole.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := aruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aruo.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: applicationrole.FieldAppID,
		})
	}
	if value, ok := aruo.mutation.RoleName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: applicationrole.FieldRoleName,
		})
	}
	if value, ok := aruo.mutation.Creator(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: applicationrole.FieldCreator,
		})
	}
	if value, ok := aruo.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: applicationrole.FieldCreateAt,
		})
	}
	if value, ok := aruo.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: applicationrole.FieldCreateAt,
		})
	}
	if value, ok := aruo.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: applicationrole.FieldUpdateAt,
		})
	}
	if value, ok := aruo.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: applicationrole.FieldUpdateAt,
		})
	}
	if value, ok := aruo.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: applicationrole.FieldDeleteAt,
		})
	}
	if value, ok := aruo.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: applicationrole.FieldDeleteAt,
		})
	}
	if value, ok := aruo.mutation.Annotation(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: applicationrole.FieldAnnotation,
		})
	}
	if aruo.mutation.AnnotationCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: applicationrole.FieldAnnotation,
		})
	}
	_node = &ApplicationRole{config: aruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, aruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{applicationrole.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}

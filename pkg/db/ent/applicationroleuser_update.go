// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/application-management/pkg/db/ent/applicationroleuser"
	"github.com/NpoolPlatform/application-management/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// ApplicationRoleUserUpdate is the builder for updating ApplicationRoleUser entities.
type ApplicationRoleUserUpdate struct {
	config
	hooks    []Hook
	mutation *ApplicationRoleUserMutation
}

// Where appends a list predicates to the ApplicationRoleUserUpdate builder.
func (aruu *ApplicationRoleUserUpdate) Where(ps ...predicate.ApplicationRoleUser) *ApplicationRoleUserUpdate {
	aruu.mutation.Where(ps...)
	return aruu
}

// SetAppID sets the "app_id" field.
func (aruu *ApplicationRoleUserUpdate) SetAppID(s string) *ApplicationRoleUserUpdate {
	aruu.mutation.SetAppID(s)
	return aruu
}

// SetRoleID sets the "role_id" field.
func (aruu *ApplicationRoleUserUpdate) SetRoleID(u uuid.UUID) *ApplicationRoleUserUpdate {
	aruu.mutation.SetRoleID(u)
	return aruu
}

// SetUserID sets the "user_id" field.
func (aruu *ApplicationRoleUserUpdate) SetUserID(u uuid.UUID) *ApplicationRoleUserUpdate {
	aruu.mutation.SetUserID(u)
	return aruu
}

// SetCreateAt sets the "create_at" field.
func (aruu *ApplicationRoleUserUpdate) SetCreateAt(u uint32) *ApplicationRoleUserUpdate {
	aruu.mutation.ResetCreateAt()
	aruu.mutation.SetCreateAt(u)
	return aruu
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (aruu *ApplicationRoleUserUpdate) SetNillableCreateAt(u *uint32) *ApplicationRoleUserUpdate {
	if u != nil {
		aruu.SetCreateAt(*u)
	}
	return aruu
}

// AddCreateAt adds u to the "create_at" field.
func (aruu *ApplicationRoleUserUpdate) AddCreateAt(u uint32) *ApplicationRoleUserUpdate {
	aruu.mutation.AddCreateAt(u)
	return aruu
}

// SetDeleteAt sets the "delete_at" field.
func (aruu *ApplicationRoleUserUpdate) SetDeleteAt(u uint32) *ApplicationRoleUserUpdate {
	aruu.mutation.ResetDeleteAt()
	aruu.mutation.SetDeleteAt(u)
	return aruu
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (aruu *ApplicationRoleUserUpdate) SetNillableDeleteAt(u *uint32) *ApplicationRoleUserUpdate {
	if u != nil {
		aruu.SetDeleteAt(*u)
	}
	return aruu
}

// AddDeleteAt adds u to the "delete_at" field.
func (aruu *ApplicationRoleUserUpdate) AddDeleteAt(u uint32) *ApplicationRoleUserUpdate {
	aruu.mutation.AddDeleteAt(u)
	return aruu
}

// Mutation returns the ApplicationRoleUserMutation object of the builder.
func (aruu *ApplicationRoleUserUpdate) Mutation() *ApplicationRoleUserMutation {
	return aruu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (aruu *ApplicationRoleUserUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(aruu.hooks) == 0 {
		affected, err = aruu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ApplicationRoleUserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			aruu.mutation = mutation
			affected, err = aruu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(aruu.hooks) - 1; i >= 0; i-- {
			if aruu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = aruu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, aruu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (aruu *ApplicationRoleUserUpdate) SaveX(ctx context.Context) int {
	affected, err := aruu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (aruu *ApplicationRoleUserUpdate) Exec(ctx context.Context) error {
	_, err := aruu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aruu *ApplicationRoleUserUpdate) ExecX(ctx context.Context) {
	if err := aruu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (aruu *ApplicationRoleUserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   applicationroleuser.Table,
			Columns: applicationroleuser.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: applicationroleuser.FieldID,
			},
		},
	}
	if ps := aruu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aruu.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: applicationroleuser.FieldAppID,
		})
	}
	if value, ok := aruu.mutation.RoleID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: applicationroleuser.FieldRoleID,
		})
	}
	if value, ok := aruu.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: applicationroleuser.FieldUserID,
		})
	}
	if value, ok := aruu.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: applicationroleuser.FieldCreateAt,
		})
	}
	if value, ok := aruu.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: applicationroleuser.FieldCreateAt,
		})
	}
	if value, ok := aruu.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: applicationroleuser.FieldDeleteAt,
		})
	}
	if value, ok := aruu.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: applicationroleuser.FieldDeleteAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, aruu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{applicationroleuser.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ApplicationRoleUserUpdateOne is the builder for updating a single ApplicationRoleUser entity.
type ApplicationRoleUserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ApplicationRoleUserMutation
}

// SetAppID sets the "app_id" field.
func (aruuo *ApplicationRoleUserUpdateOne) SetAppID(s string) *ApplicationRoleUserUpdateOne {
	aruuo.mutation.SetAppID(s)
	return aruuo
}

// SetRoleID sets the "role_id" field.
func (aruuo *ApplicationRoleUserUpdateOne) SetRoleID(u uuid.UUID) *ApplicationRoleUserUpdateOne {
	aruuo.mutation.SetRoleID(u)
	return aruuo
}

// SetUserID sets the "user_id" field.
func (aruuo *ApplicationRoleUserUpdateOne) SetUserID(u uuid.UUID) *ApplicationRoleUserUpdateOne {
	aruuo.mutation.SetUserID(u)
	return aruuo
}

// SetCreateAt sets the "create_at" field.
func (aruuo *ApplicationRoleUserUpdateOne) SetCreateAt(u uint32) *ApplicationRoleUserUpdateOne {
	aruuo.mutation.ResetCreateAt()
	aruuo.mutation.SetCreateAt(u)
	return aruuo
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (aruuo *ApplicationRoleUserUpdateOne) SetNillableCreateAt(u *uint32) *ApplicationRoleUserUpdateOne {
	if u != nil {
		aruuo.SetCreateAt(*u)
	}
	return aruuo
}

// AddCreateAt adds u to the "create_at" field.
func (aruuo *ApplicationRoleUserUpdateOne) AddCreateAt(u uint32) *ApplicationRoleUserUpdateOne {
	aruuo.mutation.AddCreateAt(u)
	return aruuo
}

// SetDeleteAt sets the "delete_at" field.
func (aruuo *ApplicationRoleUserUpdateOne) SetDeleteAt(u uint32) *ApplicationRoleUserUpdateOne {
	aruuo.mutation.ResetDeleteAt()
	aruuo.mutation.SetDeleteAt(u)
	return aruuo
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (aruuo *ApplicationRoleUserUpdateOne) SetNillableDeleteAt(u *uint32) *ApplicationRoleUserUpdateOne {
	if u != nil {
		aruuo.SetDeleteAt(*u)
	}
	return aruuo
}

// AddDeleteAt adds u to the "delete_at" field.
func (aruuo *ApplicationRoleUserUpdateOne) AddDeleteAt(u uint32) *ApplicationRoleUserUpdateOne {
	aruuo.mutation.AddDeleteAt(u)
	return aruuo
}

// Mutation returns the ApplicationRoleUserMutation object of the builder.
func (aruuo *ApplicationRoleUserUpdateOne) Mutation() *ApplicationRoleUserMutation {
	return aruuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (aruuo *ApplicationRoleUserUpdateOne) Select(field string, fields ...string) *ApplicationRoleUserUpdateOne {
	aruuo.fields = append([]string{field}, fields...)
	return aruuo
}

// Save executes the query and returns the updated ApplicationRoleUser entity.
func (aruuo *ApplicationRoleUserUpdateOne) Save(ctx context.Context) (*ApplicationRoleUser, error) {
	var (
		err  error
		node *ApplicationRoleUser
	)
	if len(aruuo.hooks) == 0 {
		node, err = aruuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ApplicationRoleUserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			aruuo.mutation = mutation
			node, err = aruuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(aruuo.hooks) - 1; i >= 0; i-- {
			if aruuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = aruuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, aruuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (aruuo *ApplicationRoleUserUpdateOne) SaveX(ctx context.Context) *ApplicationRoleUser {
	node, err := aruuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (aruuo *ApplicationRoleUserUpdateOne) Exec(ctx context.Context) error {
	_, err := aruuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aruuo *ApplicationRoleUserUpdateOne) ExecX(ctx context.Context) {
	if err := aruuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (aruuo *ApplicationRoleUserUpdateOne) sqlSave(ctx context.Context) (_node *ApplicationRoleUser, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   applicationroleuser.Table,
			Columns: applicationroleuser.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: applicationroleuser.FieldID,
			},
		},
	}
	id, ok := aruuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing ApplicationRoleUser.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := aruuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, applicationroleuser.FieldID)
		for _, f := range fields {
			if !applicationroleuser.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != applicationroleuser.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := aruuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aruuo.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: applicationroleuser.FieldAppID,
		})
	}
	if value, ok := aruuo.mutation.RoleID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: applicationroleuser.FieldRoleID,
		})
	}
	if value, ok := aruuo.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: applicationroleuser.FieldUserID,
		})
	}
	if value, ok := aruuo.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: applicationroleuser.FieldCreateAt,
		})
	}
	if value, ok := aruuo.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: applicationroleuser.FieldCreateAt,
		})
	}
	if value, ok := aruuo.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: applicationroleuser.FieldDeleteAt,
		})
	}
	if value, ok := aruuo.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: applicationroleuser.FieldDeleteAt,
		})
	}
	_node = &ApplicationRoleUser{config: aruuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, aruuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{applicationroleuser.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}

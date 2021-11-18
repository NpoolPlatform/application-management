// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/application-management/pkg/db/ent/applicationuser"
	"github.com/NpoolPlatform/application-management/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// ApplicationUserUpdate is the builder for updating ApplicationUser entities.
type ApplicationUserUpdate struct {
	config
	hooks    []Hook
	mutation *ApplicationUserMutation
}

// Where appends a list predicates to the ApplicationUserUpdate builder.
func (auu *ApplicationUserUpdate) Where(ps ...predicate.ApplicationUser) *ApplicationUserUpdate {
	auu.mutation.Where(ps...)
	return auu
}

// SetAppID sets the "app_id" field.
func (auu *ApplicationUserUpdate) SetAppID(u uuid.UUID) *ApplicationUserUpdate {
	auu.mutation.SetAppID(u)
	return auu
}

// SetUserID sets the "user_id" field.
func (auu *ApplicationUserUpdate) SetUserID(u uuid.UUID) *ApplicationUserUpdate {
	auu.mutation.SetUserID(u)
	return auu
}

// SetOriginal sets the "original" field.
func (auu *ApplicationUserUpdate) SetOriginal(b bool) *ApplicationUserUpdate {
	auu.mutation.SetOriginal(b)
	return auu
}

// SetNillableOriginal sets the "original" field if the given value is not nil.
func (auu *ApplicationUserUpdate) SetNillableOriginal(b *bool) *ApplicationUserUpdate {
	if b != nil {
		auu.SetOriginal(*b)
	}
	return auu
}

// SetCreateAt sets the "create_at" field.
func (auu *ApplicationUserUpdate) SetCreateAt(u uint32) *ApplicationUserUpdate {
	auu.mutation.ResetCreateAt()
	auu.mutation.SetCreateAt(u)
	return auu
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (auu *ApplicationUserUpdate) SetNillableCreateAt(u *uint32) *ApplicationUserUpdate {
	if u != nil {
		auu.SetCreateAt(*u)
	}
	return auu
}

// AddCreateAt adds u to the "create_at" field.
func (auu *ApplicationUserUpdate) AddCreateAt(u uint32) *ApplicationUserUpdate {
	auu.mutation.AddCreateAt(u)
	return auu
}

// SetDeleteAt sets the "delete_at" field.
func (auu *ApplicationUserUpdate) SetDeleteAt(u uint32) *ApplicationUserUpdate {
	auu.mutation.ResetDeleteAt()
	auu.mutation.SetDeleteAt(u)
	return auu
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (auu *ApplicationUserUpdate) SetNillableDeleteAt(u *uint32) *ApplicationUserUpdate {
	if u != nil {
		auu.SetDeleteAt(*u)
	}
	return auu
}

// AddDeleteAt adds u to the "delete_at" field.
func (auu *ApplicationUserUpdate) AddDeleteAt(u uint32) *ApplicationUserUpdate {
	auu.mutation.AddDeleteAt(u)
	return auu
}

// Mutation returns the ApplicationUserMutation object of the builder.
func (auu *ApplicationUserUpdate) Mutation() *ApplicationUserMutation {
	return auu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (auu *ApplicationUserUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(auu.hooks) == 0 {
		affected, err = auu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ApplicationUserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			auu.mutation = mutation
			affected, err = auu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(auu.hooks) - 1; i >= 0; i-- {
			if auu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = auu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, auu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (auu *ApplicationUserUpdate) SaveX(ctx context.Context) int {
	affected, err := auu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (auu *ApplicationUserUpdate) Exec(ctx context.Context) error {
	_, err := auu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auu *ApplicationUserUpdate) ExecX(ctx context.Context) {
	if err := auu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (auu *ApplicationUserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   applicationuser.Table,
			Columns: applicationuser.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: applicationuser.FieldID,
			},
		},
	}
	if ps := auu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auu.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: applicationuser.FieldAppID,
		})
	}
	if value, ok := auu.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: applicationuser.FieldUserID,
		})
	}
	if value, ok := auu.mutation.Original(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: applicationuser.FieldOriginal,
		})
	}
	if value, ok := auu.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: applicationuser.FieldCreateAt,
		})
	}
	if value, ok := auu.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: applicationuser.FieldCreateAt,
		})
	}
	if value, ok := auu.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: applicationuser.FieldDeleteAt,
		})
	}
	if value, ok := auu.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: applicationuser.FieldDeleteAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, auu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{applicationuser.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ApplicationUserUpdateOne is the builder for updating a single ApplicationUser entity.
type ApplicationUserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ApplicationUserMutation
}

// SetAppID sets the "app_id" field.
func (auuo *ApplicationUserUpdateOne) SetAppID(u uuid.UUID) *ApplicationUserUpdateOne {
	auuo.mutation.SetAppID(u)
	return auuo
}

// SetUserID sets the "user_id" field.
func (auuo *ApplicationUserUpdateOne) SetUserID(u uuid.UUID) *ApplicationUserUpdateOne {
	auuo.mutation.SetUserID(u)
	return auuo
}

// SetOriginal sets the "original" field.
func (auuo *ApplicationUserUpdateOne) SetOriginal(b bool) *ApplicationUserUpdateOne {
	auuo.mutation.SetOriginal(b)
	return auuo
}

// SetNillableOriginal sets the "original" field if the given value is not nil.
func (auuo *ApplicationUserUpdateOne) SetNillableOriginal(b *bool) *ApplicationUserUpdateOne {
	if b != nil {
		auuo.SetOriginal(*b)
	}
	return auuo
}

// SetCreateAt sets the "create_at" field.
func (auuo *ApplicationUserUpdateOne) SetCreateAt(u uint32) *ApplicationUserUpdateOne {
	auuo.mutation.ResetCreateAt()
	auuo.mutation.SetCreateAt(u)
	return auuo
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (auuo *ApplicationUserUpdateOne) SetNillableCreateAt(u *uint32) *ApplicationUserUpdateOne {
	if u != nil {
		auuo.SetCreateAt(*u)
	}
	return auuo
}

// AddCreateAt adds u to the "create_at" field.
func (auuo *ApplicationUserUpdateOne) AddCreateAt(u uint32) *ApplicationUserUpdateOne {
	auuo.mutation.AddCreateAt(u)
	return auuo
}

// SetDeleteAt sets the "delete_at" field.
func (auuo *ApplicationUserUpdateOne) SetDeleteAt(u uint32) *ApplicationUserUpdateOne {
	auuo.mutation.ResetDeleteAt()
	auuo.mutation.SetDeleteAt(u)
	return auuo
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (auuo *ApplicationUserUpdateOne) SetNillableDeleteAt(u *uint32) *ApplicationUserUpdateOne {
	if u != nil {
		auuo.SetDeleteAt(*u)
	}
	return auuo
}

// AddDeleteAt adds u to the "delete_at" field.
func (auuo *ApplicationUserUpdateOne) AddDeleteAt(u uint32) *ApplicationUserUpdateOne {
	auuo.mutation.AddDeleteAt(u)
	return auuo
}

// Mutation returns the ApplicationUserMutation object of the builder.
func (auuo *ApplicationUserUpdateOne) Mutation() *ApplicationUserMutation {
	return auuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auuo *ApplicationUserUpdateOne) Select(field string, fields ...string) *ApplicationUserUpdateOne {
	auuo.fields = append([]string{field}, fields...)
	return auuo
}

// Save executes the query and returns the updated ApplicationUser entity.
func (auuo *ApplicationUserUpdateOne) Save(ctx context.Context) (*ApplicationUser, error) {
	var (
		err  error
		node *ApplicationUser
	)
	if len(auuo.hooks) == 0 {
		node, err = auuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ApplicationUserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			auuo.mutation = mutation
			node, err = auuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(auuo.hooks) - 1; i >= 0; i-- {
			if auuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = auuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, auuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (auuo *ApplicationUserUpdateOne) SaveX(ctx context.Context) *ApplicationUser {
	node, err := auuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auuo *ApplicationUserUpdateOne) Exec(ctx context.Context) error {
	_, err := auuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auuo *ApplicationUserUpdateOne) ExecX(ctx context.Context) {
	if err := auuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (auuo *ApplicationUserUpdateOne) sqlSave(ctx context.Context) (_node *ApplicationUser, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   applicationuser.Table,
			Columns: applicationuser.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: applicationuser.FieldID,
			},
		},
	}
	id, ok := auuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing ApplicationUser.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := auuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, applicationuser.FieldID)
		for _, f := range fields {
			if !applicationuser.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != applicationuser.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auuo.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: applicationuser.FieldAppID,
		})
	}
	if value, ok := auuo.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: applicationuser.FieldUserID,
		})
	}
	if value, ok := auuo.mutation.Original(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: applicationuser.FieldOriginal,
		})
	}
	if value, ok := auuo.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: applicationuser.FieldCreateAt,
		})
	}
	if value, ok := auuo.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: applicationuser.FieldCreateAt,
		})
	}
	if value, ok := auuo.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: applicationuser.FieldDeleteAt,
		})
	}
	if value, ok := auuo.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: applicationuser.FieldDeleteAt,
		})
	}
	_node = &ApplicationUser{config: auuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{applicationuser.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}

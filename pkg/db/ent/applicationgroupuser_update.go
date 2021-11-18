// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/application-management/pkg/db/ent/applicationgroupuser"
	"github.com/NpoolPlatform/application-management/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// ApplicationGroupUserUpdate is the builder for updating ApplicationGroupUser entities.
type ApplicationGroupUserUpdate struct {
	config
	hooks    []Hook
	mutation *ApplicationGroupUserMutation
}

// Where appends a list predicates to the ApplicationGroupUserUpdate builder.
func (aguu *ApplicationGroupUserUpdate) Where(ps ...predicate.ApplicationGroupUser) *ApplicationGroupUserUpdate {
	aguu.mutation.Where(ps...)
	return aguu
}

// SetGroupID sets the "group_id" field.
func (aguu *ApplicationGroupUserUpdate) SetGroupID(u uuid.UUID) *ApplicationGroupUserUpdate {
	aguu.mutation.SetGroupID(u)
	return aguu
}

// SetAppID sets the "app_id" field.
func (aguu *ApplicationGroupUserUpdate) SetAppID(u uuid.UUID) *ApplicationGroupUserUpdate {
	aguu.mutation.SetAppID(u)
	return aguu
}

// SetUserID sets the "user_id" field.
func (aguu *ApplicationGroupUserUpdate) SetUserID(u uuid.UUID) *ApplicationGroupUserUpdate {
	aguu.mutation.SetUserID(u)
	return aguu
}

// SetAnnotation sets the "annotation" field.
func (aguu *ApplicationGroupUserUpdate) SetAnnotation(s string) *ApplicationGroupUserUpdate {
	aguu.mutation.SetAnnotation(s)
	return aguu
}

// SetNillableAnnotation sets the "annotation" field if the given value is not nil.
func (aguu *ApplicationGroupUserUpdate) SetNillableAnnotation(s *string) *ApplicationGroupUserUpdate {
	if s != nil {
		aguu.SetAnnotation(*s)
	}
	return aguu
}

// ClearAnnotation clears the value of the "annotation" field.
func (aguu *ApplicationGroupUserUpdate) ClearAnnotation() *ApplicationGroupUserUpdate {
	aguu.mutation.ClearAnnotation()
	return aguu
}

// SetCreateAt sets the "create_at" field.
func (aguu *ApplicationGroupUserUpdate) SetCreateAt(u uint32) *ApplicationGroupUserUpdate {
	aguu.mutation.ResetCreateAt()
	aguu.mutation.SetCreateAt(u)
	return aguu
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (aguu *ApplicationGroupUserUpdate) SetNillableCreateAt(u *uint32) *ApplicationGroupUserUpdate {
	if u != nil {
		aguu.SetCreateAt(*u)
	}
	return aguu
}

// AddCreateAt adds u to the "create_at" field.
func (aguu *ApplicationGroupUserUpdate) AddCreateAt(u uint32) *ApplicationGroupUserUpdate {
	aguu.mutation.AddCreateAt(u)
	return aguu
}

// SetDeleteAt sets the "delete_at" field.
func (aguu *ApplicationGroupUserUpdate) SetDeleteAt(u uint32) *ApplicationGroupUserUpdate {
	aguu.mutation.ResetDeleteAt()
	aguu.mutation.SetDeleteAt(u)
	return aguu
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (aguu *ApplicationGroupUserUpdate) SetNillableDeleteAt(u *uint32) *ApplicationGroupUserUpdate {
	if u != nil {
		aguu.SetDeleteAt(*u)
	}
	return aguu
}

// AddDeleteAt adds u to the "delete_at" field.
func (aguu *ApplicationGroupUserUpdate) AddDeleteAt(u uint32) *ApplicationGroupUserUpdate {
	aguu.mutation.AddDeleteAt(u)
	return aguu
}

// Mutation returns the ApplicationGroupUserMutation object of the builder.
func (aguu *ApplicationGroupUserUpdate) Mutation() *ApplicationGroupUserMutation {
	return aguu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (aguu *ApplicationGroupUserUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(aguu.hooks) == 0 {
		affected, err = aguu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ApplicationGroupUserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			aguu.mutation = mutation
			affected, err = aguu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(aguu.hooks) - 1; i >= 0; i-- {
			if aguu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = aguu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, aguu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (aguu *ApplicationGroupUserUpdate) SaveX(ctx context.Context) int {
	affected, err := aguu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (aguu *ApplicationGroupUserUpdate) Exec(ctx context.Context) error {
	_, err := aguu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aguu *ApplicationGroupUserUpdate) ExecX(ctx context.Context) {
	if err := aguu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (aguu *ApplicationGroupUserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   applicationgroupuser.Table,
			Columns: applicationgroupuser.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: applicationgroupuser.FieldID,
			},
		},
	}
	if ps := aguu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aguu.mutation.GroupID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: applicationgroupuser.FieldGroupID,
		})
	}
	if value, ok := aguu.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: applicationgroupuser.FieldAppID,
		})
	}
	if value, ok := aguu.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: applicationgroupuser.FieldUserID,
		})
	}
	if value, ok := aguu.mutation.Annotation(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: applicationgroupuser.FieldAnnotation,
		})
	}
	if aguu.mutation.AnnotationCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: applicationgroupuser.FieldAnnotation,
		})
	}
	if value, ok := aguu.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: applicationgroupuser.FieldCreateAt,
		})
	}
	if value, ok := aguu.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: applicationgroupuser.FieldCreateAt,
		})
	}
	if value, ok := aguu.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: applicationgroupuser.FieldDeleteAt,
		})
	}
	if value, ok := aguu.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: applicationgroupuser.FieldDeleteAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, aguu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{applicationgroupuser.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ApplicationGroupUserUpdateOne is the builder for updating a single ApplicationGroupUser entity.
type ApplicationGroupUserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ApplicationGroupUserMutation
}

// SetGroupID sets the "group_id" field.
func (aguuo *ApplicationGroupUserUpdateOne) SetGroupID(u uuid.UUID) *ApplicationGroupUserUpdateOne {
	aguuo.mutation.SetGroupID(u)
	return aguuo
}

// SetAppID sets the "app_id" field.
func (aguuo *ApplicationGroupUserUpdateOne) SetAppID(u uuid.UUID) *ApplicationGroupUserUpdateOne {
	aguuo.mutation.SetAppID(u)
	return aguuo
}

// SetUserID sets the "user_id" field.
func (aguuo *ApplicationGroupUserUpdateOne) SetUserID(u uuid.UUID) *ApplicationGroupUserUpdateOne {
	aguuo.mutation.SetUserID(u)
	return aguuo
}

// SetAnnotation sets the "annotation" field.
func (aguuo *ApplicationGroupUserUpdateOne) SetAnnotation(s string) *ApplicationGroupUserUpdateOne {
	aguuo.mutation.SetAnnotation(s)
	return aguuo
}

// SetNillableAnnotation sets the "annotation" field if the given value is not nil.
func (aguuo *ApplicationGroupUserUpdateOne) SetNillableAnnotation(s *string) *ApplicationGroupUserUpdateOne {
	if s != nil {
		aguuo.SetAnnotation(*s)
	}
	return aguuo
}

// ClearAnnotation clears the value of the "annotation" field.
func (aguuo *ApplicationGroupUserUpdateOne) ClearAnnotation() *ApplicationGroupUserUpdateOne {
	aguuo.mutation.ClearAnnotation()
	return aguuo
}

// SetCreateAt sets the "create_at" field.
func (aguuo *ApplicationGroupUserUpdateOne) SetCreateAt(u uint32) *ApplicationGroupUserUpdateOne {
	aguuo.mutation.ResetCreateAt()
	aguuo.mutation.SetCreateAt(u)
	return aguuo
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (aguuo *ApplicationGroupUserUpdateOne) SetNillableCreateAt(u *uint32) *ApplicationGroupUserUpdateOne {
	if u != nil {
		aguuo.SetCreateAt(*u)
	}
	return aguuo
}

// AddCreateAt adds u to the "create_at" field.
func (aguuo *ApplicationGroupUserUpdateOne) AddCreateAt(u uint32) *ApplicationGroupUserUpdateOne {
	aguuo.mutation.AddCreateAt(u)
	return aguuo
}

// SetDeleteAt sets the "delete_at" field.
func (aguuo *ApplicationGroupUserUpdateOne) SetDeleteAt(u uint32) *ApplicationGroupUserUpdateOne {
	aguuo.mutation.ResetDeleteAt()
	aguuo.mutation.SetDeleteAt(u)
	return aguuo
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (aguuo *ApplicationGroupUserUpdateOne) SetNillableDeleteAt(u *uint32) *ApplicationGroupUserUpdateOne {
	if u != nil {
		aguuo.SetDeleteAt(*u)
	}
	return aguuo
}

// AddDeleteAt adds u to the "delete_at" field.
func (aguuo *ApplicationGroupUserUpdateOne) AddDeleteAt(u uint32) *ApplicationGroupUserUpdateOne {
	aguuo.mutation.AddDeleteAt(u)
	return aguuo
}

// Mutation returns the ApplicationGroupUserMutation object of the builder.
func (aguuo *ApplicationGroupUserUpdateOne) Mutation() *ApplicationGroupUserMutation {
	return aguuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (aguuo *ApplicationGroupUserUpdateOne) Select(field string, fields ...string) *ApplicationGroupUserUpdateOne {
	aguuo.fields = append([]string{field}, fields...)
	return aguuo
}

// Save executes the query and returns the updated ApplicationGroupUser entity.
func (aguuo *ApplicationGroupUserUpdateOne) Save(ctx context.Context) (*ApplicationGroupUser, error) {
	var (
		err  error
		node *ApplicationGroupUser
	)
	if len(aguuo.hooks) == 0 {
		node, err = aguuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ApplicationGroupUserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			aguuo.mutation = mutation
			node, err = aguuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(aguuo.hooks) - 1; i >= 0; i-- {
			if aguuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = aguuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, aguuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (aguuo *ApplicationGroupUserUpdateOne) SaveX(ctx context.Context) *ApplicationGroupUser {
	node, err := aguuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (aguuo *ApplicationGroupUserUpdateOne) Exec(ctx context.Context) error {
	_, err := aguuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aguuo *ApplicationGroupUserUpdateOne) ExecX(ctx context.Context) {
	if err := aguuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (aguuo *ApplicationGroupUserUpdateOne) sqlSave(ctx context.Context) (_node *ApplicationGroupUser, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   applicationgroupuser.Table,
			Columns: applicationgroupuser.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: applicationgroupuser.FieldID,
			},
		},
	}
	id, ok := aguuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing ApplicationGroupUser.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := aguuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, applicationgroupuser.FieldID)
		for _, f := range fields {
			if !applicationgroupuser.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != applicationgroupuser.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := aguuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aguuo.mutation.GroupID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: applicationgroupuser.FieldGroupID,
		})
	}
	if value, ok := aguuo.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: applicationgroupuser.FieldAppID,
		})
	}
	if value, ok := aguuo.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: applicationgroupuser.FieldUserID,
		})
	}
	if value, ok := aguuo.mutation.Annotation(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: applicationgroupuser.FieldAnnotation,
		})
	}
	if aguuo.mutation.AnnotationCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: applicationgroupuser.FieldAnnotation,
		})
	}
	if value, ok := aguuo.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: applicationgroupuser.FieldCreateAt,
		})
	}
	if value, ok := aguuo.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: applicationgroupuser.FieldCreateAt,
		})
	}
	if value, ok := aguuo.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: applicationgroupuser.FieldDeleteAt,
		})
	}
	if value, ok := aguuo.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: applicationgroupuser.FieldDeleteAt,
		})
	}
	_node = &ApplicationGroupUser{config: aguuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, aguuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{applicationgroupuser.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}

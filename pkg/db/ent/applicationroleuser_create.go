// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/application-management/pkg/db/ent/applicationroleuser"
	"github.com/google/uuid"
)

// ApplicationRoleUserCreate is the builder for creating a ApplicationRoleUser entity.
type ApplicationRoleUserCreate struct {
	config
	mutation *ApplicationRoleUserMutation
	hooks    []Hook
}

// SetAppID sets the "app_id" field.
func (aruc *ApplicationRoleUserCreate) SetAppID(u uuid.UUID) *ApplicationRoleUserCreate {
	aruc.mutation.SetAppID(u)
	return aruc
}

// SetRoleID sets the "role_id" field.
func (aruc *ApplicationRoleUserCreate) SetRoleID(u uuid.UUID) *ApplicationRoleUserCreate {
	aruc.mutation.SetRoleID(u)
	return aruc
}

// SetUserID sets the "user_id" field.
func (aruc *ApplicationRoleUserCreate) SetUserID(u uuid.UUID) *ApplicationRoleUserCreate {
	aruc.mutation.SetUserID(u)
	return aruc
}

// SetCreateAt sets the "create_at" field.
func (aruc *ApplicationRoleUserCreate) SetCreateAt(u uint32) *ApplicationRoleUserCreate {
	aruc.mutation.SetCreateAt(u)
	return aruc
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (aruc *ApplicationRoleUserCreate) SetNillableCreateAt(u *uint32) *ApplicationRoleUserCreate {
	if u != nil {
		aruc.SetCreateAt(*u)
	}
	return aruc
}

// SetDeleteAt sets the "delete_at" field.
func (aruc *ApplicationRoleUserCreate) SetDeleteAt(u uint32) *ApplicationRoleUserCreate {
	aruc.mutation.SetDeleteAt(u)
	return aruc
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (aruc *ApplicationRoleUserCreate) SetNillableDeleteAt(u *uint32) *ApplicationRoleUserCreate {
	if u != nil {
		aruc.SetDeleteAt(*u)
	}
	return aruc
}

// SetID sets the "id" field.
func (aruc *ApplicationRoleUserCreate) SetID(u uuid.UUID) *ApplicationRoleUserCreate {
	aruc.mutation.SetID(u)
	return aruc
}

// Mutation returns the ApplicationRoleUserMutation object of the builder.
func (aruc *ApplicationRoleUserCreate) Mutation() *ApplicationRoleUserMutation {
	return aruc.mutation
}

// Save creates the ApplicationRoleUser in the database.
func (aruc *ApplicationRoleUserCreate) Save(ctx context.Context) (*ApplicationRoleUser, error) {
	var (
		err  error
		node *ApplicationRoleUser
	)
	aruc.defaults()
	if len(aruc.hooks) == 0 {
		if err = aruc.check(); err != nil {
			return nil, err
		}
		node, err = aruc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ApplicationRoleUserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = aruc.check(); err != nil {
				return nil, err
			}
			aruc.mutation = mutation
			if node, err = aruc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(aruc.hooks) - 1; i >= 0; i-- {
			if aruc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = aruc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, aruc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (aruc *ApplicationRoleUserCreate) SaveX(ctx context.Context) *ApplicationRoleUser {
	v, err := aruc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (aruc *ApplicationRoleUserCreate) Exec(ctx context.Context) error {
	_, err := aruc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aruc *ApplicationRoleUserCreate) ExecX(ctx context.Context) {
	if err := aruc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (aruc *ApplicationRoleUserCreate) defaults() {
	if _, ok := aruc.mutation.CreateAt(); !ok {
		v := applicationroleuser.DefaultCreateAt()
		aruc.mutation.SetCreateAt(v)
	}
	if _, ok := aruc.mutation.DeleteAt(); !ok {
		v := applicationroleuser.DefaultDeleteAt()
		aruc.mutation.SetDeleteAt(v)
	}
	if _, ok := aruc.mutation.ID(); !ok {
		v := applicationroleuser.DefaultID()
		aruc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (aruc *ApplicationRoleUserCreate) check() error {
	if _, ok := aruc.mutation.AppID(); !ok {
		return &ValidationError{Name: "app_id", err: errors.New(`ent: missing required field "app_id"`)}
	}
	if _, ok := aruc.mutation.RoleID(); !ok {
		return &ValidationError{Name: "role_id", err: errors.New(`ent: missing required field "role_id"`)}
	}
	if _, ok := aruc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "user_id"`)}
	}
	if _, ok := aruc.mutation.CreateAt(); !ok {
		return &ValidationError{Name: "create_at", err: errors.New(`ent: missing required field "create_at"`)}
	}
	if _, ok := aruc.mutation.DeleteAt(); !ok {
		return &ValidationError{Name: "delete_at", err: errors.New(`ent: missing required field "delete_at"`)}
	}
	return nil
}

func (aruc *ApplicationRoleUserCreate) sqlSave(ctx context.Context) (*ApplicationRoleUser, error) {
	_node, _spec := aruc.createSpec()
	if err := sqlgraph.CreateNode(ctx, aruc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		_node.ID = _spec.ID.Value.(uuid.UUID)
	}
	return _node, nil
}

func (aruc *ApplicationRoleUserCreate) createSpec() (*ApplicationRoleUser, *sqlgraph.CreateSpec) {
	var (
		_node = &ApplicationRoleUser{config: aruc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: applicationroleuser.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: applicationroleuser.FieldID,
			},
		}
	)
	if id, ok := aruc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := aruc.mutation.AppID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: applicationroleuser.FieldAppID,
		})
		_node.AppID = value
	}
	if value, ok := aruc.mutation.RoleID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: applicationroleuser.FieldRoleID,
		})
		_node.RoleID = value
	}
	if value, ok := aruc.mutation.UserID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: applicationroleuser.FieldUserID,
		})
		_node.UserID = value
	}
	if value, ok := aruc.mutation.CreateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: applicationroleuser.FieldCreateAt,
		})
		_node.CreateAt = value
	}
	if value, ok := aruc.mutation.DeleteAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: applicationroleuser.FieldDeleteAt,
		})
		_node.DeleteAt = value
	}
	return _node, _spec
}

// ApplicationRoleUserCreateBulk is the builder for creating many ApplicationRoleUser entities in bulk.
type ApplicationRoleUserCreateBulk struct {
	config
	builders []*ApplicationRoleUserCreate
}

// Save creates the ApplicationRoleUser entities in the database.
func (arucb *ApplicationRoleUserCreateBulk) Save(ctx context.Context) ([]*ApplicationRoleUser, error) {
	specs := make([]*sqlgraph.CreateSpec, len(arucb.builders))
	nodes := make([]*ApplicationRoleUser, len(arucb.builders))
	mutators := make([]Mutator, len(arucb.builders))
	for i := range arucb.builders {
		func(i int, root context.Context) {
			builder := arucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ApplicationRoleUserMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, arucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, arucb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, arucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (arucb *ApplicationRoleUserCreateBulk) SaveX(ctx context.Context) []*ApplicationRoleUser {
	v, err := arucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (arucb *ApplicationRoleUserCreateBulk) Exec(ctx context.Context) error {
	_, err := arucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (arucb *ApplicationRoleUserCreateBulk) ExecX(ctx context.Context) {
	if err := arucb.Exec(ctx); err != nil {
		panic(err)
	}
}

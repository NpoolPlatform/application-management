// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/application-management/pkg/db/ent/applicationrole"
	"github.com/google/uuid"
)

// ApplicationRoleCreate is the builder for creating a ApplicationRole entity.
type ApplicationRoleCreate struct {
	config
	mutation *ApplicationRoleMutation
	hooks    []Hook
}

// SetAppID sets the "app_id" field.
func (arc *ApplicationRoleCreate) SetAppID(s string) *ApplicationRoleCreate {
	arc.mutation.SetAppID(s)
	return arc
}

// SetRoleName sets the "role_name" field.
func (arc *ApplicationRoleCreate) SetRoleName(s string) *ApplicationRoleCreate {
	arc.mutation.SetRoleName(s)
	return arc
}

// SetCreator sets the "creator" field.
func (arc *ApplicationRoleCreate) SetCreator(u uuid.UUID) *ApplicationRoleCreate {
	arc.mutation.SetCreator(u)
	return arc
}

// SetCreateAt sets the "create_at" field.
func (arc *ApplicationRoleCreate) SetCreateAt(u uint32) *ApplicationRoleCreate {
	arc.mutation.SetCreateAt(u)
	return arc
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (arc *ApplicationRoleCreate) SetNillableCreateAt(u *uint32) *ApplicationRoleCreate {
	if u != nil {
		arc.SetCreateAt(*u)
	}
	return arc
}

// SetUpdateAt sets the "update_at" field.
func (arc *ApplicationRoleCreate) SetUpdateAt(u uint32) *ApplicationRoleCreate {
	arc.mutation.SetUpdateAt(u)
	return arc
}

// SetNillableUpdateAt sets the "update_at" field if the given value is not nil.
func (arc *ApplicationRoleCreate) SetNillableUpdateAt(u *uint32) *ApplicationRoleCreate {
	if u != nil {
		arc.SetUpdateAt(*u)
	}
	return arc
}

// SetDeleteAt sets the "delete_at" field.
func (arc *ApplicationRoleCreate) SetDeleteAt(u uint32) *ApplicationRoleCreate {
	arc.mutation.SetDeleteAt(u)
	return arc
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (arc *ApplicationRoleCreate) SetNillableDeleteAt(u *uint32) *ApplicationRoleCreate {
	if u != nil {
		arc.SetDeleteAt(*u)
	}
	return arc
}

// SetAnnotation sets the "annotation" field.
func (arc *ApplicationRoleCreate) SetAnnotation(s string) *ApplicationRoleCreate {
	arc.mutation.SetAnnotation(s)
	return arc
}

// SetNillableAnnotation sets the "annotation" field if the given value is not nil.
func (arc *ApplicationRoleCreate) SetNillableAnnotation(s *string) *ApplicationRoleCreate {
	if s != nil {
		arc.SetAnnotation(*s)
	}
	return arc
}

// SetID sets the "id" field.
func (arc *ApplicationRoleCreate) SetID(u uuid.UUID) *ApplicationRoleCreate {
	arc.mutation.SetID(u)
	return arc
}

// Mutation returns the ApplicationRoleMutation object of the builder.
func (arc *ApplicationRoleCreate) Mutation() *ApplicationRoleMutation {
	return arc.mutation
}

// Save creates the ApplicationRole in the database.
func (arc *ApplicationRoleCreate) Save(ctx context.Context) (*ApplicationRole, error) {
	var (
		err  error
		node *ApplicationRole
	)
	arc.defaults()
	if len(arc.hooks) == 0 {
		if err = arc.check(); err != nil {
			return nil, err
		}
		node, err = arc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ApplicationRoleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = arc.check(); err != nil {
				return nil, err
			}
			arc.mutation = mutation
			if node, err = arc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(arc.hooks) - 1; i >= 0; i-- {
			if arc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = arc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, arc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (arc *ApplicationRoleCreate) SaveX(ctx context.Context) *ApplicationRole {
	v, err := arc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (arc *ApplicationRoleCreate) Exec(ctx context.Context) error {
	_, err := arc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (arc *ApplicationRoleCreate) ExecX(ctx context.Context) {
	if err := arc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (arc *ApplicationRoleCreate) defaults() {
	if _, ok := arc.mutation.CreateAt(); !ok {
		v := applicationrole.DefaultCreateAt()
		arc.mutation.SetCreateAt(v)
	}
	if _, ok := arc.mutation.UpdateAt(); !ok {
		v := applicationrole.DefaultUpdateAt()
		arc.mutation.SetUpdateAt(v)
	}
	if _, ok := arc.mutation.DeleteAt(); !ok {
		v := applicationrole.DefaultDeleteAt()
		arc.mutation.SetDeleteAt(v)
	}
	if _, ok := arc.mutation.ID(); !ok {
		v := applicationrole.DefaultID()
		arc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (arc *ApplicationRoleCreate) check() error {
	if _, ok := arc.mutation.AppID(); !ok {
		return &ValidationError{Name: "app_id", err: errors.New(`ent: missing required field "app_id"`)}
	}
	if _, ok := arc.mutation.RoleName(); !ok {
		return &ValidationError{Name: "role_name", err: errors.New(`ent: missing required field "role_name"`)}
	}
	if _, ok := arc.mutation.Creator(); !ok {
		return &ValidationError{Name: "creator", err: errors.New(`ent: missing required field "creator"`)}
	}
	if _, ok := arc.mutation.CreateAt(); !ok {
		return &ValidationError{Name: "create_at", err: errors.New(`ent: missing required field "create_at"`)}
	}
	if _, ok := arc.mutation.UpdateAt(); !ok {
		return &ValidationError{Name: "update_at", err: errors.New(`ent: missing required field "update_at"`)}
	}
	if _, ok := arc.mutation.DeleteAt(); !ok {
		return &ValidationError{Name: "delete_at", err: errors.New(`ent: missing required field "delete_at"`)}
	}
	return nil
}

func (arc *ApplicationRoleCreate) sqlSave(ctx context.Context) (*ApplicationRole, error) {
	_node, _spec := arc.createSpec()
	if err := sqlgraph.CreateNode(ctx, arc.driver, _spec); err != nil {
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

func (arc *ApplicationRoleCreate) createSpec() (*ApplicationRole, *sqlgraph.CreateSpec) {
	var (
		_node = &ApplicationRole{config: arc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: applicationrole.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: applicationrole.FieldID,
			},
		}
	)
	if id, ok := arc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := arc.mutation.AppID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: applicationrole.FieldAppID,
		})
		_node.AppID = value
	}
	if value, ok := arc.mutation.RoleName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: applicationrole.FieldRoleName,
		})
		_node.RoleName = value
	}
	if value, ok := arc.mutation.Creator(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: applicationrole.FieldCreator,
		})
		_node.Creator = value
	}
	if value, ok := arc.mutation.CreateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: applicationrole.FieldCreateAt,
		})
		_node.CreateAt = value
	}
	if value, ok := arc.mutation.UpdateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: applicationrole.FieldUpdateAt,
		})
		_node.UpdateAt = value
	}
	if value, ok := arc.mutation.DeleteAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: applicationrole.FieldDeleteAt,
		})
		_node.DeleteAt = value
	}
	if value, ok := arc.mutation.Annotation(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: applicationrole.FieldAnnotation,
		})
		_node.Annotation = value
	}
	return _node, _spec
}

// ApplicationRoleCreateBulk is the builder for creating many ApplicationRole entities in bulk.
type ApplicationRoleCreateBulk struct {
	config
	builders []*ApplicationRoleCreate
}

// Save creates the ApplicationRole entities in the database.
func (arcb *ApplicationRoleCreateBulk) Save(ctx context.Context) ([]*ApplicationRole, error) {
	specs := make([]*sqlgraph.CreateSpec, len(arcb.builders))
	nodes := make([]*ApplicationRole, len(arcb.builders))
	mutators := make([]Mutator, len(arcb.builders))
	for i := range arcb.builders {
		func(i int, root context.Context) {
			builder := arcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ApplicationRoleMutation)
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
					_, err = mutators[i+1].Mutate(root, arcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, arcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, arcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (arcb *ApplicationRoleCreateBulk) SaveX(ctx context.Context) []*ApplicationRole {
	v, err := arcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (arcb *ApplicationRoleCreateBulk) Exec(ctx context.Context) error {
	_, err := arcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (arcb *ApplicationRoleCreateBulk) ExecX(ctx context.Context) {
	if err := arcb.Exec(ctx); err != nil {
		panic(err)
	}
}

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
)

// ApplicationGroupUserDelete is the builder for deleting a ApplicationGroupUser entity.
type ApplicationGroupUserDelete struct {
	config
	hooks    []Hook
	mutation *ApplicationGroupUserMutation
}

// Where appends a list predicates to the ApplicationGroupUserDelete builder.
func (agud *ApplicationGroupUserDelete) Where(ps ...predicate.ApplicationGroupUser) *ApplicationGroupUserDelete {
	agud.mutation.Where(ps...)
	return agud
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (agud *ApplicationGroupUserDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(agud.hooks) == 0 {
		affected, err = agud.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ApplicationGroupUserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			agud.mutation = mutation
			affected, err = agud.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(agud.hooks) - 1; i >= 0; i-- {
			if agud.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = agud.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, agud.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (agud *ApplicationGroupUserDelete) ExecX(ctx context.Context) int {
	n, err := agud.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (agud *ApplicationGroupUserDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: applicationgroupuser.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: applicationgroupuser.FieldID,
			},
		},
	}
	if ps := agud.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, agud.driver, _spec)
}

// ApplicationGroupUserDeleteOne is the builder for deleting a single ApplicationGroupUser entity.
type ApplicationGroupUserDeleteOne struct {
	agud *ApplicationGroupUserDelete
}

// Exec executes the deletion query.
func (agudo *ApplicationGroupUserDeleteOne) Exec(ctx context.Context) error {
	n, err := agudo.agud.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{applicationgroupuser.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (agudo *ApplicationGroupUserDeleteOne) ExecX(ctx context.Context) {
	agudo.agud.ExecX(ctx)
}

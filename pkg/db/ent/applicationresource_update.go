// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/application-management/pkg/db/ent/applicationresource"
	"github.com/NpoolPlatform/application-management/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// ApplicationResourceUpdate is the builder for updating ApplicationResource entities.
type ApplicationResourceUpdate struct {
	config
	hooks    []Hook
	mutation *ApplicationResourceMutation
}

// Where appends a list predicates to the ApplicationResourceUpdate builder.
func (aru *ApplicationResourceUpdate) Where(ps ...predicate.ApplicationResource) *ApplicationResourceUpdate {
	aru.mutation.Where(ps...)
	return aru
}

// SetAppID sets the "app_id" field.
func (aru *ApplicationResourceUpdate) SetAppID(s string) *ApplicationResourceUpdate {
	aru.mutation.SetAppID(s)
	return aru
}

// SetResourceName sets the "resource_name" field.
func (aru *ApplicationResourceUpdate) SetResourceName(s string) *ApplicationResourceUpdate {
	aru.mutation.SetResourceName(s)
	return aru
}

// SetResourceDescription sets the "resource_description" field.
func (aru *ApplicationResourceUpdate) SetResourceDescription(s string) *ApplicationResourceUpdate {
	aru.mutation.SetResourceDescription(s)
	return aru
}

// SetNillableResourceDescription sets the "resource_description" field if the given value is not nil.
func (aru *ApplicationResourceUpdate) SetNillableResourceDescription(s *string) *ApplicationResourceUpdate {
	if s != nil {
		aru.SetResourceDescription(*s)
	}
	return aru
}

// ClearResourceDescription clears the value of the "resource_description" field.
func (aru *ApplicationResourceUpdate) ClearResourceDescription() *ApplicationResourceUpdate {
	aru.mutation.ClearResourceDescription()
	return aru
}

// SetType sets the "type" field.
func (aru *ApplicationResourceUpdate) SetType(s string) *ApplicationResourceUpdate {
	aru.mutation.SetType(s)
	return aru
}

// SetNillableType sets the "type" field if the given value is not nil.
func (aru *ApplicationResourceUpdate) SetNillableType(s *string) *ApplicationResourceUpdate {
	if s != nil {
		aru.SetType(*s)
	}
	return aru
}

// SetCreator sets the "creator" field.
func (aru *ApplicationResourceUpdate) SetCreator(u uuid.UUID) *ApplicationResourceUpdate {
	aru.mutation.SetCreator(u)
	return aru
}

// SetCreateAt sets the "create_at" field.
func (aru *ApplicationResourceUpdate) SetCreateAt(u uint32) *ApplicationResourceUpdate {
	aru.mutation.ResetCreateAt()
	aru.mutation.SetCreateAt(u)
	return aru
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (aru *ApplicationResourceUpdate) SetNillableCreateAt(u *uint32) *ApplicationResourceUpdate {
	if u != nil {
		aru.SetCreateAt(*u)
	}
	return aru
}

// AddCreateAt adds u to the "create_at" field.
func (aru *ApplicationResourceUpdate) AddCreateAt(u uint32) *ApplicationResourceUpdate {
	aru.mutation.AddCreateAt(u)
	return aru
}

// SetUpdateAt sets the "update_at" field.
func (aru *ApplicationResourceUpdate) SetUpdateAt(u uint32) *ApplicationResourceUpdate {
	aru.mutation.ResetUpdateAt()
	aru.mutation.SetUpdateAt(u)
	return aru
}

// AddUpdateAt adds u to the "update_at" field.
func (aru *ApplicationResourceUpdate) AddUpdateAt(u uint32) *ApplicationResourceUpdate {
	aru.mutation.AddUpdateAt(u)
	return aru
}

// SetDeleteAt sets the "delete_at" field.
func (aru *ApplicationResourceUpdate) SetDeleteAt(u uint32) *ApplicationResourceUpdate {
	aru.mutation.ResetDeleteAt()
	aru.mutation.SetDeleteAt(u)
	return aru
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (aru *ApplicationResourceUpdate) SetNillableDeleteAt(u *uint32) *ApplicationResourceUpdate {
	if u != nil {
		aru.SetDeleteAt(*u)
	}
	return aru
}

// AddDeleteAt adds u to the "delete_at" field.
func (aru *ApplicationResourceUpdate) AddDeleteAt(u uint32) *ApplicationResourceUpdate {
	aru.mutation.AddDeleteAt(u)
	return aru
}

// Mutation returns the ApplicationResourceMutation object of the builder.
func (aru *ApplicationResourceUpdate) Mutation() *ApplicationResourceMutation {
	return aru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (aru *ApplicationResourceUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	aru.defaults()
	if len(aru.hooks) == 0 {
		affected, err = aru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ApplicationResourceMutation)
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
func (aru *ApplicationResourceUpdate) SaveX(ctx context.Context) int {
	affected, err := aru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (aru *ApplicationResourceUpdate) Exec(ctx context.Context) error {
	_, err := aru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aru *ApplicationResourceUpdate) ExecX(ctx context.Context) {
	if err := aru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (aru *ApplicationResourceUpdate) defaults() {
	if _, ok := aru.mutation.UpdateAt(); !ok {
		v := applicationresource.UpdateDefaultUpdateAt()
		aru.mutation.SetUpdateAt(v)
	}
}

func (aru *ApplicationResourceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   applicationresource.Table,
			Columns: applicationresource.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: applicationresource.FieldID,
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
			Column: applicationresource.FieldAppID,
		})
	}
	if value, ok := aru.mutation.ResourceName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: applicationresource.FieldResourceName,
		})
	}
	if value, ok := aru.mutation.ResourceDescription(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: applicationresource.FieldResourceDescription,
		})
	}
	if aru.mutation.ResourceDescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: applicationresource.FieldResourceDescription,
		})
	}
	if value, ok := aru.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: applicationresource.FieldType,
		})
	}
	if value, ok := aru.mutation.Creator(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: applicationresource.FieldCreator,
		})
	}
	if value, ok := aru.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: applicationresource.FieldCreateAt,
		})
	}
	if value, ok := aru.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: applicationresource.FieldCreateAt,
		})
	}
	if value, ok := aru.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: applicationresource.FieldUpdateAt,
		})
	}
	if value, ok := aru.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: applicationresource.FieldUpdateAt,
		})
	}
	if value, ok := aru.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: applicationresource.FieldDeleteAt,
		})
	}
	if value, ok := aru.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: applicationresource.FieldDeleteAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, aru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{applicationresource.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ApplicationResourceUpdateOne is the builder for updating a single ApplicationResource entity.
type ApplicationResourceUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ApplicationResourceMutation
}

// SetAppID sets the "app_id" field.
func (aruo *ApplicationResourceUpdateOne) SetAppID(s string) *ApplicationResourceUpdateOne {
	aruo.mutation.SetAppID(s)
	return aruo
}

// SetResourceName sets the "resource_name" field.
func (aruo *ApplicationResourceUpdateOne) SetResourceName(s string) *ApplicationResourceUpdateOne {
	aruo.mutation.SetResourceName(s)
	return aruo
}

// SetResourceDescription sets the "resource_description" field.
func (aruo *ApplicationResourceUpdateOne) SetResourceDescription(s string) *ApplicationResourceUpdateOne {
	aruo.mutation.SetResourceDescription(s)
	return aruo
}

// SetNillableResourceDescription sets the "resource_description" field if the given value is not nil.
func (aruo *ApplicationResourceUpdateOne) SetNillableResourceDescription(s *string) *ApplicationResourceUpdateOne {
	if s != nil {
		aruo.SetResourceDescription(*s)
	}
	return aruo
}

// ClearResourceDescription clears the value of the "resource_description" field.
func (aruo *ApplicationResourceUpdateOne) ClearResourceDescription() *ApplicationResourceUpdateOne {
	aruo.mutation.ClearResourceDescription()
	return aruo
}

// SetType sets the "type" field.
func (aruo *ApplicationResourceUpdateOne) SetType(s string) *ApplicationResourceUpdateOne {
	aruo.mutation.SetType(s)
	return aruo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (aruo *ApplicationResourceUpdateOne) SetNillableType(s *string) *ApplicationResourceUpdateOne {
	if s != nil {
		aruo.SetType(*s)
	}
	return aruo
}

// SetCreator sets the "creator" field.
func (aruo *ApplicationResourceUpdateOne) SetCreator(u uuid.UUID) *ApplicationResourceUpdateOne {
	aruo.mutation.SetCreator(u)
	return aruo
}

// SetCreateAt sets the "create_at" field.
func (aruo *ApplicationResourceUpdateOne) SetCreateAt(u uint32) *ApplicationResourceUpdateOne {
	aruo.mutation.ResetCreateAt()
	aruo.mutation.SetCreateAt(u)
	return aruo
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (aruo *ApplicationResourceUpdateOne) SetNillableCreateAt(u *uint32) *ApplicationResourceUpdateOne {
	if u != nil {
		aruo.SetCreateAt(*u)
	}
	return aruo
}

// AddCreateAt adds u to the "create_at" field.
func (aruo *ApplicationResourceUpdateOne) AddCreateAt(u uint32) *ApplicationResourceUpdateOne {
	aruo.mutation.AddCreateAt(u)
	return aruo
}

// SetUpdateAt sets the "update_at" field.
func (aruo *ApplicationResourceUpdateOne) SetUpdateAt(u uint32) *ApplicationResourceUpdateOne {
	aruo.mutation.ResetUpdateAt()
	aruo.mutation.SetUpdateAt(u)
	return aruo
}

// AddUpdateAt adds u to the "update_at" field.
func (aruo *ApplicationResourceUpdateOne) AddUpdateAt(u uint32) *ApplicationResourceUpdateOne {
	aruo.mutation.AddUpdateAt(u)
	return aruo
}

// SetDeleteAt sets the "delete_at" field.
func (aruo *ApplicationResourceUpdateOne) SetDeleteAt(u uint32) *ApplicationResourceUpdateOne {
	aruo.mutation.ResetDeleteAt()
	aruo.mutation.SetDeleteAt(u)
	return aruo
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (aruo *ApplicationResourceUpdateOne) SetNillableDeleteAt(u *uint32) *ApplicationResourceUpdateOne {
	if u != nil {
		aruo.SetDeleteAt(*u)
	}
	return aruo
}

// AddDeleteAt adds u to the "delete_at" field.
func (aruo *ApplicationResourceUpdateOne) AddDeleteAt(u uint32) *ApplicationResourceUpdateOne {
	aruo.mutation.AddDeleteAt(u)
	return aruo
}

// Mutation returns the ApplicationResourceMutation object of the builder.
func (aruo *ApplicationResourceUpdateOne) Mutation() *ApplicationResourceMutation {
	return aruo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (aruo *ApplicationResourceUpdateOne) Select(field string, fields ...string) *ApplicationResourceUpdateOne {
	aruo.fields = append([]string{field}, fields...)
	return aruo
}

// Save executes the query and returns the updated ApplicationResource entity.
func (aruo *ApplicationResourceUpdateOne) Save(ctx context.Context) (*ApplicationResource, error) {
	var (
		err  error
		node *ApplicationResource
	)
	aruo.defaults()
	if len(aruo.hooks) == 0 {
		node, err = aruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ApplicationResourceMutation)
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
func (aruo *ApplicationResourceUpdateOne) SaveX(ctx context.Context) *ApplicationResource {
	node, err := aruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (aruo *ApplicationResourceUpdateOne) Exec(ctx context.Context) error {
	_, err := aruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aruo *ApplicationResourceUpdateOne) ExecX(ctx context.Context) {
	if err := aruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (aruo *ApplicationResourceUpdateOne) defaults() {
	if _, ok := aruo.mutation.UpdateAt(); !ok {
		v := applicationresource.UpdateDefaultUpdateAt()
		aruo.mutation.SetUpdateAt(v)
	}
}

func (aruo *ApplicationResourceUpdateOne) sqlSave(ctx context.Context) (_node *ApplicationResource, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   applicationresource.Table,
			Columns: applicationresource.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: applicationresource.FieldID,
			},
		},
	}
	id, ok := aruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing ApplicationResource.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := aruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, applicationresource.FieldID)
		for _, f := range fields {
			if !applicationresource.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != applicationresource.FieldID {
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
			Column: applicationresource.FieldAppID,
		})
	}
	if value, ok := aruo.mutation.ResourceName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: applicationresource.FieldResourceName,
		})
	}
	if value, ok := aruo.mutation.ResourceDescription(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: applicationresource.FieldResourceDescription,
		})
	}
	if aruo.mutation.ResourceDescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: applicationresource.FieldResourceDescription,
		})
	}
	if value, ok := aruo.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: applicationresource.FieldType,
		})
	}
	if value, ok := aruo.mutation.Creator(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: applicationresource.FieldCreator,
		})
	}
	if value, ok := aruo.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: applicationresource.FieldCreateAt,
		})
	}
	if value, ok := aruo.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: applicationresource.FieldCreateAt,
		})
	}
	if value, ok := aruo.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: applicationresource.FieldUpdateAt,
		})
	}
	if value, ok := aruo.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: applicationresource.FieldUpdateAt,
		})
	}
	if value, ok := aruo.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: applicationresource.FieldDeleteAt,
		})
	}
	if value, ok := aruo.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: applicationresource.FieldDeleteAt,
		})
	}
	_node = &ApplicationResource{config: aruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, aruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{applicationresource.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}

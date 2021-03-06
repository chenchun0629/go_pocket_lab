// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"pocket_lab/ent/ent/admin"
	"pocket_lab/ent/ent/bee"
	"pocket_lab/ent/ent/cat"
	"pocket_lab/ent/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CatUpdate is the builder for updating Cat entities.
type CatUpdate struct {
	config
	hooks    []Hook
	mutation *CatMutation
}

// Where appends a list predicates to the CatUpdate builder.
func (cu *CatUpdate) Where(ps ...predicate.Cat) *CatUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetAge sets the "age" field.
func (cu *CatUpdate) SetAge(i int) *CatUpdate {
	cu.mutation.ResetAge()
	cu.mutation.SetAge(i)
	return cu
}

// AddAge adds i to the "age" field.
func (cu *CatUpdate) AddAge(i int) *CatUpdate {
	cu.mutation.AddAge(i)
	return cu
}

// SetName sets the "name" field.
func (cu *CatUpdate) SetName(s string) *CatUpdate {
	cu.mutation.SetName(s)
	return cu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cu *CatUpdate) SetNillableName(s *string) *CatUpdate {
	if s != nil {
		cu.SetName(*s)
	}
	return cu
}

// SetOwnerID sets the "owner" edge to the Admin entity by ID.
func (cu *CatUpdate) SetOwnerID(id int) *CatUpdate {
	cu.mutation.SetOwnerID(id)
	return cu
}

// SetNillableOwnerID sets the "owner" edge to the Admin entity by ID if the given value is not nil.
func (cu *CatUpdate) SetNillableOwnerID(id *int) *CatUpdate {
	if id != nil {
		cu = cu.SetOwnerID(*id)
	}
	return cu
}

// SetOwner sets the "owner" edge to the Admin entity.
func (cu *CatUpdate) SetOwner(a *Admin) *CatUpdate {
	return cu.SetOwnerID(a.ID)
}

// AddCbIDs adds the "cb" edge to the Bee entity by IDs.
func (cu *CatUpdate) AddCbIDs(ids ...int) *CatUpdate {
	cu.mutation.AddCbIDs(ids...)
	return cu
}

// AddCb adds the "cb" edges to the Bee entity.
func (cu *CatUpdate) AddCb(b ...*Bee) *CatUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return cu.AddCbIDs(ids...)
}

// Mutation returns the CatMutation object of the builder.
func (cu *CatUpdate) Mutation() *CatMutation {
	return cu.mutation
}

// ClearOwner clears the "owner" edge to the Admin entity.
func (cu *CatUpdate) ClearOwner() *CatUpdate {
	cu.mutation.ClearOwner()
	return cu
}

// ClearCb clears all "cb" edges to the Bee entity.
func (cu *CatUpdate) ClearCb() *CatUpdate {
	cu.mutation.ClearCb()
	return cu
}

// RemoveCbIDs removes the "cb" edge to Bee entities by IDs.
func (cu *CatUpdate) RemoveCbIDs(ids ...int) *CatUpdate {
	cu.mutation.RemoveCbIDs(ids...)
	return cu
}

// RemoveCb removes "cb" edges to Bee entities.
func (cu *CatUpdate) RemoveCb(b ...*Bee) *CatUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return cu.RemoveCbIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CatUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cu.hooks) == 0 {
		if err = cu.check(); err != nil {
			return 0, err
		}
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CatMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cu.check(); err != nil {
				return 0, err
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			if cu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CatUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CatUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CatUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *CatUpdate) check() error {
	if v, ok := cu.mutation.Age(); ok {
		if err := cat.AgeValidator(v); err != nil {
			return &ValidationError{Name: "age", err: fmt.Errorf("ent: validator failed for field \"age\": %w", err)}
		}
	}
	return nil
}

func (cu *CatUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   cat.Table,
			Columns: cat.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: cat.FieldID,
			},
		},
	}
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Age(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: cat.FieldAge,
		})
	}
	if value, ok := cu.mutation.AddedAge(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: cat.FieldAge,
		})
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: cat.FieldName,
		})
	}
	if cu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cat.OwnerTable,
			Columns: []string{cat.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: admin.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cat.OwnerTable,
			Columns: []string{cat.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: admin.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.CbCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   cat.CbTable,
			Columns: cat.CbPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: bee.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedCbIDs(); len(nodes) > 0 && !cu.mutation.CbCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   cat.CbTable,
			Columns: cat.CbPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: bee.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.CbIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   cat.CbTable,
			Columns: cat.CbPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: bee.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{cat.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// CatUpdateOne is the builder for updating a single Cat entity.
type CatUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CatMutation
}

// SetAge sets the "age" field.
func (cuo *CatUpdateOne) SetAge(i int) *CatUpdateOne {
	cuo.mutation.ResetAge()
	cuo.mutation.SetAge(i)
	return cuo
}

// AddAge adds i to the "age" field.
func (cuo *CatUpdateOne) AddAge(i int) *CatUpdateOne {
	cuo.mutation.AddAge(i)
	return cuo
}

// SetName sets the "name" field.
func (cuo *CatUpdateOne) SetName(s string) *CatUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cuo *CatUpdateOne) SetNillableName(s *string) *CatUpdateOne {
	if s != nil {
		cuo.SetName(*s)
	}
	return cuo
}

// SetOwnerID sets the "owner" edge to the Admin entity by ID.
func (cuo *CatUpdateOne) SetOwnerID(id int) *CatUpdateOne {
	cuo.mutation.SetOwnerID(id)
	return cuo
}

// SetNillableOwnerID sets the "owner" edge to the Admin entity by ID if the given value is not nil.
func (cuo *CatUpdateOne) SetNillableOwnerID(id *int) *CatUpdateOne {
	if id != nil {
		cuo = cuo.SetOwnerID(*id)
	}
	return cuo
}

// SetOwner sets the "owner" edge to the Admin entity.
func (cuo *CatUpdateOne) SetOwner(a *Admin) *CatUpdateOne {
	return cuo.SetOwnerID(a.ID)
}

// AddCbIDs adds the "cb" edge to the Bee entity by IDs.
func (cuo *CatUpdateOne) AddCbIDs(ids ...int) *CatUpdateOne {
	cuo.mutation.AddCbIDs(ids...)
	return cuo
}

// AddCb adds the "cb" edges to the Bee entity.
func (cuo *CatUpdateOne) AddCb(b ...*Bee) *CatUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return cuo.AddCbIDs(ids...)
}

// Mutation returns the CatMutation object of the builder.
func (cuo *CatUpdateOne) Mutation() *CatMutation {
	return cuo.mutation
}

// ClearOwner clears the "owner" edge to the Admin entity.
func (cuo *CatUpdateOne) ClearOwner() *CatUpdateOne {
	cuo.mutation.ClearOwner()
	return cuo
}

// ClearCb clears all "cb" edges to the Bee entity.
func (cuo *CatUpdateOne) ClearCb() *CatUpdateOne {
	cuo.mutation.ClearCb()
	return cuo
}

// RemoveCbIDs removes the "cb" edge to Bee entities by IDs.
func (cuo *CatUpdateOne) RemoveCbIDs(ids ...int) *CatUpdateOne {
	cuo.mutation.RemoveCbIDs(ids...)
	return cuo
}

// RemoveCb removes "cb" edges to Bee entities.
func (cuo *CatUpdateOne) RemoveCb(b ...*Bee) *CatUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return cuo.RemoveCbIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CatUpdateOne) Select(field string, fields ...string) *CatUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Cat entity.
func (cuo *CatUpdateOne) Save(ctx context.Context) (*Cat, error) {
	var (
		err  error
		node *Cat
	)
	if len(cuo.hooks) == 0 {
		if err = cuo.check(); err != nil {
			return nil, err
		}
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CatMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cuo.check(); err != nil {
				return nil, err
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			if cuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CatUpdateOne) SaveX(ctx context.Context) *Cat {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CatUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CatUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *CatUpdateOne) check() error {
	if v, ok := cuo.mutation.Age(); ok {
		if err := cat.AgeValidator(v); err != nil {
			return &ValidationError{Name: "age", err: fmt.Errorf("ent: validator failed for field \"age\": %w", err)}
		}
	}
	return nil
}

func (cuo *CatUpdateOne) sqlSave(ctx context.Context) (_node *Cat, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   cat.Table,
			Columns: cat.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: cat.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Cat.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, cat.FieldID)
		for _, f := range fields {
			if !cat.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != cat.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.Age(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: cat.FieldAge,
		})
	}
	if value, ok := cuo.mutation.AddedAge(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: cat.FieldAge,
		})
	}
	if value, ok := cuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: cat.FieldName,
		})
	}
	if cuo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cat.OwnerTable,
			Columns: []string{cat.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: admin.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cat.OwnerTable,
			Columns: []string{cat.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: admin.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.CbCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   cat.CbTable,
			Columns: cat.CbPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: bee.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedCbIDs(); len(nodes) > 0 && !cuo.mutation.CbCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   cat.CbTable,
			Columns: cat.CbPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: bee.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.CbIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   cat.CbTable,
			Columns: cat.CbPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: bee.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Cat{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{cat.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}

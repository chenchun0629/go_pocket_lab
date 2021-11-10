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

// BeeUpdate is the builder for updating Bee entities.
type BeeUpdate struct {
	config
	hooks    []Hook
	mutation *BeeMutation
}

// Where appends a list predicates to the BeeUpdate builder.
func (bu *BeeUpdate) Where(ps ...predicate.Bee) *BeeUpdate {
	bu.mutation.Where(ps...)
	return bu
}

// SetAge sets the "age" field.
func (bu *BeeUpdate) SetAge(i int) *BeeUpdate {
	bu.mutation.ResetAge()
	bu.mutation.SetAge(i)
	return bu
}

// AddAge adds i to the "age" field.
func (bu *BeeUpdate) AddAge(i int) *BeeUpdate {
	bu.mutation.AddAge(i)
	return bu
}

// SetName sets the "name" field.
func (bu *BeeUpdate) SetName(s string) *BeeUpdate {
	bu.mutation.SetName(s)
	return bu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (bu *BeeUpdate) SetNillableName(s *string) *BeeUpdate {
	if s != nil {
		bu.SetName(*s)
	}
	return bu
}

// AddOwnerIDs adds the "owner" edge to the Admin entity by IDs.
func (bu *BeeUpdate) AddOwnerIDs(ids ...int) *BeeUpdate {
	bu.mutation.AddOwnerIDs(ids...)
	return bu
}

// AddOwner adds the "owner" edges to the Admin entity.
func (bu *BeeUpdate) AddOwner(a ...*Admin) *BeeUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return bu.AddOwnerIDs(ids...)
}

// AddBcIDs adds the "bc" edge to the Cat entity by IDs.
func (bu *BeeUpdate) AddBcIDs(ids ...int) *BeeUpdate {
	bu.mutation.AddBcIDs(ids...)
	return bu
}

// AddBc adds the "bc" edges to the Cat entity.
func (bu *BeeUpdate) AddBc(c ...*Cat) *BeeUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return bu.AddBcIDs(ids...)
}

// Mutation returns the BeeMutation object of the builder.
func (bu *BeeUpdate) Mutation() *BeeMutation {
	return bu.mutation
}

// ClearOwner clears all "owner" edges to the Admin entity.
func (bu *BeeUpdate) ClearOwner() *BeeUpdate {
	bu.mutation.ClearOwner()
	return bu
}

// RemoveOwnerIDs removes the "owner" edge to Admin entities by IDs.
func (bu *BeeUpdate) RemoveOwnerIDs(ids ...int) *BeeUpdate {
	bu.mutation.RemoveOwnerIDs(ids...)
	return bu
}

// RemoveOwner removes "owner" edges to Admin entities.
func (bu *BeeUpdate) RemoveOwner(a ...*Admin) *BeeUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return bu.RemoveOwnerIDs(ids...)
}

// ClearBc clears all "bc" edges to the Cat entity.
func (bu *BeeUpdate) ClearBc() *BeeUpdate {
	bu.mutation.ClearBc()
	return bu
}

// RemoveBcIDs removes the "bc" edge to Cat entities by IDs.
func (bu *BeeUpdate) RemoveBcIDs(ids ...int) *BeeUpdate {
	bu.mutation.RemoveBcIDs(ids...)
	return bu
}

// RemoveBc removes "bc" edges to Cat entities.
func (bu *BeeUpdate) RemoveBc(c ...*Cat) *BeeUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return bu.RemoveBcIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bu *BeeUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(bu.hooks) == 0 {
		if err = bu.check(); err != nil {
			return 0, err
		}
		affected, err = bu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BeeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = bu.check(); err != nil {
				return 0, err
			}
			bu.mutation = mutation
			affected, err = bu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(bu.hooks) - 1; i >= 0; i-- {
			if bu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = bu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (bu *BeeUpdate) SaveX(ctx context.Context) int {
	affected, err := bu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bu *BeeUpdate) Exec(ctx context.Context) error {
	_, err := bu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bu *BeeUpdate) ExecX(ctx context.Context) {
	if err := bu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bu *BeeUpdate) check() error {
	if v, ok := bu.mutation.Age(); ok {
		if err := bee.AgeValidator(v); err != nil {
			return &ValidationError{Name: "age", err: fmt.Errorf("ent: validator failed for field \"age\": %w", err)}
		}
	}
	return nil
}

func (bu *BeeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   bee.Table,
			Columns: bee.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: bee.FieldID,
			},
		},
	}
	if ps := bu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bu.mutation.Age(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: bee.FieldAge,
		})
	}
	if value, ok := bu.mutation.AddedAge(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: bee.FieldAge,
		})
	}
	if value, ok := bu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: bee.FieldName,
		})
	}
	if bu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   bee.OwnerTable,
			Columns: bee.OwnerPrimaryKey,
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
	if nodes := bu.mutation.RemovedOwnerIDs(); len(nodes) > 0 && !bu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   bee.OwnerTable,
			Columns: bee.OwnerPrimaryKey,
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   bee.OwnerTable,
			Columns: bee.OwnerPrimaryKey,
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
	if bu.mutation.BcCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   bee.BcTable,
			Columns: bee.BcPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: cat.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.RemovedBcIDs(); len(nodes) > 0 && !bu.mutation.BcCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   bee.BcTable,
			Columns: bee.BcPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: cat.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.BcIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   bee.BcTable,
			Columns: bee.BcPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: cat.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{bee.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// BeeUpdateOne is the builder for updating a single Bee entity.
type BeeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BeeMutation
}

// SetAge sets the "age" field.
func (buo *BeeUpdateOne) SetAge(i int) *BeeUpdateOne {
	buo.mutation.ResetAge()
	buo.mutation.SetAge(i)
	return buo
}

// AddAge adds i to the "age" field.
func (buo *BeeUpdateOne) AddAge(i int) *BeeUpdateOne {
	buo.mutation.AddAge(i)
	return buo
}

// SetName sets the "name" field.
func (buo *BeeUpdateOne) SetName(s string) *BeeUpdateOne {
	buo.mutation.SetName(s)
	return buo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (buo *BeeUpdateOne) SetNillableName(s *string) *BeeUpdateOne {
	if s != nil {
		buo.SetName(*s)
	}
	return buo
}

// AddOwnerIDs adds the "owner" edge to the Admin entity by IDs.
func (buo *BeeUpdateOne) AddOwnerIDs(ids ...int) *BeeUpdateOne {
	buo.mutation.AddOwnerIDs(ids...)
	return buo
}

// AddOwner adds the "owner" edges to the Admin entity.
func (buo *BeeUpdateOne) AddOwner(a ...*Admin) *BeeUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return buo.AddOwnerIDs(ids...)
}

// AddBcIDs adds the "bc" edge to the Cat entity by IDs.
func (buo *BeeUpdateOne) AddBcIDs(ids ...int) *BeeUpdateOne {
	buo.mutation.AddBcIDs(ids...)
	return buo
}

// AddBc adds the "bc" edges to the Cat entity.
func (buo *BeeUpdateOne) AddBc(c ...*Cat) *BeeUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return buo.AddBcIDs(ids...)
}

// Mutation returns the BeeMutation object of the builder.
func (buo *BeeUpdateOne) Mutation() *BeeMutation {
	return buo.mutation
}

// ClearOwner clears all "owner" edges to the Admin entity.
func (buo *BeeUpdateOne) ClearOwner() *BeeUpdateOne {
	buo.mutation.ClearOwner()
	return buo
}

// RemoveOwnerIDs removes the "owner" edge to Admin entities by IDs.
func (buo *BeeUpdateOne) RemoveOwnerIDs(ids ...int) *BeeUpdateOne {
	buo.mutation.RemoveOwnerIDs(ids...)
	return buo
}

// RemoveOwner removes "owner" edges to Admin entities.
func (buo *BeeUpdateOne) RemoveOwner(a ...*Admin) *BeeUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return buo.RemoveOwnerIDs(ids...)
}

// ClearBc clears all "bc" edges to the Cat entity.
func (buo *BeeUpdateOne) ClearBc() *BeeUpdateOne {
	buo.mutation.ClearBc()
	return buo
}

// RemoveBcIDs removes the "bc" edge to Cat entities by IDs.
func (buo *BeeUpdateOne) RemoveBcIDs(ids ...int) *BeeUpdateOne {
	buo.mutation.RemoveBcIDs(ids...)
	return buo
}

// RemoveBc removes "bc" edges to Cat entities.
func (buo *BeeUpdateOne) RemoveBc(c ...*Cat) *BeeUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return buo.RemoveBcIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (buo *BeeUpdateOne) Select(field string, fields ...string) *BeeUpdateOne {
	buo.fields = append([]string{field}, fields...)
	return buo
}

// Save executes the query and returns the updated Bee entity.
func (buo *BeeUpdateOne) Save(ctx context.Context) (*Bee, error) {
	var (
		err  error
		node *Bee
	)
	if len(buo.hooks) == 0 {
		if err = buo.check(); err != nil {
			return nil, err
		}
		node, err = buo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BeeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = buo.check(); err != nil {
				return nil, err
			}
			buo.mutation = mutation
			node, err = buo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(buo.hooks) - 1; i >= 0; i-- {
			if buo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = buo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, buo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (buo *BeeUpdateOne) SaveX(ctx context.Context) *Bee {
	node, err := buo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (buo *BeeUpdateOne) Exec(ctx context.Context) error {
	_, err := buo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buo *BeeUpdateOne) ExecX(ctx context.Context) {
	if err := buo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (buo *BeeUpdateOne) check() error {
	if v, ok := buo.mutation.Age(); ok {
		if err := bee.AgeValidator(v); err != nil {
			return &ValidationError{Name: "age", err: fmt.Errorf("ent: validator failed for field \"age\": %w", err)}
		}
	}
	return nil
}

func (buo *BeeUpdateOne) sqlSave(ctx context.Context) (_node *Bee, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   bee.Table,
			Columns: bee.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: bee.FieldID,
			},
		},
	}
	id, ok := buo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Bee.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := buo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, bee.FieldID)
		for _, f := range fields {
			if !bee.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != bee.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := buo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := buo.mutation.Age(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: bee.FieldAge,
		})
	}
	if value, ok := buo.mutation.AddedAge(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: bee.FieldAge,
		})
	}
	if value, ok := buo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: bee.FieldName,
		})
	}
	if buo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   bee.OwnerTable,
			Columns: bee.OwnerPrimaryKey,
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
	if nodes := buo.mutation.RemovedOwnerIDs(); len(nodes) > 0 && !buo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   bee.OwnerTable,
			Columns: bee.OwnerPrimaryKey,
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   bee.OwnerTable,
			Columns: bee.OwnerPrimaryKey,
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
	if buo.mutation.BcCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   bee.BcTable,
			Columns: bee.BcPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: cat.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.RemovedBcIDs(); len(nodes) > 0 && !buo.mutation.BcCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   bee.BcTable,
			Columns: bee.BcPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: cat.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.BcIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   bee.BcTable,
			Columns: bee.BcPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: cat.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Bee{config: buo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, buo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{bee.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
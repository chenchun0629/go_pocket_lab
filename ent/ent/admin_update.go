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

// AdminUpdate is the builder for updating Admin entities.
type AdminUpdate struct {
	config
	hooks    []Hook
	mutation *AdminMutation
}

// Where appends a list predicates to the AdminUpdate builder.
func (au *AdminUpdate) Where(ps ...predicate.Admin) *AdminUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetAge sets the "age" field.
func (au *AdminUpdate) SetAge(i int) *AdminUpdate {
	au.mutation.ResetAge()
	au.mutation.SetAge(i)
	return au
}

// AddAge adds i to the "age" field.
func (au *AdminUpdate) AddAge(i int) *AdminUpdate {
	au.mutation.AddAge(i)
	return au
}

// SetName sets the "name" field.
func (au *AdminUpdate) SetName(s string) *AdminUpdate {
	au.mutation.SetName(s)
	return au
}

// SetNillableName sets the "name" field if the given value is not nil.
func (au *AdminUpdate) SetNillableName(s *string) *AdminUpdate {
	if s != nil {
		au.SetName(*s)
	}
	return au
}

// AddCatIDs adds the "cats" edge to the Cat entity by IDs.
func (au *AdminUpdate) AddCatIDs(ids ...int) *AdminUpdate {
	au.mutation.AddCatIDs(ids...)
	return au
}

// AddCats adds the "cats" edges to the Cat entity.
func (au *AdminUpdate) AddCats(c ...*Cat) *AdminUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return au.AddCatIDs(ids...)
}

// AddBeeIDs adds the "bees" edge to the Bee entity by IDs.
func (au *AdminUpdate) AddBeeIDs(ids ...int) *AdminUpdate {
	au.mutation.AddBeeIDs(ids...)
	return au
}

// AddBees adds the "bees" edges to the Bee entity.
func (au *AdminUpdate) AddBees(b ...*Bee) *AdminUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return au.AddBeeIDs(ids...)
}

// SetParentID sets the "parent" edge to the Admin entity by ID.
func (au *AdminUpdate) SetParentID(id int) *AdminUpdate {
	au.mutation.SetParentID(id)
	return au
}

// SetNillableParentID sets the "parent" edge to the Admin entity by ID if the given value is not nil.
func (au *AdminUpdate) SetNillableParentID(id *int) *AdminUpdate {
	if id != nil {
		au = au.SetParentID(*id)
	}
	return au
}

// SetParent sets the "parent" edge to the Admin entity.
func (au *AdminUpdate) SetParent(a *Admin) *AdminUpdate {
	return au.SetParentID(a.ID)
}

// SetSonID sets the "son" edge to the Admin entity by ID.
func (au *AdminUpdate) SetSonID(id int) *AdminUpdate {
	au.mutation.SetSonID(id)
	return au
}

// SetNillableSonID sets the "son" edge to the Admin entity by ID if the given value is not nil.
func (au *AdminUpdate) SetNillableSonID(id *int) *AdminUpdate {
	if id != nil {
		au = au.SetSonID(*id)
	}
	return au
}

// SetSon sets the "son" edge to the Admin entity.
func (au *AdminUpdate) SetSon(a *Admin) *AdminUpdate {
	return au.SetSonID(a.ID)
}

// Mutation returns the AdminMutation object of the builder.
func (au *AdminUpdate) Mutation() *AdminMutation {
	return au.mutation
}

// ClearCats clears all "cats" edges to the Cat entity.
func (au *AdminUpdate) ClearCats() *AdminUpdate {
	au.mutation.ClearCats()
	return au
}

// RemoveCatIDs removes the "cats" edge to Cat entities by IDs.
func (au *AdminUpdate) RemoveCatIDs(ids ...int) *AdminUpdate {
	au.mutation.RemoveCatIDs(ids...)
	return au
}

// RemoveCats removes "cats" edges to Cat entities.
func (au *AdminUpdate) RemoveCats(c ...*Cat) *AdminUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return au.RemoveCatIDs(ids...)
}

// ClearBees clears all "bees" edges to the Bee entity.
func (au *AdminUpdate) ClearBees() *AdminUpdate {
	au.mutation.ClearBees()
	return au
}

// RemoveBeeIDs removes the "bees" edge to Bee entities by IDs.
func (au *AdminUpdate) RemoveBeeIDs(ids ...int) *AdminUpdate {
	au.mutation.RemoveBeeIDs(ids...)
	return au
}

// RemoveBees removes "bees" edges to Bee entities.
func (au *AdminUpdate) RemoveBees(b ...*Bee) *AdminUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return au.RemoveBeeIDs(ids...)
}

// ClearParent clears the "parent" edge to the Admin entity.
func (au *AdminUpdate) ClearParent() *AdminUpdate {
	au.mutation.ClearParent()
	return au
}

// ClearSon clears the "son" edge to the Admin entity.
func (au *AdminUpdate) ClearSon() *AdminUpdate {
	au.mutation.ClearSon()
	return au
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AdminUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(au.hooks) == 0 {
		if err = au.check(); err != nil {
			return 0, err
		}
		affected, err = au.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AdminMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = au.check(); err != nil {
				return 0, err
			}
			au.mutation = mutation
			affected, err = au.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(au.hooks) - 1; i >= 0; i-- {
			if au.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = au.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, au.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (au *AdminUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AdminUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AdminUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (au *AdminUpdate) check() error {
	if v, ok := au.mutation.Age(); ok {
		if err := admin.AgeValidator(v); err != nil {
			return &ValidationError{Name: "age", err: fmt.Errorf("ent: validator failed for field \"age\": %w", err)}
		}
	}
	return nil
}

func (au *AdminUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   admin.Table,
			Columns: admin.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: admin.FieldID,
			},
		},
	}
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.Age(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: admin.FieldAge,
		})
	}
	if value, ok := au.mutation.AddedAge(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: admin.FieldAge,
		})
	}
	if value, ok := au.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: admin.FieldName,
		})
	}
	if au.mutation.CatsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   admin.CatsTable,
			Columns: []string{admin.CatsColumn},
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
	if nodes := au.mutation.RemovedCatsIDs(); len(nodes) > 0 && !au.mutation.CatsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   admin.CatsTable,
			Columns: []string{admin.CatsColumn},
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
	if nodes := au.mutation.CatsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   admin.CatsTable,
			Columns: []string{admin.CatsColumn},
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
	if au.mutation.BeesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   admin.BeesTable,
			Columns: admin.BeesPrimaryKey,
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
	if nodes := au.mutation.RemovedBeesIDs(); len(nodes) > 0 && !au.mutation.BeesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   admin.BeesTable,
			Columns: admin.BeesPrimaryKey,
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
	if nodes := au.mutation.BeesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   admin.BeesTable,
			Columns: admin.BeesPrimaryKey,
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
	if au.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   admin.ParentTable,
			Columns: []string{admin.ParentColumn},
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
	if nodes := au.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   admin.ParentTable,
			Columns: []string{admin.ParentColumn},
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
	if au.mutation.SonCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   admin.SonTable,
			Columns: []string{admin.SonColumn},
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
	if nodes := au.mutation.SonIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   admin.SonTable,
			Columns: []string{admin.SonColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{admin.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// AdminUpdateOne is the builder for updating a single Admin entity.
type AdminUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AdminMutation
}

// SetAge sets the "age" field.
func (auo *AdminUpdateOne) SetAge(i int) *AdminUpdateOne {
	auo.mutation.ResetAge()
	auo.mutation.SetAge(i)
	return auo
}

// AddAge adds i to the "age" field.
func (auo *AdminUpdateOne) AddAge(i int) *AdminUpdateOne {
	auo.mutation.AddAge(i)
	return auo
}

// SetName sets the "name" field.
func (auo *AdminUpdateOne) SetName(s string) *AdminUpdateOne {
	auo.mutation.SetName(s)
	return auo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (auo *AdminUpdateOne) SetNillableName(s *string) *AdminUpdateOne {
	if s != nil {
		auo.SetName(*s)
	}
	return auo
}

// AddCatIDs adds the "cats" edge to the Cat entity by IDs.
func (auo *AdminUpdateOne) AddCatIDs(ids ...int) *AdminUpdateOne {
	auo.mutation.AddCatIDs(ids...)
	return auo
}

// AddCats adds the "cats" edges to the Cat entity.
func (auo *AdminUpdateOne) AddCats(c ...*Cat) *AdminUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return auo.AddCatIDs(ids...)
}

// AddBeeIDs adds the "bees" edge to the Bee entity by IDs.
func (auo *AdminUpdateOne) AddBeeIDs(ids ...int) *AdminUpdateOne {
	auo.mutation.AddBeeIDs(ids...)
	return auo
}

// AddBees adds the "bees" edges to the Bee entity.
func (auo *AdminUpdateOne) AddBees(b ...*Bee) *AdminUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return auo.AddBeeIDs(ids...)
}

// SetParentID sets the "parent" edge to the Admin entity by ID.
func (auo *AdminUpdateOne) SetParentID(id int) *AdminUpdateOne {
	auo.mutation.SetParentID(id)
	return auo
}

// SetNillableParentID sets the "parent" edge to the Admin entity by ID if the given value is not nil.
func (auo *AdminUpdateOne) SetNillableParentID(id *int) *AdminUpdateOne {
	if id != nil {
		auo = auo.SetParentID(*id)
	}
	return auo
}

// SetParent sets the "parent" edge to the Admin entity.
func (auo *AdminUpdateOne) SetParent(a *Admin) *AdminUpdateOne {
	return auo.SetParentID(a.ID)
}

// SetSonID sets the "son" edge to the Admin entity by ID.
func (auo *AdminUpdateOne) SetSonID(id int) *AdminUpdateOne {
	auo.mutation.SetSonID(id)
	return auo
}

// SetNillableSonID sets the "son" edge to the Admin entity by ID if the given value is not nil.
func (auo *AdminUpdateOne) SetNillableSonID(id *int) *AdminUpdateOne {
	if id != nil {
		auo = auo.SetSonID(*id)
	}
	return auo
}

// SetSon sets the "son" edge to the Admin entity.
func (auo *AdminUpdateOne) SetSon(a *Admin) *AdminUpdateOne {
	return auo.SetSonID(a.ID)
}

// Mutation returns the AdminMutation object of the builder.
func (auo *AdminUpdateOne) Mutation() *AdminMutation {
	return auo.mutation
}

// ClearCats clears all "cats" edges to the Cat entity.
func (auo *AdminUpdateOne) ClearCats() *AdminUpdateOne {
	auo.mutation.ClearCats()
	return auo
}

// RemoveCatIDs removes the "cats" edge to Cat entities by IDs.
func (auo *AdminUpdateOne) RemoveCatIDs(ids ...int) *AdminUpdateOne {
	auo.mutation.RemoveCatIDs(ids...)
	return auo
}

// RemoveCats removes "cats" edges to Cat entities.
func (auo *AdminUpdateOne) RemoveCats(c ...*Cat) *AdminUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return auo.RemoveCatIDs(ids...)
}

// ClearBees clears all "bees" edges to the Bee entity.
func (auo *AdminUpdateOne) ClearBees() *AdminUpdateOne {
	auo.mutation.ClearBees()
	return auo
}

// RemoveBeeIDs removes the "bees" edge to Bee entities by IDs.
func (auo *AdminUpdateOne) RemoveBeeIDs(ids ...int) *AdminUpdateOne {
	auo.mutation.RemoveBeeIDs(ids...)
	return auo
}

// RemoveBees removes "bees" edges to Bee entities.
func (auo *AdminUpdateOne) RemoveBees(b ...*Bee) *AdminUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return auo.RemoveBeeIDs(ids...)
}

// ClearParent clears the "parent" edge to the Admin entity.
func (auo *AdminUpdateOne) ClearParent() *AdminUpdateOne {
	auo.mutation.ClearParent()
	return auo
}

// ClearSon clears the "son" edge to the Admin entity.
func (auo *AdminUpdateOne) ClearSon() *AdminUpdateOne {
	auo.mutation.ClearSon()
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AdminUpdateOne) Select(field string, fields ...string) *AdminUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Admin entity.
func (auo *AdminUpdateOne) Save(ctx context.Context) (*Admin, error) {
	var (
		err  error
		node *Admin
	)
	if len(auo.hooks) == 0 {
		if err = auo.check(); err != nil {
			return nil, err
		}
		node, err = auo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AdminMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = auo.check(); err != nil {
				return nil, err
			}
			auo.mutation = mutation
			node, err = auo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(auo.hooks) - 1; i >= 0; i-- {
			if auo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = auo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, auo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AdminUpdateOne) SaveX(ctx context.Context) *Admin {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AdminUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AdminUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (auo *AdminUpdateOne) check() error {
	if v, ok := auo.mutation.Age(); ok {
		if err := admin.AgeValidator(v); err != nil {
			return &ValidationError{Name: "age", err: fmt.Errorf("ent: validator failed for field \"age\": %w", err)}
		}
	}
	return nil
}

func (auo *AdminUpdateOne) sqlSave(ctx context.Context) (_node *Admin, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   admin.Table,
			Columns: admin.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: admin.FieldID,
			},
		},
	}
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Admin.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, admin.FieldID)
		for _, f := range fields {
			if !admin.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != admin.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.Age(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: admin.FieldAge,
		})
	}
	if value, ok := auo.mutation.AddedAge(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: admin.FieldAge,
		})
	}
	if value, ok := auo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: admin.FieldName,
		})
	}
	if auo.mutation.CatsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   admin.CatsTable,
			Columns: []string{admin.CatsColumn},
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
	if nodes := auo.mutation.RemovedCatsIDs(); len(nodes) > 0 && !auo.mutation.CatsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   admin.CatsTable,
			Columns: []string{admin.CatsColumn},
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
	if nodes := auo.mutation.CatsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   admin.CatsTable,
			Columns: []string{admin.CatsColumn},
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
	if auo.mutation.BeesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   admin.BeesTable,
			Columns: admin.BeesPrimaryKey,
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
	if nodes := auo.mutation.RemovedBeesIDs(); len(nodes) > 0 && !auo.mutation.BeesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   admin.BeesTable,
			Columns: admin.BeesPrimaryKey,
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
	if nodes := auo.mutation.BeesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   admin.BeesTable,
			Columns: admin.BeesPrimaryKey,
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
	if auo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   admin.ParentTable,
			Columns: []string{admin.ParentColumn},
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
	if nodes := auo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   admin.ParentTable,
			Columns: []string{admin.ParentColumn},
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
	if auo.mutation.SonCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   admin.SonTable,
			Columns: []string{admin.SonColumn},
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
	if nodes := auo.mutation.SonIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   admin.SonTable,
			Columns: []string{admin.SonColumn},
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
	_node = &Admin{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{admin.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}

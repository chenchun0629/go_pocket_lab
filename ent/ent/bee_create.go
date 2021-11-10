// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"pocket_lab/ent/ent/admin"
	"pocket_lab/ent/ent/bee"
	"pocket_lab/ent/ent/cat"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BeeCreate is the builder for creating a Bee entity.
type BeeCreate struct {
	config
	mutation *BeeMutation
	hooks    []Hook
}

// SetAge sets the "age" field.
func (bc *BeeCreate) SetAge(i int) *BeeCreate {
	bc.mutation.SetAge(i)
	return bc
}

// SetName sets the "name" field.
func (bc *BeeCreate) SetName(s string) *BeeCreate {
	bc.mutation.SetName(s)
	return bc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (bc *BeeCreate) SetNillableName(s *string) *BeeCreate {
	if s != nil {
		bc.SetName(*s)
	}
	return bc
}

// AddOwnerIDs adds the "owner" edge to the Admin entity by IDs.
func (bc *BeeCreate) AddOwnerIDs(ids ...int) *BeeCreate {
	bc.mutation.AddOwnerIDs(ids...)
	return bc
}

// AddOwner adds the "owner" edges to the Admin entity.
func (bc *BeeCreate) AddOwner(a ...*Admin) *BeeCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return bc.AddOwnerIDs(ids...)
}

// AddBcIDs adds the "bc" edge to the Cat entity by IDs.
func (bc *BeeCreate) AddBcIDs(ids ...int) *BeeCreate {
	bc.mutation.AddBcIDs(ids...)
	return bc
}

// AddBc adds the "bc" edges to the Cat entity.
func (bc *BeeCreate) AddBc(c ...*Cat) *BeeCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return bc.AddBcIDs(ids...)
}

// Mutation returns the BeeMutation object of the builder.
func (bc *BeeCreate) Mutation() *BeeMutation {
	return bc.mutation
}

// Save creates the Bee in the database.
func (bc *BeeCreate) Save(ctx context.Context) (*Bee, error) {
	var (
		err  error
		node *Bee
	)
	bc.defaults()
	if len(bc.hooks) == 0 {
		if err = bc.check(); err != nil {
			return nil, err
		}
		node, err = bc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BeeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = bc.check(); err != nil {
				return nil, err
			}
			bc.mutation = mutation
			if node, err = bc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(bc.hooks) - 1; i >= 0; i-- {
			if bc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = bc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (bc *BeeCreate) SaveX(ctx context.Context) *Bee {
	v, err := bc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bc *BeeCreate) Exec(ctx context.Context) error {
	_, err := bc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bc *BeeCreate) ExecX(ctx context.Context) {
	if err := bc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bc *BeeCreate) defaults() {
	if _, ok := bc.mutation.Name(); !ok {
		v := bee.DefaultName
		bc.mutation.SetName(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bc *BeeCreate) check() error {
	if _, ok := bc.mutation.Age(); !ok {
		return &ValidationError{Name: "age", err: errors.New(`ent: missing required field "age"`)}
	}
	if v, ok := bc.mutation.Age(); ok {
		if err := bee.AgeValidator(v); err != nil {
			return &ValidationError{Name: "age", err: fmt.Errorf(`ent: validator failed for field "age": %w`, err)}
		}
	}
	if _, ok := bc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "name"`)}
	}
	return nil
}

func (bc *BeeCreate) sqlSave(ctx context.Context) (*Bee, error) {
	_node, _spec := bc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (bc *BeeCreate) createSpec() (*Bee, *sqlgraph.CreateSpec) {
	var (
		_node = &Bee{config: bc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: bee.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: bee.FieldID,
			},
		}
	)
	if value, ok := bc.mutation.Age(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: bee.FieldAge,
		})
		_node.Age = value
	}
	if value, ok := bc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: bee.FieldName,
		})
		_node.Name = value
	}
	if nodes := bc.mutation.OwnerIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bc.mutation.BcIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// BeeCreateBulk is the builder for creating many Bee entities in bulk.
type BeeCreateBulk struct {
	config
	builders []*BeeCreate
}

// Save creates the Bee entities in the database.
func (bcb *BeeCreateBulk) Save(ctx context.Context) ([]*Bee, error) {
	specs := make([]*sqlgraph.CreateSpec, len(bcb.builders))
	nodes := make([]*Bee, len(bcb.builders))
	mutators := make([]Mutator, len(bcb.builders))
	for i := range bcb.builders {
		func(i int, root context.Context) {
			builder := bcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BeeMutation)
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
					_, err = mutators[i+1].Mutate(root, bcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bcb.driver, spec); err != nil {
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
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, bcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bcb *BeeCreateBulk) SaveX(ctx context.Context) []*Bee {
	v, err := bcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bcb *BeeCreateBulk) Exec(ctx context.Context) error {
	_, err := bcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcb *BeeCreateBulk) ExecX(ctx context.Context) {
	if err := bcb.Exec(ctx); err != nil {
		panic(err)
	}
}
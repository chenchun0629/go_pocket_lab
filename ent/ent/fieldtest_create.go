// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"pocket_lab/ent/ent/fieldtest"
	"pocket_lab/ent/ent/schema"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// FieldTestCreate is the builder for creating a FieldTest entity.
type FieldTestCreate struct {
	config
	mutation *FieldTestMutation
	hooks    []Hook
}

// SetDateD sets the "date_d" field.
func (ftc *FieldTestCreate) SetDateD(t time.Time) *FieldTestCreate {
	ftc.mutation.SetDateD(t)
	return ftc
}

// SetDate sets the "date" field.
func (ftc *FieldTestCreate) SetDate(t time.Time) *FieldTestCreate {
	ftc.mutation.SetDate(t)
	return ftc
}

// SetTitle sets the "title" field.
func (ftc *FieldTestCreate) SetTitle(s string) *FieldTestCreate {
	ftc.mutation.SetTitle(s)
	return ftc
}

// SetJF sets the "j_f" field.
func (ftc *FieldTestCreate) SetJF(s []string) *FieldTestCreate {
	ftc.mutation.SetJF(s)
	return ftc
}

// SetJSF sets the "j_s_f" field.
func (ftc *FieldTestCreate) SetJSF(sc *schema.FTConfig) *FieldTestCreate {
	ftc.mutation.SetJSF(sc)
	return ftc
}

// SetAppID sets the "app_id" field.
func (ftc *FieldTestCreate) SetAppID(s string) *FieldTestCreate {
	ftc.mutation.SetAppID(s)
	return ftc
}

// Mutation returns the FieldTestMutation object of the builder.
func (ftc *FieldTestCreate) Mutation() *FieldTestMutation {
	return ftc.mutation
}

// Save creates the FieldTest in the database.
func (ftc *FieldTestCreate) Save(ctx context.Context) (*FieldTest, error) {
	var (
		err  error
		node *FieldTest
	)
	if len(ftc.hooks) == 0 {
		if err = ftc.check(); err != nil {
			return nil, err
		}
		node, err = ftc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FieldTestMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ftc.check(); err != nil {
				return nil, err
			}
			ftc.mutation = mutation
			if node, err = ftc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ftc.hooks) - 1; i >= 0; i-- {
			if ftc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ftc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ftc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ftc *FieldTestCreate) SaveX(ctx context.Context) *FieldTest {
	v, err := ftc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ftc *FieldTestCreate) Exec(ctx context.Context) error {
	_, err := ftc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ftc *FieldTestCreate) ExecX(ctx context.Context) {
	if err := ftc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ftc *FieldTestCreate) check() error {
	if _, ok := ftc.mutation.DateD(); !ok {
		return &ValidationError{Name: "date_d", err: errors.New(`ent: missing required field "date_d"`)}
	}
	if _, ok := ftc.mutation.Date(); !ok {
		return &ValidationError{Name: "date", err: errors.New(`ent: missing required field "date"`)}
	}
	if _, ok := ftc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "title"`)}
	}
	if _, ok := ftc.mutation.JF(); !ok {
		return &ValidationError{Name: "j_f", err: errors.New(`ent: missing required field "j_f"`)}
	}
	if _, ok := ftc.mutation.JSF(); !ok {
		return &ValidationError{Name: "j_s_f", err: errors.New(`ent: missing required field "j_s_f"`)}
	}
	if _, ok := ftc.mutation.AppID(); !ok {
		return &ValidationError{Name: "app_id", err: errors.New(`ent: missing required field "app_id"`)}
	}
	return nil
}

func (ftc *FieldTestCreate) sqlSave(ctx context.Context) (*FieldTest, error) {
	_node, _spec := ftc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ftc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ftc *FieldTestCreate) createSpec() (*FieldTest, *sqlgraph.CreateSpec) {
	var (
		_node = &FieldTest{config: ftc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: fieldtest.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: fieldtest.FieldID,
			},
		}
	)
	if value, ok := ftc.mutation.DateD(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: fieldtest.FieldDateD,
		})
		_node.DateD = value
	}
	if value, ok := ftc.mutation.Date(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: fieldtest.FieldDate,
		})
		_node.Date = value
	}
	if value, ok := ftc.mutation.Title(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: fieldtest.FieldTitle,
		})
		_node.Title = value
	}
	if value, ok := ftc.mutation.JF(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: fieldtest.FieldJF,
		})
		_node.JF = value
	}
	if value, ok := ftc.mutation.JSF(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: fieldtest.FieldJSF,
		})
		_node.JSF = value
	}
	if value, ok := ftc.mutation.AppID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: fieldtest.FieldAppID,
		})
		_node.AppID = value
	}
	return _node, _spec
}

// FieldTestCreateBulk is the builder for creating many FieldTest entities in bulk.
type FieldTestCreateBulk struct {
	config
	builders []*FieldTestCreate
}

// Save creates the FieldTest entities in the database.
func (ftcb *FieldTestCreateBulk) Save(ctx context.Context) ([]*FieldTest, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ftcb.builders))
	nodes := make([]*FieldTest, len(ftcb.builders))
	mutators := make([]Mutator, len(ftcb.builders))
	for i := range ftcb.builders {
		func(i int, root context.Context) {
			builder := ftcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FieldTestMutation)
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
					_, err = mutators[i+1].Mutate(root, ftcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ftcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ftcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ftcb *FieldTestCreateBulk) SaveX(ctx context.Context) []*FieldTest {
	v, err := ftcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ftcb *FieldTestCreateBulk) Exec(ctx context.Context) error {
	_, err := ftcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ftcb *FieldTestCreateBulk) ExecX(ctx context.Context) {
	if err := ftcb.Exec(ctx); err != nil {
		panic(err)
	}
}

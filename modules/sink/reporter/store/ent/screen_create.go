// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/blushft/strana/modules/sink/reporter/store/ent/screen"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// ScreenCreate is the builder for creating a Screen entity.
type ScreenCreate struct {
	config
	mutation *ScreenMutation
	hooks    []Hook
}

// Mutation returns the ScreenMutation object of the builder.
func (sc *ScreenCreate) Mutation() *ScreenMutation {
	return sc.mutation
}

// Save creates the Screen in the database.
func (sc *ScreenCreate) Save(ctx context.Context) (*Screen, error) {
	if err := sc.preSave(); err != nil {
		return nil, err
	}
	var (
		err  error
		node *Screen
	)
	if len(sc.hooks) == 0 {
		node, err = sc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ScreenMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sc.mutation = mutation
			node, err = sc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(sc.hooks) - 1; i >= 0; i-- {
			mut = sc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sc *ScreenCreate) SaveX(ctx context.Context) *Screen {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (sc *ScreenCreate) preSave() error {
	return nil
}

func (sc *ScreenCreate) sqlSave(ctx context.Context) (*Screen, error) {
	s, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	s.ID = int(id)
	return s, nil
}

func (sc *ScreenCreate) createSpec() (*Screen, *sqlgraph.CreateSpec) {
	var (
		s     = &Screen{config: sc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: screen.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: screen.FieldID,
			},
		}
	)
	return s, _spec
}

// ScreenCreateBulk is the builder for creating a bulk of Screen entities.
type ScreenCreateBulk struct {
	config
	builders []*ScreenCreate
}

// Save creates the Screen entities in the database.
func (scb *ScreenCreateBulk) Save(ctx context.Context) ([]*Screen, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Screen, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				if err := builder.preSave(); err != nil {
					return nil, err
				}
				mutation, ok := m.(*ScreenMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (scb *ScreenCreateBulk) SaveX(ctx context.Context) []*Screen {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
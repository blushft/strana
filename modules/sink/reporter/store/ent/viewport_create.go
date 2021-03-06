// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/blushft/strana/modules/sink/reporter/store/ent/event"
	"github.com/blushft/strana/modules/sink/reporter/store/ent/viewport"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
)

// ViewportCreate is the builder for creating a Viewport entity.
type ViewportCreate struct {
	config
	mutation *ViewportMutation
	hooks    []Hook
}

// SetDensity sets the density field.
func (vc *ViewportCreate) SetDensity(i int) *ViewportCreate {
	vc.mutation.SetDensity(i)
	return vc
}

// SetWidth sets the width field.
func (vc *ViewportCreate) SetWidth(i int) *ViewportCreate {
	vc.mutation.SetWidth(i)
	return vc
}

// SetHeight sets the height field.
func (vc *ViewportCreate) SetHeight(i int) *ViewportCreate {
	vc.mutation.SetHeight(i)
	return vc
}

// AddEventIDs adds the events edge to Event by ids.
func (vc *ViewportCreate) AddEventIDs(ids ...uuid.UUID) *ViewportCreate {
	vc.mutation.AddEventIDs(ids...)
	return vc
}

// AddEvents adds the events edges to Event.
func (vc *ViewportCreate) AddEvents(e ...*Event) *ViewportCreate {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return vc.AddEventIDs(ids...)
}

// Mutation returns the ViewportMutation object of the builder.
func (vc *ViewportCreate) Mutation() *ViewportMutation {
	return vc.mutation
}

// Save creates the Viewport in the database.
func (vc *ViewportCreate) Save(ctx context.Context) (*Viewport, error) {
	if err := vc.preSave(); err != nil {
		return nil, err
	}
	var (
		err  error
		node *Viewport
	)
	if len(vc.hooks) == 0 {
		node, err = vc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ViewportMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			vc.mutation = mutation
			node, err = vc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(vc.hooks) - 1; i >= 0; i-- {
			mut = vc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, vc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (vc *ViewportCreate) SaveX(ctx context.Context) *Viewport {
	v, err := vc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (vc *ViewportCreate) preSave() error {
	if _, ok := vc.mutation.Density(); !ok {
		return &ValidationError{Name: "density", err: errors.New("ent: missing required field \"density\"")}
	}
	if _, ok := vc.mutation.Width(); !ok {
		return &ValidationError{Name: "width", err: errors.New("ent: missing required field \"width\"")}
	}
	if _, ok := vc.mutation.Height(); !ok {
		return &ValidationError{Name: "height", err: errors.New("ent: missing required field \"height\"")}
	}
	return nil
}

func (vc *ViewportCreate) sqlSave(ctx context.Context) (*Viewport, error) {
	v, _spec := vc.createSpec()
	if err := sqlgraph.CreateNode(ctx, vc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	v.ID = int(id)
	return v, nil
}

func (vc *ViewportCreate) createSpec() (*Viewport, *sqlgraph.CreateSpec) {
	var (
		v     = &Viewport{config: vc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: viewport.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: viewport.FieldID,
			},
		}
	)
	if value, ok := vc.mutation.Density(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: viewport.FieldDensity,
		})
		v.Density = value
	}
	if value, ok := vc.mutation.Width(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: viewport.FieldWidth,
		})
		v.Width = value
	}
	if value, ok := vc.mutation.Height(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: viewport.FieldHeight,
		})
		v.Height = value
	}
	if nodes := vc.mutation.EventsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   viewport.EventsTable,
			Columns: []string{viewport.EventsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: event.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return v, _spec
}

// ViewportCreateBulk is the builder for creating a bulk of Viewport entities.
type ViewportCreateBulk struct {
	config
	builders []*ViewportCreate
}

// Save creates the Viewport entities in the database.
func (vcb *ViewportCreateBulk) Save(ctx context.Context) ([]*Viewport, error) {
	specs := make([]*sqlgraph.CreateSpec, len(vcb.builders))
	nodes := make([]*Viewport, len(vcb.builders))
	mutators := make([]Mutator, len(vcb.builders))
	for i := range vcb.builders {
		func(i int, root context.Context) {
			builder := vcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				if err := builder.preSave(); err != nil {
					return nil, err
				}
				mutation, ok := m.(*ViewportMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, vcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, vcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, vcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (vcb *ViewportCreateBulk) SaveX(ctx context.Context) []*Viewport {
	v, err := vcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

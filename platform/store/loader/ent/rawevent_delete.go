// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/blushft/strana/platform/store/loader/ent/predicate"
	"github.com/blushft/strana/platform/store/loader/ent/rawevent"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// RawEventDelete is the builder for deleting a RawEvent entity.
type RawEventDelete struct {
	config
	hooks      []Hook
	mutation   *RawEventMutation
	predicates []predicate.RawEvent
}

// Where adds a new predicate to the delete builder.
func (red *RawEventDelete) Where(ps ...predicate.RawEvent) *RawEventDelete {
	red.predicates = append(red.predicates, ps...)
	return red
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (red *RawEventDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(red.hooks) == 0 {
		affected, err = red.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RawEventMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			red.mutation = mutation
			affected, err = red.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(red.hooks) - 1; i >= 0; i-- {
			mut = red.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, red.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (red *RawEventDelete) ExecX(ctx context.Context) int {
	n, err := red.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (red *RawEventDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: rawevent.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: rawevent.FieldID,
			},
		},
	}
	if ps := red.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, red.driver, _spec)
}

// RawEventDeleteOne is the builder for deleting a single RawEvent entity.
type RawEventDeleteOne struct {
	red *RawEventDelete
}

// Exec executes the deletion query.
func (redo *RawEventDeleteOne) Exec(ctx context.Context) error {
	n, err := redo.red.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{rawevent.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (redo *RawEventDeleteOne) ExecX(ctx context.Context) {
	redo.red.ExecX(ctx)
}

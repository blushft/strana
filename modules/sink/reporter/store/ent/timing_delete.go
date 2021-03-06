// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/blushft/strana/modules/sink/reporter/store/ent/predicate"
	"github.com/blushft/strana/modules/sink/reporter/store/ent/timing"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// TimingDelete is the builder for deleting a Timing entity.
type TimingDelete struct {
	config
	hooks      []Hook
	mutation   *TimingMutation
	predicates []predicate.Timing
}

// Where adds a new predicate to the delete builder.
func (td *TimingDelete) Where(ps ...predicate.Timing) *TimingDelete {
	td.predicates = append(td.predicates, ps...)
	return td
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (td *TimingDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(td.hooks) == 0 {
		affected, err = td.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TimingMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			td.mutation = mutation
			affected, err = td.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(td.hooks) - 1; i >= 0; i-- {
			mut = td.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, td.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (td *TimingDelete) ExecX(ctx context.Context) int {
	n, err := td.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (td *TimingDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: timing.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: timing.FieldID,
			},
		},
	}
	if ps := td.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, td.driver, _spec)
}

// TimingDeleteOne is the builder for deleting a single Timing entity.
type TimingDeleteOne struct {
	td *TimingDelete
}

// Exec executes the deletion query.
func (tdo *TimingDeleteOne) Exec(ctx context.Context) error {
	n, err := tdo.td.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{timing.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tdo *TimingDeleteOne) ExecX(ctx context.Context) {
	tdo.td.ExecX(ctx)
}

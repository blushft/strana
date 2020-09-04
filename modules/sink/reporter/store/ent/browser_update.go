// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/blushft/strana/modules/sink/reporter/store/ent/browser"
	"github.com/blushft/strana/modules/sink/reporter/store/ent/event"
	"github.com/blushft/strana/modules/sink/reporter/store/ent/predicate"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
)

// BrowserUpdate is the builder for updating Browser entities.
type BrowserUpdate struct {
	config
	hooks      []Hook
	mutation   *BrowserMutation
	predicates []predicate.Browser
}

// Where adds a new predicate for the builder.
func (bu *BrowserUpdate) Where(ps ...predicate.Browser) *BrowserUpdate {
	bu.predicates = append(bu.predicates, ps...)
	return bu
}

// SetName sets the name field.
func (bu *BrowserUpdate) SetName(s string) *BrowserUpdate {
	bu.mutation.SetName(s)
	return bu
}

// SetVersion sets the version field.
func (bu *BrowserUpdate) SetVersion(s string) *BrowserUpdate {
	bu.mutation.SetVersion(s)
	return bu
}

// SetUseragent sets the useragent field.
func (bu *BrowserUpdate) SetUseragent(s string) *BrowserUpdate {
	bu.mutation.SetUseragent(s)
	return bu
}

// SetNillableUseragent sets the useragent field if the given value is not nil.
func (bu *BrowserUpdate) SetNillableUseragent(s *string) *BrowserUpdate {
	if s != nil {
		bu.SetUseragent(*s)
	}
	return bu
}

// ClearUseragent clears the value of useragent.
func (bu *BrowserUpdate) ClearUseragent() *BrowserUpdate {
	bu.mutation.ClearUseragent()
	return bu
}

// AddEventIDs adds the events edge to Event by ids.
func (bu *BrowserUpdate) AddEventIDs(ids ...uuid.UUID) *BrowserUpdate {
	bu.mutation.AddEventIDs(ids...)
	return bu
}

// AddEvents adds the events edges to Event.
func (bu *BrowserUpdate) AddEvents(e ...*Event) *BrowserUpdate {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return bu.AddEventIDs(ids...)
}

// Mutation returns the BrowserMutation object of the builder.
func (bu *BrowserUpdate) Mutation() *BrowserMutation {
	return bu.mutation
}

// RemoveEventIDs removes the events edge to Event by ids.
func (bu *BrowserUpdate) RemoveEventIDs(ids ...uuid.UUID) *BrowserUpdate {
	bu.mutation.RemoveEventIDs(ids...)
	return bu
}

// RemoveEvents removes events edges to Event.
func (bu *BrowserUpdate) RemoveEvents(e ...*Event) *BrowserUpdate {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return bu.RemoveEventIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (bu *BrowserUpdate) Save(ctx context.Context) (int, error) {

	var (
		err      error
		affected int
	)
	if len(bu.hooks) == 0 {
		affected, err = bu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BrowserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bu.mutation = mutation
			affected, err = bu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(bu.hooks) - 1; i >= 0; i-- {
			mut = bu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (bu *BrowserUpdate) SaveX(ctx context.Context) int {
	affected, err := bu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bu *BrowserUpdate) Exec(ctx context.Context) error {
	_, err := bu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bu *BrowserUpdate) ExecX(ctx context.Context) {
	if err := bu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bu *BrowserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   browser.Table,
			Columns: browser.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: browser.FieldID,
			},
		},
	}
	if ps := bu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: browser.FieldName,
		})
	}
	if value, ok := bu.mutation.Version(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: browser.FieldVersion,
		})
	}
	if value, ok := bu.mutation.Useragent(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: browser.FieldUseragent,
		})
	}
	if bu.mutation.UseragentCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: browser.FieldUseragent,
		})
	}
	if nodes := bu.mutation.RemovedEventsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   browser.EventsTable,
			Columns: []string{browser.EventsColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.EventsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   browser.EventsTable,
			Columns: []string{browser.EventsColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{browser.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// BrowserUpdateOne is the builder for updating a single Browser entity.
type BrowserUpdateOne struct {
	config
	hooks    []Hook
	mutation *BrowserMutation
}

// SetName sets the name field.
func (buo *BrowserUpdateOne) SetName(s string) *BrowserUpdateOne {
	buo.mutation.SetName(s)
	return buo
}

// SetVersion sets the version field.
func (buo *BrowserUpdateOne) SetVersion(s string) *BrowserUpdateOne {
	buo.mutation.SetVersion(s)
	return buo
}

// SetUseragent sets the useragent field.
func (buo *BrowserUpdateOne) SetUseragent(s string) *BrowserUpdateOne {
	buo.mutation.SetUseragent(s)
	return buo
}

// SetNillableUseragent sets the useragent field if the given value is not nil.
func (buo *BrowserUpdateOne) SetNillableUseragent(s *string) *BrowserUpdateOne {
	if s != nil {
		buo.SetUseragent(*s)
	}
	return buo
}

// ClearUseragent clears the value of useragent.
func (buo *BrowserUpdateOne) ClearUseragent() *BrowserUpdateOne {
	buo.mutation.ClearUseragent()
	return buo
}

// AddEventIDs adds the events edge to Event by ids.
func (buo *BrowserUpdateOne) AddEventIDs(ids ...uuid.UUID) *BrowserUpdateOne {
	buo.mutation.AddEventIDs(ids...)
	return buo
}

// AddEvents adds the events edges to Event.
func (buo *BrowserUpdateOne) AddEvents(e ...*Event) *BrowserUpdateOne {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return buo.AddEventIDs(ids...)
}

// Mutation returns the BrowserMutation object of the builder.
func (buo *BrowserUpdateOne) Mutation() *BrowserMutation {
	return buo.mutation
}

// RemoveEventIDs removes the events edge to Event by ids.
func (buo *BrowserUpdateOne) RemoveEventIDs(ids ...uuid.UUID) *BrowserUpdateOne {
	buo.mutation.RemoveEventIDs(ids...)
	return buo
}

// RemoveEvents removes events edges to Event.
func (buo *BrowserUpdateOne) RemoveEvents(e ...*Event) *BrowserUpdateOne {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return buo.RemoveEventIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (buo *BrowserUpdateOne) Save(ctx context.Context) (*Browser, error) {

	var (
		err  error
		node *Browser
	)
	if len(buo.hooks) == 0 {
		node, err = buo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BrowserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			buo.mutation = mutation
			node, err = buo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(buo.hooks) - 1; i >= 0; i-- {
			mut = buo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, buo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (buo *BrowserUpdateOne) SaveX(ctx context.Context) *Browser {
	b, err := buo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return b
}

// Exec executes the query on the entity.
func (buo *BrowserUpdateOne) Exec(ctx context.Context) error {
	_, err := buo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buo *BrowserUpdateOne) ExecX(ctx context.Context) {
	if err := buo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (buo *BrowserUpdateOne) sqlSave(ctx context.Context) (b *Browser, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   browser.Table,
			Columns: browser.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: browser.FieldID,
			},
		},
	}
	id, ok := buo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Browser.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := buo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: browser.FieldName,
		})
	}
	if value, ok := buo.mutation.Version(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: browser.FieldVersion,
		})
	}
	if value, ok := buo.mutation.Useragent(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: browser.FieldUseragent,
		})
	}
	if buo.mutation.UseragentCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: browser.FieldUseragent,
		})
	}
	if nodes := buo.mutation.RemovedEventsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   browser.EventsTable,
			Columns: []string{browser.EventsColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.EventsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   browser.EventsTable,
			Columns: []string{browser.EventsColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	b = &Browser{config: buo.config}
	_spec.Assign = b.assignValues
	_spec.ScanValues = b.scanValues()
	if err = sqlgraph.UpdateNode(ctx, buo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{browser.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return b, nil
}
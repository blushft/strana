// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/blushft/strana/platform/store/ent/device"
	"github.com/blushft/strana/platform/store/ent/predicate"
	"github.com/blushft/strana/platform/store/ent/session"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/google/uuid"
)

// DeviceUpdate is the builder for updating Device entities.
type DeviceUpdate struct {
	config
	hooks      []Hook
	mutation   *DeviceMutation
	predicates []predicate.Device
}

// Where adds a new predicate for the builder.
func (du *DeviceUpdate) Where(ps ...predicate.Device) *DeviceUpdate {
	du.predicates = append(du.predicates, ps...)
	return du
}

// SetName sets the name field.
func (du *DeviceUpdate) SetName(s string) *DeviceUpdate {
	du.mutation.SetName(s)
	return du
}

// SetVersion sets the version field.
func (du *DeviceUpdate) SetVersion(s string) *DeviceUpdate {
	du.mutation.SetVersion(s)
	return du
}

// AddSessionIDs adds the sessions edge to Session by ids.
func (du *DeviceUpdate) AddSessionIDs(ids ...uuid.UUID) *DeviceUpdate {
	du.mutation.AddSessionIDs(ids...)
	return du
}

// AddSessions adds the sessions edges to Session.
func (du *DeviceUpdate) AddSessions(s ...*Session) *DeviceUpdate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return du.AddSessionIDs(ids...)
}

// Mutation returns the DeviceMutation object of the builder.
func (du *DeviceUpdate) Mutation() *DeviceMutation {
	return du.mutation
}

// RemoveSessionIDs removes the sessions edge to Session by ids.
func (du *DeviceUpdate) RemoveSessionIDs(ids ...uuid.UUID) *DeviceUpdate {
	du.mutation.RemoveSessionIDs(ids...)
	return du
}

// RemoveSessions removes sessions edges to Session.
func (du *DeviceUpdate) RemoveSessions(s ...*Session) *DeviceUpdate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return du.RemoveSessionIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (du *DeviceUpdate) Save(ctx context.Context) (int, error) {

	var (
		err      error
		affected int
	)
	if len(du.hooks) == 0 {
		affected, err = du.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DeviceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			du.mutation = mutation
			affected, err = du.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(du.hooks) - 1; i >= 0; i-- {
			mut = du.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, du.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (du *DeviceUpdate) SaveX(ctx context.Context) int {
	affected, err := du.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (du *DeviceUpdate) Exec(ctx context.Context) error {
	_, err := du.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (du *DeviceUpdate) ExecX(ctx context.Context) {
	if err := du.Exec(ctx); err != nil {
		panic(err)
	}
}

func (du *DeviceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   device.Table,
			Columns: device.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: device.FieldID,
			},
		},
	}
	if ps := du.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := du.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: device.FieldName,
		})
	}
	if value, ok := du.mutation.Version(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: device.FieldVersion,
		})
	}
	if nodes := du.mutation.RemovedSessionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   device.SessionsTable,
			Columns: []string{device.SessionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: session.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.SessionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   device.SessionsTable,
			Columns: []string{device.SessionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: session.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, du.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{device.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// DeviceUpdateOne is the builder for updating a single Device entity.
type DeviceUpdateOne struct {
	config
	hooks    []Hook
	mutation *DeviceMutation
}

// SetName sets the name field.
func (duo *DeviceUpdateOne) SetName(s string) *DeviceUpdateOne {
	duo.mutation.SetName(s)
	return duo
}

// SetVersion sets the version field.
func (duo *DeviceUpdateOne) SetVersion(s string) *DeviceUpdateOne {
	duo.mutation.SetVersion(s)
	return duo
}

// AddSessionIDs adds the sessions edge to Session by ids.
func (duo *DeviceUpdateOne) AddSessionIDs(ids ...uuid.UUID) *DeviceUpdateOne {
	duo.mutation.AddSessionIDs(ids...)
	return duo
}

// AddSessions adds the sessions edges to Session.
func (duo *DeviceUpdateOne) AddSessions(s ...*Session) *DeviceUpdateOne {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return duo.AddSessionIDs(ids...)
}

// Mutation returns the DeviceMutation object of the builder.
func (duo *DeviceUpdateOne) Mutation() *DeviceMutation {
	return duo.mutation
}

// RemoveSessionIDs removes the sessions edge to Session by ids.
func (duo *DeviceUpdateOne) RemoveSessionIDs(ids ...uuid.UUID) *DeviceUpdateOne {
	duo.mutation.RemoveSessionIDs(ids...)
	return duo
}

// RemoveSessions removes sessions edges to Session.
func (duo *DeviceUpdateOne) RemoveSessions(s ...*Session) *DeviceUpdateOne {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return duo.RemoveSessionIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (duo *DeviceUpdateOne) Save(ctx context.Context) (*Device, error) {

	var (
		err  error
		node *Device
	)
	if len(duo.hooks) == 0 {
		node, err = duo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DeviceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			duo.mutation = mutation
			node, err = duo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(duo.hooks) - 1; i >= 0; i-- {
			mut = duo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, duo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (duo *DeviceUpdateOne) SaveX(ctx context.Context) *Device {
	d, err := duo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return d
}

// Exec executes the query on the entity.
func (duo *DeviceUpdateOne) Exec(ctx context.Context) error {
	_, err := duo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duo *DeviceUpdateOne) ExecX(ctx context.Context) {
	if err := duo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (duo *DeviceUpdateOne) sqlSave(ctx context.Context) (d *Device, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   device.Table,
			Columns: device.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: device.FieldID,
			},
		},
	}
	id, ok := duo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Device.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := duo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: device.FieldName,
		})
	}
	if value, ok := duo.mutation.Version(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: device.FieldVersion,
		})
	}
	if nodes := duo.mutation.RemovedSessionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   device.SessionsTable,
			Columns: []string{device.SessionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: session.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.SessionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   device.SessionsTable,
			Columns: []string{device.SessionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: session.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	d = &Device{config: duo.config}
	_spec.Assign = d.assignValues
	_spec.ScanValues = d.scanValues()
	if err = sqlgraph.UpdateNode(ctx, duo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{device.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return d, nil
}

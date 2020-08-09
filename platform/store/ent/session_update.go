// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/blushft/strana/platform/store/ent/app"
	"github.com/blushft/strana/platform/store/ent/device"
	"github.com/blushft/strana/platform/store/ent/pageview"
	"github.com/blushft/strana/platform/store/ent/predicate"
	"github.com/blushft/strana/platform/store/ent/session"
	"github.com/blushft/strana/platform/store/ent/user"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/google/uuid"
)

// SessionUpdate is the builder for updating Session entities.
type SessionUpdate struct {
	config
	hooks      []Hook
	mutation   *SessionMutation
	predicates []predicate.Session
}

// Where adds a new predicate for the builder.
func (su *SessionUpdate) Where(ps ...predicate.Session) *SessionUpdate {
	su.predicates = append(su.predicates, ps...)
	return su
}

// SetNewUser sets the new_user field.
func (su *SessionUpdate) SetNewUser(b bool) *SessionUpdate {
	su.mutation.SetNewUser(b)
	return su
}

// SetIsUnique sets the is_unique field.
func (su *SessionUpdate) SetIsUnique(b bool) *SessionUpdate {
	su.mutation.SetIsUnique(b)
	return su
}

// SetIsBounce sets the is_bounce field.
func (su *SessionUpdate) SetIsBounce(b bool) *SessionUpdate {
	su.mutation.SetIsBounce(b)
	return su
}

// SetIsFinished sets the is_finished field.
func (su *SessionUpdate) SetIsFinished(b bool) *SessionUpdate {
	su.mutation.SetIsFinished(b)
	return su
}

// SetDuration sets the duration field.
func (su *SessionUpdate) SetDuration(i int) *SessionUpdate {
	su.mutation.ResetDuration()
	su.mutation.SetDuration(i)
	return su
}

// SetNillableDuration sets the duration field if the given value is not nil.
func (su *SessionUpdate) SetNillableDuration(i *int) *SessionUpdate {
	if i != nil {
		su.SetDuration(*i)
	}
	return su
}

// AddDuration adds i to duration.
func (su *SessionUpdate) AddDuration(i int) *SessionUpdate {
	su.mutation.AddDuration(i)
	return su
}

// ClearDuration clears the value of duration.
func (su *SessionUpdate) ClearDuration() *SessionUpdate {
	su.mutation.ClearDuration()
	return su
}

// SetStartedAt sets the started_at field.
func (su *SessionUpdate) SetStartedAt(t time.Time) *SessionUpdate {
	su.mutation.SetStartedAt(t)
	return su
}

// SetFinishedAt sets the finished_at field.
func (su *SessionUpdate) SetFinishedAt(t time.Time) *SessionUpdate {
	su.mutation.SetFinishedAt(t)
	return su
}

// SetNillableFinishedAt sets the finished_at field if the given value is not nil.
func (su *SessionUpdate) SetNillableFinishedAt(t *time.Time) *SessionUpdate {
	if t != nil {
		su.SetFinishedAt(*t)
	}
	return su
}

// ClearFinishedAt clears the value of finished_at.
func (su *SessionUpdate) ClearFinishedAt() *SessionUpdate {
	su.mutation.ClearFinishedAt()
	return su
}

// SetAppID sets the app edge to App by id.
func (su *SessionUpdate) SetAppID(id int) *SessionUpdate {
	su.mutation.SetAppID(id)
	return su
}

// SetApp sets the app edge to App.
func (su *SessionUpdate) SetApp(a *App) *SessionUpdate {
	return su.SetAppID(a.ID)
}

// SetUserID sets the user edge to User by id.
func (su *SessionUpdate) SetUserID(id string) *SessionUpdate {
	su.mutation.SetUserID(id)
	return su
}

// SetUser sets the user edge to User.
func (su *SessionUpdate) SetUser(u *User) *SessionUpdate {
	return su.SetUserID(u.ID)
}

// SetDeviceID sets the device edge to Device by id.
func (su *SessionUpdate) SetDeviceID(id string) *SessionUpdate {
	su.mutation.SetDeviceID(id)
	return su
}

// SetNillableDeviceID sets the device edge to Device by id if the given value is not nil.
func (su *SessionUpdate) SetNillableDeviceID(id *string) *SessionUpdate {
	if id != nil {
		su = su.SetDeviceID(*id)
	}
	return su
}

// SetDevice sets the device edge to Device.
func (su *SessionUpdate) SetDevice(d *Device) *SessionUpdate {
	return su.SetDeviceID(d.ID)
}

// AddPageviewIDs adds the pageviews edge to PageView by ids.
func (su *SessionUpdate) AddPageviewIDs(ids ...uuid.UUID) *SessionUpdate {
	su.mutation.AddPageviewIDs(ids...)
	return su
}

// AddPageviews adds the pageviews edges to PageView.
func (su *SessionUpdate) AddPageviews(p ...*PageView) *SessionUpdate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return su.AddPageviewIDs(ids...)
}

// Mutation returns the SessionMutation object of the builder.
func (su *SessionUpdate) Mutation() *SessionMutation {
	return su.mutation
}

// ClearApp clears the app edge to App.
func (su *SessionUpdate) ClearApp() *SessionUpdate {
	su.mutation.ClearApp()
	return su
}

// ClearUser clears the user edge to User.
func (su *SessionUpdate) ClearUser() *SessionUpdate {
	su.mutation.ClearUser()
	return su
}

// ClearDevice clears the device edge to Device.
func (su *SessionUpdate) ClearDevice() *SessionUpdate {
	su.mutation.ClearDevice()
	return su
}

// RemovePageviewIDs removes the pageviews edge to PageView by ids.
func (su *SessionUpdate) RemovePageviewIDs(ids ...uuid.UUID) *SessionUpdate {
	su.mutation.RemovePageviewIDs(ids...)
	return su
}

// RemovePageviews removes pageviews edges to PageView.
func (su *SessionUpdate) RemovePageviews(p ...*PageView) *SessionUpdate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return su.RemovePageviewIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (su *SessionUpdate) Save(ctx context.Context) (int, error) {

	if _, ok := su.mutation.AppID(); su.mutation.AppCleared() && !ok {
		return 0, errors.New("ent: clearing a unique edge \"app\"")
	}

	if _, ok := su.mutation.UserID(); su.mutation.UserCleared() && !ok {
		return 0, errors.New("ent: clearing a unique edge \"user\"")
	}

	var (
		err      error
		affected int
	)
	if len(su.hooks) == 0 {
		affected, err = su.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SessionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			su.mutation = mutation
			affected, err = su.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(su.hooks) - 1; i >= 0; i-- {
			mut = su.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, su.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (su *SessionUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SessionUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SessionUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

func (su *SessionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   session.Table,
			Columns: session.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: session.FieldID,
			},
		},
	}
	if ps := su.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.NewUser(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: session.FieldNewUser,
		})
	}
	if value, ok := su.mutation.IsUnique(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: session.FieldIsUnique,
		})
	}
	if value, ok := su.mutation.IsBounce(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: session.FieldIsBounce,
		})
	}
	if value, ok := su.mutation.IsFinished(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: session.FieldIsFinished,
		})
	}
	if value, ok := su.mutation.Duration(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: session.FieldDuration,
		})
	}
	if value, ok := su.mutation.AddedDuration(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: session.FieldDuration,
		})
	}
	if su.mutation.DurationCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: session.FieldDuration,
		})
	}
	if value, ok := su.mutation.StartedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: session.FieldStartedAt,
		})
	}
	if value, ok := su.mutation.FinishedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: session.FieldFinishedAt,
		})
	}
	if su.mutation.FinishedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: session.FieldFinishedAt,
		})
	}
	if su.mutation.AppCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   session.AppTable,
			Columns: []string{session.AppColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: app.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.AppIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   session.AppTable,
			Columns: []string{session.AppColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: app.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if su.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   session.UserTable,
			Columns: []string{session.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   session.UserTable,
			Columns: []string{session.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if su.mutation.DeviceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   session.DeviceTable,
			Columns: []string{session.DeviceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: device.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.DeviceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   session.DeviceTable,
			Columns: []string{session.DeviceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: device.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := su.mutation.RemovedPageviewsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   session.PageviewsTable,
			Columns: []string{session.PageviewsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: pageview.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.PageviewsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   session.PageviewsTable,
			Columns: []string{session.PageviewsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: pageview.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{session.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// SessionUpdateOne is the builder for updating a single Session entity.
type SessionUpdateOne struct {
	config
	hooks    []Hook
	mutation *SessionMutation
}

// SetNewUser sets the new_user field.
func (suo *SessionUpdateOne) SetNewUser(b bool) *SessionUpdateOne {
	suo.mutation.SetNewUser(b)
	return suo
}

// SetIsUnique sets the is_unique field.
func (suo *SessionUpdateOne) SetIsUnique(b bool) *SessionUpdateOne {
	suo.mutation.SetIsUnique(b)
	return suo
}

// SetIsBounce sets the is_bounce field.
func (suo *SessionUpdateOne) SetIsBounce(b bool) *SessionUpdateOne {
	suo.mutation.SetIsBounce(b)
	return suo
}

// SetIsFinished sets the is_finished field.
func (suo *SessionUpdateOne) SetIsFinished(b bool) *SessionUpdateOne {
	suo.mutation.SetIsFinished(b)
	return suo
}

// SetDuration sets the duration field.
func (suo *SessionUpdateOne) SetDuration(i int) *SessionUpdateOne {
	suo.mutation.ResetDuration()
	suo.mutation.SetDuration(i)
	return suo
}

// SetNillableDuration sets the duration field if the given value is not nil.
func (suo *SessionUpdateOne) SetNillableDuration(i *int) *SessionUpdateOne {
	if i != nil {
		suo.SetDuration(*i)
	}
	return suo
}

// AddDuration adds i to duration.
func (suo *SessionUpdateOne) AddDuration(i int) *SessionUpdateOne {
	suo.mutation.AddDuration(i)
	return suo
}

// ClearDuration clears the value of duration.
func (suo *SessionUpdateOne) ClearDuration() *SessionUpdateOne {
	suo.mutation.ClearDuration()
	return suo
}

// SetStartedAt sets the started_at field.
func (suo *SessionUpdateOne) SetStartedAt(t time.Time) *SessionUpdateOne {
	suo.mutation.SetStartedAt(t)
	return suo
}

// SetFinishedAt sets the finished_at field.
func (suo *SessionUpdateOne) SetFinishedAt(t time.Time) *SessionUpdateOne {
	suo.mutation.SetFinishedAt(t)
	return suo
}

// SetNillableFinishedAt sets the finished_at field if the given value is not nil.
func (suo *SessionUpdateOne) SetNillableFinishedAt(t *time.Time) *SessionUpdateOne {
	if t != nil {
		suo.SetFinishedAt(*t)
	}
	return suo
}

// ClearFinishedAt clears the value of finished_at.
func (suo *SessionUpdateOne) ClearFinishedAt() *SessionUpdateOne {
	suo.mutation.ClearFinishedAt()
	return suo
}

// SetAppID sets the app edge to App by id.
func (suo *SessionUpdateOne) SetAppID(id int) *SessionUpdateOne {
	suo.mutation.SetAppID(id)
	return suo
}

// SetApp sets the app edge to App.
func (suo *SessionUpdateOne) SetApp(a *App) *SessionUpdateOne {
	return suo.SetAppID(a.ID)
}

// SetUserID sets the user edge to User by id.
func (suo *SessionUpdateOne) SetUserID(id string) *SessionUpdateOne {
	suo.mutation.SetUserID(id)
	return suo
}

// SetUser sets the user edge to User.
func (suo *SessionUpdateOne) SetUser(u *User) *SessionUpdateOne {
	return suo.SetUserID(u.ID)
}

// SetDeviceID sets the device edge to Device by id.
func (suo *SessionUpdateOne) SetDeviceID(id string) *SessionUpdateOne {
	suo.mutation.SetDeviceID(id)
	return suo
}

// SetNillableDeviceID sets the device edge to Device by id if the given value is not nil.
func (suo *SessionUpdateOne) SetNillableDeviceID(id *string) *SessionUpdateOne {
	if id != nil {
		suo = suo.SetDeviceID(*id)
	}
	return suo
}

// SetDevice sets the device edge to Device.
func (suo *SessionUpdateOne) SetDevice(d *Device) *SessionUpdateOne {
	return suo.SetDeviceID(d.ID)
}

// AddPageviewIDs adds the pageviews edge to PageView by ids.
func (suo *SessionUpdateOne) AddPageviewIDs(ids ...uuid.UUID) *SessionUpdateOne {
	suo.mutation.AddPageviewIDs(ids...)
	return suo
}

// AddPageviews adds the pageviews edges to PageView.
func (suo *SessionUpdateOne) AddPageviews(p ...*PageView) *SessionUpdateOne {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return suo.AddPageviewIDs(ids...)
}

// Mutation returns the SessionMutation object of the builder.
func (suo *SessionUpdateOne) Mutation() *SessionMutation {
	return suo.mutation
}

// ClearApp clears the app edge to App.
func (suo *SessionUpdateOne) ClearApp() *SessionUpdateOne {
	suo.mutation.ClearApp()
	return suo
}

// ClearUser clears the user edge to User.
func (suo *SessionUpdateOne) ClearUser() *SessionUpdateOne {
	suo.mutation.ClearUser()
	return suo
}

// ClearDevice clears the device edge to Device.
func (suo *SessionUpdateOne) ClearDevice() *SessionUpdateOne {
	suo.mutation.ClearDevice()
	return suo
}

// RemovePageviewIDs removes the pageviews edge to PageView by ids.
func (suo *SessionUpdateOne) RemovePageviewIDs(ids ...uuid.UUID) *SessionUpdateOne {
	suo.mutation.RemovePageviewIDs(ids...)
	return suo
}

// RemovePageviews removes pageviews edges to PageView.
func (suo *SessionUpdateOne) RemovePageviews(p ...*PageView) *SessionUpdateOne {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return suo.RemovePageviewIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (suo *SessionUpdateOne) Save(ctx context.Context) (*Session, error) {

	if _, ok := suo.mutation.AppID(); suo.mutation.AppCleared() && !ok {
		return nil, errors.New("ent: clearing a unique edge \"app\"")
	}

	if _, ok := suo.mutation.UserID(); suo.mutation.UserCleared() && !ok {
		return nil, errors.New("ent: clearing a unique edge \"user\"")
	}

	var (
		err  error
		node *Session
	)
	if len(suo.hooks) == 0 {
		node, err = suo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SessionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			suo.mutation = mutation
			node, err = suo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(suo.hooks) - 1; i >= 0; i-- {
			mut = suo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, suo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SessionUpdateOne) SaveX(ctx context.Context) *Session {
	s, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return s
}

// Exec executes the query on the entity.
func (suo *SessionUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SessionUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (suo *SessionUpdateOne) sqlSave(ctx context.Context) (s *Session, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   session.Table,
			Columns: session.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: session.FieldID,
			},
		},
	}
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Session.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := suo.mutation.NewUser(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: session.FieldNewUser,
		})
	}
	if value, ok := suo.mutation.IsUnique(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: session.FieldIsUnique,
		})
	}
	if value, ok := suo.mutation.IsBounce(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: session.FieldIsBounce,
		})
	}
	if value, ok := suo.mutation.IsFinished(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: session.FieldIsFinished,
		})
	}
	if value, ok := suo.mutation.Duration(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: session.FieldDuration,
		})
	}
	if value, ok := suo.mutation.AddedDuration(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: session.FieldDuration,
		})
	}
	if suo.mutation.DurationCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: session.FieldDuration,
		})
	}
	if value, ok := suo.mutation.StartedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: session.FieldStartedAt,
		})
	}
	if value, ok := suo.mutation.FinishedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: session.FieldFinishedAt,
		})
	}
	if suo.mutation.FinishedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: session.FieldFinishedAt,
		})
	}
	if suo.mutation.AppCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   session.AppTable,
			Columns: []string{session.AppColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: app.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.AppIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   session.AppTable,
			Columns: []string{session.AppColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: app.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if suo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   session.UserTable,
			Columns: []string{session.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   session.UserTable,
			Columns: []string{session.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if suo.mutation.DeviceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   session.DeviceTable,
			Columns: []string{session.DeviceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: device.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.DeviceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   session.DeviceTable,
			Columns: []string{session.DeviceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: device.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := suo.mutation.RemovedPageviewsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   session.PageviewsTable,
			Columns: []string{session.PageviewsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: pageview.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.PageviewsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   session.PageviewsTable,
			Columns: []string{session.PageviewsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: pageview.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	s = &Session{config: suo.config}
	_spec.Assign = s.assignValues
	_spec.ScanValues = s.scanValues()
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{session.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return s, nil
}

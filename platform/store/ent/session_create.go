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
	"github.com/blushft/strana/platform/store/ent/session"
	"github.com/blushft/strana/platform/store/ent/user"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/google/uuid"
)

// SessionCreate is the builder for creating a Session entity.
type SessionCreate struct {
	config
	mutation *SessionMutation
	hooks    []Hook
}

// SetNewUser sets the new_user field.
func (sc *SessionCreate) SetNewUser(b bool) *SessionCreate {
	sc.mutation.SetNewUser(b)
	return sc
}

// SetIsUnique sets the is_unique field.
func (sc *SessionCreate) SetIsUnique(b bool) *SessionCreate {
	sc.mutation.SetIsUnique(b)
	return sc
}

// SetIsBounce sets the is_bounce field.
func (sc *SessionCreate) SetIsBounce(b bool) *SessionCreate {
	sc.mutation.SetIsBounce(b)
	return sc
}

// SetIsFinished sets the is_finished field.
func (sc *SessionCreate) SetIsFinished(b bool) *SessionCreate {
	sc.mutation.SetIsFinished(b)
	return sc
}

// SetDuration sets the duration field.
func (sc *SessionCreate) SetDuration(i int) *SessionCreate {
	sc.mutation.SetDuration(i)
	return sc
}

// SetNillableDuration sets the duration field if the given value is not nil.
func (sc *SessionCreate) SetNillableDuration(i *int) *SessionCreate {
	if i != nil {
		sc.SetDuration(*i)
	}
	return sc
}

// SetStartedAt sets the started_at field.
func (sc *SessionCreate) SetStartedAt(t time.Time) *SessionCreate {
	sc.mutation.SetStartedAt(t)
	return sc
}

// SetFinishedAt sets the finished_at field.
func (sc *SessionCreate) SetFinishedAt(t time.Time) *SessionCreate {
	sc.mutation.SetFinishedAt(t)
	return sc
}

// SetNillableFinishedAt sets the finished_at field if the given value is not nil.
func (sc *SessionCreate) SetNillableFinishedAt(t *time.Time) *SessionCreate {
	if t != nil {
		sc.SetFinishedAt(*t)
	}
	return sc
}

// SetID sets the id field.
func (sc *SessionCreate) SetID(u uuid.UUID) *SessionCreate {
	sc.mutation.SetID(u)
	return sc
}

// SetAppID sets the app edge to App by id.
func (sc *SessionCreate) SetAppID(id int) *SessionCreate {
	sc.mutation.SetAppID(id)
	return sc
}

// SetApp sets the app edge to App.
func (sc *SessionCreate) SetApp(a *App) *SessionCreate {
	return sc.SetAppID(a.ID)
}

// SetUserID sets the user edge to User by id.
func (sc *SessionCreate) SetUserID(id string) *SessionCreate {
	sc.mutation.SetUserID(id)
	return sc
}

// SetUser sets the user edge to User.
func (sc *SessionCreate) SetUser(u *User) *SessionCreate {
	return sc.SetUserID(u.ID)
}

// SetDeviceID sets the device edge to Device by id.
func (sc *SessionCreate) SetDeviceID(id string) *SessionCreate {
	sc.mutation.SetDeviceID(id)
	return sc
}

// SetNillableDeviceID sets the device edge to Device by id if the given value is not nil.
func (sc *SessionCreate) SetNillableDeviceID(id *string) *SessionCreate {
	if id != nil {
		sc = sc.SetDeviceID(*id)
	}
	return sc
}

// SetDevice sets the device edge to Device.
func (sc *SessionCreate) SetDevice(d *Device) *SessionCreate {
	return sc.SetDeviceID(d.ID)
}

// AddPageviewIDs adds the pageviews edge to PageView by ids.
func (sc *SessionCreate) AddPageviewIDs(ids ...uuid.UUID) *SessionCreate {
	sc.mutation.AddPageviewIDs(ids...)
	return sc
}

// AddPageviews adds the pageviews edges to PageView.
func (sc *SessionCreate) AddPageviews(p ...*PageView) *SessionCreate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return sc.AddPageviewIDs(ids...)
}

// Mutation returns the SessionMutation object of the builder.
func (sc *SessionCreate) Mutation() *SessionMutation {
	return sc.mutation
}

// Save creates the Session in the database.
func (sc *SessionCreate) Save(ctx context.Context) (*Session, error) {
	if _, ok := sc.mutation.NewUser(); !ok {
		return nil, &ValidationError{Name: "new_user", err: errors.New("ent: missing required field \"new_user\"")}
	}
	if _, ok := sc.mutation.IsUnique(); !ok {
		return nil, &ValidationError{Name: "is_unique", err: errors.New("ent: missing required field \"is_unique\"")}
	}
	if _, ok := sc.mutation.IsBounce(); !ok {
		return nil, &ValidationError{Name: "is_bounce", err: errors.New("ent: missing required field \"is_bounce\"")}
	}
	if _, ok := sc.mutation.IsFinished(); !ok {
		return nil, &ValidationError{Name: "is_finished", err: errors.New("ent: missing required field \"is_finished\"")}
	}
	if _, ok := sc.mutation.StartedAt(); !ok {
		return nil, &ValidationError{Name: "started_at", err: errors.New("ent: missing required field \"started_at\"")}
	}
	if _, ok := sc.mutation.AppID(); !ok {
		return nil, &ValidationError{Name: "app", err: errors.New("ent: missing required edge \"app\"")}
	}
	if _, ok := sc.mutation.UserID(); !ok {
		return nil, &ValidationError{Name: "user", err: errors.New("ent: missing required edge \"user\"")}
	}
	var (
		err  error
		node *Session
	)
	if len(sc.hooks) == 0 {
		node, err = sc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SessionMutation)
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
func (sc *SessionCreate) SaveX(ctx context.Context) *Session {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (sc *SessionCreate) sqlSave(ctx context.Context) (*Session, error) {
	s, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return s, nil
}

func (sc *SessionCreate) createSpec() (*Session, *sqlgraph.CreateSpec) {
	var (
		s     = &Session{config: sc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: session.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: session.FieldID,
			},
		}
	)
	if id, ok := sc.mutation.ID(); ok {
		s.ID = id
		_spec.ID.Value = id
	}
	if value, ok := sc.mutation.NewUser(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: session.FieldNewUser,
		})
		s.NewUser = value
	}
	if value, ok := sc.mutation.IsUnique(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: session.FieldIsUnique,
		})
		s.IsUnique = value
	}
	if value, ok := sc.mutation.IsBounce(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: session.FieldIsBounce,
		})
		s.IsBounce = value
	}
	if value, ok := sc.mutation.IsFinished(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: session.FieldIsFinished,
		})
		s.IsFinished = value
	}
	if value, ok := sc.mutation.Duration(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: session.FieldDuration,
		})
		s.Duration = value
	}
	if value, ok := sc.mutation.StartedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: session.FieldStartedAt,
		})
		s.StartedAt = value
	}
	if value, ok := sc.mutation.FinishedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: session.FieldFinishedAt,
		})
		s.FinishedAt = &value
	}
	if nodes := sc.mutation.AppIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.UserIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.DeviceIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.PageviewsIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return s, _spec
}

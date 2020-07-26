// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/blushft/strana/platform/store/ent/app"
	"github.com/blushft/strana/platform/store/ent/pageview"
	"github.com/blushft/strana/platform/store/ent/session"
	"github.com/blushft/strana/platform/store/ent/user"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/google/uuid"
)

// PageViewCreate is the builder for creating a PageView entity.
type PageViewCreate struct {
	config
	mutation *PageViewMutation
	hooks    []Hook
}

// SetHostname sets the hostname field.
func (pvc *PageViewCreate) SetHostname(s string) *PageViewCreate {
	pvc.mutation.SetHostname(s)
	return pvc
}

// SetPathname sets the pathname field.
func (pvc *PageViewCreate) SetPathname(s string) *PageViewCreate {
	pvc.mutation.SetPathname(s)
	return pvc
}

// SetReferrer sets the referrer field.
func (pvc *PageViewCreate) SetReferrer(s string) *PageViewCreate {
	pvc.mutation.SetReferrer(s)
	return pvc
}

// SetIsEntry sets the is_entry field.
func (pvc *PageViewCreate) SetIsEntry(b bool) *PageViewCreate {
	pvc.mutation.SetIsEntry(b)
	return pvc
}

// SetIsFinished sets the is_finished field.
func (pvc *PageViewCreate) SetIsFinished(b bool) *PageViewCreate {
	pvc.mutation.SetIsFinished(b)
	return pvc
}

// SetDuration sets the duration field.
func (pvc *PageViewCreate) SetDuration(i int) *PageViewCreate {
	pvc.mutation.SetDuration(i)
	return pvc
}

// SetTimestamp sets the timestamp field.
func (pvc *PageViewCreate) SetTimestamp(t time.Time) *PageViewCreate {
	pvc.mutation.SetTimestamp(t)
	return pvc
}

// SetUserAgent sets the user_agent field.
func (pvc *PageViewCreate) SetUserAgent(s string) *PageViewCreate {
	pvc.mutation.SetUserAgent(s)
	return pvc
}

// SetNillableUserAgent sets the user_agent field if the given value is not nil.
func (pvc *PageViewCreate) SetNillableUserAgent(s *string) *PageViewCreate {
	if s != nil {
		pvc.SetUserAgent(*s)
	}
	return pvc
}

// SetIPAddress sets the ip_address field.
func (pvc *PageViewCreate) SetIPAddress(s string) *PageViewCreate {
	pvc.mutation.SetIPAddress(s)
	return pvc
}

// SetNillableIPAddress sets the ip_address field if the given value is not nil.
func (pvc *PageViewCreate) SetNillableIPAddress(s *string) *PageViewCreate {
	if s != nil {
		pvc.SetIPAddress(*s)
	}
	return pvc
}

// SetScreenDim sets the screen_dim field.
func (pvc *PageViewCreate) SetScreenDim(s string) *PageViewCreate {
	pvc.mutation.SetScreenDim(s)
	return pvc
}

// SetNillableScreenDim sets the screen_dim field if the given value is not nil.
func (pvc *PageViewCreate) SetNillableScreenDim(s *string) *PageViewCreate {
	if s != nil {
		pvc.SetScreenDim(*s)
	}
	return pvc
}

// SetExtra sets the extra field.
func (pvc *PageViewCreate) SetExtra(m map[string]interface{}) *PageViewCreate {
	pvc.mutation.SetExtra(m)
	return pvc
}

// SetID sets the id field.
func (pvc *PageViewCreate) SetID(u uuid.UUID) *PageViewCreate {
	pvc.mutation.SetID(u)
	return pvc
}

// SetAppID sets the app edge to App by id.
func (pvc *PageViewCreate) SetAppID(id int) *PageViewCreate {
	pvc.mutation.SetAppID(id)
	return pvc
}

// SetNillableAppID sets the app edge to App by id if the given value is not nil.
func (pvc *PageViewCreate) SetNillableAppID(id *int) *PageViewCreate {
	if id != nil {
		pvc = pvc.SetAppID(*id)
	}
	return pvc
}

// SetApp sets the app edge to App.
func (pvc *PageViewCreate) SetApp(a *App) *PageViewCreate {
	return pvc.SetAppID(a.ID)
}

// SetSessionID sets the session edge to Session by id.
func (pvc *PageViewCreate) SetSessionID(id uuid.UUID) *PageViewCreate {
	pvc.mutation.SetSessionID(id)
	return pvc
}

// SetNillableSessionID sets the session edge to Session by id if the given value is not nil.
func (pvc *PageViewCreate) SetNillableSessionID(id *uuid.UUID) *PageViewCreate {
	if id != nil {
		pvc = pvc.SetSessionID(*id)
	}
	return pvc
}

// SetSession sets the session edge to Session.
func (pvc *PageViewCreate) SetSession(s *Session) *PageViewCreate {
	return pvc.SetSessionID(s.ID)
}

// SetUserID sets the user edge to User by id.
func (pvc *PageViewCreate) SetUserID(id string) *PageViewCreate {
	pvc.mutation.SetUserID(id)
	return pvc
}

// SetNillableUserID sets the user edge to User by id if the given value is not nil.
func (pvc *PageViewCreate) SetNillableUserID(id *string) *PageViewCreate {
	if id != nil {
		pvc = pvc.SetUserID(*id)
	}
	return pvc
}

// SetUser sets the user edge to User.
func (pvc *PageViewCreate) SetUser(u *User) *PageViewCreate {
	return pvc.SetUserID(u.ID)
}

// Mutation returns the PageViewMutation object of the builder.
func (pvc *PageViewCreate) Mutation() *PageViewMutation {
	return pvc.mutation
}

// Save creates the PageView in the database.
func (pvc *PageViewCreate) Save(ctx context.Context) (*PageView, error) {
	if _, ok := pvc.mutation.Hostname(); !ok {
		return nil, &ValidationError{Name: "hostname", err: errors.New("ent: missing required field \"hostname\"")}
	}
	if _, ok := pvc.mutation.Pathname(); !ok {
		return nil, &ValidationError{Name: "pathname", err: errors.New("ent: missing required field \"pathname\"")}
	}
	if _, ok := pvc.mutation.Referrer(); !ok {
		return nil, &ValidationError{Name: "referrer", err: errors.New("ent: missing required field \"referrer\"")}
	}
	if _, ok := pvc.mutation.IsEntry(); !ok {
		return nil, &ValidationError{Name: "is_entry", err: errors.New("ent: missing required field \"is_entry\"")}
	}
	if _, ok := pvc.mutation.IsFinished(); !ok {
		return nil, &ValidationError{Name: "is_finished", err: errors.New("ent: missing required field \"is_finished\"")}
	}
	if _, ok := pvc.mutation.Duration(); !ok {
		return nil, &ValidationError{Name: "duration", err: errors.New("ent: missing required field \"duration\"")}
	}
	if _, ok := pvc.mutation.Timestamp(); !ok {
		return nil, &ValidationError{Name: "timestamp", err: errors.New("ent: missing required field \"timestamp\"")}
	}
	var (
		err  error
		node *PageView
	)
	if len(pvc.hooks) == 0 {
		node, err = pvc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PageViewMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pvc.mutation = mutation
			node, err = pvc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(pvc.hooks) - 1; i >= 0; i-- {
			mut = pvc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pvc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pvc *PageViewCreate) SaveX(ctx context.Context) *PageView {
	v, err := pvc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pvc *PageViewCreate) sqlSave(ctx context.Context) (*PageView, error) {
	pv, _spec := pvc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pvc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return pv, nil
}

func (pvc *PageViewCreate) createSpec() (*PageView, *sqlgraph.CreateSpec) {
	var (
		pv    = &PageView{config: pvc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: pageview.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: pageview.FieldID,
			},
		}
	)
	if id, ok := pvc.mutation.ID(); ok {
		pv.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pvc.mutation.Hostname(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pageview.FieldHostname,
		})
		pv.Hostname = value
	}
	if value, ok := pvc.mutation.Pathname(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pageview.FieldPathname,
		})
		pv.Pathname = value
	}
	if value, ok := pvc.mutation.Referrer(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pageview.FieldReferrer,
		})
		pv.Referrer = value
	}
	if value, ok := pvc.mutation.IsEntry(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: pageview.FieldIsEntry,
		})
		pv.IsEntry = value
	}
	if value, ok := pvc.mutation.IsFinished(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: pageview.FieldIsFinished,
		})
		pv.IsFinished = value
	}
	if value, ok := pvc.mutation.Duration(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: pageview.FieldDuration,
		})
		pv.Duration = value
	}
	if value, ok := pvc.mutation.Timestamp(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: pageview.FieldTimestamp,
		})
		pv.Timestamp = value
	}
	if value, ok := pvc.mutation.UserAgent(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pageview.FieldUserAgent,
		})
		pv.UserAgent = value
	}
	if value, ok := pvc.mutation.IPAddress(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pageview.FieldIPAddress,
		})
		pv.IPAddress = value
	}
	if value, ok := pvc.mutation.ScreenDim(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pageview.FieldScreenDim,
		})
		pv.ScreenDim = value
	}
	if value, ok := pvc.mutation.Extra(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: pageview.FieldExtra,
		})
		pv.Extra = value
	}
	if nodes := pvc.mutation.AppIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   pageview.AppTable,
			Columns: []string{pageview.AppColumn},
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
	if nodes := pvc.mutation.SessionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   pageview.SessionTable,
			Columns: []string{pageview.SessionColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pvc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   pageview.UserTable,
			Columns: []string{pageview.UserColumn},
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
	return pv, _spec
}

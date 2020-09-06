package entity

import (
	"context"

	"github.com/blushft/strana/event/contexts"
	"github.com/blushft/strana/modules/sink/reporter/store"
	"github.com/blushft/strana/modules/sink/reporter/store/ent"
	"github.com/blushft/strana/modules/sink/reporter/store/ent/action"
	"github.com/blushft/strana/platform/logger"
	"github.com/google/uuid"
	"github.com/vmihailenco/msgpack"
)

type Action struct {
	ID       int `json:"id"`
	Action   string
	Category string
	Label    string
	Property string
	Value    interface{}
}

func NewAction(a contexts.Action) *Action {
	return &Action{
		Action:   a.Action,
		Category: a.Category,
		Label:    a.Label,
		Property: a.Property,
		Value:    a.Value,
	}
}

type ActionReader interface {
	List() ([]*Action, error)
	Get(id int) (*Action, error)
}

type ActionWriter interface {
	Create(*Action, ...CreateActionEdge) error
	Update(*Action, ...UpdateActionEdge) error
	Delete(*Action) error
}

type actionRepo interface {
	ActionReader
	ActionWriter
}

type ActionManager interface {
	actionRepo
}

type actionManager struct {
	store *store.Store
}

func NewActionService(s *store.Store) *actionManager {
	return &actionManager{
		store: s,
	}
}

func (mgr *actionManager) List() ([]*Action, error) {
	c := mgr.store.Client().Action

	recs, err := c.Query().All(context.TODO())
	if err != nil {
		return nil, err
	}

	res := make([]*Action, 0, len(recs))
	for _, rec := range recs {
		res = append(res, actionSchemaToEntity(rec))
	}

	return res, nil
}

func (mgr *actionManager) Get(id int) (*Action, error) {
	c := mgr.store.Client().Action

	rec, err := c.Get(context.TODO(), id)
	if err != nil {
		return nil, err
	}

	return actionSchemaToEntity(rec), nil
}

func (mgr *actionManager) Create(a *Action, edges ...CreateActionEdge) error {
	c := mgr.store.Client().Action

	sa, err := actionEntityCreate(c, a, edges...).Save(context.TODO())
	if err != nil {
		return err
	}

	a.ID = sa.ID

	return nil
}

func (mgr *actionManager) Update(a *Action, edges ...UpdateActionEdge) error {
	c := mgr.store.Client().Action

	_, err := actionEntityUpdate(c, a, edges...).Save(context.TODO())

	return err
}

func (mgr *actionManager) Delete(a *Action) error {
	return mgr.store.Client().Action.DeleteOneID(a.ID).Exec(context.TODO())
}

func actionSchemaToEntity(sch *ent.Action) *Action {
	v, err := decValue(sch.Value)
	if err != nil {
		logger.Log().WithError(err).Error("decoding action value")
	}

	return &Action{
		ID:       sch.ID,
		Action:   sch.Action,
		Label:    sch.ActionLabel,
		Property: sch.ActionLabel,
		Value:    v,
	}
}

func actionEntityCreate(c *ent.ActionClient, e *Action, edges ...CreateActionEdge) *ent.ActionCreate {
	b, err := encValue(e.Value)
	if err != nil {
		logger.Log().WithError(err).Error("encode value")
	}

	cr := c.Create().
		SetAction(e.Action).
		SetCategory(e.Category).
		SetActionLabel(e.Label).
		SetProperty(e.Property).
		SetValue(b)

	for _, edge := range edges {
		edge(cr)
	}

	return cr
}

func actionEntityUpdate(c *ent.ActionClient, e *Action, edges ...UpdateActionEdge) *ent.ActionUpdate {
	b, err := encValue(e.Value)
	if err != nil {
		logger.Log().WithError(err).Error("encode value")
	}

	up := c.Update().
		SetAction(e.Action).
		SetCategory(e.Category).
		SetActionLabel(e.Label).
		SetProperty(e.Property).
		SetValue(b).
		Where(action.ID(e.ID))

	for _, edge := range edges {
		edge(up)
	}

	return up
}

type CreateActionEdge func(*ent.ActionCreate)
type UpdateActionEdge func(*ent.ActionUpdate)

func ActionWithEventEdge(id uuid.UUID) CreateActionEdge {
	return func(c *ent.ActionCreate) {
		c.SetEventID(id)
	}
}

func encValue(v interface{}) ([]byte, error) {
	return msgpack.Marshal(v)
}

func decValue(b []byte) (interface{}, error) {
	var v interface{}
	if err := msgpack.Unmarshal(b, v); err != nil {
		return nil, err
	}

	return v, nil
}

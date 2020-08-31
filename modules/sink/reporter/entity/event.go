package entity

import (
	"context"
	"time"

	"github.com/blushft/strana/modules/sink/reporter/store"
	"github.com/blushft/strana/modules/sink/reporter/store/ent"
	"github.com/google/uuid"
)

type Event struct {
	ID         uuid.UUID
	TrackingID string
	UserID     string
	GroupID    int
	SessionID  uuid.UUID
	DeviceID   string

	Type      string
	Channel   string
	Platform  string
	Timestamp time.Time
}

type EventReader interface {
	List(QueryParams) ([]*Event, error)
	Get(id uuid.UUID) (*Event, error)
}

type EventWriter interface {
	Create(*Event) error
	Update(*Event) error
	Delete(*Event) error
}

type eventRepo interface {
	EventReader
	EventWriter
}

type EventManager interface {
	eventRepo
}

type eventManager struct {
	store *store.Store
}

func NewEventService(s *store.Store) *eventManager {
	return &eventManager{
		store: s,
	}
}

func (mgr *eventManager) List(qp QueryParams) ([]*Event, error) {
	c := mgr.store.Client().Event
	var res []*Event

	q := c.Query()

	if qp.Limit > 0 {
		q.Limit(qp.Limit)

		if qp.Offset > 0 {
			q.Offset(qp.Offset)
		}
	}

	recs, err := eventGraphEdges(q).All(context.TODO())
	if err != nil {
		return nil, err
	}

	for _, rec := range recs {
		res = append(res, eventSchemaToEntity(rec))
	}

	return res, nil

}

func eventGraphEdges(q *ent.EventQuery) *ent.EventQuery {
	return q.WithAction().
		WithDevice().
		WithGroup().
		WithSession().
		WithUser()
}

func eventSchemaToEntity(sch *ent.Event) *Event {
	uid := ""
	if sch.Edges.User != nil {
		uid = sch.Edges.User.ID
	}

	gid := 0
	if sch.Edges.Group != nil {
		gid = sch.Edges.Group.ID
	}

	sid := uuid.New()
	if sch.Edges.Session != nil {
		sid = sch.Edges.Session.ID
	}

	cid := ""
	if sch.Edges.Device != nil {
		cid = sch.Edges.Device.ID
	}

	return &Event{
		ID:         sch.ID,
		TrackingID: sch.TrackingID,
		UserID:     uid,
		GroupID:    gid,
		SessionID:  sid,
		DeviceID:   cid,
		Type:       sch.Event.String(),
		Channel:    sch.Channel,
		Platform:   sch.Platform,
		Timestamp:  sch.Timestamp,
	}
}

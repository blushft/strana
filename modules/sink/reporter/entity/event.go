package entity

import (
	"context"
	"strconv"
	"time"

	revent "github.com/blushft/strana/event"
	"github.com/blushft/strana/modules/sink/reporter/store"
	"github.com/blushft/strana/modules/sink/reporter/store/ent"
	"github.com/blushft/strana/modules/sink/reporter/store/ent/event"
	"github.com/google/uuid"
)

type Event struct {
	ID         uuid.UUID `json:"id"`
	TrackingID string    `json:"trackingID"`
	UserID     string    `json:"userID"`
	GroupID    int       `json:"groupID"`
	SessionID  uuid.UUID `json:"sessionID"`
	DeviceID   string    `json:"deviceID"`

	Anonymous      bool `json:"anonymous"`
	NonInteractive bool `json:"nonInteractive"`

	Type      string    `json:"type"`
	Channel   string    `json:"channel"`
	Platform  string    `json:"platform"`
	Timestamp time.Time `json:"timestamp"`
}

func NewEvent(evt *revent.Event) (*Event, error) {
	id, err := uuid.Parse(evt.ID)
	if err != nil {
		return nil, err
	}

	if evt.GroupID != "" {

	}

	var sid uuid.UUID
	if evt.SessionID != "" {
		sid, err = uuid.Parse(evt.SessionID)
		if err != nil {
			return nil, err
		}
	}

	var gid int
	if evt.GroupID != "" {
		gid, err = strconv.Atoi(evt.GroupID)
		if err != nil {
			return nil, err
		}
	}

	return &Event{
		ID:             id,
		TrackingID:     evt.TrackingID,
		UserID:         evt.UserID,
		GroupID:        gid,
		SessionID:      sid,
		DeviceID:       evt.DeviceID,
		Anonymous:      evt.Anonymous,
		NonInteractive: evt.NonInteractive,
		Type:           string(evt.Event),
		Channel:        evt.Channel,
		Platform:       evt.Platform,
		Timestamp:      evt.Timestamp,
	}, nil
}

type EventReader interface {
	List(QueryParams) ([]*Event, error)
	Get(id uuid.UUID) (*Event, error)
	Query(...QueryEvent) ([]*Event, error)
}

type EventWriter interface {
	Create(*Event, ...CreateEventEdge) error
	Update(*Event, ...UpdateEventEdge) error
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

func (mgr *eventManager) Get(id uuid.UUID) (*Event, error) {
	c := mgr.store.Client().Event

	q := eventGraphEdges(c.Query().Where(event.ID(id)))

	e, err := q.Only(context.TODO())
	if err != nil {
		return nil, err
	}

	return eventSchemaToEntity(e), nil
}

func (mgr *eventManager) Query(qs ...QueryEvent) ([]*Event, error) {
	c := mgr.store.Client().Event

	q := eventGraphEdges(c.Query())

	for _, qe := range qs {
		qe(q)
	}

	recs, err := q.All(context.TODO())
	if err != nil {
		return nil, err
	}

	res := make([]*Event, 0, len(recs))
	for _, rec := range recs {
		res = append(res, eventSchemaToEntity(rec))
	}

	return res, nil
}

func (mgr *eventManager) Create(e *Event, edges ...CreateEventEdge) error {
	c := mgr.store.Client().Event
	cr := eventEntityCreate(c, e)

	for _, edge := range edges {
		edge(cr)
	}

	_, err := cr.Save(context.TODO())

	return err
}

func (mgr *eventManager) Update(e *Event, edges ...UpdateEventEdge) error {
	c := mgr.store.Client().Event
	up := eventEntityUpdate(c, e)

	for _, edge := range edges {
		edge(up)
	}

	_, err := up.Save(context.TODO())

	return err
}

func (mgr *eventManager) Delete(e *Event) error {
	c := mgr.store.Client().Event
	return c.DeleteOneID(e.ID).Exec(context.TODO())
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
		ID:             sch.ID,
		TrackingID:     sch.TrackingID,
		UserID:         uid,
		GroupID:        gid,
		SessionID:      sid,
		DeviceID:       cid,
		NonInteractive: sch.NonInteractive,
		Type:           sch.Event.String(),
		Channel:        sch.Channel,
		Platform:       sch.Platform,
		Timestamp:      sch.Timestamp,
	}
}

type CreateEventEdge func(c *ent.EventCreate)
type UpdateEventEdge func(c *ent.EventUpdate)

func eventEntityCreate(c *ent.EventClient, e *Event) *ent.EventCreate {
	et := event.Event(e.Type)
	cr := c.Create().
		SetID(e.ID).
		SetTrackingID(e.TrackingID).
		SetEvent(et).
		SetChannel(e.Channel).
		SetPlatform(e.Platform).
		SetTimestamp(e.Timestamp).
		SetNonInteractive(e.NonInteractive)

	if e.UserID != "" {
		cr.SetUserID(e.UserID)
	}

	return cr
}

func eventEntityUpdate(c *ent.EventClient, e *Event) *ent.EventUpdate {
	et := event.Event(e.Type)
	up := c.Update().
		SetTrackingID(e.TrackingID).
		SetEvent(et).
		SetChannel(e.Channel).
		SetPlatform(e.Platform).
		SetTimestamp(e.Timestamp).
		SetNonInteractive(e.NonInteractive).
		Where(event.ID(e.ID))

	if e.UserID != "" {
		up.SetUserID(e.UserID)
	}

	return up
}

package entity

import (
	"context"
	"time"

	"github.com/blushft/strana/event"
	"github.com/blushft/strana/platform/store/loader"
	"github.com/blushft/strana/platform/store/loader/ent"
	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
)

type RawEvent struct {
	*event.Event
}

type RawEventQuery struct {
	TrackingID string    `json:"tid,omitempty" query:"tid"`
	UserID     string    `json:"uid,omitempty" query:"uid"`
	GroupID    string    `json:"gid,omitempty" query:"gid"`
	SessionID  string    `json:"sid,omitempty" query:"sid"`
	DeviceID   string    `json:"cid,omitempty" query:"cid"`
	Event      string    `json:"e,omitempty" query:"e"`
	Channel    string    `json:"c,omitempty" query:"c"`
	Platform   string    `json:"p,omitempty" query:"p"`
	Before     time.Time `json:"before,omitempty" query:"before"`
	After      time.Time `json:"after,omitempty" query:"after"`
}

type RawEventReader interface {
	List() ([]*RawEvent, error)
	Get(uuid.UUID) (*RawEvent, error)
}

type RawEventWriter interface {
	Create(*RawEvent) error
	Update(*RawEvent) error
	Delete(*RawEvent) error
}

type rawEventRepo interface {
	RawEventReader
	RawEventWriter
}

type RawEventManager interface {
	rawEventRepo
}

type RawEventReporter interface {
	RawEventReader
}

type rawEventManager struct {
	store *loader.Store
}

func NewRawEventService(s *loader.Store) *rawEventManager {
	return &rawEventManager{
		store: s,
	}
}

func (mgr *rawEventManager) List() ([]*RawEvent, error) {
	c := mgr.store.Client().RawEvent
	var res []*RawEvent

	recs, err := c.Query().All(context.TODO())
	if err != nil {
		return nil, err
	}

	for _, rec := range recs {
		re, err := reSchemaToEntity(rec)
		if err != nil {
			return nil, err
		}

		res = append(res, re)
	}

	return res, nil
}

func (mgr *rawEventManager) Get(id uuid.UUID) (*RawEvent, error) {
	c := mgr.store.Client().RawEvent

	rec, err := c.Get(context.TODO(), id)
	if err != nil {
		return nil, err
	}

	return reSchemaToEntity(rec)
}

func (mgr *rawEventManager) Query(q RawEventQuery) ([]*RawEvent, error) {
	panic("not implemented")
}

func (mgr *rawEventManager) Create(re *RawEvent) error {
	c := mgr.store.Client().RawEvent

	_, err := reEntityCreate(c, re).Save(context.TODO())

	return err
}

func (mgr *rawEventManager) Update(re *RawEvent) error {
	c := mgr.store.Client().RawEvent

	_, err := reEntityUpdate(c, re).Save(context.TODO())

	return err
}

func (mgr *rawEventManager) Delete(re *RawEvent) error {
	c := mgr.store.Client().RawEvent

	id, err := uuid.Parse(re.ID)
	if err != nil {
		return err
	}

	return c.DeleteOneID(id).Exec(context.TODO())
}

func reSchemaToEntity(sch *ent.RawEvent) (*RawEvent, error) {
	evt := &event.Event{
		ID:             sch.ID.String(),
		TrackingID:     sch.TrackingID.String(),
		UserID:         sch.UserID,
		Anonymous:      sch.Anonymous,
		GroupID:        sch.GroupID,
		SessionID:      sch.SessionID,
		DeviceID:       sch.DeviceID,
		Event:          event.Type(sch.Event),
		NonInteractive: sch.NonInteractive,
		Channel:        sch.Channel,
		Platform:       sch.Platform,
		Timestamp:      sch.Timestamp,
	}

	ctx := event.Contexts{}
	if err := mapstructure.WeakDecode(sch.Context, &ctx); err != nil {
		return nil, err
	}

	evt.Context = ctx

	return &RawEvent{
		Event: evt,
	}, nil
}

func reEntityCreate(c *ent.RawEventClient, re *RawEvent) *ent.RawEventCreate {
	return c.Create().
		SetID(uuid.MustParse(re.ID)).
		SetTrackingID(uuid.MustParse(re.TrackingID)).
		SetUserID(re.UserID).
		SetAnonymous(re.Anonymous).
		SetGroupID(re.GroupID).
		SetSessionID(re.SessionID).
		SetDeviceID(re.DeviceID).
		SetEvent(string(re.Event.Event)).
		SetNonInteractive(re.NonInteractive).
		SetChannel(re.Channel).
		SetPlatform(re.Platform).
		SetTimestamp(re.Timestamp).
		SetContext(re.Context.Map())
}

func reEntityUpdate(c *ent.RawEventClient, re *RawEvent) *ent.RawEventUpdate {
	return c.Update().
		SetTrackingID(uuid.MustParse(re.TrackingID)).
		SetUserID(re.UserID).
		SetAnonymous(re.Anonymous).
		SetGroupID(re.GroupID).
		SetSessionID(re.SessionID).
		SetDeviceID(re.DeviceID).
		SetEvent(string(re.Event.Event)).
		SetNonInteractive(re.NonInteractive).
		SetChannel(re.Channel).
		SetPlatform(re.Platform).
		SetTimestamp(re.Timestamp).
		SetContext(re.Context.Map())
}

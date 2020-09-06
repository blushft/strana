package entity

import (
	"context"
	"time"

	"github.com/blushft/strana/modules/sink/reporter/store/ent"
	"github.com/blushft/strana/modules/sink/reporter/store/ent/event"
	"github.com/google/uuid"
)

const (
	DefaultLimit  = 10
	DefaultOffset = 0
)

type QueryParams struct {
	DeviceIDs  []string          `json:"cid" query:"cid"`
	SessionIDs []uuid.UUID       `json:"sid" query:"sid"`
	UserIDs    []uuid.UUID       `json:"uid" query:"uid"`
	GroupIDs   []int             `json:"gid" query:"gid"`
	Offset     int               `json:"offset" query:"offset"`
	Limit      int               `json:"limit" query:"limit"`
	Start      int               `json:"start" query:"start"`
	End        int               `json:"end" query:"end"`
	Params     map[string]string `json:"params" query:"params"`
}

func (qp QueryParams) QueryEvent() []QueryEvent {
	var qs []QueryEvent

	return qs
}

type QueryEvent func(q *ent.EventQuery)

type EventQuery interface {
	All(context.Context) ([]*ent.Event, error)
}
type ScannableQuery interface {
	Scan(context.Context, interface{}) error
}

type EventQueryBuild func(q *ent.EventQuery, qs ...QueryEvent)
type ScannableQueryBuild func(q *ent.EventQuery, qs ...QueryEvent) ScannableQuery

func NewEventQuery(q *ent.EventQuery, qs ...QueryEvent) {
	for _, eq := range qs {
		eq(q)
	}
}

func EventLimit(l int) QueryEvent {
	return func(eq *ent.EventQuery) {
		eq.Limit(l)
	}
}

func EventOffset(o int) QueryEvent {
	return func(eq *ent.EventQuery) {
		eq.Offset(o)
	}
}

func EventHasAction() QueryEvent {
	return func(eq *ent.EventQuery) {
		eq.Where(event.HasAction())
	}
}

func EventTimeBefore(t time.Time) QueryEvent {
	return func(eq *ent.EventQuery) {
		eq.Where(event.TimestampLT(t))
	}
}

func EventTimeAfter(t time.Time) QueryEvent {
	return func(eq *ent.EventQuery) {
		eq.Where(event.TimestampGT(t))
	}
}

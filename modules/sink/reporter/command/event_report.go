package command

import (
	"github.com/blushft/strana/modules/sink/reporter/entity"
	"github.com/blushft/strana/modules/sink/reporter/store"
	pstore "github.com/blushft/strana/platform/store"
)

type EventReporter struct {
	s *store.Store

	evtsvc entity.EventReader
	aggsvc entity.EventAggregate
}

func NewEventReporter(dbs *pstore.SQLStore) (*EventReporter, error) {
	rs, err := store.New(dbs)
	if err != nil {
		return nil, err
	}

	return &EventReporter{
		s:      rs,
		evtsvc: entity.NewEventService(rs),
		aggsvc: entity.NewEventAggregate(rs),
	}, nil
}

func (rpt *EventReporter) Events(params entity.QueryParams) ([]*entity.Event, error) {
	return rpt.evtsvc.List(params)
}

func (rpt *EventReporter) EventsCount() (int, error) {
	return rpt.aggsvc.Count()
}

func (rpt *EventReporter) EventsWithAction(params entity.QueryParams) ([]*entity.Event, error) {
	return rpt.evtsvc.Query(entity.EventHasAction())
}

type CategoryCount struct {
	Category string `json:"category"`
	Count    int    `json:"count"`
}

func (rpt *EventReporter) TopActionCatgories(params entity.QueryParams) ([]*CategoryCount, error) {
	var cc []*CategoryCount
	if err := rpt.aggsvc.GroupBy(entity.TopActionCategories, &cc, params.QueryEvent()...); err != nil {
		return nil, err
	}

	return cc, nil
}

package command

import (
	revent "github.com/blushft/strana/event"
	"github.com/blushft/strana/event/contexts"
	"github.com/blushft/strana/modules/sink/reporter/entity"
	"github.com/blushft/strana/modules/sink/reporter/store"
	sstore "github.com/blushft/strana/platform/store"
)

type EventExtractor struct {
	rs *store.Store

	actsvc entity.ActionManager
	evtsvc entity.EventManager
}

func NewEventExtractor(s *sstore.SQLStore) (*EventExtractor, error) {
	rs, err := store.New(s)
	if err != nil {
		return nil, err
	}

	return &EventExtractor{
		actsvc: entity.NewActionService(rs),
		evtsvc: entity.NewEventService(rs),
	}, nil
}

func (ee *EventExtractor) Save(evt *revent.Event) error {
	ne, err := entity.NewEvent(evt)
	if err != nil {
		return err
	}

	if err := ee.evtsvc.Create(ne); err != nil {
		return err
	}

	if ne.Type == "action" {
		if err := ee.extractAction(evt, ne); err != nil {
			return err
		}
	}

	return nil
}

func (ee *EventExtractor) extractAction(revt *revent.Event, evt *entity.Event) error {
	a, ok := revt.Context.Get(contexts.ContextAction)
	if ok {
		ea, ok := a.Interface().(*contexts.Action)
		if ok {
			na := entity.NewAction(*ea)
			return ee.actsvc.Create(na, entity.ActionWithEventEdge(evt.ID))
		}
	}

	return nil
}

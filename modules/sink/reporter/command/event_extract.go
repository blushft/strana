package command

import (
	"context"

	revent "github.com/blushft/strana/event"
	"github.com/blushft/strana/modules/sink/reporter/entity"
	"github.com/blushft/strana/modules/sink/reporter/store"
	sstore "github.com/blushft/strana/platform/store"
)

type EventExtractor struct {
	rs *store.Store

	txsvc  entity.EventTransaction
	actsvc entity.ActionManager
	evtsvc entity.EventManager
}

func NewEventExtractor(s *sstore.SQLStore) (*EventExtractor, error) {
	rs, err := store.New(s)
	if err != nil {
		return nil, err
	}

	return &EventExtractor{
		txsvc:  entity.NewEventTransaction(rs),
		actsvc: entity.NewActionService(rs),
		evtsvc: entity.NewEventService(rs),
	}, nil
}

func (ee *EventExtractor) Save(evt *revent.Event) error {
	if err := ee.txsvc.NewEvent(context.Background(), evt); err != nil {
		return err
	}

	return nil
}

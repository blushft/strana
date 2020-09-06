package entity

import (
	"context"

	"github.com/blushft/strana/modules/sink/reporter/store"
	"github.com/blushft/strana/modules/sink/reporter/store/ent"
	"github.com/blushft/strana/modules/sink/reporter/store/ent/action"
)

type EventAggregate interface {
	Count(...QueryEvent) (int, error)
	GroupBy(ScannableQueryBuild, interface{}, ...QueryEvent) error
}

type evtAggregate struct {
	store *store.Store
}

func NewEventAggregate(s *store.Store) *evtAggregate {
	return &evtAggregate{
		store: s,
	}
}

func (agg *evtAggregate) Count(qs ...QueryEvent) (int, error) {
	c := agg.store.Client().Event

	q := c.Query()

	for _, aq := range qs {
		aq(q)
	}

	return q.Count(context.TODO())
}

func (agg *evtAggregate) GroupBy(sq ScannableQueryBuild, v interface{}, qs ...QueryEvent) error {
	c := agg.store.Client().Event

	q := c.Query()

	return sq(q, qs...).Scan(context.TODO(), v)
}

func TopActionCategories(q *ent.EventQuery, qs ...QueryEvent) ScannableQuery {
	NewEventQuery(q, qs...)

	return q.QueryAction().GroupBy(action.FieldCategory).Aggregate(ent.Count())
}

package entity

import (
	"context"

	"github.com/blushft/strana/event"
	"github.com/blushft/strana/event/contexts"
	"github.com/blushft/strana/modules/sink/reporter/store"
	"github.com/blushft/strana/modules/sink/reporter/store/ent"
	"github.com/pkg/errors"
)

type EventTransaction interface {
	NewEvent(context.Context, *event.Event) error
}

type evtTx struct {
	s *store.Store
}

func NewEventTransaction(s *store.Store) EventTransaction {
	return &evtTx{
		s: s,
	}
}

func (etx *evtTx) NewEvent(ctx context.Context, evt *event.Event) (txerr error) {
	tx, err := etx.s.Client().Tx(ctx)
	if err != nil {
		return err
	}

	w := &txWorker{}

	if err := w.buildTx(ctx, evt); err != nil {
		return err
	}

	var werr error

	defer func() {
		if werr != nil {
			txerr = rollback(tx, werr)
		}
	}()

	for _, b := range w.before {
		if werr = b(tx); werr != nil {
			return werr
		}
	}

	if werr = w.exec(tx); werr != nil {
		return werr
	}

	for _, a := range w.after {
		if werr = a(tx); werr != nil {
			return werr
		}
	}

	txerr = tx.Commit()

	return
}

type workFn func(*ent.Tx) error

type txWorker struct {
	before []workFn
	exec   workFn
	after  []workFn
}

func (w *txWorker) buildTx(ctx context.Context, evt *event.Event) error {
	e, err := NewEvent(evt)
	if err != nil {
		return err
	}

	w.exec = func(tx *ent.Tx) error {
		_, err := eventEntityCreate(tx.Event, e).Save(ctx)
		return err
	}

	cs := evt.Context.Iter()

	for c := cs.First(); c != nil; c = cs.Next() {
		switch c.Type() {
		case contexts.ContextAction:
			this := c

			w.after = append(w.before, func(tx *ent.Tx) error {
				a := NewAction(this.Interface().(*contexts.Action))

				_, err := actionEntityCreate(tx.Action, a, ActionWithEventEdge(e.ID)).Save(ctx)
				return err
			})
		}
	}

	return nil
}

func rollback(tx *ent.Tx, txerr error) error {
	if err := tx.Rollback(); err != nil {
		return errors.Wrap(
			errors.Wrap(txerr, "transaction failed"), "rollback failed")
	}

	return txerr
}

package entity

import (
	"context"

	"github.com/blushft/strana/event"
	"github.com/blushft/strana/modules/sink/reporter/store"
	"github.com/blushft/strana/modules/sink/reporter/store/ent"
	"github.com/pkg/errors"
)

type EventTransaction interface {
	New(context.Context, *event.Event) error
}

type evtTx struct {
	s *store.Store
	w *txWorker
}

func NewEventTransaction(s *store.Store) EventTransaction {
	return &evtTx{
		s: s,
	}
}

func (etx *evtTx) New(ctx context.Context, evt *event.Event) (txerr error) {
	etx.w = &txWorker{
		evt: evt,
	}

	tx, err := etx.s.Client().Tx(ctx)
	if err != nil {
		return err
	}

	var werr error

	defer func() {
		if werr != nil {
			txerr = rollback(tx, werr)
		}
	}()

	for _, b := range etx.w.before {
		if werr = b(tx); werr != nil {
			return werr
		}
	}

	if werr = etx.w.exec(tx); werr != nil {
		return werr
	}

	for _, a := range etx.w.after {
		if werr = a(tx); werr != nil {
			return werr
		}
	}

	return nil
}

type workFn func(*ent.Tx) error

type txWorker struct {
	evt    *event.Event
	before []workFn
	exec   workFn
	after  []workFn
}

func (w *txWorker) buildTx() error {
	return nil
}

func rollback(tx *ent.Tx, txerr error) error {
	if err := tx.Rollback(); err != nil {
		return errors.Wrap(
			errors.Wrap(txerr, "transaction failed"), "rollback failed")
	}

	return txerr
}

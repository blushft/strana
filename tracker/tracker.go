package tracker

import (
	"log"
	"time"

	"github.com/blushft/strana/event"
	"gopkg.in/resty.v1"
)

type Tracker struct {
	options Options
	httpc   *resty.Client
	store   Store

	q  chan *event.Event
	cl chan struct{}

	copts []event.Option
}

func New(opts ...Option) (*Tracker, error) {
	options := defaultOptions(opts...)

	httpc := resty.New().SetHostURL(options.CollectorURL)
	store, err := NewMemStore()
	if err != nil {
		return nil, err
	}

	t := &Tracker{
		options: options,
		httpc:   httpc,
		store:   store,
		q:       make(chan *event.Event, options.QueueBuffer),
		cl:      make(chan struct{}),
		copts:   options.EventOptions(),
	}

	go t.collect()
	go t.emit()

	return t, nil
}

func (t *Tracker) SetOption(opt event.Option) *Tracker {
	t.copts = append(t.copts, opt)
	return t
}

func (t *Tracker) Track(typ event.Type, opts ...event.Option) error {
	evtOpts := t.eventOptions(opts...)
	evt := event.New(typ, evtOpts...)

	t.q <- evt

	return nil
}

func (t *Tracker) Action(a *event.Action, opts ...event.Option) error {
	opts = append(opts, event.WithActionContext(a))
	return t.Track(event.EventTypeAction, opts...)
}

func (t *Tracker) Identify(u *event.User, opts ...event.Option) error {
	opts = append(opts, event.WithUserContext(u))
	return t.Track(event.EventTypeIdentify, opts...)
}

func (t *Tracker) Alias(opts ...event.Option) error {
	return t.Track(event.EventTypeAlias, opts...)
}

func (t *Tracker) Page(opts ...event.Option) error {
	return t.Track(event.EventTypePageview, opts...)
}

func (t *Tracker) Screen(opts ...event.Option) error {
	return t.Track(event.EventTypeScreenview, opts...)
}

func (t *Tracker) Session(opts ...event.Option) error {
	return t.Track(event.EventTypeSession, opts...)
}

func (t *Tracker) Group(opts ...event.Option) error {
	return t.Track(event.EventTypeGroup, opts...)
}

func (t *Tracker) Transaction(opts ...event.Option) error {
	return t.Track(event.EventTypeTransaction, opts...)
}

func (t *Tracker) Timing(te *event.Timing, opts ...event.Option) error {
	opts = append(opts, event.WithTimingContext(te))
	return t.Track(event.EventTypeTiming, opts...)
}

func (t *Tracker) Close() error {
	close(t.cl)

	return nil
}

func (t *Tracker) eventOptions(opts ...event.Option) []event.Option {
	var eopts []event.Option

	eopts = append(eopts, t.copts...)
	eopts = append(eopts, opts...)

	return eopts
}

func (t *Tracker) collect() {
	for {
		select {
		case e := <-t.q:
			if err := t.store.Set(e); err != nil {
				log.Printf("error storing event: %s\n", err.Error())
			}
		case <-t.cl:
			return
		}
	}
}

func (t *Tracker) emit() {
	for {
		evts, err := t.store.GetAll()
		if err != nil {
			log.Printf("error getting items from store: %v", err)
			time.Sleep(time.Second)
		}

		if len(evts) == 0 {
			time.Sleep(time.Millisecond * 250)
			continue
		}

		for _, e := range evts {
			if err := t.send(e.Event); err != nil {
				e.Attempted = true
				e.Attempts++
				e.LastAttempt = time.Now()

				if err := t.store.Update(e); err != nil {
					log.Printf("error updating event: %v", err)
				}
			} else {
				if err := t.store.Remove(e); err != nil {
					log.Printf("error removing event: %v", err)
				}
			}
		}

		time.Sleep(time.Millisecond * 250)
	}
}

func (t *Tracker) send(e []byte) error {
	_, err := t.httpc.R().
		SetBody(e).
		SetHeader("Content-Type", "application/json").
		Post("/analytics/collect")

	return err
}

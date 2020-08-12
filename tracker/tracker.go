package tracker

import (
	"log"
	"time"

	"gopkg.in/resty.v1"
)

type Context interface {
	Values() map[string]interface{}
}

type Tracker struct {
	options Options
	httpc   *resty.Client
	store   Store

	q  chan *Event
	cl chan struct{}

	copts []EventOption
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
		q:       make(chan *Event, options.QueueBuffer),
		cl:      make(chan struct{}),
		copts:   options.EventOptions(),
	}

	go t.collect()
	go t.emit()

	return t, nil
}

func (t *Tracker) SetOption(opt EventOption) *Tracker {
	t.copts = append(t.copts, opt)
	return t
}

func (t *Tracker) Track(typ EventType, opts ...EventOption) error {
	evtOpts := t.eventOptions(opts...)
	evt := NewEvent(typ, evtOpts...)

	t.q <- evt

	return nil
}

func (t *Tracker) TrackAction(opts ...EventOption) error {
	return t.Track(EventTypeAction, opts...)
}

func (t *Tracker) TrackPageview(opts ...EventOption) error {
	return t.Track(EventTypePageview, opts...)
}

func (t *Tracker) Close() error {
	close(t.cl)

	return nil
}

func (t *Tracker) eventOptions(opts ...EventOption) []EventOption {
	var eopts []EventOption

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

package tracker

import (
	"math"
	"sync"
	"time"

	"github.com/blushft/strana/pkg/event"
)

type timer struct {
	lock  sync.Mutex
	start time.Time

	timing *event.Timing
}

func (t *Tracker) TimingStart(cat, label, variable string) *timer {
	return &timer{
		start: time.Now(),
		timing: &event.Timing{
			Category: cat,
			Label:    label,
			Variable: variable,
			Unit:     "seconds",
			Value:    -1,
		},
	}
}

func (timer *timer) End() *event.Timing {
	timer.lock.Lock()
	defer timer.lock.Unlock()
	dur := time.Now().Sub(timer.start).Seconds()

	timer.timing.Value = math.Floor(dur*100) / 100
	return timer.timing
}

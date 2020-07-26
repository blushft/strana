package entity

import (
	"fmt"
	"time"
)

type SiteStats struct {
	New            bool      `db:"-" json:"-" `
	SiteID         int64     `db:"site_id" json:"-"`
	Visitors       int64     `db:"visitors"`
	Pageviews      int64     `db:"pageviews"`
	Sessions       int64     `db:"sessions"`
	BounceRate     float64   `db:"bounce_rate"`
	AvgDuration    float64   `db:"avg_duration"`
	KnownDurations int64     `db:"known_durations" json:",omitempty"`
	Date           time.Time `db:"ts" json:",omitempty"`
}

func (s *SiteStats) FormattedDuration() string {
	return fmt.Sprintf("%d:%d", int(s.AvgDuration/60.00), (int(s.AvgDuration) % 60))
}

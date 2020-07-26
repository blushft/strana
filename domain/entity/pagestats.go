package entity

import (
	"time"
)

type PageStats struct {
	New            bool      `db:"-" json:"-"`
	SiteID         int64     `db:"site_id" json:"-"`
	HostnameID     int64     `db:"hostname_id" json:"-"`
	PathnameID     int64     `db:"pathname_id" json:"-"`
	Hostname       string    `db:"hostname"`
	Pathname       string    `db:"pathname"`
	Pageviews      int64     `db:"pageviews"`
	Visitors       int64     `db:"visitors"`
	Entries        int64     `db:"entries"`
	BounceRate     float64   `db:"bounce_rate"`
	AvgDuration    float64   `db:"avg_duration"`
	KnownDurations int64     `db:"known_durations"`
	Date           time.Time `db:"ts" json:",omitempty"`
}

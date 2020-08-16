package entity

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID         uuid.UUID
	TrackingID string
	UserID     string
	GroupID    string
	SessionID  string
	DeviceID   string

	Type      string
	Channel   string
	Platform  string
	Timestamp time.Time
}

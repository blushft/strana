package entity

import "github.com/google/uuid"

type Action struct {
	ID             uuid.UUID   `json:"id"`
	AppID          int         `json:"app_id"`
	SessionID      uuid.UUID   `json:"session_id"`
	UserID         string      `json:"user_id"`
	NonInteractive bool        `json:"non_interactive"`
	Category       string      `json:"category"`
	Action         string      `json:"action"`
	Label          string      `json:"label"`
	Property       string      `json:"property"`
	Value          interface{} `json:"value"`
}

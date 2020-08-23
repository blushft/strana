package entity

type Group struct {
	ID     int               `json:"id"`
	Name   string            `json:"name"`
	Meta   map[string]string `json:"meta"`
	Users  []*User           `json:"users"`
	Groups []*Group          `json:"groups"`
}

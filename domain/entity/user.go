package entity

type User struct {
	ID        string   `json:"id"`
	Anonymous bool     `json:"anonymous"`
	Name      string   `json:"name"`
	Groups    []*Group `json:"groups"`
}

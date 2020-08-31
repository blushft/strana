package entity

import (
	"github.com/facebook/ent"
	"github.com/google/uuid"
)

const (
	DefaultLimit  = 10
	DefaultOffset = 0
)

type QueryParams struct {
	DeviceIDs  []string          `json:"cid" query:"cid"`
	SessionIDs []uuid.UUID       `json:"sid" query:"sid"`
	UserIDs    []uuid.UUID       `json:"uid" query:"uid"`
	GroupIDs   []int             `json:"gid" query:"gid"`
	Offset     int               `json:"offset" query:"offset"`
	Limit      int               `json:"limit" query:"limit"`
	Start      int               `json:"start" query:"start"`
	End        int               `json:"end" query:"end"`
	Params     map[string]string `json:"params" query:"params"`
}

type Pageable interface {
	Limit(int) ent.Query
	Offset(int) ent.Query
}

func (qp QueryParams) SetLimit(v Pageable) {
	v.Limit(qp.Limit)
}

func (qp QueryParams) SetOffset(v Pageable) {
	v.Offset(qp.Offset)
}

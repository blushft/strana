package entity

import (
	"context"

	"github.com/blushft/strana/modules/sink/reporter/store/ent"
)

type Group struct {
	ID     int               `json:"id"`
	Name   string            `json:"name"`
	Meta   map[string]string `json:"meta"`
	Users  []*User           `json:"users"`
	Groups []*Group          `json:"groups"`
}

type GroupReader interface {
	List(QueryParams) ([]*Group, error)
	Get(id int) (*Group, error)
}

type GroupWriter interface {
	Create(*Group) error
	Update(*Group) error
	Delete(*Group) error
}

type groupRepo interface {
	GroupReader
	GroupWriter
}

type GroupManager interface {
	groupRepo
}

type groupManager struct {
	c *ent.GroupClient
}

func NewGroupService(c *ent.GroupClient) *groupManager {
	return &groupManager{
		c: c,
	}
}

func (mgr *groupManager) List(qp QueryParams) ([]*Group, error) {
	var res []*Group

	q := mgr.c.Query()

	recs, err := q.All(context.TODO())
	if err != nil {
		return nil, err
	}

	for _, rec := range recs {
		res = append(res, groupSchemaToEntity(rec))
	}

	return res, nil
}

func groupEntityCreate(c *ent.GroupClient, e *Group) *ent.GroupCreate {
	return c.Create().
		SetName(e.Name)
}

func groupEntityUpdate(c *ent.GroupClient, e *Group) *ent.GroupUpdate {
	return c.Update().
		SetName(e.Name)
}

func groupSchemaToEntity(sch *ent.Group) *Group {
	return &Group{
		ID:   sch.ID,
		Name: sch.Name,
	}
}

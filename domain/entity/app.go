package entity

import (
	"context"

	"github.com/blushft/strana/platform/store"
	"github.com/blushft/strana/platform/store/ent"
)

type App struct {
	ID         int    `db:"id" json:"id"`
	TrackingID string `db:"tracking_id" json:"trackingId"`
	Name       string `db:"name" json:"name"`
}

type AppReader interface {
	List() ([]*App, error)
	Get(int) (*App, error)
}

type AppWriter interface {
	Create(*App) error
	Update(*App) error
	Delete(*App) error
}

type appRepo interface {
	AppReader
	AppWriter
}

type AppManager interface {
	appRepo
}

type AppReporter interface {
	AppReader
}

type appManager struct {
	store *store.Store
}

func NewAppService(s *store.Store) *appManager {
	return &appManager{
		store: s,
	}
}

func (mgr *appManager) List() ([]*App, error) {
	c := mgr.store.Client().App
	var res []*App

	recs, err := c.Query().All(context.TODO())
	if err != nil {
		return nil, err
	}

	for _, rec := range recs {
		res = append(res, siteSchemaToEntity(rec))
	}

	return res, nil
}

func (mgr *appManager) Get(id int) (*App, error) {
	c := mgr.store.Client().App

	rec, err := c.Get(context.TODO(), id)
	if err != nil {
		return nil, err
	}

	return siteSchemaToEntity(rec), nil
}

func siteSchemaToEntity(sch *ent.App) *App {
	return &App{
		ID:         sch.ID,
		Name:       sch.Name,
		TrackingID: sch.TrackingID,
	}
}

func siteEntityCreate(c *ent.AppClient, e *App) *ent.AppCreate {
	return c.Create().
		SetName(e.Name).
		SetTrackingID(e.TrackingID)
}

func siteEntityUpdate(c *ent.AppClient, e *App) *ent.AppUpdate {
	return c.Update().
		SetName(e.Name).
		SetTrackingID(e.TrackingID)
}

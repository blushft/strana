package entity

import (
	"context"

	"github.com/blushft/strana/platform/store/reporter"
	"github.com/blushft/strana/platform/store/reporter/ent"
	"github.com/blushft/strana/platform/store/reporter/ent/app"
)

type App struct {
	ID         int    `db:"id" json:"id"`
	Name       string `db:"name" json:"name"`
	Version    string
	Build      string
	Properties map[string]interface{}
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
	store *reporter.Store
}

func NewAppService(s *reporter.Store) *appManager {
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

func (mgr *appManager) GetByTrackingID(tid string) (*App, error) {
	c := mgr.store.Client().App

	rec, err := c.Query().Where(app.TrackingIDEqualFold(tid)).Only(context.TODO())
	if err != nil {
		return nil, err
	}

	return siteSchemaToEntity(rec), nil
}

func siteSchemaToEntity(sch *ent.App) *App {
	return &App{
		ID:   sch.ID,
		Name: sch.Name,
	}
}

func siteEntityCreate(c *ent.AppClient, e *App) *ent.AppCreate {
	return c.Create().
		SetName(e.Name)

}

func siteEntityUpdate(c *ent.AppClient, e *App) *ent.AppUpdate {
	return c.Update().
		SetName(e.Name)
}

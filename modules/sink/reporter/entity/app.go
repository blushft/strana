package entity

import (
	"context"

	"github.com/blushft/strana/modules/sink/reporter/store"
	"github.com/blushft/strana/modules/sink/reporter/store/ent"
	"github.com/blushft/strana/modules/sink/reporter/store/ent/app"
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
	store *store.Store
}

func NewAppService(s *store.Store) *appManager {
	return &appManager{
		store: s,
	}
}

func (mgr *appManager) List() ([]*App, error) {
	c := mgr.store.Client().App

	recs, err := c.Query().All(context.TODO())
	if err != nil {
		return nil, err
	}

	return siteSchemasToEntities(recs), nil
}

func (mgr *appManager) Get(id int) (*App, error) {
	c := mgr.store.Client().App

	rec, err := c.Get(context.TODO(), id)
	if err != nil {
		return nil, err
	}

	return siteSchemaToEntity(rec), nil
}

func (mgr *appManager) GetByName(name string) ([]*App, error) {
	c := mgr.store.Client().App

	rec, err := c.Query().Where(app.Name(name)).All(context.TODO())
	if err != nil {
		return nil, err
	}

	return siteSchemasToEntities(rec), nil
}

func siteSchemasToEntities(sch []*ent.App) []*App {
	var res []*App
	for _, rec := range sch {
		res = append(res, siteSchemaToEntity(rec))
	}

	return res
}

func siteSchemaToEntity(sch *ent.App) *App {
	return &App{
		ID:         sch.ID,
		Name:       sch.Name,
		Version:    sch.Version,
		Build:      sch.Build,
		Properties: sch.Properties,
	}
}

func siteEntityCreate(c *ent.AppClient, e *App) *ent.AppCreate {
	return c.Create().
		SetName(e.Name).
		SetVersion(e.Version).
		SetBuild(e.Build).
		SetProperties(e.Properties)
}

func siteEntityUpdate(c *ent.AppClient, e *App) *ent.AppUpdate {
	return c.Update().
		SetName(e.Name).
		SetVersion(e.Version).
		SetBuild(e.Build).
		SetProperties(e.Properties)
}

package entity

import (
	"context"
	"time"

	"github.com/blushft/strana/modules/sink/reporter/store"
	"github.com/blushft/strana/modules/sink/reporter/store/ent"
	"github.com/blushft/strana/modules/sink/reporter/store/ent/app"
	"github.com/blushft/strana/modules/sink/reporter/store/ent/pageview"
	"github.com/blushft/strana/modules/sink/reporter/store/ent/predicate"

	"github.com/google/uuid"
)

type Pageview struct {
	ID         uuid.UUID `json:"id"`
	AppID      int       `json:"app_id"`
	SessionID  uuid.UUID `json:"session_id"`
	Hostname   string    `json:"hostname"`
	Pathname   string    `json:"pathname"`
	IsEntry    bool
	IsFinished bool                   `json:"is_finished"`
	Referrer   string                 `json:"referrer"`
	Duration   int                    `json:"duration"`
	Timestamp  time.Time              `json:"timestamp"`
	UserAgent  string                 `json:"user_agent"`
	IPAddress  string                 `json:"ip_address"`
	ScreenDim  string                 `json:"screen_dim"`
	Extra      map[string]interface{} `json:"extra"`
}

type PageviewReader interface {
	Get(uuid.UUID) (*Pageview, error)
	List(limit int, offset int, sites ...int) ([]*Pageview, error)
	GetProcessable(limit int) ([]*Pageview, error)
}

type PageviewWriter interface {
	Create(*Pageview) error
	InsertBatch([]*Pageview) error
	UpdateBatch([]*Pageview) error
	DeleteBatch([]*Pageview) error
}

type pageviewRepo interface {
	PageviewReader
	PageviewWriter
}

type PageviewManager interface {
	pageviewRepo
}

type PageviewReporter interface {
	PageviewReader
}

type pvManager struct {
	store *store.Store
}

func NewPageviewService(s *store.Store) *pvManager {
	return &pvManager{
		store: s,
	}
}

func (mgr *pvManager) Get(id uuid.UUID) (*Pageview, error) {
	c := mgr.store.Client().PageView

	rec, err := c.Query().Where(pageview.IDEQ(id)).
		WithApp().
		Only(context.TODO())
	if err != nil {
		return nil, err
	}

	return pvSchemaToEntity(rec), nil
}

func (mgr *pvManager) List(limit, offset int, sites ...int) ([]*Pageview, error) {
	c := mgr.store.Client().PageView
	var res []*Pageview
	var q predicate.PageView

	if len(sites) > 0 {
		var siteq []predicate.PageView
		for _, sid := range sites {
			siteq = append(siteq, pageview.HasAppWith(app.ID(sid)))
		}

		q = pageview.Or(siteq...)
	}

	recs, err := c.Query().Where(q).WithApp().Limit(limit).Offset(offset).All(context.TODO())
	if err != nil {
		return nil, err
	}

	for _, rec := range recs {
		res = append(res, pvSchemaToEntity(rec))
	}

	return res, nil
}

func (mgr *pvManager) GetProcessable(limit int) ([]*Pageview, error) {
	c := mgr.store.Client().PageView
	var res []*Pageview

	prevHH := time.Now().Add(-30 * time.Minute)

	qry := pageview.Or(
		pageview.IsFinished(true),
		pageview.TimestampLT(prevHH),
	)

	recs, err := c.Query().Where(qry).WithApp().Limit(limit).All(context.TODO())
	if err != nil {
		return nil, err
	}

	for _, rec := range recs {
		res = append(res, pvSchemaToEntity(rec))
	}

	return res, nil
}

func (mgr *pvManager) Create(pv *Pageview) error {
	c := mgr.store.Client().PageView

	_, err := pvEntityCreate(c, pv).Save(context.TODO())
	if err != nil {
		return err
	}

	return nil
}

func (mgr *pvManager) Update(pv *Pageview) error {
	c := mgr.store.Client().PageView

	_, err := pvEntityUpdate(c, pv).Save(context.TODO())
	if err != nil {
		return err
	}

	return nil
}

func (mgr *pvManager) Delete(pv *Pageview) error {
	c := mgr.store.Client().PageView

	return c.DeleteOneID(pv.ID).Exec(context.TODO())
}

func (mgr *pvManager) InsertBatch(pvs []*Pageview) error {
	//TODO: create a tx

	for _, pv := range pvs {
		if err := mgr.Create(pv); err != nil {
			return err
		}
	}

	return nil
}

func (mgr *pvManager) UpdateBatch(pvs []*Pageview) error {
	//TODO: create a tx

	for _, pv := range pvs {
		if err := mgr.Update(pv); err != nil {
			return err
		}
	}

	return nil
}

func (mgr *pvManager) DeleteBatch(pvs []*Pageview) error {
	for _, pv := range pvs {
		if err := mgr.Delete(pv); err != nil {
			return err
		}
	}

	return nil
}

func pvSchemaToEntity(sch *ent.PageView) *Pageview {
	return &Pageview{
		ID:         sch.ID,
		AppID:      sch.Edges.App.ID,
		SessionID:  sch.Edges.Session.ID,
		Hostname:   sch.Hostname,
		Pathname:   sch.Pathname,
		IsEntry:    sch.IsEntry,
		IsFinished: sch.IsFinished,
		Duration:   sch.Duration,
		Timestamp:  sch.Timestamp,
		UserAgent:  sch.UserAgent,
		IPAddress:  sch.IPAddress,
		ScreenDim:  sch.ScreenDim,
		Extra:      sch.Extra,
	}
}

func pvEntityCreate(c *ent.PageViewClient, e *Pageview) *ent.PageViewCreate {
	return c.Create().
		SetID(e.ID).
		SetSessionID(e.SessionID).
		SetAppID(e.AppID).
		SetHostname(e.Hostname).
		SetPathname(e.Pathname).
		SetReferrer(e.Referrer).
		SetIsEntry(e.IsEntry).
		SetIsFinished(e.IsFinished).
		SetDuration(e.Duration).
		SetTimestamp(e.Timestamp).
		SetUserAgent(e.UserAgent).
		SetIPAddress(e.IPAddress).
		SetScreenDim(e.ScreenDim).
		SetExtra(e.Extra)
}

func pvEntityUpdate(c *ent.PageViewClient, e *Pageview) *ent.PageViewUpdate {
	return c.Update().
		SetSessionID(e.SessionID).
		SetAppID(e.AppID).
		SetHostname(e.Hostname).
		SetPathname(e.Pathname).
		SetReferrer(e.Referrer).
		SetIsEntry(e.IsEntry).
		SetIsFinished(e.IsFinished).
		SetDuration(e.Duration).
		SetTimestamp(e.Timestamp).
		SetUserAgent(e.UserAgent).
		SetIPAddress(e.IPAddress).
		SetScreenDim(e.ScreenDim).
		SetExtra(e.Extra).
		Where(pageview.ID(e.ID))
}

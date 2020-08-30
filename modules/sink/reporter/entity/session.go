package entity

import (
	"context"
	"time"

	"github.com/blushft/strana/modules/sink/reporter/store"
	"github.com/blushft/strana/modules/sink/reporter/store/ent"
	"github.com/blushft/strana/modules/sink/reporter/store/ent/session"
	"github.com/blushft/strana/platform/cache"

	"github.com/google/uuid"
)

type Session struct {
	ID         uuid.UUID
	UserID     string
	DeviceID   *string
	AppID      int
	NewUser    bool
	IsUnique   bool
	IsBounce   bool
	IsFinished bool
	Duration   int
	StartedAt  time.Time
	FinishedAt *time.Time
}

type SessionReader interface {
	List(QueryParams) ([]*Session, error)
	Get(uuid.UUID) (*Session, error)
}

type SessionWriter interface {
	Create(*Session) error
	Update(*Session) error
	Delete(*Session) error
}

type sessionRepo interface {
	SessionReader
	SessionWriter
}

type SessionReporter interface {
	SessionReader
}

type SessionManager interface {
	sessionRepo
}

type sessionManager struct {
	store *store.Store
}

func NewSessionService(s *store.Store) *sessionManager {
	return &sessionManager{
		store: s,
	}
}

func (mgr *sessionManager) List(qp QueryParams) ([]*Session, error) {
	c := mgr.store.Client().Session
	var res []*Session

	q := c.Query()

	if qp.Limit > 0 {
		q.Limit(qp.Limit)

		if qp.Offset > 0 {
			q.Offset(qp.Offset)
		}
	}

	recs, err := q.All(context.TODO())
	if err != nil {
		return nil, err
	}

	for _, rec := range recs {
		res = append(res, sessionSchemaToEntity(rec))
	}

	return res, nil
}

func (mgr *sessionManager) Get(id uuid.UUID) (*Session, error) {
	c := mgr.store.Client().Session

	rec, err := c.Query().Where(session.ID(id)).Only(context.TODO())
	if err != nil {
		return nil, err
	}

	return sessionSchemaToEntity(rec), nil
}

func (mgr *sessionManager) Create(s *Session) error {
	c := mgr.store.Client().Session

	_, err := sessionEntityCreate(c, s).Save(context.TODO())

	return err
}

func (mgr *sessionManager) Update(s *Session) error {
	c := mgr.store.Client().Session

	_, err := sessionEntityUpdate(c, s).Save(context.TODO())

	return err
}

func (mgr *sessionManager) Delete(s *Session) error {
	c := mgr.store.Client().Session
	return c.DeleteOneID(s.ID).Exec(context.TODO())
}

type cachedSessionManager struct {
	mgr   *sessionManager
	cache *cache.Cache
}

func NewCachedSessionService(s *store.Store, c *cache.Cache) *cachedSessionManager {
	return &cachedSessionManager{
		mgr:   NewSessionService(s),
		cache: c,
	}
}

func (mgr *cachedSessionManager) List(qp QueryParams) ([]*Session, error) {
	return mgr.mgr.List(qp)
}

func (mgr *cachedSessionManager) Get(id uuid.UUID) (*Session, error) {
	cv, err := mgr.cache.Get(id)
	if err == nil {
		ses, ok := cv.(*Session)
		if ok {
			return ses, nil
		}
	}

	return mgr.mgr.Get(id)
}

func (mgr *cachedSessionManager) Create(s *Session) error {
	if err := mgr.mgr.Create(s); err != nil {
		return err
	}

	return mgr.cache.Set(s.ID, s, nil)
}

func (mgr *cachedSessionManager) Update(s *Session) error {
	if err := mgr.mgr.Update(s); err != nil {
		return err
	}

	return mgr.cache.Set(s.ID, s, nil)
}

func (mgr *cachedSessionManager) Delete(s *Session) error {
	if err := mgr.mgr.Delete(s); err != nil {
		return err
	}

	return mgr.cache.Delete(s.ID)
}

func sessionEntityCreate(c *ent.SessionClient, e *Session) *ent.SessionCreate {
	cr := c.Create().
		SetID(e.ID).
		SetNewUser(e.NewUser).
		SetIsUnique(e.IsUnique).
		SetIsBounce(e.IsBounce).
		SetIsFinished(e.IsFinished).
		SetStartedAt(e.StartedAt).
		SetNillableFinishedAt(e.FinishedAt)

	return cr
}

func sessionEntityUpdate(c *ent.SessionClient, e *Session) *ent.SessionUpdate {
	cr := c.Update().
		SetNewUser(e.NewUser).
		SetIsUnique(e.IsUnique).
		SetIsBounce(e.IsBounce).
		SetIsFinished(e.IsFinished).
		SetStartedAt(e.StartedAt).
		SetNillableFinishedAt(e.FinishedAt)

	return cr.Where(session.ID(e.ID))
}

func sessionSchemaToEntity(sch *ent.Session) *Session {
	return &Session{
		ID:         sch.ID,
		NewUser:    sch.NewUser,
		IsUnique:   sch.IsUnique,
		IsBounce:   sch.IsBounce,
		IsFinished: sch.IsFinished,
		StartedAt:  sch.StartedAt,
		FinishedAt: sch.FinishedAt,
	}
}

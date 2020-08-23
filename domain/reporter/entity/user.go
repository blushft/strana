package entity

import (
	"context"

	"github.com/blushft/strana/platform/store/reporter"
	"github.com/blushft/strana/platform/store/reporter/ent"
	"github.com/blushft/strana/platform/store/reporter/ent/predicate"
	"github.com/blushft/strana/platform/store/reporter/ent/session"
	"github.com/blushft/strana/platform/store/reporter/ent/user"
	"github.com/google/uuid"
)

type User struct {
	ID        string     `json:"id"`
	Anonymous bool       `json:"anonymous"`
	Name      string     `json:"name"`
	Sessions  []*Session `json:"sessions"`
	Groups    []*Group   `json:"groups"`
}

type UserReader interface {
	List(QueryParams) ([]*User, error)
	Get(id string) (*User, error)
}

type UserWriter interface {
	Create(*User) error
	Update(*User) error
	Delete(*User) error
}

type userRepo interface {
	UserReader
	UserWriter
}

type UserReporter interface {
	UserReader
}

type UserManager interface {
	userRepo
}

type userManager struct {
	store *reporter.Store
}

func NewUserService(s *reporter.Store) *userManager {
	return &userManager{
		store: s,
	}
}

func (mgr *userManager) List(qp QueryParams) ([]*User, error) {
	c := mgr.store.Client().User
	var res []*User

	q := c.Query()

	if len(qp.SessionIDs) > 0 {
		sesq := make([]predicate.User, 0, len(qp.SessionIDs))

		for _, id := range qp.SessionIDs {
			sid, err := uuid.Parse(id)
			if err != nil {
				continue
			}

			sesq = append(sesq, user.HasSessionsWith(session.ID(sid)))
		}

		q.Where(user.Or(sesq...))
	}

	if qp.Limit > 0 {
		q.Limit(qp.Limit)

		if qp.Offset > 0 {
			q.Offset(qp.Offset)
		}
	}

	recs, err := q.WithSessions().All(context.TODO())
	if err != nil {
		return nil, err
	}

	for _, rec := range recs {
		res = append(res, userSchemaToEntity(rec))
	}

	return res, nil
}

func (mgr *userManager) Get(id string) (*User, error) {
	c := mgr.store.Client().User

	u, err := c.Get(context.TODO(), id)
	if err != nil {
		return nil, err
	}

	return userSchemaToEntity(u), nil
}

func (mgr *userManager) Create(u *User) error {
	c := mgr.store.Client().User
	_, err := userEntityCreate(c, u).Save(context.TODO())

	return err
}

func (mgr *userManager) Update(u *User) error {
	c := mgr.store.Client().User
	_, err := userEntityUpdate(c, u).Save(context.TODO())

	return err
}

func (mgr *userManager) Delete(u *User) error {
	c := mgr.store.Client().User
	return c.DeleteOneID(u.ID).Exec(context.TODO())
}

func userEntityCreate(c *ent.UserClient, e *User) *ent.UserCreate {
	return c.Create().
		SetID(e.ID).
		SetIsAnonymous(e.Anonymous).
		SetName(e.Name)
}

func userEntityUpdate(c *ent.UserClient, e *User) *ent.UserUpdate {
	return c.Update().
		SetIsAnonymous(e.Anonymous).
		SetName(e.Name).
		Where(user.ID(e.ID))
}

func userSchemaToEntity(sch *ent.User) *User {
	ses := make([]*Session, 0, len(sch.Edges.Sessions))
	for _, s := range sch.Edges.Sessions {
		ses = append(ses, sessionSchemaToEntity(s))
	}

	return &User{
		ID:        sch.ID,
		Anonymous: sch.IsAnonymous,
		Name:      sch.Name,
		Sessions:  ses,
	}
}

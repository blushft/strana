package entity

import (
	"context"
	"time"

	"github.com/blushft/strana/modules/sink/reporter/store"
	"github.com/blushft/strana/modules/sink/reporter/store/ent"
	"github.com/blushft/strana/modules/sink/reporter/store/ent/user"
)

type User struct {
	ID        string                 `json:"id"`
	Anonymous bool                   `json:"anonymous"`
	Name      string                 `json:"name"`
	Title     string                 `json:"title"`
	FirstName string                 `json:"first_name"`
	LastName  string                 `json:"last_name"`
	Email     string                 `json:"email"`
	Username  string                 `json:"username"`
	Age       int                    `json:"age"`
	Birthday  time.Time              `json:"birthday"`
	Gender    string                 `json:"gender"`
	Phone     string                 `json:"phone"`
	Website   string                 `json:"website"`
	Extra     map[string]interface{} `json:"extra"`
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
	store *store.Store
}

func NewUserService(s *store.Store) *userManager {
	return &userManager{
		store: s,
	}
}

func (mgr *userManager) List(qp QueryParams) ([]*User, error) {
	c := mgr.store.Client().User
	var res []*User

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
	return &User{
		ID:        sch.ID,
		Anonymous: sch.IsAnonymous,
		Name:      sch.Name,
	}
}

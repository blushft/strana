package tracker

import (
	"testing"

	"github.com/blushft/strana/pkg/event"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/suite"
)

type StoreSuite struct {
	suite.Suite
	store Store
}

func TestStoreSuite(t *testing.T) {
	suite.Run(t, new(StoreSuite))
}

func (s *StoreSuite) SetupSuite() {
	store, err := NewMemStore()
	if err != nil {
		s.FailNow(err.Error())
	}

	s.store = store
}

func (s *StoreSuite) TestASet() {
	evt := event.New(event.EventTypeAction,
		event.TrackingID("123"),
		event.SessionID("321"),
		event.UserID("someguy"),
		event.WithActionContext(&event.Action{
			Category: "test",
		}),
	)

	spew.Dump(evt)

	s.Require().NoError(s.store.Set(evt))

	se, err := s.store.Get(evt.ID)
	s.Require().NoError(err)

	spew.Dump(se)
}

func (s *StoreSuite) TestBGetAll() {
	se, err := s.store.GetAll()
	s.Require().NoError(err)

	s.Equal(1, len(se))
}

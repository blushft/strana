package tracker

import (
	"testing"

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
	evt := NewEvent(EventTypeAction,
		TrackingID("123"),
		SessionID("321"),
		UserID("someguy"),
		WithActionContext(&Action{
			Category: "test",
		}),
	)

	s.Require().NoError(s.store.Set(evt))
}

func (s *StoreSuite) TestBGetAll() {
	se, err := s.store.GetAll()
	s.Require().NoError(err)

	s.Equal(1, len(se))
}

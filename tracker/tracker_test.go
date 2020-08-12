package tracker

import (
	"testing"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type TrackerSuite struct {
	suite.Suite
	tr *Tracker
}

func TestRunTrackerSuite(t *testing.T) {
	suite.Run(t, new(TrackerSuite))
}

func (s *TrackerSuite) SetupSuite() {
	tr, err := New(WithTrackingID("1234"))
	s.Require().NoError(err)
	s.tr = tr
}

func (s *TrackerSuite) TearDownSuite() {
	s.tr.Close()
}

func (s *TrackerSuite) TestTrackAction() {
	if err := s.tr.TrackAction(
		SessionID(uuid.New().String()),
		UserID("me@nomail.int"),
		AppVersion("0.0.1"),
		WithActionContext(&Action{
			Category: "test",
			Action:   "testing",
			Label:    "test",
			Property: "prop",
			Value:    "val",
		}),
	); err != nil {
		s.Fail(err.Error())
	}

	time.Sleep(time.Second * 5)

	v, err := s.tr.store.GetAll()
	if err != nil {
		s.Fail(err.Error())
	}

	spew.Dump(v)
}

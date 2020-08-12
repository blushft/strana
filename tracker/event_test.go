package tracker

import (
	"encoding/json"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type EventSuite struct {
	suite.Suite
}

func TestRunEventSuite(t *testing.T) {
	suite.Run(t, new(EventSuite))
}

func (s *EventSuite) TestAction() {
	evt := NewEvent("action",
		SessionID(uuid.New().String()),
		UserID("me@nomail.int"),
		AppVersion("0.0.1"),
		WithActionContext(&Action{
			Category: "test",
			Action:   "testing",
			Label:    "test",
			Property: "prop",
			Value:    "val",
		}))

	b, err := evt.Payload()
	s.Require().NoError(err)

	spew.Dump(string(b))

	var jevt *Event
	s.Require().NoError(json.Unmarshal(b, &jevt))

	s.Equal(evt.ID, jevt.ID)
}

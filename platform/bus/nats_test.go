package bus

import (
	"testing"
	"time"

	"github.com/blushft/strana/pkg/event"
	"github.com/stretchr/testify/suite"
)

type NatsSuite struct {
	suite.Suite
	b *natsBus

	sub      *subscriber
	subcount int
}

func TestRunNatsSuite(t *testing.T) {
	suite.Run(t, new(NatsSuite))
}

func (s *NatsSuite) SetupSuite() {
	b, err := newNatsBus(4222, 4224, "supersecrettokenthing")
	if err != nil {
		s.FailNow("unable to setup bus", err)
	}

	s.b = b

	go s.b.Start()

	s.True(s.b.svr.ReadyForConnections(time.Second * 10))
}

func (s *NatsSuite) TeardownSuite() {
	s.NoError(s.sub.Close())

	s.b.Shutdown()
}

func (s *NatsSuite) TestASub() {
	sub, err := s.b.NewSubscriber("test.events.*", s.handle)
	if err != nil {
		s.FailNow("unable to create subscriber", err)
	}

	s.sub = sub
}

func (s *NatsSuite) TestBPub() {
	evt := event.New(event.EventTypeAction,
		event.WithActionContext(
			&event.Action{
				Category: "tests",
				Action:   "tested",
				Label:    "testing",
				Property: "test-val",
				Value:    22,
			},
		),
	)

	s.Require().NoError(s.b.Publish("test.events.action", evt))

	time.Sleep(time.Second * 3)

	s.Equal(1, s.subcount)
}

func (s *NatsSuite) handle(evt *event.Event) error {
	s.subcount++

	return nil
}

package event

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/suite"
)

type EventSuite struct {
	suite.Suite
}

func TestRunEventSuite(t *testing.T) {
	suite.Run(t, new(EventSuite))
}

func (s *EventSuite) TestMarshal() {
	action := &context{
		typ: ContextAction,
		v: Action{
			Category: "tests",
			Action:   "testing",
			Label:    "test_marshal",
			Property: "test_value",
			Value:    "abcd",
		},
	}

	app := &context{
		typ: ContextApp,
		v: App{
			Name:    "testapp",
			Version: "0.0.1",
			Build:   "-",
			Properties: map[string]interface{}{
				"commit": "abcdefg",
			},
		},
	}

	evt := New(
		EventTypeAction,
		TrackingID("abcd"),
		WithContext(action),
		WithContext(app),
		WithContext(NewNetworkContext("127.0.0.1", "")),
	)

	b, err := json.MarshalIndent(evt, "  ", "  ")
	s.Require().NoError(err, "marshal event")

	fmt.Println(string(b))
}

func (s *EventSuite) TestUnmarshal() {
	var evt *Event
	s.Require().NoError(json.Unmarshal([]byte(evtFixture), &evt))

	s.Equal(EventTypeAction, evt.Event)
	spew.Dump(evt)
}

var evtFixture = `
{
	"id": "c21ab6dc-9ed8-4e83-9a7d-8092782b90e0",
	"trackingId": "abcd",
	"event": "action",
	"timestamp": "2020-08-15T12:18:08.787165623-04:00",
	"context": {
		"action": {
			"action": "testing",
			"category": "tests",
			"label": "test_marshal",
			"property": "test_value",
			"value": "abcd"
		},
		"app": {
			"build": "-",
			"name": "testapp",
			"properties": {
				"commit": "abcdefg"
			},
			"version": "0.0.1"
		},
		"network": {
			"ip": "127.0.0.1"
		}
	}
}
`

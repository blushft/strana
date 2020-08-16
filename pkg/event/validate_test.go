package event

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ValidateSuite struct {
	suite.Suite
}

func TestRunValidateSuite(t *testing.T) {
	suite.Run(t, new(ValidateSuite))
}

func (s *ValidateSuite) TestHasContext() {
	ctx := &context{
		typ: ContextAction,
		v: Action{
			Category: "test",
			Action:   "testing",
			Label:    "validate_suite",
			Property: "test_value",
			Value:    3,
		},
	}

	evt := New(
		EventTypeAction,
		TrackingID("testing"),
		WithContext(ctx),
	)

	validator := NewValidator(
		HasContext(ContextAction),
		ContextContains(ContextAction, "value", true),
	)

	s.True(validator.Validate(evt))
}

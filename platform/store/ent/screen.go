// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/blushft/strana/platform/store/ent/screen"
	"github.com/facebookincubator/ent/dialect/sql"
)

// Screen is the model entity for the Screen schema.
type Screen struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Screen) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // id
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Screen fields.
func (s *Screen) assignValues(values ...interface{}) error {
	if m, n := len(values), len(screen.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	s.ID = int(value.Int64)
	values = values[1:]
	return nil
}

// Update returns a builder for updating this Screen.
// Note that, you need to call Screen.Unwrap() before calling this method, if this Screen
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Screen) Update() *ScreenUpdateOne {
	return (&ScreenClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (s *Screen) Unwrap() *Screen {
	tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Screen is not a transactional entity")
	}
	s.config.driver = tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Screen) String() string {
	var builder strings.Builder
	builder.WriteString("Screen(")
	builder.WriteString(fmt.Sprintf("id=%v", s.ID))
	builder.WriteByte(')')
	return builder.String()
}

// Screens is a parsable slice of Screen.
type Screens []*Screen

func (s Screens) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}

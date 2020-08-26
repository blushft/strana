// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/blushft/strana/modules/sink/reporter/store/ent/pathname"
	"github.com/facebook/ent/dialect/sql"
)

// Pathname is the model entity for the Pathname schema.
type Pathname struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PathnameQuery when eager-loading is set.
	Edges PathnameEdges `json:"edges"`
}

// PathnameEdges holds the relations/edges for other nodes in the graph.
type PathnameEdges struct {
	// PageStats holds the value of the page_stats edge.
	PageStats []*PageStat
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// PageStatsOrErr returns the PageStats value or an error if the edge
// was not loaded in eager-loading.
func (e PathnameEdges) PageStatsOrErr() ([]*PageStat, error) {
	if e.loadedTypes[0] {
		return e.PageStats, nil
	}
	return nil, &NotLoadedError{edge: "page_stats"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Pathname) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // name
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Pathname fields.
func (pa *Pathname) assignValues(values ...interface{}) error {
	if m, n := len(values), len(pathname.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	pa.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[0])
	} else if value.Valid {
		pa.Name = value.String
	}
	return nil
}

// QueryPageStats queries the page_stats edge of the Pathname.
func (pa *Pathname) QueryPageStats() *PageStatQuery {
	return (&PathnameClient{config: pa.config}).QueryPageStats(pa)
}

// Update returns a builder for updating this Pathname.
// Note that, you need to call Pathname.Unwrap() before calling this method, if this Pathname
// was returned from a transaction, and the transaction was committed or rolled back.
func (pa *Pathname) Update() *PathnameUpdateOne {
	return (&PathnameClient{config: pa.config}).UpdateOne(pa)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (pa *Pathname) Unwrap() *Pathname {
	tx, ok := pa.config.driver.(*txDriver)
	if !ok {
		panic("ent: Pathname is not a transactional entity")
	}
	pa.config.driver = tx.drv
	return pa
}

// String implements the fmt.Stringer.
func (pa *Pathname) String() string {
	var builder strings.Builder
	builder.WriteString("Pathname(")
	builder.WriteString(fmt.Sprintf("id=%v", pa.ID))
	builder.WriteString(", name=")
	builder.WriteString(pa.Name)
	builder.WriteByte(')')
	return builder.String()
}

// Pathnames is a parsable slice of Pathname.
type Pathnames []*Pathname

func (pa Pathnames) config(cfg config) {
	for _i := range pa {
		pa[_i].config = cfg
	}
}
// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebook/ent/dialect/sql/schema"
	"github.com/facebook/ent/schema/field"
)

var (
	// RawEventsColumns holds the columns for the "raw_events" table.
	RawEventsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "tracking_id", Type: field.TypeUUID},
		{Name: "user_id", Type: field.TypeString},
		{Name: "anonymous", Type: field.TypeBool},
		{Name: "group_id", Type: field.TypeString, Nullable: true},
		{Name: "session_id", Type: field.TypeString, Nullable: true},
		{Name: "device_id", Type: field.TypeString, Nullable: true},
		{Name: "event", Type: field.TypeString},
		{Name: "non_interactive", Type: field.TypeBool},
		{Name: "channel", Type: field.TypeString, Nullable: true},
		{Name: "platform", Type: field.TypeString, Nullable: true},
		{Name: "timestamp", Type: field.TypeTime},
		{Name: "context", Type: field.TypeJSON},
	}
	// RawEventsTable holds the schema information for the "raw_events" table.
	RawEventsTable = &schema.Table{
		Name:        "raw_events",
		Columns:     RawEventsColumns,
		PrimaryKey:  []*schema.Column{RawEventsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
		Indexes: []*schema.Index{
			{
				Name:    "rawevent_user_id_tracking_id_group_id_session_id_device_id_event",
				Unique:  false,
				Columns: []*schema.Column{RawEventsColumns[2], RawEventsColumns[1], RawEventsColumns[4], RawEventsColumns[5], RawEventsColumns[6], RawEventsColumns[7]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		RawEventsTable,
	}
)

func init() {
}

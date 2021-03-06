// Code generated by entc, DO NOT EDIT.

package oscontext

const (
	// Label holds the string label denoting the oscontext type in the database.
	Label = "os_context"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldFamily holds the string denoting the family field in the database.
	FieldFamily = "family"
	// FieldPlatform holds the string denoting the platform field in the database.
	FieldPlatform = "platform"
	// FieldVersion holds the string denoting the version field in the database.
	FieldVersion = "version"

	// EdgeEvents holds the string denoting the events edge name in mutations.
	EdgeEvents = "events"

	// Table holds the table name of the oscontext in the database.
	Table = "os"
	// EventsTable is the table the holds the events relation/edge.
	EventsTable = "events"
	// EventsInverseTable is the table name for the Event entity.
	// It exists in this package in order to avoid circular dependency with the "event" package.
	EventsInverseTable = "events"
	// EventsColumn is the table column denoting the events relation/edge.
	EventsColumn = "event_os"
)

// Columns holds all SQL columns for oscontext fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldFamily,
	FieldPlatform,
	FieldVersion,
}

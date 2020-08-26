// Code generated by entc, DO NOT EDIT.

package user

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldIsAnonymous holds the string denoting the is_anonymous field in the database.
	FieldIsAnonymous = "is_anonymous"

	// EdgeSessions holds the string denoting the sessions edge name in mutations.
	EdgeSessions = "sessions"

	// Table holds the table name of the user in the database.
	Table = "users"
	// SessionsTable is the table the holds the sessions relation/edge.
	SessionsTable = "sessions"
	// SessionsInverseTable is the table name for the Session entity.
	// It exists in this package in order to avoid circular dependency with the "session" package.
	SessionsInverseTable = "sessions"
	// SessionsColumn is the table column denoting the sessions relation/edge.
	SessionsColumn = "session_user"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldIsAnonymous,
}
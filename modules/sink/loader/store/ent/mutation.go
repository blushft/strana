// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/blushft/strana/modules/sink/loader/store/ent/rawevent"
	"github.com/google/uuid"

	"github.com/facebook/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeRawEvent = "RawEvent"
)

// RawEventMutation represents an operation that mutate the RawEvents
// nodes in the graph.
type RawEventMutation struct {
	config
	op              Op
	typ             string
	id              *uuid.UUID
	tracking_id     *uuid.UUID
	user_id         *string
	anonymous       *bool
	group_id        *string
	session_id      *string
	device_id       *string
	event           *string
	non_interactive *bool
	channel         *string
	platform        *string
	timestamp       *time.Time
	context         *map[string]interface{}
	clearedFields   map[string]struct{}
	done            bool
	oldValue        func(context.Context) (*RawEvent, error)
}

var _ ent.Mutation = (*RawEventMutation)(nil)

// raweventOption allows to manage the mutation configuration using functional options.
type raweventOption func(*RawEventMutation)

// newRawEventMutation creates new mutation for $n.Name.
func newRawEventMutation(c config, op Op, opts ...raweventOption) *RawEventMutation {
	m := &RawEventMutation{
		config:        c,
		op:            op,
		typ:           TypeRawEvent,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withRawEventID sets the id field of the mutation.
func withRawEventID(id uuid.UUID) raweventOption {
	return func(m *RawEventMutation) {
		var (
			err   error
			once  sync.Once
			value *RawEvent
		)
		m.oldValue = func(ctx context.Context) (*RawEvent, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().RawEvent.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withRawEvent sets the old RawEvent of the mutation.
func withRawEvent(node *RawEvent) raweventOption {
	return func(m *RawEventMutation) {
		m.oldValue = func(context.Context) (*RawEvent, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m RawEventMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m RawEventMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that, this
// operation is accepted only on RawEvent creation.
func (m *RawEventMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the id value in the mutation. Note that, the id
// is available only if it was provided to the builder.
func (m *RawEventMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetTrackingID sets the tracking_id field.
func (m *RawEventMutation) SetTrackingID(u uuid.UUID) {
	m.tracking_id = &u
}

// TrackingID returns the tracking_id value in the mutation.
func (m *RawEventMutation) TrackingID() (r uuid.UUID, exists bool) {
	v := m.tracking_id
	if v == nil {
		return
	}
	return *v, true
}

// OldTrackingID returns the old tracking_id value of the RawEvent.
// If the RawEvent object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *RawEventMutation) OldTrackingID(ctx context.Context) (v uuid.UUID, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldTrackingID is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldTrackingID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTrackingID: %w", err)
	}
	return oldValue.TrackingID, nil
}

// ResetTrackingID reset all changes of the "tracking_id" field.
func (m *RawEventMutation) ResetTrackingID() {
	m.tracking_id = nil
}

// SetUserID sets the user_id field.
func (m *RawEventMutation) SetUserID(s string) {
	m.user_id = &s
}

// UserID returns the user_id value in the mutation.
func (m *RawEventMutation) UserID() (r string, exists bool) {
	v := m.user_id
	if v == nil {
		return
	}
	return *v, true
}

// OldUserID returns the old user_id value of the RawEvent.
// If the RawEvent object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *RawEventMutation) OldUserID(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldUserID is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldUserID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUserID: %w", err)
	}
	return oldValue.UserID, nil
}

// ResetUserID reset all changes of the "user_id" field.
func (m *RawEventMutation) ResetUserID() {
	m.user_id = nil
}

// SetAnonymous sets the anonymous field.
func (m *RawEventMutation) SetAnonymous(b bool) {
	m.anonymous = &b
}

// Anonymous returns the anonymous value in the mutation.
func (m *RawEventMutation) Anonymous() (r bool, exists bool) {
	v := m.anonymous
	if v == nil {
		return
	}
	return *v, true
}

// OldAnonymous returns the old anonymous value of the RawEvent.
// If the RawEvent object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *RawEventMutation) OldAnonymous(ctx context.Context) (v bool, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldAnonymous is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldAnonymous requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAnonymous: %w", err)
	}
	return oldValue.Anonymous, nil
}

// ResetAnonymous reset all changes of the "anonymous" field.
func (m *RawEventMutation) ResetAnonymous() {
	m.anonymous = nil
}

// SetGroupID sets the group_id field.
func (m *RawEventMutation) SetGroupID(s string) {
	m.group_id = &s
}

// GroupID returns the group_id value in the mutation.
func (m *RawEventMutation) GroupID() (r string, exists bool) {
	v := m.group_id
	if v == nil {
		return
	}
	return *v, true
}

// OldGroupID returns the old group_id value of the RawEvent.
// If the RawEvent object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *RawEventMutation) OldGroupID(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldGroupID is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldGroupID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldGroupID: %w", err)
	}
	return oldValue.GroupID, nil
}

// ClearGroupID clears the value of group_id.
func (m *RawEventMutation) ClearGroupID() {
	m.group_id = nil
	m.clearedFields[rawevent.FieldGroupID] = struct{}{}
}

// GroupIDCleared returns if the field group_id was cleared in this mutation.
func (m *RawEventMutation) GroupIDCleared() bool {
	_, ok := m.clearedFields[rawevent.FieldGroupID]
	return ok
}

// ResetGroupID reset all changes of the "group_id" field.
func (m *RawEventMutation) ResetGroupID() {
	m.group_id = nil
	delete(m.clearedFields, rawevent.FieldGroupID)
}

// SetSessionID sets the session_id field.
func (m *RawEventMutation) SetSessionID(s string) {
	m.session_id = &s
}

// SessionID returns the session_id value in the mutation.
func (m *RawEventMutation) SessionID() (r string, exists bool) {
	v := m.session_id
	if v == nil {
		return
	}
	return *v, true
}

// OldSessionID returns the old session_id value of the RawEvent.
// If the RawEvent object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *RawEventMutation) OldSessionID(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldSessionID is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldSessionID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldSessionID: %w", err)
	}
	return oldValue.SessionID, nil
}

// ClearSessionID clears the value of session_id.
func (m *RawEventMutation) ClearSessionID() {
	m.session_id = nil
	m.clearedFields[rawevent.FieldSessionID] = struct{}{}
}

// SessionIDCleared returns if the field session_id was cleared in this mutation.
func (m *RawEventMutation) SessionIDCleared() bool {
	_, ok := m.clearedFields[rawevent.FieldSessionID]
	return ok
}

// ResetSessionID reset all changes of the "session_id" field.
func (m *RawEventMutation) ResetSessionID() {
	m.session_id = nil
	delete(m.clearedFields, rawevent.FieldSessionID)
}

// SetDeviceID sets the device_id field.
func (m *RawEventMutation) SetDeviceID(s string) {
	m.device_id = &s
}

// DeviceID returns the device_id value in the mutation.
func (m *RawEventMutation) DeviceID() (r string, exists bool) {
	v := m.device_id
	if v == nil {
		return
	}
	return *v, true
}

// OldDeviceID returns the old device_id value of the RawEvent.
// If the RawEvent object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *RawEventMutation) OldDeviceID(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldDeviceID is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldDeviceID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDeviceID: %w", err)
	}
	return oldValue.DeviceID, nil
}

// ClearDeviceID clears the value of device_id.
func (m *RawEventMutation) ClearDeviceID() {
	m.device_id = nil
	m.clearedFields[rawevent.FieldDeviceID] = struct{}{}
}

// DeviceIDCleared returns if the field device_id was cleared in this mutation.
func (m *RawEventMutation) DeviceIDCleared() bool {
	_, ok := m.clearedFields[rawevent.FieldDeviceID]
	return ok
}

// ResetDeviceID reset all changes of the "device_id" field.
func (m *RawEventMutation) ResetDeviceID() {
	m.device_id = nil
	delete(m.clearedFields, rawevent.FieldDeviceID)
}

// SetEvent sets the event field.
func (m *RawEventMutation) SetEvent(s string) {
	m.event = &s
}

// Event returns the event value in the mutation.
func (m *RawEventMutation) Event() (r string, exists bool) {
	v := m.event
	if v == nil {
		return
	}
	return *v, true
}

// OldEvent returns the old event value of the RawEvent.
// If the RawEvent object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *RawEventMutation) OldEvent(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldEvent is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldEvent requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldEvent: %w", err)
	}
	return oldValue.Event, nil
}

// ResetEvent reset all changes of the "event" field.
func (m *RawEventMutation) ResetEvent() {
	m.event = nil
}

// SetNonInteractive sets the non_interactive field.
func (m *RawEventMutation) SetNonInteractive(b bool) {
	m.non_interactive = &b
}

// NonInteractive returns the non_interactive value in the mutation.
func (m *RawEventMutation) NonInteractive() (r bool, exists bool) {
	v := m.non_interactive
	if v == nil {
		return
	}
	return *v, true
}

// OldNonInteractive returns the old non_interactive value of the RawEvent.
// If the RawEvent object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *RawEventMutation) OldNonInteractive(ctx context.Context) (v bool, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldNonInteractive is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldNonInteractive requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldNonInteractive: %w", err)
	}
	return oldValue.NonInteractive, nil
}

// ResetNonInteractive reset all changes of the "non_interactive" field.
func (m *RawEventMutation) ResetNonInteractive() {
	m.non_interactive = nil
}

// SetChannel sets the channel field.
func (m *RawEventMutation) SetChannel(s string) {
	m.channel = &s
}

// Channel returns the channel value in the mutation.
func (m *RawEventMutation) Channel() (r string, exists bool) {
	v := m.channel
	if v == nil {
		return
	}
	return *v, true
}

// OldChannel returns the old channel value of the RawEvent.
// If the RawEvent object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *RawEventMutation) OldChannel(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldChannel is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldChannel requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldChannel: %w", err)
	}
	return oldValue.Channel, nil
}

// ClearChannel clears the value of channel.
func (m *RawEventMutation) ClearChannel() {
	m.channel = nil
	m.clearedFields[rawevent.FieldChannel] = struct{}{}
}

// ChannelCleared returns if the field channel was cleared in this mutation.
func (m *RawEventMutation) ChannelCleared() bool {
	_, ok := m.clearedFields[rawevent.FieldChannel]
	return ok
}

// ResetChannel reset all changes of the "channel" field.
func (m *RawEventMutation) ResetChannel() {
	m.channel = nil
	delete(m.clearedFields, rawevent.FieldChannel)
}

// SetPlatform sets the platform field.
func (m *RawEventMutation) SetPlatform(s string) {
	m.platform = &s
}

// Platform returns the platform value in the mutation.
func (m *RawEventMutation) Platform() (r string, exists bool) {
	v := m.platform
	if v == nil {
		return
	}
	return *v, true
}

// OldPlatform returns the old platform value of the RawEvent.
// If the RawEvent object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *RawEventMutation) OldPlatform(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldPlatform is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldPlatform requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPlatform: %w", err)
	}
	return oldValue.Platform, nil
}

// ClearPlatform clears the value of platform.
func (m *RawEventMutation) ClearPlatform() {
	m.platform = nil
	m.clearedFields[rawevent.FieldPlatform] = struct{}{}
}

// PlatformCleared returns if the field platform was cleared in this mutation.
func (m *RawEventMutation) PlatformCleared() bool {
	_, ok := m.clearedFields[rawevent.FieldPlatform]
	return ok
}

// ResetPlatform reset all changes of the "platform" field.
func (m *RawEventMutation) ResetPlatform() {
	m.platform = nil
	delete(m.clearedFields, rawevent.FieldPlatform)
}

// SetTimestamp sets the timestamp field.
func (m *RawEventMutation) SetTimestamp(t time.Time) {
	m.timestamp = &t
}

// Timestamp returns the timestamp value in the mutation.
func (m *RawEventMutation) Timestamp() (r time.Time, exists bool) {
	v := m.timestamp
	if v == nil {
		return
	}
	return *v, true
}

// OldTimestamp returns the old timestamp value of the RawEvent.
// If the RawEvent object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *RawEventMutation) OldTimestamp(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldTimestamp is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldTimestamp requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTimestamp: %w", err)
	}
	return oldValue.Timestamp, nil
}

// ResetTimestamp reset all changes of the "timestamp" field.
func (m *RawEventMutation) ResetTimestamp() {
	m.timestamp = nil
}

// SetContext sets the context field.
func (m *RawEventMutation) SetContext(value map[string]interface{}) {
	m.context = &value
}

// Context returns the context value in the mutation.
func (m *RawEventMutation) Context() (r map[string]interface{}, exists bool) {
	v := m.context
	if v == nil {
		return
	}
	return *v, true
}

// OldContext returns the old context value of the RawEvent.
// If the RawEvent object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *RawEventMutation) OldContext(ctx context.Context) (v map[string]interface{}, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldContext is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldContext requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldContext: %w", err)
	}
	return oldValue.Context, nil
}

// ResetContext reset all changes of the "context" field.
func (m *RawEventMutation) ResetContext() {
	m.context = nil
}

// Op returns the operation name.
func (m *RawEventMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (RawEvent).
func (m *RawEventMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during
// this mutation. Note that, in order to get all numeric
// fields that were in/decremented, call AddedFields().
func (m *RawEventMutation) Fields() []string {
	fields := make([]string, 0, 12)
	if m.tracking_id != nil {
		fields = append(fields, rawevent.FieldTrackingID)
	}
	if m.user_id != nil {
		fields = append(fields, rawevent.FieldUserID)
	}
	if m.anonymous != nil {
		fields = append(fields, rawevent.FieldAnonymous)
	}
	if m.group_id != nil {
		fields = append(fields, rawevent.FieldGroupID)
	}
	if m.session_id != nil {
		fields = append(fields, rawevent.FieldSessionID)
	}
	if m.device_id != nil {
		fields = append(fields, rawevent.FieldDeviceID)
	}
	if m.event != nil {
		fields = append(fields, rawevent.FieldEvent)
	}
	if m.non_interactive != nil {
		fields = append(fields, rawevent.FieldNonInteractive)
	}
	if m.channel != nil {
		fields = append(fields, rawevent.FieldChannel)
	}
	if m.platform != nil {
		fields = append(fields, rawevent.FieldPlatform)
	}
	if m.timestamp != nil {
		fields = append(fields, rawevent.FieldTimestamp)
	}
	if m.context != nil {
		fields = append(fields, rawevent.FieldContext)
	}
	return fields
}

// Field returns the value of a field with the given name.
// The second boolean value indicates that this field was
// not set, or was not define in the schema.
func (m *RawEventMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case rawevent.FieldTrackingID:
		return m.TrackingID()
	case rawevent.FieldUserID:
		return m.UserID()
	case rawevent.FieldAnonymous:
		return m.Anonymous()
	case rawevent.FieldGroupID:
		return m.GroupID()
	case rawevent.FieldSessionID:
		return m.SessionID()
	case rawevent.FieldDeviceID:
		return m.DeviceID()
	case rawevent.FieldEvent:
		return m.Event()
	case rawevent.FieldNonInteractive:
		return m.NonInteractive()
	case rawevent.FieldChannel:
		return m.Channel()
	case rawevent.FieldPlatform:
		return m.Platform()
	case rawevent.FieldTimestamp:
		return m.Timestamp()
	case rawevent.FieldContext:
		return m.Context()
	}
	return nil, false
}

// OldField returns the old value of the field from the database.
// An error is returned if the mutation operation is not UpdateOne,
// or the query to the database was failed.
func (m *RawEventMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case rawevent.FieldTrackingID:
		return m.OldTrackingID(ctx)
	case rawevent.FieldUserID:
		return m.OldUserID(ctx)
	case rawevent.FieldAnonymous:
		return m.OldAnonymous(ctx)
	case rawevent.FieldGroupID:
		return m.OldGroupID(ctx)
	case rawevent.FieldSessionID:
		return m.OldSessionID(ctx)
	case rawevent.FieldDeviceID:
		return m.OldDeviceID(ctx)
	case rawevent.FieldEvent:
		return m.OldEvent(ctx)
	case rawevent.FieldNonInteractive:
		return m.OldNonInteractive(ctx)
	case rawevent.FieldChannel:
		return m.OldChannel(ctx)
	case rawevent.FieldPlatform:
		return m.OldPlatform(ctx)
	case rawevent.FieldTimestamp:
		return m.OldTimestamp(ctx)
	case rawevent.FieldContext:
		return m.OldContext(ctx)
	}
	return nil, fmt.Errorf("unknown RawEvent field %s", name)
}

// SetField sets the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *RawEventMutation) SetField(name string, value ent.Value) error {
	switch name {
	case rawevent.FieldTrackingID:
		v, ok := value.(uuid.UUID)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTrackingID(v)
		return nil
	case rawevent.FieldUserID:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUserID(v)
		return nil
	case rawevent.FieldAnonymous:
		v, ok := value.(bool)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAnonymous(v)
		return nil
	case rawevent.FieldGroupID:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetGroupID(v)
		return nil
	case rawevent.FieldSessionID:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetSessionID(v)
		return nil
	case rawevent.FieldDeviceID:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDeviceID(v)
		return nil
	case rawevent.FieldEvent:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetEvent(v)
		return nil
	case rawevent.FieldNonInteractive:
		v, ok := value.(bool)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetNonInteractive(v)
		return nil
	case rawevent.FieldChannel:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetChannel(v)
		return nil
	case rawevent.FieldPlatform:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPlatform(v)
		return nil
	case rawevent.FieldTimestamp:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTimestamp(v)
		return nil
	case rawevent.FieldContext:
		v, ok := value.(map[string]interface{})
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetContext(v)
		return nil
	}
	return fmt.Errorf("unknown RawEvent field %s", name)
}

// AddedFields returns all numeric fields that were incremented
// or decremented during this mutation.
func (m *RawEventMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was in/decremented
// from a field with the given name. The second value indicates
// that this field was not set, or was not define in the schema.
func (m *RawEventMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *RawEventMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown RawEvent numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared
// during this mutation.
func (m *RawEventMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(rawevent.FieldGroupID) {
		fields = append(fields, rawevent.FieldGroupID)
	}
	if m.FieldCleared(rawevent.FieldSessionID) {
		fields = append(fields, rawevent.FieldSessionID)
	}
	if m.FieldCleared(rawevent.FieldDeviceID) {
		fields = append(fields, rawevent.FieldDeviceID)
	}
	if m.FieldCleared(rawevent.FieldChannel) {
		fields = append(fields, rawevent.FieldChannel)
	}
	if m.FieldCleared(rawevent.FieldPlatform) {
		fields = append(fields, rawevent.FieldPlatform)
	}
	return fields
}

// FieldCleared returns a boolean indicates if this field was
// cleared in this mutation.
func (m *RawEventMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value for the given name. It returns an
// error if the field is not defined in the schema.
func (m *RawEventMutation) ClearField(name string) error {
	switch name {
	case rawevent.FieldGroupID:
		m.ClearGroupID()
		return nil
	case rawevent.FieldSessionID:
		m.ClearSessionID()
		return nil
	case rawevent.FieldDeviceID:
		m.ClearDeviceID()
		return nil
	case rawevent.FieldChannel:
		m.ClearChannel()
		return nil
	case rawevent.FieldPlatform:
		m.ClearPlatform()
		return nil
	}
	return fmt.Errorf("unknown RawEvent nullable field %s", name)
}

// ResetField resets all changes in the mutation regarding the
// given field name. It returns an error if the field is not
// defined in the schema.
func (m *RawEventMutation) ResetField(name string) error {
	switch name {
	case rawevent.FieldTrackingID:
		m.ResetTrackingID()
		return nil
	case rawevent.FieldUserID:
		m.ResetUserID()
		return nil
	case rawevent.FieldAnonymous:
		m.ResetAnonymous()
		return nil
	case rawevent.FieldGroupID:
		m.ResetGroupID()
		return nil
	case rawevent.FieldSessionID:
		m.ResetSessionID()
		return nil
	case rawevent.FieldDeviceID:
		m.ResetDeviceID()
		return nil
	case rawevent.FieldEvent:
		m.ResetEvent()
		return nil
	case rawevent.FieldNonInteractive:
		m.ResetNonInteractive()
		return nil
	case rawevent.FieldChannel:
		m.ResetChannel()
		return nil
	case rawevent.FieldPlatform:
		m.ResetPlatform()
		return nil
	case rawevent.FieldTimestamp:
		m.ResetTimestamp()
		return nil
	case rawevent.FieldContext:
		m.ResetContext()
		return nil
	}
	return fmt.Errorf("unknown RawEvent field %s", name)
}

// AddedEdges returns all edge names that were set/added in this
// mutation.
func (m *RawEventMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all ids (to other nodes) that were added for
// the given edge name.
func (m *RawEventMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this
// mutation.
func (m *RawEventMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all ids (to other nodes) that were removed for
// the given edge name.
func (m *RawEventMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this
// mutation.
func (m *RawEventMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean indicates if this edge was
// cleared in this mutation.
func (m *RawEventMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value for the given name. It returns an
// error if the edge name is not defined in the schema.
func (m *RawEventMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown RawEvent unique edge %s", name)
}

// ResetEdge resets all changes in the mutation regarding the
// given edge name. It returns an error if the edge is not
// defined in the schema.
func (m *RawEventMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown RawEvent edge %s", name)
}

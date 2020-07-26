// Code generated by entc, DO NOT EDIT.

package pageview

import (
	"time"

	"github.com/blushft/strana/platform/store/ent/predicate"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// ID filters vertices based on their identifier.
func ID(id uuid.UUID) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Hostname applies equality check predicate on the "hostname" field. It's identical to HostnameEQ.
func Hostname(v string) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHostname), v))
	})
}

// Pathname applies equality check predicate on the "pathname" field. It's identical to PathnameEQ.
func Pathname(v string) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPathname), v))
	})
}

// Referrer applies equality check predicate on the "referrer" field. It's identical to ReferrerEQ.
func Referrer(v string) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldReferrer), v))
	})
}

// IsEntry applies equality check predicate on the "is_entry" field. It's identical to IsEntryEQ.
func IsEntry(v bool) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsEntry), v))
	})
}

// IsFinished applies equality check predicate on the "is_finished" field. It's identical to IsFinishedEQ.
func IsFinished(v bool) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsFinished), v))
	})
}

// Duration applies equality check predicate on the "duration" field. It's identical to DurationEQ.
func Duration(v int) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDuration), v))
	})
}

// Timestamp applies equality check predicate on the "timestamp" field. It's identical to TimestampEQ.
func Timestamp(v time.Time) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTimestamp), v))
	})
}

// UserAgent applies equality check predicate on the "user_agent" field. It's identical to UserAgentEQ.
func UserAgent(v string) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserAgent), v))
	})
}

// IPAddress applies equality check predicate on the "ip_address" field. It's identical to IPAddressEQ.
func IPAddress(v string) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIPAddress), v))
	})
}

// ScreenDim applies equality check predicate on the "screen_dim" field. It's identical to ScreenDimEQ.
func ScreenDim(v string) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldScreenDim), v))
	})
}

// HostnameEQ applies the EQ predicate on the "hostname" field.
func HostnameEQ(v string) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHostname), v))
	})
}

// HostnameNEQ applies the NEQ predicate on the "hostname" field.
func HostnameNEQ(v string) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldHostname), v))
	})
}

// HostnameIn applies the In predicate on the "hostname" field.
func HostnameIn(vs ...string) predicate.PageView {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PageView(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldHostname), v...))
	})
}

// HostnameNotIn applies the NotIn predicate on the "hostname" field.
func HostnameNotIn(vs ...string) predicate.PageView {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PageView(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldHostname), v...))
	})
}

// HostnameGT applies the GT predicate on the "hostname" field.
func HostnameGT(v string) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldHostname), v))
	})
}

// HostnameGTE applies the GTE predicate on the "hostname" field.
func HostnameGTE(v string) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldHostname), v))
	})
}

// HostnameLT applies the LT predicate on the "hostname" field.
func HostnameLT(v string) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldHostname), v))
	})
}

// HostnameLTE applies the LTE predicate on the "hostname" field.
func HostnameLTE(v string) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldHostname), v))
	})
}

// HostnameContains applies the Contains predicate on the "hostname" field.
func HostnameContains(v string) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldHostname), v))
	})
}

// HostnameHasPrefix applies the HasPrefix predicate on the "hostname" field.
func HostnameHasPrefix(v string) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldHostname), v))
	})
}

// HostnameHasSuffix applies the HasSuffix predicate on the "hostname" field.
func HostnameHasSuffix(v string) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldHostname), v))
	})
}

// HostnameEqualFold applies the EqualFold predicate on the "hostname" field.
func HostnameEqualFold(v string) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldHostname), v))
	})
}

// HostnameContainsFold applies the ContainsFold predicate on the "hostname" field.
func HostnameContainsFold(v string) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldHostname), v))
	})
}

// PathnameEQ applies the EQ predicate on the "pathname" field.
func PathnameEQ(v string) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPathname), v))
	})
}

// PathnameNEQ applies the NEQ predicate on the "pathname" field.
func PathnameNEQ(v string) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPathname), v))
	})
}

// PathnameIn applies the In predicate on the "pathname" field.
func PathnameIn(vs ...string) predicate.PageView {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PageView(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPathname), v...))
	})
}

// PathnameNotIn applies the NotIn predicate on the "pathname" field.
func PathnameNotIn(vs ...string) predicate.PageView {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PageView(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPathname), v...))
	})
}

// PathnameGT applies the GT predicate on the "pathname" field.
func PathnameGT(v string) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPathname), v))
	})
}

// PathnameGTE applies the GTE predicate on the "pathname" field.
func PathnameGTE(v string) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPathname), v))
	})
}

// PathnameLT applies the LT predicate on the "pathname" field.
func PathnameLT(v string) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPathname), v))
	})
}

// PathnameLTE applies the LTE predicate on the "pathname" field.
func PathnameLTE(v string) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPathname), v))
	})
}

// PathnameContains applies the Contains predicate on the "pathname" field.
func PathnameContains(v string) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPathname), v))
	})
}

// PathnameHasPrefix applies the HasPrefix predicate on the "pathname" field.
func PathnameHasPrefix(v string) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPathname), v))
	})
}

// PathnameHasSuffix applies the HasSuffix predicate on the "pathname" field.
func PathnameHasSuffix(v string) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPathname), v))
	})
}

// PathnameEqualFold applies the EqualFold predicate on the "pathname" field.
func PathnameEqualFold(v string) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPathname), v))
	})
}

// PathnameContainsFold applies the ContainsFold predicate on the "pathname" field.
func PathnameContainsFold(v string) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPathname), v))
	})
}

// ReferrerEQ applies the EQ predicate on the "referrer" field.
func ReferrerEQ(v string) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldReferrer), v))
	})
}

// ReferrerNEQ applies the NEQ predicate on the "referrer" field.
func ReferrerNEQ(v string) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldReferrer), v))
	})
}

// ReferrerIn applies the In predicate on the "referrer" field.
func ReferrerIn(vs ...string) predicate.PageView {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PageView(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldReferrer), v...))
	})
}

// ReferrerNotIn applies the NotIn predicate on the "referrer" field.
func ReferrerNotIn(vs ...string) predicate.PageView {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PageView(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldReferrer), v...))
	})
}

// ReferrerGT applies the GT predicate on the "referrer" field.
func ReferrerGT(v string) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldReferrer), v))
	})
}

// ReferrerGTE applies the GTE predicate on the "referrer" field.
func ReferrerGTE(v string) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldReferrer), v))
	})
}

// ReferrerLT applies the LT predicate on the "referrer" field.
func ReferrerLT(v string) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldReferrer), v))
	})
}

// ReferrerLTE applies the LTE predicate on the "referrer" field.
func ReferrerLTE(v string) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldReferrer), v))
	})
}

// ReferrerContains applies the Contains predicate on the "referrer" field.
func ReferrerContains(v string) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldReferrer), v))
	})
}

// ReferrerHasPrefix applies the HasPrefix predicate on the "referrer" field.
func ReferrerHasPrefix(v string) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldReferrer), v))
	})
}

// ReferrerHasSuffix applies the HasSuffix predicate on the "referrer" field.
func ReferrerHasSuffix(v string) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldReferrer), v))
	})
}

// ReferrerEqualFold applies the EqualFold predicate on the "referrer" field.
func ReferrerEqualFold(v string) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldReferrer), v))
	})
}

// ReferrerContainsFold applies the ContainsFold predicate on the "referrer" field.
func ReferrerContainsFold(v string) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldReferrer), v))
	})
}

// IsEntryEQ applies the EQ predicate on the "is_entry" field.
func IsEntryEQ(v bool) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsEntry), v))
	})
}

// IsEntryNEQ applies the NEQ predicate on the "is_entry" field.
func IsEntryNEQ(v bool) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIsEntry), v))
	})
}

// IsFinishedEQ applies the EQ predicate on the "is_finished" field.
func IsFinishedEQ(v bool) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsFinished), v))
	})
}

// IsFinishedNEQ applies the NEQ predicate on the "is_finished" field.
func IsFinishedNEQ(v bool) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIsFinished), v))
	})
}

// DurationEQ applies the EQ predicate on the "duration" field.
func DurationEQ(v int) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDuration), v))
	})
}

// DurationNEQ applies the NEQ predicate on the "duration" field.
func DurationNEQ(v int) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDuration), v))
	})
}

// DurationIn applies the In predicate on the "duration" field.
func DurationIn(vs ...int) predicate.PageView {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PageView(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDuration), v...))
	})
}

// DurationNotIn applies the NotIn predicate on the "duration" field.
func DurationNotIn(vs ...int) predicate.PageView {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PageView(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDuration), v...))
	})
}

// DurationGT applies the GT predicate on the "duration" field.
func DurationGT(v int) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDuration), v))
	})
}

// DurationGTE applies the GTE predicate on the "duration" field.
func DurationGTE(v int) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDuration), v))
	})
}

// DurationLT applies the LT predicate on the "duration" field.
func DurationLT(v int) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDuration), v))
	})
}

// DurationLTE applies the LTE predicate on the "duration" field.
func DurationLTE(v int) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDuration), v))
	})
}

// TimestampEQ applies the EQ predicate on the "timestamp" field.
func TimestampEQ(v time.Time) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTimestamp), v))
	})
}

// TimestampNEQ applies the NEQ predicate on the "timestamp" field.
func TimestampNEQ(v time.Time) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTimestamp), v))
	})
}

// TimestampIn applies the In predicate on the "timestamp" field.
func TimestampIn(vs ...time.Time) predicate.PageView {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PageView(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTimestamp), v...))
	})
}

// TimestampNotIn applies the NotIn predicate on the "timestamp" field.
func TimestampNotIn(vs ...time.Time) predicate.PageView {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PageView(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTimestamp), v...))
	})
}

// TimestampGT applies the GT predicate on the "timestamp" field.
func TimestampGT(v time.Time) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTimestamp), v))
	})
}

// TimestampGTE applies the GTE predicate on the "timestamp" field.
func TimestampGTE(v time.Time) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTimestamp), v))
	})
}

// TimestampLT applies the LT predicate on the "timestamp" field.
func TimestampLT(v time.Time) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTimestamp), v))
	})
}

// TimestampLTE applies the LTE predicate on the "timestamp" field.
func TimestampLTE(v time.Time) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTimestamp), v))
	})
}

// UserAgentEQ applies the EQ predicate on the "user_agent" field.
func UserAgentEQ(v string) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserAgent), v))
	})
}

// UserAgentNEQ applies the NEQ predicate on the "user_agent" field.
func UserAgentNEQ(v string) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserAgent), v))
	})
}

// UserAgentIn applies the In predicate on the "user_agent" field.
func UserAgentIn(vs ...string) predicate.PageView {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PageView(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUserAgent), v...))
	})
}

// UserAgentNotIn applies the NotIn predicate on the "user_agent" field.
func UserAgentNotIn(vs ...string) predicate.PageView {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PageView(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUserAgent), v...))
	})
}

// UserAgentGT applies the GT predicate on the "user_agent" field.
func UserAgentGT(v string) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUserAgent), v))
	})
}

// UserAgentGTE applies the GTE predicate on the "user_agent" field.
func UserAgentGTE(v string) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUserAgent), v))
	})
}

// UserAgentLT applies the LT predicate on the "user_agent" field.
func UserAgentLT(v string) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUserAgent), v))
	})
}

// UserAgentLTE applies the LTE predicate on the "user_agent" field.
func UserAgentLTE(v string) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUserAgent), v))
	})
}

// UserAgentContains applies the Contains predicate on the "user_agent" field.
func UserAgentContains(v string) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldUserAgent), v))
	})
}

// UserAgentHasPrefix applies the HasPrefix predicate on the "user_agent" field.
func UserAgentHasPrefix(v string) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldUserAgent), v))
	})
}

// UserAgentHasSuffix applies the HasSuffix predicate on the "user_agent" field.
func UserAgentHasSuffix(v string) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldUserAgent), v))
	})
}

// UserAgentIsNil applies the IsNil predicate on the "user_agent" field.
func UserAgentIsNil() predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUserAgent)))
	})
}

// UserAgentNotNil applies the NotNil predicate on the "user_agent" field.
func UserAgentNotNil() predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUserAgent)))
	})
}

// UserAgentEqualFold applies the EqualFold predicate on the "user_agent" field.
func UserAgentEqualFold(v string) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldUserAgent), v))
	})
}

// UserAgentContainsFold applies the ContainsFold predicate on the "user_agent" field.
func UserAgentContainsFold(v string) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldUserAgent), v))
	})
}

// IPAddressEQ applies the EQ predicate on the "ip_address" field.
func IPAddressEQ(v string) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIPAddress), v))
	})
}

// IPAddressNEQ applies the NEQ predicate on the "ip_address" field.
func IPAddressNEQ(v string) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIPAddress), v))
	})
}

// IPAddressIn applies the In predicate on the "ip_address" field.
func IPAddressIn(vs ...string) predicate.PageView {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PageView(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldIPAddress), v...))
	})
}

// IPAddressNotIn applies the NotIn predicate on the "ip_address" field.
func IPAddressNotIn(vs ...string) predicate.PageView {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PageView(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldIPAddress), v...))
	})
}

// IPAddressGT applies the GT predicate on the "ip_address" field.
func IPAddressGT(v string) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldIPAddress), v))
	})
}

// IPAddressGTE applies the GTE predicate on the "ip_address" field.
func IPAddressGTE(v string) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldIPAddress), v))
	})
}

// IPAddressLT applies the LT predicate on the "ip_address" field.
func IPAddressLT(v string) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldIPAddress), v))
	})
}

// IPAddressLTE applies the LTE predicate on the "ip_address" field.
func IPAddressLTE(v string) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldIPAddress), v))
	})
}

// IPAddressContains applies the Contains predicate on the "ip_address" field.
func IPAddressContains(v string) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldIPAddress), v))
	})
}

// IPAddressHasPrefix applies the HasPrefix predicate on the "ip_address" field.
func IPAddressHasPrefix(v string) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldIPAddress), v))
	})
}

// IPAddressHasSuffix applies the HasSuffix predicate on the "ip_address" field.
func IPAddressHasSuffix(v string) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldIPAddress), v))
	})
}

// IPAddressIsNil applies the IsNil predicate on the "ip_address" field.
func IPAddressIsNil() predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldIPAddress)))
	})
}

// IPAddressNotNil applies the NotNil predicate on the "ip_address" field.
func IPAddressNotNil() predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldIPAddress)))
	})
}

// IPAddressEqualFold applies the EqualFold predicate on the "ip_address" field.
func IPAddressEqualFold(v string) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldIPAddress), v))
	})
}

// IPAddressContainsFold applies the ContainsFold predicate on the "ip_address" field.
func IPAddressContainsFold(v string) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldIPAddress), v))
	})
}

// ScreenDimEQ applies the EQ predicate on the "screen_dim" field.
func ScreenDimEQ(v string) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldScreenDim), v))
	})
}

// ScreenDimNEQ applies the NEQ predicate on the "screen_dim" field.
func ScreenDimNEQ(v string) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldScreenDim), v))
	})
}

// ScreenDimIn applies the In predicate on the "screen_dim" field.
func ScreenDimIn(vs ...string) predicate.PageView {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PageView(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldScreenDim), v...))
	})
}

// ScreenDimNotIn applies the NotIn predicate on the "screen_dim" field.
func ScreenDimNotIn(vs ...string) predicate.PageView {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PageView(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldScreenDim), v...))
	})
}

// ScreenDimGT applies the GT predicate on the "screen_dim" field.
func ScreenDimGT(v string) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldScreenDim), v))
	})
}

// ScreenDimGTE applies the GTE predicate on the "screen_dim" field.
func ScreenDimGTE(v string) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldScreenDim), v))
	})
}

// ScreenDimLT applies the LT predicate on the "screen_dim" field.
func ScreenDimLT(v string) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldScreenDim), v))
	})
}

// ScreenDimLTE applies the LTE predicate on the "screen_dim" field.
func ScreenDimLTE(v string) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldScreenDim), v))
	})
}

// ScreenDimContains applies the Contains predicate on the "screen_dim" field.
func ScreenDimContains(v string) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldScreenDim), v))
	})
}

// ScreenDimHasPrefix applies the HasPrefix predicate on the "screen_dim" field.
func ScreenDimHasPrefix(v string) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldScreenDim), v))
	})
}

// ScreenDimHasSuffix applies the HasSuffix predicate on the "screen_dim" field.
func ScreenDimHasSuffix(v string) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldScreenDim), v))
	})
}

// ScreenDimIsNil applies the IsNil predicate on the "screen_dim" field.
func ScreenDimIsNil() predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldScreenDim)))
	})
}

// ScreenDimNotNil applies the NotNil predicate on the "screen_dim" field.
func ScreenDimNotNil() predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldScreenDim)))
	})
}

// ScreenDimEqualFold applies the EqualFold predicate on the "screen_dim" field.
func ScreenDimEqualFold(v string) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldScreenDim), v))
	})
}

// ScreenDimContainsFold applies the ContainsFold predicate on the "screen_dim" field.
func ScreenDimContainsFold(v string) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldScreenDim), v))
	})
}

// ExtraIsNil applies the IsNil predicate on the "extra" field.
func ExtraIsNil() predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldExtra)))
	})
}

// ExtraNotNil applies the NotNil predicate on the "extra" field.
func ExtraNotNil() predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldExtra)))
	})
}

// HasApp applies the HasEdge predicate on the "app" edge.
func HasApp() predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AppTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, AppTable, AppColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAppWith applies the HasEdge predicate on the "app" edge with a given conditions (other predicates).
func HasAppWith(preds ...predicate.App) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AppInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, AppTable, AppColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSession applies the HasEdge predicate on the "session" edge.
func HasSession() predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SessionTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, SessionTable, SessionColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSessionWith applies the HasEdge predicate on the "session" edge with a given conditions (other predicates).
func HasSessionWith(preds ...predicate.Session) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SessionInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, SessionTable, SessionColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, UserTable, UserColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.PageView) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.PageView) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.PageView) predicate.PageView {
	return predicate.PageView(func(s *sql.Selector) {
		p(s.Not())
	})
}

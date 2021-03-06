// Code generated by entc, DO NOT EDIT.

package campaign

import (
	"github.com/blushft/strana/modules/sink/reporter/store/ent/predicate"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Campaign {
	return predicate.Campaign(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Campaign {
	return predicate.Campaign(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Campaign {
	return predicate.Campaign(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Campaign {
	return predicate.Campaign(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Campaign {
	return predicate.Campaign(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Campaign {
	return predicate.Campaign(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Campaign {
	return predicate.Campaign(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Campaign {
	return predicate.Campaign(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Campaign {
	return predicate.Campaign(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Campaign {
	return predicate.Campaign(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Source applies equality check predicate on the "source" field. It's identical to SourceEQ.
func Source(v string) predicate.Campaign {
	return predicate.Campaign(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSource), v))
	})
}

// Medium applies equality check predicate on the "medium" field. It's identical to MediumEQ.
func Medium(v string) predicate.Campaign {
	return predicate.Campaign(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMedium), v))
	})
}

// Term applies equality check predicate on the "term" field. It's identical to TermEQ.
func Term(v string) predicate.Campaign {
	return predicate.Campaign(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTerm), v))
	})
}

// Content applies equality check predicate on the "content" field. It's identical to ContentEQ.
func Content(v string) predicate.Campaign {
	return predicate.Campaign(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContent), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Campaign {
	return predicate.Campaign(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Campaign {
	return predicate.Campaign(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Campaign {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Campaign(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Campaign {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Campaign(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Campaign {
	return predicate.Campaign(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Campaign {
	return predicate.Campaign(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Campaign {
	return predicate.Campaign(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Campaign {
	return predicate.Campaign(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Campaign {
	return predicate.Campaign(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Campaign {
	return predicate.Campaign(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Campaign {
	return predicate.Campaign(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Campaign {
	return predicate.Campaign(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Campaign {
	return predicate.Campaign(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// SourceEQ applies the EQ predicate on the "source" field.
func SourceEQ(v string) predicate.Campaign {
	return predicate.Campaign(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSource), v))
	})
}

// SourceNEQ applies the NEQ predicate on the "source" field.
func SourceNEQ(v string) predicate.Campaign {
	return predicate.Campaign(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSource), v))
	})
}

// SourceIn applies the In predicate on the "source" field.
func SourceIn(vs ...string) predicate.Campaign {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Campaign(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSource), v...))
	})
}

// SourceNotIn applies the NotIn predicate on the "source" field.
func SourceNotIn(vs ...string) predicate.Campaign {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Campaign(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSource), v...))
	})
}

// SourceGT applies the GT predicate on the "source" field.
func SourceGT(v string) predicate.Campaign {
	return predicate.Campaign(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSource), v))
	})
}

// SourceGTE applies the GTE predicate on the "source" field.
func SourceGTE(v string) predicate.Campaign {
	return predicate.Campaign(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSource), v))
	})
}

// SourceLT applies the LT predicate on the "source" field.
func SourceLT(v string) predicate.Campaign {
	return predicate.Campaign(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSource), v))
	})
}

// SourceLTE applies the LTE predicate on the "source" field.
func SourceLTE(v string) predicate.Campaign {
	return predicate.Campaign(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSource), v))
	})
}

// SourceContains applies the Contains predicate on the "source" field.
func SourceContains(v string) predicate.Campaign {
	return predicate.Campaign(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSource), v))
	})
}

// SourceHasPrefix applies the HasPrefix predicate on the "source" field.
func SourceHasPrefix(v string) predicate.Campaign {
	return predicate.Campaign(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSource), v))
	})
}

// SourceHasSuffix applies the HasSuffix predicate on the "source" field.
func SourceHasSuffix(v string) predicate.Campaign {
	return predicate.Campaign(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSource), v))
	})
}

// SourceIsNil applies the IsNil predicate on the "source" field.
func SourceIsNil() predicate.Campaign {
	return predicate.Campaign(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldSource)))
	})
}

// SourceNotNil applies the NotNil predicate on the "source" field.
func SourceNotNil() predicate.Campaign {
	return predicate.Campaign(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldSource)))
	})
}

// SourceEqualFold applies the EqualFold predicate on the "source" field.
func SourceEqualFold(v string) predicate.Campaign {
	return predicate.Campaign(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSource), v))
	})
}

// SourceContainsFold applies the ContainsFold predicate on the "source" field.
func SourceContainsFold(v string) predicate.Campaign {
	return predicate.Campaign(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSource), v))
	})
}

// MediumEQ applies the EQ predicate on the "medium" field.
func MediumEQ(v string) predicate.Campaign {
	return predicate.Campaign(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMedium), v))
	})
}

// MediumNEQ applies the NEQ predicate on the "medium" field.
func MediumNEQ(v string) predicate.Campaign {
	return predicate.Campaign(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMedium), v))
	})
}

// MediumIn applies the In predicate on the "medium" field.
func MediumIn(vs ...string) predicate.Campaign {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Campaign(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMedium), v...))
	})
}

// MediumNotIn applies the NotIn predicate on the "medium" field.
func MediumNotIn(vs ...string) predicate.Campaign {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Campaign(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMedium), v...))
	})
}

// MediumGT applies the GT predicate on the "medium" field.
func MediumGT(v string) predicate.Campaign {
	return predicate.Campaign(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMedium), v))
	})
}

// MediumGTE applies the GTE predicate on the "medium" field.
func MediumGTE(v string) predicate.Campaign {
	return predicate.Campaign(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMedium), v))
	})
}

// MediumLT applies the LT predicate on the "medium" field.
func MediumLT(v string) predicate.Campaign {
	return predicate.Campaign(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMedium), v))
	})
}

// MediumLTE applies the LTE predicate on the "medium" field.
func MediumLTE(v string) predicate.Campaign {
	return predicate.Campaign(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMedium), v))
	})
}

// MediumContains applies the Contains predicate on the "medium" field.
func MediumContains(v string) predicate.Campaign {
	return predicate.Campaign(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldMedium), v))
	})
}

// MediumHasPrefix applies the HasPrefix predicate on the "medium" field.
func MediumHasPrefix(v string) predicate.Campaign {
	return predicate.Campaign(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldMedium), v))
	})
}

// MediumHasSuffix applies the HasSuffix predicate on the "medium" field.
func MediumHasSuffix(v string) predicate.Campaign {
	return predicate.Campaign(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldMedium), v))
	})
}

// MediumIsNil applies the IsNil predicate on the "medium" field.
func MediumIsNil() predicate.Campaign {
	return predicate.Campaign(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldMedium)))
	})
}

// MediumNotNil applies the NotNil predicate on the "medium" field.
func MediumNotNil() predicate.Campaign {
	return predicate.Campaign(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldMedium)))
	})
}

// MediumEqualFold applies the EqualFold predicate on the "medium" field.
func MediumEqualFold(v string) predicate.Campaign {
	return predicate.Campaign(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldMedium), v))
	})
}

// MediumContainsFold applies the ContainsFold predicate on the "medium" field.
func MediumContainsFold(v string) predicate.Campaign {
	return predicate.Campaign(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldMedium), v))
	})
}

// TermEQ applies the EQ predicate on the "term" field.
func TermEQ(v string) predicate.Campaign {
	return predicate.Campaign(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTerm), v))
	})
}

// TermNEQ applies the NEQ predicate on the "term" field.
func TermNEQ(v string) predicate.Campaign {
	return predicate.Campaign(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTerm), v))
	})
}

// TermIn applies the In predicate on the "term" field.
func TermIn(vs ...string) predicate.Campaign {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Campaign(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTerm), v...))
	})
}

// TermNotIn applies the NotIn predicate on the "term" field.
func TermNotIn(vs ...string) predicate.Campaign {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Campaign(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTerm), v...))
	})
}

// TermGT applies the GT predicate on the "term" field.
func TermGT(v string) predicate.Campaign {
	return predicate.Campaign(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTerm), v))
	})
}

// TermGTE applies the GTE predicate on the "term" field.
func TermGTE(v string) predicate.Campaign {
	return predicate.Campaign(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTerm), v))
	})
}

// TermLT applies the LT predicate on the "term" field.
func TermLT(v string) predicate.Campaign {
	return predicate.Campaign(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTerm), v))
	})
}

// TermLTE applies the LTE predicate on the "term" field.
func TermLTE(v string) predicate.Campaign {
	return predicate.Campaign(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTerm), v))
	})
}

// TermContains applies the Contains predicate on the "term" field.
func TermContains(v string) predicate.Campaign {
	return predicate.Campaign(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTerm), v))
	})
}

// TermHasPrefix applies the HasPrefix predicate on the "term" field.
func TermHasPrefix(v string) predicate.Campaign {
	return predicate.Campaign(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTerm), v))
	})
}

// TermHasSuffix applies the HasSuffix predicate on the "term" field.
func TermHasSuffix(v string) predicate.Campaign {
	return predicate.Campaign(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTerm), v))
	})
}

// TermIsNil applies the IsNil predicate on the "term" field.
func TermIsNil() predicate.Campaign {
	return predicate.Campaign(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldTerm)))
	})
}

// TermNotNil applies the NotNil predicate on the "term" field.
func TermNotNil() predicate.Campaign {
	return predicate.Campaign(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldTerm)))
	})
}

// TermEqualFold applies the EqualFold predicate on the "term" field.
func TermEqualFold(v string) predicate.Campaign {
	return predicate.Campaign(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTerm), v))
	})
}

// TermContainsFold applies the ContainsFold predicate on the "term" field.
func TermContainsFold(v string) predicate.Campaign {
	return predicate.Campaign(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTerm), v))
	})
}

// ContentEQ applies the EQ predicate on the "content" field.
func ContentEQ(v string) predicate.Campaign {
	return predicate.Campaign(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContent), v))
	})
}

// ContentNEQ applies the NEQ predicate on the "content" field.
func ContentNEQ(v string) predicate.Campaign {
	return predicate.Campaign(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldContent), v))
	})
}

// ContentIn applies the In predicate on the "content" field.
func ContentIn(vs ...string) predicate.Campaign {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Campaign(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldContent), v...))
	})
}

// ContentNotIn applies the NotIn predicate on the "content" field.
func ContentNotIn(vs ...string) predicate.Campaign {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Campaign(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldContent), v...))
	})
}

// ContentGT applies the GT predicate on the "content" field.
func ContentGT(v string) predicate.Campaign {
	return predicate.Campaign(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldContent), v))
	})
}

// ContentGTE applies the GTE predicate on the "content" field.
func ContentGTE(v string) predicate.Campaign {
	return predicate.Campaign(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldContent), v))
	})
}

// ContentLT applies the LT predicate on the "content" field.
func ContentLT(v string) predicate.Campaign {
	return predicate.Campaign(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldContent), v))
	})
}

// ContentLTE applies the LTE predicate on the "content" field.
func ContentLTE(v string) predicate.Campaign {
	return predicate.Campaign(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldContent), v))
	})
}

// ContentContains applies the Contains predicate on the "content" field.
func ContentContains(v string) predicate.Campaign {
	return predicate.Campaign(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldContent), v))
	})
}

// ContentHasPrefix applies the HasPrefix predicate on the "content" field.
func ContentHasPrefix(v string) predicate.Campaign {
	return predicate.Campaign(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldContent), v))
	})
}

// ContentHasSuffix applies the HasSuffix predicate on the "content" field.
func ContentHasSuffix(v string) predicate.Campaign {
	return predicate.Campaign(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldContent), v))
	})
}

// ContentIsNil applies the IsNil predicate on the "content" field.
func ContentIsNil() predicate.Campaign {
	return predicate.Campaign(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldContent)))
	})
}

// ContentNotNil applies the NotNil predicate on the "content" field.
func ContentNotNil() predicate.Campaign {
	return predicate.Campaign(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldContent)))
	})
}

// ContentEqualFold applies the EqualFold predicate on the "content" field.
func ContentEqualFold(v string) predicate.Campaign {
	return predicate.Campaign(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldContent), v))
	})
}

// ContentContainsFold applies the ContainsFold predicate on the "content" field.
func ContentContainsFold(v string) predicate.Campaign {
	return predicate.Campaign(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldContent), v))
	})
}

// HasEvent applies the HasEdge predicate on the "event" edge.
func HasEvent() predicate.Campaign {
	return predicate.Campaign(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EventTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, EventTable, EventColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEventWith applies the HasEdge predicate on the "event" edge with a given conditions (other predicates).
func HasEventWith(preds ...predicate.Event) predicate.Campaign {
	return predicate.Campaign(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EventInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, EventTable, EventColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Campaign) predicate.Campaign {
	return predicate.Campaign(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Campaign) predicate.Campaign {
	return predicate.Campaign(func(s *sql.Selector) {
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
func Not(p predicate.Campaign) predicate.Campaign {
	return predicate.Campaign(func(s *sql.Selector) {
		p(s.Not())
	})
}

package event

import (
	"reflect"

	"github.com/fatih/structs"
)

type Rule func(*Event) bool

type Validator struct {
	rules []Rule
}

func NewValidator(rules ...Rule) Validator {
	return Validator{
		rules: defaultRules(rules...),
	}
}

func (v Validator) Validate(evt *Event) bool {
	for _, r := range v.rules {
		if !r(evt) {
			return false
		}
	}

	return true
}

func defaultRules(rules ...Rule) []Rule {
	r := []Rule{
		HasID("event"),
		HasID("tracking"),
	}

	r = append(r, rules...)

	return r
}

func HasContext(ct ContextType) Rule {
	return func(evt *Event) bool {
		_, ok := evt.Context[string(ct)]
		return ok
	}
}

func ContextContains(ct ContextType, k string, nonZero bool) Rule {
	return func(evt *Event) bool {
		c, ok := evt.Context[string(ct)]
		if !ok {
			return false
		}

		v, ok := c.Values()[k]
		if !ok {
			return false
		}

		if nonZero {
			return !isZero(v)
		}

		return true
	}
}

func HasID(id string) Rule {
	return func(evt *Event) bool {
		switch id {
		case "event":
			return len(evt.ID) > 0
		case "tracking":
			return len(evt.TrackingID) > 0
		case "user":
			return len(evt.UserID) > 0
		case "group":
			return len(evt.GroupID) > 0
		case "session":
			return len(evt.SessionID) > 0
		case "device":
			return len(evt.DeviceID) > 0
		default:
			return false
		}
	}
}

func isZero(v interface{}) bool {
	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return rv.Len() == 0
	case reflect.Bool:
		return !rv.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return rv.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return rv.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return rv.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return rv.IsNil()
	case reflect.Struct:
		return structs.IsZero(v)
	case reflect.Invalid:
		return true
	}

	return false
}

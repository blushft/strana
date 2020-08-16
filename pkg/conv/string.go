package conv

import (
	"fmt"
	"strconv"
)

func ToString(v interface{}) string {
	sr, ok := v.(fmt.Stringer)
	if ok {
		return sr.String()
	}

	switch s := v.(type) {
	case string:
		return s
	case int:
		return strconv.Itoa(s)
	}

	return ""
}

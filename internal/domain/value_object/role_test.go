package value_object

import (
	"testing"
)

func TestParseRole(t *testing.T) {
	m := make(map[string]Role)
	m["OWNER"] = OWNER
	m["ADMIN"] = ADMIN
	m["MEMBER"] = MEMBER

	for k, v := range m {
		r, err := ParseRole(k)
		if err != nil || r != v {
			t.Errorf("ParseRole(%s) = %v, want %v", k, r, v)
		}
	}
}

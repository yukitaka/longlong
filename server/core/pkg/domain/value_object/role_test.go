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

func TestRole_IsAbove(t *testing.T) {
	m := make(map[Role]map[Role]bool)
	m[OWNER] = map[Role]bool{OWNER: false, ADMIN: true, MEMBER: true}
	m[ADMIN] = map[Role]bool{OWNER: false, ADMIN: false, MEMBER: true}
	m[MEMBER] = map[Role]bool{OWNER: false, ADMIN: false, MEMBER: false}

	for operator, targetExpects := range m {
		for target, correct := range targetExpects {
			if operator.IsAbove(target) != correct {
				t.Errorf("%s is above %s = %v, want %v", operator, target, operator.IsAbove(target), correct)
			}
		}
	}
}

func TestRole_IsBelow(t *testing.T) {
	m := make(map[Role]map[Role]bool)
	m[OWNER] = map[Role]bool{OWNER: false, ADMIN: false, MEMBER: false}
	m[ADMIN] = map[Role]bool{OWNER: true, ADMIN: false, MEMBER: false}
	m[MEMBER] = map[Role]bool{OWNER: true, ADMIN: true, MEMBER: false}

	for operator, targetExpects := range m {
		for target, correct := range targetExpects {
			if operator.IsBelow(target) != correct {
				t.Errorf("%s is below %s = %v, want %v", operator, target, operator.IsBelow(target), correct)
			}
		}
	}
}

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
	if OWNER.IsAbove(OWNER) {
		t.Errorf("Owner is above owner = %v, want %v", OWNER.IsAbove(OWNER), false)
	}
	if !OWNER.IsAbove(ADMIN) {
		t.Errorf("Owner is above admin = %v, want %v", OWNER.IsAbove(ADMIN), false)
	}
	if !ADMIN.IsAbove(MEMBER) {
		t.Errorf("Admin is above member = %v, want %v", ADMIN.IsAbove(MEMBER), false)
	}
	if MEMBER.IsAbove(MEMBER) {
		t.Errorf("Member is above = %v, want %v", MEMBER.IsAbove(MEMBER), false)
	}
}

func TestRole_IsBelow(t *testing.T) {
	if OWNER.IsBelow(OWNER) {
		t.Errorf("Owner is below owner = %v, want %v", OWNER.IsBelow(OWNER), false)
	}
	if OWNER.IsBelow(ADMIN) {
		t.Errorf("Owner is below admin = %v, want %v", OWNER.IsBelow(ADMIN), false)
	}
	if ADMIN.IsBelow(MEMBER) {
		t.Errorf("Admin is below member = %v, want %v", ADMIN.IsBelow(MEMBER), false)
	}
	if !MEMBER.IsBelow(ADMIN) {
		t.Errorf("Member is below admin = %v, want %v", MEMBER.IsBelow(ADMIN), false)
	}
	if MEMBER.IsBelow(MEMBER) {
		t.Errorf("Member is below = %v, want %v", MEMBER.IsBelow(MEMBER), false)
	}
}

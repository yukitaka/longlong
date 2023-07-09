//go:generate stringer -type=Role
package value_object

import "fmt"

type Role int

const (
	OWNER Role = iota
	ADMIN
	MEMBER
)

var MapEnumStringToUSState = func() map[string]Role {
	m := make(map[string]Role)
	for i := OWNER; i <= MEMBER; i++ {
		m[i.String()] = i
	}
	return m
}()

func ParseRole(name string) (Role, error) {
	if v, ok := MapEnumStringToUSState[name]; ok {
		return v, nil
	}

	return -1, fmt.Errorf("Unparse string '%s' to Role\n", name)
}

func (it Role) IsAbove(target Role) bool {
	return it < target
}

func (it Role) IsBelow(target Role) bool {
	return it > target
}

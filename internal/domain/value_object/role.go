//go:generate stringer -type=Role
package value_object

type Role int

const (
	OWNER Role = iota
	ADMIN
	MEMBER
)

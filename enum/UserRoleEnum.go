package enum

// role enum
type UserRoleEnum int

const (
	Admin UserRoleEnum = iota
	Owner
	User
)

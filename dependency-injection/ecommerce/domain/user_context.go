package domain

type UserContext interface {
	IsInRole(role Role) bool
}

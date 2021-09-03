package domain

type StubUserContext struct {
	roles []Role
}

func (s StubUserContext) IsInRole(role Role) bool {
	for _, r := range s.roles {
		if r == role {
			return true
		}
	}
	return false
}

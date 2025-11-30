package household

import "time"

type ID string

type MemberRole string

const (
	RoleOwner  MemberRole = "owner"
	RoleMember MemberRole = "member"
)

type Household struct {
	ID        ID
	Name      string
	Currency  string // "ARS", "USD", etc.
	CreatedAt time.Time
}

type Member struct {
	HouseholdID ID
	UserID      string // después podemos usar user.ID si queremos
	Role        MemberRole
	CreatedAt   time.Time
}

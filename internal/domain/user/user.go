package user

import "time"

type ID string

type User struct {
	ID        ID
	Email     string
	Name      string
	GoogleSub string // sub que viene de Google / Cognito
	CreatedAt time.Time
	UpdatedAt time.Time
}

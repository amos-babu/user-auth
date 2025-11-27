package types

import "time"

type User struct {
	ID        string
	Email     string
	Name      string
	Age       int
	CreatedAt time.Time
	UpdatedAt time.Time
}

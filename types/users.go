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

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Token string `json:"token"`
}

package models

import (
	"time"
)

type User struct {
	ID           string
	Username     string
	PasswordHash string
	Role         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

const (
	RoleUser  = "user"
	RoleAdmin = "admin"
)

func (u *User) IsAdmin() bool {
	return u.Role == RoleAdmin
}

func (u *User) IsUser() bool {
	return u.Role == RoleUser
}

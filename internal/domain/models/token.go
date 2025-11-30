package models

import (
	"time"
)

type PersonalAccessToken struct {
	ID        string
	UserID    string
	Token     string
	ExpiresAt time.Time
	RevokedAt *time.Time
	CreatedAt time.Time
}

func (t *PersonalAccessToken) IsExpired() bool {
	return time.Now().After(t.ExpiresAt)
}

func (t *PersonalAccessToken) IsRevoked() bool {
	return t.RevokedAt != nil
}

func (t *PersonalAccessToken) IsValid() bool {
	return !t.IsExpired() && !t.IsRevoked()
}


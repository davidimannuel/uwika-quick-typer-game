package models

import (
	"time"
)

type Theme struct {
	ID          string
	Name        string
	Description string
	CreatedAt   time.Time
}


package models

import (
	"time"
)

type Phrase struct {
	ID             string
	StageID        string
	Text           string
	SequenceNumber int
	BaseMultiplier float64
	CreatedAt      time.Time
	UpdatedAt      time.Time
}


package models

import (
	"time"
)

type Score struct {
	UserID      string
	StageID     string
	FinalScore  float64
	TotalTimeMs int
	TotalErrors int
	CompletedAt time.Time
}


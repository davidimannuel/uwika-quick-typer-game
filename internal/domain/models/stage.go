package models

import (
	"time"
)

type Stage struct {
	ID         string
	Name       string
	ThemeID    string
	Difficulty string
	IsActive   bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

const (
	DifficultyEasy   = "easy"
	DifficultyMedium = "medium"
	DifficultyHard   = "hard"
)


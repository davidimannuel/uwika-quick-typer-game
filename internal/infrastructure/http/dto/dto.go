package dto

// Auth DTOs
type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthResponse struct {
	UserID         string `json:"user_id"`
	Username       string `json:"username,omitempty"`
	Role           string `json:"role,omitempty"`
	AccessToken    string `json:"access_token"`
	TokenExpiresAt string `json:"token_expires_at"`
}

// Theme DTOs
type ThemeResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

// Stage DTOs
type CreateStageRequest struct {
	Name       string `json:"name" binding:"required"`
	ThemeID    string `json:"theme_id" binding:"required"`
	Difficulty string `json:"difficulty" binding:"required"`
	IsActive   bool   `json:"is_active"`
}

type UpdateStageRequest struct {
	Name       string `json:"name" binding:"required"`
	ThemeID    string `json:"theme_id" binding:"required"`
	Difficulty string `json:"difficulty" binding:"required"`
	IsActive   bool   `json:"is_active"`
}

type StageResponse struct {
	ID         string           `json:"id"`
	Name       string           `json:"name"`
	ThemeID    string           `json:"theme_id,omitempty"`
	ThemeName  string           `json:"theme_name,omitempty"`
	Difficulty string           `json:"difficulty"`
	IsActive   bool             `json:"is_active"`
	Phrases    []PhraseResponse `json:"phrases,omitempty"`
}

// Phrase DTOs
type CreatePhraseRequest struct {
	StageID        string  `json:"stage_id" binding:"required"`
	Text           string  `json:"text" binding:"required"`
	SequenceNumber int     `json:"sequence_number" binding:"required"`
	BaseMultiplier float64 `json:"base_multiplier" binding:"required"`
}

type UpdatePhraseRequest struct {
	StageID        string  `json:"stage_id" binding:"required"`
	Text           string  `json:"text" binding:"required"`
	SequenceNumber int     `json:"sequence_number" binding:"required"`
	BaseMultiplier float64 `json:"base_multiplier" binding:"required"`
}

type PhraseResponse struct {
	ID             string  `json:"id"`
	StageID        string  `json:"stage_id,omitempty"`
	Text           string  `json:"text"`
	SequenceNumber int     `json:"sequence_number"`
	Multiplier     float64 `json:"multiplier"`
}

// Score DTOs
type SubmitScoreRequest struct {
	StageID     string `json:"stage_id" binding:"required"`
	TotalTimeMs int    `json:"total_time_ms" binding:"required,min=1"`
	TotalErrors int    `json:"total_errors" binding:"min=0"`
}

type SubmitScoreResponse struct {
	Status     string  `json:"status"`
	FinalScore float64 `json:"final_score"`
}

type LeaderboardEntry struct {
	Username    string  `json:"username"`
	FinalScore  float64 `json:"final_score"`
	TotalTimeMs int     `json:"total_time_ms"`
}

// Generic Response
type ErrorResponse struct {
	Error string `json:"error"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

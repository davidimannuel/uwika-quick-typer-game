package services

import (
	"math"
)

// ScoreCalculator adalah domain service untuk menghitung score
type ScoreCalculator struct{}

func NewScoreCalculator() *ScoreCalculator {
	return &ScoreCalculator{}
}

// CalculationInput adalah input untuk perhitungan score
type CalculationInput struct {
	Accuracy       float64
	TypingSpeed    float64
	TimeTaken      float64
	MaxCombo       int
	PerfectPhrases int
	BaseMultiplier float64
}

// CalculationResult adalah hasil perhitungan score
type CalculationResult struct {
	BaseScore       int
	AccuracyBonus   int
	SpeedBonus      int
	ComboBonus      int
	TimeBonus       int
	FinalScore      int
	FinalMultiplier float64
}

// CalculateScore menghitung final score berdasarkan input metrics
func (sc *ScoreCalculator) CalculateScore(input CalculationInput) CalculationResult {
	result := CalculationResult{}

	// 1. Base Score: accuracy × typing speed
	result.BaseScore = int(math.Round(input.Accuracy * input.TypingSpeed))

	// 2. Accuracy Bonus: perfect accuracy gives extra points
	if input.Accuracy >= 100 {
		result.AccuracyBonus = 500
	} else if input.Accuracy >= 95 {
		result.AccuracyBonus = 200
	} else if input.Accuracy >= 90 {
		result.AccuracyBonus = 100
	}

	// 3. Speed Bonus: fast typing (> 60 WPM) gets bonus
	if input.TypingSpeed >= 100 {
		result.SpeedBonus = 300
	} else if input.TypingSpeed >= 80 {
		result.SpeedBonus = 200
	} else if input.TypingSpeed >= 60 {
		result.SpeedBonus = 100
	}

	// 4. Combo Bonus: reward high combos
	result.ComboBonus = input.MaxCombo * 10

	// 5. Time Bonus: faster completion
	// Assuming target time is calculated based on phrase complexity
	// For now, simple: less time = more bonus
	if input.TimeTaken > 0 {
		timeBonus := 1000 / input.TimeTaken // Inverse relationship
		result.TimeBonus = int(math.Round(timeBonus * 100))
		if result.TimeBonus > 500 {
			result.TimeBonus = 500 // Cap at 500
		}
	}

	// 6. Calculate total before multiplier
	totalBeforeMultiplier := result.BaseScore +
		result.AccuracyBonus +
		result.SpeedBonus +
		result.ComboBonus +
		result.TimeBonus

	// 7. Apply base multiplier
	result.FinalMultiplier = input.BaseMultiplier
	result.FinalScore = int(math.Round(float64(totalBeforeMultiplier) * result.FinalMultiplier))

	// Ensure minimum score of 0
	if result.FinalScore < 0 {
		result.FinalScore = 0
	}

	return result
}

// ValidateMetrics validates that the metrics are within reasonable bounds
func (sc *ScoreCalculator) ValidateMetrics(accuracy, typingSpeed, timeTaken float64) error {
	if accuracy < 0 || accuracy > 100 {
		return ErrInvalidAccuracy
	}

	if typingSpeed < 0 || typingSpeed > 300 { // Max 300 WPM is reasonable
		return ErrInvalidTypingSpeed
	}

	if timeTaken < 0 {
		return ErrInvalidTimeTaken
	}

	return nil
}

// CalculateTimeBonus calculates bonus based on completion time
func (sc *ScoreCalculator) CalculateTimeBonus(timeTaken, targetTime float64) int {
	if timeTaken <= 0 || targetTime <= 0 {
		return 0
	}

	ratio := targetTime / timeTaken
	if ratio >= 1.5 {
		return 500
	} else if ratio >= 1.2 {
		return 300
	} else if ratio >= 1.0 {
		return 100
	}

	return 0
}

// CalculateStars determines star rating (1-3) based on accuracy
func (sc *ScoreCalculator) CalculateStars(accuracy float64) int {
	if accuracy >= 95 {
		return 3
	} else if accuracy >= 80 {
		return 2
	}
	return 1
}

// Domain errors
var (
	ErrInvalidAccuracy    = &DomainError{Code: "INVALID_ACCURACY", Message: "accuracy must be between 0 and 100"}
	ErrInvalidTypingSpeed = &DomainError{Code: "INVALID_TYPING_SPEED", Message: "typing speed is out of reasonable range"}
	ErrInvalidTimeTaken   = &DomainError{Code: "INVALID_TIME_TAKEN", Message: "time taken must be positive"}
)

type DomainError struct {
	Code    string
	Message string
}

func (e *DomainError) Error() string {
	return e.Message
}

package services

import (
	"context"
	"errors"

	"uwika_quick_typer_game/internal/domain/models"
	"uwika_quick_typer_game/internal/domain/repositories"
	domainservices "uwika_quick_typer_game/internal/domain/services"
)

var (
	ErrStageNotFound = errors.New("stage not found")
)

type GameService struct {
	stageRepo       repositories.StageRepository
	phraseRepo      repositories.PhraseRepository
	scoreRepo       repositories.ScoreRepository
	scoreCalculator *domainservices.ScoreCalculator
}

func NewGameService(
	stageRepo repositories.StageRepository,
	phraseRepo repositories.PhraseRepository,
	scoreRepo repositories.ScoreRepository,
) *GameService {
	return &GameService{
		stageRepo:       stageRepo,
		phraseRepo:      phraseRepo,
		scoreRepo:       scoreRepo,
		scoreCalculator: domainservices.NewScoreCalculator(),
	}
}

func (s *GameService) GetActiveStages(ctx context.Context) ([]*models.Stage, error) {
	return s.stageRepo.FindAllActive(ctx)
}

func (s *GameService) GetStageWithPhrases(ctx context.Context, stageID string) (*models.Stage, []*models.Phrase, error) {
	stage, err := s.stageRepo.FindByID(ctx, stageID)
	if err != nil {
		return nil, nil, err
	}
	if stage == nil {
		return nil, nil, ErrStageNotFound
	}

	phrases, err := s.phraseRepo.FindByStageID(ctx, stageID)
	if err != nil {
		return nil, nil, err
	}

	return stage, phrases, nil
}

// SubmitScore - calculation dilakukan di domain service
func (s *GameService) SubmitScore(ctx context.Context, userID, stageID string, totalTimeMs, totalErrors int) (*models.Score, string, error) {
	// Get stage and phrases
	stage, phrases, err := s.GetStageWithPhrases(ctx, stageID)
	if err != nil {
		return nil, "", err
	}
	if stage == nil {
		return nil, "", ErrStageNotFound
	}

	// Calculate metrics for domain service
	totalChars := 0
	totalMultiplier := 0.0
	for _, phrase := range phrases {
		totalChars += len(phrase.Text)
		totalMultiplier += phrase.BaseMultiplier
	}

	avgMultiplier := 1.0
	if len(phrases) > 0 {
		avgMultiplier = totalMultiplier / float64(len(phrases))
	}

	// Simple accuracy calculation based on errors
	accuracy := 100.0
	if totalChars > 0 && totalErrors > 0 {
		accuracy = float64(totalChars-totalErrors) / float64(totalChars) * 100
		if accuracy < 0 {
			accuracy = 0
		}
	}

	// Calculate typing speed (WPM)
	timeTakenSeconds := float64(totalTimeMs) / 1000.0
	if timeTakenSeconds == 0 {
		timeTakenSeconds = 0.001
	}
	typingSpeed := (float64(totalChars) / timeTakenSeconds) * 60.0 / 5.0 // WPM

	// Use domain service untuk calculate score
	calcInput := domainservices.CalculationInput{
		Accuracy:       accuracy,
		TypingSpeed:    typingSpeed,
		TimeTaken:      timeTakenSeconds,
		MaxCombo:       0,
		PerfectPhrases: 0,
		BaseMultiplier: avgMultiplier,
	}

	// Validate metrics
	if err := s.scoreCalculator.ValidateMetrics(accuracy, typingSpeed, timeTakenSeconds); err != nil {
		return nil, "", err
	}

	// Calculate final score using domain service
	calcResult := s.scoreCalculator.CalculateScore(calcInput)

	score := &models.Score{
		UserID:      userID,
		StageID:     stageID,
		FinalScore:  float64(calcResult.FinalScore),
		TotalTimeMs: totalTimeMs,
		TotalErrors: totalErrors,
	}

	// Allow multiple attempts - always insert
	err = s.scoreRepo.Create(ctx, score)
	if err != nil {
		return nil, "", err
	}

	return score, "INSERTED", nil
}

func (s *GameService) GetLeaderboard(ctx context.Context, stageID string, limit int) ([]*models.Score, error) {
	if limit <= 0 {
		limit = 20
	}
	return s.scoreRepo.FindLeaderboardByStage(ctx, stageID, limit)
}

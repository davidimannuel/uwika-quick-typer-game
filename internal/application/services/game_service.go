package services

import (
	"context"
	"errors"
	"math"

	"uwika_quick_typer_game/internal/domain/models"
	"uwika_quick_typer_game/internal/domain/repositories"
)

var (
	ErrStageNotFound = errors.New("stage not found")
)

const (
	ErrorPenaltyPerMistake = 50.0
)

type GameService struct {
	stageRepo  repositories.StageRepository
	phraseRepo repositories.PhraseRepository
	scoreRepo  repositories.ScoreRepository
}

func NewGameService(
	stageRepo repositories.StageRepository,
	phraseRepo repositories.PhraseRepository,
	scoreRepo repositories.ScoreRepository,
) *GameService {
	return &GameService{
		stageRepo:  stageRepo,
		phraseRepo: phraseRepo,
		scoreRepo:  scoreRepo,
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

func (s *GameService) SubmitScore(ctx context.Context, userID, stageID string, totalTimeMs, totalErrors int) (*models.Score, string, error) {
	// Get stage and phrases to calculate score
	stage, phrases, err := s.GetStageWithPhrases(ctx, stageID)
	if err != nil {
		return nil, "", err
	}
	if stage == nil {
		return nil, "", ErrStageNotFound
	}

	// Calculate final score
	finalScore := s.calculateScore(phrases, totalTimeMs, totalErrors)

	score := &models.Score{
		UserID:      userID,
		StageID:     stageID,
		FinalScore:  finalScore,
		TotalTimeMs: totalTimeMs,
		TotalErrors: totalErrors,
	}

	isInserted, err := s.scoreRepo.Upsert(ctx, score)
	if err != nil {
		return nil, "", err
	}

	status := "IGNORED"
	if isInserted {
		status = "UPSERTED"
	} else {
		// Check if it's actually updated (better score)
		existingScore, _ := s.scoreRepo.FindByUserAndStage(ctx, userID, stageID)
		if existingScore != nil && existingScore.FinalScore == finalScore {
			status = "UPSERTED"
		}
	}

	return score, status, nil
}

func (s *GameService) GetLeaderboard(ctx context.Context, stageID string, limit int) ([]*models.Score, error) {
	if limit <= 0 {
		limit = 20
	}
	return s.scoreRepo.FindLeaderboardByStage(ctx, stageID, limit)
}

func (s *GameService) calculateScore(phrases []*models.Phrase, totalTimeMs, totalErrors int) float64 {
	// Calculate: (sum(phraseLength * multiplier) / totalTimeInSeconds) - errorPenalty
	totalTimeSeconds := float64(totalTimeMs) / 1000.0
	if totalTimeSeconds == 0 {
		totalTimeSeconds = 0.001 // Prevent division by zero
	}

	var scoreSum float64
	for _, phrase := range phrases {
		phraseLength := float64(len(phrase.Text))
		scoreSum += phraseLength * phrase.BaseMultiplier
	}

	scorePerSecond := scoreSum / totalTimeSeconds
	errorPenalty := float64(totalErrors) * ErrorPenaltyPerMistake

	finalScore := scorePerSecond - errorPenalty
	if finalScore < 0 {
		finalScore = 0
	}

	return math.Round(finalScore*100) / 100 // Round to 2 decimal places
}


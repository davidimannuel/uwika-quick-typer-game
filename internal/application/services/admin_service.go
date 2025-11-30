package services

import (
	"context"
	"errors"

	"uwika_quick_typer_game/internal/domain/models"
	"uwika_quick_typer_game/internal/domain/repositories"
)

type AdminService struct {
	stageRepo  repositories.StageRepository
	phraseRepo repositories.PhraseRepository
	userRepo   repositories.UserRepository
	themeRepo  repositories.ThemeRepository
}

func NewAdminService(
	stageRepo repositories.StageRepository,
	phraseRepo repositories.PhraseRepository,
	userRepo repositories.UserRepository,
	themeRepo repositories.ThemeRepository,
) *AdminService {
	return &AdminService{
		stageRepo:  stageRepo,
		phraseRepo: phraseRepo,
		userRepo:   userRepo,
		themeRepo:  themeRepo,
	}
}

// Theme Management
func (s *AdminService) GetAllThemes(ctx context.Context) ([]*models.Theme, error) {
	return s.themeRepo.FindAll(ctx)
}

// Stage Management
func (s *AdminService) CreateStage(ctx context.Context, name, themeID, difficulty string, isActive bool) (*models.Stage, error) {
	stage := &models.Stage{
		Name:       name,
		ThemeID:    themeID,
		Difficulty: difficulty,
		IsActive:   isActive,
	}
	err := s.stageRepo.Create(ctx, stage)
	if err != nil {
		return nil, err
	}
	return stage, nil
}

func (s *AdminService) UpdateStage(ctx context.Context, stageID, name, themeID, difficulty string, isActive bool) (*models.Stage, error) {
	stage, err := s.stageRepo.FindByID(ctx, stageID)
	if err != nil {
		return nil, err
	}
	if stage == nil {
		return nil, ErrStageNotFound
	}

	stage.Name = name
	stage.ThemeID = themeID
	stage.Difficulty = difficulty
	stage.IsActive = isActive

	err = s.stageRepo.Update(ctx, stage)
	if err != nil {
		return nil, err
	}
	return stage, nil
}

func (s *AdminService) DeleteStage(ctx context.Context, stageID string) error {
	return s.stageRepo.Delete(ctx, stageID)
}

func (s *AdminService) GetAllStages(ctx context.Context) ([]*models.Stage, error) {
	return s.stageRepo.FindAll(ctx)
}

// Phrase Management
func (s *AdminService) CreatePhrase(ctx context.Context, stageID, text string, sequenceNumber int, baseMultiplier float64) (*models.Phrase, error) {
	phrase := &models.Phrase{
		StageID:        stageID,
		Text:           text,
		SequenceNumber: sequenceNumber,
		BaseMultiplier: baseMultiplier,
	}
	err := s.phraseRepo.Create(ctx, phrase)
	if err != nil {
		return nil, err
	}
	return phrase, nil
}

func (s *AdminService) UpdatePhrase(ctx context.Context, phraseID, stageID, text string, sequenceNumber int, baseMultiplier float64) (*models.Phrase, error) {
	phrase, err := s.phraseRepo.FindByID(ctx, phraseID)
	if err != nil {
		return nil, err
	}
	if phrase == nil {
		return nil, errors.New("phrase not found")
	}

	phrase.StageID = stageID
	phrase.Text = text
	phrase.SequenceNumber = sequenceNumber
	phrase.BaseMultiplier = baseMultiplier

	err = s.phraseRepo.Update(ctx, phrase)
	if err != nil {
		return nil, err
	}
	return phrase, nil
}

func (s *AdminService) DeletePhrase(ctx context.Context, phraseID string) error {
	return s.phraseRepo.Delete(ctx, phraseID)
}

func (s *AdminService) GetPhrasesByStage(ctx context.Context, stageID string) ([]*models.Phrase, error) {
	return s.phraseRepo.FindByStageID(ctx, stageID)
}

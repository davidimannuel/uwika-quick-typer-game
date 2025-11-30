package repositories

import (
	"context"
	"uwika_quick_typer_game/internal/domain/models"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	FindByID(ctx context.Context, userID string) (*models.User, error)
	FindByUsername(ctx context.Context, username string) (*models.User, error)
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, userID string) error
}

type TokenRepository interface {
	Create(ctx context.Context, token *models.PersonalAccessToken) error
	FindByToken(ctx context.Context, tokenHash string) (*models.PersonalAccessToken, error)
	RevokeToken(ctx context.Context, tokenHash string) error
	RevokeAllUserTokens(ctx context.Context, userID string) error
	DeleteExpiredTokens(ctx context.Context) error
}

type ThemeRepository interface {
	FindAll(ctx context.Context) ([]*models.Theme, error)
	FindByID(ctx context.Context, themeID string) (*models.Theme, error)
}

type StageRepository interface {
	Create(ctx context.Context, stage *models.Stage) error
	FindByID(ctx context.Context, stageID string) (*models.Stage, error)
	FindAll(ctx context.Context) ([]*models.Stage, error)
	FindAllActive(ctx context.Context) ([]*models.Stage, error)
	Update(ctx context.Context, stage *models.Stage) error
	Delete(ctx context.Context, stageID string) error
}

type PhraseRepository interface {
	Create(ctx context.Context, phrase *models.Phrase) error
	FindByID(ctx context.Context, phraseID string) (*models.Phrase, error)
	FindByStageID(ctx context.Context, stageID string) ([]*models.Phrase, error)
	Update(ctx context.Context, phrase *models.Phrase) error
	Delete(ctx context.Context, phraseID string) error
}

type ScoreRepository interface {
	Create(ctx context.Context, score *models.Score) error
	FindByUserAndStage(ctx context.Context, userID, stageID string) (*models.Score, error)
	FindLeaderboardByStage(ctx context.Context, stageID string, limit int) ([]*models.Score, error)
	FindByUserID(ctx context.Context, userID string) ([]*models.Score, error)
}


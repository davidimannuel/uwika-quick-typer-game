package postgres

import (
	"context"
	"database/sql"
	"time"

	"uwika_quick_typer_game/internal/domain/models"
	"uwika_quick_typer_game/internal/domain/repositories"

	"github.com/google/uuid"
)

type tokenRepository struct {
	db *sql.DB
}

func NewTokenRepository(db *sql.DB) repositories.TokenRepository {
	return &tokenRepository{db: db}
}

func (r *tokenRepository) Create(ctx context.Context, token *models.PersonalAccessToken) error {
	if token.ID == "" {
		token.ID = uuid.New().String()
	}
	token.CreatedAt = time.Now()

	query := `
		INSERT INTO personal_access_tokens (id, user_id, token, expires_at, revoked_at, created_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err := r.db.ExecContext(ctx, query,
		token.ID, token.UserID, token.Token, token.ExpiresAt, token.RevokedAt, token.CreatedAt,
	)
	return err
}

func (r *tokenRepository) FindByToken(ctx context.Context, tokenHash string) (*models.PersonalAccessToken, error) {
	query := `
		SELECT id, user_id, token, expires_at, revoked_at, created_at
		FROM personal_access_tokens 
		WHERE token = $1
	`
	token := &models.PersonalAccessToken{}
	err := r.db.QueryRowContext(ctx, query, tokenHash).Scan(
		&token.ID, &token.UserID, &token.Token, &token.ExpiresAt, &token.RevokedAt, &token.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (r *tokenRepository) RevokeToken(ctx context.Context, tokenHash string) error {
	query := `
		UPDATE personal_access_tokens 
		SET revoked_at = $2
		WHERE token = $1 AND revoked_at IS NULL
	`
	_, err := r.db.ExecContext(ctx, query, tokenHash, time.Now())
	return err
}

func (r *tokenRepository) RevokeAllUserTokens(ctx context.Context, userID string) error {
	query := `
		UPDATE personal_access_tokens 
		SET revoked_at = $2
		WHERE user_id = $1 AND revoked_at IS NULL
	`
	_, err := r.db.ExecContext(ctx, query, userID, time.Now())
	return err
}

func (r *tokenRepository) DeleteExpiredTokens(ctx context.Context) error {
	query := `
		DELETE FROM personal_access_tokens 
		WHERE expires_at < $1 OR revoked_at IS NOT NULL
	`
	_, err := r.db.ExecContext(ctx, query, time.Now())
	return err
}


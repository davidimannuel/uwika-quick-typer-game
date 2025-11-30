package postgres

import (
	"context"
	"database/sql"
	"time"

	"uwika_quick_typer_game/internal/domain/models"
	"uwika_quick_typer_game/internal/domain/repositories"

	"github.com/google/uuid"
)

type phraseRepository struct {
	db *sql.DB
}

func NewPhraseRepository(db *sql.DB) repositories.PhraseRepository {
	return &phraseRepository{db: db}
}

func (r *phraseRepository) Create(ctx context.Context, phrase *models.Phrase) error {
	if phrase.ID == "" {
		phrase.ID = uuid.New().String()
	}
	phrase.CreatedAt = time.Now()
	phrase.UpdatedAt = time.Now()

	query := `
		INSERT INTO phrases (id, stage_id, text, sequence_number, base_multiplier, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	_, err := r.db.ExecContext(ctx, query,
		phrase.ID, phrase.StageID, phrase.Text, phrase.SequenceNumber, phrase.BaseMultiplier, phrase.CreatedAt, phrase.UpdatedAt,
	)
	return err
}

func (r *phraseRepository) FindByID(ctx context.Context, phraseID string) (*models.Phrase, error) {
	query := `
		SELECT id, stage_id, text, sequence_number, base_multiplier, created_at, updated_at
		FROM phrases WHERE id = $1
	`
	phrase := &models.Phrase{}
	err := r.db.QueryRowContext(ctx, query, phraseID).Scan(
		&phrase.ID, &phrase.StageID, &phrase.Text, &phrase.SequenceNumber, &phrase.BaseMultiplier, &phrase.CreatedAt, &phrase.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return phrase, nil
}

func (r *phraseRepository) FindByStageID(ctx context.Context, stageID string) ([]*models.Phrase, error) {
	query := `
		SELECT id, stage_id, text, sequence_number, base_multiplier, created_at, updated_at
		FROM phrases 
		WHERE stage_id = $1
		ORDER BY sequence_number ASC
	`
	rows, err := r.db.QueryContext(ctx, query, stageID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var phrases []*models.Phrase
	for rows.Next() {
		phrase := &models.Phrase{}
		err := rows.Scan(
			&phrase.ID, &phrase.StageID, &phrase.Text, &phrase.SequenceNumber, &phrase.BaseMultiplier, &phrase.CreatedAt, &phrase.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		phrases = append(phrases, phrase)
	}
	return phrases, nil
}

func (r *phraseRepository) Update(ctx context.Context, phrase *models.Phrase) error {
	phrase.UpdatedAt = time.Now()
	query := `
		UPDATE phrases 
		SET stage_id = $2, text = $3, sequence_number = $4, base_multiplier = $5, updated_at = $6
		WHERE id = $1
	`
	_, err := r.db.ExecContext(ctx, query,
		phrase.ID, phrase.StageID, phrase.Text, phrase.SequenceNumber, phrase.BaseMultiplier, phrase.UpdatedAt,
	)
	return err
}

func (r *phraseRepository) Delete(ctx context.Context, phraseID string) error {
	query := `DELETE FROM phrases WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, phraseID)
	return err
}

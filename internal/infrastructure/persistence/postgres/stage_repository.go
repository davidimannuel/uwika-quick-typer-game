package postgres

import (
	"context"
	"database/sql"
	"time"

	"uwika_quick_typer_game/internal/domain/models"
	"uwika_quick_typer_game/internal/domain/repositories"

	"github.com/google/uuid"
)

type stageRepository struct {
	db *sql.DB
}

func NewStageRepository(db *sql.DB) repositories.StageRepository {
	return &stageRepository{db: db}
}

func (r *stageRepository) Create(ctx context.Context, stage *models.Stage) error {
	if stage.ID == "" {
		stage.ID = uuid.New().String()
	}
	stage.CreatedAt = time.Now()
	stage.UpdatedAt = time.Now()

	query := `
		INSERT INTO stages (id, name, theme_id, difficulty, is_active, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	_, err := r.db.ExecContext(ctx, query,
		stage.ID, stage.Name, stage.ThemeID, stage.Difficulty, stage.IsActive, stage.CreatedAt, stage.UpdatedAt,
	)
	return err
}

func (r *stageRepository) FindByID(ctx context.Context, stageID string) (*models.Stage, error) {
	query := `
		SELECT id, name, theme_id, difficulty, is_active, created_at, updated_at
		FROM stages WHERE id = $1
	`
	stage := &models.Stage{}
	err := r.db.QueryRowContext(ctx, query, stageID).Scan(
		&stage.ID, &stage.Name, &stage.ThemeID, &stage.Difficulty, &stage.IsActive, &stage.CreatedAt, &stage.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return stage, nil
}

func (r *stageRepository) FindAll(ctx context.Context) ([]*models.Stage, error) {
	query := `
		SELECT id, name, theme_id, difficulty, is_active, created_at, updated_at
		FROM stages
		ORDER BY created_at DESC
	`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stages []*models.Stage
	for rows.Next() {
		stage := &models.Stage{}
		err := rows.Scan(
			&stage.ID, &stage.Name, &stage.ThemeID, &stage.Difficulty, &stage.IsActive, &stage.CreatedAt, &stage.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		stages = append(stages, stage)
	}
	return stages, nil
}

func (r *stageRepository) FindAllActive(ctx context.Context) ([]*models.Stage, error) {
	query := `
		SELECT id, name, theme_id, difficulty, is_active, created_at, updated_at
		FROM stages
		WHERE is_active = true
		ORDER BY created_at DESC
	`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stages []*models.Stage
	for rows.Next() {
		stage := &models.Stage{}
		err := rows.Scan(
			&stage.ID, &stage.Name, &stage.ThemeID, &stage.Difficulty, &stage.IsActive, &stage.CreatedAt, &stage.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		stages = append(stages, stage)
	}
	return stages, nil
}

func (r *stageRepository) Update(ctx context.Context, stage *models.Stage) error {
	stage.UpdatedAt = time.Now()
	query := `
		UPDATE stages 
		SET name = $2, theme_id = $3, difficulty = $4, is_active = $5, updated_at = $6
		WHERE id = $1
	`
	_, err := r.db.ExecContext(ctx, query,
		stage.ID, stage.Name, stage.ThemeID, stage.Difficulty, stage.IsActive, stage.UpdatedAt,
	)
	return err
}

func (r *stageRepository) Delete(ctx context.Context, stageID string) error {
	query := `DELETE FROM stages WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, stageID)
	return err
}

package postgres

import (
	"context"
	"database/sql"
	"time"

	"uwika_quick_typer_game/internal/domain/models"
	"uwika_quick_typer_game/internal/domain/repositories"
)

type scoreRepository struct {
	db *sql.DB
}

func NewScoreRepository(db *sql.DB) repositories.ScoreRepository {
	return &scoreRepository{db: db}
}

func (r *scoreRepository) Create(ctx context.Context, score *models.Score) error {
	score.CompletedAt = time.Now()

	query := `
		INSERT INTO scores (user_id, stage_id, final_score, total_time_ms, total_errors, completed_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err := r.db.ExecContext(ctx, query,
		score.UserID, score.StageID, score.FinalScore, score.TotalTimeMs, score.TotalErrors, score.CompletedAt,
	)

	return err
}

func (r *scoreRepository) FindByUserAndStage(ctx context.Context, userID, stageID string) (*models.Score, error) {
	// Get best score for this user on this stage
	query := `
		SELECT user_id, stage_id, final_score, total_time_ms, total_errors, completed_at
		FROM scores 
		WHERE user_id = $1 AND stage_id = $2
		ORDER BY final_score DESC, total_time_ms ASC
		LIMIT 1
	`
	score := &models.Score{}
	err := r.db.QueryRowContext(ctx, query, userID, stageID).Scan(
		&score.UserID, &score.StageID, &score.FinalScore, &score.TotalTimeMs, &score.TotalErrors, &score.CompletedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return score, nil
}

func (r *scoreRepository) FindLeaderboardByStage(ctx context.Context, stageID string, limit int) ([]*models.Score, error) {
	// Get best score per user for the leaderboard
	query := `
		WITH best_scores AS (
			SELECT DISTINCT ON (user_id)
				user_id, stage_id, final_score, total_time_ms, total_errors, completed_at
			FROM scores 
			WHERE stage_id = $1
			ORDER BY user_id, final_score DESC, total_time_ms ASC
		)
		SELECT user_id, stage_id, final_score, total_time_ms, total_errors, completed_at
		FROM best_scores
		ORDER BY final_score DESC, total_time_ms ASC
		LIMIT $2
	`
	rows, err := r.db.QueryContext(ctx, query, stageID, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var scores []*models.Score
	for rows.Next() {
		score := &models.Score{}
		err := rows.Scan(
			&score.UserID, &score.StageID, &score.FinalScore, &score.TotalTimeMs, &score.TotalErrors, &score.CompletedAt,
		)
		if err != nil {
			return nil, err
		}
		scores = append(scores, score)
	}
	return scores, nil
}

func (r *scoreRepository) FindByUserID(ctx context.Context, userID string) ([]*models.Score, error) {
	query := `
		SELECT user_id, stage_id, final_score, total_time_ms, total_errors, completed_at
		FROM scores 
		WHERE user_id = $1
		ORDER BY completed_at DESC
	`
	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var scores []*models.Score
	for rows.Next() {
		score := &models.Score{}
		err := rows.Scan(
			&score.UserID, &score.StageID, &score.FinalScore, &score.TotalTimeMs, &score.TotalErrors, &score.CompletedAt,
		)
		if err != nil {
			return nil, err
		}
		scores = append(scores, score)
	}
	return scores, nil
}

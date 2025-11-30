package postgres

import (
	"context"
	"database/sql"

	"uwika_quick_typer_game/internal/domain/models"
	"uwika_quick_typer_game/internal/domain/repositories"
)

type themeRepository struct {
	db *sql.DB
}

func NewThemeRepository(db *sql.DB) repositories.ThemeRepository {
	return &themeRepository{db: db}
}

func (r *themeRepository) FindAll(ctx context.Context) ([]*models.Theme, error) {
	query := `
		SELECT id, name, description, created_at
		FROM themes
		ORDER BY name ASC
	`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var themes []*models.Theme
	for rows.Next() {
		theme := &models.Theme{}
		var description sql.NullString
		err := rows.Scan(&theme.ID, &theme.Name, &description, &theme.CreatedAt)
		if err != nil {
			return nil, err
		}
		if description.Valid {
			theme.Description = description.String
		}
		themes = append(themes, theme)
	}
	return themes, nil
}

func (r *themeRepository) FindByID(ctx context.Context, themeID string) (*models.Theme, error) {
	query := `
		SELECT id, name, description, created_at
		FROM themes
		WHERE id = $1
	`
	theme := &models.Theme{}
	var description sql.NullString
	err := r.db.QueryRowContext(ctx, query, themeID).Scan(
		&theme.ID, &theme.Name, &description, &theme.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	if description.Valid {
		theme.Description = description.String
	}
	return theme, nil
}


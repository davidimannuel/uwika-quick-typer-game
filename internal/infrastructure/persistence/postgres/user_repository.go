package postgres

import (
	"context"
	"database/sql"
	"time"

	"uwika_quick_typer_game/internal/domain/models"
	"uwika_quick_typer_game/internal/domain/repositories"

	"github.com/google/uuid"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repositories.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, user *models.User) error {
	if user.ID == "" {
		user.ID = uuid.New().String()
	}
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	query := `
		INSERT INTO users (id, username, password_hash, role, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err := r.db.ExecContext(ctx, query,
		user.ID, user.Username, user.PasswordHash, user.Role, user.CreatedAt, user.UpdatedAt,
	)
	return err
}

func (r *userRepository) FindByID(ctx context.Context, userID string) (*models.User, error) {
	query := `
		SELECT id, username, password_hash, role, created_at, updated_at
		FROM users WHERE id = $1
	`
	user := &models.User{}
	err := r.db.QueryRowContext(ctx, query, userID).Scan(
		&user.ID, &user.Username, &user.PasswordHash, &user.Role, &user.CreatedAt, &user.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) FindByUsername(ctx context.Context, username string) (*models.User, error) {
	query := `
		SELECT id, username, password_hash, role, created_at, updated_at
		FROM users WHERE username = $1
	`
	user := &models.User{}
	err := r.db.QueryRowContext(ctx, query, username).Scan(
		&user.ID, &user.Username, &user.PasswordHash, &user.Role, &user.CreatedAt, &user.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) Update(ctx context.Context, user *models.User) error {
	user.UpdatedAt = time.Now()
	query := `
		UPDATE users 
		SET username = $2, password_hash = $3, role = $4, updated_at = $5
		WHERE id = $1
	`
	_, err := r.db.ExecContext(ctx, query,
		user.ID, user.Username, user.PasswordHash, user.Role, user.UpdatedAt,
	)
	return err
}

func (r *userRepository) Delete(ctx context.Context, userID string) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, userID)
	return err
}


package services

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"time"

	"uwika_quick_typer_game/internal/domain/models"
	"uwika_quick_typer_game/internal/domain/repositories"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserAlreadyExists  = errors.New("user already exists")
	ErrInvalidToken       = errors.New("invalid or expired token")
	ErrUnauthorized       = errors.New("unauthorized")
)

type AuthService struct {
	userRepo  repositories.UserRepository
	tokenRepo repositories.TokenRepository
}

func NewAuthService(userRepo repositories.UserRepository, tokenRepo repositories.TokenRepository) *AuthService {
	return &AuthService{
		userRepo:  userRepo,
		tokenRepo: tokenRepo,
	}
}

func (s *AuthService) Register(ctx context.Context, username, password string) (*models.User, string, time.Time, error) {
	// Check if user exists
	existingUser, err := s.userRepo.FindByUsername(ctx, username)
	if err != nil {
		return nil, "", time.Time{}, err
	}
	if existingUser != nil {
		return nil, "", time.Time{}, ErrUserAlreadyExists
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, "", time.Time{}, err
	}

	// Create user
	user := &models.User{
		ID:           uuid.New().String(),
		Username:     username,
		PasswordHash: string(hashedPassword),
		Role:         models.RoleUser,
	}

	err = s.userRepo.Create(ctx, user)
	if err != nil {
		return nil, "", time.Time{}, err
	}

	// Generate token
	token, tokenHash, expiresAt := s.generateToken()

	// Save token
	personalToken := &models.PersonalAccessToken{
		ID:        uuid.New().String(),
		UserID:    user.ID,
		Token:     tokenHash,
		ExpiresAt: expiresAt,
	}

	err = s.tokenRepo.Create(ctx, personalToken)
	if err != nil {
		return nil, "", time.Time{}, err
	}

	return user, token, expiresAt, nil
}

func (s *AuthService) Login(ctx context.Context, username, password string) (*models.User, string, time.Time, error) {
	// Find user
	user, err := s.userRepo.FindByUsername(ctx, username)
	if err != nil {
		return nil, "", time.Time{}, err
	}
	if user == nil {
		return nil, "", time.Time{}, ErrInvalidCredentials
	}

	// Verify password
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, "", time.Time{}, ErrInvalidCredentials
	}

	// Generate token
	token, tokenHash, expiresAt := s.generateToken()

	// Save token
	personalToken := &models.PersonalAccessToken{
		ID:        uuid.New().String(),
		UserID:    user.ID,
		Token:     tokenHash,
		ExpiresAt: expiresAt,
	}

	err = s.tokenRepo.Create(ctx, personalToken)
	if err != nil {
		return nil, "", time.Time{}, err
	}

	return user, token, expiresAt, nil
}

func (s *AuthService) ValidateToken(ctx context.Context, token string) (*models.User, error) {
	// Hash token
	tokenHash := hashToken(token)

	// Find token
	personalToken, err := s.tokenRepo.FindByToken(ctx, tokenHash)
	if err != nil {
		return nil, err
	}
	if personalToken == nil || !personalToken.IsValid() {
		return nil, ErrInvalidToken
	}

	// Find user
	user, err := s.userRepo.FindByID(ctx, personalToken.UserID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrInvalidToken
	}

	return user, nil
}

func (s *AuthService) Logout(ctx context.Context, token string) error {
	tokenHash := hashToken(token)
	return s.tokenRepo.RevokeToken(ctx, tokenHash)
}

func (s *AuthService) generateToken() (token, tokenHash string, expiresAt time.Time) {
	// Generate random token
	token = uuid.New().String() + uuid.New().String()
	tokenHash = hashToken(token)
	expiresAt = time.Now().Add(30 * 24 * time.Hour) // 30 days
	return
}

func hashToken(token string) string {
	hash := sha256.Sum256([]byte(token))
	return hex.EncodeToString(hash[:])
}


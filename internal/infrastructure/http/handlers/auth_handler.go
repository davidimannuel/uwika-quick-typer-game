package handlers

import (
	"net/http"

	"uwika_quick_typer_game/internal/application/services"
	"uwika_quick_typer_game/internal/domain/models"
	"uwika_quick_typer_game/internal/infrastructure/http/dto"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	user, token, expiresAt, err := h.authService.Register(c.Request.Context(), req.Username, req.Password)
	if err != nil {
		if err == services.ErrUserAlreadyExists {
			c.JSON(http.StatusConflict, dto.ErrorResponse{Error: "username already exists"})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dto.AuthResponse{
		UserID:         user.ID,
		Username:       user.Username,
		Role:           user.Role,
		AccessToken:    token,
		TokenExpiresAt: expiresAt.Format("2006-01-02T15:04:05Z07:00"),
	})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	user, token, expiresAt, err := h.authService.Login(c.Request.Context(), req.Username, req.Password)
	if err != nil {
		if err == services.ErrInvalidCredentials {
			c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Error: "invalid credentials"})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.AuthResponse{
		UserID:         user.ID,
		AccessToken:    token,
		TokenExpiresAt: expiresAt.Format("2006-01-02T15:04:05Z07:00"),
	})
}

func (h *AuthHandler) Profile(c *gin.Context) {
	user := getUserFromContext(c)
	if user == nil {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Error: "unauthorized"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user_id":  user.ID,
		"username": user.Username,
		"role":     user.Role,
	})
}

func getUserFromContext(c *gin.Context) *models.User {
	userInterface, exists := c.Get("user")
	if !exists {
		return nil
	}
	user, _ := userInterface.(*models.User)
	return user
}


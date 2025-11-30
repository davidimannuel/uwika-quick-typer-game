package handlers

import (
	"net/http"
	"strconv"

	"uwika_quick_typer_game/internal/application/services"
	"uwika_quick_typer_game/internal/domain/repositories"
	"uwika_quick_typer_game/internal/infrastructure/http/dto"
	"uwika_quick_typer_game/internal/infrastructure/http/middleware"

	"github.com/gin-gonic/gin"
)

type GameHandler struct {
	gameService *services.GameService
	userRepo    repositories.UserRepository
}

func NewGameHandler(gameService *services.GameService, userRepo repositories.UserRepository) *GameHandler {
	return &GameHandler{
		gameService: gameService,
		userRepo:    userRepo,
	}
}

func (h *GameHandler) GetStages(c *gin.Context) {
	stages, err := h.gameService.GetActiveStages(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	var response []dto.StageResponse
	for _, stage := range stages {
		response = append(response, dto.StageResponse{
			ID:         stage.ID,
			Name:       stage.Name,
			Difficulty: stage.Difficulty,
		})
	}

	c.JSON(http.StatusOK, response)
}

func (h *GameHandler) GetStageDetail(c *gin.Context) {
	stageID := c.Param("id")

	stage, phrases, err := h.gameService.GetStageWithPhrases(c.Request.Context(), stageID)
	if err != nil {
		if err == services.ErrStageNotFound {
			c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "stage not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	var phrasesResponse []dto.PhraseResponse
	for _, phrase := range phrases {
		phrasesResponse = append(phrasesResponse, dto.PhraseResponse{
			ID:             phrase.ID,
			Text:           phrase.Text,
			SequenceNumber: phrase.SequenceNumber,
			Multiplier:     phrase.BaseMultiplier,
		})
	}

	response := dto.StageResponse{
		ID:         stage.ID,
		Name:       stage.Name,
		ThemeID:    stage.ThemeID,
		Difficulty: stage.Difficulty,
		Phrases:    phrasesResponse,
	}

	c.JSON(http.StatusOK, response)
}

func (h *GameHandler) SubmitScore(c *gin.Context) {
	user := middleware.GetUserFromContext(c)
	if user == nil {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Error: "unauthorized"})
		return
	}

	var req dto.SubmitScoreRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	score, status, err := h.gameService.SubmitScore(
		c.Request.Context(),
		user.ID,
		req.StageID,
		req.TotalTimeMs,
		req.TotalErrors,
	)
	if err != nil {
		if err == services.ErrStageNotFound {
			c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "stage not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SubmitScoreResponse{
		Status:     status,
		FinalScore: score.FinalScore,
	})
}

func (h *GameHandler) GetLeaderboard(c *gin.Context) {
	stageID := c.Query("stage_id")
	if stageID == "" {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "stage_id is required"})
		return
	}

	limitStr := c.DefaultQuery("limit", "20")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 20
	}

	scores, err := h.gameService.GetLeaderboard(c.Request.Context(), stageID, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	var response []dto.LeaderboardEntry
	for _, score := range scores {
		// Get username
		user, err := h.userRepo.FindByID(c.Request.Context(), score.UserID)
		if err != nil {
			continue
		}
		response = append(response, dto.LeaderboardEntry{
			Username:    user.Username,
			FinalScore:  score.FinalScore,
			TotalTimeMs: score.TotalTimeMs,
		})
	}

	c.JSON(http.StatusOK, response)
}

package handlers

import (
	"net/http"

	"uwika_quick_typer_game/internal/application/services"
	"uwika_quick_typer_game/internal/infrastructure/http/dto"

	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	adminService *services.AdminService
}

func NewAdminHandler(adminService *services.AdminService) *AdminHandler {
	return &AdminHandler{adminService: adminService}
}

// Theme Management
func (h *AdminHandler) GetAllThemes(c *gin.Context) {
	themes, err := h.adminService.GetAllThemes(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	var response []dto.ThemeResponse
	for _, theme := range themes {
		response = append(response, dto.ThemeResponse{
			ID:          theme.ID,
			Name:        theme.Name,
			Description: theme.Description,
		})
	}

	c.JSON(http.StatusOK, response)
}

// Stage Management
func (h *AdminHandler) CreateStage(c *gin.Context) {
	var req dto.CreateStageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	stage, err := h.adminService.CreateStage(
		c.Request.Context(),
		req.Name,
		req.ThemeID,
		req.Difficulty,
		req.IsActive,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dto.StageResponse{
		ID:         stage.ID,
		Name:       stage.Name,
		ThemeID:    stage.ThemeID,
		Difficulty: stage.Difficulty,
		IsActive:   stage.IsActive,
	})
}

func (h *AdminHandler) UpdateStage(c *gin.Context) {
	stageID := c.Param("id")

	var req dto.UpdateStageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	stage, err := h.adminService.UpdateStage(
		c.Request.Context(),
		stageID,
		req.Name,
		req.ThemeID,
		req.Difficulty,
		req.IsActive,
	)
	if err != nil {
		if err == services.ErrStageNotFound {
			c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "stage not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.StageResponse{
		ID:         stage.ID,
		Name:       stage.Name,
		ThemeID:    stage.ThemeID,
		Difficulty: stage.Difficulty,
		IsActive:   stage.IsActive,
	})
}

func (h *AdminHandler) DeleteStage(c *gin.Context) {
	stageID := c.Param("id")

	err := h.adminService.DeleteStage(c.Request.Context(), stageID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.MessageResponse{Message: "stage deleted successfully"})
}

func (h *AdminHandler) GetAllStages(c *gin.Context) {
	stages, err := h.adminService.GetAllStages(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	var response []dto.StageResponse
	for _, stage := range stages {
		response = append(response, dto.StageResponse{
			ID:         stage.ID,
			Name:       stage.Name,
			ThemeID:    stage.ThemeID,
			Difficulty: stage.Difficulty,
			IsActive:   stage.IsActive,
		})
	}

	c.JSON(http.StatusOK, response)
}

// Phrase Management
func (h *AdminHandler) CreatePhrase(c *gin.Context) {
	var req dto.CreatePhraseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	phrase, err := h.adminService.CreatePhrase(
		c.Request.Context(),
		req.StageID,
		req.Text,
		req.SequenceNumber,
		req.BaseMultiplier,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dto.PhraseResponse{
		ID:             phrase.ID,
		StageID:        phrase.StageID,
		Text:           phrase.Text,
		SequenceNumber: phrase.SequenceNumber,
		Multiplier:     phrase.BaseMultiplier,
	})
}

func (h *AdminHandler) UpdatePhrase(c *gin.Context) {
	phraseID := c.Param("id")

	var req dto.UpdatePhraseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	phrase, err := h.adminService.UpdatePhrase(
		c.Request.Context(),
		phraseID,
		req.StageID,
		req.Text,
		req.SequenceNumber,
		req.BaseMultiplier,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.PhraseResponse{
		ID:             phrase.ID,
		StageID:        phrase.StageID,
		Text:           phrase.Text,
		SequenceNumber: phrase.SequenceNumber,
		Multiplier:     phrase.BaseMultiplier,
	})
}

func (h *AdminHandler) DeletePhrase(c *gin.Context) {
	phraseID := c.Param("id")

	err := h.adminService.DeletePhrase(c.Request.Context(), phraseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.MessageResponse{Message: "phrase deleted successfully"})
}

func (h *AdminHandler) GetPhrasesByStage(c *gin.Context) {
	stageID := c.Query("stage_id")
	if stageID == "" {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "stage_id is required"})
		return
	}

	phrases, err := h.adminService.GetPhrasesByStage(c.Request.Context(), stageID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	var response []dto.PhraseResponse
	for _, phrase := range phrases {
		response = append(response, dto.PhraseResponse{
			ID:             phrase.ID,
			StageID:        phrase.StageID,
			Text:           phrase.Text,
			SequenceNumber: phrase.SequenceNumber,
			Multiplier:     phrase.BaseMultiplier,
		})
	}

	c.JSON(http.StatusOK, response)
}

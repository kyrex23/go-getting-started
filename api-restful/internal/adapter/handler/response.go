package handler

import (
	"getting-started/api-restful/internal/core/domain"
)

// ==================== HeroResponse ====================

type heroResponse struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	ActualName string `json:"actual_name"`
}

func newHeroResponse(hero *domain.Hero) heroResponse {
	return heroResponse{
		ID:         hero.ID,
		Name:       hero.Name,
		ActualName: hero.ActualName,
	}
}

// ==================== ErrorResponse ====================

type errorResponse struct {
	Success bool   `json:"success" example:"false"`
	Message string `json:"message" example:"Error message"`
}

func newErrorResponse(message string) errorResponse {
	return errorResponse{
		Success: false,
		Message: message,
	}
}

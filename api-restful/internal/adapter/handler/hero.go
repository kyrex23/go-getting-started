package handler

import (
	"getting-started/api-restful/internal/core/port"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// HeroHandler represents the HTTP handler for hero-related requests
type HeroHandler struct {
	heroService port.HeroService
}

func NewHeroHandler(heroService port.HeroService) *HeroHandler {
	return &HeroHandler{heroService: heroService}
}

// newHeroRequest represents the request body for creating a hero
type newHeroRequest struct {
	Name       string `json:"name" binding:"required" example:"Mr. Hero"`
	ActualName string `json:"actual_name" binding:"required" example:"John Doe"`
}

func (heroHandler *HeroHandler) GetHeroes(ctx *gin.Context) {
	var heroesList []heroResponse
	heroes, err := heroHandler.heroService.GetHeroes(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	for _, hero := range heroes {
		heroesList = append(heroesList, newHeroResponse(&hero))
	}
	ctx.JSON(http.StatusOK, heroesList)
}

func (heroHandler *HeroHandler) GetHeroById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Params.ByName("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	hero, err := heroHandler.heroService.GetHeroById(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	heroResponse := newHeroResponse(hero)
	ctx.JSON(http.StatusOK, heroResponse)
}

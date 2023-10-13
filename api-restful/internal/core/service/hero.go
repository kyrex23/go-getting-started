package service

import (
	"context"
	"getting-started/api-restful/internal/core/domain"
	"getting-started/api-restful/internal/core/port"
)

// HeroService implements port.HeroService interface
// and provides an access to the user repository
type HeroService struct {
	heroRepository port.HeroRepository
}

func NewHeroService(heroRepository port.HeroRepository) *HeroService {
	return &HeroService{
		heroRepository: heroRepository,
	}
}

func (heroService *HeroService) GetHeroes(ctx context.Context) ([]domain.Hero, error) {
	heroes, err := heroService.heroRepository.GetHeroes(ctx)
	if err != nil {
		return nil, err
	}
	return heroes, nil
}

func (heroService *HeroService) GetHeroById(ctx context.Context, id int) (*domain.Hero, error) {
	hero, err := heroService.heroRepository.GetHeroById(ctx, id)
	if err != nil {
		return nil, err
	}
	return hero, nil
}

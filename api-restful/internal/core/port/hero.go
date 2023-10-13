package port

import (
	"context"
	"getting-started/api-restful/internal/core/domain"
)

// HeroRepository is an interface for interacting with hero-related data
type HeroRepository interface {
	// Gets all the heroes stored in the database
	GetHeroes(ctx context.Context) ([]domain.Hero, error)
	// Gets a hero by its id
	GetHeroById(ctx context.Context, id int) (*domain.Hero, error)
}

// HeroService is an interface for interacting with hero-related business logic
type HeroService interface {
	// Gets a list of heroes
	GetHeroes(ctx context.Context) ([]domain.Hero, error)
	// Gets a hero by its id
	GetHeroById(ctx context.Context, id int) (*domain.Hero, error)
}

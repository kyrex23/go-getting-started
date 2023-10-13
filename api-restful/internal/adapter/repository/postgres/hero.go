package postgres

import (
	"context"
	"getting-started/api-restful/internal/core/domain"

	"github.com/Masterminds/squirrel"
)

// HeroRepository implements port.HeroRepository interface
// and provides an access to the postgres database
type HeroRepository struct {
	db *DB
}

func NewHeroRepository(db *DB) *HeroRepository {
	return &HeroRepository{
		db,
	}
}

func (hr *HeroRepository) GetHeroes(ctx context.Context) ([]domain.Hero, error) {
	var hero domain.Hero
	var heroes []domain.Hero

	query := psql.Select("*").
		From("heroes").
		OrderBy("id")

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := hr.db.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&hero.ID,
			&hero.Name,
			&hero.ActualName,
		)
		if err != nil {
			return nil, err
		}
		heroes = append(heroes, hero)
	}
	return heroes, nil
}

func (hr *HeroRepository) GetHeroById(ctx context.Context, id int) (*domain.Hero, error) {
	var hero domain.Hero

	query := psql.Select("*").
		From("heroes").
		Where(squirrel.Eq{"id": id}).
		Limit(1)

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	err = hr.db.QueryRow(ctx, sql, args...).Scan(
		&hero.ID,
		&hero.Name,
		&hero.ActualName,
	)
	if err != nil {
		return nil, err
	}
	return &hero, nil
}

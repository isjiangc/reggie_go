package repository

import (
	"context"
)

type DishRepository interface {
	QueryCountByCategoryId(ctx context.Context, id int64) (*int, error)
}

func NewDishRepository(repository *Repository) DishRepository {
	return &dishRepository{
		Repository: repository,
	}
}

type dishRepository struct {
	*Repository
}

func (d *dishRepository) QueryCountByCategoryId(ctx context.Context, id int64) (*int, error) {
	selStr := `
			SELECT
				COUNT(*)
			FROM
				dish
			WHERE
				1 = 1
				AND category_id = ?`
	var count int
	err := d.db2.Get(&count, selStr, id)
	if err != nil {
		return nil, err
	}
	return &count, nil

}

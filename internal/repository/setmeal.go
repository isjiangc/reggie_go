package repository

import (
	"context"
	"fmt"
	"reggie_go/internal/model"
)

type SetmealRepository interface {
	GetSetmealCountByname(ctx context.Context, name string) (int, error)
	GetSetmealByPage(ctx context.Context, page int, size int, name string) ([]model.Setmeal, error)
}

type setmealRepository struct {
	*Repository
}

func (s *setmealRepository) GetSetmealCountByname(ctx context.Context, name string) (int, error) {
	var count int
	sqlStr := `
		SELECT
			COUNT(*)
		FROM
			setmeal
		WHERE
			1 = 1
			AND is_deleted = 0`
	if name != "" {
		sqlStr += fmt.Sprintf(` AND name LIKE '%%%s%%'`, name)
	}
	err := s.db2.Get(&count, sqlStr)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (s *setmealRepository) GetSetmealByPage(ctx context.Context, page int, size int, name string) ([]model.Setmeal, error) {
	sqlStr := `
	SELECT
		id,
		category_id,
		name,
		price,
		status,
		code,
		description,
		image,
		create_time,
		update_time,
		create_user,
		update_user,
		is_deleted
	FROM
		setmeal
	WHERE
		1 = 1
		AND is_deleted = 0`
	if name != "" {
		sqlStr += fmt.Sprintf(` AND name LIKE '%%%s%%'`, name)
	}
	var setmeal []model.Setmeal
	offset := (page - 1) * size
	sqlStr += fmt.Sprintf(` ORDER BY update_time DESC LIMIT %d OFFSET %d`, size, offset)
	err := s.db2.Select(&setmeal, sqlStr)
	if err != nil {
		return nil, err
	}
	return setmeal, nil
}

func NewSetmealRepository(r *Repository) SetmealRepository {
	return &setmealRepository{
		Repository: r,
	}
}

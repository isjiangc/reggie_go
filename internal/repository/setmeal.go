package repository

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"reggie_go/internal/model"
)

type SetmealRepository interface {
	DeleteSetmeal(ctx context.Context, ids []string) (int, error)
	GetDeleteSetmealStatusCount(ctx context.Context, ids []string) (int, error)
	GetSetmealCountByname(ctx context.Context, name string) (int, error)
	GetSetmealByPage(ctx context.Context, page int, size int, name string) ([]model.Setmeal, error)
}

type setmealRepository struct {
	*Repository
}

func (s *setmealRepository) DeleteSetmeal(ctx context.Context, ids []string) (int, error) {
	tx, err := s.db2.Begin()
	if err != nil {
		return -1, err
	}
	defer func() { _ = tx.Rollback() }()
	sqlStr := `
		UPDATE
			setmeal
		SET
			is_deleted = 1
		WHERE
			1=1
			AND id IN (?)`
	query, id, err := sqlx.In(sqlStr, ids)
	if err != nil {
		return -1, nil
	}
	query = s.db2.Rebind(query)
	ret, err := s.db2.Exec(query, id...)
	if err != nil {
		return -1, err
	}
	affected, err := ret.RowsAffected()
	if err != nil || affected <= 0 {
		return -1, nil
	}
	sqlStr2 := `
		UPDATE
			setmeal_dish 
		SET
			is_deleted = 1
		WHERE
			1=1
			AND setmeal_id  IN (?)`
	query2, setmeal_id, err := sqlx.In(sqlStr2, ids)
	if err != nil {
		return -1, nil
	}
	query2 = s.db2.Rebind(query2)
	ret2, err := s.db2.Exec(query2, setmeal_id...)
	if err != nil {
		return -1, err
	}
	rowsAffected, err := ret2.RowsAffected()
	if err != nil {
		return -1, nil
	}
	return int(rowsAffected), nil
}

func (s *setmealRepository) GetDeleteSetmealStatusCount(ctx context.Context, ids []string) (int, error) {
	var count int
	sqlStr := `
	SELECT
		COUNT(*)
	FROM
		setmeal
	WHERE
		1 = 1
		AND status = 1
        AND id in(?)`
	query, id, err := sqlx.In(sqlStr, ids)
	if err != nil {
		return -1, nil
	}
	query = s.db2.Rebind(query)
	err = s.db2.Get(&count, query, id...)
	if err != nil {
		return -1, nil
	}
	return count, nil
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

package repository

import (
	"context"
	"reggie_go/internal/model"
)

type DishRepository interface {
	SaveDishWithFlavor(ctx context.Context, dish model.Dish, flavors []model.DishFlavor) (int64, error)
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

func (d *dishRepository) SaveDishWithFlavor(ctx context.Context, dish model.Dish, flavors []model.DishFlavor) (int64, error) {
	tx, err := d.db2.Begin()
	if err != nil {
		return 0, err
	}
	defer func() { _ = tx.Rollback() }()
	sqlStr := `INSERT INTO dish (name, category_id, price, code, image, description,status, sort,
								create_time, update_time, create_user, update_user, is_deleted) 
								VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`
	ret, err := tx.Exec(sqlStr, dish.Name, dish.CategoryId, dish.Price, dish.Code, dish.Image, dish.Description, dish.Status, dish.Sort,
		dish.CreateTime, dish.UpdateTime, dish.CreateUser, dish.UpdateUser, dish.IsDeleted)
	if err != nil {
		return 0, nil
	}
	dishId, err := ret.LastInsertId()
	if err != nil {
		return 0, err
	}
	if dishId <= 0 {
		return 0, err
	}
	sqlStr2 := `INSERT INTO dish_flavor (dish_id, name, value, create_time, update_time,create_user, update_user, is_deleted) 
         VALUES(?, ?, ?, ?, ?, ?, ?,?)`
	for _, v := range flavors {
		_, err := tx.Exec(sqlStr2, dishId, v.Name, v.Value, dish.CreateTime, dish.UpdateTime, dish.CreateUser, dish.UpdateUser, 0)
		if err != nil {
			return 0, err
		}
	}
	return dishId, tx.Commit()
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

package repository

import (
	"context"
	"reggie_go/internal/model"
)

type DishFlavorRepository interface {
	GetDishFlavorByDishId(ctx context.Context, dishId int64) ([]*model.DishFlavor, error)
	SaveDishFlavor(ctx context.Context, flavors []*model.DishFlavor) (int64, error)
}

func NewDishFlavorRepository(repository *Repository) DishFlavorRepository {
	return &dishflavorRepository{
		Repository: repository,
	}
}

type dishflavorRepository struct {
	*Repository
}

func (d *dishflavorRepository) GetDishFlavorByDishId(ctx context.Context, dishId int64) ([]*model.DishFlavor, error) {
	sqlStr := `
			SELECT
				id,
				dish_id,
				name,
				value,
				create_time,
				update_time,
				create_user,
				update_user,
				is_deleted
			FROM
				dish_flavor
				WHERE 1=1
				AND is_deleted = 0
				AND dish_id = ?`
	var dishFlavor []*model.DishFlavor
	err := d.db2.Select(&dishFlavor, sqlStr, dishId)
	if err != nil {
		return nil, err
	}
	return dishFlavor, nil
}

func (d *dishflavorRepository) SaveDishFlavor(ctx context.Context, flavors []*model.DishFlavor) (int64, error) {
	sqlStr := `INSERT INTO dish_flavor (dish_id, name, value, create_time, update_time,create_user, update_user, is_deleted) 
         VALUES(:dish_id, :name, :value, :create_time, :update_time, :create_user, :update_user,:is_deleted)`
	result, err := d.db2.NamedExec(sqlStr, flavors)
	if err != nil {
		return 0, err
	}
	var rows_affected int64
	rows_affected, err = result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rows_affected, nil
}

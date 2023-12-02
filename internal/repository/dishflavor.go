package repository

import (
	"context"
	"reggie_go/internal/model"
)

type DishFlavorRepository interface {
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

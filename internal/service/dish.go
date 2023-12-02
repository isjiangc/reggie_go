package service

import (
	"context"
	"reggie_go/internal/model"
	"reggie_go/internal/repository"
	"time"
)

type DishService interface {
	SaveDishWithFlavor(ctx context.Context, dis model.Dish, flavors []model.DishFlavor) (int64, error)
}

type TransactionSqlx interface {
	InTx(context.Context, func(ctx context.Context) error) error
}

func NewDishService(service *Service, dishRepository repository.DishRepository,
	dishFlavorRepository repository.DishFlavorRepository, repository *repository.Repository) DishService {
	return &dishService{
		Service:              service,
		dishRepository:       dishRepository,
		dishFlavorRepository: dishFlavorRepository,
		Repository:           repository,
	}
}

type dishService struct {
	*Service
	*repository.Repository
	dishRepository       repository.DishRepository
	dishFlavorRepository repository.DishFlavorRepository
}

func (d *dishService) SaveDishWithFlavor(ctx context.Context, dis model.Dish, flavors []model.DishFlavor) (int64, error) {
	var (
		dishId int64
		err    error
	)
	dishId, err = d.dishRepository.SaveDishWithFlavor(ctx, model.Dish{
		Name:        dis.Name,
		CategoryId:  dis.CategoryId,
		Price:       dis.Price,
		Code:        dis.Code,
		Image:       dis.Image,
		Description: dis.Description,
		Status:      dis.Status,
		Sort:        0,
		CreateTime:  time.Now(),
		UpdateTime:  time.Now(),
		CreateUser:  dis.CreateUser,
		UpdateUser:  dis.UpdateUser,
		IsDeleted:   0,
	}, flavors)
	if err != nil {
		d.Service.logger.Error("SaveDishWithFlavor is failed")
		return 0, err
	}
	return dishId, nil
}

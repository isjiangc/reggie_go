package service

import (
	"context"
	v1 "reggie_go/api/v1"
	"reggie_go/internal/model"
	"reggie_go/internal/repository"
	"time"
)

type DishService interface {
	GetDishByPage(ctx context.Context, req *v1.GetDishByPageRequest) (*v1.GetDishByPageData, error)
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

func (d *dishService) GetDishByPage(ctx context.Context, req *v1.GetDishByPageRequest) (*v1.GetDishByPageData, error) {
	if req.PageNum < 1 || req.PageSize < 1 {
		return nil, nil
	}
	count, err := d.dishRepository.GetDishCountByName(ctx, req.Name)
	if err != nil {
		return nil, err
	}
	dishDtoList, err := d.dishRepository.GetDishByPage(ctx, req.PageNum, req.PageSize, req.Name)
	if err != nil {
		return nil, err
	}
	var dishList []*v1.Dish
	for _, dishDto := range dishDtoList {
		dis := &v1.Dish{}
		dis.Id = dishDto.Id
		dis.Name = dishDto.Name
		dis.CategoryId = dishDto.CategoryId
		dis.Price = dishDto.Price
		dis.Code = dishDto.Code
		dis.Image = dishDto.Image
		dis.Description = dishDto.Description
		dis.Status = dishDto.Status
		dis.Sort = dishDto.Sort
		dis.CreateTime = dishDto.CreateTime
		dis.UpdateTime = dishDto.UpdateTime
		dis.UpdateUser = dishDto.UpdateUser
		dis.CategoryName = dishDto.CategoryName
		dishList = append(dishList, dis)
	}
	return &v1.GetDishByPageData{
		Records: dishList,
		Total:   count,
		Size:    req.PageSize,
	}, nil

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

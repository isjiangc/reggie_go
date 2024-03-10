package service

import (
	"context"
	"time"

	v1 "reggie_go/api/v1"
	"reggie_go/internal/model"
	"reggie_go/internal/repository"
)

type DishService interface {
	GetDishById(ctx context.Context, req *v1.GetDishByIdRequest) (*v1.GetDishByIdData, error)
	GetDishByPage(ctx context.Context, req *v1.GetDishByPageRequest) (*v1.GetDishByPageData, error)
	SaveDishWithFlavor(ctx context.Context, dis model.Dish, flavors []model.DishFlavor) (int64, error)
}

type TransactionSqlx interface {
	InTx(context.Context, func(ctx context.Context) error) error
}

func NewDishService(service *Service, dishRepository repository.DishRepository,
	dishFlavorRepository repository.DishFlavorRepository, repository *repository.Repository,
) DishService {
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

func (d *dishService) GetDishById(ctx context.Context, req *v1.GetDishByIdRequest) (*v1.GetDishByIdData, error) {
	dish, err := d.dishRepository.GetDishById(ctx, req.Id)
	if err != nil {
		return nil, v1.ErrDishNotExit
	}
	dishFlavors, err := d.dishFlavorRepository.GetDishFlavorByDishId(ctx, req.Id)
	if err != nil {
		return nil, v1.ErrDishFlavorNotExit
	}
	var dishFlavorList []*v1.DishFlavor
	for _, dishFlavorDto := range dishFlavors {
		dishFl := &v1.DishFlavor{}
		dishFl.Id = dishFlavorDto.Id
		dishFl.DishId = dishFlavorDto.DishId
		dishFl.Name = dishFlavorDto.Name
		dishFl.Value = dishFlavorDto.Value
		dishFl.CreateTime = dishFlavorDto.CreateTime
		dishFl.UpdateTime = dishFlavorDto.UpdateTime
		dishFl.CreateUser = dishFlavorDto.CreateUser
		dishFl.UpdateUser = dishFlavorDto.UpdateUser
		dishFl.IsDeleted = dishFlavorDto.IsDeleted
		dishFlavorList = append(dishFlavorList, dishFl)
	}
	return &v1.GetDishByIdData{
		Id:          dish.Id,
		Name:        dish.Name,
		CategoryId:  dish.CategoryId,
		Price:       dish.Price,
		Code:        dish.Code,
		Image:       dish.Image,
		Description: dish.Description,
		Status:      dish.Status,
		Sort:        dish.Sort,
		CreateTime:  dish.CreateTime,
		UpdateTime:  dish.UpdateTime,
		CreateUser:  dish.CreateUser,
		UpdateUser:  dish.UpdateUser,
		IsDeleted:   dish.IsDeleted,
		Falavors:    dishFlavorList,
	}, nil
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
	if err != nil && dishId == 0 {
		d.Service.logger.Error("SaveDishWithFlavor is failed")
		return 0, err
	}
	return dishId, nil
}

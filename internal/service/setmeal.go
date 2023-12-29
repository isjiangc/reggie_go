package service

import (
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	v1 "reggie_go/api/v1"
	"reggie_go/internal/repository"
)

type SetmealService interface {
	GetSetmealByPage(ctx context.Context, req *v1.GetSetmealByPageRequest) (*v1.GetSetmealByPageData, error)
}

func NewSetmealService(service *Service, setmealRepo repository.SetmealRepository) SetmealService {
	return &setmealService{
		setmealRepo: setmealRepo,
		Service:     service,
	}
}

type setmealService struct {
	setmealRepo repository.SetmealRepository
	*Service
}

func (s *setmealService) GetSetmealByPage(ctx context.Context, req *v1.GetSetmealByPageRequest) (*v1.GetSetmealByPageData, error) {
	if req.PageNum < 1 || req.PageSize < 1 {
		return nil, nil
	}
	count, err := s.setmealRepo.GetSetmealCountByname(ctx, req.Name)
	if err != nil {
		return nil, err
	}
	setmealList, err := s.setmealRepo.GetSetmealByPage(ctx, req.PageNum, req.PageSize, req.Name)
	if err != nil {
		return nil, err
	}
	var setmeal []v1.Setmeal
	err = copier.Copy(&setmeal, setmealList)
	if err != nil {
		return nil, err
	}
	fmt.Println(setmeal)
	return &v1.GetSetmealByPageData{
		Records: setmeal,
		Total:   count,
		Size:    req.PageSize,
	}, nil

}

package service

import (
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	v1 "reggie_go/api/v1"
	"reggie_go/internal/repository"
)

type SetmealService interface {
	UpdateSetmealStatus(ctx context.Context, req *v1.UpdateSellSetmealStatusRequest) error
	DeleteSetmeal(ctx context.Context, req *v1.DeleteSetmealRequest) error
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

func (s *setmealService) UpdateSetmealStatus(ctx context.Context, req *v1.UpdateSellSetmealStatusRequest) error {
	affected, err := s.setmealRepo.UpdateSellSetmealStatus(ctx, req.Status, req.Ids)
	if err != nil || affected < 0 {
		return err
	}
	return nil
}

func (s *setmealService) DeleteSetmeal(ctx context.Context, req *v1.DeleteSetmealRequest) error {
	count, err := s.setmealRepo.GetDeleteSetmealStatusCount(ctx, req.Ids)
	if err != nil {
		return err
	}
	if count > 0 {
		return v1.ErrSetmealTheSetmealIsSellIng
	}
	ret, err := s.setmealRepo.DeleteSetmeal(ctx, req.Ids)
	if err != nil {
		return err
	}
	if err == nil && ret > 0 {
		return nil
	}
	return nil
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

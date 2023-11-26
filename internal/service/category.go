package service

import (
	"context"
	"errors"
	"fmt"
	v1 "reggie_go/api/v1"
	"reggie_go/internal/model"
	"reggie_go/internal/repository"
	"time"
)

type CategoryService interface {
	CreateCategory(ctx context.Context, req *CreateCategoryData) error
}
type CreateCategoryData struct {
	Type       int       `json:"type"`       //类型   1 菜品分类 2 套餐分类
	Name       string    `json:"name"`       //分类名称
	Sort       int       `json:"sort"`       //顺序
	CreateTime time.Time `json:"createTime"` //创建时间
	UpdateTime time.Time `json:"updateTime"` //更新时间
	CreateUser int64     `json:"createUser"` //创建人
	UpdateUser int64     `json:"updateUser"` //修改人
}

func NewCategoryService(service *Service, categoryRepository repository.CategoryRepository) CategoryService {
	return &categoryService{
		Service:            service,
		categoryRepository: categoryRepository,
	}
}

type categoryService struct {
	*Service
	categoryRepository repository.CategoryRepository
}

func (c *categoryService) CreateCategory(ctx context.Context, req *CreateCategoryData) error {
	// 使用name查询是否存在
	category, err := c.categoryRepository.FirstByName(ctx, req.Name)
	if err == nil && category != nil {
		// 分类存在
		return errors.New(fmt.Sprintf("%s%s", category.Name, "菜品分类已存在"))
	}
	theId, err := c.categoryRepository.Save(ctx, &model.Category{
		Type:       req.Type,
		Name:       req.Name,
		Sort:       req.Sort,
		CreateTime: req.CreateTime,
		UpdateTime: req.UpdateTime,
		CreateUser: req.CreateUser,
		UpdateUser: req.UpdateUser,
	})
	if err != nil {
		return v1.ErrInternalServerError
	} else if err == nil && theId > 0 {
		return nil
	}
	return nil

}

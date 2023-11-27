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
	DeleteCategory(ctx context.Context, req *v1.DeleteCategoryRequest) error
	GetCategoryPage(ctx context.Context, req *v1.GetCategoryPageRequest) (*v1.GetCategoryPageData, error)
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

func NewCategoryService(service *Service, categoryRepository repository.CategoryRepository, dishRepository repository.DishRepository) CategoryService {
	return &categoryService{
		Service:            service,
		categoryRepository: categoryRepository,
		dishRepository:     dishRepository,
	}
}

type categoryService struct {
	*Service
	categoryRepository repository.CategoryRepository
	dishRepository     repository.DishRepository
}

func (c *categoryService) DeleteCategory(ctx context.Context, req *v1.DeleteCategoryRequest) error {
	// 查询分类下是否关联了菜品
	count, err2 := c.dishRepository.QueryCountByCategoryId(ctx, req.Id)
	if err2 != nil {
		return v1.ErrInternalServerError
	}
	if count == nil {
		return v1.ErrInternalServerError
	}
	// 表明已有关联了
	if count != nil && *count > 0 {
		return v1.ErrCategoryHaveSomeDish
	}
	ret, err := c.categoryRepository.DeleteCategory(ctx, req.Id)
	if err != nil {
		return v1.ErrDeleteCategoryFailed
	} else if err == nil && ret > 0 {
		return nil
	}
	return nil
}

func (c *categoryService) GetCategoryPage(ctx context.Context, req *v1.GetCategoryPageRequest) (*v1.GetCategoryPageData, error) {
	if req.PageNum < 1 || req.PageSize < 1 {
		return nil, nil
	}
	count, err := c.categoryRepository.GetCount(ctx)
	if err != nil {
		return nil, err
	}
	categories, err := c.categoryRepository.GetByPage(ctx, req.PageNum, req.PageSize)
	if err != nil {
		return nil, err
	}
	return &v1.GetCategoryPageData{
		Records: categories,
		Total:   count,
		Size:    req.PageSize,
	}, nil
}

func (c *categoryService) CreateCategory(ctx context.Context, req *CreateCategoryData) error {
	if req.Type != 1 || req.Type != 2 {
		return v1.ErrCategoryTypeIsIllegal
	}
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

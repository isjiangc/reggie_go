package v1

import "reggie_go/internal/model"

type CreateCategoryRequest struct {
	Name string `json:"name"` // 分类名称
	Type string `json:"type"` // 类型   1 菜品分类 2 套餐分类
	Sort string `json:"sort"` // 顺序
}

type GetCategoryPageRequest struct {
	PageNum  int `json:"page_num"`
	PageSize int `json:"page_size"`
}
type GetCategoryPageData struct {
	Records []*model.Category `json:"records"`
	Total   int               `json:"total"`
	Size    int               `json:"size"`
}
type GetCategoryPageResponse struct {
	Response
	Data GetCategoryPageData
}

type DeleteCategoryRequest struct {
	Id int64 `json:"id"`
}

type UpdateCategoryRequest struct {
	Id   string `json:"id"`   // 主键
	Name string `json:"name"` // 分类名称
	Sort int    `json:"sort"` // 顺序
}

package v1

type CreateCategoryRequest struct {
	Name string `json:"name"` //分类名称
	Type string `json:"type"` //类型   1 菜品分类 2 套餐分类
	Sort string `json:"sort"` //顺序
}

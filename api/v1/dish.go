package v1

import "time"

type CreateDishRequest struct {
	Name        string    `json:"name"`        //菜品名称
	Price       float64   `json:"price"`       //菜品价格
	Code        string    `json:"code"`        //商品码
	Image       string    `json:"image"`       //图片
	Description string    `json:"description"` //描述信息
	Status      int       `json:"status"`      //0 停售 1 起售
	CategoryId  string    `json:"categoryId"`  //菜品分类id
	Flavors     []Flavors `json:"flavors"`     // 口味
}
type Flavors struct {
	Name       string `json:"name"`  // 口味名称
	Value      string `json:"value"` // 口味数据list
	ShowOption bool   `json:"showOption"`
}

type GetDishByPageRequest struct {
	PageNum  int
	PageSize int
	Name     string
}

type Dish struct {
	Id           int64     `json:"id"`           //主键
	Name         string    `json:"name"`         //菜品名称
	CategoryId   int64     `json:"categoryId"`   //菜品分类id
	Price        float64   `json:"price"`        //菜品价格
	Code         string    `json:"code"`         //商品码
	Image        string    `json:"image"`        //图片
	Description  string    `json:"description"`  //描述信息
	Status       int       `json:"status"`       //0 停售 1 起售
	Sort         int       `json:"sort"`         //顺序
	CreateTime   time.Time `json:"createTime"`   //创建时间
	UpdateTime   time.Time `json:"updateTime"`   //更新时间
	CreateUser   int64     `json:"createUser"`   //创建人
	UpdateUser   int64     `json:"updateUser"`   //修改人
	CategoryName string    `json:"categoryName"` //分类名称

}
type GetDishByPageData struct {
	Records []*Dish `json:"records"`
	Total   int     `json:"total"`
	Size    int     `json:"size"`
}

type GetDishByPageResponse struct {
	Response
	Data GetCategoryPageData
}

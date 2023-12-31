package v1

import (
	"time"
)

type GetSetmealByPageRequest struct {
	PageNum  int
	PageSize int
	Name     string
}

type Setmeal struct {
	Id          int64     `json:"id"`          //主键
	CategoryId  int64     `json:"categoryId"`  //菜品分类id
	Name        string    `json:"name"`        //套餐名称
	Price       float64   `json:"price"`       //套餐价格
	Status      int       `json:"status"`      //状态 0:停用 1:启用
	Code        string    `json:"code"`        //编码
	Description string    `json:"description"` //描述信息
	Image       string    `json:"image"`       //图片
	CreateTime  time.Time `json:"createTime"`  //创建时间
	UpdateTime  time.Time `json:"updateTime"`  //更新时间
	CreateUser  int64     `json:"createUser"`  //创建人
	UpdateUser  int64     `json:"updateUser"`  //修改人
	IsDeleted   int       `json:"isDeleted"`   //是否删除
}
type GetSetmealByPageData struct {
	Records []Setmeal `json:"records"`
	Total   int       `json:"total"`
	Size    int       `json:"size"`
}

type GetSetmealByPageResponse struct {
	Response
	Data GetEmployeeByPageData
}

type DeleteSetmealRequest struct {
	Ids []string `json:"ids"`
}

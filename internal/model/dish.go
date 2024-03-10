package model

import (
	"time"
)

// 菜品管理
type Dish struct {
	Id          int64     `db:"id"`          // 主键
	Name        string    `db:"name"`        // 菜品名称
	CategoryId  int64     `db:"category_id"` // 菜品分类id
	Price       float64   `db:"price"`       // 菜品价格
	Code        string    `db:"code"`        // 商品码
	Image       string    `db:"image"`       // 图片
	Description string    `db:"description"` // 描述信息
	Status      int       `db:"status"`      // 0 停售 1 起售
	Sort        int       `db:"sort"`        // 顺序
	CreateTime  time.Time `db:"create_time"` // 创建时间
	UpdateTime  time.Time `db:"update_time"` // 更新时间
	CreateUser  int64     `db:"create_user"` // 创建人
	UpdateUser  int64     `db:"update_user"` // 修改人
	IsDeleted   int       `db:"is_deleted"`  // 是否删除
}

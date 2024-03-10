package model

import (
	"time"
)

// 套餐
type Setmeal struct {
	Id          int64     `db:"id"`          // 主键
	CategoryId  int64     `db:"category_id"` // 菜品分类id
	Name        string    `db:"name"`        // 套餐名称
	Price       float64   `db:"price"`       // 套餐价格
	Status      int       `db:"status"`      // 状态 0:停用 1:启用
	Code        string    `db:"code"`        // 编码
	Description string    `db:"description"` // 描述信息
	Image       string    `db:"image"`       // 图片
	CreateTime  time.Time `db:"create_time"` // 创建时间
	UpdateTime  time.Time `db:"update_time"` // 更新时间
	CreateUser  int64     `db:"create_user"` // 创建人
	UpdateUser  int64     `db:"update_user"` // 修改人
	IsDeleted   int       `db:"is_deleted"`  // 是否删除
}

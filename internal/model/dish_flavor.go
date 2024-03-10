package model

import (
	"time"
)

// 菜品口味关系表
type DishFlavor struct {
	Id         int64     `db:"id"`          // 主键
	DishId     int64     `db:"dish_id"`     // 菜品
	Name       string    `db:"name"`        // 口味名称
	Value      string    `db:"value"`       // 口味数据list
	CreateTime time.Time `db:"create_time"` // 创建时间
	UpdateTime time.Time `db:"update_time"` // 更新时间
	CreateUser int64     `db:"create_user"` // 创建人
	UpdateUser int64     `db:"update_user"` // 修改人
	IsDeleted  int       `db:"is_deleted"`  // 是否删除
}

package model

import (
	"time"
)

// 套餐菜品关系
type SetmealDish struct {
	Id         int64     `db:"id"`          // 主键
	SetmealId  string    `db:"setmeal_id"`  // 套餐id
	DishId     string    `db:"dish_id"`     // 菜品id
	Name       string    `db:"name"`        // 菜品名称 （冗余字段）
	Price      float64   `db:"price"`       // 菜品原价（冗余字段）
	Copies     int       `db:"copies"`      // 份数
	Sort       int       `db:"sort"`        // 排序
	CreateTime time.Time `db:"create_time"` // 创建时间
	UpdateTime time.Time `db:"update_time"` // 更新时间
	CreateUser int64     `db:"create_user"` // 创建人
	UpdateUser int64     `db:"update_user"` // 修改人
	IsDeleted  int       `db:"is_deleted"`  // 是否删除
}

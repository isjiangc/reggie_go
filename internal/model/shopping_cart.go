package model

import (
	"time"
)

// 购物车
type ShoppingCart struct {
	Id         int64     `db:"id"`          // 主键
	Name       string    `db:"name"`        // 名称
	Image      string    `db:"image"`       // 图片
	UserId     int64     `db:"user_id"`     // 主键
	DishId     int64     `db:"dish_id"`     // 菜品id
	SetmealId  int64     `db:"setmeal_id"`  // 套餐id
	DishFlavor string    `db:"dish_flavor"` // 口味
	Number     int       `db:"number"`      // 数量
	Amount     float64   `db:"amount"`      // 金额
	CreateTime time.Time `db:"create_time"` // 创建时间
}

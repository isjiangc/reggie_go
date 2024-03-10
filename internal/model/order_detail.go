package model

// 订单明细表
type OrderDetail struct {
	Id         int64   `db:"id"`          // 主键
	Name       string  `db:"name"`        // 名字
	Image      string  `db:"image"`       // 图片
	OrderId    int64   `db:"order_id"`    // 订单id
	DishId     int64   `db:"dish_id"`     // 菜品id
	SetmealId  int64   `db:"setmeal_id"`  // 套餐id
	DishFlavor string  `db:"dish_flavor"` // 口味
	Number     int     `db:"number"`      // 数量
	Amount     float64 `db:"amount"`      // 金额
}

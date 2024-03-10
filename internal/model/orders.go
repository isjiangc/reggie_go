package model

import (
	"time"
)

// 订单表
type Orders struct {
	Id            int64     `db:"id"`              // 主键
	Number        string    `db:"number"`          // 订单号
	Status        int       `db:"status"`          // 订单状态 1待付款，2待派送，3已派送，4已完成，5已取消
	UserId        int64     `db:"user_id"`         // 下单用户
	AddressBookId int64     `db:"address_book_id"` // 地址id
	OrderTime     time.Time `db:"order_time"`      // 下单时间
	CheckoutTime  time.Time `db:"checkout_time"`   // 结账时间
	PayMethod     int       `db:"pay_method"`      // 支付方式 1微信,2支付宝
	Amount        float64   `db:"amount"`          // 实收金额
	Remark        string    `db:"remark"`          // 备注
	Phone         string    `db:"phone"`
	Address       string    `db:"address"`
	UserName      string    `db:"user_name"`
	Consignee     string    `db:"consignee"`
}

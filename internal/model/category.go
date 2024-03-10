package model

import (
	"time"
)

// 菜品及套餐分类
type Category struct {
	Id         int64     `db:"id"json:"id"`                  // 主键
	Type       int       `db:"type"json:"type"`              // 类型   1 菜品分类 2 套餐分类
	Name       string    `db:"name"json:"name"`              // 分类名称
	Sort       int       `db:"sort"json:"sort"`              // 顺序
	CreateTime time.Time `db:"create_time"json:"createTime"` // 创建时间
	UpdateTime time.Time `db:"update_time"json:"updateTime"` // 更新时间
	CreateUser int64     `db:"create_user"json:"createUser"` // 创建人
	UpdateUser int64     `db:"update_user"json:"updateUser"` // 修改人
}

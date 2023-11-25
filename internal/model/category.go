package model

import (
	"time"
)

// 菜品及套餐分类
type Category struct {
	Id         int64     `db:"id"`          //主键
	Type       int       `db:"type"`        //类型   1 菜品分类 2 套餐分类
	Name       string    `db:"name"`        //分类名称
	Sort       int       `db:"sort"`        //顺序
	CreateTime time.Time `db:"create_time"` //创建时间
	UpdateTime time.Time `db:"update_time"` //更新时间
	CreateUser int64     `db:"create_user"` //创建人
	UpdateUser int64     `db:"update_user"` //修改人
}

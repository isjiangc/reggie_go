package model

import (
	"time"
)

// Employee 员工信息
type Employee struct {
	Id         int64     `db:"id"`          //主键
	Name       string    `db:"name"`        //姓名
	Username   string    `db:"username"`    //用户名
	Password   string    `db:"password"`    //密码
	Phone      string    `db:"phone"`       //手机号
	Sex        string    `db:"sex"`         //性别
	IdNumber   string    `db:"id_number"`   //身份证号
	Status     int       `db:"status"`      //状态 0:禁用，1:正常
	CreateTime time.Time `db:"create_time"` //创建时间
	UpdateTime time.Time `db:"update_time"` //更新时间
	CreateUser int64     `db:"create_user"` //创建人
	UpdateUser int64     `db:"update_user"` //修改人
}

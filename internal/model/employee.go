package model

import (
	"time"
)

// Employee 员工信息
type Employee struct {
	Id         int64     `db:"id"json:"id"`                  //主键
	Name       string    `db:"name"json:"name"`              //姓名
	Username   string    `db:"username"json:"username"`      //用户名
	Password   string    `db:"password"json:"password"`      //密码
	Phone      string    `db:"phone"json:"phone"`            //手机号
	Sex        string    `db:"sex"json:"sex"`                //性别
	IdNumber   string    `db:"id_number"json:"idNumber"`     //身份证号
	Status     int       `db:"status"json:"status"`          //状态 0:禁用，1:正常
	CreateTime time.Time `db:"create_time"json:"createTime"` //创建时间
	UpdateTime time.Time `db:"update_time"json:"updateTime"` //更新时间
	CreateUser int64     `db:"create_user"json:"createUser"` //创建人
	UpdateUser int64     `db:"update_user"json:"updateUser"` //修改人
}

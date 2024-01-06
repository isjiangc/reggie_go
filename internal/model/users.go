package model

// 用户信息
type Users struct {
	Id       int64  `db:"id"`        //主键
	Name     string `db:"name"`      //姓名
	Phone    string `db:"phone"`     //手机号
	Sex      string `db:"sex"`       //性别
	IdNumber string `db:"id_number"` //身份证号
	Avatar   string `db:"avatar"`    //头像
	Status   int    `db:"status"`    //状态 0:禁用，1:正常
}

package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id        uint   `gorm:"primarykey"`
	UserId    string `gorm:"unique;not null"`
	Nickname  string `gorm:"not null"`
	Password  string `gorm:"not null"`
	Email     string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (u *User) TableName() string {
	return "users"
}

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

package model

import (
	"time"
)

// 地址管理
type AddressBook struct {
	Id           int64     `db:"id"`            // 主键
	UserId       int64     `db:"user_id"`       // 用户id
	Consignee    string    `db:"consignee"`     // 收货人
	Sex          int8      `db:"sex"`           // 性别 0 女 1 男
	Phone        string    `db:"phone"`         // 手机号
	ProvinceCode string    `db:"province_code"` // 省级区划编号
	ProvinceName string    `db:"province_name"` // 省级名称
	CityCode     string    `db:"city_code"`     // 市级区划编号
	CityName     string    `db:"city_name"`     // 市级名称
	DistrictCode string    `db:"district_code"` // 区级区划编号
	DistrictName string    `db:"district_name"` // 区级名称
	Detail       string    `db:"detail"`        // 详细地址
	Label        string    `db:"label"`         // 标签
	IsDefault    int8      `db:"is_default"`    // 默认 0 否 1是
	CreateTime   time.Time `db:"create_time"`   // 创建时间
	UpdateTime   time.Time `db:"update_time"`   // 更新时间
	CreateUser   int64     `db:"create_user"`   // 创建人
	UpdateUser   int64     `db:"update_user"`   // 修改人
	IsDeleted    int       `db:"is_deleted"`    // 是否删除
}

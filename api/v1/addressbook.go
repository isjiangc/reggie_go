package v1

import "time"

type GetAddressBookByUserIdRequest struct {
	UserId int64 `json:"userId"` //用户id
}

type UpdateAddressBookIsDefaultRequest struct {
	Id     int64 `json:"id"`     //主键
	UserId int64 `json:"userId"` //用户id
}

type GetAddressBookByIdRequest struct {
	Id int64 `json:"id"` //主键
}

type SaveAddressBookRequest struct {
	AddressBook
}

type GetDefaultAddressBookRequest struct {
	UserId int64 `json:"userId"` //用户id

}

type AddressBook struct {
	Id           int64     `json:"id"`           //主键
	UserId       int64     `json:"userId"`       //用户id
	Consignee    string    `json:"consignee"`    //收货人
	Sex          int8      `json:"sex"`          //性别 0 女 1 男
	Phone        string    `json:"phone"`        //手机号
	ProvinceCode string    `json:"provinceCode"` //省级区划编号
	ProvinceName string    `json:"provinceName"` //省级名称
	CityCode     string    `json:"cityCode"`     //市级区划编号
	CityName     string    `json:"cityName"`     //市级名称
	DistrictCode string    `json:"districtCode"` //区级区划编号
	DistrictName string    `json:"districtName"` //区级名称
	Detail       string    `json:"detail"`       //详细地址
	Label        string    `json:"label"`        //标签
	IsDefault    int8      `json:"isDefault"`    //默认 0 否 1是
	CreateTime   time.Time `json:"createTime"`   //创建时间
	UpdateTime   time.Time `json:"updateTime"`   //更新时间
	CreateUser   int64     `json:"createUser"`   //创建人
	UpdateUser   int64     `json:"updateUser"`   //修改人
	IsDeleted    int       `json:"isDeleted"`    //是否删除
}

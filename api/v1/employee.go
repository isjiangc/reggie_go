package v1

import (
	"time"

	"reggie_go/internal/model"
)

type EmployeeLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type EmployeeLoginResponseData struct {
	Id         int64     `json:"id"`         // 主键
	Name       string    `json:"name"`       // 姓名
	Username   string    `json:"username"`   // 用户名
	Password   string    `json:"password"`   // 密码
	Phone      string    `json:"phone"`      // 手机号
	Sex        string    `json:"sex"`        // 性别
	IdNumber   string    `json:"idNumber"`   // 身份证号
	Status     int       `json:"status"`     // 状态 0:禁用，1:正常
	CreateTime time.Time `json:"createTime"` // 创建时间
	UpdateTime time.Time `json:"updateTime"` // 更新时间
	CreateUser int64     `json:"createUser"` // 创建人
	UpdateUser int64     `json:"updateUser"` // 修改人
}

type EmployeeLoginResponse struct {
	Response
	Data EmployeeLoginResponseData
}

type CreateEmployeeRequest struct {
	Name     string `json:"name"`     // 姓名
	Username string `json:"username"` // 用户名
	Phone    string `json:"phone"`    // 手机号
	Sex      string `json:"sex"`      // 性别
	IdNumber string `json:"idNumber"` // 身份证号
}

type GetEmployeeByPageRequest struct {
	PageNum  int
	PageSize int
	Name     string
}
type GetEmployeeByPageData struct {
	Records []*model.Employee `json:"records"`
	Total   int               `json:"total"`
	Size    int               `json:"size"`
}
type GetEmployeeByPageResponse struct {
	Response
	Data GetEmployeeByPageData
}

type UpdateEmployeeRequest struct {
	Id         int64     `json:"id"`         // 主键
	Name       string    `json:"name"`       // 姓名
	Username   string    `json:"username"`   // 用户名
	Password   string    `json:"password"`   // 密码
	Phone      string    `json:"phone"`      // 手机号
	Sex        string    `json:"sex"`        // 性别
	IdNumber   string    `json:"idNumber"`   // 身份证号
	Status     int       `json:"status"`     // 状态 0:禁用，1:正常
	CreateTime time.Time `json:"createTime"` // 创建时间
	UpdateTime time.Time `json:"updateTime"` // 更新时间
	CreateUser int64     `json:"createUser"` // 创建人
	UpdateUser int64     `json:"updateUser"` // 修改人
}

type GetEmployeeByIdData struct {
	Id         int64     `json:"id"`         // 主键
	Name       string    `json:"name"`       // 姓名
	Username   string    `json:"username"`   // 用户名
	Password   string    `json:"password"`   // 密码
	Phone      string    `json:"phone"`      // 手机号
	Sex        string    `json:"sex"`        // 性别
	IdNumber   string    `json:"idNumber"`   // 身份证号
	Status     int       `json:"status"`     // 状态 0:禁用，1:正常
	CreateTime time.Time `json:"createTime"` // 创建时间
	UpdateTime time.Time `json:"updateTime"` // 更新时间
	CreateUser int64     `json:"createUser"` // 创建人
	UpdateUser int64     `json:"updateUser"` // 修改人`
}

type GetEmployeeByPaIdResponse struct {
	Response
	Data GetEmployeeByIdData
}

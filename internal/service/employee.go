package service

import (
	"context"
	"errors"
	"fmt"
	v1 "reggie_go/api/v1"
	"reggie_go/internal/model"
	"reggie_go/internal/repository"
	"reggie_go/pkg/helper/md5"
	"time"
)

type EmployeeService interface {
	Login(ctx context.Context, req *v1.EmployeeLoginRequest) (*v1.EmployeeLoginResponseData, error)
	CreateEmployee(ctx context.Context, req *CreateEmployeeData) error
}

type CreateEmployeeData struct {
	Name       string `json:"name"`       //姓名
	Username   string `json:"username"`   //用户名
	Phone      string `json:"phone"`      //手机号
	Sex        string `json:"sex"`        //性别
	IdNumber   string `json:"idNumber"`   //身份证号
	CreateUser int64  `json:"createUser"` //创建人
	UpdateUser int64  `json:"updateUser"` //修改人
}

func NewEmployeeService(service *Service, employeeRepo repository.EmployeeRepository) EmployeeService {
	return &employeeService{
		employeeRepo: employeeRepo,
		Service:      service,
	}
}

type employeeService struct {
	employeeRepo repository.EmployeeRepository
	*Service
}

func (e *employeeService) CreateEmployee(ctx context.Context, req *CreateEmployeeData) error {
	// 使用用户名进行查询
	employee, err := e.employeeRepo.GetByUsername(ctx, req.Username)
	// 用户存在
	if err == nil && employee != nil {
		// 用户已经存在
		return errors.New(fmt.Sprintf("%s%s", employee.Username, "员工已存在"))
	}
	theId, err := e.employeeRepo.CreateEmployee(ctx, &model.Employee{
		Name:       req.Name,
		Username:   req.Username,
		Password:   md5.Md5("123456"),
		Phone:      req.Phone,
		Sex:        req.Sex,
		IdNumber:   req.IdNumber,
		Status:     1,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		CreateUser: req.CreateUser,
		UpdateUser: req.UpdateUser,
	})
	if err != nil {
		return v1.ErrInternalServerError
	} else if err == nil && theId > 0 {
		return nil
	}
	return nil
}

func (e *employeeService) Login(ctx context.Context, req *v1.EmployeeLoginRequest) (*v1.EmployeeLoginResponseData, error) {
	password := md5.Md5(req.Password)
	employee, err := e.employeeRepo.GetByUsername(ctx, req.Username)
	if err != nil || employee == nil {
		return nil, v1.ErrEmployeeLoginFailed
	}
	if employee.Password != password {
		return nil, v1.ErrEmployeeLoginFailed
	}
	if employee.Status == 0 {
		return nil, v1.ErrEmployeeAccountIsDisabled
	}
	return &v1.EmployeeLoginResponseData{
		Id:         employee.Id,
		Name:       employee.Name,
		Username:   employee.Username,
		Password:   employee.Password,
		Phone:      employee.Phone,
		Sex:        employee.Sex,
		IdNumber:   employee.IdNumber,
		Status:     employee.Status,
		CreateTime: employee.CreateTime,
		UpdateTime: employee.UpdateTime,
		CreateUser: employee.CreateUser,
		UpdateUser: employee.UpdateUser,
	}, nil
}

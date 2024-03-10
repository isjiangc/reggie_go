package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	v1 "reggie_go/api/v1"
	"reggie_go/internal/model"
	"reggie_go/internal/repository"
	"reggie_go/pkg/helper/md5"
)

type EmployeeService interface {
	Login(ctx context.Context, req *v1.EmployeeLoginRequest) (*v1.EmployeeLoginResponseData, error)
	CreateEmployee(ctx context.Context, req *CreateEmployeeData) error
	GetEmployeeByPage(ctx context.Context, req *v1.GetEmployeeByPageRequest) (*v1.GetEmployeeByPageData, error)
	UpdateEmployee(ctx context.Context, req *v1.UpdateEmployeeRequest) error
	GetEmployeeById(ctx context.Context, id int64) (*v1.GetEmployeeByIdData, error)
}

type CreateEmployeeData struct {
	Name       string `json:"name"`       // 姓名
	Username   string `json:"username"`   // 用户名
	Phone      string `json:"phone"`      // 手机号
	Sex        string `json:"sex"`        // 性别
	IdNumber   string `json:"idNumber"`   // 身份证号
	CreateUser int64  `json:"createUser"` // 创建人
	UpdateUser int64  `json:"updateUser"` // 修改人
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

func (e *employeeService) GetEmployeeByPage(ctx context.Context, req *v1.GetEmployeeByPageRequest) (*v1.GetEmployeeByPageData, error) {
	if req.PageNum < 1 || req.PageSize < 1 {
		return nil, nil
	}
	// 查询总数
	count, err := e.employeeRepo.GetEmployeeCountByUsername(ctx, req.Name)
	if err != nil {
		return nil, err
	}
	// 分页查询
	empList, err := e.employeeRepo.GetEmployeeByPage(ctx, req.PageNum, req.PageSize, req.Name)
	if err != nil {
		return nil, err
	}
	return &v1.GetEmployeeByPageData{
		Records: empList,
		Total:   count,
		Size:    req.PageSize,
	}, nil
}

// UpdateEmployee 更新员工信息
func (e *employeeService) UpdateEmployee(ctx context.Context, req *v1.UpdateEmployeeRequest) error {
	ret, err := e.employeeRepo.UpdateEmployee(ctx, &model.Employee{
		Id:         req.Id,
		Name:       req.Name,
		Username:   req.Username,
		Password:   req.Password,
		Phone:      req.Phone,
		Sex:        req.Sex,
		IdNumber:   req.IdNumber,
		Status:     req.Status,
		CreateTime: req.CreateTime,
		UpdateTime: req.UpdateTime,
		CreateUser: req.CreateUser,
		UpdateUser: req.UpdateUser,
	})
	if err != nil {
		return v1.ErrEmployeeUpdatedFailed
	} else if err == nil && ret > 0 {
		return nil
	}
	return nil
}

func (e *employeeService) GetEmployeeById(ctx context.Context, id int64) (*v1.GetEmployeeByIdData, error) {
	employee, err := e.employeeRepo.GetEmployeeById(ctx, id)
	if err != nil {
		return nil, v1.ErrEmployeeNotExit
	}
	return &v1.GetEmployeeByIdData{
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

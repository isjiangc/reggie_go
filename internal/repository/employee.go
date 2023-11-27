package repository

import (
	"context"
	"fmt"
	"reggie_go/internal/model"
)

type EmployeeRepository interface {
	GetByUsername(ctx context.Context, username string) (*model.Employee, error)
	CreateEmployee(ctx context.Context, employee *model.Employee) (int64, error)
	GetEmployeeByPage(ctx context.Context, page int, size int, name string) ([]*model.Employee, error)
	GetEmployeeCountByUsername(ctx context.Context, name string) (int, error)
	UpdateEmployee(cxt context.Context, employee *model.Employee) (int64, error)
	GetEmployeeById(ctx context.Context, id int64) (*model.Employee, error)
}

func NewEmployeeRepository(r *Repository) EmployeeRepository {
	return &employeeRepository{
		Repository: r,
	}
}

type employeeRepository struct {
	*Repository
}

func (e *employeeRepository) GetByUsername(ctx context.Context, username string) (*model.Employee, error) {
	sqlStr := `SELECT
					id,
					name,
					username,
					password,
					phone,
					sex,
					id_number,
					status,
					create_time,
					update_time,
					create_user,
					update_user
				FROM
					employee
				WHERE
					1 = 1
					AND username = ?;`
	var employee = model.Employee{}
	err := e.db2.Get(&employee, sqlStr, username)
	if err != nil {
		return nil, err
	}
	return &employee, nil
}

// CreateEmployee 新增员工
func (e *employeeRepository) CreateEmployee(ctx context.Context, employee *model.Employee) (int64, error) {
	sqlStr := `INSERT INTO employee (name, username, password, phone, sex, id_number, status, create_time, update_time, create_user, update_user)
                VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`
	ret, err := e.db2.Exec(sqlStr, employee.Name, employee.Username, employee.Password, employee.Phone,
		employee.Sex, employee.IdNumber, employee.Status, employee.CreateTime,
		employee.UpdateTime, employee.CreateUser, employee.UpdateUser)
	if err != nil {
		return 0, err
	}
	theId, err := ret.LastInsertId()
	if err != nil {
		return 0, err
	}
	return theId, nil
}

// GetEmployeeByPage 分页查询员工信息
func (e *employeeRepository) GetEmployeeByPage(ctx context.Context, page int, size int, name string) ([]*model.Employee, error) {
	sqlStr := `  
			SELECT
				id,
				name,
				username,
				password,
				phone,
				sex,
				id_number,
				status,
				create_time,
				update_time,
				create_user,
				update_user
			FROM
				employee
			WHERE
				1 = 1`
	if name != "" {
		sqlStr += fmt.Sprintf(` AND name LIKE '%%%s%%'`, name)
	}
	var employee []*model.Employee
	offset := (page - 1) * size
	sqlStr += fmt.Sprintf(` ORDER BY update_time DESC LIMIT %d OFFSET %d`, size, offset)
	err := e.db2.Select(&employee, sqlStr)
	if err != nil {
		return nil, err
	}
	return employee, nil
}

// GetEmployeeCountByUsername 查询总数
func (e *employeeRepository) GetEmployeeCountByUsername(ctx context.Context, name string) (int, error) {
	var count int
	sqlStr := `
		SELECT
			COUNT(*)
		FROM
			employee
		WHERE
			1 = 1`
	if name != "" {
		sqlStr += fmt.Sprintf(` AND name LIKE '%%%s%%'`, name)
	}
	err := e.db2.Get(&count, sqlStr)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// UpdateEmployee 根据id更新员工信息
func (e *employeeRepository) UpdateEmployee(cxt context.Context, employee *model.Employee) (int64, error) {
	sqlStr := `
			UPDATE
				employee
			SET
				name = ?,
				password = ?,
				phone = ?,
				sex = ?,
				id_number = ?,
				status = ?,
				create_time = ?,
				update_time = ?,
				create_user = ?,
				update_user = ?
			WHERE
				1=1
				AND id  = ?;
			`
	ret, err := e.db2.Exec(sqlStr, employee.Name, employee.Password, employee.Phone, employee.Sex,
		employee.IdNumber, employee.Status, employee.CreateTime, employee.UpdateTime,
		employee.CreateUser, employee.UpdateUser, employee.Id)
	if err != nil {
		return 0, err
	}
	n, err := ret.RowsAffected()
	if err != nil {
		return 0, err
	}
	fmt.Println(n)
	return n, nil
}

func (e *employeeRepository) GetEmployeeById(ctx context.Context, id int64) (*model.Employee, error) {
	sqlStr := `
		SELECT
			id,
			name,
			username,
			password,
			phone,
			sex,
			id_number,
			status,
			create_time,
			update_time,
			create_user,
			update_user
		FROM
			employee
		WHERE
			1 = 1
			AND id = ?`
	var employee model.Employee
	err := e.db2.Get(&employee, sqlStr, id)
	if err != nil {
		return nil, err
	}
	return &employee, nil
}

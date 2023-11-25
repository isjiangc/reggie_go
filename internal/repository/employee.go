package repository

import (
	"context"
	"reggie_go/internal/model"
)

type EmployeeRepository interface {
	GetByUsername(ctx context.Context, username string) (*model.Employee, error)
	CreateEmployee(ctx context.Context, employee *model.Employee) (int64, error)
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
	//TODO implement me
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

package repository

import (
	"context"

	"reggie_go/internal/model"
)

type UsersRepository interface {
	GetUserByPhone(ctx context.Context, phone string) (*model.Users, error)
	SaveUser(ctx context.Context, users model.Users) (int64, error)
	FirstCountByPhone(ctx context.Context, phone string) (int, error)
}

func NewUsersRepository(repository *Repository) UsersRepository {
	return &usersRepository{
		Repository: repository,
	}
}

type usersRepository struct {
	*Repository
}

func (u *usersRepository) GetUserByPhone(ctx context.Context, phone string) (*model.Users, error) {
	sqlStr := `
		SELECT
			id,
			COALESCE(name, '') as name,
			phone,
			COALESCE(sex, '') as sex,
			COALESCE(id_number, '') as id_number,
			COALESCE(avatar, '') as avatar,
			COALESCE(status, '') as status
		FROM
			user
		WHERE
			1 = 1
			AND phone = ?`
	user := model.Users{}
	err := u.db2.Get(&user, sqlStr, phone)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *usersRepository) SaveUser(ctx context.Context, users model.Users) (int64, error) {
	sqlStr := `INSERT INTO user (name, phone, sex, id_number, avatar, status) VALUES(?, ?, ?, ?, ?, ?);`
	ret, err := u.db2.Exec(sqlStr, users.Name, users.Phone, users.Sex, users.IdNumber, users.Avatar, users.Status)
	if err != nil {
		return -1, err
	}
	theId, err := ret.LastInsertId()
	if err != nil {
		return -1, err
	}
	return theId, nil
}

func (u *usersRepository) FirstCountByPhone(ctx context.Context, phone string) (int, error) {
	var count int
	sqlStr := `SELECT COUNT(*)  FROM user WHERE 1=1 AND phone = ?`
	err := u.db2.Get(&count, sqlStr, phone)
	if err != nil {
		return -1, err
	}
	return count, nil
}

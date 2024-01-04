package repository

import (
	"context"
	"reggie_go/internal/model"
	"time"
)

type AddressbookRepository interface {
	QueryAddressById(ctx context.Context, id int64) (*model.AddressBook, error)
	UpdataAddressIsDefault(ctx context.Context, userId int64, Id int64, updateTime time.Time, updateUser int64) (int, error)
	FirstByUserId(ctx context.Context, userId int64) ([]model.AddressBook, error)
}

func NewAddressbookRepository(repository *Repository) AddressbookRepository {
	return &addressbookRepository{
		Repository: repository,
	}
}

type addressbookRepository struct {
	*Repository
}

func (s *addressbookRepository) QueryAddressById(ctx context.Context, id int64) (*model.AddressBook, error) {
	sqlStr := `
	SELECT
	    id,
		user_id,
		consignee,
		sex,
		phone,
		COALESCE(province_code, '') as province_code,
		COALESCE(province_name, '') as province_name,
		COALESCE(city_code, '') as city_code,
		COALESCE(city_name, '') as city_name,
		COALESCE(district_code, '') as district_code,
		COALESCE(district_name, '') as district_name,
		detail,
		label,
		is_default,
		create_time,
		update_time,
		create_user,
		update_user,
		is_deleted
	FROM
	address_book
	WHERE
	1 = 1
	AND id = ?`
	var addressBook model.AddressBook
	err := s.db2.Get(&addressBook, sqlStr, id)
	if err != nil {
		return nil, err
	}
	return &addressBook, nil
}

func (s *addressbookRepository) UpdataAddressIsDefault(ctx context.Context, userId int64, Id int64, updateTime time.Time, updateUser int64) (int, error) {
	tx, err := s.db2.Begin()
	if err != nil {
		return -1, err
	}
	defer func() { _ = tx.Rollback() }()
	sqlStr := `
		UPDATE
			address_book
		SET
			is_default = 0
		WHERE
			1 = 1
			AND user_id = ?`
	ret, err := s.db2.Exec(sqlStr, userId)
	if err != nil {
		return -1, err
	}
	affected, err := ret.RowsAffected()
	if err != nil || affected < 0 {
		return -1, nil
	}
	sqlStr2 := `
		UPDATE
			address_book
		SET
			is_default = 1,
			update_time = ?,
			update_user = ?
		WHERE
			1 = 1
			AND id = ?`
	ret2, err := s.db2.Exec(sqlStr2, updateTime, updateUser, Id)
	if err != nil {
		return -1, err
	}
	rowsAffected, err := ret2.RowsAffected()
	if err != nil {
		return -1, nil
	}
	return int(rowsAffected), nil
}

func (a *addressbookRepository) FirstByUserId(ctx context.Context, userId int64) ([]model.AddressBook, error) {
	sqlStr := `
		SELECT
			id,
			user_id,
			consignee,
			sex,
			phone,
			COALESCE(province_code, '') as province_code,
	        COALESCE(province_name, '') as province_name,
	        COALESCE(city_code, '') as city_code,
			COALESCE(city_name, '') as city_name,
			COALESCE(district_code, '') as district_code,
			COALESCE(district_name, '') as district_name,
			detail,
			label,
			is_default,
			create_time,
			update_time,
			create_user,
			update_user,
			is_deleted
		FROM
			address_book
		WHERE
			1 = 1
			AND user_id = ?
		ORDER BY 
			update_time DESC`
	var addressBooks []model.AddressBook
	err := a.db2.Select(&addressBooks, sqlStr, userId)
	if err != nil {
		return nil, err
	}
	return addressBooks, nil
}

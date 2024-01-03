package repository

import (
	"context"
	"reggie_go/internal/model"
)

type AddressbookRepository interface {
	FirstById(ctx context.Context, id int64) ([]model.AddressBook, error)
}

func NewAddressbookRepository(repository *Repository) AddressbookRepository {
	return &addressbookRepository{
		Repository: repository,
	}
}

type addressbookRepository struct {
	*Repository
}

func (a *addressbookRepository) FirstById(ctx context.Context, id int64) ([]model.AddressBook, error) {
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
	err := a.db2.Select(&addressBooks, sqlStr, id)
	if err != nil {
		return nil, err
	}
	return addressBooks, nil
}

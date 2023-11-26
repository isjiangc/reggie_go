package repository

import (
	"context"
	"reggie_go/internal/model"
)

type CategoryRepository interface {
	Save(ctx context.Context, category *model.Category) (int64, error)
	FirstByName(ctx context.Context, name string) (*model.Category, error)
}

func NewCategoryRepository(repository *Repository) CategoryRepository {
	return &categoryRepository{
		Repository: repository,
	}
}

type categoryRepository struct {
	*Repository
}

func (c *categoryRepository) Save(ctx context.Context, category *model.Category) (int64, error) {
	sqlStr := `INSERT INTO category (type, name, sort, create_time, update_time, create_user, update_user) 
    							VALUES(?, ?, ?, ?, ?, ?, ?);`
	ret, err := c.db2.Exec(sqlStr, category.Type, category.Name, category.Sort, category.CreateTime,
		category.UpdateTime, category.CreateUser, category.UpdateUser)
	if err != nil {
		return 0, err
	}
	theId, err := ret.LastInsertId()
	if err != nil {
		return 0, err
	}
	return theId, nil
}

func (c *categoryRepository) FirstByName(ctx context.Context, name string) (*model.Category, error) {
	sqlStr := `
			SELECT
				id,
				type,
				name,
				sort,
				create_time,
				update_time,
				create_user,
				update_user
			FROM
				category
			WHERE
				1 = 1
				AND name = ?`
	var category = model.Category{}
	err := c.db2.Get(&category, sqlStr, name)
	if err != nil {
		return nil, err
	}
	return &category, nil
}

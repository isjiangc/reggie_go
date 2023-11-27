package repository

import (
	"context"
	"reggie_go/internal/model"
)

type CategoryRepository interface {
	DeleteCategory(ctx context.Context, id int64) (int64, error)
	GetCount(ctx context.Context) (int, error)
	GetByPage(ctx context.Context, page int, size int) ([]*model.Category, error)
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

func (c *categoryRepository) DeleteCategory(ctx context.Context, id int64) (int64, error) {
	delStr := `
		DELETE
		FROM
			category
		WHERE
			1 = 1
			and id = ?;`
	ret, err := c.db2.Exec(delStr, id)
	if err != nil {
		return 0, err
	}
	affected, err := ret.RowsAffected()
	if err != nil {
		return 0, err
	}
	return affected, nil

}

func (c *categoryRepository) GetCount(ctx context.Context) (int, error) {
	sqlStr := `
		SELECT
			COUNT(*)
		FROM
			category
		WHERE
			1 = 1`
	var count int
	err := c.db2.Get(&count, sqlStr)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (c *categoryRepository) GetByPage(ctx context.Context, page int, size int) ([]*model.Category, error) {
	selStr := `
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
			ORDER BY update_time desc LIMIT ? OFFSET ?;`
	var category []*model.Category
	offset := (page - 1) * size
	err := c.db2.Select(&category, selStr, size, offset)
	if err != nil {
		return nil, err
	}
	return category, err

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

package repository

import (
	"context"
	"fmt"
	"reggie_go/internal/model"
	"time"
)

type DishRepository interface {
	GetDishByPage(ctx context.Context, page int, size int, name string) ([]*DishDto, error)
	GetDishCountByName(ctx context.Context, name string) (int, error)
	SaveDishWithFlavor(ctx context.Context, dish model.Dish, flavors []model.DishFlavor) (int64, error)
	QueryCountByCategoryId(ctx context.Context, id int64) (*int, error)
}

func NewDishRepository(repository *Repository) DishRepository {
	return &dishRepository{
		Repository: repository,
	}
}

// 菜品管理
type DishDto struct {
	Id           int64     `db:"id"`           //主键
	Name         string    `db:"name"`         //菜品名称
	CategoryId   int64     `db:"category_id"`  //菜品分类id
	Price        float64   `db:"price"`        //菜品价格
	Code         string    `db:"code"`         //商品码
	Image        string    `db:"image"`        //图片
	Description  string    `db:"description"`  //描述信息
	Status       int       `db:"status"`       //0 停售 1 起售
	Sort         int       `db:"sort"`         //顺序
	CreateTime   time.Time `db:"create_time"`  //创建时间
	UpdateTime   time.Time `db:"update_time"`  //更新时间
	CreateUser   int64     `db:"create_user"`  //创建人
	UpdateUser   int64     `db:"update_user"`  //修改人
	CategoryName string    `db:"categoryName"` //分类名称
}

type dishRepository struct {
	*Repository
}

func (d *dishRepository) GetDishByPage(ctx context.Context, page int, size int, name string) ([]*DishDto, error) {
	sqlStr := `
	select
		 t1.id,
		 t1.name,
		 t1.category_id,
		 t1.price,
		 t1.code,
		 t1.image,
		 t1.description,
		 t1.status,
		 t1.sort,
		 t1.create_time,
		 t1.update_time,
		 t1.create_user,
		 t1.update_user,
		 t2.name as categoryName
   from
		 dish t1
   left join category t2
		 on (t1.category_id = t2.id and  t1.is_deleted = 0)
		where
		 1 = 1 `
	if name != "" {
		sqlStr += fmt.Sprintf(`and t1.name like '%%%s%%'`, name)
	}
	var dishDto []*DishDto
	offset := (page - 1) * size
	sqlStr += fmt.Sprintf(`order by t1.update_time desc limit %d offset %d`, size, offset)
	err := d.db2.Select(&dishDto, sqlStr)
	if err != nil {
		return nil, err
	}
	return dishDto, nil
}

func (d *dishRepository) GetDishCountByName(ctx context.Context, name string) (int, error) {
	var count int
	sqlStr := `
	SELECT
		COUNT(*)
	FROM
		dish
	WHERE
		1 = 1`
	if name != "" {
		sqlStr += fmt.Sprintf(` AND name LIKE '%%%s%%'`, name)
	}
	err := d.db2.Get(&count, sqlStr)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (d *dishRepository) SaveDishWithFlavor(ctx context.Context, dish model.Dish, flavors []model.DishFlavor) (int64, error) {
	tx, err := d.db2.Begin()
	if err != nil {
		return 0, err
	}
	defer func() { _ = tx.Rollback() }()
	sqlStr := `INSERT INTO dish (name, category_id, price, code, image, description,status, sort,
								create_time, update_time, create_user, update_user, is_deleted) 
								VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`
	ret, err := tx.Exec(sqlStr, dish.Name, dish.CategoryId, dish.Price, dish.Code, dish.Image, dish.Description, dish.Status, dish.Sort,
		dish.CreateTime, dish.UpdateTime, dish.CreateUser, dish.UpdateUser, dish.IsDeleted)
	if err != nil {
		return 0, nil
	}
	dishId, err := ret.LastInsertId()
	if err != nil {
		return 0, err
	}
	if dishId <= 0 {
		return 0, err
	}
	sqlStr2 := `INSERT INTO dish_flavor (dish_id, name, value, create_time, update_time,create_user, update_user, is_deleted) 
         VALUES(?, ?, ?, ?, ?, ?, ?,?)`
	for _, v := range flavors {
		_, err := tx.Exec(sqlStr2, dishId, v.Name, v.Value, dish.CreateTime, dish.UpdateTime, dish.CreateUser, dish.UpdateUser, 0)
		if err != nil {
			return 0, err
		}
	}
	return dishId, tx.Commit()
}

func (d *dishRepository) QueryCountByCategoryId(ctx context.Context, id int64) (*int, error) {
	selStr := `
			SELECT
				COUNT(*)
			FROM
				dish
			WHERE
				1 = 1
				AND category_id = ?`
	var count int
	err := d.db2.Get(&count, selStr, id)
	if err != nil {
		return nil, err
	}
	return &count, nil

}

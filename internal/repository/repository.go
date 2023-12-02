package repository

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"moul.io/zapgorm2"
	"reggie_go/pkg/log"
	"time"
)

const ctxTxKey = "TxKey"

type Repository struct {
	db     *gorm.DB
	rdb    *redis.Client
	logger *log.Logger
	db2    *sqlx.DB
}

func NewRepository(db *gorm.DB, rdb *redis.Client, logger *log.Logger, db2 *sqlx.DB) *Repository {
	return &Repository{
		db:     db,
		rdb:    rdb,
		logger: logger,
		db2:    db2,
	}
}

type Transaction interface {
	Transaction(ctx context.Context, fn func(ctx context.Context) error) error
}

type SqlxTransaction interface {
	SqlxTran(ctx context.Context, fn func(ctx context.Context) error) error
}

func NewTransaction(r *Repository) Transaction {
	return r
}
func NewSqlxTransaction(r *Repository) SqlxTransaction {
	return r
}

// DB return tx
// If you need to create a Transaction, you must call DB(ctx) and Transaction(ctx,fn)
func (r *Repository) DB(ctx context.Context) *gorm.DB {
	v := ctx.Value(ctxTxKey)
	if v != nil {
		if tx, ok := v.(*gorm.DB); ok {
			return tx
		}
	}
	return r.db.WithContext(ctx)
}

func (r *Repository) DB2() *sqlx.DB {
	return r.db2
}

func (r *Repository) Transaction(ctx context.Context, fn func(ctx context.Context) error) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, ctxTxKey, tx)
		return fn(ctx)
	})
}

func (r *Repository) SqlxTran(ctx context.Context, fn func(ctx context.Context) error) error {
	tx, err := r.db2.Begin()
	if err != nil {
		return err
	}
	defer func() { _ = tx.Rollback() }()
	err = fn(context.WithValue(ctx, ctxTxKey, tx))
	return tx.Commit()
}

func NewDB(conf *viper.Viper, l *log.Logger) *gorm.DB {
	logger := zapgorm2.New(l.Logger)
	logger.SetAsDefault()
	db, err := gorm.Open(mysql.Open(conf.GetString("data.mysql.user")), &gorm.Config{Logger: logger})
	if err != nil {
		panic(err)
	}
	db = db.Debug()
	return db
}
func NewRedis(conf *viper.Viper) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     conf.GetString("data.redis.addr"),
		Password: conf.GetString("data.redis.password"),
		DB:       conf.GetInt("data.redis.db"),
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("redis error: %s", err.Error()))
	}

	return rdb
}

// NewSqlxDB sqlx链接设置
func NewSqlxDB(conf *viper.Viper) *sqlx.DB {
	db2, err := sqlx.Connect("mysql", conf.GetString("data.mysql.user"))
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return nil
	}
	db2.SetMaxOpenConns(conf.GetInt("data.mysql.max_open_conns"))
	db2.SetMaxIdleConns(conf.GetInt("data.mysql.max_idle_conns"))
	return db2
}

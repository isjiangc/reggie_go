// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/spf13/viper"
	"reggie_go/internal/handler"
	"reggie_go/internal/repository"
	"reggie_go/internal/server"
	"reggie_go/internal/service"
	"reggie_go/pkg/app"
	"reggie_go/pkg/helper/sid"
	"reggie_go/pkg/jwt"
	"reggie_go/pkg/log"
	"reggie_go/pkg/server/http"
)

// Injectors from wire.go:

func NewWire(viperViper *viper.Viper, logger *log.Logger) (*app.App, func(), error) {
	jwtJWT := jwt.NewJwt(viperViper)
	handlerHandler := handler.NewHandler(logger)
	db := repository.NewDB(viperViper, logger)
	client := repository.NewRedis(viperViper)
	sqlxDB := repository.NewSqlxDB(viperViper)
	repositoryRepository := repository.NewRepository(db, client, logger, sqlxDB)
	transaction := repository.NewTransaction(repositoryRepository)
	sidSid := sid.NewSid()
	serviceService := service.NewService(transaction, logger, sidSid, jwtJWT)
	userRepository := repository.NewUserRepository(repositoryRepository)
	userService := service.NewUserService(serviceService, userRepository)
	userHandler := handler.NewUserHandler(handlerHandler, userService)
	employeeRepository := repository.NewEmployeeRepository(repositoryRepository)
	employeeService := service.NewEmployeeService(serviceService, employeeRepository)
	employeeHandler := handler.NewEmployeeHandler(handlerHandler, employeeService)
	categoryRepository := repository.NewCategoryRepository(repositoryRepository)
	dishRepository := repository.NewDishRepository(repositoryRepository)
	categoryService := service.NewCategoryService(serviceService, categoryRepository, dishRepository)
	categoryHandler := handler.NewCategoryHandler(handlerHandler, categoryService)
	dishFlavorRepository := repository.NewDishFlavorRepository(repositoryRepository)
	dishService := service.NewDishService(serviceService, dishRepository, dishFlavorRepository, repositoryRepository)
	dishHandler := handler.NewDishHandler(handlerHandler, dishService)
	setmealRepository := repository.NewSetmealRepository(repositoryRepository)
	setmealService := service.NewSetmealService(serviceService, setmealRepository)
	setmealHandler := handler.NewSetmealHandler(handlerHandler, setmealService)
	httpServer := server.NewHTTPServer(logger, viperViper, jwtJWT, userHandler, employeeHandler, categoryHandler, dishHandler, setmealHandler)
	job := server.NewJob(logger)
	appApp := newApp(httpServer, job)
	return appApp, func() {
	}, nil
}

// wire.go:

var repositorySet = wire.NewSet(repository.NewDB, repository.NewRedis, repository.NewSqlxDB, repository.NewRepository, repository.NewTransaction, repository.NewUserRepository, repository.NewEmployeeRepository, repository.NewCategoryRepository, repository.NewDishRepository, repository.NewDishFlavorRepository, repository.NewSetmealRepository, repository.NewSqlxTransaction)

var serviceSet = wire.NewSet(service.NewService, service.NewUserService, service.NewEmployeeService, service.NewCategoryService, service.NewDishService, service.NewSetmealService)

var handlerSet = wire.NewSet(handler.NewHandler, handler.NewUserHandler, handler.NewEmployeeHandler, handler.NewCategoryHandler, handler.NewDishHandler, handler.NewSetmealHandler)

var serverSet = wire.NewSet(server.NewHTTPServer, server.NewJob, server.NewTask)

// build App
func newApp(httpServer *http.Server, job *server.Job) *app.App {
	return app.NewApp(app.WithServer(httpServer, job), app.WithName("demo-server"))
}

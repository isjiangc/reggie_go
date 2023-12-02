//go:build wireinject
// +build wireinject

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

var repositorySet = wire.NewSet(
	repository.NewDB,
	repository.NewRedis,
	repository.NewSqlxDB,
	repository.NewRepository,
	repository.NewTransaction,
	repository.NewUserRepository,
	repository.NewEmployeeRepository,
	repository.NewCategoryRepository,
	repository.NewDishRepository,
	repository.NewDishFlavorRepository,
	repository.NewSqlxTransaction,
)

var serviceSet = wire.NewSet(
	service.NewService,
	service.NewUserService,
	service.NewEmployeeService,
	service.NewCategoryService,
	service.NewDishService,
)

var handlerSet = wire.NewSet(
	handler.NewHandler,
	handler.NewUserHandler,
	handler.NewEmployeeHandler,
	handler.NewCategoryHandler,
	handler.NewDishHandler,
)
var serverSet = wire.NewSet(
	server.NewHTTPServer,
	server.NewJob,
	server.NewTask,
)

// build App
func newApp(httpServer *http.Server, job *server.Job) *app.App {
	return app.NewApp(
		app.WithServer(httpServer, job),
		app.WithName("demo-server"),
	)
}

func NewWire(*viper.Viper, *log.Logger) (*app.App, func(), error) {

	panic(wire.Build(
		repositorySet,
		serviceSet,
		handlerSet,
		serverSet,
		sid.NewSid,
		jwt.NewJwt,
		newApp,
	))
}

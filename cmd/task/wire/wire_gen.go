// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/spf13/viper"
	"reggie_go/internal/server"
	"reggie_go/pkg/app"
	"reggie_go/pkg/log"
)

// Injectors from wire.go:

func NewWire(viperViper *viper.Viper, logger *log.Logger) (*app.App, func(), error) {
	task := server.NewTask(logger)
	appApp := newApp(task)
	return appApp, func() {
	}, nil
}

// wire.go:

var taskSet = wire.NewSet(server.NewTask)

// build App
func newApp(task *server.Task) *app.App {
	return app.NewApp(app.WithServer(task), app.WithName("demo-task"))
}

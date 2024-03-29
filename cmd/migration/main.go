package main

import (
	"context"
	"flag"

	"reggie_go/cmd/migration/wire"
	"reggie_go/pkg/config"
	"reggie_go/pkg/log"
)

func main() {
	envConf := flag.String("conf", "config/local.yml", "config path, eg: -conf ./config/local.yml")
	flag.Parse()
	conf := config.NewConfig(*envConf)

	logger := log.NewLog(conf)

	app, cleanup, err := wire.NewWire(conf, logger)
	defer cleanup()
	if err != nil {
		panic(err)
	}
	if err = app.Run(context.Background()); err != nil {
		panic(err)
	}
}

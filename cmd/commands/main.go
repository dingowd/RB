package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/dingowd/RB/app"
	c "github.com/dingowd/RB/internal/cache"
	"github.com/dingowd/RB/internal/config"
	"github.com/dingowd/RB/internal/logger/lrus"
	"github.com/dingowd/RB/internal/storage"
	"github.com/dingowd/RB/internal/storage/mongo"
	"github.com/dingowd/RB/server"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "./config/config.toml", "Path to configuration file")
}

func main() {
	// Init config
	conf := config.NewConfig()
	if _, err := toml.DecodeFile(configFile, &conf); err != nil {
		fmt.Fprintln(os.Stdout, "ошибка чтения toml файла "+err.Error()+", установка параметров по умолчанию")
		conf = config.Default()
	}
	// init logger
	logg := lrus.New()
	logg.SetLevel(conf.Logger.Level)
	file, err := os.OpenFile(conf.Logger.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o666)
	if err == nil {
		logg.SetOutput(file)
		defer file.Close()
	} else {
		logg.SetOutput(os.Stdout)
	}
	// init storage
	var store storage.Storage
	store = mongo.New(logg, conf)
	if err := store.Connect(context.Background(), conf.DB.DSN); err != nil {
		logg.Error("failed to connect database" + err.Error())
		os.Exit(1) // nolint:gocritic
	}
	defer func() {
		logg.Info("Закрытие соединения с БД")
		store.Close()
	}()

	// Init cache
	var cache c.CacheInterface
	stop := make(chan struct{})
	cache = c.NewCache(logg, store, conf.CacheTick)
	go cache.WriteToCache(stop)

	// Init app
	app := app.New(logg, store, cache)

	// init http server
	server := server.NewServer(app, conf.HTTPSrv)

	exit := make(chan os.Signal)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-exit
		logg.Info("Приложение останавливается")
		server.Stop()
		stop <- struct{}{}
		logg.Info("Приложение остановлено")
		time.Sleep(5 * time.Second)
	}()

	// start http server
	server.Start()
}

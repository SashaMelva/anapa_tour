package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/SashaMelva/anapa_tour/internal/app"
	"github.com/SashaMelva/anapa_tour/internal/config"
	"github.com/SashaMelva/anapa_tour/internal/log"
	"github.com/SashaMelva/anapa_tour/internal/storage/memory"
	"github.com/SashaMelva/anapa_tour/internal/storage/sql"
	"github.com/SashaMelva/anapa_tour/server/http"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "../configfile/", "Path to configuration file")
}

func main() {
	//flag.Parse()

	// if flag.Arg(0) == "version" {
	// 	printVersion()
	// 	return
	// }
	config := config.NewConfigApp(configFile)
	log := log.NewLogger(config.Logger, "../logfile/")

	//Соединение с бд
	connection := sql.New(config.DataBase, log)

	memstorage := memory.New(connection.StorageDb)
	app := app.New(log, memstorage, config.JWTKey)

	httpServer := http.NewServer(log, app, config.HttpServer)

	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	go func() {
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		if err := httpServer.Stop(ctx); err != nil {
			log.Error("failed to stop http server: " + err.Error())
		}
	}()

	log.Info("calendar is running...")

	if err := httpServer.Start(ctx); err != nil {
		log.Error("failed to start http server: " + err.Error())
		cancel()
		os.Exit(1) //nolint:gocritic
	}
}

package main

import (
	"context"
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/skullkon/info-app/internal/repository"
	"github.com/skullkon/info-app/internal/server"
	"github.com/skullkon/info-app/internal/service"
	"github.com/skullkon/info-app/internal/transport/http"
	"github.com/skullkon/info-app/pkg/client"
	"github.com/skullkon/info-app/pkg/logging"
	"github.com/skullkon/info-app/pkg/shutdown"
	"os"
	"syscall"
	"time"
)

type Info struct {
	Id             int32     `ch:"id"`
	Ip             string    `ch:"ip"`
	TypeOfDevice   string    `ch:"type"`
	Os             string    `ch:"os"`
	OsVersion      string    `ch:"osVersion"`
	Browser        string    `ch:"browser"`
	BrowserVersion string    `ch:"browserVersion"`
	Brand          string    `ch:"brand"`
	Model          string    `ch:"model"`
	Resolution     string    `ch:"resolution"`
	Time           time.Time `ch:"time"`
}

func main() {
	logging.Init()
	logger := logging.GetLogger()
	logger.Println("logger initialized")

	ctx := clickhouse.Context(context.Background(), clickhouse.WithSettings(clickhouse.Settings{
		"max_block_size": 10,
	}), clickhouse.WithProgress(func(p *clickhouse.Progress) {
		fmt.Println("progress: ", p)
	}))
	ch, err := client.NewClient(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}

	repos := repository.NewRepositories(ch, &logger)

	deps := service.Deps{Repos: repos, Logger: logger}

	services := service.NewServices(deps)

	handlers := http.NewHandler(services, logger)

	srv := server.NewServer(handlers.Init())

	// repos.Information.SeedData(ctx) если нужно заполнить базу рандомной инфой

	err = srv.Run()
	if err != nil {
		fmt.Println(err)
		return
	}

	go shutdown.Graceful([]os.Signal{syscall.SIGABRT, syscall.SIGQUIT, syscall.SIGHUP, os.Interrupt, syscall.SIGTERM},
		srv)

	if err := srv.Stop(ctx); err != nil {
		logger.Errorf("failed to stop server: %v", err)
	}

	if err := ch.Close(); err != nil {
		logger.Error(err.Error())
	}

}

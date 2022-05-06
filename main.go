package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
)

func main() {
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{"127.0.0.1:9000"},
		Auth: clickhouse.Auth{
			Database: "information",
			Username: "default",
			Password: "",
		},
		//Debug:           true,
		DialTimeout:     time.Second,
		MaxOpenConns:    10,
		MaxIdleConns:    5,
		ConnMaxLifetime: time.Hour,
		Compression: &clickhouse.Compression{
			Method: clickhouse.CompressionLZ4,
		},
	})
	if err != nil {
		return
	}
	ctx := clickhouse.Context(context.Background(), clickhouse.WithSettings(clickhouse.Settings{
		"max_block_size": 10,
	}), clickhouse.WithProgress(func(p *clickhouse.Progress) {
		fmt.Println("progress: ", p)
	}))
	if err := conn.Ping(ctx); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			fmt.Printf("Catch exception [%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		}
		return
	}

	var result []struct {
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

	if err := conn.Select(ctx, &result, "SELECT id, ip, type FROM info"); err != nil {
		log.Println(err)
		return
	}
	fmt.Println(result)
}

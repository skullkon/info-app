package main

import (
	"log"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
	_ "github.com/mailru/go-clickhouse/v2"
)

func main() {
	conn := clickhouse.OpenDB(&clickhouse.Options{

		Addr: []string{"127.0.0.1:9000"},
		Auth: clickhouse.Auth{
			Database: "default",
			Username: "default",
			Password: "",
		},
		Settings: clickhouse.Settings{
			"max_execution_time": 60,
		},
		DialTimeout: 10 * time.Second,
		Compression: &clickhouse.Compression{
			clickhouse.CompressionLZ4,
		},
		Debug: true,
	})
	log.Print(conn.Ping())
}

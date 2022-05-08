package client

import (
	"context"
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"time"
)

type Client interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) error
	PrepareBatch(ctx context.Context, sql string) (driver.Batch, error)
	Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	Query(ctx context.Context, query string, args ...interface{}) (rows driver.Rows, err error)
}

func NewClient(ctx context.Context, cfg *ClickHouseConfig) (driver.Conn, error) {
	connAddr := cfg.Host + ":" + cfg.Port
	fmt.Println(connAddr)
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{connAddr},
		Auth: clickhouse.Auth{
			Database: cfg.Dbname,
			Username: cfg.Username,
			Password: cfg.Password,
		},
		DialTimeout:     time.Second,
		MaxOpenConns:    10,
		MaxIdleConns:    5,
		ConnMaxLifetime: time.Hour,
		Compression: &clickhouse.Compression{
			Method: clickhouse.CompressionLZ4,
		},
	})
	if err != nil {
		return nil, err
	}

	if err := conn.Ping(ctx); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			fmt.Printf("Catch exception [%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		}
		return nil, err
	}

	return conn, nil
}

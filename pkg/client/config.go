package client

import (
	"fmt"
	"os"
)

type ClickHouseConfig struct {
	Host     string
	Port     string
	Dbname   string
	Username string
	Password string
}

func NewConfig() *ClickHouseConfig {
	return &ClickHouseConfig{}
}

func (c *ClickHouseConfig) Init() error {
	c.Host = os.Getenv("DB_HOST")
	c.Port = os.Getenv("DB_PORT")
	c.Dbname = os.Getenv("DB_NAME")
	c.Username = os.Getenv("DB_USERNAME")
	c.Password = os.Getenv("DB_PASSWORD")
	fmt.Println(c.Username)
	return nil
}

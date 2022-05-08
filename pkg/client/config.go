package client

import (
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
	return &ClickHouseConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Dbname:   os.Getenv("DB_NAME"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
	}
}

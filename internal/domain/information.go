package domain

import (
	"github.com/google/uuid"
	"time"
)

type Info struct {
	Id             uuid.UUID `ch:"id"`
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

type ClientInfo struct {
	Id         string `json:"id"`
	Ip         string `json:"ip"`
	Resolution string `json:"resolution"`
}

package domain

import "time"

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

type ClientInfo struct {
	Id         int32  `json:"id"`
	Ip         string `json:"ip"`
	Resolution string `json:"resolution"`
}

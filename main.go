package main

import (
	"context"
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ua-parser/uap-go/uaparser"
	"log"
	"math/rand"
	"time"

	browser "github.com/EDDYCJY/fake-useragent"
)

type Info struct {
	Id             int32     `ch:"id"`
	Ip             string    `ch:"ip"`
	TypeOfDevice   string    `ch:"type"`
	Os             string    `ch:"os"`
	OsVersion      string    `ch:"osVersion"`
	Browser        string    `ch:"browser"`
	BrowserVersion string    `ch:"browserVersion"`
	Brand          string    `ch:"Brand"`
	Model          string    `ch:"model"`
	Resolution     string    `ch:"resolution"`
	Time           time.Time `ch:"time"`
}

type Browser struct {
	Name    string
	Version string
}

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

	var result []Info

	if err := conn.Select(ctx, &result, "SELECT id, ip, type FROM info"); err != nil {
		log.Println(err)
		return
	}

	parser, err := uaparser.New("./regexes.yaml")
	if err != nil {
		log.Fatal(err)
	}
	var uaList []Info
	for i := 0; i < 10000; i++ {
		t := browser.Random()
		ua := parser.Parse(t)
		agent := Info{
			Id:             GenID(100),
			Ip:             GenIP(),
			TypeOfDevice:   ua.Device.Family,
			Os:             ua.Os.Family,
			OsVersion:      ua.Os.Major,
			Browser:        ua.UserAgent.Family,
			BrowserVersion: ua.UserAgent.Major + "." + ua.UserAgent.Minor,
			Brand:          ua.Device.Brand,
			Model:          ua.Device.Model,
			Resolution:     GenResolution(),
			Time:           time.Now(),
		}
		uaList = append(uaList, agent)

	}

	batch, err := conn.PrepareBatch(ctx, "INSERT INTO info")
	if err != nil {
		return
	}

	for i := 0; i < 10000; i++ {
		err := batch.Append(
			uaList[i].Id,
			uaList[i].Ip,
			uaList[i].TypeOfDevice,
			uaList[i].Os,
			uaList[i].OsVersion,
			uaList[i].Browser,
			uaList[i].BrowserVersion,
			uaList[i].Brand,
			uaList[i].Model,
			uaList[i].Resolution,
			uaList[i].Time,
		)
		if err != nil {
			fmt.Print(err)
			return
		}
	}
	err = batch.Send()
	if err != nil {
		fmt.Print(err)

		return
	}
}

func GenID(max int) int32 {
	rand.Seed(time.Now().UnixNano())
	return int32(rand.Intn(max))
}

func GenIP() string {
	ip := fmt.Sprintf("%d:%d:%d:%d", GenID(255), GenID(255), GenID(255), GenID(255))
	return ip
}

func GenResolution() string {
	resolutions := []string{
		"320×240", "352×240", "352×288",
		"400×240", "480×576", "640×240",
		"320×480", "640×360", "640×480",
		"800×480", "800×600", "848×480",
		"960×540", "1024×600", "1024×768",
		"1152×864", "1200×600", "1280×720",
		"1280×768", "1280×1024", "1440×900",
		"1400×1050", "1536×960", "1536×1024",
		"1600×900", "1600×1024", "1600×1200",
		"1680×1050", "1920×1080", "1920×1200",
		"2048×1080", "2048×1152", "2048×1536",
		"2560×1440", "2560×1600", "2560×2048",
		"3072×1620", "3200×1800", "3200×2048",
		"3200×2400", "3440×1440", "3840×2400",
	}
	return resolutions[GenID(len(resolutions))]
}

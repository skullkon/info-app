package main

import (
	"encoding/json"
	"fmt"
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
	// conn, err := clickhouse.Open(&clickhouse.Options{
	// 	Addr: []string{"127.0.0.1:9000"},
	// 	Auth: clickhouse.Auth{
	// 		Database: "information",
	// 		Username: "default",
	// 		Password: "",
	// 	},
	// 	//Debug:           true,
	// 	DialTimeout:     time.Second,
	// 	MaxOpenConns:    10,
	// 	MaxIdleConns:    5,
	// 	ConnMaxLifetime: time.Hour,
	// 	Compression: &clickhouse.Compression{
	// 		Method: clickhouse.CompressionLZ4,
	// 	},
	// })
	// if err != nil {
	// 	return
	// }
	// ctx := clickhouse.Context(context.Background(), clickhouse.WithSettings(clickhouse.Settings{
	// 	"max_block_size": 10,
	// }), clickhouse.WithProgress(func(p *clickhouse.Progress) {
	// 	fmt.Println("progress: ", p)
	// }))
	// if err := conn.Ping(ctx); err != nil {
	// 	if exception, ok := err.(*clickhouse.Exception); ok {
	// 		fmt.Printf("Catch exception [%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
	// 	}
	// 	return
	// }

	// var result []Info

	// if err := conn.Select(ctx, &result, "SELECT id, ip, type FROM info"); err != nil {
	// 	log.Println(err)
	// 	return
	// }

	// batch, err := conn.PrepareBatch(ctx, "INSERT INTO info")
	// if err != nil {
	// 	return
	// }

	// for i := 0; i < 500_000; i++ {
	// 	err := batch.Append(
	// 		uint8(42),
	// 		"ClickHouse", "Inc",
	// 		uuid.New(),
	// 		map[string]uint8{"key": 1},             // Map(String, UInt8)
	// 		[]string{"Q", "W", "E", "R", "T", "Y"}, // Array(String)
	// 		[]interface{}{ // Tuple(String, UInt8, Array(Map(String, String)))
	// 			"String Value", uint8(5), []map[string]string{
	// 				map[string]string{"key": "value"},
	// 				map[string]string{"key": "value"},
	// 				map[string]string{"key": "value"},
	// 			},
	// 		},
	// 		time.Now(),
	// 	)
	// 	if err != nil {
	// 		return err
	// 	}
	// }
	parser, err := uaparser.New("./regexes.yaml")
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < 100; i++ {
		t := browser.Random()
		ua := parser.Parse(t)
		agent := &Info{
			Id:             GenID(100),
			Ip:             GenIP(),
			Os:             ua.Os.Family,
			OsVersion:      ua.Os.Major,
			Browser:        ua.UserAgent.Family,
			BrowserVersion: ua.UserAgent.Major + "." + ua.UserAgent.Minor,
			Brand:          ua.Device.Brand,
			Model:          ua.Device.Model,
			Resolution:     GenResolution(),
			Time:           time.Time{},
		}
		b, err := json.Marshal(agent)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(b))
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
	resolution := fmt.Sprintf("%dx%d", GenID(4000), GenID(4000)/2)
	return resolution
}

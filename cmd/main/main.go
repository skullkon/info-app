package main

import (
	"errors"
	"github.com/skullkon/info-app/pkg/logging"
	"github.com/skullkon/info-app/pkg/shutdown"
	"net"
	"net/http"
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
	var listener net.Listener
	logging.Init()
	logger := logging.GetLogger()
	logger.Println("logger initialized")
	//ctx := clickhouse.Context(context.Background(), clickhouse.WithSettings(clickhouse.Settings{
	//	"max_block_size": 10,
	//}), clickhouse.WithProgress(func(p *clickhouse.Progress) {
	//	fmt.Println("progress: ", p)
	//}))
	//ch, err := client.NewClient(ctx)
	//if err != nil {
	//	return
	//}
	//parser, err := uaparser.New("./regexes.yaml")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//var uaList []Info
	//for i := 0; i < 10000; i++ {
	//	t := browser.Random()
	//	ua := parser.Parse(t)
	//	agent := Info{
	//		Id:             utils.GenID(100),
	//		Ip:             utils.GenIP(),
	//		TypeOfDevice:   ua.Device.Family,
	//		Os:             ua.Os.Family,
	//		OsVersion:      ua.Os.Major,
	//		Browser:        ua.UserAgent.Family,
	//		BrowserVersion: ua.UserAgent.Major + "." + ua.UserAgent.Minor,
	//		Brand:          ua.Device.Brand,
	//		Model:          ua.Device.Model,
	//		Resolution:     utils.GenResolution(),
	//		Time:           time.Now(),
	//	}
	//	uaList = append(uaList, agent)
	//
	//}

	//batch, err := ch.PrepareBatch(ctx, "INSERT INTO info")
	//if err != nil {
	//	return
	//}
	//
	//for i := 0; i < 10000; i++ {
	//	err := batch.Append(
	//		uaList[i].Id,
	//		uaList[i].Ip,
	//		uaList[i].TypeOfDevice,
	//		uaList[i].Os,
	//		uaList[i].OsVersion,
	//		uaList[i].Browser,
	//		uaList[i].BrowserVersion,
	//		uaList[i].Brand,
	//		uaList[i].Model,
	//		uaList[i].Resolution,
	//		uaList[i].Time,
	//	)
	//	if err != nil {
	//		fmt.Print(err)
	//		return
	//	}
	//}
	//err = batch.Send()
	//if err != nil {
	//	fmt.Print(err)
	//
	//	return
	//}
	//

	server := &http.Server{
		//Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go shutdown.Graceful([]os.Signal{syscall.SIGABRT, syscall.SIGQUIT, syscall.SIGHUP, os.Interrupt, syscall.SIGTERM},
		server)

	logger.Println("application initialized and started")

	if err := server.Serve(listener); err != nil {
		switch {
		case errors.Is(err, http.ErrServerClosed):
			logger.Warn("server shutdown")
		default:
			logger.Fatal(err)
		}
	}

}

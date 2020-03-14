package main

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"

	"github.com/riceChuang/marryme/config"
	"github.com/riceChuang/marryme/marryme"
)

func main() {
	logLevel, err := log.ParseLevel("debug")
	if err != nil {
		log.Warnf("log level invalid, set log to info level, config data:%s", "debug")
		logLevel = log.InfoLevel
	}
	log.SetLevel(logLevel)

	conf := config.Initial()

	ctx := context.Background()
	srv, err := sheets.NewService(ctx, option.WithAPIKey(conf.APIKey))
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	readRange := "Class Data!A2:E"
	resp, err := srv.Spreadsheets.Values.Get(conf.SpreadSheetID, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	_ = marryme.InitialGuests(resp.Values)
	fmt.Println(marryme.SearchGuests("狗"))

	marryme.RunHTTPServer()
}

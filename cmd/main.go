package main

import (
	"github.com/riceChuang/marryme/config"
	"github.com/riceChuang/marryme/db"
	"github.com/riceChuang/marryme/handler"
	"github.com/riceChuang/marryme/repo"
	"github.com/riceChuang/marryme/service"
	log "github.com/sirupsen/logrus"
)

func main() {
	logLevel, err := log.ParseLevel("debug")
	if err != nil {
		log.Warnf("log level invalid, set log to info level, config data:%s", "debug")
		logLevel = log.InfoLevel
	}
	log.SetLevel(logLevel)

	conf := config.Initial()
	googleSheetClient, err := db.NewGoogleSheet(conf)
	if err != nil {
		log.Panic(err)
	}
	guestRepo, err := repo.NewGuestsRepo(googleSheetClient)
	if err != nil {
		log.Panicf("New guest repo err: %v", err)
	}
	service.NewMerryMeSvc(guestRepo)
	
	
	handler.RunHTTPServer()
}

package main

import (
	"fmt"
	"github.com/riceChuang/marryme/config"
	"github.com/riceChuang/marryme/handler"
	"github.com/riceChuang/marryme/repo"
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
	guestRepo, err  := repo.NewGuestsRepo(conf)
	if err != nil {
		log.Panicf("New guest repo err: %v",err)
	}
	
	fmt.Println(guestRepo.SearchGuests("狗"))

	handler.RunHTTPServer()
}

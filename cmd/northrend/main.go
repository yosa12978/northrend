package main

import (
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/yosa12978/northrend/config"
	"github.com/yosa12978/northrend/db"
	"github.com/yosa12978/northrend/server"
)

func init() {
	config.NewJsonConfigParser("config.json").Parse()
}

func main() {
	listener, err := net.Listen("tcp", config.Config.Api.Addr)
	if err != nil {
		panic(err)
	}
	db.GetDB()
	serv := server.NewServer()
	go serv.Listen(listener)

	out := make(chan os.Signal, 1)
	signal.Notify(out, os.Interrupt, syscall.SIGTERM)
	<-out
}

package main

import (
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/yosa12978/northrend/config"
	"github.com/yosa12978/northrend/db"
	"github.com/yosa12978/northrend/server"
	"github.com/yosa12978/northrend/services"
)

func init() {
	gin.SetMode(gin.ReleaseMode)

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
	services.NewConsoleLogger("server").Info("Server started")
	out := make(chan os.Signal, 1)
	signal.Notify(out, os.Interrupt, syscall.SIGTERM)
	<-out
}

package main

import (
	"awesomeProject/library-app/global"
	server2 "awesomeProject/library-app/server"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("[server] Starting app...")

	server := server2.Server{}
	server.InitializeConfig(global.DefaultConfigurationFilePath)
	server.Initialize()
	server.StartWebServer()
}

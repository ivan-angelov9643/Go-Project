package main

import (
	"awesomeProject/library-app/global"
	server2 "awesomeProject/library-app/server"
	"context"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("[server] Starting app...")

	ctx, cancel := context.WithCancel(context.Background())

	server := server2.Server{}
	server.InitializeConfig(global.DefaultConfigurationFilePath)
	server.Initialize(ctx)
	server.StartWebServer()

	cancel()
}

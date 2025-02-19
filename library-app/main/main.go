package main

import (
	"context"
	"github.com/ivan-angelov9643/go-project/library-app/global"
	server2 "github.com/ivan-angelov9643/go-project/library-app/server"
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

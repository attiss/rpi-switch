package main

import (
	"flag"
	"os"
	"os/signal"

	"github.com/attiss/rpi-switch/config"
	"github.com/attiss/rpi-switch/server"
	"go.uber.org/zap"
)

var (
	configFile = flag.String("config", "config.yaml", "Path for the YAML configuration file.")
)

func main() {
	flag.Parse()

	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	config, err := config.GetConfig(*configFile)
	if err != nil {
		logger.Error("failed to get config", zap.Error(err))
		panic(err)
	}

	shutdownChannel := make(chan os.Signal, 1)
	signal.Notify(shutdownChannel, os.Interrupt)

	server, err := server.NewAPIServer(config, logger)
	if err != nil {
		logger.Error("failed to create api server", zap.Error(err))
		panic(err)
	}

	server.Start(shutdownChannel)
}

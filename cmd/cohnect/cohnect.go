package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/davidaburns/cohnect/config"
	"github.com/davidaburns/cohnect/internal/cache"
	"github.com/davidaburns/cohnect/internal/logger"
	"github.com/davidaburns/cohnect/internal/server"
)

func main() {
	defer os.Exit(0)
	sig := make(chan os.Signal,2)
	done := make(chan bool, 1)

	signal.Notify(sig, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	build := config.NewBuildInfo();
	serverConfig, err := config.LoadConfigFile("./config.yaml")
	if err != nil {
		log.Fatalf("error loading config file: %s", err)
	}

	log := logger.CreateNew(serverConfig.Logger.Level, logger.ConsoleLog)
	cache := cache.NewInMemoryCache()

	serverListener := server.CreateNew(serverConfig, log, cache)

	log.Infof("Starting: %s", build.ToString())
	if err = serverListener.Start(done); err != nil {
		log.Errorf("Failed to startup udp server: %s", err)
		os.Exit(1)
	}

	// Wait for some signal from the os to shutdown the server
	<-sig
	done <- true

	log.Info("Server shutting down")

	close(done)
	close(sig)
}
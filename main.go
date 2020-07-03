package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/Floor-Gang/prevent-link-redirection/internal"
)

const (
	configPath = "./config.yml"
)

func main() {
	err := internal.Start(configPath)

	if err != nil {
		panic(err)
	}

	keepalive()
}

func keepalive() {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}

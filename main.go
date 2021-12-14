package main

import (
	"fmt"
	"github.com/gaecoli/config"
	"github.com/gaecoli/utils/logger"
	"github.com/gaecoli/tcp"
	"os"
)

var banner = "Hello GuYu"

var defaultProperties = &config.ServerProperties{
	Bind: "0.0.0.0",
	Port: 6666,
	AppendOnly: false,
	AppendFilename: "",
	MaxClients: 1000,
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	return err == nil && !info.IsDir()
}

type Handler struct {
	a string
}

func TestHandler() *tcp.EchoHandler {
	a := "This is a test"
	print(a)
	return nil
}

func main() {
	print(banner)
	logger.Setup(&logger.Settings{
		Path: "logs",
		Name: "go-redis",
		Ext: "log",
		TimeFormat: "2021-12-11",
	})

	configFilename := os.Getenv("CONFIG")
	if configFilename == "" {
		if fileExists("go-redis.conf") {
			config.SetupConfig("go-redis.conf")
		} else {
			config.Properties = defaultProperties
		}
	} else {
		config.SetupConfig(configFilename)
	}

	err := tcp.ListenAndServeWithSignal(&tcp.Config{
		Address: fmt.Sprintf("%s:%d", config.Properties.Bind, config.Properties.Port),
	}, TestHandler())
	if err != nil {
		logger.Error(err)
	}

}

package main

import (
	"flag"

	"github.com/Abhi-Harsha/transaction-service/app/server"
)

const (
	configFileKey     = "configFile"
	defaultConfigFile = ""
	configFileUsage   = "config file path"
)

func main() {
	var configFile string
	flag.StringVar(&configFile, configFileKey, defaultConfigFile, configFileUsage)
	flag.Parse()

	server.InitializeApp(configFile)
}

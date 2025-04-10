package main

import (
	"flag"
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
}

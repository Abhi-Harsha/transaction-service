package config

import "github.com/spf13/viper"

type ServerConfig struct {
	Host string
	Port int
}

type AppConfig struct {
	ServerConfig ServerConfig
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func NewAppConfig(config string) AppConfig {
	err := loadConfigFile(config)
	check(err)
	return parserConfigFile()
}

func loadConfigFile(config string) error {
	viper.AddConfigPath(config)
	return viper.ReadInConfig()
}

func parserConfigFile() AppConfig {
	var appConfig AppConfig
	err := viper.Unmarshal(&appConfig)
	check(err)
	return appConfig
}

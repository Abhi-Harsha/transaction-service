package config

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ServerConfig struct {
	Host    string
	Port    int
	GinMode string
}

type Logger struct {
	Level string
}

func (l Logger) GetZapLevel() zapcore.Level {
	levelMap := make(map[string]zapcore.Level)
	levelMap["debug"] = zap.DebugLevel
	levelMap["info"] = zap.InfoLevel
	levelMap["warn"] = zap.WarnLevel
	levelMap["error"] = zap.ErrorLevel

	return levelMap[l.Level]
}

type AppConfig struct {
	Server ServerConfig
	Logger Logger
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
	viper.SetConfigType("yaml")
	viper.SetConfigFile(config)
	viper.AutomaticEnv()
	return viper.ReadInConfig()
}

func parserConfigFile() AppConfig {
	var appConfig AppConfig
	err := viper.Unmarshal(&appConfig)
	check(err)
	return appConfig
}

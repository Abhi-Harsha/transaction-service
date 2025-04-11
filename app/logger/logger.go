package logger

import (
	"os"
	"sync"

	"github.com/Abhi-Harsha/transaction-service/app/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var once sync.Once
var logger *zap.Logger

func InitLogger(logConfig config.Logger) {
	once.Do(func() {
		logger = configureLogger(logConfig.GetZapLevel())
	})
}

func configureLogger(level zapcore.Level) *zap.Logger {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:    "timestamp",
		LevelKey:   "level",
		MessageKey: "msg",
	}
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(os.Stdout),
		zap.NewAtomicLevelAt(level),
	)
	return zap.New(core)
}

func GetLogger() *zap.Logger {
	if logger == nil {
		lConfig := config.Logger{Level: "info"}
		logger = configureLogger(lConfig.GetZapLevel())
	}

	return logger
}

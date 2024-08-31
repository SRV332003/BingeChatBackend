package config

import (
	"HangAroundBackend/logger"
	"os"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

var configLogger *zap.Logger

func GetEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		configLogger.Panic("Please set the environment variable " + key)
	}
	return val
}

func init() {
	configLogger = logger.GetLoggerWithName("config")
	LoadEnvs()
}

func LoadEnvs() (err error) {
	err = godotenv.Load()
	return
}

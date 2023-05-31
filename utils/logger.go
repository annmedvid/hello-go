package utils

import (
	"encoding/json"

	"go.uber.org/zap"
)

var Logger *zap.Logger

func getLogger() *zap.Logger {
	rawJSON := []byte(`{
	  "level": "debug",
	  "encoding": "json",
	  "outputPaths": ["stdout", "/tmp/hello-go-logs.log"],
	  "errorOutputPaths": ["stderr", "/tmp/hello-go-errors.log"],
	  "encoderConfig": {
	    "messageKey": "message",
	    "levelKey": "level",
	    "levelEncoder": "lowercase"
	  }
	}`)

	var cfg zap.Config
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		panic(err)
	}
	logger := zap.Must(cfg.Build())

	return logger
}

func InitializeLogger() {
	Logger = getLogger()
	defer Logger.Sync()

	Logger.Info("Logger initialized")
}

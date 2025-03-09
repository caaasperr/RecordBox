package utils

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger = log.New()

func InitializeLogger() {
	Logger.SetFormatter(&log.JSONFormatter{})
	file := &lumberjack.Logger{
		Filename:   "Recordbox.json",
		MaxSize:    10,
		MaxBackups: 3,
		MaxAge:     7,
		Compress:   true,
	}
	Logger.SetOutput(file)
	Logger.SetLevel(log.TraceLevel)
}

package main

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	// Configure log rotation using Lumberjack
	log := logrus.New()
	log.SetOutput(&lumberjack.Logger{
		Filename:   "server.log", // Log file name
		MaxSize:    2,            // Max size in MB before rotation
		MaxBackups: 3,            // Keep last 3 old log files
		MaxAge:     7,            // Keep logs for 7 days
		Compress:   true,         // Compress old log files
	})

	// Logging messages
	for i := 1; i <= 100000; i++ {
		log.Info("This is log message ", i)
	}
}

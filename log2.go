package main

import (
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()

	log.Info("this is info message")
	log.Warn("this is warn message")
	log.Error("this is error message")

}

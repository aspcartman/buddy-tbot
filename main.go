package main

import (
	"time"

	"github.com/aspcartman/buddy-tbot/api"
	"github.com/aspcartman/buddy-tbot/bot"
	"github.com/aspcartman/buddy-tbot/e/elogrus"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.StandardLogger()
	logger.Formatter = &logrus.TextFormatter{DisableTimestamp: true}
	elogrus.AddLogger(logger)

	logrus.Info("Starting")
	b := bot.Run()

	api.ListenForAPIEvents(func(str string) {
		b.SendNotification(str)
	})

	time.Sleep(1 * time.Hour)
}

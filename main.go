package main

import (
	"time"

	"github.com/aspcartman/buddy-tbot/api"
	"github.com/aspcartman/buddy-tbot/bot"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("Starting")
	b := bot.Run()

	api.ListenForAPIEvents(func() {
		b.SendNotification()
	})

	time.Sleep(1 * time.Hour)
}

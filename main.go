package main

import (
	"time"

	"github.com/aspcartman/buddy-tbot/api"
	"github.com/aspcartman/buddy-tbot/bot"
	_ "github.com/aspcartman/buddy-tbot/env"
)

func main() {
	b := bot.Run()

	api.ListenForAPIEvents(func(str string) {
		b.SendNotification(str)
	})

	time.Sleep(1 * time.Hour)
}

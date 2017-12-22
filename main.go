package main

import (
	"time"

	"github.com/aspcartman/buddy-tbot/api"
	"github.com/aspcartman/buddy-tbot/bot"
	_ "github.com/aspcartman/buddy-tbot/env"
)

func main() {
	b := bot.Run()
	b.SendMessage("I'm back, boss")

	api.ListenForAPIEvents(func(str string) {
		b.SendMessage("I have something for you, man!")
		b.SendNotification(str)
	})

	time.Sleep(1 * time.Hour)
}

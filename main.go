package main

import (
	"time"

	"github.com/aspcartman/buddy-tbot/api"
	"github.com/aspcartman/buddy-tbot/bot"
	"github.com/aspcartman/buddy-tbot/env"
	_ "github.com/aspcartman/buddy-tbot/env"
)

func main() {
	env.Log.Info("Launching")
	tg := bot.Run()
	api.ListenForAPIEvents(tg)
	tg.SendMessage("I'm here, boss.")
	time.Sleep(1000 * time.Hour)
}

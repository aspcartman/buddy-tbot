package main

import (
	"time"

	"github.com/aspcartman/buddy-tbot/api"
	"github.com/aspcartman/buddy-tbot/bot"
	"github.com/aspcartman/buddy-tbot/env"
)

func main() {
	env.Log.Info("Launching")
	api.ListenForAPIEvents(bot.Run())
	time.Sleep(1000 * time.Hour)
}

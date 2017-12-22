package main

import (
	"time"

	"github.com/aspcartman/buddy-tbot/api"
	"github.com/aspcartman/buddy-tbot/bot"
)

func main() {
	api.ListenForAPIEvents(bot.Run())
	time.Sleep(1000 * time.Hour)
}

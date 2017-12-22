package main

import (
	"github.com/aspcartman/buddy-tbot/api/buddy"
	"github.com/aspcartman/buddy-tbot/bot"
	_ "github.com/aspcartman/buddy-tbot/env"
)

func main() {
	b := bot.Run()

	event := buddy.WebHookEvent{}
	event.Workspace.Name = "lol"
	event.Project.DisplayName = "TestProject"
	event.Project.HTMLURL = "https://buddy.aspc.me/aspc/buddy-tbot"
	event.Project.Status = "waiting"
	event.Invoker.Name = "yourmum"
	event.Invoker.HTMLURL = "http://ya.ru"
	event.Invoker.Email = "asp@gmail.com"
	b.HandleBuddyEvent(&event)
}

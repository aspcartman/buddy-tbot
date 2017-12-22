package bot

import (
	"time"

	"github.com/aspcartman/buddy-tbot/e"
	"github.com/aspcartman/buddy-tbot/env"
	"github.com/sirupsen/logrus"
	"gopkg.in/tucnak/telebot.v2"
)

var logger logrus.FieldLogger

func init() {
	logger = env.Log
}

type Bot struct {
	tg *telebot.Bot
}

func Run() *Bot {
	logger.Info("Starting telegram Bot")

	telegramBot, err := telebot.NewBot(telebot.Settings{
		Token:  "487229393:AAEWArHxDgx2tRaQmUJu7qLNRl4hRobPdsI",
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		e.Throw("Failed to authorize Bot a instance", err)
	}

	bot := &Bot{telegramBot}
	telegramBot.Handle(telebot.OnText, bot.handleMessage)

	go bot.runloop()

	return &Bot{telegramBot}
}

func (b *Bot) runloop() {
	defer e.Catch(func(e *e.Exception) {
		logger.Error("Telegram Bot runloop panic")
		go b.runloop()
	})

	b.tg.Start()
}

func (b Bot) handleMessage(m *telebot.Message) {
	logger.Info("Here comes the owner", m.Sender, m.Chat.ID)
	b.tg.Send(m.Sender, "hello world")
}

func (b Bot) SendNotification(data string) {
	logger.WithField("length", len(data)).Info("Sending notification")
	b.SendMessage(data)
	logger.Info("Sent notification")
}

func (b Bot) SendMessage(msg string) {
	ch, err := b.tg.ChatByID("45944997")
	if err != nil {
		e.Throw("Failed getting chat by id", err)
	}

	_, err = b.tg.Send(ch, msg)
	if err != nil {
		e.Throw("Failed sending message", err)
	}
}

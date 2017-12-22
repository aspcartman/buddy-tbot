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

type TelegramNotificator interface {
	SendNotification()
}

type bot struct {
	telegram *telebot.Bot
}

func Run() *bot {
	logger.Info("Starting telegram bot")

	b, err := telebot.NewBot(telebot.Settings{
		Token:  "487229393:AAEWArHxDgx2tRaQmUJu7qLNRl4hRobPdsI",
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		e.Throw("Failed to authorize bot a instance", err)
	}

	b.Handle("/start", func(m *telebot.Message) {
		logger.Info("Here comes the owner", m.Sender, m.Chat.ID)
		b.Send(m.Sender, "hello world")
	})

	go func() {
		logger.Info("Running the bot runloop")
		b.Start()
	}()

	return &bot{b}
}

func (b bot) SendNotification(data string) {
	logger.WithField("length", len(data)).Info("Sending notification")
	ch, err := b.telegram.ChatByID("45944997")
	if err != nil {
		e.Throw("Failed getting chat by id", err)
	}

	_, err = b.telegram.Send(ch, data)
	if err != nil {
		e.Throw("Failed sending the notification", err)
	}
	logger.Info("Sent notification")
}

package env

import (
	"github.com/aspcartman/buddy-tbot/e"
	"github.com/getsentry/raven-go"
	"github.com/sirupsen/logrus"
)

var Log logrus.FieldLogger

func init() {
	logger := logrus.StandardLogger()
	logger.Formatter = &logrus.TextFormatter{DisableTimestamp: true}
	e.RegisterHook(func(ex *e.Exception) {
		logger.WithFields(logrus.Fields(ex.Args)).WithError(ex.Cause).Error(ex.Info)
	})
	Log = logger

	raven.SetDSN("https://0e4779deb4be44cc9400e800ff0486bb:84b84a0654644c918956b8f01dc14e83@sentry.aspc.me/3")
}

package env

import (
	"fmt"

	"github.com/aspcartman/darkside"
	"github.com/aspcartman/darkside/g"
	"github.com/aspcartman/exceptions"
	"github.com/getsentry/raven-go"
	"github.com/sirupsen/logrus"
)

var Log logrus.FieldLogger

func init() {
	setupPanicSafenet()
	setupLogging()
	//setupGoroutineLocal()
}

func setupPanicSafenet() {
	raven.SetDSN("https://0e4779deb4be44cc9400e800ff0486bb:84b84a0654644c918956b8f01dc14e83@sentry.aspc.me/3")

	darkside.SetUnrecoveredPanicHandler(func(p *g.Panic) {
		var title string
		var reason string
		var message string
		switch arg := p.Arg.(type) {
		case *e.Exception:
			title = arg.BottommostError().Error()
			reason = arg.Info
			message = arg.Error()
		case error:
			title = arg.Error()
			message = arg.Error()
		case fmt.Stringer:
			title = arg.String()
			message = arg.String()
		default:
			title = fmt.Sprint(arg)
			message = fmt.Sprint(arg)
		}
		trace := raven.NewStacktrace(3, 3, []string{})
		exc := raven.Exception{
			Type:       title,
			Value:      reason,
			Stacktrace: trace,
		}

		packet := raven.NewPacket(message, &exc)
		raven.Capture(packet, nil)
		raven.Wait()
		fmt.Println("SENT CRASH PANIC TO SENTRY")
	})
}

func setupLogging() {
	logger := logrus.StandardLogger()
	logger.Formatter = &logrus.TextFormatter{DisableTimestamp: true}
	e.RegisterHook(func(ex *e.Exception) {
		logger.WithFields(logrus.Fields(ex.Args)).WithError(ex.Cause).Error(ex.Info)
	})
	Log = logger
}

package env

import (
	"os"

	"github.com/aspcartman/buddy-tbot/e"
	"github.com/getsentry/raven-go"
	"github.com/mitchellh/panicwrap"
	"github.com/sirupsen/logrus"
)

var Log logrus.FieldLogger

func init() {
	//
	// Panics
	//
	raven.SetDSN("https://0e4779deb4be44cc9400e800ff0486bb:84b84a0654644c918956b8f01dc14e83@sentry.aspc.me/3")
	ret, err := panicwrap.BasicWrap(func(panicInfo string) {
		//dumpInfo, err := stack.ParseDump(strings.NewReader(panicInfo), ioutil.Discard)
		//NewPacket(rvalStr, append(append(interfaces, client.context.interfaces()...), NewException(errors.New(rvalStr), NewStacktrace(2, 3, client.includePaths)))...)
		raven.CaptureMessage("I have paniced!", nil)
	})
	switch {
	case err != nil:
		e.Throw("failure wrapping application with a panicwrapper", err)
	case ret == -1:
	default:
		os.Exit(ret)
	}

	//
	// Logging
	//
	logger := logrus.StandardLogger()
	logger.Formatter = &logrus.TextFormatter{DisableTimestamp: true}
	e.RegisterHook(func(ex *e.Exception) {
		logger.WithFields(logrus.Fields(ex.Args)).WithError(ex.Cause).Error(ex.Info)
	})
	Log = logger

}

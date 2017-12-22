package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/aspcartman/buddy-tbot/api/buddy"
	"github.com/aspcartman/buddy-tbot/e"
	"github.com/aspcartman/buddy-tbot/env"
	"github.com/sirupsen/logrus"
)

var logger logrus.FieldLogger

func init() {
	logger = env.Log
}

type Backend interface {
	HandleBuddyEvent(event *buddy.WebHookEvent)
}

func ListenForAPIEvents(backend Backend) {
	logger.Info("starting to listen for http events")

	srv := http.Server{}
	srv.Handler = handler{backend}
	srv.Addr = "0.0.0.0:8080"

	go e.Must(srv.ListenAndServe(), "failed listening and serving", e.Map{
		"addr": srv.Addr,
	})
}

type handler struct {
	Backend
}

func (h handler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	defer e.Catch(func(e *e.Exception) {
		logger.Error("failure while handling event")
		rw.WriteHeader(400)
	})

	data, _ := ioutil.ReadAll(r.Body)
	r.Body.Close()

	event := buddy.WebHookEvent{}
	e.Must(json.Unmarshal(data, &event), "failure unmarshalling the json", e.Map{
		"data": string(data),
	})

	go h.HandleBuddyEvent(&event)

	rw.WriteHeader(200)
}

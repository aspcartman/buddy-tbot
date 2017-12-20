package api

import (
	"net/http"

	"github.com/aspcartman/buddy-tbot/e"
)

func ListenForAPIEvents(clb func()) {
	srv := http.Server{}
	srv.Handler = handler{clb}
	srv.Addr = "0.0.0.0:8080"
	go e.Must(srv.ListenAndServe())
}

type handler struct {
	clb func()
}

func (h handler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	go h.clb()
	rw.WriteHeader(200)
}

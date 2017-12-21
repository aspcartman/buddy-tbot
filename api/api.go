package api

import (
	"io/ioutil"
	"net/http"

	"github.com/aspcartman/buddy-tbot/e"
)

func ListenForAPIEvents(clb func(payload string)) {
	srv := http.Server{}
	srv.Handler = handler{clb}
	srv.Addr = "0.0.0.0:8080"
	go e.Must(srv.ListenAndServe())
}

type handler struct {
	clb func(string)
}

func (h handler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	data, _ := ioutil.ReadAll(r.Body)
	r.Body.Close()

	go h.clb(string(data))

	rw.WriteHeader(200)
}

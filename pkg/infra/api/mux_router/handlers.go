package mux_router

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type muxEnginer struct {
}

func NewEnginerHandler() *muxEnginer {
	return &muxEnginer{}
}

func (m *muxEnginer) InitializeRoutes() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/ping", m.ping())
	return r
}

func (m *muxEnginer) ping() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		logrus.WithFields(logrus.Fields{"component": "mux - handler", "function": "ping"})
		json.NewEncoder(w).Encode(map[string]string{"status": "OK"})
	}
}

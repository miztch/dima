package main

import (
	"github.com/gorilla/mux"
)

type Router struct {
	muxRouter *mux.Router
}

func NewRouter(h *Handlers) *Router {
	r := &Router{muxRouter: mux.NewRouter()}
	r.RegisterRoutes(h)
	return r
}

func (r *Router) RegisterRoutes(h *Handlers) {
	r.muxRouter.HandleFunc("/", h.IndexHandler).Methods("GET")
	r.muxRouter.HandleFunc("/matches", h.GetMatchesHandler).Methods("GET")
}

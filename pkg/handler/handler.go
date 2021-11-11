package handler

import "github.com/gorilla/mux"

type Handler interface {
	Init(*mux.Router)
}

func GetHandlers() []Handler {
	return []Handler{&SummonerHandler{}}
}

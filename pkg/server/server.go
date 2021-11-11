package server

import (
	"fmt"
	"net/http"
	"sotater/copy-opgg-go/pkg/handler"

	"github.com/gorilla/mux"
)

type Server struct {
	router   *mux.Router
	handlers []handler.Handler
}

func GetServer() *Server {
	rootRouter := mux.NewRouter()
	handlers := handler.GetHandlers()
	for _, handler := range handlers {
		handler.Init(rootRouter)
	}
	return &Server{
		rootRouter,
		handlers,
	}
}

func (s Server) Run() {
	fmt.Println("Server started!")
	http.Handle("/", s.router)
	http.ListenAndServe(":8080", nil)
}

package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type WebServer struct {
	Router   chi.Router
	Handlers map[string]http.HandlerFunc
	Port     string
}

func NewWebServer(port string) *WebServer {
	return &WebServer{
		Router:   chi.NewRouter(),
		Handlers: make(map[string]http.HandlerFunc),
		Port:     port,
	}
}

func (s *WebServer) AddHandler(path string, handler http.HandlerFunc) {
	s.Handlers[path] = handler
}

func (s *WebServer) Start() {
	log.Default().Println("About to start web server")
	s.Router.Use(middleware.Logger)
	for path, handler := range s.Handlers {
		s.Router.Get(path, handler)
	}
	http.ListenAndServe(s.Port, s.Router)
	fmt.Println("Server is running")
}

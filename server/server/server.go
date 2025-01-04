package server

import (
	"fmt"
	"net/http"
)

type Server struct {
	Port   int
	routes map[string]func(http.ResponseWriter, *http.Request)
	server *http.ServeMux
}

func NewServer(port int) Server {
	httpServer := http.NewServeMux()
	s := Server{Port: port, server: httpServer, routes: make(map[string]func(http.ResponseWriter, *http.Request), 10)}
	return s
}

func (s *Server) AddRoute(route string, handler func(http.ResponseWriter, *http.Request)) {
	s.routes[route] = handler
}

func (s *Server) Start() {
	for route, handler := range s.routes {
		s.server.HandleFunc(route, handler)
	}

	port := fmt.Sprintf(":%v", s.Port)

	fmt.Printf("Server running on port %v\n", port)

	err := http.ListenAndServe(port, s.server)

	if err != nil {
		fmt.Println(err)
	}
}

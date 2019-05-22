package server

import "net/http"

func New() *Server {
	return &Server{}
}

type Server struct{}

func (s *Server) Run() error {
	return http.ListenAndServe(":3000", nil)
}

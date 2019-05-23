package server

import "net/http"

type Server struct{}

func New() *Server {
	return &Server{}
}

func (s *Server) Run() error {
	return http.ListenAndServe(":3000", nil)
}

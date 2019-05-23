package server

import (
	"fmt"
	"net/http"
)

type Config struct {
	Port int
	Log  bool
	Auth string
}

type Server struct {
	Config
}

func New(c Config) *Server {
	return &Server{Config: c}
}

func (s *Server) Run() error {
	addr := fmt.Sprintf("0.0.0.0:%d", s.Port)
	return http.ListenAndServe(addr, nil)
}

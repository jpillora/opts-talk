package server

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/jpillora/cookieauth"
	"github.com/jpillora/requestlog"
)

type Config struct {
	Port int    `opts:"help=listening port, env=PORT"`
	Log  bool   `opts:"help=enable request logging"`
	Auth string `opts:"help=enable basic-auth (user:pass)"`
}

type Server struct {
	Config
}

func New(c Config) *Server {
	return &Server{Config: c}
}

func (s *Server) hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world!\n"))
}

func (s *Server) Run() error {
	//router
	m := http.NewServeMux()
	m.HandleFunc("/", s.hello)
	//middleware
	h := http.Handler(m)
	if a := strings.Split(s.Auth, ":"); len(a) == 2 {
		h = cookieauth.Wrap(h, a[0], a[1])
	}
	if s.Log {
		h = requestlog.Wrap(h)
	}
	//listen
	addr := fmt.Sprintf("0.0.0.0:%d", s.Port)
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	log.Printf("listening on %d...", s.Port)
	return http.Serve(l, h)
}

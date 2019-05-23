package server

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/jpillora/cookieauth"
	"github.com/jpillora/opts"
	"github.com/jpillora/requestlog"
)

type command struct {
	Port     int    `opts:"help=listening port, env=PORT"`
	Log      bool   `opts:"help=enable request logging"`
	Auth     string `opts:"help=enable basic-auth (user:pass)"`
	KeyPath  string `opts:"help=path to TLS key"`
	CertPath string `opts:"help=path to TLS certificate"`
}

func New() opts.Opts {
	return opts.New(&command{
		Port: 3000,
	})
}

func (s *command) hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world!\n"))
}

func (s *command) Run() error {
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
	//optional tls
	if s.CertPath != "" && s.KeyPath != "" {
		cert, err := tls.LoadX509KeyPair(s.CertPath, s.KeyPath)
		if err != nil {
			return err
		}
		l = tls.NewListener(l, &tls.Config{
			Certificates: []tls.Certificate{cert},
		})
		log.Printf("enabled tls")
	}
	log.Printf("listening on %d...", s.Port)
	return http.Serve(l, h)
}

package main

import (
	"flag"

	"github.com/jpillora/opts-talk/step3/server"
)

func main() {
	c := server.Config{}
	flag.IntVar(&c.Port, "port", 3000, "http listening port")
	flag.BoolVar(&c.Log, "log", false, "enable http request logging")
	flag.StringVar(&c.Auth, "auth", "", "enable http basic-auth (user:pass)")
	flag.Parse()

	s := server.New(c)
	s.Run()
}

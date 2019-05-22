package main

import (
	"flag"
	"os"

	"github.com/jpillora/opts-talk/step3/server"
)

func main() {
	c := server.Config{}
	f := flag.NewFlagSet("server", flag.ExitOnError)
	f.IntVar(&c.Port, "port", 3000, "listening port")
	f.BoolVar(&c.Log, "log", false, "enable request logging")
	f.StringVar(&c.Auth, "auth", "", "enable basic-auth (user:pass)")
	f.Parse(os.Args[1:])
	s := server.New(c)
	s.Run()
}

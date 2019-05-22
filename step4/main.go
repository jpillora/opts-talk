package main

import (
	"github.com/jpillora/opts"
	"github.com/jpillora/opts-talk/step4/server"
)

func main() {
	c := server.Config{Port: 3000}
	opts.Parse(&c)
	s := server.New(c)
	s.Run()
}

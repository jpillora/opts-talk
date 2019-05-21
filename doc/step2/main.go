package main

import (
	"github.com/jpillora/opts-talk/step2/server"
)

func main() {
	c := server.Config{Port: 3000}
	s := server.New(c)
}

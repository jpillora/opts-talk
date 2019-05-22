package present

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jpillora/opts"
	"github.com/jpillora/present/handler"
	"github.com/jpillora/requestlog"
)

type Command struct {
	Host string `opts:"help=listening interface"`
	Port int    `opts:"help=listening port, env"`
	handler.Config
}

func New() opts.Opts {
	cmd := Command{
		Host: "localhost",
		Port: 3000,
		Config: handler.Config{
			ContentPath: "present",
		},
	}
	return opts.New(&cmd).
		Summary("runs an http server hosting the target presentation")
}

func (c *Command) Run() error {
	addr := fmt.Sprintf("%s:%d", c.Host, c.Port)
	//prepare present handler
	h, err := handler.New(c.Config)
	if err != nil {
		return err
	}
	//add http request logging
	h = requestlog.Wrap(h)
	//listen!
	log.Printf("listening at %s...", addr)
	return http.ListenAndServe(addr, h)
}

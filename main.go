package main

import (
	"github.com/jpillora/opts"
	"github.com/jpillora/opts-talk/present"
	"github.com/jpillora/opts-talk/search"
)

type cmd struct{}

func main() {
	opts.New(&cmd{}).
		AddCommand(present.New()).
		AddCommand(search.New()).
		Complete().
		Parse().
		RunFatal()
}

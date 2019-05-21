package main

import (
	"github.com/jpillora/opts"
	"github.com/jpillora/opts-talk/present"
)

func main() {
	opts.New(&struct{}{}).
		AddCommand(present.New()).
		Complete().
		Parse().
		RunFatal()
}

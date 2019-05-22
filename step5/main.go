package main

import (
	"github.com/jpillora/opts-talk/step5/cert"
	"github.com/jpillora/opts-talk/step5/server"
)

func main() {
	server.New().
		Complete().
		PkgRepo().
		AddCommand(cert.New()).
		Parse().
		RunFatal()
}

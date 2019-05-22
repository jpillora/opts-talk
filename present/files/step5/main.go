package main

import (
	"github.com/jpillora/opts-talk/present/files/step5/cert"
	"github.com/jpillora/opts-talk/present/files/step5/server"
)

func main() {
	server.New().
		Complete().
		PkgRepo().
		AddCommand(cert.New()).
		Parse().
		RunFatal()
}

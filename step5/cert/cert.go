package cert

import (
	"github.com/jpillora/opts"
)

type Paths struct {
	CertPath string `opts:"help=path to TLS certificate"`
	KeyPath  string `opts:"help=path to TLS key"`
}

var defaultPaths = Paths{
	CertPath: "server.cert",
	KeyPath:  "server.key",
}

func New() opts.Opts {
	type command struct{}
	o := opts.New(&command{}).
		Name("cert").
		Summary("quickly generate tls certificates for development").
		AddCommand(newGen()).
		AddCommand(newInspect())
	return o
}

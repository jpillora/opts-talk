package cert

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"

	"github.com/jpillora/opts"
)

type inspect struct {
	Paths
}

func newInspect() opts.Opts {
	return opts.New(&inspect{Paths: defaultPaths}).Name("inspect")
}

func (c inspect) Run() error {
	if c.CertPath == "" {
		return errors.New("missing cert path")
	}
	if c.KeyPath == "" {
		return errors.New("missing key path")
	}
	cert, err := tls.LoadX509KeyPair(c.CertPath, c.KeyPath)
	if err != nil {
		return err
	}
	for _, b := range cert.Certificate {
		cs, err := x509.ParseCertificates(b)
		if err != nil {
			return err
		}
		for _, c := range cs {
			fmt.Printf("issuer: %s\n", c.Issuer)
		}
	}
	fmt.Println("certificate pair is valid")
	return nil
}

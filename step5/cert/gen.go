package cert

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	"fmt"
	"math/big"
	"os"

	"github.com/jpillora/opts"
)

type gen struct {
	Paths
}

func newGen() opts.Opts {
	return opts.New(&gen{Paths: defaultPaths}).Name("gen")
}

func (c gen) Run() error {
	if c.CertPath == "" {
		return errors.New("missing cert path")
	}
	if c.KeyPath == "" {
		return errors.New("missing key path")
	}
	r := rand.Reader
	curve := elliptic.P256()
	priv, err := ecdsa.GenerateKey(curve, r)
	if err != nil {
		return fmt.Errorf("Failed to generate ecdsa key: %s", err)
	}
	pub := priv.Public()
	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(r, serialNumberLimit)
	if err != nil {
		return fmt.Errorf("Failed to generate serial number: %s", err)
	}
	//provide mininal information!
	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			CommonName:   "dev.cert",
			Organization: []string{"opts-demo"},
		},
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth},
		BasicConstraintsValid: true,
	}
	//write template, embed public key, sign with private key
	cert, err := x509.CreateCertificate(r, &template, &template, pub, priv)
	if err != nil {
		return fmt.Errorf("Failed to create certificate: %s", err)
	}
	//write private key
	key, err := x509.MarshalECPrivateKey(priv)
	if err != nil {
		return fmt.Errorf("Unable to marshal ECDSA private key: %s", err)
	}
	//write as files
	f, err := os.Create(c.CertPath)
	if err != nil {
		return err
	}
	defer f.Close()
	if err := pem.Encode(f, &pem.Block{Type: "CERTIFICATE", Bytes: cert}); err != nil {
		return err
	}
	fmt.Printf("created %s\n", c.CertPath)
	f, err = os.OpenFile(c.KeyPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer f.Close()
	if err := pem.Encode(f, &pem.Block{Type: "EC PRIVATE KEY", Bytes: key}); err != nil {
		return err
	}
	fmt.Printf("created %s\n", c.KeyPath)
	return nil
}

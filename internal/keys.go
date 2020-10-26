package internal

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	TypeRSAPrivateKey = "RSA PRIVATE KEY"
	TypeCertificate   = "CERTIFICATE"
)

func NewPrivateKeyFromPEM(data []byte) (*rsa.PrivateKey, error) {

	block, _ := pem.Decode(data)
	if block == nil {
		return nil, fmt.Errorf("no PEM data found")
	}

	if block.Type != TypeRSAPrivateKey {
		return nil, fmt.Errorf("expected private key, but got %s", block.Type)
	}

	return x509.ParsePKCS1PrivateKey(block.Bytes)
}

func NewPrivateKeyFromFile(path string) (*rsa.PrivateKey, error) {

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, fmt.Errorf("private key does not exist at %s", path)
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return NewPrivateKeyFromPEM(data)
}

func NewPublicKeyFromPEM(data []byte) (*rsa.PublicKey, error) {

	block, _ := pem.Decode(data)
	if block == nil {
		return nil, fmt.Errorf("no PEM data found")
	}

	if block.Type != TypeCertificate {
		return nil, fmt.Errorf("expected certificate, but got %s", block.Type)
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, err
	}

	if pub, ok := cert.PublicKey.(*rsa.PublicKey); ok {
		return pub, nil
	}

	return nil, fmt.Errorf("cannot create public key")
}

func NewPublicKeyFromFile(path string) (*rsa.PublicKey, error) {

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, fmt.Errorf("public key does not exist at %s", path)
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return NewPublicKeyFromPEM(data)
}

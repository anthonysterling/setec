package internal

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"
)

func NewPrivateKeyFromPEM(data []byte) (*rsa.PrivateKey, error) {

	block, _ := pem.Decode(data)
	if block == nil {
		return nil, fmt.Errorf("no PEM data found")
	}

	return x509.ParsePKCS1PrivateKey(block.Bytes)
}

func NewPrivateKeyFromFile(path string) (*rsa.PrivateKey, error) {

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, fmt.Errorf("key does not exist at %s", path)
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return NewPrivateKeyFromPEM(data)
}

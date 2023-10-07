package ceffu

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"errors"
	"log"
	"net/http"
)

type Client struct {
	key     *rsa.PrivateKey
	apiKey  string
	http    *http.Client
	logger  *log.Logger
	baseUrl string
}

// New creates a new Client from a base64 encoded x509 private key.
// The key must be in PKCS8 format. Also, the key must be RSA.
// x509KeyEncoded is the base64 encoded x509 private key. aka, the part between BEGIN PRIVATE KEY line and END PRIVATE KEY line.
// client is the http client to use. If nil, http.DefaultClient is used.
// logger is the logger to use. If nil, no logging is done.
// baseUrl is the base url to use. If empty, the default url is used.
func New(apiKey string, x509KeyBase64 string, client *http.Client, logger *log.Logger, baseUrl string) (*Client, error) {
	// Base 64 Decode x509KeyEncoded
	decodedLength := base64.StdEncoding.DecodedLen(len(x509KeyBase64))
	decodedBytes := make([]byte, decodedLength)
	decodedLength, err :=
		base64.StdEncoding.Decode([]byte(x509KeyBase64), decodedBytes)
	if err != nil {
		return nil, err
	}
	decodedBytes = decodedBytes[:decodedLength]
	key, err := x509.ParsePKCS8PrivateKey(decodedBytes)
	if err != nil {
		return nil, err
	}
	// Try Casting to RSA
	keyRsa := key.(*rsa.PrivateKey)
	if keyRsa == nil {
		return nil, errors.New("key is not RSA")
	}
	// If no client provided, use default
	if client == nil {
		client = http.DefaultClient
	}
	if baseUrl == "" {
		baseUrl = CeffuApiBaseUrl
	}
	return &Client{
		apiKey:  apiKey,
		baseUrl: baseUrl,
		key:     keyRsa,
		http:    client,
		logger:  logger,
	}, nil
}

func (c *Client) GetPublicKey() rsa.PublicKey {
	return c.key.PublicKey
}

func (c *Client) Logf(format string, v ...interface{}) {
	if c.logger != nil {
		c.logger.Printf(format, v...)
	}
}

package ceffu

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"encoding/base64"
)

func (c *Client) SignString(message string) (string, error) {
	msgHash := sha512.New()
	msgHash.Write([]byte(message))
	hash := msgHash.Sum(nil)
	sign, err := rsa.SignPKCS1v15(rand.Reader, c.key, crypto.SHA512, hash)
	//sign, err := c.key.Sign(nil, []byte(message))
	if err != nil {
		return "", err
	}
	// Base64 Encode sign
	signBase64 := base64.StdEncoding.EncodeToString(sign)
	return signBase64, nil
}

package ceffu

import (
	"crypto"
	"encoding/base64"
)

func (c *Client) SignString(message string) (string, error) {
	sign, err := c.key.Sign(nil, []byte(message), crypto.SHA512)
	if err != nil {
		return "", err
	}
	// Base64 Encode sign
	signBase64 := base64.StdEncoding.EncodeToString(sign)
	return signBase64, nil
}

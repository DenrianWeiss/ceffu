package ceffu

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"
)

func (c *Client) post(endpoint string, message map[string]interface{}) ([]byte, error) {
	// Encode message to JSON
	// Check if timestamp is present
	if _, ok := message["timestamp"]; !ok {
		message["timestamp"] = time.Now().UnixMilli()
	}
	encoded, err := json.Marshal(message)
	if err != nil {
		return nil, err
	}
	// Sign message
	signature, err := c.SignString(string(encoded))
	// Assemble request
	requestPath := fmt.Sprintf("%s%s%s", c.baseUrl, CeffuVersionPath, endpoint)
	// Add http header
	request, err := http.NewRequest(http.MethodPost, requestPath, bytes.NewReader(encoded))
	if err != nil {
		return nil, err
	}
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("signature", signature)
	request.Header.Add("open-apikey", c.apiKey)
	request.Header.Add("User-Agent", "ceffu-go-sdk/0.0.0")
	// Send request
	response, err := c.http.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	// Read response
	return io.ReadAll(response.Body)
}

func (c *Client) postV2(endpoint string, message map[string]interface{}) ([]byte, error) {
	// Encode message to JSON
	// Check if timestamp is present
	if _, ok := message["timestamp"]; !ok {
		message["timestamp"] = time.Now().UnixMilli()
	}
	encoded, err := json.Marshal(message)
	if err != nil {
		return nil, err
	}
	// Sign message
	signature, err := c.SignString(string(encoded))
	// Assemble request
	requestPath := fmt.Sprintf("%s%s%s", c.baseUrl, CeffuVersion2Path, endpoint)
	// Add http header
	request, err := http.NewRequest(http.MethodPost, requestPath, bytes.NewReader(encoded))
	if err != nil {
		return nil, err
	}
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("signature", signature)
	request.Header.Add("open-apikey", c.apiKey)
	request.Header.Add("User-Agent", "ceffu-go-sdk/0.0.0")
	// Send request
	response, err := c.http.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	// Read response
	return io.ReadAll(response.Body)
}

func (c *Client) get(endpoint string, params map[string]string) ([]byte, error) {
	// Assemble query string
	queryString := ""
	// todo: sort params - although it's not required, it's better to do so
	for key, value := range params {
		queryString += fmt.Sprintf("%s=%s&", key, value)
	}
	if _, ok := params["timestamp"]; !ok {
		queryString += fmt.Sprintf("timestamp=%d", time.Now().UnixMilli())
	}
	// Assemble request
	requestPath := fmt.Sprintf("%s%s%s?%s", c.baseUrl, CeffuVersionPath, endpoint, queryString)
	// Add http header
	request, err := http.NewRequest(http.MethodGet, requestPath, nil)
	if err != nil {
		return nil, err
	}
	// Sign query params
	signature, err := c.SignString(queryString)
	if err != nil {
		return nil, err
	}
	request.Header.Add("signature", signature)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("open-apikey", c.apiKey)
	request.Header.Add("User-Agent", "ceffu-go-sdk/0.0.0")
	// Send request
	response, err := c.http.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	// Read response
	return io.ReadAll(response.Body)
}

func (c *Client) getV2(endpoint string, params map[string]string) ([]byte, error) {
	// Assemble query string
	queryString := ""
	// todo: sort params - although it's not required, it's better to do so
	for key, value := range params {
		queryString += fmt.Sprintf("%s=%s&", key, value)
	}
	if _, ok := params["timestamp"]; !ok {
		queryString += fmt.Sprintf("timestamp=%d", time.Now().UnixMilli())
	}
	// Assemble request
	requestPath := fmt.Sprintf("%s%s%s?%s", c.baseUrl, CeffuVersion2Path, endpoint, queryString)
	// Add http header
	request, err := http.NewRequest(http.MethodGet, requestPath, nil)
	if err != nil {
		return nil, err
	}
	// Sign query params
	signature, err := c.SignString(queryString)
	if err != nil {
		return nil, err
	}
	request.Header.Add("signature", signature)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("open-apikey", c.apiKey)
	request.Header.Add("User-Agent", "ceffu-go-sdk/0.0.0")
	// Send request
	response, err := c.http.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	// Read response
	return io.ReadAll(response.Body)
}

func GetReqId() int64 {
	// generate int64 using random
	return rand.Int63()
}

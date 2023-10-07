package ceffu

import "encoding/json"

type CreateSubWalletResp struct {
	Data struct {
		WalletID       int64  `json:"walletId"`
		WalletName     string `json:"walletName"`
		WalletType     int    `json:"walletType"`
		ParentWalletID int64  `json:"parentWalletId"`
	} `json:"data"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

// CreateSubWallet creates a sub wallet for certain organization
// parentWalletId: required
// walletName: required, max 20 char
// autoCollection: optional, default false(0)
// requestId: optional, default random
func (c *Client) CreateSubWallet(parentWalletId int64, walletName string, autoCollection bool, requestId ...int64) (*CreateSubWalletResp, error) {
	params := map[string]interface{}{
		"parentWalletId": parentWalletId,
		"walletName":     walletName,
		"autoCollection": false,
	}
	if len(requestId) > 0 {
		params["requestId"] = requestId[0]
	} else {
		// Generate Random Request ID
		params["requestId"] = GetReqId()
	}
	post, err := c.post("subwallet/create", params)
	if err != nil {
		return nil, err
	}
	response := &CreateSubWalletResp{}
	err = json.Unmarshal(post, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

type UpdateSubWalletResp struct {
	ParentWalletID int    `json:"parentWalletId"`
	WalletID       int    `json:"walletId"`
	WalletName     string `json:"walletName"`
	WalletType     int    `json:"walletType"`
	AutoCollection int    `json:"autoCollection"`
}

// UpdateSubWallet updates a sub wallet for certain organization
// autoCollection: optional, default false(0)
// walletId: required
// walletName: optional, max 20 char
// requestId: optional, default random
func (c *Client) UpdateSubWallet(autoCollection bool, walletId int64, walletName string, requestId ...int64) (*UpdateSubWalletResp, error) {
	params := map[string]interface{}{
		"walletId":       walletId,
		"walletName":     walletName,
		"autoCollection": autoCollection,
	}
	if len(requestId) > 0 {
		params["requestId"] = requestId[0]
	} else {
		// Generate Random Request ID
		params["requestId"] = GetReqId()
	}
	post, err := c.post("subwallet/update", params)
	if err != nil {
		return nil, err
	}
	response := &UpdateSubWalletResp{}
	err = json.Unmarshal(post, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

type TransferWithSubWalletResp struct {
	Data struct {
		OrderViewID string `json:"orderViewId"`
		Status      int    `json:"status"`
		Direction   int    `json:"direction"`
	} `json:"data"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

// TransferWithSubWallet transfer assets between sub wallets and main wallet
// coinSymbol: required
// amount: required
// fromWalletId: required
// toWalletId: required
// requestId: optional, default random
func (c *Client) TransferWithSubWallet(coinSymbol string, amount float64, fromWalletId int64, toWalletId int64, requestId ...int64) (*TransferWithSubWalletResp, error) {
	params := map[string]interface{}{
		"coinSymbol":   coinSymbol,
		"amount":       amount,
		"fromWalletId": fromWalletId,
		"toWalletId":   toWalletId,
	}
	if len(requestId) > 0 {
		params["requestId"] = requestId[0]
	} else {
		// Generate Random Request ID
		params["requestId"] = GetReqId()
	}
	post, err := c.post("subwallet/transfer", params)
	if err != nil {
		return nil, err
	}
	response := &TransferWithSubWalletResp{}
	err = json.Unmarshal(post, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

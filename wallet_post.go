package ceffu

import (
	"encoding/json"
)

type CreateWalletResp struct {
	Data struct {
		WalletID   int64  `json:"walletId"`
		WalletName string `json:"walletName"`
		WalletType int    `json:"walletType"`
	} `json:"data"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

// CreateWallet creates a wallet for certain organization
// walletName: required
// walletType: required, Use WalletTypeInt*
// requestId: optional, default random
func (c *Client) CreateWallet(walletName string, walletType int, requestId ...int64) (*CreateWalletResp, error) {
	params := map[string]interface{}{
		"walletName": walletName,
		"walletType": walletType,
	}
	if len(requestId) > 0 {
		params["requestId"] = requestId[0]
	} else {
		// Generate Random Request ID
		params["requestId"] = GetReqId()
	}
	post, err := c.post("wallet/create", params)
	if err != nil {
		return nil, err
	}
	response := &CreateWalletResp{}
	err = json.Unmarshal(post, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

type UpdateWalletResp struct {
	Data struct {
		WalletID   int64  `json:"walletId"`
		WalletName string `json:"walletName"`
		WalletType int    `json:"walletType"`
	} `json:"data"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

// UpdateWallet updates a wallet for certain organization
// walletId: required
// walletName: required
// requestId: optional, default random
func (c *Client) UpdateWallet(walletId int64, walletName string, requestId ...int64) (*UpdateWalletResp, error) {
	params := map[string]interface{}{
		"walletId":   walletId,
		"walletName": walletName,
	}
	if len(requestId) > 0 {
		params["requestId"] = requestId[0]
	} else {
		// Generate Random Request ID
		params["requestId"] = GetReqId()
	}
	post, err := c.post("wallet/updateWallet", params)
	if err != nil {
		return nil, err
	}
	response := &UpdateWalletResp{}
	err = json.Unmarshal(post, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

type WithdrawalResp struct {
	Data struct {
		OrderViewID  string         `json:"orderViewId"`
		Status       WithdrawStatus `json:"status"` // See WithdrawStatusInt*
		TransferType TransferType   `json:"transferType"`
	} `json:"data"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

// Withdrawal withdraws from ceffu
// amount: withdrawal amount, decimal string
// coinSymbol: coin symbol, e.g. BTC
// memo: optional, memo for tx
// network: string, network for tx
// requestId: optional, default random
// walletId: required
// withdrawalAddress: required
func (c *Client) Withdrawal(amount string, coinSymbol string, memo string, network string, walletId int64, withdrawalAddress string, requestId ...int64) (*WithdrawalResp, error) {
	params := map[string]interface{}{
		"amount":            amount,
		"coinSymbol":        coinSymbol,
		"network":           network,
		"walletId":          walletId,
		"withdrawalAddress": withdrawalAddress,
	}
	if memo != "" {
		params["memo"] = memo
	}
	if len(requestId) > 0 {
		params["requestId"] = requestId[0]
	} else {
		// Generate Random Request ID
		params["requestId"] = GetReqId()
	}
	post, err := c.postV2("wallet/withdrawal", params)
	if err != nil {
		return nil, err
	}
	response := &WithdrawalResp{}
	err = json.Unmarshal(post, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

type TransferWithExchangeResp struct {
	Data struct {
		OrderViewID string            `json:"orderViewId"`
		Status      WithdrawStatus    `json:"status"`
		Direction   TransferDirection `json:"direction"`
	} `json:"data"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

// TransferWithExchange transfers from ceffu to binance exchange(currently only 1 direction is supported)
// amount: transfer amount, decimal string
// coinSymbol: coin symbol, e.g. BTC
// direction: transfer direction, use TransferDirectionInt*
// exchangeCode: only 10(Binance) supported.
// exchangeUserId: string, binance UID
// parentWalletId: if using parent shared wallet, required.
// requestId: optional, default random
func (c *Client) TransferWithExchange(amount string, coinSymbol string, direction int, exchangeCode int, exchangeUserId string, parentWalletId ...int64) (*TransferWithExchangeResp, error) {
	params := map[string]interface{}{
		"amount":         amount,
		"coinSymbol":     coinSymbol,
		"direction":      direction,
		"exchangeCode":   exchangeCode,
		"exchangeUserId": exchangeUserId,
	}
	if len(parentWalletId) > 0 {
		params["parentWalletId"] = parentWalletId[0]
	}
	// Generate Random Request ID
	params["requestId"] = GetReqId()
	post, err := c.post("wallet/transferWithExchange", params)
	if err != nil {
		return nil, err
	}
	response := &TransferWithExchangeResp{}
	err = json.Unmarshal(post, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

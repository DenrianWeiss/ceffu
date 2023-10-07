package ceffu

import (
	"encoding/json"
	"strconv"
	"time"
)

type GetSubWalletAssetDetailsResp struct {
	Data struct {
		Data []struct {
			CoinSymbol      string      `json:"coinSymbol"`
			Network         interface{} `json:"network"`
			Amount          string      `json:"amount"`
			AvailableAmount string      `json:"availableAmount"`
		} `json:"data"`
		TotalPage int `json:"totalPage"`
		PageNo    int `json:"pageNo"`
		PageLimit int `json:"pageLimit"`
	} `json:"data"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

// GetSubWalletAssetDetails gets asset details of a sub wallet
// coinSymbol: optional
// network: optional
// pageLimit: optional, default 25 max 25
// pageNo: optional, default 1
// walletId: required
func (c *Client) GetSubWalletAssetDetails(walletId int64, coinSymbol string, network string, pageLimit int, pageNo int) (*GetSubWalletAssetDetailsResp, error) {
	if pageLimit == 0 || pageLimit > 25 {
		pageLimit = 25
	}
	if pageNo == 0 {
		pageNo = 1
	}
	params := map[string]string{
		"walletId": strconv.FormatInt(walletId, 10),
	}
	if coinSymbol != "" {
		params["coinSymbol"] = coinSymbol
	}
	if network != "" {
		params["network"] = network
	}

	get, err := c.get("subwallet/asset/details", params)
	if err != nil {
		return nil, err
	}
	response := &GetSubWalletAssetDetailsResp{}
	err = json.Unmarshal(get, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

type GetSubWalletSummaryResp struct {
	Data struct {
		WalletIDStr      string `json:"walletIdStr"`
		TotalAmountInBTC string `json:"totalAmountInBTC"`
		TotalAmountInUSD string `json:"totalAmountInUSD"`
		Data             []struct {
			WalletIDStr         string `json:"walletIdStr"`
			SubTotalAmountInBTC string `json:"subTotalAmountInBTC"`
			SubTotalAmountInUSD string `json:"subTotalAmountInUSD"`
		} `json:"data"`
	} `json:"data"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

// GetSubWalletSummary return asset summary for subaccounts in certain prime/qualified account
// walletIdStr: prime or qualified account id
func (c *Client) GetSubWalletSummary(walletIdStr string) (*GetSubWalletSummaryResp, error) {
	params := map[string]string{
		"walletIdStr": walletIdStr,
	}

	get, err := c.get("subwallet/asset/summary", params)
	if err != nil {
		return nil, err
	}
	response := &GetSubWalletSummaryResp{}
	err = json.Unmarshal(get, response)
	if err != nil {
		return nil, err
	}
	return response, nil

}

type GetSubWalletDepositAddressResp struct {
	Data struct {
		WalletAddress string `json:"walletAddress"`
		Memo          string `json:"memo"`
	} `json:"data"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

// GetSubWalletDepositAddress gets deposit address for a sub wallet
// coinSymbol: required for prime, not required for qualified
// network: required
// walletId: required, sub wallet id
func (c *Client) GetSubWalletDepositAddress(walletId int64, coinSymbol string, network string) (*GetSubWalletDepositAddressResp, error) {
	params := map[string]string{
		"walletId": strconv.FormatInt(walletId, 10),
		"network":  network,
	}
	if coinSymbol != "" {
		params["coinSymbol"] = coinSymbol
	}

	get, err := c.get("subwallet/deposit/address", params)
	if err != nil {
		return nil, err
	}
	response := &GetSubWalletDepositAddressResp{}
	err = json.Unmarshal(get, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

type GetSubWalletDepositHistoryResp struct {
	Data []struct {
		Direction           int     `json:"direction"`
		Network             string  `json:"network"`
		Memo                int     `json:"memo"`
		CoinSymbol          string  `json:"coinSymbol"`
		Amount              float64 `json:"amount"`
		FeeSymbol           string  `json:"feeSymbol"`
		FeeAmount           float64 `json:"feeAmount"`
		WalletID            int64   `json:"walletId"`
		FromAddress         string  `json:"fromAddress"`
		ToAddress           string  `json:"toAddress"`
		OrderViewID         int64   `json:"orderViewId"`
		TransferType        int     `json:"transferType"`
		Status              int     `json:"status"`
		TxID                string  `json:"txId"`
		TxTime              int64   `json:"txTime"`
		ConfirmedBlockCount int     `json:"confirmedBlockCount"`
		MaxConfirmedBlock   string  `json:"maxConfirmedBlock"`
		UnlockConfirm       string  `json:"unlockConfirm"`
	} `json:"data"`
	PageLimit int `json:"pageLimit"`
	PageNo    int `json:"pageNo"`
	TotalPage int `json:"totalPage"`
}

// GetSubWalletDepositHistory gets deposit history for a sub wallet, v2 api
// walletId: sub wallet id
// coinSymbol: optional
// network: optional
// startTime: required, unix timestamp in millisecond
// endTime: optional, unix timestamp in millisecond, default now
// pageLimit: optional, default 25 max 25
// pageNo: optional, default 1
func (c *Client) GetSubWalletDepositHistory(walletId int64, coinSymbol string, network string, startTime int64, endTime int64, pageLimit int, pageNo int) (*GetSubWalletDepositHistoryResp, error) {
	if pageLimit == 0 || pageLimit > 25 {
		pageLimit = 25
	}
	if pageNo == 0 {
		pageNo = 1
	}
	params := map[string]string{
		"walletId":   strconv.FormatInt(walletId, 10),
		"startTime":  strconv.FormatInt(startTime, 10),
		"pageLimit":  strconv.Itoa(pageLimit),
		"pageNo":     strconv.Itoa(pageNo),
		"network":    network,
		"coinSymbol": coinSymbol,
	}
	if endTime != 0 {
		params["endTime"] = strconv.FormatInt(endTime, 10)
	} else {
		params["endTime"] = strconv.FormatInt(time.Now().UnixMilli(), 10)
	}

	get, err := c.get("subwallet/deposit/history", params)
	if err != nil {
		return nil, err
	}
	response := &GetSubWalletDepositHistoryResp{}
	err = json.Unmarshal(get, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

type GetAllSubWalletDepositHistoryResp struct {
	Data struct {
		Data []struct {
			OrderViewId         interface{} `json:"orderViewId"`
			TxId                *string     `json:"txId"`
			TransferType        int         `json:"transferType"`
			Direction           int         `json:"direction"`
			FromAddress         string      `json:"fromAddress"`
			ToAddress           string      `json:"toAddress"`
			Network             *string     `json:"network"`
			CoinSymbol          string      `json:"coinSymbol"`
			Amount              string      `json:"amount"`
			FeeSymbol           interface{} `json:"feeSymbol"`
			FeeAmount           string      `json:"feeAmount"`
			Status              int         `json:"status"`
			ConfirmedBlockCount interface{} `json:"confirmedBlockCount"`
			UnlockConfirm       interface{} `json:"unlockConfirm"`
			MaxConfirmBlock     interface{} `json:"maxConfirmBlock"`
			Memo                interface{} `json:"memo"`
			TxTime              int64       `json:"txTime"`
			WalletIdStr         string      `json:"walletIdStr"`
			RequestId           interface{} `json:"requestId"`
		} `json:"data"`
		TotalPage int `json:"totalPage"`
		PageNo    int `json:"pageNo"`
		PageLimit int `json:"pageLimit"`
	} `json:"data"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

// GetAllSubWalletDepositHistory gets deposit history for all sub wallets
// parentWalletId required, prime or qualified account id
// coinSymbol: optional
// network: optional
// pageLimit: optional, default 25 max 25
// pageNo: optional, default 1
func (c *Client) GetAllSubWalletDepositHistory(parentWalletId int64, coinSymbol string, network string, pageLimit int, pageNo int) (*GetAllSubWalletDepositHistoryResp, error) {
	if pageLimit == 0 || pageLimit > 25 {
		pageLimit = 25
	}
	if pageNo == 0 {
		pageNo = 1
	}
	params := map[string]string{
		"parentWalletId": strconv.FormatInt(parentWalletId, 10),
		"pageLimit":      strconv.Itoa(pageLimit),
		"pageNo":         strconv.Itoa(pageNo),
		"network":        network,
		"coinSymbol":     coinSymbol,
	}

	get, err := c.getV2("subwallet/deposit/history", params)
	if err != nil {
		return nil, err
	}
	response := &GetAllSubWalletDepositHistoryResp{}
	err = json.Unmarshal(get, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

type GetAllSubWalletDepositAddressResp struct {
	Data struct {
		Data []struct {
			WalletAddress string `json:"walletAddress"`
			Memo          string `json:"memo"`
			WalletID      int64  `json:"walletId"`
		} `json:"data"`
		TotalPage int `json:"totalPage"`
		PageNo    int `json:"pageNo"`
		PageLimit int `json:"pageLimit"`
	} `json:"data"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

// GetAllSubWalletDepositAddress get All sub wallet deposit address under the requested Parent Wallet ID (Prime), coinSymbol and network. (Only applicable to Parent Wallet Id(Prime))
// parentWalletId: prime account id
// coinSymbol: required
// network: required
// pageLimit: optional, default 25 max 25
// pageNo: optional, default 1
func (c *Client) GetAllSubWalletDepositAddress(parentWalletId int64, coinSymbol string, network string, pageLimit int, pageNo int) (*GetAllSubWalletDepositAddressResp, error) {
	if pageLimit == 0 || pageLimit > 25 {
		pageLimit = 25
	}
	if pageNo == 0 {
		pageNo = 1
	}
	params := map[string]string{
		"parentWalletId": strconv.FormatInt(parentWalletId, 10),
		"pageLimit":      strconv.Itoa(pageLimit),
		"pageNo":         strconv.Itoa(pageNo),
		"network":        network,
		"coinSymbol":     coinSymbol,
	}

	get, err := c.get("subwallet/deposit/address", params)
	if err != nil {
		return nil, err
	}
	response := &GetAllSubWalletDepositAddressResp{}
	err = json.Unmarshal(get, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

type GetAllSubWalletResp struct {
	Data struct {
		Data      []int64 `json:"data"`
		TotalPage int     `json:"totalPage"`
		PageNo    int     `json:"pageNo"`
		PageLimit int     `json:"pageLimit"`
	} `json:"data"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

// GetAllSubWallet gets all sub wallets under a prime account
// parentWalletId: prime account id
// pageLimit: optional, default 25 max 25
// pageNo: optional, default 1
func (c *Client) GetAllSubWallet(parentWalletId int64, pageLimit int, pageNo int) (*GetAllSubWalletResp, error) {
	if pageLimit == 0 || pageLimit > 25 {
		pageLimit = 25
	}
	if pageNo == 0 {
		pageNo = 1
	}
	params := map[string]string{
		"parentWalletId": strconv.FormatInt(parentWalletId, 10),
		"pageLimit":      strconv.Itoa(pageLimit),
		"pageNo":         strconv.Itoa(pageNo),
	}

	get, err := c.get("subwallet/list", params)
	if err != nil {
		return nil, err
	}
	response := &GetAllSubWalletResp{}
	err = json.Unmarshal(get, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

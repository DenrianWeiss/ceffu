package ceffu

import (
	"encoding/json"
	"strconv"
	"time"
)

type GetPrimeSupportedCoinListResp struct {
	Data []struct {
		CoinID            int    `json:"coinId"`
		CoinSymbol        string `json:"coinSymbol"`
		CoinFullName      string `json:"coinFullName"`
		NetworkConfigList []struct {
			CoinSymbol       string `json:"coinSymbol"`
			CoinFullName     string `json:"coinFullName"`
			Network          string `json:"network"`
			DepositEnable    bool   `json:"depositEnable"`
			WithdrawalEnable bool   `json:"withdrawalEnable"`
			WithdrawalMin    string `json:"withdrawalMin"`
			WithdrawalMax    string `json:"withdrawalMax"`
			Precision        int    `json:"precision"`
			WithdrawalFee    string `json:"withdrawalFee"`
			AddressRegex     string `json:"addressRegex"`
		} `json:"networkConfigList"`
		DepositEnable    bool `json:"depositEnable"`
		WithdrawalEnable bool `json:"withdrawalEnable"`
	} `json:"data"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

// GetPrimeSupportedCoinList returns the list of supported coins
func (c *Client) GetPrimeSupportedCoinList() (*GetPrimeSupportedCoinListResp, error) {
	get, err := c.get("wallet/shared/coin", map[string]string{})
	if err != nil {
		return nil, err
	}
	response := &GetPrimeSupportedCoinListResp{}
	err = json.Unmarshal(get, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

type GetQualifiedSupportedCoinListResp struct {
	Data []struct {
		CoinID           int         `json:"coinId"`
		CoinSymbol       string      `json:"coinSymbol"`
		CoinFullName     interface{} `json:"coinFullName"`
		Network          string      `json:"network"`
		Protocol         interface{} `json:"protocol"`
		DepositEnable    bool        `json:"depositEnable"`
		WithdrawalEnable bool        `json:"withdrawalEnable"`
		WithdrawalMin    string      `json:"withdrawalMin"`
		WithdrawalMax    interface{} `json:"withdrawalMax"`
		Precision        int         `json:"precision"`
		AddressRegex     string      `json:"addressRegex"`
	} `json:"data"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

// GetQualifiedSupportedCoinList returns the list of coins that are supported by Ceffu's qualified wallet
func (c *Client) GetQualifiedSupportedCoinList() (*GetQualifiedSupportedCoinListResp, error) {
	get, err := c.get("wallet/qualified/coin", map[string]string{})
	if err != nil {
		return nil, err
	}
	response := &GetQualifiedSupportedCoinListResp{}
	err = json.Unmarshal(get, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

type GetWalletListResp struct {
	Data struct {
		Data []struct {
			WalletID    int64  `json:"walletId"`
			WalletName  string `json:"walletName"`
			WalletType  int    `json:"walletType"`
			WalletIDStr string `json:"walletIdStr"`
		} `json:"data"`
		TotalPage int `json:"totalPage"`
		PageNo    int `json:"pageNo"`
		PageLimit int `json:"pageLimit"`
	} `json:"data"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

// GetWalletList returns the list of wallets for certain organization
// pageLimit: optional, default 25, max 25
// pageNo: optional, default 1
func (c *Client) GetWalletList(pageLimit int, pageNo int) (*GetWalletListResp, error) {
	if pageLimit > 25 || pageLimit == 0 {
		pageLimit = 25
	}
	if pageNo == 0 {
		pageNo = 1
	}
	get, err := c.get("wallet/list", map[string]string{
		"pageLimit": strconv.Itoa(pageLimit),
		"pageNo":    strconv.Itoa(pageNo),
	})
	if err != nil {
		return nil, err
	}
	response := &GetWalletListResp{}
	err = json.Unmarshal(get, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

type GetAssetDetailsResp struct {
	Data struct {
		Data []struct {
			CoinSymbol      string `json:"coinSymbol"`
			Network         string `json:"network"`
			Amount          string `json:"amount"`
			AvailableAmount string `json:"availableAmount"`
		} `json:"data"`
		TotalPage int `json:"totalPage"`
		PageNo    int `json:"pageNo"`
		PageLimit int `json:"pageLimit"`
	} `json:"data"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

// GetAssetDetails returns the asset details of a wallet
// coinSymbol: optional, if not provided, all coins will be returned
// network: optional, if not provided, all networks will be returned
// walletId: required
// pageLimit: optional, default 25, max 25
// pageNo: optional, default 1
func (c *Client) GetAssetDetails(coinSymbol string, network string, walletId string, pageLimit int, pageNo int) (*GetAssetDetailsResp, error) {
	if pageLimit > 25 || pageLimit == 0 {
		pageLimit = 25
	}
	if pageNo == 0 {
		pageNo = 1
	}
	params := map[string]string{
		"pageLimit": strconv.Itoa(pageLimit),
		"pageNo":    strconv.Itoa(pageNo),
		"walletId":  walletId,
	}
	if coinSymbol != "" {
		params["coinSymbol"] = coinSymbol
	}
	if network != "" {
		params["network"] = network
	}
	get, err := c.get("wallet/asset/list", params)
	if err != nil {
		return nil, err
	}
	response := &GetAssetDetailsResp{}
	err = json.Unmarshal(get, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

type GetAssetSummaryResp struct {
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

// GetAssetSummary This method allows you to fetch the specified wallet's asset summary,
//
//	represented in its equivalent BTC & USD value.
//	Please note that this equivalent value is provided for reference only and is based on our internal calculation.
//
// WalletId required
func (c *Client) GetAssetSummary(walletId string) (*GetAssetSummaryResp, error) {
	get, err := c.get("wallet/asset/summary", map[string]string{
		"walletIdStr": walletId,
	})
	if err != nil {
		return nil, err
	}
	response := &GetAssetSummaryResp{}
	err = json.Unmarshal(get, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

type GetWithdrawalFeeResp struct {
	Data struct {
		FeeAmount string `json:"feeAmount"`
		FeeSymbol string `json:"feeSymbol"`
	} `json:"data"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

// GetWithdrawalFee returns the withdrawal fee of a coin
// walletId: required
// coinSymbol: required
// network: required, network symbol in capital letters
// amount: optional, if not provided, the minimum withdrawal amount will be returned
func (c *Client) GetWithdrawalFee(walletId string, coinSymbol string, network string, amount string) (*GetWithdrawalFeeResp, error) {
	params := map[string]string{
		"walletId":   walletId,
		"coinSymbol": coinSymbol,
		"network":    network,
	}
	if amount != "" {
		params["amount"] = amount
	}
	get, err := c.get("wallet/withdrawal/fee", params)
	if err != nil {
		return nil, err
	}
	response := &GetWithdrawalFeeResp{}
	err = json.Unmarshal(get, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

type GetDepositAddressResp struct {
	Data struct {
		WalletAddress string `json:"walletAddress"`
		Memo          string `json:"memo"`
	} `json:"data"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

// GetDepositAddress returns the deposit address of a coin
// coinSymbol: required
// network: required, network symbol in capital letters
// walletId: required
func (c *Client) GetDepositAddress(coinSymbol string, network string, walletId string) (*GetDepositAddressResp, error) {
	get, err := c.get("wallet/deposit/address", map[string]string{
		"coinSymbol": coinSymbol,
		"network":    network,
		"walletId":   walletId,
	})
	if err != nil {
		return nil, err
	}
	response := &GetDepositAddressResp{}
	err = json.Unmarshal(get, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

type GetDepositHistoryResp struct {
	Data struct {
		Data []struct {
			OrderViewID         string            `json:"orderViewId"`
			TxID                interface{}       `json:"txId"` // String or null
			TransferType        TransferType      `json:"transferType"`
			Direction           TransferDirection `json:"direction"` // See constants.go/TransferDirectionInt*
			FromAddress         string            `json:"fromAddress"`
			ToAddress           string            `json:"toAddress"`
			Network             interface{}       `json:"network"` // String or null
			CoinSymbol          string            `json:"coinSymbol"`
			Amount              string            `json:"amount"`
			FeeSymbol           interface{}       `json:"feeSymbol"`
			FeeAmount           string            `json:"feeAmount"`
			Status              int               `json:"status"`
			ConfirmedBlockCount int               `json:"confirmedBlockCount"`
			UnlockConfirm       int               `json:"unlockConfirm"`
			MaxConfirmBlock     interface{}       `json:"maxConfirmBlock"`
			Memo                interface{}       `json:"memo"` // String or null
			TxTime              int64             `json:"txTime"`
			WalletID            int64             `json:"walletId"`
		} `json:"data"`
		TotalPage int `json:"totalPage"`
		PageNo    int `json:"pageNo"`
		PageLimit int `json:"pageLimit"`
	} `json:"data"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

// GetDepositHistory returns the deposit history of a coin
// walletId: required
// coinSymbol: optional, if not provided, all coins will be returned
// network: optional, if not provided, all networks will be returned
// startTime: required, unix timestamp in milliseconds
// endTime: optional, default to current time, unix timestamp in milliseconds
// pageLimit: optional, default 25, max 25
// pageNo: optional, default 1
func (c *Client) GetDepositHistory(walletId string, coinSymbol string, network string, startTime int64, endTime int64, pageLimit int, pageNo int) (*GetDepositHistoryResp, error) {
	if pageLimit > 25 || pageLimit == 0 {
		pageLimit = 25
	}
	if pageNo == 0 {
		pageNo = 1
	}
	params := map[string]string{
		"pageLimit": strconv.Itoa(pageLimit),
		"pageNo":    strconv.Itoa(pageNo),
		"walletId":  walletId,
		"startTime": strconv.FormatInt(startTime, 10),
	}
	if coinSymbol != "" {
		params["coinSymbol"] = coinSymbol
	}
	if network != "" {
		params["network"] = network
	}
	if endTime != 0 {
		params["endTime"] = strconv.FormatInt(endTime, 10)
	} else {
		params["endTime"] = strconv.FormatInt(time.Now().UnixMilli(), 10)
	}
	get, err := c.get("wallet/deposit/history", params)
	if err != nil {
		return nil, err
	}
	response := &GetDepositHistoryResp{}
	err = json.Unmarshal(get, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

type GetDepositDetailResp struct {
	Data []struct {
		OrderViewId         string      `json:"orderViewId"`
		TxId                interface{} `json:"txId"` // String or null
		TransferType        int         `json:"transferType"`
		Direction           int         `json:"direction"`
		FromAddress         string      `json:"fromAddress"`
		ToAddress           string      `json:"toAddress"`
		Network             interface{} `json:"network"` // String or null
		CoinSymbol          string      `json:"coinSymbol"`
		Amount              string      `json:"amount"`
		FeeSymbol           interface{} `json:"feeSymbol"` // String or null
		FeeAmount           string      `json:"feeAmount"`
		Status              int         `json:"status"`
		ConfirmedBlockCount interface{} `json:"confirmedBlockCount"` // int or null
		UnlockConfirm       interface{} `json:"unlockConfirm"`       // int or null
		MaxConfirmBlock     interface{} `json:"maxConfirmBlock"`     // int or null
		Memo                interface{} `json:"memo"`
		TxTime              int64       `json:"txTime"`
		WalletId            int64       `json:"walletId"`
		RequestId           interface{} `json:"requestId"` // String or null
	} `json:"data"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

// GetDepositDetail queries the deposit detail of a transaction
// txId: required, transaction id for corresponding deposit
func (c *Client) GetDepositDetail(txId string) (*GetDepositDetailResp, error) {
	get, err := c.getV2("wallet/deposit/detail", map[string]string{
		"txId": txId,
	})
	if err != nil {
		return nil, err
	}
	response := &GetDepositDetailResp{}
	err = json.Unmarshal(get, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

type GetWithdrawalHistoryResp struct {
	Data []struct {
		Direction           int          `json:"direction"`
		Network             string       `json:"network"`
		Memo                int          `json:"memo"`
		CoinSymbol          string       `json:"coinSymbol"`
		Amount              float64      `json:"amount"`
		FeeSymbol           string       `json:"feeSymbol"`
		FeeAmount           float64      `json:"feeAmount"`
		WalletID            int64        `json:"walletId"`
		FromAddress         string       `json:"fromAddress"`
		ToAddress           string       `json:"toAddress"`
		OrderViewID         int64        `json:"orderViewId"`
		TransferType        TransferType `json:"transferType"`
		Status              int          `json:"status"`
		TxID                string       `json:"txId"`
		TxTime              int64        `json:"txTime"`
		ConfirmedBlockCount int          `json:"confirmedBlockCount"`
		MaxConfirmedBlock   string       `json:"maxConfirmedBlock"`
		UnlockConfirm       string       `json:"unlockConfirm"`
	} `json:"data"`
	PageLimit int `json:"pageLimit"`
	PageNo    int `json:"pageNo"`
	TotalPage int `json:"totalPage"`
}

// GetWithdrawalHistory returns the withdrawal history of a coin
// walletId required
// network optional, if not provided, all networks will be returned
// coinSymbol optional, if not provided, all coins will be returned
// status optional, if not provided, all statuses will be returned
// startTime required, unix timestamp in milliseconds
// endTime optional, default to current time, unix timestamp in milliseconds
// pageLimit optional, default 25, max 25
// pageNo optional, default 1
func (c *Client) GetWithdrawalHistory(walletId string, network string, coinSymbol string, status WithdrawStatus, startTime int64, endTime int64, pageLimit int, pageNo int) (*GetWithdrawalHistoryResp, error) {
	if pageLimit > 25 || pageLimit == 0 {
		pageLimit = 25
	}
	if pageNo == 0 {
		pageNo = 1
	}
	params := map[string]string{
		"pageLimit": strconv.Itoa(pageLimit),
		"pageNo":    strconv.Itoa(pageNo),
		"walletId":  walletId,
		"startTime": strconv.FormatInt(startTime, 10),
	}
	if network != "" {
		params["network"] = network
	}
	if coinSymbol != "" {
		params["coinSymbol"] = coinSymbol
	}
	if status != 0 {
		params["status"] = strconv.Itoa(int(status))
	}
	if endTime != 0 {
		params["endTime"] = strconv.FormatInt(endTime, 10)
	} else {
		params["endTime"] = strconv.FormatInt(time.Now().UnixMilli(), 10)
	}
	get, err := c.get("wallet/withdrawal/history", params)
	if err != nil {
		return nil, err
	}
	response := &GetWithdrawalHistoryResp{}
	err = json.Unmarshal(get, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

type GetWithdrawalDetailResp struct {
	Data struct {
		OrderViewID         string `json:"orderViewId"`
		TxID                string `json:"txId"`
		TransferType        int    `json:"transferType"`
		Direction           int    `json:"direction"`
		FromAddress         string `json:"fromAddress"`
		ToAddress           string `json:"toAddress"`
		Network             string `json:"network"`
		CoinSymbol          string `json:"coinSymbol"`
		Amount              string `json:"amount"`
		FeeSymbol           string `json:"feeSymbol"`
		FeeAmount           string `json:"feeAmount"`
		Status              int    `json:"status"`
		ConfirmedBlockCount int    `json:"confirmedBlockCount"`
		Memo                string `json:"memo"`
		TxTime              int64  `json:"txTime"`
		WalletID            int64  `json:"walletId"`
	} `json:"data"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

// GetWithdrawalDetail queries the withdrawal detail of a transaction
// orderViewId: required, order view id for corresponding withdrawal
func (c *Client) GetWithdrawalDetail(orderViewId string) (*GetWithdrawalDetailResp, error) {
	get, err := c.get("wallet/withdrawal/detail", map[string]string{
		"orderViewId": orderViewId,
	})
	if err != nil {
		return nil, err
	}
	response := &GetWithdrawalDetailResp{}
	err = json.Unmarshal(get, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

type GetTransferHistoryWithExchangeResp struct {
	Data struct {
		Data []struct {
			OrderViewID    string      `json:"orderViewId"`
			Direction      int         `json:"direction"`
			WalletID       int64       `json:"walletId"`
			CreateTime     int64       `json:"createTime"`
			ExchangeCode   int         `json:"exchangeCode"`
			ExchangeUserID string      `json:"exchangeUserId"`
			CoinSymbol     string      `json:"coinSymbol"`
			Amount         string      `json:"amount"`
			Status         int         `json:"status"`
			RequestID      interface{} `json:"requestId"`
		} `json:"data"`
		TotalPage int `json:"totalPage"`
		PageNo    int `json:"pageNo"`
		PageLimit int `json:"pageLimit"`
	} `json:"data"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

// GetTransferHistoryWithExchange returns the transfer history of a coin
// walletId required
// coinSymbol optional, if not provided, all coins will be returned
// direction optional, if not provided, all directions will be returned
// status optional, if not provided, all statuses will be returned
// startTime required, unix timestamp in milliseconds
// endTime optional, default to current time, unix timestamp in milliseconds
// pageLimit optional, default 25, max 25
// pageNo optional, default 1
func (c *Client) GetTransferHistoryWithExchange(walletId string, coinSymbol string, direction TransferDirection, status WithdrawStatus, startTime int64, endTime int64, pageLimit int, pageNo int) (*GetTransferHistoryWithExchangeResp, error) {
	if pageLimit > 25 || pageLimit == 0 {
		pageLimit = 25
	}
	if pageNo == 0 {
		pageNo = 1
	}
	params := map[string]string{
		"pageLimit": strconv.Itoa(pageLimit),
		"pageNo":    strconv.Itoa(pageNo),
		"walletId":  walletId,
		"startTime": strconv.FormatInt(startTime, 10),
	}
	if coinSymbol != "" {
		params["coinSymbol"] = coinSymbol
	}
	if direction != 0 {
		params["direction"] = strconv.Itoa(int(direction))
	}
	if status != 0 {
		params["status"] = strconv.Itoa(int(status))
	}
	if endTime != 0 {
		params["endTime"] = strconv.FormatInt(endTime, 10)
	} else {
		params["endTime"] = strconv.FormatInt(time.Now().UnixMilli(), 10)
	}
	get, err := c.get("wallet/transfer/exchange/history", params)
	if err != nil {
		return nil, err
	}
	response := &GetTransferHistoryWithExchangeResp{}
	err = json.Unmarshal(get, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

type GetTransferDetailWithExchangeResp struct {
	Data struct {
		OrderViewID    string            `json:"orderViewId"`
		Direction      TransferDirection `json:"direction"`
		WalletID       int64             `json:"walletId"`
		CreateTime     int64             `json:"createTime"`
		ExchangeCode   int               `json:"exchangeCode"`
		ExchangeUserID string            `json:"exchangeUserId"`
		CoinSymbol     string            `json:"coinSymbol"`
		Amount         string            `json:"amount"`
		Status         int               `json:"status"`
		RequestID      string            `json:"requestId"`
	} `json:"data"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

// GetTransferDetailWithExchange queries the transfer detail of a transaction
// orderViewId: required, order view id for corresponding transfer
// walletId: required
func (c *Client) GetTransferDetailWithExchange(orderViewId string, walletId string) (*GetTransferDetailWithExchangeResp, error) {
	get, err := c.get("wallet/transfer/exchange/detail", map[string]string{
		"orderViewId": orderViewId,
		"walletId":    walletId,
	})
	if err != nil {
		return nil, err
	}
	response := &GetTransferDetailWithExchangeResp{}
	err = json.Unmarshal(get, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

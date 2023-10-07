package ceffu

import "encoding/json"

const BusinessTypeDeposit = "10"
const BusinessTypeWithdraw = "20"
const BusinessTypeTransferToBinanceExchange = "30"

type GetStatusResp struct {
	Data struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	} `json:"data"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (c *Client) GetStatus(business string, walletType string) (resp *GetStatusResp, err error) {
	params := map[string]string{
		"business":   business,
		"walletType": walletType,
	}
	get, err := c.get("/status", params)
	if err != nil {
		return nil, err
	}
	response := &GetStatusResp{}
	err = json.Unmarshal(get, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

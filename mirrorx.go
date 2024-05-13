package ceffu

import (
	"encoding/json"
	"math"
	"math/rand"
	"strconv"
	"time"
)

const (
	GetMirrorXLinkIdListApi       = "mirrorX/mirrorXLinkId/list"
	GetMirrorXDelegationOrdersApi = "mirrorX/order/list"
	GetMirrorXAvailableAmountApi  = "mirrorX/order/check"
	GetMirrorXAssetPositionsApi   = "mirrorX/positions/list"
	CreateMirrorXOrderApi         = "mirrorX/order"
)

type GetMirrorXLinkListResp struct {
	Data struct {
		Data []struct {
			MirrorXLinkId string `json:"mirrorXLinkId"`
			BinanceUID    string `json:"binanceUID"`
			WalletIdStr   string `json:"walletIdStr"`
			Label         string `json:"label"`
			Status        int    `json:"status"`
			CreateDate    string `json:"createDate"`
		} `json:"data"`
		TotalPage int `json:"totalPage"`
		PageNo    int `json:"pageNo"`
		PageLimit int `json:"pageLimit"`
	} `json:"data"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

type GetMirrorXDelegationOrdersResp struct {
	Data struct {
		Data []struct {
			MirrorXLinkId string `json:"mirrorXLinkId"`
			BinanceUID    string `json:"binanceUID"`
			WalletIdStr   string `json:"walletIdStr"`
			OrderType     int    `json:"orderType"`
			Amount        string `json:"amount"`
			CoinSymbol    string `json:"coinSymbol"`
			Status        int    `json:"status"`
			OrderTime     string `json:"orderTime"`
			OrderViewId   string `json:"orderViewId"`
		} `json:"data"`
		TotalPage int `json:"totalPage"`
		PageNo    int `json:"pageNo"`
		PageLimit int `json:"pageLimit"`
	} `json:"data"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

type GetMirrorXAvailableAmountResp struct {
	Data struct {
		CoinSymbol         string `json:"coinSymbol"`
		MaxAvailableAmount string `json:"maxAvailableAmount"`
	} `json:"data"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

type GetMirrorXAssetPositionsResp struct {
	Data struct {
		Data []struct {
			MirrorXLinkId  string `json:"mirrorXLinkId"`
			BinanceUID     string `json:"binanceUID"`
			WalletIdStr    string `json:"walletIdStr"`
			CoinSymbol     string `json:"coinSymbol"`
			MirrorXBalance string `json:"mirrorXBalance"`
		} `json:"data"`
		TotalPage int `json:"totalPage"`
		PageNo    int `json:"pageNo"`
		PageLimit int `json:"pageLimit"`
	} `json:"data"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

type CreateMirrorXOrderReq struct {
	MirrorXLinkId int    `json:"mirrorXLinkId"`
	OrderType     int    `json:"orderType"`
	CoinSymbol    string `json:"coinSymbol"`
	Amount        string `json:"amount"`
	RequestId     string `json:"requestId"`
}

type CreateMirrorXOrderResp struct {
	Data struct {
		OrderViewId string `json:"orderViewId"`
		Status      int    `json:"status"`
		RequestId   string `json:"requestId"`
	} `json:"data"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (c *Client) GetMirrorXLinkList(pageLimit, pageNo int) (*GetMirrorXLinkListResp, error) {
	params := map[string]string{
		"pageLimit": strconv.Itoa(pageLimit),
		"pageNo":    strconv.Itoa(pageNo),
	}
	resp, err := c.get(GetMirrorXLinkIdListApi, params)
	if err != nil {
		return nil, err
	}
	var result GetMirrorXLinkListResp
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetMirrorXDelegationOrders get mirrorX delegation orders
// @param mirrorXLinkId: mirrorXLinkId, must have, binded binance uid
// @param coinSymbol: coin symbol optional, if not set, will return all orders, example: "USDT"
// @param orderType: order type, optional, if not set, will return all orders, 10: buy, 20: sell
// @param startTime: start time, required
// @param endTime: end time, required
// @param pageLimit: page limit, optional, default 10
// @param pageNo: page no, optional, default 1
func (c *Client) GetMirrorXDelegationOrders(mirrorXLinkId string, coinSymbol string, orderType MirrorXOrderType, startTime int, endTime int, pageLimit int, pageNo int) (*GetMirrorXDelegationOrdersResp, error) {
	params := map[string]string{
		"mirrorXLinkId": mirrorXLinkId,
		"startTime":     strconv.Itoa(startTime),
	}
	if endTime == 0 {
		// Set endTime to current time
		endTime = int(time.Now().Unix())
	}
	params["endTime"] = strconv.Itoa(endTime)
	if coinSymbol != "" {
		params["coinSymbol"] = coinSymbol
	}
	if orderType != 0 {
		params["orderType"] = strconv.Itoa(int(orderType))
	}
	if pageLimit != 0 {
		params["pageLimit"] = strconv.Itoa(pageLimit)
	} else {
		params["pageLimit"] = "10"
	}
	if pageNo != 0 {
		params["pageNo"] = strconv.Itoa(pageNo)
	} else {
		params["pageNo"] = "1"
	}
	resp, err := c.get(GetMirrorXDelegationOrdersApi, params)
	if err != nil {
		return nil, err
	}
	var result GetMirrorXDelegationOrdersResp
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetMirrorXAvailableAmount get mirrorX available amount
// @param mirrorXLinkId: mirrorXLinkId, must have, binded binance uid
// @param coinSymbol: coin symbol, must have, example: "USDT"
// @param orderType: order type, must have, 10: buy, 20: sell
func (c *Client) GetMirrorXAvailableAmount(mirrorXLinkId string, coinSymbol string, orderType MirrorXOrderType) (*GetMirrorXAvailableAmountResp, error) {
	params := map[string]string{
		"mirrorXLinkId": mirrorXLinkId,
		"coinSymbol":    coinSymbol,
		"orderType":     strconv.Itoa(int(orderType)),
	}
	resp, err := c.get(GetMirrorXAvailableAmountApi, params)
	if err != nil {
		return nil, err
	}
	var result GetMirrorXAvailableAmountResp
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetMirrorXAssetPositions get mirrorX asset positions
// @param mirrorXLinkId: mirrorXLinkId, must have, binded binance uid
// @param excludeZeroAmountFlag: exclude zero amount flag, optional, default false
// @param pageLimit: page limit, optional, default 10
// @param pageNo: page no, optional, default 1
func (c *Client) GetMirrorXAssetPositions(mirrorXLinkId string, excludeZeroAmountFlag bool, pageLimit int, pageNo int) (*GetMirrorXAssetPositionsResp, error) {
	params := map[string]string{
		"mirrorXLinkId": mirrorXLinkId,
	}
	if excludeZeroAmountFlag {
		params["excludeZeroAmountFlag"] = "true"
	}
	if pageLimit != 0 {
		params["pageLimit"] = strconv.Itoa(pageLimit)
	} else {
		params["pageLimit"] = "10"
	}
	if pageNo != 0 {
		params["pageNo"] = strconv.Itoa(pageNo)
	} else {
		params["pageNo"] = "1"
	}
	resp, err := c.get(GetMirrorXAssetPositionsApi, params)
	if err != nil {
		return nil, err
	}
	var result GetMirrorXAssetPositionsResp
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateMirrorXOrder place mirrorX order
// @param mirrorXLinkId: mirrorXLinkId, must have, binded binance uid
// @param orderType: order type, must have, 10: buy, 20: sell
// @param coinSymbol: coin symbol, must have, example: "USDT"
// @param amount: amount, must have, example: "100"
// @param timestamp: timestamp, must have, example: 1630000000
func (c *Client) CreateMirrorXOrder(req *CreateMirrorXOrderReq) (*CreateMirrorXOrderResp, error) {
	if req.RequestId == "" {
		// Fill requestId with random int
		randId := rand.Intn(math.MaxInt)
		req.RequestId = strconv.Itoa(randId)
	}
	// Convert req to map[string]interface{}
	reqMap := map[string]interface{}{}
	reqJson, _ := json.Marshal(req)
	_ = json.Unmarshal(reqJson, &reqMap)
	resp, err := c.post(CreateMirrorXOrderApi, reqMap)
	if err != nil {
		return nil, err
	}
	var result CreateMirrorXOrderResp
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

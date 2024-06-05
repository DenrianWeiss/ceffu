package ceffu

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"testing"
)

func TestGetDeposit(t *testing.T) {
	apiKey, _ := os.LookupEnv("CEFFU_API_KEY")
	apiSecret, _ := os.LookupEnv("CEFFU_API_SECRET")
	depositTx, _ := os.LookupEnv("CEFFU_TX")
	cl, err := New(apiKey, apiSecret, http.DefaultClient, nil, CeffuApiBaseUrl)
	if err != nil {
		panic(err)
	}
	detail, err := cl.GetDepositDetail(depositTx)
	if err != nil {
		log.Printf("GetDepositDetail error %s", err)
		return
	}
	log.Printf("GetDepositDetail %+v", detail)
	// Get Deposit History

}

func TestCeffuClient(t *testing.T) {
	// Load Ceffu Args from env
	apiKey, _ := os.LookupEnv("CEFFU_API_KEY")
	apiSecret, _ := os.LookupEnv("CEFFU_API_SECRET")
	cl, err := New(apiKey, apiSecret, http.DefaultClient, nil, CeffuApiBaseUrl)
	if err != nil {
		panic(err)
	}
	list, err := cl.GetWalletList(10, 1)
	if err != nil {
		t.Error(err)
	}
	t.Log(list)
	// Find a wallet with name hybird_test_ent
	var walletId int64
	for _, datum := range list.Data.Data {
		if datum.WalletName == "hybird_test_ent" {
			t.Log(datum)
			walletId = datum.WalletID
		}
	}
	detail, err := cl.GetDepositDetail("0x49b865a694d13d92c22555a4e024171752479b3265af56fd8e38d72dce79ce75")
	if err != nil {
		return
	}
	t.Log(detail)
	// Get wallet charge address
	address, err := cl.GetDepositAddress("ETH", "ETH", fmt.Sprintf("%d", walletId))
	if err != nil {
		t.Error(err)
	}
	t.Log(address)
	// Get Balance
	summary, err := cl.GetAssetSummary(fmt.Sprintf("%d", walletId))
	if err != nil {
		t.Error(err)
	}
	t.Log(summary)
	details, err := cl.GetAssetDetails("USDT", "ETH", fmt.Sprintf("%d", walletId), 10, 1)
	if err != nil {
		t.Error(err)
	}
	t.Log(details)
}

func TestGetMirrorXInfo(t *testing.T) {
	// Load Ceffu Args from env
	apiKey, _ := os.LookupEnv("CEFFU_API_KEY")
	apiSecret, _ := os.LookupEnv("CEFFU_API_SECRET")
	cl, err := New(apiKey, apiSecret, http.DefaultClient, nil, CeffuApiBaseUrl)
	if err != nil {
		panic(err)
	}
	list, err := cl.GetMirrorXLinkList(10, 1)
	if err != nil {
		t.Error(err)
	}
	t.Log(list)
	mirrorXLinkIdActive := 11282
	orders, err := cl.GetMirrorXDelegationOrders(fmt.Sprintf("%d", mirrorXLinkIdActive), "", 0, 0, 0, 10, 1)
	if err != nil {
		t.Error(err)
	}
	t.Log(orders)
	// Get MirrorX Order Detail
	amount, err := cl.GetMirrorXAvailableAmount("11282", "USDT", MirrorXOrderTypeDeposit)
	if err != nil {
		t.Error(err)
	}
	t.Log(amount)
	// Create Mirror Using API
	order, err := cl.CreateMirrorXOrder(&CreateMirrorXOrderReq{
		MirrorXLinkId: mirrorXLinkIdActive,
		OrderType:     int(MirrorXOrderTypeDeposit),
		CoinSymbol:    "USDT",
		Amount:        amount.Data.MaxAvailableAmount,
		RequestId:     "",
	})
	if err != nil {
		return
	}
	t.Log(order)
}

func TestWithdraw(t *testing.T) {
	apiKey, _ := os.LookupEnv("CEFFU_API_KEY")
	apiSecret, _ := os.LookupEnv("CEFFU_API_SECRET")
	cl, err := New(apiKey, apiSecret, http.DefaultClient, nil, CeffuApiBaseUrl)
	if err != nil {
		panic(err)
	}
	list, err := cl.GetWalletList(10, 1)
	if err != nil {
		t.Error(err)
	}
	t.Log(list)
	// Find a wallet with name hybird_test_ent
	var walletId int64
	for _, datum := range list.Data.Data {
		if datum.WalletName == "hybird_test_ent" {
			t.Log(datum)
			walletId = datum.WalletID
		}
	}
	// Get USDC Amount
	ret, _ := cl.GetAssetDetails("USDT", "ETH", fmt.Sprintf("%d", walletId), 10, 1)
	var usdcAmount float64
	for _, datum := range ret.Data.Data {
		if datum.CoinSymbol == "USDT" {
			usdcAmount, _ = strconv.ParseFloat(datum.Amount, 64)
		}
	}
	usdcAmountInt := usdcAmount
	fee, err := cl.GetWithdrawalFee(strconv.FormatInt(walletId, 10), "USDT", "ETH", fmt.Sprintf("%f", usdcAmountInt))
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", fee)
	feeAmount, _ := strconv.ParseFloat(fee.Data.FeeAmount, 64)
	// Get Fee From Resp & Call Withdraw
	outAccount := "0xd3BdD5B82B4a75cb2081405C35B9DDd6875fdC03"
	rq, err := cl.Withdrawal(fmt.Sprintf("%f", usdcAmountInt-feeAmount), "USDT", "", "ETH", walletId, outAccount)
	t.Logf("%+v", rq)
}

func TestGetWithdraw(t *testing.T) {
	apiKey, _ := os.LookupEnv("CEFFU_API_KEY")
	apiSecret, _ := os.LookupEnv("CEFFU_API_SECRET")
	cl, err := New(apiKey, apiSecret, http.DefaultClient, nil, CeffuApiBaseUrl)
	if err != nil {
		panic(err)
	}
	detail, err := cl.GetWithdrawalDetail("24648095850261443323360001")
	if err != nil {
		log.Printf("GetDepositDetail error %s", err)
		return
	}
	log.Printf("GetDepositDetail %+v", detail)
	// Get Deposit History

}

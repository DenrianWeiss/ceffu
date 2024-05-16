package ceffu

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"testing"
)

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
		Amount:        "1",
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
	fee, err := cl.GetWithdrawalFee(strconv.FormatInt(walletId, 10), "USDC", "ETH", "2")
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", fee)
	feeAmount, _ := strconv.ParseFloat(fee.Data.FeeAmount, 64)
	// Get Fee From Resp & Call Withdraw
	outAccount, _ := os.LookupEnv("CEFFU_OUT_ACCOUNT")
	rq, err := cl.Withdrawal(fmt.Sprintf("%.2f", 100-feeAmount), "USDC", "", "ETH", walletId, outAccount)
	t.Logf("%+v", rq)
}

package ceffu

import (
	"fmt"
	"net/http"
	"os"
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
	details, err := cl.GetAssetDetails("", "ETH", fmt.Sprintf("%d", walletId), 10, 1)
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

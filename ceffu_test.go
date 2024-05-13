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
}

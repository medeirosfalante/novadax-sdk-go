package novadax_test

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/rafaeltokyo/novadax-sdk-go"
)

func TestGetOrders(t *testing.T) {
	godotenv.Load()
	client := novadax.New("", os.Getenv("ENV"))
	response, errAPI, err := client.Market().Depth(novadax.DepthQuery{Symbol: "BTC_BRL"})
	if err != nil {
		t.Errorf("err : %#v", err)
		return
	}
	if errAPI != nil {
		t.Errorf("errAPI : %#v", errAPI)
		return
	}
	if response == nil {
		t.Error("response is null")
		return
	}
}

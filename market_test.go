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
	response, errAPI, err := client.Market().Depth(&novadax.DepthQuery{Symbol: "BTC_BRL"})
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

func TestGetTrades(t *testing.T) {
	godotenv.Load()
	client := novadax.New("", os.Getenv("ENV"))
	response, errAPI, err := client.Market().Trades(&novadax.DepthQuery{Symbol: "BTC_BRL"})
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

func TestGetTradesQueryNill(t *testing.T) {
	godotenv.Load()
	client := novadax.New("", os.Getenv("ENV"))
	_, _, err := client.Market().Trades(nil)
	if err == nil {
		t.Errorf("err : %v", err)
		return
	}
}

func TestGetTradesInvalidSymbol(t *testing.T) {
	godotenv.Load()
	client := novadax.New("", os.Getenv("ENV"))
	_, errAPI, err := client.Market().Trades(&novadax.DepthQuery{Symbol: "BTC"})
	if err != nil {
		t.Errorf("err : %#v", err)
		return
	}
	if errAPI == nil {
		t.Errorf("errAPI : %#v", errAPI)
		return
	}
}

func TestGetTradesEqual15(t *testing.T) {
	godotenv.Load()
	client := novadax.New("", os.Getenv("ENV"))
	response, errAPI, err := client.Market().Trades(&novadax.DepthQuery{Symbol: "BTC_BRL", Limit: 15})
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

	if len(*response) < 15 {
		t.Error("invalid total")
	}

}

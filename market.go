package novadax

import (
	"errors"
	"strconv"
)

type DepthQuery struct {
	Symbol string `json:"symbol"`
	Limit  int    `json:"limit"`
}

// Market struct
type Market struct {
	client *APIClient
	Crypto string
}

//Market - Create a new instance struct
func (c *APIClient) Market() *Market {
	return &Market{client: c}
}

type Book struct {
	Asks []BookItem `json:"asks"`
	Bids []BookItem `json:"bids"`
}

type Trades []*TradeItem

type TradeItem struct {
	Price     string `json:"price" `
	Amount    string `json:"amount"`
	Side      string `json:"side" `
	Timestamp int64  `json:"timestamp" `
}

func (c BookItem) Price() float64 {
	f, _ := strconv.ParseFloat(c[0], 64)
	return f
}
func (c BookItem) Amount() float64 {
	f, _ := strconv.ParseFloat(c[1], 64)
	return f
}

type BookItem []string

// Depth - OrderBook in exchange
func (p Market) Depth(query *DepthQuery) (*Book, *Error, error) {
	var response *Book
	if query.Limit == 0 {
		query.Limit = 10
	}
	if query.Symbol == "" {
		return nil, nil, errors.New("symbol is required")
	}
	err, errAPI := p.client.Request("GET", "/market/depth", nil, query, &response)
	if err != nil {
		return nil, nil, err
	}
	if errAPI != nil {
		return nil, errAPI, nil
	}
	return response, nil, nil
}

// Trades - Trades in exchange
func (p Market) Trades(query *DepthQuery) (*Trades, *Error, error) {
	var response *Trades
	if query != nil {
		if query.Limit == 0 {
			query.Limit = 10
		}
		if query.Symbol == "" {
			return nil, nil, errors.New("symbol is required")
		}
	} else {
		return nil, nil, errors.New("query is nill")
	}
	err, errAPI := p.client.Request("GET", "/market/trades", nil, query, &response)
	if err != nil {
		return nil, nil, err
	}
	if errAPI != nil {
		return nil, errAPI, nil
	}
	return response, nil, nil
}

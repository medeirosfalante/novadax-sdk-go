package novadax

import "errors"

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

func (c BookItem) Price() float64 {
	return c[0]
}
func (c BookItem) Amount() float64 {
	return c[1]
}

type BookItem []float64

// Depth - OrderBook in exchange
func (p Market) Depth(query DepthQuery) (*Book, *Error, error) {
	var response *Book
	if query.Limit == 0 {
		query.Limit = 10
	}
	if query.Symbol == "" {
		return nil, nil, errors.New("symbol is required")
	}
	err, errAPI := p.client.Request("GET", "/market/depth", nil, &query, &response)
	if err != nil {
		return nil, nil, err
	}
	if errAPI != nil {
		return nil, errAPI, nil
	}
	return response, nil, nil
}

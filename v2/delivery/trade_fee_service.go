package delivery

import (
	"context"
	"encoding/json"
	"net/http"
)

// TradeFeeService shows current trade fee for all symbols available
type TradeFeeService struct {
	c      *Client
	symbol string
}

type TradeFeeDetails struct {
	Symbol              string `json:"symbol"`
	MakerCommissionRate string `json:"makerCommissionRate"`
	TakerCommissionRate string `json:"takerCommissionRate"`
}

// Symbol set the symbol parameter for the request
func (s *TradeFeeService) Symbol(symbol string) *TradeFeeService {
	s.symbol = symbol

	return s
}

// Do send request
func (s *TradeFeeService) Do(ctx context.Context, opts ...RequestOption) (res *TradeFeeDetails, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/dapi/v1/commissionRate",
		secType:  secTypeSigned,
	}

	r.setParam("symbol", s.symbol)

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res = new(TradeFeeDetails)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

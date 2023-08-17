package futures

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
	Symbol          string `json:"symbol"`
	MakerCommission string `json:"makerCommissionRate"`
	TakerCommission string `json:"takerCommissionRate"`
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
		endpoint: "/fapi/v1/commissionRate",
		secType:  secTypeSigned,
	}

	r.setParam("symbol", s.symbol)

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res = new(TradeFeeDetails)

	if err = json.Unmarshal(data, &res); err != nil {
		return nil, err
	}

	return res, nil
}

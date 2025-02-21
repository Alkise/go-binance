package delivery

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
)

type assetTradeFeeServiceSuite struct {
	baseTestSuite
}

func (a *assetTradeFeeServiceSuite) assertTradeFeeServiceEqual(expected, other *TradeFeeDetails) {
	r := a.r()

	r.Equal(expected.Symbol, other.Symbol)
	r.Equal(expected.MakerCommissionRate, other.MakerCommissionRate)
	r.Equal(expected.TakerCommissionRate, other.TakerCommissionRate)
}

func TestTradeFeeService(t *testing.T) {
	suite.Run(t, new(assetTradeFeeServiceSuite))
}

func (s *assetTradeFeeServiceSuite) TestSingleSymbolTradeFee() {
	data := []byte(`
	{
		"symbol": "ADABNB",
		"makerCommissionRate": "0.001",
		"takerCommissionRate": "0.001"
	}
	`)

	s.mockDo(data, nil)
	defer s.assertDo()

	fees, err := s.client.NewTradeFeeService().Symbol("ADABNB").Do(context.Background())
	s.r().NoError(err)

	s.assertTradeFeeServiceEqual(&TradeFeeDetails{
		Symbol:              "ADABNB",
		MakerCommissionRate: "0.001",
		TakerCommissionRate: "0.001"},
		fees)
}

package binance

import (
	"context"
	"encoding/json"
	"net/http"
)

// StakingProductPositionService fetches the staking product positions
type StakingProductPositionService struct {
	c          *Client
	product    StakingProduct
	productId  *string
	positionId *int64
	asset      *string
	current    *int32
	size       *int32
}

// Product sets the product parameter.
func (s *StakingProductPositionService) Product(product StakingProduct) *StakingProductPositionService {
	s.product = product
	return s
}

// ProductId sets the productId parameter.
func (s *StakingProductPositionService) ProductId(productId string) *StakingProductPositionService {
	s.productId = &productId
	return s
}

func (s *StakingProductPositionService) PositionId(positionId int64) *StakingProductPositionService {
	s.positionId = &positionId
	return s
}

// Asset sets the asset parameter.
func (s *StakingProductPositionService) Asset(asset string) *StakingProductPositionService {
	s.asset = &asset
	return s
}

// Current sets the current parameter.
func (s *StakingProductPositionService) Current(current int32) *StakingProductPositionService {
	s.current = &current
	return s
}

// Size sets the size parameter.
func (s *StakingProductPositionService) Size(size int32) *StakingProductPositionService {
	s.size = &size
	return s
}

// Do sends the request.
func (s *StakingProductPositionService) Do(ctx context.Context) (*StakingProductPositions, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/staking/position",
		secType:  secTypeSigned,
	}
	r.setParam("product", s.product)
	if s.productId != nil {
		r.setParam("productId", *s.productId)
	}
	if s.positionId != nil {
		r.setParam("positionId", *s.positionId)
	}
	if s.asset != nil {
		r.setParam("asset", *s.asset)
	}
	if s.current != nil {
		r.setParam("current", *s.current)
	}
	if s.size != nil {
		r.setParam("size", *s.size)
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := new(StakingProductPositions)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// StakingProductPositions represents a list of staking product positions.
type StakingProductPositions []StakingProductPosition

// StakingProductPosition represents a staking product position.
type StakingProductPosition struct {
	PositionId                 int64  `json:"positionId"`
	ProductId                  string `json:"productId"`
	Asset                      string `json:"asset"`
	Amount                     string `json:"amount"`
	PurchaseTime               int64  `json:"purchaseTime"`
	Duration                   int64  `json:"duration"`
	AccrualDays                int64  `json:"accrualDays"`
	RewardAsset                string `json:"rewardAsset"`
	APY                        string `json:"apy"`
	RewardAmount               string `json:"rewardAmt"`
	ExtraRewardAsset           string `json:"extraRewardAsset"`
	ExtraRewardAPY             string `json:"extraRewardAPY"`
	EstimatedExtraRewardAmount string `json:"estExtraRewardAmt"`
	NextInterestPay            string `json:"nextInterestPay"`
	NextInterestPayDate        int64  `json:"nextInterestPayDate"`
	PayInterestPeriod          int64  `json:"payInterestPeriod"`
	RedeemAmountEarly          string `json:"redeemAmountEarly"`
	InterestEndDate            int64  `json:"interestEndDate"`
	DeliverDate                int64  `json:"deliverDate"`
	RedeemPeriod               int64  `json:"redeemPeriod"`
	RedeemingAmount            string `json:"redeemingAmt"`
	PartialAmountDeliverDate   int64  `json:"partialAmtDeliverDate"`
	CanRedeemEarly             bool   `json:"canRedeemEarly"`
	Renewable                  bool   `json:"renewable"`
	Type                       string `json:"type"`
	Status                     string `json:"status"`
}

// StakingHistoryService fetches the staking history
type StakingHistoryService struct {
	c               *Client
	product         StakingProduct
	transactionType StakingTransactionType
	asset           *string
	startTime       *int64
	endTime         *int64
	current         *int32
	size            *int32
}

// Product sets the product parameter.
func (s *StakingHistoryService) Product(product StakingProduct) *StakingHistoryService {
	s.product = product
	return s
}

// TransactionType sets the txnType parameter.
func (s *StakingHistoryService) TransactionType(transactionType StakingTransactionType) *StakingHistoryService {
	s.transactionType = transactionType
	return s
}

// Asset sets the asset parameter.
func (s *StakingHistoryService) Asset(asset string) *StakingHistoryService {
	s.asset = &asset
	return s
}

// StartTime sets the startTime parameter.
func (s *StakingHistoryService) StartTime(startTime int64) *StakingHistoryService {
	s.startTime = &startTime
	return s
}

// EndTime sets the endTime parameter.
func (s *StakingHistoryService) EndTime(endTime int64) *StakingHistoryService {
	s.endTime = &endTime
	return s
}

// Current sets the current parameter.
func (s *StakingHistoryService) Current(current int32) *StakingHistoryService {
	s.current = &current
	return s
}

// Size sets the size parameter.
func (s *StakingHistoryService) Size(size int32) *StakingHistoryService {
	s.size = &size
	return s
}

// Do sends the request.
func (s *StakingHistoryService) Do(ctx context.Context) (*StakingHistory, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/staking/stakingRecord",
		secType:  secTypeSigned,
	}
	r.setParam("product", s.product)
	r.setParam("txnType", s.transactionType)
	if s.asset != nil {
		r.setParam("asset", *s.asset)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.current != nil {
		r.setParam("current", *s.current)
	}
	if s.size != nil {
		r.setParam("size", *s.size)
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := new(StakingHistory)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// StakingHistory represents a list of staking history transactions.
type StakingHistory []StakingHistoryTransaction

// StakingHistoryTransaction represents a staking history transaction.
type StakingHistoryTransaction struct {
	PositionId  int64  `json:"positionId"`
	Time        int64  `json:"time"`
	Asset       string `json:"asset"`
	Project     string `json:"project"`
	Amount      string `json:"amount"`
	LockPeriod  int64  `json:"lockPeriod"`
	DeliverDate int64  `json:"deliverDate"`
	Type        string `json:"type"`
	Status      string `json:"status"`
}

// StakingProductListService fetches available staking products
type StakingProductListService struct {
	c       *Client
	product StakingProduct
	asset   *string
	current *int32
	size    *int32
}

func (s *StakingProductListService) Product(product StakingProduct) *StakingProductListService {
	s.product = product
	return s
}

func (s *StakingProductListService) Asset(asset string) *StakingProductListService {
	s.asset = &asset
	return s
}

func (s *StakingProductListService) Current(current int32) *StakingProductListService {
	s.current = &current
	return s
}

func (s *StakingProductListService) Size(size int32) *StakingProductListService {
	s.size = &size
	return s
}

// Do sends the request.
func (s *StakingProductListService) Do(ctx context.Context) (*StakingProducts, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/staking/productList",
		secType:  secTypeSigned,
	}
	r.setParam("product", s.product)
	if s.asset != nil {
		r.setParam("asset", *s.asset)
	}
	if s.current != nil {
		r.setParam("current", *s.current)
	}
	if s.size != nil {
		r.setParam("size", *s.size)
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := new(StakingProducts)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// StakingProducts represents a list of staking products
type StakingProducts []StakingProduct

// StakingProduct represents staking product
type StakingProduct struct {
	ProjectId string               `json:"projectId"`
	Detail    StakingProductDetail `json:"detail"`
	Quota     StakingProductQuota  `json:"quota"`
}

type StakingProductDetail struct {
	Asset       string `json:"asset"`
	RewardAsset string `json:"rewardAsset"`
	Duration    int32  `json:"duration"`
	Renewable   bool   `json:"renewable"`
	APY         string `json:"apy"`
}

type StakingProductQuota struct {
	TotalPersonalQuota string `json:"totalPersonalQuota"`
	Minimum            string `json:"minimum"`
}

type PurchaseStakingProductService struct {
	c         *Client
	product   StakingProduct
	productId string
	amount    string
	renewable *bool
}

func (s *PurchaseStakingProductService) Product(product StakingProduct) *PurchaseStakingProductService {
	s.product = product
	return s
}

func (s *PurchaseStakingProductService) ProductId(productId string) *PurchaseStakingProductService {
	s.productId = productId
	return s
}

func (s *PurchaseStakingProductService) Amount(amount string) *PurchaseStakingProductService {
	s.amount = amount
	return s
}

func (s *PurchaseStakingProductService) Renewable(renewable bool) *PurchaseStakingProductService {
	s.renewable = &renewable
	return s
}

func (s *PurchaseStakingProductService) Do(ctx context.Context, opts ...RequestOption) (*PurchaseStakingProductResponse, error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/staking/purchase",
		secType:  secTypeSigned,
	}
	m := params{
		"product":   s.product,
		"productId": s.productId,
		"amount":    s.amount,
	}
	if s.renewable != nil {
		m["renewable"] = *s.renewable
	}
	r.setParams(m)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := new(PurchaseStakingProductResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type PurchaseStakingProductResponse struct {
	PositionId int64 `json:"positionId"`
	Success    bool  `json:"success"`
}

type RedeemStakingProductService struct {
	c          *Client
	product    StakingProduct
	productId  string
	positionId *int64
	amount     *string
}

func (s *RedeemStakingProductService) Product(product StakingProduct) *RedeemStakingProductService {
	s.product = product
	return s
}

func (s *RedeemStakingProductService) ProductId(productId string) *RedeemStakingProductService {
	s.productId = productId
	return s
}

func (s *RedeemStakingProductService) PositionId(positionId int64) *RedeemStakingProductService {
	s.positionId = &positionId
	return s
}

func (s *RedeemStakingProductService) Amount(amount string) *RedeemStakingProductService {
	s.amount = &amount
	return s
}

func (s *RedeemStakingProductService) Do(ctx context.Context, opts ...RequestOption) (*RedeemStakingProductResponse, error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/staking/redeem",
		secType:  secTypeSigned,
	}
	m := params{
		"product":   s.product,
		"productId": s.productId,
	}
	if s.positionId != nil {
		m["positionId"] = *s.positionId
	}
	if s.amount != nil {
		m["amount"] = *s.amount
	}
	r.setParams(m)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res := new(RedeemStakingProductResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type RedeemStakingProductResponse struct {
	Success bool `json:"success"`
}

type StakingProductQuotaService struct {
	c         *Client
	product   StakingProduct
	productId string
}

func (s *StakingProductQuotaService) Product(product StakingProduct) *StakingProductQuotaService {
	s.product = product
	return s
}

func (s *StakingProductQuotaService) ProductId(productId string) *StakingProductQuotaService {
	s.productId = productId
	return s
}

func (s *StakingProductQuotaService) Do(ctx context.Context) (*StakingProductPersonalQuota, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/staking/personalLeftQuota",
		secType:  secTypeSigned,
	}
	r.setParam("product", s.product)
	r.setParam("productId", s.productId)
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := new(StakingProductPersonalQuota)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type StakingProductPersonalQuota struct {
	LeftPersonalQuota string `json:"leftPersonalQuota"`
}

type SetAutoStakingService struct {
	c          *Client
	product    StakingProduct
	positionId int64
	renewable  bool
}

func (s *SetAutoStakingService) Product(product StakingProduct) *SetAutoStakingService {
	s.product = product
	return s
}

func (s *SetAutoStakingService) PositionId(positionId int64) *SetAutoStakingService {
	s.positionId = positionId
	return s
}

func (s *SetAutoStakingService) Renewable(renewable bool) *SetAutoStakingService {
	s.renewable = renewable
	return s
}

func (s *SetAutoStakingService) Do(ctx context.Context, opts ...RequestOption) (*SetAutoStakingResponse, error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/staking/setAutoStaking",
		secType:  secTypeSigned,
	}
	m := params{
		"product":    s.product,
		"positionId": s.positionId,
		"renewable":  s.renewable,
	}
	r.setParams(m)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res := new(SetAutoStakingResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type SetAutoStakingResponse struct {
	Success bool `json:"success"`
}

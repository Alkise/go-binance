package binance

import (
	"context"
)

type (
	CreateBrokerSubAccountService struct {
		c   *Client
		tag *string
	}
	CreateBrokerSubAccountResponse struct {
		SubAccountId string `json:"subaccountId"`
		Email        string `json:"email"`
		Tag          string `json:"tag"`
	}
	CreateBrokerSubAccountApiKeyService struct {
		c            *Client
		subAccountId string
		canTrade     bool
		marginTrade  *bool
		futuresTrade *bool
	}
	CreateBrokerSubAccountApiKeyResponse struct {
		SubAccountId string `json:"subaccountId"`
		ApiKey       string `json:"apiKey"`
		SecretKey    string `json:"secretKey"`
		CanTrade     bool   `json:"canTrade"`
		MarginTrade  bool   `json:"marginTrade"`
		FuturesTrade bool   `json:"futuresTrade"`
	}
	DeleteBrokerSubAccountApiKeyService struct {
		c                *Client
		subAccountId     string
		subAccountApiKey string
	}
	ListBrokerSubAccountsService struct {
		c            *Client
		subAccountId *string
		page         *int64
		size         *int64
	}
	BrokerSubAccount struct {
		SubAccountId          string `json:"subaccountId"`
		Email                 string `json:"email"`
		Tag                   string `json:"tag"`
		MakerCommission       string `json:"makerCommission"`
		TakerCommission       string `json:"takerCommission"`
		MarginMakerCommission string `json:"marginMakerCommission"`
		MarginTakerCommission string `json:"marginTakerCommission"`
		CreateTime            int64  `json:"createTime"`
	}
	EnableMarginForBrokerSubAccountService struct {
		c            *Client
		subAccountId string
		margin       bool
	}
	EnableFuturesForBrokerSubAccountService struct {
		c            *Client
		subAccountId string
		futures      bool
	}
	EnableMarginForBrokerSubAccountResponse struct {
		SubAccountId string `json:"subaccountId"`
		EnableMargin bool   `json:"enableMargin"`
		UpdateTime   int64  `json:"updateTime"`
	}
	EnableFuturesForBrokerSubAccountResponse struct {
		SubAccountId  string `json:"subaccountId"`
		EnableFutures bool   `json:"enableFutures"`
		UpdateTime    int64  `json:"updateTime"`
	}
	AddIPRestrictionForBrokerSubAccountService struct {
		c                *Client
		subAccountId     string
		subAccountApiKey string
		ipAddress        string
	}
	AddIPRestrictionForBrokerSubAccountResponse struct {
		SubAccountId string `json:"subaccountId"`
		Apikey       string `json:"apikey"`
		Ip           string `json:"ip"`
		UpdateTime   int64  `json:"updateTime"`
	}
	IPRestrictionForBrokerSubAccountService struct {
		c                *Client
		subAccountId     string
		subAccountApiKey string
		status           string
		ipAddress        string
	}
	IPRestrictionForBrokerSubAccountResponse struct {
		SubAccountId string   `json:"subaccountId"`
		IpRestrict   string   `json:"ipRestrict"`
		Apikey       string   `json:"apikey"`
		IpList       []string `json:"ipList"`
		UpdateTime   int64    `json:"updateTime"`
	}
	UniversalTransferService struct {
		c               *Client
		fromAccountType string
		toAccountType   string
		asset           string
		amount          float64
		fromId          *string
		toId            *string
		clientTranId    *string
	}
	UniversalTransferResponse struct {
		TxnId        int64  `json:"txnId"`
		ClientTranId string `json:"clientTranId"`
	}
	UniversalTransferHistoryService struct {
		c             *Client
		fromId        *string
		toId          *string
		clientTranId  *string
		startTime     *int64
		endTime       *int64
		page          *int32
		limit         *int32
		showAllStatus *bool
	}
	UniversalTransfer struct {
		FromId          string `json:"fromId,omitempty"`
		ToId            string `json:"toId,omitempty"`
		Asset           string `json:"asset"`
		Qty             string `json:"qty"`
		Time            int64  `json:"time"`
		Status          string `json:"status"`
		TxnId           int64  `json:"txnId"`
		ClientTranId    string `json:"clientTranId"`
		FromAccountType string `json:"fromAccountType"`
		ToAccountType   string `json:"toAccountType"`
	}
	CommissionRebateRecentRecordService struct {
		c            *Client
		subAccountId *string
		startTime    *int64
		endTime      *int64
		page         *int64
		size         *int64
	}
	FuturesCommissionRebateRecentRecordService struct {
		c           *Client
		futuresType uint8
		startTime   int64
		endTime     int64
		page        *int64
		size        *int64
	}
	BrokerRebateRecord struct {
		SubAccountId string `json:"subAccountId"`
		Income       string `json:"income"`
		Asset        string `json:"asset"`
		Symbol       string `json:"symbol"`
		TradeId      uint64 `json:"tradeId"`
		Time         uint64 `json:"time"`
		Status       uint8  `json:"status"` // 0: Pending, 1: Successful, 2: Failed
	}
	BNBBurnStatusService struct {
		c            *Client
		subAccountId string
	}
	BNBBurnStatus struct {
		SubAccountId    uint64 `json:"subAccountId"`
		SpotBNBBurn     bool   `json:"spotBNBBurn"`
		InterestBNBBurn bool   `json:"interestBNBBurn"`
	}
	SpotBNBBurnService struct {
		c            *Client
		subAccountId string
		spotBNBBurn  bool
	}
	SpotBNBBurnStatus struct {
		SubAccountId uint64 `json:"subAccountId"`
		SpotBNBBurn  bool   `json:"spotBNBBurn"`
	}
	MarginInterestBNBBurnService struct {
		c               *Client
		subAccountId    string
		interestBNBBurn bool
	}
	MarginInterestBNBBurnStatus struct {
		SubAccountId    uint64 `json:"subAccountId"`
		InterestBNBBurn bool   `json:"interestBNBBurn"`
	}
)

func (s *CreateBrokerSubAccountService) Tag(tag string) *CreateBrokerSubAccountService {
	s.tag = &tag
	return s
}

func (s *CreateBrokerSubAccountService) createBrokerSubAccount(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, err error) {
	r := &request{
		method:   "POST",
		endpoint: endpoint,
		secType:  secTypeSigned,
	}
	m := params{}
	if s.tag != nil {
		m["tag"] = *s.tag
	}
	r.setFormParams(m)
	data, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

func (s *CreateBrokerSubAccountService) Do(ctx context.Context, opts ...RequestOption) (res *CreateBrokerSubAccountResponse, err error) {
	data, err := s.createBrokerSubAccount(ctx, "/sapi/v1/broker/subAccount", opts...)
	if err != nil {
		return nil, err
	}
	res = new(CreateBrokerSubAccountResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *CreateBrokerSubAccountApiKeyService) SubAccountId(subAccountId string) *CreateBrokerSubAccountApiKeyService {
	s.subAccountId = subAccountId
	return s
}

func (s *CreateBrokerSubAccountApiKeyService) CanTrade(canTrade bool) *CreateBrokerSubAccountApiKeyService {
	s.canTrade = canTrade
	return s
}

func (s *CreateBrokerSubAccountApiKeyService) MarginTrade(marginTrade bool) *CreateBrokerSubAccountApiKeyService {
	s.marginTrade = &marginTrade
	return s
}

func (s *CreateBrokerSubAccountApiKeyService) FuturesTrade(futuresTrade bool) *CreateBrokerSubAccountApiKeyService {
	s.futuresTrade = &futuresTrade
	return s
}

func (s *CreateBrokerSubAccountApiKeyService) createBrokerSubAccountApiKey(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, err error) {
	r := &request{
		method:   "POST",
		endpoint: endpoint,
		secType:  secTypeSigned,
	}
	m := params{
		"subAccountId": s.subAccountId,
		"canTrade":     s.canTrade,
	}
	if s.marginTrade != nil {
		m["marginTrade"] = *s.marginTrade
	}
	if s.futuresTrade != nil {
		m["futuresTrade"] = *s.futuresTrade
	}
	r.setFormParams(m)
	data, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

func (s *CreateBrokerSubAccountApiKeyService) Do(ctx context.Context, opts ...RequestOption) (res *CreateBrokerSubAccountApiKeyResponse, err error) {
	data, err := s.createBrokerSubAccountApiKey(ctx, "/sapi/v1/broker/subAccountApi", opts...)
	if err != nil {
		return nil, err
	}
	res = new(CreateBrokerSubAccountApiKeyResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *DeleteBrokerSubAccountApiKeyService) SubAccountId(subAccountId string) *DeleteBrokerSubAccountApiKeyService {
	s.subAccountId = subAccountId
	return s
}

func (s *DeleteBrokerSubAccountApiKeyService) SubAccountApiKey(subAccountApiKey string) *DeleteBrokerSubAccountApiKeyService {
	s.subAccountApiKey = subAccountApiKey
	return s
}

func (s *DeleteBrokerSubAccountApiKeyService) deleteSubAccountApiKey(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, err error) {
	r := &request{
		method:   "DELETE",
		endpoint: endpoint,
		secType:  secTypeSigned,
	}
	m := params{
		"subAccountId":     s.subAccountId,
		"subAccountApiKey": s.subAccountApiKey,
	}
	r.setFormParams(m)
	data, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

func (s *DeleteBrokerSubAccountApiKeyService) Do(ctx context.Context, opts ...RequestOption) (err error) {
	_, err = s.deleteSubAccountApiKey(ctx, "/sapi/v1/broker/subAccountApi", opts...)
	return err
}

func (s *ListBrokerSubAccountsService) SubAccountId(subAccountId string) *ListBrokerSubAccountsService {
	s.subAccountId = &subAccountId
	return s
}

func (s *ListBrokerSubAccountsService) Page(page int64) *ListBrokerSubAccountsService {
	s.page = &page
	return s
}

func (s *ListBrokerSubAccountsService) Size(size int64) *ListBrokerSubAccountsService {
	s.size = &size
	return s
}

func (s *ListBrokerSubAccountsService) Do(ctx context.Context, opts ...RequestOption) (res []*BrokerSubAccount, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/broker/subAccount",
		secType:  secTypeSigned,
	}
	if s.subAccountId != nil {
		r.setParam("subAccountId", *s.subAccountId)
	}
	if s.page != nil {
		r.setParam("page", *s.page)
	}
	if s.size != nil {
		r.setParam("size", *s.size)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*BrokerSubAccount{}, err
	}
	res = make([]*BrokerSubAccount, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*BrokerSubAccount{}, err
	}
	return res, nil
}

func (s *EnableMarginForBrokerSubAccountService) SubAccountId(subAccountId string) *EnableMarginForBrokerSubAccountService {
	s.subAccountId = subAccountId
	return s
}

func (s *EnableMarginForBrokerSubAccountService) Margin(margin bool) *EnableMarginForBrokerSubAccountService {
	s.margin = margin
	return s
}

func (s *EnableMarginForBrokerSubAccountService) Do(ctx context.Context, opts ...RequestOption) (res *EnableMarginForBrokerSubAccountResponse, err error) {
	r := &request{
		method:   "POST",
		endpoint: "/sapi/v1/broker/subAccount/margin",
		secType:  secTypeSigned,
	}
	m := params{
		"subAccountId": s.subAccountId,
		"margin":       s.margin,
	}
	r.setFormParams(m)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(EnableMarginForBrokerSubAccountResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *EnableFuturesForBrokerSubAccountService) SubAccountId(subAccountId string) *EnableFuturesForBrokerSubAccountService {
	s.subAccountId = subAccountId
	return s
}

func (s *EnableFuturesForBrokerSubAccountService) Futures(futures bool) *EnableFuturesForBrokerSubAccountService {
	s.futures = futures
	return s
}

func (s *EnableFuturesForBrokerSubAccountService) Do(ctx context.Context, opts ...RequestOption) (res *EnableFuturesForBrokerSubAccountResponse, err error) {
	r := &request{
		method:   "POST",
		endpoint: "/sapi/v1/broker/subAccount/futures",
		secType:  secTypeSigned,
	}
	m := params{
		"subAccountId": s.subAccountId,
		"futures":      s.futures,
	}
	r.setFormParams(m)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(EnableFuturesForBrokerSubAccountResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *AddIPRestrictionForBrokerSubAccountService) SubAccountId(subAccountId string) *AddIPRestrictionForBrokerSubAccountService {
	s.subAccountId = subAccountId
	return s
}

func (s *AddIPRestrictionForBrokerSubAccountService) SubAccountApiKey(subAccountApiKey string) *AddIPRestrictionForBrokerSubAccountService {
	s.subAccountApiKey = subAccountApiKey
	return s
}

func (s *AddIPRestrictionForBrokerSubAccountService) IPAddress(ipAddress string) *AddIPRestrictionForBrokerSubAccountService {
	s.ipAddress = ipAddress
	return s
}

func (s *AddIPRestrictionForBrokerSubAccountService) Do(ctx context.Context, opts ...RequestOption) (res *AddIPRestrictionForBrokerSubAccountResponse, err error) {
	r := &request{
		method:   "POST",
		endpoint: "/sapi/v1/broker/subAccountApi/ipRestriction/ipList",
		secType:  secTypeSigned,
	}
	m := params{
		"subAccountId":     s.subAccountId,
		"subAccountApiKey": s.subAccountApiKey,
		"ipAddress":        s.ipAddress,
	}
	r.setFormParams(m)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(AddIPRestrictionForBrokerSubAccountResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *IPRestrictionForBrokerSubAccountService) SubAccountId(subAccountId string) *IPRestrictionForBrokerSubAccountService {
	s.subAccountId = subAccountId
	return s
}

func (s *IPRestrictionForBrokerSubAccountService) SubAccountApiKey(subAccountApiKey string) *IPRestrictionForBrokerSubAccountService {
	s.subAccountApiKey = subAccountApiKey
	return s
}

func (s *IPRestrictionForBrokerSubAccountService) Status(status string) *IPRestrictionForBrokerSubAccountService {
	s.status = status
	return s
}

func (s *IPRestrictionForBrokerSubAccountService) IPAddress(ipAddress string) *IPRestrictionForBrokerSubAccountService {
	s.ipAddress = ipAddress
	return s
}

func (s *IPRestrictionForBrokerSubAccountService) Do(ctx context.Context, opts ...RequestOption) (res *IPRestrictionForBrokerSubAccountResponse, err error) {
	r := &request{
		method:   "POST",
		endpoint: "/sapi/v1/broker/subAccountApi/ipRestriction",
		secType:  secTypeSigned,
	}
	m := params{
		"subAccountId":     s.subAccountId,
		"subAccountApiKey": s.subAccountApiKey,
		"status":           s.status,
	}
	if len(ipAddress) != 0 {
		m["ipAddress"] = s.ipAddress
	}
	r.setFormParams(m)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(IPRestrictionForBrokerSubAccountResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *UniversalTransferService) FromAccountType(fromAccountType string) *UniversalTransferService {
	s.fromAccountType = fromAccountType
	return s
}
func (s *UniversalTransferService) ToAccountType(toAccountType string) *UniversalTransferService {
	s.toAccountType = toAccountType
	return s
}
func (s *UniversalTransferService) Asset(asset string) *UniversalTransferService {
	s.asset = asset
	return s
}
func (s *UniversalTransferService) Amount(amount float64) *UniversalTransferService {
	s.amount = amount
	return s
}
func (s *UniversalTransferService) FromId(fromId string) *UniversalTransferService {
	s.fromId = &fromId
	return s
}
func (s *UniversalTransferService) ToId(toId string) *UniversalTransferService {
	s.toId = &toId
	return s
}
func (s *UniversalTransferService) ClientTranId(clientTranId string) *UniversalTransferService {
	s.clientTranId = &clientTranId
	return s
}

func (s *UniversalTransferService) Do(ctx context.Context, opts ...RequestOption) (res *UniversalTransferResponse, err error) {
	r := &request{
		method:   "POST",
		endpoint: "/sapi/v1/broker/universalTransfer",
		secType:  secTypeSigned,
	}
	m := params{
		"fromAccountType": s.fromAccountType,
		"toAccountType":   s.toAccountType,
		"asset":           s.asset,
		"amount":          s.amount,
	}
	if s.clientTranId != nil {
		m["clientTranId"] = *s.clientTranId
	}
	if s.fromId != nil {
		m["fromId"] = *s.fromId
	}
	if s.toId != nil {
		m["toId"] = *s.toId
	}

	s.c.debug("/sapi/v1/broker/universalTransfer: %q", m)

	r.setFormParams(m)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(UniversalTransferResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *UniversalTransferHistoryService) FromId(fromId string) *UniversalTransferHistoryService {
	s.fromId = &fromId
	return s
}
func (s *UniversalTransferHistoryService) ToId(toId string) *UniversalTransferHistoryService {
	s.toId = &toId
	return s
}
func (s *UniversalTransferHistoryService) ClientTranId(clientTranId string) *UniversalTransferHistoryService {
	s.clientTranId = &clientTranId
	return s
}
func (s *UniversalTransferHistoryService) StartTime(startTime int64) *UniversalTransferHistoryService {
	s.startTime = &startTime
	return s
}
func (s *UniversalTransferHistoryService) EndTime(endTime int64) *UniversalTransferHistoryService {
	s.endTime = &endTime
	return s
}
func (s *UniversalTransferHistoryService) Page(page int32) *UniversalTransferHistoryService {
	s.page = &page
	return s
}
func (s *UniversalTransferHistoryService) Limit(limit int32) *UniversalTransferHistoryService {
	s.limit = &limit
	return s
}
func (s *UniversalTransferHistoryService) ShowAllStatus(showAllStatus bool) *UniversalTransferHistoryService {
	s.showAllStatus = &showAllStatus
	return s
}

func (s *UniversalTransferHistoryService) Do(ctx context.Context, opts ...RequestOption) (res []*UniversalTransfer, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/broker/universalTransfer",
		secType:  secTypeSigned,
	}
	if s.fromId != nil {
		r.setParam("fromId", *s.fromId)
	}
	if s.toId != nil {
		r.setParam("toId", *s.toId)
	}
	if s.clientTranId != nil {
		r.setParam("clientTranId", *s.clientTranId)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.showAllStatus != nil {
		r.setParam("showAllStatus", *s.showAllStatus)
	}
	if s.page != nil {
		r.setParam("page", *s.page)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*UniversalTransfer{}, err
	}
	res = make([]*UniversalTransfer, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*UniversalTransfer{}, err
	}
	return res, nil
}

func (s *CommissionRebateRecentRecordService) SubAccountId(subAccountId string) *CommissionRebateRecentRecordService {
	s.subAccountId = &subAccountId
	return s
}

func (s *CommissionRebateRecentRecordService) StartTime(startTime int64) *CommissionRebateRecentRecordService {
	s.startTime = &startTime
	return s
}

func (s *CommissionRebateRecentRecordService) EndTime(endTime int64) *CommissionRebateRecentRecordService {
	s.endTime = &endTime
	return s
}

func (s *CommissionRebateRecentRecordService) Page(page int64) *CommissionRebateRecentRecordService {
	s.page = &page
	return s
}

func (s *CommissionRebateRecentRecordService) Size(size int64) *CommissionRebateRecentRecordService {
	s.size = &size
	return s
}

func (s *CommissionRebateRecentRecordService) Do(ctx context.Context, opts ...RequestOption) (res []*BrokerRebateRecord, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/broker/rebate/recentRecord",
		secType:  secTypeSigned,
	}

	if s.subAccountId != nil {
		r.setParam("subAccountId", *s.subAccountId)
	}

	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}

	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}

	if s.page != nil {
		r.setParam("page", *s.page)
	}

	if s.size != nil {
		r.setParam("size", *s.size)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*BrokerRebateRecord{}, err
	}

	res = make([]*BrokerRebateRecord, 0)

	if err = json.Unmarshal(data, &res); err != nil {
		return []*BrokerRebateRecord{}, err
	}

	return res, nil
}

func (s *FuturesCommissionRebateRecentRecordService) FuturesType(futuresType uint8) *FuturesCommissionRebateRecentRecordService {
	s.futuresType = futuresType
	return s

}

func (s *FuturesCommissionRebateRecentRecordService) StartTime(startTime int64) *FuturesCommissionRebateRecentRecordService {
	s.startTime = startTime
	return s
}

func (s *FuturesCommissionRebateRecentRecordService) EndTime(endTime int64) *FuturesCommissionRebateRecentRecordService {
	s.endTime = endTime
	return s
}

func (s *FuturesCommissionRebateRecentRecordService) Page(page int64) *FuturesCommissionRebateRecentRecordService {
	s.page = &page
	return s
}

func (s *FuturesCommissionRebateRecentRecordService) Size(size int64) *FuturesCommissionRebateRecentRecordService {
	s.size = &size
	return s
}

func (s *FuturesCommissionRebateRecentRecordService) Do(ctx context.Context, opts ...RequestOption) (res []*BrokerRebateRecord, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/broker/rebate/futures/recentRecord",
		secType:  secTypeSigned,
	}

	r.setParams(params{
		"futuresType": s.futuresType,
		"startTime":   s.startTime,
		"endTime":     s.endTime,
	})

	if s.page != nil {
		r.setParam("page", *s.page)
	}

	if s.size != nil {
		r.setParam("size", *s.size)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*BrokerRebateRecord{}, err
	}

	res = make([]*BrokerRebateRecord, 0)

	if err = json.Unmarshal(data, &res); err != nil {
		return []*BrokerRebateRecord{}, err
	}

	return res, nil
}

func (s *BNBBurnStatusService) SubAccountId(subAccountId string) *BNBBurnStatusService {
	s.subAccountId = subAccountId
	return s
}

func (s *BNBBurnStatusService) Do(ctx context.Context, opts ...RequestOption) (res *BNBBurnStatus, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/broker/subAccount/bnbBurn/status",
		secType:  secTypeSigned,
	}

	r.setParams(params{
		"subAccountId": s.subAccountId,
	})

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}

	res = new(BNBBurnStatus)

	err = json.Unmarshal(data, res)

	return
}

func (s *SpotBNBBurnService) SubAccountId(subAccountId string) *SpotBNBBurnService {
	s.subAccountId = subAccountId
	return s
}

func (s *SpotBNBBurnService) SpotBNBBurn(spotBNBBurn bool) *SpotBNBBurnService {
	s.spotBNBBurn = spotBNBBurn
	return s
}

func (s *SpotBNBBurnService) Do(ctx context.Context, opts ...RequestOption) (res *SpotBNBBurnStatus, err error) {
	r := &request{
		method:   "POST",
		endpoint: "/sapi/v1/broker/subAccount/bnbBurn/spot",
		secType:  secTypeSigned,
	}
	m := params{
		"subAccountId": s.subAccountId,
		"spotBNBBurn":  s.spotBNBBurn,
	}

	s.c.debug("/sapi/v1/broker/subAccount/bnbBurn/spot: %q", m)

	r.setFormParams(m)

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res = new(SpotBNBBurnStatus)

	err = json.Unmarshal(data, res)

	return
}

func (s *MarginInterestBNBBurnService) SubAccountId(subAccountId string) *MarginInterestBNBBurnService {
	s.subAccountId = subAccountId
	return s
}

func (s *MarginInterestBNBBurnService) MarginInterestBNBBurn(interestBNBBurn bool) *MarginInterestBNBBurnService {
	s.interestBNBBurn = interestBNBBurn
	return s
}

func (s *MarginInterestBNBBurnService) Do(ctx context.Context, opts ...RequestOption) (res *MarginInterestBNBBurnStatus, err error) {
	r := &request{
		method:   "POST",
		endpoint: "/sapi/v1/broker/subAccount/bnbBurn/marginInterest",
		secType:  secTypeSigned,
	}
	m := params{
		"subAccountId":    s.subAccountId,
		"interestBNBBurn": s.interestBNBBurn,
	}

	s.c.debug("/sapi/v1/broker/subAccount/bnbBurn/marginInterest: %q", m)

	r.setFormParams(m)

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res = new(MarginInterestBNBBurnStatus)

	err = json.Unmarshal(data, res)

	return
}

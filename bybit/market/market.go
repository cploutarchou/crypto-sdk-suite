package market

import (
	"fmt"
	"github.com/cploutarchou/crypto-sdk-suite/bybit/client"
)

type Market interface {
	ServerTime(params *client.Params) (*ServerTimeResponse, error)
	Kline(params *client.Params) (*KlineResponse, error)
	Announcement(params *client.Params) (*AnnouncementsResponse, error)
	MarkPriceKline(params *client.Params) (*KlineResponse, error)
	IndexPriceKline(params *client.Params) (*KlineResponse, error)
	PremiumIndexKline(params *client.Params) (*KlineResponse, error)
	OrderBook(params *client.Params) (*OrderBook, error)
	InstrumentsInfo(params *client.Params) (*InstrumentsInfoResponse, error)
	Tickers(params *client.Params) (*TickerResponse, error)
	FundingHistory(params *client.Params) (*FundingRateHistory, error)
	RiskLimit(params *client.Params) (*RiskLimit, error)
	OpenInterest(params *client.Params) (*OpenHistory, error)
	Insurance(params *client.Params) (*Insurance, error)
	RecentTrade(params *client.Params) (*ResendTrade, error)
	DeliveryPrice(params *client.Params) (*DeliveryPrice, error)
	HistoricalVolatility(params *client.Params) (*HistoricalVolatility, error)
}

type marketImpl struct {
	c *client.Client
}

func New(c *client.Client) Market {
	return &marketImpl{c}
}

func (m *marketImpl) ServerTime(params *client.Params) (*ServerTimeResponse, error) {
	res, err := m.c.Get(fmt.Sprintf("/%s/market/time", client.ApiVersion), *params)
	if err != nil {
		return nil, err
	}
	var serverTime ServerTimeResponse
	if err := res.Unmarshal(&serverTime); err != nil {
		return nil, err
	}
	return &serverTime, nil
}
func (m *marketImpl) Kline(params *client.Params) (*KlineResponse, error) {
	res, err := m.c.Get(fmt.Sprintf("/%s/market/kline", client.ApiVersion), *params)

	if err != nil {
		return nil, err
	}
	var kline KlineResponse
	err = res.Unmarshal(&kline)
	if err != nil {
		return nil, err
	}
	return &kline, nil
}

func (m *marketImpl) Announcement(params *client.Params) (*AnnouncementsResponse, error) {
	res, err := m.c.Get(fmt.Sprintf("/%s/announcements/index", client.ApiVersion), *params)
	if err != nil {
		return nil, err
	}
	var announcement AnnouncementsResponse
	err = res.Unmarshal(&announcement)
	if err != nil {
		return nil, err
	}
	return &announcement, nil
}

func (m *marketImpl) MarkPriceKline(params *client.Params) (*KlineResponse, error) {
	res, err := m.c.Get(fmt.Sprintf("/%s/market/mark-price-kline", client.ApiVersion), *params)
	if err != nil {
		return nil, err
	}
	var markPriceKline KlineResponse
	err = res.Unmarshal(&markPriceKline)
	if err != nil {
		return nil, err
	}
	return &markPriceKline, nil
}

func (m *marketImpl) IndexPriceKline(params *client.Params) (*KlineResponse, error) {
	res, err := m.c.Get(fmt.Sprintf("/%s/market/index-price-kline", client.ApiVersion), *params)
	if err != nil {
		return nil, err
	}
	var indexPriceKline KlineResponse
	err = res.Unmarshal(&indexPriceKline)
	if err != nil {
		return nil, err
	}
	return &indexPriceKline, nil
}

func (m *marketImpl) PremiumIndexKline(params *client.Params) (*KlineResponse, error) {
	res, err := m.c.Get(fmt.Sprintf("/%s/market/premium-index-kline", client.ApiVersion), *params)
	if err != nil {
		return nil, err
	}
	var premiumIndexKline KlineResponse
	err = res.Unmarshal(&premiumIndexKline)
	if err != nil {
		return nil, err
	}
	return &premiumIndexKline, nil
}

func (m *marketImpl) OrderBook(params *client.Params) (*OrderBook, error) {
	res, err := m.c.Get(fmt.Sprintf("/%s/market/orderbook", client.ApiVersion), *params)
	if err != nil {
		return nil, err
	}
	var orderBook OrderBook
	err = res.Unmarshal(&orderBook)
	if err != nil {
		return nil, err
	}
	return &orderBook, nil
}

func (m *marketImpl) InstrumentsInfo(params *client.Params) (*InstrumentsInfoResponse, error) {
	res, err := m.c.Get(fmt.Sprintf("/%s/market/instruments-info", client.ApiVersion), *params)
	if err != nil {
		return nil, err
	}
	var instrumentsInfo InstrumentsInfoResponse
	err = res.Unmarshal(&instrumentsInfo)
	if err != nil {
		return nil, err
	}
	return &instrumentsInfo, nil
}

func (m *marketImpl) Tickers(params *client.Params) (*TickerResponse, error) {
	res, err := m.c.Get(fmt.Sprintf("/%s/market/tickers", client.ApiVersion), *params)
	if err != nil {
		return nil, err
	}
	var tickers TickerResponse
	err = res.Unmarshal(&tickers)
	if err != nil {
		return nil, err
	}
	return &tickers, nil
}

func (m *marketImpl) FundingHistory(params *client.Params) (*FundingRateHistory, error) {
	res, err := m.c.Get(fmt.Sprintf("/%s/market/funding/history", client.ApiVersion), *params)
	if err != nil {
		return nil, err
	}
	var fundingHistory FundingRateHistory
	err = res.Unmarshal(&fundingHistory)
	if err != nil {
		return nil, err
	}
	return &fundingHistory, nil
}

func (m *marketImpl) RiskLimit(params *client.Params) (*RiskLimit, error) {
	res, err := m.c.Get(fmt.Sprintf("/%s/market/insurance", client.ApiVersion), *params)
	if err != nil {
		return nil, err
	}
	var riskLimit RiskLimit
	err = res.Unmarshal(&riskLimit)
	if err != nil {
		return nil, err
	}
	return &riskLimit, nil
}

func (m *marketImpl) OpenInterest(params *client.Params) (*OpenHistory, error) {
	res, err := m.c.Get(fmt.Sprintf("/%s/market/open-interest", client.ApiVersion), *params)
	if err != nil {
		return nil, err
	}
	var openInterest OpenHistory
	err = res.Unmarshal(&openInterest)
	if err != nil {
		return nil, err
	}
	return &openInterest, nil
}

func (m *marketImpl) Insurance(params *client.Params) (*Insurance, error) {
	res, err := m.c.Get(fmt.Sprintf("/%s/market/insurance", client.ApiVersion), *params)
	if err != nil {
		return nil, err
	}
	var insurance Insurance
	err = res.Unmarshal(&insurance)
	if err != nil {
		return nil, err
	}
	return &insurance, nil
}

func (m *marketImpl) RecentTrade(params *client.Params) (*ResendTrade, error) {
	res, err := m.c.Get(fmt.Sprintf("/%s/market/trading-records", client.ApiVersion), *params)
	if err != nil {
		return nil, err
	}
	var recentTrade ResendTrade
	err = res.Unmarshal(&recentTrade)
	if err != nil {
		return nil, err
	}
	return &recentTrade, nil
}

func (m *marketImpl) DeliveryPrice(params *client.Params) (*DeliveryPrice, error) {
	res, err := m.c.Get(fmt.Sprintf("/%s/public/delivery-price", client.ApiVersion), *params)
	if err != nil {
		return nil, err
	}
	var deliveryPrice DeliveryPrice
	err = res.Unmarshal(&deliveryPrice)
	if err != nil {
		return nil, err
	}
	return &deliveryPrice, nil
}

func (m *marketImpl) HistoricalVolatility(params *client.Params) (*HistoricalVolatility, error) {
	res, err := m.c.Get(fmt.Sprintf("/%s/public/historical-volatility", client.ApiVersion), *params)
	if err != nil {
		return nil, err
	}
	var historicalVolatility HistoricalVolatility
	err = res.Unmarshal(&historicalVolatility)
	if err != nil {
		return nil, err
	}
	return &historicalVolatility, nil
}

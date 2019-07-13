package iex

// Quote ...
type Quote struct {
	Symbol                string  `json:"symbol"`
	CompanyName           string  `json:"companyName"`
	CalculationPrice      string  `json:"calculationPrice"`
	Open                  float64 `json:"open"`
	OpenTime              uint64  `json:"openTime"`
	Close                 float64 `json:"close"`
	CloseTime             uint64  `json:"closeTime"`
	High                  float64 `json:"high"`
	Low                   float64 `json:"low"`
	LatestPrice           float64 `json:"latestPrice"`
	LatestSource          string  `json:"latestSource"`
	LatestTime            string  `json:"latestTime"`
	LatestUpdate          uint64  `json:"latestUpdate"`
	LatestVolume          uint64  `json:"latestVolume"`
	IEXRealtimePrice      float64 `json:"iexRealtimePrice"`
	IEXRealtimeSize       uint64  `json:"iexRealtimeSize"`
	IEXLastUpdated        uint64  `json:"iexLastUpdated"`
	DelayedPrice          float64 `json:"delayedPrice"`
	DelayedPriceTime      uint64  `json:"delayedPriceTime"`
	ExtendedPrice         float64 `json:"extendedPrice"`
	ExtendedChange        float64 `json:"extendedChange"`
	ExtendedChangePercent float64 `json:"extendedChangePercent"`
	ExtendedPriceTime     uint64  `json:"extendedPriceTime"`
	PreviousClose         float64 `json:"previousClose"`
	Change                float64 `json:"change"`
	ChangePercent         float64 `json:"changePercent"`
	IEXMarketPercent      float64 `json:"iexMarketPercent"`
	IEXVolume             uint64  `json:"iexVolume"`
	AvgTotalVolume        uint64  `json:"avgTotalVolume"`
	IEXBidPrice           float64 `json:"iexBidPrice"`
	IEXBidSize            uint64  `json:"iexBidSize"`
	IEXAskPrice           float64 `json:"iexAskPrice"`
	IEXAskSize            uint64  `json:"iexAskSize"`
	MarketCap             uint64  `json:"marketCap"`
	Week52High            float64 `json:"week52High"`
	Week52Low             float64 `json:"week52Low"`
	YTDChange             float64 `json:"ytdChange"`
}

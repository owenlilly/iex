package iex

// IntradayPrices ...
type IntradayPrices []struct {
	Date                 string  `json:"date"`
	Minute               string  `json:"minute"`
	Label                string  `json:"label"`
	MarktOpen            float64 `json:"marktOpen"`
	MarketClose          float64 `json:"marketClose"`
	MarktHigh            float64 `json:"marktHigh"`
	MarketLow            float64 `json:"marketLow"`
	MarketAverage        float64 `json:"marketAverage"`
	MarketVolume         uint64  `json:"marketVolume"`
	MarketNotional       float64 `json:"marketNotional"`
	MarketNumberOfTrades uint64  `json:"marketNumberOfTrades"`
	MarketChangeOverTime float64 `json:"marketChangeOverTime"`
	High                 float64 `json:"high"`
	Low                  float64 `json:"low"`
	Open                 float64 `json:"open"`
	Close                float64 `json:"close"`
	Average              float64 `json:"average"`
	Volume               uint64  `json:"volume"`
	Notional             float64 `json:"notional"`
	NumberOfTrades       uint64  `json:"numberOfTrades"`
	ChangeOverTime       float64 `json:"changeOverTime"`
}

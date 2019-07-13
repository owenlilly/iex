package iex

// OHLC ...
type OHLC []struct {
	Open struct {
		Price uint64 `json:"price"`
		Time  uint64 `json:"time"`
	} `json:"open"`

	Close struct {
		Price float64 `json:"price"`
		Time  uint64  `json:"time"`
	} `json:"close"`

	High float64 `json:"high"`
	Low  float64 `json:"low"`
}

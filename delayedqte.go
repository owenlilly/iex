package iex

// DelayedQuote returns the 15 minute delayed market quote.
type DelayedQuote struct {
	Symbol           string  `json:"symbol"`
	DelayedPrice     float64 `json:"delayedPrice"`
	DelayedSize      uint64  `json:"delayedSize"`
	DelayedPriceTime uint64  `json:"delayedPriceTime"`
	High             float64 `json:"high"`
	Low              float64 `json:"low"`
	TotalVolume      uint64  `json:"totalVolume"`
	ProcessedTime    uint64  `json:"processedTime"`
}

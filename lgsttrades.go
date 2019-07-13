package iex

// LgstTrades returns 15 minute delayed, last sale eligible trades.
type LgstTrades []struct {
	Price     float64 `json:"price"`
	Size      uint64  `json:"size"`
	Time      uint64  `json:"time"`
	TimeLabel string  `json:"timeLabel"`
	Venue     string  `json:"venue"`
	VenueName string  `json:"venueName"`
}

package iex

// VolByVenue ...
type VolByVenue []struct {
	Volume        uint64  `json:"volume"`
	Venue         string  `json:"venue"`
	VenueName     string  `json:"venueName"`
	Date          string  `json:"date"`
	MarketPercent float64 `json:"marketPercent"`
}

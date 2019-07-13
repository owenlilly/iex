package iex

// MarketVol ...
type MarketVol []struct {
	Mic           string  `json:"mic"`
	TapeID        string  `json:"tapeId"`
	VenueName     string  `json:"venueName"`
	Volume        uint64  `json:"volume"`
	TapeA         uint64  `json:"tapeA"`
	TapeB         uint64  `json:"tapeB"`
	TapeC         uint64  `json:"tapeC"`
	MarketPercent float64 `json:"marketPercent"`
	LastUpdated   uint64  `json:"lastUpdated"`
}

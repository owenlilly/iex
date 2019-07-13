package iex

// PriceTgt ...
type PriceTgt struct {
	Symbol             string  `json:"symbol"`
	UpdatedDate        string  `json:"updatedDate"`
	PriceTargetAverage float64 `json:"priceTargetAverage"`
	PriceTargetHigh    float64 `json:"priceTargetHigh"`
	PriceTargetLow     float64 `json:"priceTargetLow"`
	NumberOfAnalysts   uint64  `json:"numberOfAnalysts"`
}
